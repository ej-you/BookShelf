// Package cookie provides builder interface to create cookies.
package cookie

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

const _expiresToClear = -(24 * time.Hour)

var _ Builder = (*builder)(nil)

type Builder interface {
	CreateCookie(key, value string) *fiber.Cookie
	ClearCookie(key string) *fiber.Cookie
}

// Builder implementation.
type builder struct {
	expires  time.Duration
	path     string
	secure   bool
	httpOnly bool
	sameSite string
}

// Type for options for Builder initializing.
type Option func(*builder)

// Returns new Builder. Only required
// parameter is expires duration for cookies.
// Options can be set with "WithSmth" funcs.
func NewBuilder(expires time.Duration, options ...Option) Builder {
	builder := &builder{
		expires:  expires,
		path:     "/",
		secure:   false,
		httpOnly: false,
		sameSite: "",
	}

	// apply all options to customize Builder
	for _, opt := range options {
		opt(builder)
	}
	return builder
}

// Set path for cookie. Optional.
func WithPath(path string) Option {
	return func(b *builder) {
		b.path = path
	}
}

// Set secure for cookie. Optional.
func WithSecure(secure bool) Option {
	return func(b *builder) {
		b.secure = secure
	}
}

// Set httpOnly for cookie. Optional.
func WithHTTPOnly(httpOnly bool) Option {
	return func(b *builder) {
		b.httpOnly = httpOnly
	}
}

// Set sameSite for cookie. Optional.
// Supported values: Strict, Lax.
func WithSameSite(rawSameSite string) Option {
	sameSite := rawSameSite
	if rawSameSite != "Strict" && rawSameSite != "Lax" {
		sameSite = ""
	}
	return func(b *builder) {
		b.sameSite = sameSite
	}
}

// Create and return cookie with parameters from builder.
func (b builder) CreateCookie(key, value string) *fiber.Cookie {
	emptyCookie := b.emptyCookie()
	emptyCookie.Name = key
	emptyCookie.Value = value
	emptyCookie.Expires = emptyCookie.Expires.Add(b.expires)
	return emptyCookie
}

// Create and return clear cookie with parameters from builder.
func (b builder) ClearCookie(key string) *fiber.Cookie {
	emptyCookie := b.emptyCookie()
	emptyCookie.Name = key
	emptyCookie.Expires = emptyCookie.Expires.Add(_expiresToClear)
	return emptyCookie
}

// Create and return new empty cookie with builder parameters.
func (b builder) emptyCookie() *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "",
		Value:    "",
		Path:     b.path,
		HTTPOnly: b.httpOnly,
		Secure:   b.secure,
		SameSite: b.sameSite,
		Expires:  time.Now().UTC(),
	}
}
