package services

import (
	"lld/todo/models"
	"time"
)

func CreateUser(name string, id int) *models.User {
	return &models.User{
		Name: name,
		Id:   id,
	}
}

func CreateTask(id int, title string, description string, date time.Time, userId int) *models.Task {
	return &models.Task{
		Id:          id,
		UserId:      userId,
		Title:       title,
		Description: description,
		DueDate:     date,
		IsCompleted: false,
	}
}

func (usr *models.User) AddTask(task *models.Task) {

	usr.Tasks = append(usr.Tasks, task)

}
