package gojit

import (
	"encoding/json"
	"net/http"

	"goji.io"

	"goji.io/pat"
)

type gojiContent struct {
	Hello string `json:"hello"`
	Foo   string `json:"foo"`
	A     string `json:"a"`
	B     string `json:"b"`
	C     string `json:"c"`
	D     string `json:"d"`
}

// Binding from JSON
type gojiJSONContent struct {
	A int `json:"a" binding:"required"`
	B int `json:"b" binding:"required"`
}

func RenderJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func RenderText(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func GojiMux() *goji.Mux {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello"), helloController)
	mux.HandleFunc(pat.Get("/text"), textController)
	mux.HandleFunc(pat.Get("/query"), querryController)

	mux.HandleFunc(pat.Post("/form"), postFormController)
	mux.HandleFunc(pat.Post("/json"), jsonController)
	mux.HandleFunc(pat.Put("/update"), putController)
	mux.HandleFunc(pat.Delete("/delete"), deleteController)

	mux.HandleFunc(pat.Patch("/patch"), helloController)
	mux.HandleFunc(pat.Options("/options"), helloController)
	mux.HandleFunc(pat.Head("/head"), helloController)
	return mux
}

func helloController(w http.ResponseWriter, r *http.Request) {
	RenderJSON(w, gojiContent{
		Hello: "world",
	})
}

func textController(w http.ResponseWriter, r *http.Request) {
	RenderText(w, "Hello World")
}

func querryController(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	foo := r.FormValue("foo")
	RenderJSON(w, gojiContent{
		Hello: text,
		Foo:   foo,
	})
}

func postFormController(w http.ResponseWriter, r *http.Request) {
	a := r.FormValue("a")
	b := r.FormValue("b")
	RenderJSON(w, gojiContent{
		A: a,
		B: b,
	})
}

func jsonController(w http.ResponseWriter, req *http.Request) {
	jsonContent := new(gojiJSONContent)
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(jsonContent)
	if err != nil {
		panic(err)
	}
	RenderJSON(w, jsonContent)
}

func putController(w http.ResponseWriter, r *http.Request) {
	foo := r.FormValue("c")
	bar := r.FormValue("d")
	RenderJSON(w, gojiContent{
		C: foo,
		D: bar,
	})
}

func deleteController(w http.ResponseWriter, r *http.Request) {
	RenderJSON(w, &gojiContent{
		Hello: "world",
	})
}
