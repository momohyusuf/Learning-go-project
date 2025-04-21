package cmd

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/momohyusuf/task-app/utils"
	"github.com/spf13/cobra"
)

var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user task",
	Long:  "Deleting of a user task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		existingTasks, err := os.ReadFile(FileName)
		task_id := strings.Join(args, " ")
		userTasks := []Task{}

		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(existingTasks, &userTasks)

		if err != nil {
			log.Fatal(err)
		}
		var itemToDelete int
		for i := range userTasks {
			if userTasks[i].ID == task_id {
				itemToDelete = i
			}
		}

		userTasks = append(userTasks[:itemToDelete], userTasks[itemToDelete+1:]...)
		jsonData, _ := json.Marshal(userTasks)
		utils.SaveTaskToFile(jsonData, FileName)
	},
}

func init() {
	rootCmd.AddCommand(deleteTaskCmd)
}
