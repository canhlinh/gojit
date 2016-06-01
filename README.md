# Gojit

API Handler Testing for Goji framework.

## Installation
$ go get -u github.com/canhlinh/gojit

## Usage

Code:
```go

package main

import (
	"encoding/json"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func RenderJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func helloController(w http.ResponseWriter, r *http.Request) {
	RenderJSON(w, gojiContent{
		Hello: "world",
	})
}

func GojiMux() *goji.Mux {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello"), helloController)
    return mux
}

```

Unit test:

```go

package main

import (
	"net/http"
	"runtime"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
)

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

```