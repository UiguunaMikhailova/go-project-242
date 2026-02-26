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
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(context.Context, *cli.Command) error {
			fmt.Println("boom! I say!", os.Args[1])

			size, err := code.GetSize(os.Args[1])
			if err != nil {
				fmt.Println("Error: ", err)
			}
			fmt.Printf("%s\t%s", size, os.Args[1])
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println("Error: ", err)
	}

	// size, err := GetSize("/Users/ujguunamihajlova/my_study/go_learn/go-project-242/cmd/hexlet-path-size/main.go")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Size: ", size)

	// fmt.Println("Hello from Hexlet!")
}
