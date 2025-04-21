package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/momohyusuf/task-app/utils"
	"github.com/spf13/cobra"
)

var updateTaskCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing task",
	Long:  "Update your task",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		existingTasks, err := os.ReadFile(FileName)
		update_args := strings.Split(strings.Join(args, " "), " ")

		taskId := update_args[0]
		taskContent := strings.Join(update_args[1:], " ")
		userTasks := []Task{}
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(existingTasks, &userTasks)

		if err != nil {
			log.Fatal(err)
		}

		for i := range userTasks {
			if userTasks[i].ID == taskId {
				userTasks[i].Description = taskContent
			}
		}

		fmt.Println(userTasks)
		jsonData, _ := json.Marshal(userTasks)
		utils.SaveTaskToFile(jsonData, FileName)

	},
}

func init() {
	rootCmd.AddCommand(updateTaskCmd)
}
