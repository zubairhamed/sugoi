package main

import . "github.com/zubairhamed/sugoi"

func main() {
	s := NewSugoi("8085")


	s.Before(func (req Request, chain *Chain) {
		chain.NextBefore(req)
	})

	s.After(func (res Response, chain *Chain) {
		chain.NextAfter(res)
	})

	s.GET("/", func(req Request) Response {
		return "Welcome"
	})

	s.GET("/notfound", func(req Request) Response {
		return NotFound()
	})

	s.GET("/causeserror", func(req Request) Response {
		return InternalServerError()
	})

	s.GET("/hello/:name", func(req Request) Response {
		name := req.GetAttribute("name")

		return "hello " + name
	})

	s.GET("/person", func(req Request) Response {

		var person struct {
			Name 	string
			Age 	int
		}

		person.Name = "Zoob"
		person.Age = 37

		return person
	})

	s.DELETE("/", func(req Request) Response {
		return nil
	})

	s.PUT("/", func(req Request) Response {
		return nil
	})

	s.POST("/", func(req Request) Response {
		return nil
	})

	s.OPTIONS("/", func(req Request) Response {
		return nil
	})

	s.PATCH("/", func(req Request) Response {
		return nil
	})

	s.Serve()
}