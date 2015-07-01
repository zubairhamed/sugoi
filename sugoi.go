package sugoi

import (
	"errors"
)

type Content interface{}
type BeforeFilter func(*Request, *Chain)
type AfterFilter  func(*Response, *Chain)
type RouteHandler func(*Request) Content
type ErrorHandler func(*Request, error)

var ERR_NO_MATCHING_ROUTE = errors.New("No matching route found")
var ERR_UNSUPPORTED_CONTENT_FORMAT = errors.New("Unsupported Content-Format")
var ERR_NO_MATCHING_METHOD = errors.New("No matching method")

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
