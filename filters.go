package sugoi

func NewPreFilterChain(preFilters []PreFilter) (*Chain) {
	l := len(preFilters)

	if l > 0 {
		fc := preFilters[0]
		var nextCh *Chain = nil

		if l > 1 {
			nextCh = NewPreFilterChain(preFilters[1:])
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

func (c *Chain) NextPre(req *Request) {
	c.lastReq = req
	if c.nextChain != nil {
		c.filter.(PreFilter)(req, c.nextChain)
	}
}

func (c *Chain) GetFilteredRequest() *Request  {
	if c.nextChain != nil {
		return c.nextChain.GetFilteredRequest()
	} else {
		return c.lastReq
	}
}