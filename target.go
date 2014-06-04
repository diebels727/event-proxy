package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "log"
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

func handleHTTP(response http.ResponseWriter,request *http.Request) {
  logger.Printf("[target] Received request: %s",request.Method)
}

func main() {
  log_file,_ = os.Create("target.log")
  logger = log.New(log_file,"",LOG_FORMAT)
  logger.Printf("[target] Starting up ...")

  router := mux.NewRouter()
  // m.HandleFunc("/",handleHTTP)
  serveFunc := Server{}
  router.NotFoundHandler = serveFunc
  http.ListenAndServe(":5000",router)
}