package cmd

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/momohyusuf/task-app/utils"
	"github.com/spf13/cobra"
)

var updateTaskProgressCmd = &cobra.Command{
	Use:   "update-status",
	Short: "Update an existing task",
	Long:  "Update your task",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		existingTasks, err := os.ReadFile(FileName)
		update_args := strings.Split(strings.Join(args, " "), " ")

		taskStatus := update_args[0]
		taskId := strings.Join(update_args[1:], " ")
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
				switch taskStatus {
				case "mark-in-progress":
					userTasks[i].Status = "in-progress"
				case "mark-done":
					userTasks[i].Status = "done"
				default:
					log.Fatal("Invalid task status")
				}
			}
		}

		jsonData, _ := json.Marshal(userTasks)
		utils.SaveTaskToFile(jsonData, FileName)

	},
}

func init() {
	rootCmd.AddCommand(updateTaskProgressCmd)
}
