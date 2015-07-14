package sugoi

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRequests(t *testing.T) {
	attrs := make(map[string]string)
	attrs["a"] = "val_a"
	attrs["b"] = "val_b"
	attrs["c"] = "val_c"

	attrs["intValue"] = "12345"

	req := NewRequestFromHttp(attrs, &http.Request{})

	assert.NotNil(t, req)
	assert.NotNil(t, req.GetHttpRequest())

	assert.Equal(t, "val_a", req.GetAttribute("a"))
	assert.Equal(t, "val_b", req.GetAttribute("b"))
	assert.Equal(t, "val_c", req.GetAttribute("c"))
	assert.Equal(t, "12345", req.GetAttribute("intValue"))
	assert.Equal(t, 12345, req.GetAttributeAsInt("intValue"))
	assert.Equal(t, 4, len(req.GetAttributes()))
}
