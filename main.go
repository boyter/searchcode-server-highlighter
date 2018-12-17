package main

import (
	"encoding/json"
	"fmt"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/quick"
	"github.com/alecthomas/chroma/styles"
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
		fmt.Println("Error reading body", err.Error())
		return
	}

	var result InputLanguage
	err = json.Unmarshal(body, &result)

	if err != nil {
		fmt.Println("Error reading to json", err.Error())
		return
	}

	lexer := lexers.Match("foo.go")
	fmt.Println(lexer)
	lexer = lexers.Get("go")
	fmt.Println(lexer)

	if lexer == nil {
		lexer = lexers.Fallback
	}

	fmt.Println(lexers.Names(true))


	style := styles.Get("swapoff")
	if style == nil {
		style = styles.Fallback
	}
	formatter := formatters.Get("html")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	iterator, err := lexer.Tokenise(nil, result.Content)
	formatter.Format(w, style, iterator)


	fmt.Println(result)
	quick.Highlight(os.Stdout, result.Content, "go", "html", "monokai")
}


type InputLanguage struct {
	LanguageName string
	FileName string
	Content string
}