package httpietocurl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddHeader(t *testing.T) {
	r := Request{}
	r.AddHeader(Header{"foo", "bar"})
	assert.Equal(t, r.Headers, []Header{{"foo", "bar"}})
}

func TestAddForm(t *testing.T) {
	r := Request{}
	r.AddForm(Form{"foo", "bar"})
	assert.Equal(t, r.Forms, []Form{{"foo", "bar"}})
}

func TestAddJsonField(t *testing.T) {
	r := Request{}
	r.AddJsonField("foo", "bar")
	assert.Equal(t, r.Json, map[string]any{"foo": "bar"})
}

func TestAddQuery(t *testing.T) {
	r := Request{}
	r.AddQuery(Query{"foo", "bar"})
	assert.Equal(t, r.Queries, []Query{{"foo", "bar"}})
}
