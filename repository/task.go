package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"fmt"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(taskID int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	filebased *filebased.Data
}

func NewTaskRepo(filebasedDb *filebased.Data) *taskRepository {
	return &taskRepository{
		filebased: filebasedDb,
	}
}

func (t *taskRepository) Store(task *model.Task) error {
	t.filebased.StoreTask(*task)

	return nil
}

func (t *taskRepository) Update(taskID int, task *model.Task) error {
	// Ambil task yang ada berdasarkan ID
	existingTask, err := t.filebased.GetTaskByID(taskID)
	if err != nil {
		return fmt.Errorf("error fetching task: %v", err)
	}

	// Perbarui task yang ada dengan data baru
	existingTask.Title = task.Title
	existingTask.Deadline = task.Deadline
	existingTask.Priority = task.Priority
	existingTask.CategoryID = task.CategoryID
	existingTask.Status = task.Status

	// Simpan task yang telah diperbarui ke database
	err = t.filebased.UpdateTask(taskID, *existingTask)
	if err != nil {
		return fmt.Errorf("error updating task: %v", err)
	}

	return nil
}

func (t *taskRepository) Delete(id int) error {
	err := t.filebased.DeleteTask(id)
	if err != nil {
		if err.Error() == "record not found" {
			return fmt.Errorf("record not found")
		}
		return fmt.Errorf("error deleting task: %v", err) 
	}
	return nil
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	task, err := t.filebased.GetTaskByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("record not found") 
		}
		return nil, err 
	}
	return task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	tasks, err := t.filebased.GetTasks()
	if err != nil {
		return nil, fmt.Errorf("error fetching tasks: %v", err)
	}
	return tasks, nil
}

func (t *taskRepository) GetTaskCategory(categoryID int) ([]model.TaskCategory, error) {
	var taskCategories []model.TaskCategory
	tasks, err := t.filebased.GetTasks()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve tasks: %w", err)
	}

	for _, task := range tasks {
		if task.CategoryID == categoryID {
			category, err := t.filebased.GetCategoryByID(categoryID)
			if err != nil {
				return nil, fmt.Errorf("failed to retrieve category: %w", err)
			}
			taskCategories = append(taskCategories, model.TaskCategory{
				ID:       task.ID,
				Title:    task.Title,
				Category: category.Name,
			})
		}
	}

	if len(taskCategories) == 0 {
		return nil, fmt.Errorf("no tasks found for category ID: %d", categoryID)
	}

	return taskCategories, nil
}