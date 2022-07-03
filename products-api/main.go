package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mohdfaizkhan/building-microservices-golang/main/products-api/handlers"
)

func main() {

	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

	// 	log.Println("Hello World")
	// 	data, err := ioutil.ReadAll(req.Body)

	// 	if err != nil {
	// 		http.Error(res, "Oops", http.StatusBadGateway)
	// 		return
	// 	}

	// 	log.Printf("data %s\n", data)

	// 	fmt.Fprint(res, "Hello %S", data)

	// })

	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {

	// 	log.Println("Goodbye World")
	// })

	l := log.New(os.Stdout, "product api", log.LstdFlags)

	hh := handlers.NewHello(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)

	// http.ListenAndServe(":8080", nil)

	http.ListenAndServe(":8080", sm)

}
