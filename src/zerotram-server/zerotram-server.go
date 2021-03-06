package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"os"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))
	log.Fatal(http.ListenAndServe(os.Getenv("OPENSHIFT_GO_PORT"), api.MakeHandler()))
}

