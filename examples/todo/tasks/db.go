package tasks
import "log"

func NewTaskDB() *TasksDB {
	return &TasksDB{
		tasks: make(map[string]*Task),
	}
}

type TasksDB struct {
	tasks 	map[string]*Task
}

func (db *TasksDB) GetAll() []*Task {
	tasks := []*Task{}

	log.Println(len(db.tasks))
	for _, v := range db.tasks {
		tasks = append(tasks, v)
	}
	return tasks
}

func (db *TasksDB) Put(task *Task) {
	db.tasks[task.Id] = task
}

func (db *TasksDB) Delete(t *Task) {
	delete(db.tasks, t.Id)
}

func (db *TasksDB) Get(id string) *Task {
	return db.tasks[id]
}

func NewTask(id, description string) *Task {
	return &Task{
		Id: id,
		Description: description,
	}
}

type Task struct {
	Id 			string
	Description	string
}