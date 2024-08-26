package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "os"
    "time"
)

// Task represents a to-do item
type Task struct {
    ID      int
    Name    string
    Done    bool
    DueDate time.Time
}

var tasks []Task

// addTask adds a new task to the to-do list
func addTask(name string, due string) {
    var dueDate time.Time
    var err error
    if due != "" {
        dueDate, err = time.Parse("2006-01-02", due)
        if err != nil {
            fmt.Println("Invalid due date format. Use YYYY-MM-DD.")
            return
        }
    } else {
        // Set due date to the end of the current day if not provided
        now := time.Now()
        dueDate = time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
    }

    task := Task{
        ID:      len(tasks) + 1,
        Name:    name,
        Done:    false,
        DueDate: dueDate,
    }
    tasks = append(tasks, task)
    fmt.Println("Added:", name)
}

// listTasks lists all tasks in the to-do list
func listTasks() {
    fmt.Println("To-Do List:")
    for _, task := range tasks {
        status := " "
        if task.Done {
            status = "x"
        }

        due := ""
        if !task.DueDate.IsZero() {
            due = fmt.Sprintf("(Due: %s)", task.DueDate.Format("2024-08-26"))
            if time.Now().After(task.DueDate) && !task.Done {
                due += " [Overdue]"
            }
        }

        fmt.Printf("[%s] %d: %s %s\n", status, task.ID, task.Name, due)
    }
}

// markTaskDone marks a task as completed
func markTaskDone(id int) {
    for i := range tasks {
        if tasks[i].ID == id {
            tasks[i].Done = true
            fmt.Println("Marked as done:", tasks[i].Name)
            return
        }
    }
    fmt.Println("Task not found.")
}

// deleteTask deletes a task by its ID
func deleteTask(id int) {
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            fmt.Println("Deleted:", task.Name)
            return
        }
    }
    fmt.Println("Task not found.")
}

// saveTasks saves the tasks to a JSON file
func saveTasks() {
    file, err := os.Create("tasks.json")
    if err != nil {
        fmt.Println("Error saving tasks:", err)
        return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    err = encoder.Encode(tasks)
    if err != nil {
        fmt.Println("Error encoding tasks:", err)
    }
}

// loadTasks loads the tasks from a JSON file
func loadTasks() {
    file, err := os.Open("tasks.json")
    if err != nil {
        if os.IsNotExist(err) {
            tasks = []Task{}
            return
        }
        fmt.Println("Error loading tasks:", err)
        return
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&tasks)
    if err != nil {
        fmt.Println("Error decoding tasks:", err)
    }
}

func main() {
    loadTasks()

    add := flag.String("add", "", "Task to add")
    due := flag.String("due", "", "Due date for the task (YYYY-MM-DD)")
    list := flag.Bool("list", false, "List all tasks")
    done := flag.Int("done", 0, "Mark a task as done")
    delete := flag.Int("delete", 0, "Delete a task by ID")
    flag.Parse()

    if *add != "" {
        addTask(*add, *due)
    } else if *list {
        listTasks()
    } else if *done != 0 {
        markTaskDone(*done)
    } else if *delete != 0 {
        deleteTask(*delete)
    } else {
        fmt.Println("Usage:")
        fmt.Println("-add 'task name' [-due YYYY-MM-DD] to add a task")
        fmt.Println("-list to list all tasks")
        fmt.Println("-done 1 to mark task 1 as done")
        fmt.Println("-delete 1 to delete task 1")
    }

    saveTasks()
}
