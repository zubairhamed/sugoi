package main

import (
	. "github.com/zubairhamed/sugoi"
	"github.com/zubairhamed/sugoi/examples/todo/tasks"
	"log"
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
	server.SetStatic("/static", "statiwwc")

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
	server.POST("/api/task/:id", handleAddTask)
	server.DELETE("/api/task/:id", handleDeleteTask)
}

func setupDB(db *tasks.TasksDB) {
	t := tasks.NewTask("0", "Add a new Task")
	db.Put(t)
}

// Route Handlers
func handleGetAllTasks(req *Request) Content {
	log.Println("handleGetAllTasks")
	return taskDB.GetAll()
}

func handleDeleteAllTasks(req *Request) Content {
	log.Println("handlDeleteAllTasks")
	return taskDB.GetAll()
}

func handleGetTask(req *Request) Content {
	log.Println("handleGetTask")
	return OK()
}

func handleAddTask(req *Request) Content {
	log.Println("handleAddTask")
	return OK()
}

func handleDeleteTask(req *Request) Content {
	log.Println("Deleting task", req.GetAttribute("id"))
	return OK()
}

