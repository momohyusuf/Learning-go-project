# Task App

Task App is a command-line application designed to help you manage your daily tasks efficiently. It allows you to create, update, delete, and list tasks with ease.

## Features

- **Create Task**: Add a new task to your task list.
- **Update Task**: Modify the description of an existing task.
- **Update Task Status**: Change the status of a task (e.g., mark as "in-progress" or "done").
- **Delete Task**: Remove a task from your task list.
- **List Tasks**: View tasks filtered by their status.

## Project Structure

```
task_app/
├── cmd/
│   ├── createTask.go       # Command to create a new task
│   ├── deleteTask.go       # Command to delete a task
│   ├── listAllTask.go      # Command to list tasks by status
│   ├── taskRoot.go         # Root command definition
│   ├── updateTask.go       # Command to update task description
│   ├── updateTaskProgess.go # Command to update task status
├── utils/
│   ├── helper.go           # Utility functions (e.g., save tasks, generate random strings)
├── tasks.json              # JSON file to store tasks
├── main.go                 # Entry point of the application
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies checksum file
```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/momohyusuf/task-app.git
   cd task-app
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build -o task-app
   ```

## Usage

Run the application using the following commands:

### Create a Task

```bash
./task-app createTask "Your task description"
```

### Update a Task Description

```bash
./task-app update <task_id> "Updated task description"
```

### Update Task Status

```bash
./task-app update-status <status> <task_id>
# Example: ./task-app update-status mark-done BGKMcOR
```

### Delete a Task

```bash
./task-app delete <task_id>
```

### List Tasks by Status

```bash
./task-app list <status>
# Example: ./task-app list todo
```

## Task Statuses

- `todo`: Task is yet to be started.
- `in-progress`: Task is currently being worked on.
- `done`: Task is completed.

## Configuration

The tasks are stored in a JSON file named `tasks.json` in the project directory. You can modify the file path by changing the `FileName` variable in `cmd/taskRoot.go`.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
