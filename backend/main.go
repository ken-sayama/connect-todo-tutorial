package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"connect-todo-tutorial-backend/cmd"
)

func main() {
	c := &cobra.Command{Use: "select cli mode"}
	cmd.Register(c)
	err := c.Execute()
	if err != nil {
		log.Printf("An error has occurred: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}
