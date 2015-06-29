package sugoi

import (
	"errors"
)

type Response interface{}
type BeforeChainHandler func(Request, *Chain)
type AfterChainHandler  func(Response, *Chain)
type RouteHandler func(Request) Response
type ErrorHandler func(Request, error)

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


// Return Content Types
	// text
	// object : implicit json
	// various error code objects
	// specific content type (with converters)
		// xml, json, plain text, binary etc
	// with html template