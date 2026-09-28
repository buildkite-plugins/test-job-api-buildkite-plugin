package main

import (
	"context"
	"fmt"
	"os"

	"github.com/buildkite/agent/v4/jobapi"
	"github.com/kr/pretty"
)

func main() {
	ctx := context.Background()
	c, err := jobapi.NewDefaultClient(ctx)
	if err != nil {
		fatal(fmt.Errorf("starting jobapi client: %w", err))
	}

	switch os.Args[1] {
	case "environment":
		fmt.Println("Adding environment variables MOUNTAIN=cotopaxi and OCEAN=pacific")
		if _, err := c.EnvUpdate(ctx, &jobapi.EnvUpdateRequest{
			Env: map[string]string{
				"MOUNTAIN": "cotopaxi",
				"OCEAN":    "pacific",
			},
		}); err != nil {
			fatal(fmt.Errorf("updating environment variables: %w", err))
		}

	case "post-command":
		fmt.Println("Removing environment variable OCEAN")
		deleted, err := c.EnvDelete(ctx, []string{"OCEAN"})
		if err != nil {
			fatal(fmt.Errorf("deleting environment variable: %w", err))
		}
		pretty.Printf("Deleted: %v\n", deleted)

	default:
		panic("unknown command")
	}
}

func fatal(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
