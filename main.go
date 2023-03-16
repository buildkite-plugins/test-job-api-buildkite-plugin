package main

import (
	"context"
	"fmt"
	"os"

	"github.com/buildkite/agent/v3/jobapi"
	"github.com/kr/pretty"
)

func main() {
	c, err := jobapi.NewDefaultClient()
	if err != nil {
		fatal(fmt.Errorf("starting jobapi client: %w", err))
	}

	switch os.Args[1] {
	case "environment":
		fmt.Println("Adding environment variables MOUNTAIN=cotopaxi and OCEAN=pacific")
		c.EnvUpdate(context.Background(), &jobapi.EnvUpdateRequest{
			Env: map[string]*string{
				"MOUNTAIN": pt("cotopaxi"),
				"OCEAN":    pt("pacific"),
			},
		})

	case "post-command":
		fmt.Println("Removing environment variable OCEAN")
		deleted, err := c.EnvDelete(context.Background(), []string{"OCEAN"})
		if err != nil {
			fatal(fmt.Errorf("deleting environment variable: %w", err))
		}
		pretty.Printf("Deleted: %v\n", deleted)

	default:
		panic("unknown command")
	}
}

func pt(s string) *string {
	return &s
}

func fatal(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
