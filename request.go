package sugoi
import "strconv"

func NewRequestFromHttp(attrs map[string]string) Request {
	return Request{
		attrs: attrs,
	}
}

type Request struct {
	attrs 	map[string]string
	// Query
	// Cookies
	// Sessions
	// get attribute
	// set attribute
}

func (c *Request) GetAttributes() map[string]string {
	return c.attrs
}

func (c *Request) GetAttribute(o string) string {
	return c.attrs[o]
}

func (c *Request) GetAttributeAsInt(o string) int {
	attr := c.GetAttribute(o)
	i, _ := strconv.Atoi(attr)

	return i
}


func NewWrappedHandler() (*WrappedHandler) {
	return &WrappedHandler{
		routes : []*Route{},
	}
}