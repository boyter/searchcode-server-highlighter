package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	addr := flag.String("addr", "127.0.0.1:8089", "HTTP network address")

	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(Home)).Methods("POST")

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  router,
	}

	infoLog.Println("Starting server on", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

// Possibly create routes to expose the below in the future
// fmt.Println(lexers.Names(true))
// fmt.Println(styles.Names())

// See below fod details
// https://github.com/alecthomas/chroma
func Home(w http.ResponseWriter, r *http.Request) {
	startTime := makeTimestampMilli()
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := "Error reading body " + err.Error()
		errorLog.Println(msg)
		w.WriteHeader(400)

		output, _ := json.Marshal(OutputError{
			Message: msg,
		})

		w.Write(output)
		return
	}

	var result InputLanguage
	err = json.Unmarshal(body, &result)

	if err != nil {
		msg := "Error reading JSON " + err.Error()
		errorLog.Println(msg)
		w.WriteHeader(400)

		output, _ := json.Marshal(OutputError{
			Message: msg,
		})

		w.Write(output)
		return
	}

	lexer := lexers.Match(result.FileName)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	style := styles.Get(result.Style)
	if style == nil {
		style = styles.Fallback
	}

	// Parse the content
	iterator, err := lexer.Tokenise(nil, result.Content)
	if err != nil {
		msg := "Error running tokeniser " + err.Error()
		errorLog.Println(msg)
		w.WriteHeader(500)

		output, _ := json.Marshal(OutputError{
			Message: msg,
		})

		_, _ = w.Write(output)
		return
	}

	var cssBytes bytes.Buffer
	var htmlBytes bytes.Buffer

	formatter := html.New(html.WithLineNumbers(), html.WithClasses())
	if formatter.WriteCSS(&cssBytes, style) != nil {
		msg := "Error writing CSS " + err.Error()
		errorLog.Println(msg)
		w.WriteHeader(500)

		output, _ := json.Marshal(OutputError{
			Message: msg,
		})

		_, _ = w.Write(output)
		return
	}

	if formatter.Format(&htmlBytes, style, iterator) != nil {
		msg := "Error writing HTML " + err.Error()
		errorLog.Println(msg)
		w.WriteHeader(500)

		output, _ := json.Marshal(OutputError{
			Message: msg,
		})

		_, _ = w.Write(output)
		return
	}

	output, _ := json.Marshal(OutputLanguage{
		Css:  cssBytes.String(),
		Html: htmlBytes.String(),
	})

	infoLog.Println("Processed in", (makeTimestampMilli() - startTime), "milliseconds", memUsage())
	w.WriteHeader(200)
	_, _ = w.Write(output)
}

// Helper to return the current unix time
func makeTimestampMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Get random memory stats to help spot leaks
func memUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	result := fmt.Sprintf("memoryusage::Alloc = %v MB::TotalAlloc = %v MB::Sys = %v MB::tNumGC = %v", bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
	return result
}

// Helper to convert bytes to megabytes
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

type InputLanguage struct {
	LanguageName    string `json:"languageName"`
	FileName        string `json:"fileName"`
	Style           string `json:"style"`
	Content         string `json:"content"`
	WithLineNumbers bool   `json:"withLineNumbers"`
}

type OutputLanguage struct {
	Css  string `json:"css"`
	Html string `json:"html"`
}

type OutputError struct {
	Message string
}
