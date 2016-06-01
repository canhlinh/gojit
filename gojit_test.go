package gojit

import (
	"net/http"
	"runtime"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
)

var goVersion = runtime.Version()

func TestHelloWorld(t *testing.T) {
	r := New()

	r.GET("/hello").
		SetDebug(true).
		Run(GojiMux(), func(res HTTPResponse, rq HTTPRequest) {
			data := []byte(res.Body.String())

			value, _ := jsonparser.GetString(data, "hello")

			assert.Equal(t, "world", value)
			assert.Equal(t, http.StatusOK, res.Code)
		})
}

func TestHeader(t *testing.T) {
	r := New()

	r.GET("/text").
		SetHeader(H{
			"Content-Type": "text/plain",
			"Go-Version":   goVersion,
		}).
		Run(GojiMux(), func(res HTTPResponse, rq HTTPRequest) {

			assert.Equal(t, goVersion, rq.Header.Get("Go-Version"))
			assert.Equal(t, res.Body.String(), "Hello World")
			assert.Equal(t, http.StatusOK, res.Code)
		})
}

func TestQuery(t *testing.T) {
	r := New()

	r.GET("/query?text=world&foo=bar").
		Run(GojiMux(), func(res HTTPResponse, rq HTTPRequest) {
			data := []byte(res.Body.String())

			hello, _ := jsonparser.GetString(data, "hello")
			foo, _ := jsonparser.GetString(data, "foo")

			assert.Equal(t, "world", hello)
			assert.Equal(t, "bar", foo)
			assert.Equal(t, http.StatusOK, res.Code)
		})
}

func TestPostFormData(t *testing.T) {
	r := New()

	r.POST("/form").
		SetBody("a=1&b=2").
		Run(GojiMux(), func(res HTTPResponse, rq HTTPRequest) {
			data := []byte(res.Body.String())

			a, _ := jsonparser.GetString(data, "a")
			b, _ := jsonparser.GetString(data, "b")

			assert.Equal(t, "1", a)
			assert.Equal(t, "2", b)
			assert.Equal(t, http.StatusOK, res.Code)
		})
}

func TestPostJSONData(t *testing.T) {
	r := New()

	r.POST("/json").
		SetJSON(D{
			"a": 1,
			"b": 2,
		}).
		Run(GojiMux(), func(res HTTPResponse, rq HTTPRequest) {
			data := []byte(res.Body.String())

			a, _ := jsonparser.GetInt(data, "a")
			b, _ := jsonparser.GetInt(data, "b")

			assert.Equal(t, 1, int(a))
			assert.Equal(t, 2, int(b))
			assert.Equal(t, http.StatusOK, res.Code)
		})
}

func TestPut(t *testing.T) {
	r := New()

	r.PUT("/update").
		SetBody("c=1&d=2").
		Run(GojiMux(), func(res HTTPResponse, rq HTTPRequest) {
			data := []byte(res.Body.String())

			c, _ := jsonparser.GetString(data, "c")
			d, _ := jsonparser.GetString(data, "d")

			assert.Equal(t, "1", c)
			assert.Equal(t, "2", d)
			assert.Equal(t, http.StatusOK, res.Code)
		})
}

func TestDelete(t *testing.T) {
	r := New()

	r.DELETE("/delete").
		Run(GojiMux(), func(res HTTPResponse, rq HTTPRequest) {
			data := []byte(res.Body.String())

			hello, _ := jsonparser.GetString(data, "hello")

			assert.Equal(t, "world", hello)
			assert.Equal(t, http.StatusOK, res.Code)
		})
}

func TestPatch(t *testing.T) {
	r := New()

	r.PATCH("/patch").
		Run(GojiMux(), func(res HTTPResponse, rq HTTPRequest) {
			data := []byte(res.Body.String())

			value, _ := jsonparser.GetString(data, "hello")

			assert.Equal(t, "world", value)
			assert.Equal(t, http.StatusOK, res.Code)
		})
}

func TestHead(t *testing.T) {
	r := New()

	r.HEAD("/head").
		Run(GojiMux(), func(res HTTPResponse, rq HTTPRequest) {
			data := []byte(res.Body.String())

			value, _ := jsonparser.GetString(data, "hello")

			assert.Equal(t, "world", value)
			assert.Equal(t, http.StatusOK, res.Code)
		})
}

func TestOptions(t *testing.T) {
	r := New()

	r.OPTIONS("/options").
		Run(GojiMux(), func(res HTTPResponse, rq HTTPRequest) {
			data := []byte(res.Body.String())

			value, _ := jsonparser.GetString(data, "hello")

			assert.Equal(t, "world", value)
			assert.Equal(t, http.StatusOK, res.Code)
		})
}
