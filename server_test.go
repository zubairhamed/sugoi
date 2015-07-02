package sugoi
import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestServer(t *testing.T) {
	s := NewSugoi("8080")

	assert.NotNil(t, s)
	assert.NotNil(t, s.handler)
	assert.Equal(t, s.port, "8080")
	assert.Equal(t, 0, len(s.handler.defaultHandlers))

	fn := func(*Request) Content {
		return nil
	}

	Set404Page(s, fn)
	assert.Equal(t, 1, len(s.handler.defaultHandlers))

	Set500Page(s, fn)
	assert.Equal(t, 2, len(s.handler.defaultHandlers))

	Set401Page(s, fn)
	assert.Equal(t, 3, len(s.handler.defaultHandlers))

	Set403Page(s, fn)
	assert.Equal(t, 4, len(s.handler.defaultHandlers))

	Set403Page(s, fn)
	assert.Equal(t, 4, len(s.handler.defaultHandlers))

	s.SetStatic("/static", "static")
	assert.Equal(t, "/static", s.handler.staticUrl)
	assert.Equal(t, "static", s.handler.staticDir)

	// Routes
	s.GET("/", fn)
	assert.Equal(t, 1, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("get")))

	s.DELETE("/", fn)
	assert.Equal(t, 2, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("delete")))

	s.PUT("/", fn)
	assert.Equal(t, 3, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("get")))

	s.POST("/", fn)
	assert.Equal(t, 4, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("get")))

	s.OPTIONS("/", fn)
	assert.Equal(t, 5, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("get")))

	s.PATCH("/", fn)
	assert.Equal(t, 6, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("get")))

	s.GET("/get", fn)
	assert.Equal(t, 7, len(s.handler.routes))
	assert.Equal(t, 2, len(s.GetRoutes("get")))

	bfFn := func(*Request, *Chain) {}

	assert.Equal(t, 0, len(s.handler.beforeFilters))
	s.Before(bfFn)
	assert.Equal(t, 1, len(s.handler.beforeFilters))
	s.Before(bfFn)
	assert.Equal(t, 2, len(s.handler.beforeFilters))
}
