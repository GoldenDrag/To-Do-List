package main

import (
	"ToDoList/internal/repository"
	"context" // Front
	"database/sql"
)

// App struct
type App struct {
	ctx      context.Context
	taskRepo repository.TaskRepository
}

// NewApp creates a new App application struct
func NewApp(db *sql.DB) *App {
	return &App{taskRepo: &repository.PSQLTaskRepository{DB: db}}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GET tasks and 7 more
func (a *App) GetTasks() ([]repository.Task, error) {
	return a.taskRepo.GetTasks()
}

func (a *App) GetActiveTasks() ([]repository.Task, error) {
	return a.taskRepo.GetActiveTasks()
}

func (a *App) GetCompletedTasks() ([]repository.Task, error) {
	return a.taskRepo.GetCompletedTasks()
}

func (a *App) GetTasksByDate() ([]repository.Task, error) {
	return a.taskRepo.GetTasksByDate()
}

func (a *App) GetTasksByPriority() ([]repository.Task, error) {
	return a.taskRepo.GetTasksByPriority()
}

func (a *App) GetTasksDueToday() ([]repository.Task, error) {
	return a.taskRepo.GetTasksDueToday()
}

func (a *App) GetTasksDueWeek() ([]repository.Task, error) {
	return a.taskRepo.GetTasksDueWeek()
}

func (a *App) GetExpiredTasks() ([]repository.Task, error) {
	return a.taskRepo.GetExpiredTasks()
}

func (a *App) AddTask(task repository.Task) (int, error) {
	return a.taskRepo.AddTask(task)
}

func (a *App) UpdateTask(task repository.Task) error {
	return a.taskRepo.UpdateTask(task)
}

func (a *App) DeleteTask(id int) error {
	return a.taskRepo.DeleteTask(id)
}
