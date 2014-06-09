package controllers

import (
	"encoding/json"
	"math"

	mgt "github.com/donnpebe/todoapirevel/app/errors"
	"github.com/donnpebe/todoapirevel/app/models"
	"github.com/jinzhu/gorm"

	"github.com/revel/revel"
)

type TasksController struct {
	GormController
}

type ResponseTasks struct {
	Page       int           `json:"page"`
	PrevPage   int           `json:"prevPage,omitempty"`
	NextPage   int           `json:"nextPage,omitempty"`
	PerPage    int           `json:"perPage"`
	TotalPages int           `json:"totalPages"`
	Tasks      []models.Task `json:"tasks"`
}

type ResponseTask struct {
	Task *models.Task `json:"task"`
}

type ResponseDeletedTask struct {
	Deleted bool `json:"deleted"`
	Id      int  `json:"id"`
}

func (c *TasksController) List(page, perPage int) revel.Result {
	if page == 0 {
		page = 1
	}
	prevPage := page - 1
	if perPage == 0 {
		perPage = 30
	} else if perPage > 100 {
		perPage = 100
	}

	totalPages := int(math.Ceil(float64(models.CountTasks(Dbm)) / float64(perPage)))
	nextPage := page + 1
	if nextPage > totalPages {
		nextPage = 0
	}
	tasks, err := models.GetTasks(Dbm, page, perPage)
	checkPANIC(err)
	return c.RenderJson(&ResponseTasks{
		page, prevPage, nextPage, perPage, totalPages, tasks})
}

func (c *TasksController) Show(id int) revel.Result {
	if id == 0 {
		panic(mgt.NewMGTError(gorm.RecordNotFound))
	}
	task, err := models.GetTaskById(Dbm, id)
	checkPANIC(err)
	return c.RenderJson(&ResponseTask{task})
}

func (c *TasksController) Create() revel.Result {
	task := models.NewTask("")
	err := json.NewDecoder(c.Request.Body).Decode(task)
	checkPANIC(err)
	task.Id = 0
	task.Sanitize().Validate(c.Validation)
	if c.Validation.HasErrors() {
		panic(mgt.NewMGTError(mgt.MGTInvalidParams, c.Validation.Errors))
	}
	task, err = models.CreateTask(Dbm, task)
	checkPANIC(err)
	return c.RenderJson(&ResponseTask{task})
}

func (c *TasksController) Update(id int) revel.Result {
	if id == 0 {
		panic(mgt.NewMGTError(gorm.RecordNotFound))
	}
	task := models.NewTask("")
	err := json.NewDecoder(c.Request.Body).Decode(task)
	checkPANIC(err)
	task.Sanitize().Validate(c.Validation)
	task, err = models.UpdateTaskName(Dbm, id, task.Name)
	checkPANIC(err)
	return c.RenderJson(&ResponseTask{task})
}

func (c *TasksController) UpdateDone(id int, status string) revel.Result {
	var done bool
	if status == "done" {
		done = true

	} else if status == "undone" {
		done = false
	} else {
		panic(mgt.NewMGTError(mgt.MGTUnrecognizedURL))
	}
	if id == 0 {
		panic(mgt.NewMGTError(gorm.RecordNotFound))
	}
	task, err := models.UpdateTaskDone(Dbm, id, done)
	checkPANIC(err)

	return c.RenderJson(&ResponseTask{task})
}

func (c *TasksController) Delete(id int) revel.Result {
	if id == 0 {
		panic(mgt.NewMGTError(gorm.RecordNotFound))
	}
	id, err := models.DeleteTask(Dbm, id)
	checkPANIC(err)
	return c.RenderJson(&ResponseDeletedTask{true, id})
}
