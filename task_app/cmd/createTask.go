package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/momohyusuf/task-app/utils"
	"github.com/spf13/cobra"
)

var createTaskCmd = &cobra.Command{
	Use:   "createTask",
	Short: "Create new task",
	Long:  "Create a new task for the day",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		userTasks := []Task{}
		task := Task{
			ID:          utils.GenerateRandomString(7),
			Description: strings.Join(args, " "),
			Status:      "todo",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		existingTask, err := os.ReadFile(FileName)

		if err != nil {
			userTasks = append(userTasks, task)
			jsonData, err := json.Marshal(userTasks)
			if err != nil {
				fmt.Print(err)
			}
			utils.SaveTaskToFile(jsonData, FileName)
			os.Exit(1)

		}
		err = json.Unmarshal(existingTask, &userTasks)

		if err != nil {
			fmt.Print(err)
		}
		userTasks = append([]Task{task}, userTasks...)
		jsonData, _ := json.Marshal(userTasks)
		utils.SaveTaskToFile(jsonData, FileName)
	},
}

func init() {
	rootCmd.AddCommand(createTaskCmd)
}
