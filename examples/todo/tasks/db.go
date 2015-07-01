package tasks

func NewTaskDB() *TasksDB {
	return &TasksDB{}
}

type TasksDB struct {
	tasks 	map[string]*Task
}

func (db *TasksDB) GetAll() []*Task {
	tasks := make([]*Task, len(db.tasks))
	for _, v := range db.tasks {
		tasks = append(tasks, v)
	}
	return tasks
}

func (db *TasksDB) Put(task *Task) {
	db.tasks[task.id] = task
}

func (db *TasksDB) Delete(t *Task) {
	delete(db.tasks, t.id)
}

func (db *TasksDB) Get(id string) *Task {
	return db.tasks[id]
}

func NewTask(id, description string) *Task {
	return &Task{
		id: id,
		description: description,
	}
}

type Task struct {
	id 			string
	description	string
}