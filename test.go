package main

type Func func(Request, Response) interface{}

func main() {
	GET("/", func(req Request, res Response) {
		return "Hello"
	})

	DELETE("/", func(req Request, res Response){
		return 200
	})

	PUT("/", func(req Request, res Response){
		return 200
	})

	POST("/", func(req Request, res Response){
		return 200
	})

	OPTIONS("/", func(req Request, res Response){
		return 200
	})

	PATCH("/", func(req Request, res Response){
		return 200
	})
}

func GET(string, Func) {}
func DELETE(string, Func) {}
func PUT(string, Func) {}
func POST(string, Func) {}
func OPTIONS(string, Func) {}
func PATCH(string, Func) {}

type Request struct {
	// Query
	// Cookies
	// Sessions
}

type Response struct {

}


// Functions
func Halt(msg string, code int) {
	// halt(int)
	// halt(msg)
	// halt(int, msg)
}

func HaltWithCode(code int) {

}

func HaltWithMessage(msg string) {

}

func Before() {

}

func After() {

}

