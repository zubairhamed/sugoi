package sugoi

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestResponse(t *testing.T) {
	var content string

	content = msgContent("This is default")
	assert.Equal(t, "This is default", content)

	content = msgContent("This is default", "Not default")
	assert.Equal(t, "Not default", content)

	//////////////////////////////////////////////////////////////////////////////////

	var httpCode HttpCode
	httpCode = InternalServerError()
	assert.Equal(t, http.StatusInternalServerError, httpCode.GetCode())
	assert.Equal(t, "500 - Internal Server Error", httpCode.GetContent())

	httpCode = NotFound()
	assert.Equal(t, http.StatusNotFound, httpCode.GetCode())
	assert.Equal(t, "404 - Not Found", httpCode.GetContent())

	httpCode = OK()
	assert.Equal(t, http.StatusOK, httpCode.GetCode())
	assert.Equal(t, "200 - OK", httpCode.GetContent())

	httpCode = NoContent()
	assert.Equal(t, http.StatusNoContent, httpCode.GetCode())
	assert.Equal(t, "204 - No Content", httpCode.GetContent())

	httpCode = Accepted()
	assert.Equal(t, http.StatusAccepted, httpCode.GetCode())
	assert.Equal(t, "202 - Accepted", httpCode.GetContent())

	httpCode = ServiceUnavailable()
	assert.Equal(t, http.StatusServiceUnavailable, httpCode.GetCode())
	assert.Equal(t, "503 - Service Unavailable", httpCode.GetContent())

	httpCode = BadRequest()
	assert.Equal(t, http.StatusBadRequest, httpCode.GetCode())
	assert.Equal(t, "400 - Bad Request", httpCode.GetContent())

	httpCode = Unauthorized()
	assert.Equal(t, http.StatusUnauthorized, httpCode.GetCode())
	assert.Equal(t, "401 -Unauthorized", httpCode.GetContent())

	httpCode = Forbidden()
	assert.Equal(t, http.StatusForbidden, httpCode.GetCode())
	assert.Equal(t, "403 - Forbidden", httpCode.GetContent())

	httpCode = MethodNotAllowed()
	assert.Equal(t, http.StatusMethodNotAllowed, httpCode.GetCode())
	assert.Equal(t, "405 - Not Allowed", httpCode.GetContent())

	httpCode = NotImplemented()
	assert.Equal(t, http.StatusNotImplemented, httpCode.GetCode())
	assert.Equal(t, "501 - Not Implemented", httpCode.GetContent())

	httpCode = NotModified()
	assert.Equal(t, http.StatusNotModified, httpCode.GetCode())
	assert.Equal(t, "304 - Not Modified", httpCode.GetContent())

	httpCode = UnsupportedMediaType()
	assert.Equal(t, http.StatusUnsupportedMediaType, httpCode.GetCode())
	assert.Equal(t, "415 - Unsupported Media Type", httpCode.GetContent())

	httpCode = Conflict()
	assert.Equal(t, http.StatusConflict, httpCode.GetCode())
	assert.Equal(t, "409 - Conflict", httpCode.GetContent())

	httpCode = NotAcceptable()
	assert.Equal(t, http.StatusNotAcceptable, httpCode.GetCode())
	assert.Equal(t, "406 - Not Acceptable", httpCode.GetContent())

	httpCode = Created()
	assert.Equal(t, http.StatusCreated, httpCode.GetCode())
	assert.Equal(t, "201 - Created", httpCode.GetContent())

	httpCode = Gone()
	assert.Equal(t, http.StatusGone, httpCode.GetCode())
	assert.Equal(t, "410 - Gone", httpCode.GetContent())

	httpCode = Found()
	assert.Equal(t, http.StatusFound, httpCode.GetCode())
	assert.Equal(t, "302 - Found", httpCode.GetContent())

	httpCode = MovedPermanently()
	assert.Equal(t, http.StatusMovedPermanently, httpCode.GetCode())
	assert.Equal(t, "301 - Moved Permanently", httpCode.GetContent())

	//////////////////////////////////////////////////////////////////////////////////

}
