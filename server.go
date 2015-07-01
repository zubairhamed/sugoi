package sugoi
import (
	"log"
	"net/http"
)

func NewSugoi(port string) *SugoiServer {
	return &SugoiServer{
		handler: NewWrappedHandler(),
		port: port,
	}
}

type SugoiServer struct {
	handler 		*WrappedHandler
	port 			string
}

func (s *SugoiServer) SetStatic(url, dir string) {
	s.handler.staticUrl = url
	s.handler.staticDir = dir
}

func (s *SugoiServer) add(method string, path string, fn RouteHandler) {
	s.handler.routes = append(s.handler.routes, CreateNewRoute(path, method, fn))
}

func (s *SugoiServer) GET(path string, fn RouteHandler) {
	s.add("get", path, fn)
}

func (s *SugoiServer) DELETE(path string, fn RouteHandler) {
	s.add("delete", path, fn)
}

func (s *SugoiServer) PUT(path string, fn RouteHandler) {
	s.add("put", path, fn)
}

func (s *SugoiServer) POST(path string, fn RouteHandler) {
	s.add("post", path, fn)
}

func (s *SugoiServer) OPTIONS(path string, fn RouteHandler) {
	s.add("options", path, fn)
}

func (s *SugoiServer) PATCH(path string, fn RouteHandler) {
	s.add("patch", path, fn)
}

func Set404Page(s *SugoiServer, fn RouteHandler) {
	s.handler.defaultHandlers["404"] = fn
}

func Set500Page(s *SugoiServer, fn RouteHandler) {
	s.handler.defaultHandlers["500"] = fn
}

func (s *SugoiServer) Serve() {
	log.Println("Starting Sugoi!")

	err := http.ListenAndServe(":" + s.port, s.handler)
	if err != nil {
		log.Fatal(err)
	}
}

// Before every method
func (s *SugoiServer) Before(fn BeforeFilter) {
	s.handler.beforeFilters = append(s.handler.beforeFilters, fn)
}

// After every methods
func (s *SugoiServer) After(fn AfterFilter) {
	s.handler.afterFilters = append(s.handler.afterFilters, fn)
}

// Handling Exceptions
func (s *SugoiServer) Error(fn ErrorHandler) {
	s.handler.errorHandlers = append(s.handler.errorHandlers, fn)
}
