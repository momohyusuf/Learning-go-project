# Guessing Game CLI

## Overview
The Guessing Game CLI is a command-line application built with Go. It allows users to play a number guessing game with varying difficulty levels. The project is structured using the Cobra library to manage commands and subcommands.

## Features
- **Difficulty Levels**: Choose from Easy, Medium, or Hard difficulty levels, each with a different number of attempts.
- **Interactive CLI**: Uses the `promptui` library for an interactive and user-friendly command-line interface.
- **Error Handling**: Graceful handling of invalid inputs and errors.

## Project Structure
```
├── go.mod
├── go.sum
├── main.go
├── cmd/
│   ├── execute.go
│   ├── startGame.go
```
- **`main.go`**: The entry point of the application. It initializes and executes the root command.
- **`cmd/execute.go`**: Contains the root command definition and the `Execute` function to run the CLI.
- **`cmd/startGame.go`**: Implements the logic for starting the guessing game, including difficulty selection and game mechanics.

## Prerequisites
- Go 1.20 or later
- macOS, Linux, or Windows

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/guessing-game.git
   cd guessing-game
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Build the application:
   ```bash
   go build -o guessing-game
   ```

## Usage
Run the application using the following command:
```bash
./guessing-game
```

### Commands
- **Root Command**: Displays the CLI tool's description and usage.
- **Start Game**: Prompts the user to select a difficulty level and starts the guessing game.

## Example
1. Run the application:
   ```bash
   ./guessing-game
   ```
2. Select a difficulty level from the prompt.
3. Start guessing the number within the allowed attempts.

## Development
To add new features or modify the application:
1. Create a new branch:
   ```bash
   git checkout -b feature-name
   ```
2. Make your changes and commit them:
   ```bash
   git commit -m "Add feature description"
   ```
3. Push the branch and create a pull request:
   ```bash
   git push origin feature-name
   ```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request for review.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

## Acknowledgments
- [Cobra Library](https://github.com/spf13/cobra) for command-line interface management.
- [PromptUI Library](https://github.com/manifoldco/promptui) for interactive prompts.