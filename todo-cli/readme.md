Here's a README file for your Command-Line To-Do List application in Go:

---

# Command-Line To-Do List Application

This is a simple Command-Line To-Do List application written in Go. It allows you to manage tasks by adding, listing, marking them as done, and deleting them. Tasks can have optional due dates, and if no due date is provided, the current date (end of the day) is automatically assigned.

## Features

- **Add Tasks:** Add new tasks to your to-do list with an optional due date.
- **List Tasks:** View all tasks, their statuses, and due dates.
- **Mark Tasks as Done:** Mark tasks as completed.
- **Delete Tasks:** Remove tasks from the list.
- **Automatic Due Date:** If no due date is provided, the task is due by the end of the current day.

## Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Build the application:**
   ```bash
   go build -o todo main.go
   ```

3. **Run the application:**
   ```bash
   ./todo [options]
   ```

## Usage

### Add a Task

To add a task, use the `-add` flag followed by the task name. You can optionally specify a due date using the `-due` flag.

- **Without Due Date:**
  ```bash
  ./todo -add "Buy groceries"
  ```
  - The task will be due by the end of the current day.

- **With Due Date:**
  ```bash
  ./todo -add "Submit report" -due 2024-08-31
  ```
  - The task will be due on `2024-08-31`.

### List All Tasks

To list all tasks, use the `-list` flag:

```bash
./todo -list
```

### Mark a Task as Done

To mark a task as done, use the `-done` flag followed by the task ID:

```bash
./todo -done 1
```

### Delete a Task

To delete a task, use the `-delete` flag followed by the task ID:

```bash
./todo -delete 1
```

### Example Workflow

1. **Add tasks:**
   ```bash
   ./todo -add "Buy groceries"
   ./todo -add "Submit report" -due 2024-08-31
   ```

2. **List tasks:**
   ```bash
   ./todo -list
   ```

3. **Mark a task as done:**
   ```bash
   ./todo -done 1
   ```

4. **Delete a task:**
   ```bash
   ./todo -delete 2
   ```

## File Storage

Tasks are saved in a `tasks.json` file in the same directory as the application. This file is automatically loaded when the application starts and saved when the application exits.

## Contributing

If you'd like to contribute to this project, feel free to fork the repository and submit a pull request.

## License

This project is just for learning pupose.

---

This README provides clear instructions on how to use the application, including installation, usage examples, and how tasks are managed.