package commands

import (
	"context"
	"fmt"

	cli "github.com/urfave/cli/v3"

	"BookShelf/internal/pkg/migrate"
)

// Up command instance.
func NewUp(migrateManager migrate.Migrate) *cli.Command {
	return &cli.Command{
		Name:   "up",
		Usage:  "Apply migrations (apply all migrations if the flag -n is not specified)",
		Action: newUpAction(migrateManager),
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:      "number",
				Aliases:   []string{"n"},
				Value:     0,
				Usage:     "Set the number of migrations to apply",
				Validator: positiveFlagValidator,
			},
		},
	}
}

// Handler for up command.
func newUpAction(migrateManager migrate.Migrate) cli.ActionFunc {
	return func(_ context.Context, cmd *cli.Command) error {
		step := cmd.Int("n")

		if step == 0 {
			fmt.Println("Apply all migrations...")
			if err := migrateManager.Up(); err != nil {
				return err
			}
		} else {
			fmt.Printf("Apply %d migrations... \n", step)
			if err := migrateManager.Step(step); err != nil {
				return err
			}
		}
		fmt.Println("Successfully!")
		return nil
	}
}
