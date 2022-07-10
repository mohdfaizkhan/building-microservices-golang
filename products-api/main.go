package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mohdfaizkhan/building-microservices-golang/products-api/handlers"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

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

	ph := handlers.NewProducts(l)
	gb := handlers.NewGoodbye(l)

	sm := http.NewServeMux()

	sm.Handle("/", ph)
	sm.Handle("/goodbye", gb)

	// http.ListenAndServe(":8080", nil)

	//http.ListenAndServe(":8080", sm)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}

	}()

	sigchan := make(chan os.Signal)

	signal.Notify(sigchan, os.Interrupt)

	signal.Notify(sigchan, os.Kill)

	sig := <-sigchan
	l.Printf("Recived terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)

}
