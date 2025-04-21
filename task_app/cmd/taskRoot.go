package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var FileName = "tasks.json"

var rootCmd = &cobra.Command{
	Use:   "day-task",
	Short: "day-task is task app for managing your tasks efficiently",
	Long:  `don't sweat it, we no this is not so much`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
