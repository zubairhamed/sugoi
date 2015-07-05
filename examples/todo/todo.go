package main

import (
	. "github.com/zubairhamed/sugoi"
	"github.com/zubairhamed/sugoi/examples/todo/tasks"
	"crypto/rand"
	"encoding/base64"
)

var taskDB *tasks.TasksDB

func main() {
	taskDB = tasks.NewTaskDB()
	setupDB(taskDB)

	server := NewSugoi("8080")

	setupRoutes(server)
	setupDefaults(server)

	server.Serve()
}

func setupDefaults(server *SugoiServer) {
	server.SetStatic("/static", "static")

	Set404Page(server, func(req *Request) Content {
		return StaticHtml("404.html")
	})

	Set500Page(server, func(req *Request) Content {
		return StaticHtml("500.html")
	})
}

func setupRoutes(server *SugoiServer) {
	server.GET("/", func(req *Request) Content {
		return StaticHtml("index.html")
	})

	server.GET("/api/tasks", handleGetAllTasks)
	server.DELETE("/api/tasks", handleDeleteAllTasks)
	server.GET("/api/task/:id", handleGetTask)
	server.POST("/api/task/:description", handleAddTask)
	server.DELETE("/api/task/:id", handleDeleteTask)
}

func setupDB(db *tasks.TasksDB) {
	db.Put(tasks.NewTask("1", "Run 100 miles"))
	db.Put(tasks.NewTask("2", "Buy milk"))
	db.Put(tasks.NewTask("3", "Clean the house"))
	db.Put(tasks.NewTask("0", "Do laundry"))
}

// Route Handlers
func handleGetAllTasks(req *Request) Content {
	return taskDB.GetAll()
}

func handleDeleteAllTasks(req *Request) Content {
	return taskDB.GetAll()
}

func handleGetTask(req *Request) Content {
	id := req.GetAttribute("id")
	return taskDB.Get(id)
}

func handleAddTask(req *Request) Content {
	description := req.GetAttribute("description")

	taskObj := tasks.NewTask(generateId(), description)

	taskDB.Put(taskObj)

	return OK()
}

func handleDeleteTask(req *Request) Content {
	id := req.GetAttribute("id")
	taskDB.Delete(id)

	return OK("Task deleted id " + id)
}

func generateId() string {
	rb := make([]byte, 32)

	rand.Read(rb)

	rs := base64.URLEncoding.EncodeToString(rb)

	return rs
}
