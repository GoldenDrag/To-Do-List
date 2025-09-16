package repository

import (
	"database/sql"
	"fmt"
	"log"
)

type Task struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
	Priority  string `json:"priority"`
	DueDate   string `json:"due_date"`
}

type TaskRepository interface {
	AddTask(task Task) (int, error)
	GetTasks() ([]Task, error)
	GetActiveTasks() ([]Task, error)
	GetCompletedTasks() ([]Task, error)
	GetTasksByDate() ([]Task, error)
	GetTasksByPriority() ([]Task, error)
	GetTasksDueToday() ([]Task, error)
	GetTasksDueWeek() ([]Task, error)
	GetExpiredTasks() ([]Task, error)
	UpdateTask(task Task) error
	DeleteTask(id int) error
}

type PSQLTaskRepository struct {
	DB *sql.DB
}

// REST functions
// 7 GET funcs because I forgor app.go needs to be modified
// Generally function constructed next way
// 1. prepare query for psql
// 2. execute
// 3. check for error
// 4. return results

func (r *PSQLTaskRepository) AddTask(task Task) (int, error) {
	insertSQL := `
	INSERT INTO tasks (description, completed, priority, due_date)
	VALUES ($1, $2, $3, $4);`

	res, err := r.DB.Exec(insertSQL, task.Text, task.Completed, task.Priority, task.DueDate)
	if err != nil {
		log.Fatalf("Failed to insert task: %v", err)
		return 0, err
	}
	fmt.Println("Task inserted successfully.")

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *PSQLTaskRepository) GetTasks() ([]Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatalf("Failed to fetch tasks: %v", err)
		return nil, err
	}
	fmt.Println("Tasks fetched successfully.")
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Text, &task.Completed, &task.DueDate, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *PSQLTaskRepository) GetActiveTasks() ([]Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks WHERE NOT completed")
	if err != nil {
		log.Fatalf("Failed to fetch tasks: %v", err)
		return nil, err
	}
	fmt.Println("Active tasks fetched successfully.")
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Text, &task.Completed, &task.DueDate, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *PSQLTaskRepository) GetCompletedTasks() ([]Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks WHERE completed")
	if err != nil {
		log.Fatalf("Failed to fetch tasks: %v", err)
		return nil, err
	}
	fmt.Println("Completed Tasks fetched successfully.")
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Text, &task.Completed, &task.DueDate, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *PSQLTaskRepository) GetTasksByDate() ([]Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks ORDER BY due_date")
	if err != nil {
		log.Fatalf("Failed to fetch tasks: %v", err)
		return nil, err
	}
	fmt.Println("Tasks By Date fetched successfully.")
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Text, &task.Completed, &task.DueDate, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *PSQLTaskRepository) GetTasksByPriority() ([]Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks ORDER BY priority desc")
	if err != nil {
		log.Fatalf("Failed to fetch tasks: %v", err)
		return nil, err
	}
	fmt.Println("Tasks by prority fetched successfully.")
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Text, &task.Completed, &task.DueDate, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *PSQLTaskRepository) GetTasksDueToday() ([]Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks WHERE due_date::DATE = CURRENT_DATE")
	if err != nil {
		log.Fatalf("Failed to fetch tasks: %v", err)
		return nil, err
	}
	fmt.Println("Tasks due today fetched successfully.")
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Text, &task.Completed, &task.DueDate, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *PSQLTaskRepository) GetTasksDueWeek() ([]Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks WHERE due_date BETWEEN NOW() AND NOW() + INTERVAL '1 week'")
	if err != nil {
		log.Fatalf("Failed to fetch tasks: %v", err)
		return nil, err
	}
	fmt.Println("Tasks due in a week fetched successfully.")
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Text, &task.Completed, &task.DueDate, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *PSQLTaskRepository) GetExpiredTasks() ([]Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks WHERE due_date < NOW()")
	if err != nil {
		log.Fatalf("Failed to fetch tasks: %v", err)
		return nil, err
	}
	fmt.Println("Expired Tasks fetched successfully.")
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Text, &task.Completed, &task.DueDate, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *PSQLTaskRepository) UpdateTask(task Task) error {
	updateSQL := `
	UPDATE tasks
	   SET description = $2, completed = $3, due_date = $4, priority = $5 
	 WHERE id = $1;`

	_, err := r.DB.Exec(updateSQL, task.ID, task.Text, task.Completed, task.Priority, task.DueDate)
	if err != nil {
		log.Fatalf("Failed to update task: %v", err)
	}
	fmt.Println("Task updated.")

	return err
}

func (r *PSQLTaskRepository) DeleteTask(id int) error {
	deleteSQL := `
	DELETE FROM tasks
	WHERE id = $1;`

	_, err := r.DB.Exec(deleteSQL, id)
	if err != nil {
		log.Fatalf("Failed to delete task: %v", err)
	}
	fmt.Println("Task deleted successfully.")

	return err
}
