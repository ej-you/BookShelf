package commands

import (
	"context"
	"fmt"

	cli "github.com/urfave/cli/v3"

	"BookShelf/internal/pkg/migrate"
)

// Down command instance.
func NewDown(migrateManager migrate.Migrate) *cli.Command {
	return &cli.Command{
		Name:   "down",
		Usage:  "Rollback migrations (rollback all migrations if the flag -n is not specified)",
		Action: newDownAction(migrateManager),
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:      "number",
				Aliases:   []string{"n"},
				Value:     0,
				Usage:     "Set the number of migrations to rollback",
				Validator: positiveFlagValidator,
			},
		},
	}
}

// Handler for down command.
func newDownAction(migrateManager migrate.Migrate) cli.ActionFunc {
	return func(_ context.Context, cmd *cli.Command) error {
		step := cmd.Int("n")
		if step == 0 {
			fmt.Println("Rollback all migrations...")
			if err := migrateManager.Down(); err != nil {
				return err
			}
		} else {
			fmt.Printf("Rollback %d migrations... \n", step)
			if err := migrateManager.Step(-step); err != nil {
				return err
			}
		}
		fmt.Println("Successfully!")
		return nil
	}
}
