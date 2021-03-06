# sugoi
[![GoDoc](https://godoc.org/github.com/zubairhamed/sugoi?status.svg)](https://godoc.org/github.com/zubairhamed/sugoi)
[![Build Status](https://drone.io/github.com/zubairhamed/sugoi/status.png)](https://drone.io/github.com/zubairhamed/sugoi/latest)
[![Coverage Status](https://coveralls.io/repos/zubairhamed/sugoi/badge.svg?branch=master)](https://coveralls.io/r/zubairhamed/sugoi?branch=master)

#### Sugoi is a minimalist, no-fuss web framework for Go.

### Example
```go
package main

import (
	. "github.com/zubairhamed/sugoi"
)

func main() {
	server := NewSugoi("8080")
	
	server.GET("/hello", func(req *Request) Content {
		return "Hello, Sugoi!"
	})
	
	server.Serve()
}
```

### Defining and Parameterizing Routes
```go
	server.GET("/hello/:name", func(req *Request) Content {
		name := req.GetAttribute("name")
		return "Hello, !" + name
	})

```

### Returning values

##### Plain Text
```go
	server.GET("/ep", func(req *Request) Content {
		return "Hello, Sugoi!"
	})

```

##### JSON

Any objects returned are automatically converted to JSON via the Go JSON Marshaler

```go
	server.GET("/ep", func(req *Request) Content {
		p := NewPerson("Joe", 25)
		
		return p
	})

```

##### Static HTML
```go
	server.GET("/ep", func(req *Request) Content {
		return StaticHtml("index.html")
	})
```

##### Go HTML Template
```go
	server.GET("/ep", func(req *Request) Content {
		model := NewPerson("Joe", 25)
	
		return TemplateHtml("index.html", m)
	})
```

##### Http Codes
```go
	server.GET("/ep", func(req *Request) Content {
		name := "Joe"
		rec := db.GetByName(name)
		
		if rec == nil {
			// Returns a 404
			return NotFound("Record " + name + " was not found") 	
		} else {
			return rec
		}
	})
```

### Pre-filters
Pre-filters are code-chains which are executed before a request is passed onto a user-defined RouteHandler

### More examples
See /examples/todo for a Todo List example using AngularJS






