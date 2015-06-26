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
		// get attribute
		// set attribute
}

type Response struct {
	// redirect

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

// Before every method
func BeforeAll(Func) {

}

func Before(pattern string, Func) {

}

// After every methods
func After(pattern string, Func) {

}

func AfterAll(Func) {

}

// Handling Exceptions
func Error(err error, fn Func) {

}