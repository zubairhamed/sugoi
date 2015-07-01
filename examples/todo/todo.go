package main

import (
	. "github.com/zubairhamed/sugoi"
	"github.com/zubairhamed/sugoi/examples/todo/tasks"
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
	server.Set404(func(req *Request) Content {
		return "Content was not found"
	})
}

func setupRoutes(server *SugoiServer) {
	server.GET("/", func(req *Request) Content {
		return Html("index.html", nil)
	})

	server.GET("/tasks", func(req *Request) Content {
		return taskDB.GetAll()
	})

	server.DELETE("/tasks", func(req *Request) Content {
		return OK()
	})

	server.GET("/task/:id", func(req *Request) Content {
		return OK()
	})

	server.PUT("/task/:id", func(req *Request) Content {
		return OK()
	})

	server.POST("/task/:id", func(req *Request) Content {
		return OK()
	})

	server.DELETE("/task/:id", func(req *Request) Content {
		return OK()
	})
}

func setupDB(db *tasks.TasksDB) {

}
