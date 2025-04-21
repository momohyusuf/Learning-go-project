package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var listTaskCmd = &cobra.Command{
	Use:   "list",
	Short: "Update an existing task",
	Long:  "Update your task",
	Run: func(cmd *cobra.Command, args []string) {
		existingTasks, err := os.ReadFile(FileName)
		task_status := strings.Join(args, " ")

		userTasks := []Task{}
		filteredTasks := []Task{}
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(existingTasks, &userTasks)

		if err != nil {
			log.Fatal(err)
		}

		for i := range userTasks {
			if userTasks[i].Status == task_status {
				filteredTasks = append(filteredTasks, userTasks[i])
			}
		}

		fmt.Printf("Totals %s tasks: %d \n", task_status, len(filteredTasks))
		fmt.Println(filteredTasks)
	},
}

func init() {
	rootCmd.AddCommand(listTaskCmd)
}
