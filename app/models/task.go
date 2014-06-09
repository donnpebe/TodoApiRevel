package models

import (
	"github.com/donnpebe/todoapirevel/app/services"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"time"
)

type Task struct {
	Id   int    		`json:"id"`
	Name string 		`json:"name"`
	Done bool   		`json:"done"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTask(name string, attrs ...interface{}) *Task {
	task := new(Task)
	task.Name = name
	if attrs != nil {
		if b, ok := attrs[0].(bool); ok {
			task.Done = b
		}
	}
	return task
}

func CountTasks(db *gorm.DB) int {
	count := 0
	db.Model(Task{}).Count(&count)
	return count
}

func GetTasks(db *gorm.DB, page, perPage int) ([]Task, error) {
	var tasks []Task
	revel.INFO.Println(page)
	err := db.Limit(perPage).Offset((page - 1) * perPage).Order("Id asc").Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskById(db *gorm.DB, id int) (*Task, error) {
	task := new(Task)
	err := db.First(task, id).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func CreateTask(db *gorm.DB, task *Task) (*Task, error) {
	err := db.Save(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func UpdateTaskName(db *gorm.DB, id int, name string) (*Task, error) {
	task := new(Task)
	err := db.First(task, id).Update("name", name).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func UpdateTaskDone(db *gorm.DB, id int, done bool) (*Task, error) {
	task := new(Task)
	err := db.First(task, id).Update("done", done).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func DeleteTask(db *gorm.DB, id int) (int, error) {
	task := new(Task)
	err := db.First(task, id).Delete(task).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t *Task) Sanitize() *Task {
	t.Name = services.SanitizeString(t.Name)
	return t
}

func (t *Task) Validate(v *revel.Validation) *Task {
	v.Required(t.Name).Message("Task's name required.").Key("Task.Name")
	v.MaxSize(t.Name, 140).Message("Task's name must be at most 140 chars long.").Key("Task.Name")
	return t
}
