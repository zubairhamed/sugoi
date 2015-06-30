package sugoi

func NewBeforeFilterChain(beforeFilters []BeforeFilter) (*Chain) {
	l := len(beforeFilters)
	if l > 0 {
		fc := beforeFilters[0]
		var nextCh *Chain = nil

		if l > 1 {
			nextCh = NewBeforeFilterChain(beforeFilters[1:])
		}

		return &Chain{
			filter: fc,
			nextChain: nextCh,
		}
	} else {
		return nil
	}
}

func NewAfterFilterChain(afterFilters []*AfterFilter) (*Chain) {
	l := len(afterFilters)
	if l > 0 {
		fc := afterFilters[0]
		var nextCh *Chain = nil

		if l > 1 {
			nextCh = NewAfterFilterChain(afterFilters[1:])
		}

		return &Chain{
			filter: fc,
			nextChain: nextCh,
		}
	} else {
		return nil
	}
}

type Chain struct {
	filter 		interface{}
	nextChain 	*Chain
	lastReq 	*Request
}

func (c *Chain) NextBefore(req *Request) {
	c.lastReq = req
	if c.nextChain != nil {
		c.nextChain.filter.(BeforeFilter)(req, c.nextChain)
	}
}

func (c *Chain) NextAfter(res *Response) {

}

func (c *Chain) GetFilteredRequest() *Request  {
	if c.nextChain != nil {
		return c.nextChain.GetFilteredRequest()
	} else {
		return c.lastReq
	}
}