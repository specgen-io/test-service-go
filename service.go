package main

import (
	"flag"
	"fmt"
	"github.com/husobee/vestigo"
	"log"
	"net/http"
	"test-service-go/spec"
)

func main() {
	port := flag.String("port", "8081", "port number")
	flag.Parse()

	router := vestigo.NewRouter()

	router.Get("/", func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin:      []string{"*", "*"},
	})

	spec.AddEchoRoutes(router, &spec.EchoService{})
	spec.AddCheckRoutes(router, &spec.CheckService{})

	fmt.Println("Starting service on port: "+*port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}