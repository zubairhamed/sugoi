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

	server.GET("/api/tasks", func(req *Request) Content {
		return taskDB.GetAll()
	})

	server.DELETE("/api/tasks", func(req *Request) Content {
		return OK()
	})

	server.GET("/api/task/:id", func(req *Request) Content {
		return OK()
	})

	server.PUT("/api/task/:id", func(req *Request) Content {
		return OK()
	})

	server.POST("/api/task/:id", func(req *Request) Content {
		return OK()
	})

	server.DELETE("/api/task/:id", func(req *Request) Content {
		return OK()
	})
}

func setupDB(db *tasks.TasksDB) {
	t := tasks.NewTask("0", "Add a new Task")
	db.Put(t)
}
