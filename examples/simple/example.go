package main

import (
	. "github.com/zubairhamed/sugoi"
)

func main() {
	s := NewSugoi("8085")

	s.Before(func (req *Request, chain *Chain) {
		chain.NextBefore(req)
	})

	s.Before(func (req *Request, chain *Chain) {
		chain.NextBefore(req)
	})

//	s.After(func (res *Response, chain *Chain) {
//		chain.NextAfter(res)
//	})

	s.GET("/", func(req *Request) Response {
		req.GetHeaders()
		return "Welcome"
	})

	s.GET("/err404", func(req *Request) Response {
		return NotFound()
	})

	s.GET("/err500", func(req *Request) Response {
		return InternalServerError()
	})

	s.GET("/params/json/:id", func(req *Request) Response {
		id := req.GetAttribute("id")

		var thing struct {
			Name 	string
			Id 		string
		}

		thing.Id   = id
		thing.Name = "Thing One"

		return OK()
	})

	s.GET("/params/:val1/:val2", func(req *Request) Response {
		val1 := req.GetAttribute("val1")
		val2 := req.GetAttribute("val2")

		return "Values: " + val1 + "," + val2
	})

	s.DELETE("/", func(req *Request) Response {
		return nil
	})

	s.PUT("/", func(req *Request) Response {
		return nil
	})

	s.POST("/", func(req *Request) Response {
		return nil
	})

	s.OPTIONS("/", func(req *Request) Response {
		return nil
	})

	s.PATCH("/", func(req *Request) Response {
		return nil
	})

	s.Serve()
}
