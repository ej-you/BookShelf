// Package app contains all internall app logic.
// Main function of this package is Run that run
// full application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	fiber "github.com/gofiber/fiber/v2"
	django "github.com/gofiber/template/django/v3"
	"gorm.io/gorm"

	"BookShelf/config"
	"BookShelf/internal/app/delivery/http"
	"BookShelf/internal/app/middleware"
	"BookShelf/internal/app/repo/sqlite"
	"BookShelf/internal/app/usecase"
	"BookShelf/internal/pkg/cookie"
	"BookShelf/internal/pkg/db"
	"BookShelf/internal/pkg/logger"
	"BookShelf/internal/pkg/validator"
)

var _ App = (*app)(nil)

// Client interface
type App interface {
	Run() error
}

// Fiber client
type app struct {
	cfg            *config.Config
	log            logger.Logger
	templateEngine fiber.Views
	valid          validator.Validator
	cookieBuilder  cookie.Builder
	dbStorage      *gorm.DB
}

// Client constructor
func New(cfg *config.Config) (App, error) {
	log := logger.NewLogger()
	templateEngine := django.New("./web/template", ".html")
	valid := validator.NewValidator()
	cookieBuilder := cookie.NewBuilder(cfg.App.AuthTokenTTL,
		cookie.WithPath(cfg.Cookie.Path),
		cookie.WithSecure(cfg.Cookie.Secure),
		cookie.WithHTTPOnly(cfg.Cookie.HTTPOnly),
		cookie.WithSameSite(cfg.Cookie.SameSite),
	)
	dbStorage, err := db.New(cfg.DB.DSN,
		db.WithLogger(log),
		db.WithWarnLogLevel(),
		db.WithIgnoreNotFound(),
		db.WithTranslateError(),
	)
	if err != nil {
		return nil, err
	}

	return &app{
		cfg:            cfg,
		log:            log,
		templateEngine: templateEngine,
		valid:          valid,
		cookieBuilder:  cookieBuilder,
		dbStorage:      dbStorage,
	}, nil
}

func (a *app) Run() error {
	// app init
	fiberApp := fiber.New(fiber.Config{
		AppName:      "BookShelf App",
		ErrorHandler: http.CustomErrorHandler,
		ServerHeader: "BookShelf",
		Views:        a.templateEngine,
	})

	// set up base middlewares
	fiberApp.Use(middleware.Logger())
	fiberApp.Use(middleware.Recover())
	fiberApp.Use(middleware.Compression())

	// set up static
	fiberApp.Static("/favicon.ico", "./web/static/img/favicon.ico")
	fiberApp.Static("/static", "./web/static")

	// create repo
	userRepoDB := sqlite.NewUserRepoDB(a.dbStorage)
	genreRepoDB := sqlite.NewGenreRepoDB(a.dbStorage)

	// create usecases
	userUC := usecase.NewUserUsecase(userRepoDB, a.cfg.AuthTokenSecret, a.cfg.AuthTokenTTL)
	genreUC := usecase.NewGenreUsecase(genreRepoDB)

	// register endpoints
	http.RegisterIndexEndpoints(fiberApp)
	http.RegisterUserEndpoints(fiberApp, userUC, a.valid, a.cookieBuilder)
	http.RegisterGenreEndpoints(fiberApp, genreUC, a.valid)

	// handle shutdown process signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	shutdownDone := make(chan struct{})
	// create gracefully shutdown task
	go func() {
		handledSignal := <-quit
		a.log.Printf("Get %q signal. Shutdown app...", handledSignal.String())
		// shutdown app
		fiberApp.ShutdownWithTimeout(a.cfg.App.KeepAliveTimeout)
		shutdownDone <- struct{}{}
	}()

	// start app
	if err := fiberApp.Listen(fmt.Sprintf(":%s", a.cfg.App.Port)); err != nil {
		return fmt.Errorf("start app: %w", err)
	}

	// wait for gracefully shutdown
	<-shutdownDone
	a.log.Print("App shutdown successfully!")
	return nil
}
