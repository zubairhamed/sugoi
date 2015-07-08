package sugoi

import (
	"errors"
)

type Content interface{}
type PreFilter func(*Request, *Chain)
type RouteHandler func(*Request) Content
type ErrorHandler func(*Request, error)

var ERR_NO_MATCHING_ROUTE = errors.New("No matching route found")
var ERR_UNSUPPORTED_CONTENT_FORMAT = errors.New("Unsupported Content-Format")
var ERR_NO_MATCHING_METHOD = errors.New("No matching method")
