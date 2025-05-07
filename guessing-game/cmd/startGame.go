package cmd

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var startGameCmd = &cobra.Command{
	Example: "Welcome to the Number Guessing Game! I'm thinking of a number between 1 and 100. You have 5 chances to guess the correct number.",
	Use:     "start",
	Short:   "Start the game",
	Long:    "Start the game and play",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Please select the difficulty level:",
			Items: []string{"1. Easy (10 chances)", "2. Medium (5 chances)", "3. Hard (3 chances)"},
		}

		numOfAttempt := 0

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		userSelection := strings.Split(result, "")[0]

		chances, _ := strconv.Atoi(userSelection)

		switch chances {
		case 1:
			numOfAttempt = 10
		case 2:
			numOfAttempt = 5
		case 3:
			numOfAttempt = 3
		}

		randomGuess := generateRandomNumber(1, 100)

		playGame(numOfAttempt, randomGuess)

		fmt.Printf("You failed to guess the correct number. the NUMBER is %d", randomGuess)

		continueGame(numOfAttempt, randomGuess)
	},
}

func init() {
	rootCmd.AddCommand(startGameCmd)
}

func generateRandomNumber(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func playGame(numOfAttempt, randomGuess int) {
	for i := 0; i < numOfAttempt; i++ {

		prompt := promptui.Prompt{
			Label: "Please enter you number",
		}

		output, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		userGuessNumber, _ := strconv.Atoi(output)
		if randomGuess == userGuessNumber {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", i)
			os.Exit(1)
		}

		if randomGuess > userGuessNumber {
			fmt.Printf("the number is greater than %d \n", userGuessNumber)
		} else {
			fmt.Printf("the number is less than %d \n", userGuessNumber)
		}

	}

	continueGame(numOfAttempt, randomGuess)

}

func continueGame(numberAttempt, randGuess int) {
	prompt := promptui.Select{
		Label: "Do you want to continue",
		Items: []string{"Yes", "Quit"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "Yes" {
		playGame(numberAttempt, randGuess)
	} else {
		fmt.Print("Thank you for playing my game🎊")
		os.Exit(1)
	}
}
