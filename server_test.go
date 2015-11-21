package sugoi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer(t *testing.T) {
	s := NewSugoi("8080")

	assert.NotNil(t, s)
	assert.NotNil(t, s.handler)
	assert.Equal(t, s.host, "8080")
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
	s.Get("/", fn)
	assert.Equal(t, 1, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("get")))

	s.Delete("/", fn)
	assert.Equal(t, 2, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("delete")))

	s.Put("/", fn)
	assert.Equal(t, 3, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("get")))

	s.Post("/", fn)
	assert.Equal(t, 4, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("get")))

	s.Options("/", fn)
	assert.Equal(t, 5, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("get")))

	s.Patch("/", fn)
	assert.Equal(t, 6, len(s.handler.routes))
	assert.Equal(t, 1, len(s.GetRoutes("get")))

	s.Get("/get", fn)
	assert.Equal(t, 7, len(s.handler.routes))
	assert.Equal(t, 2, len(s.GetRoutes("get")))

	bfFn := func(*Request, *Chain) {}

	assert.Equal(t, 0, len(s.handler.preFilters))
	s.Before(bfFn)
	assert.Equal(t, 1, len(s.handler.preFilters))
	s.Before(bfFn)
	assert.Equal(t, 2, len(s.handler.preFilters))
}
