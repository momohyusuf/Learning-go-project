# GitHub Activity CLI

GitHub Activity CLI is a command-line tool that allows you to fetch and display the recent public activity of a GitHub user. It uses the GitHub REST API to retrieve the data and outputs it in a readable format.

## Features

- Fetch recent public activity of a GitHub user.
- Save the raw activity data to a `user-activity.json` file.
- Display activity details such as type, repository, and creation date.

## Prerequisites

- Go 1.23.2 or later installed on your system.
- An active internet connection to access the GitHub API.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/github_activity.git
   cd github_activity
   ```

2. Build the project:

   ```sh
   go build -o github-activity
   ```

## Usage

Run the CLI tool with the GitHub username as an argument:

```sh
./github-activity <github-username>
```

For example:

```sh
./github-activity octocat
```

This will fetch the recent public activity of the user `octocat` and display it in the terminal. The raw data will also be saved to a file named `user-activity.json`.

## Project Structure

```
github_activity/
├── cmd/
│   └── root.go          # Contains the main CLI logic
├── main.go              # Entry point of the application
├── user-activity.json   # Output file for raw activity data
├── go.mod               # Go module file
├── go.sum               # Dependency lock file
```

## Example Output

```plaintext
Your username octocat
Type: PushEvent
Repo: map[name:octocat/Hello-World]
Created At: 2023-10-01T12:34:56Z
---
Type: ForkEvent
Repo: map[name:octocat/Spoon-Knife]
Created At: 2023-10-01T11:22:33Z
---
```

## Notes

- The `user-activity.json` file will be overwritten each time the tool is run.
- If the username is invalid or the user has no public activity, the tool will display an error message.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
