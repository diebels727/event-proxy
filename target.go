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

func handleHTTP(response http.ResponseWriter,request *http.Request) {
  logger.Printf("[target] Received request: %s",request.Method)
}

func main() {
  log_file,_ = os.Create("target.log")
  logger = log.New(log_file,"",LOG_FORMAT)
  logger.Printf("[target] Starting up ...")

  m := mux.NewRouter()
  m.HandleFunc("/",handleHTTP)
  http.ListenAndServe(":5000",m)
}