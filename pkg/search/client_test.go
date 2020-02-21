package search_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bakito/go-search.ch/pkg/search"

	. "gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func Test_Search_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("../../testdata/api-response.xml")
		Assert(t, is.Nil(err))
		_, err = w.Write(content)
		Assert(t, is.Nil(err))
	}))
	defer ts.Close()

	cl := search.NewFor(search.Config{
		Key: "xxx",
		URL: ts.URL,
	})

	res, err := cl.Search("0111111111")
	Assert(t, is.Nil(err))
	Assert(t, res != nil)
	Assert(t, is.Equal("", res.ErrorMessage))
	Assert(t, is.Equal("", res.ErrorReason))
	Assert(t, is.Equal(0, res.ErrorCode))

	Assert(t, is.Len(res.Entry, 2))

	Assert(t, is.Equal("Meier", res.Entry[0].Name))
	Assert(t, is.Equal("John", res.Entry[0].Firstname))

	Assert(t, is.Equal("John Meier IT Consulting", res.Entry[1].Name))
	Assert(t, is.Equal("", res.Entry[1].Firstname))
}

func Test_Search_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("../../testdata/api-error.xml")
		Assert(t, is.Nil(err))
		w.WriteHeader(http.StatusForbidden)
		_, err = w.Write(content)
		Assert(t, is.Nil(err))
	}))
	defer ts.Close()

	cl := search.NewFor(search.Config{
		Key: "xxx",
		URL: ts.URL,
	})

	res, err := cl.Search("0111111111")
	Assert(t, err != nil)
	Assert(t, res != nil)
	Assert(t, is.Equal("The submitted API-Key is invalid or blocked", res.ErrorMessage))
	Assert(t, is.Equal("Forbidden", res.ErrorReason))
	Assert(t, is.Equal(403, res.ErrorCode))

	Assert(t, is.Len(res.Entry, 0))
}
