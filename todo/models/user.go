package models

type User struct {
	Name  string
	Id    int
	Tasks []*Task
}

func (usr *User) AddTask(task *Task) {
	// sort and append
	usr.Tasks = append(usr.Tasks, task)
}

func (usr *User) MarkCompleted(taskId int) {

	for _, task := range usr.Tasks {
		if task.Id == taskId {
			task.markAsCompleted()
			usr.DeleteTask(taskId)
			return
		}
	}

}

func (user *User) DeleteTask(taskId int) {
	id := -1
	for i, task := range user.Tasks {
		if task.Id == taskId {
			id = i
		}
	}

	user.Tasks = append(user.Tasks[:id], user.Tasks[id+1:]...)
}

func (usr *User) ModifyTask(taskId int) {

	for _, task := range usr.Tasks {
		if task.Id == taskId {
			// task.markAsCompleted()
			return
		}
	}

}
