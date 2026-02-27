package main

import (
	"context"
	"fmt"
	"os"

	"code"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:   "hexlet-path-size - print size of a file or directory",
		Usage:  "hexlet-path-size [global options]",
		Flags:  flags,
		Action: action,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}

var flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "human",
		Aliases: []string{"H"},
		Value:   false,
		Usage:   "human-readable sizes (auto-select unit)",
	},
	&cli.BoolFlag{
		Name:    "all",
		Aliases: []string{"a"},
		Value:   false,
		Usage:   "include hidden files and directories",
	},
}

func action(_ context.Context, cmd *cli.Command) error {
	path := cmd.Args().Get(0)
	human := cmd.Bool("human")
	all := cmd.Bool("all")

	print(path, human, all)

	size, err := code.GetSize(path, human, all)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("%s\t%s", size, path)
	return nil
}

func print(path string, human, all bool) {
	fmt.Println("Path: ", path)

	if human {
		fmt.Println("Human: ", human)
	} else {
		fmt.Println("Not human: ", human)
	}

	if all {
		fmt.Println("All: ", all)
	} else {
		fmt.Println("Not all: ", all)
	}
}
