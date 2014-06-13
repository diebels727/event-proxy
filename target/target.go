package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "log"
  "fmt"
)

const (
  LOG_FORMAT = 3
)

var log_file *os.File
var logger *log.Logger

type Server struct {}

func (s Server) ServeHTTP(responseWriter http.ResponseWriter,request *http.Request) {
  logger.Printf("[target] Received request: %s",request.Method)
}


func ArtificialBodyHandler(responseWriter http.ResponseWriter,request *http.Request) {
  logger.Printf("[target::artificial] Received request: %s",request.Method)
  fmt.Fprintf(responseWriter,"Fabricated response!")
}

func main() {
  log_file,_ = os.Create("target.log")
  logger = log.New(log_file,"",LOG_FORMAT)
  logger.Printf("[target] Starting up ...")

  router := mux.NewRouter()
  serveFunc := Server{}
  router.NotFoundHandler = serveFunc
  router.HandleFunc("/artificial",ArtificialBodyHandler)
  http.ListenAndServe(":5000",router)
}