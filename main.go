package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(Home)).Methods("POST")

	srv := &http.Server{
		Addr:     ":8080",
		ErrorLog: errorLog,
		Handler:  router,
	}

	infoLog.Println("Starting server on :8080")
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}


func Home(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	var result InputLanguage
	err = json.Unmarshal(body, &result)

	if err != nil {
		return
	}

	fmt.Println(result)
}


type InputLanguage struct {
	Name string
	Content string
}