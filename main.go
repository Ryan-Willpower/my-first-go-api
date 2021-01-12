package main

import (
	"encoding/json"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type requestStatus struct {
	Status string `json:"status"`
}

// var (
// 	contentType     = []byte("content-type")
// 	applicationJSON = []byte("application/json")
// )

func index(ctx *fasthttp.RequestCtx) {
	log.Println("GET /")

	response := requestStatus{
		Status: "ok!",
	}

	// ctx.Response.Header.SetCanonical(contentType, applicationJSON)
	ctx.Response.Header.SetContentType("application/json")

	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		panic(err)
	}
}

func main() {
	r := router.New()

	r.GET("/", index)

	log.Println("> Server is started at http://localhost:8000")
	log.Fatal(fasthttp.ListenAndServe(":8000", r.Handler))
}
