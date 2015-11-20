package main

import (
	"crypto/rand"
	"encoding/base64"
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

	server.Start()
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
	server.Get("/", func(req *Request) Content {
		return StaticHtml("index.html")
	})

	server.Get("/api/tasks", getAllItems)
	server.Delete("/api/tasks", deleteAllItems)
	server.Get("/api/task/:id", getItem)
	server.Post("/api/task/:description", addItem)
	server.Put("/api/task/:id/complete", completeItem)
	server.Put("/api/task/:id/uncomplete", uncompleteItem)
	server.Delete("/api/task/:id", deleteItem)
}

func setupDB(db *tasks.TasksDB) {
	db.Put(tasks.NewTask("1", "Run 100 miles"))
	db.Put(tasks.NewTask("2", "Buy milk"))
	db.Put(tasks.NewTask("3", "Clean the house"))
	db.Put(tasks.NewTask("0", "Do laundry"))
}

// Route Handlers
func completeItem(req *Request) Content {
	return OK()
}

func uncompleteItem(req *Request) Content {
	return OK()
}

func getAllItems(req *Request) Content {
	return taskDB.GetAll()
}

func deleteAllItems(req *Request) Content {
	return taskDB.GetAll()
}

func getItem(req *Request) Content {
	id := req.GetAttribute("id")
	return taskDB.Get(id)
}

func addItem(req *Request) Content {
	description := req.GetAttribute("description")

	taskObj := tasks.NewTask(generateId(), description)

	taskDB.Put(taskObj)

	return OK()
}

func deleteItem(req *Request) Content {
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
