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

	s.GET("/", func(req *Request) Content {
		return "Welcome"
	})

	s.GET("/err404", func(req *Request) Content {
		return NotFound()
	})

	s.GET("/err500", func(req *Request) Content {
		return InternalServerError()
	})

	s.GET("/params/json/:id", func(req *Request) Content {
		id := req.GetAttribute("id")

		var thing struct {
			Name 	string
			Id 		string
		}

		thing.Id   = id
		thing.Name = "Thing One"

		return thing
	})

	s.GET("/params/:val1/:val2", func(req *Request) Content {
		val1 := req.GetAttribute("val1")
		val2 := req.GetAttribute("val2")

		return "Values: " + val1 + "," + val2
	})

	s.DELETE("/", func(req *Request) Content {
		return nil
	})

	s.PUT("/", func(req *Request) Content {
		return nil
	})

	s.POST("/", func(req *Request) Content {
		return nil
	})

	s.OPTIONS("/", func(req *Request) Content {
		return nil
	})

	s.PATCH("/", func(req *Request) Content {
		return nil
	})

	s.Serve()
}
