package sugoi
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
)


func TestFilters(t *testing.T) {
	chain1Called := false
	chain2Called := false
	chain3Called := false

	chain1 := func(req *Request, ch *Chain) {
		chain1Called = true

		ch.NextPre(req)
	}

	chain2 := func(req *Request, ch *Chain) {
		chain2Called = true

		ch.NextPre(req)
	}

	chain3 := func(req *Request, ch *Chain) {
		chain3Called = true

		ch.NextPre(req)
	}

	preFilters := []PreFilter{
		chain1, chain2, chain3,
	}

	chain := NewPreFilterChain(preFilters)
	chain.NextPre(nil)

	assert.True(t, chain1Called)
	assert.True(t, chain2Called)
	assert.True(t, chain3Called)
}
