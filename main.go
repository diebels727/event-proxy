package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "log"
)

const (
  LOG_FORMAT = 3
  TARGET = "http://localhost:5000/"
)

var log_file *os.File
var logger *log.Logger

func Relay(response http.ResponseWriter,req *http.Request) {
  //forward request to target server
  logger.Printf("[event-proxy] Issuing GET.")
  resp,_ := http.Get(TARGET)
  defer resp.Body.Close()
  logger.Printf("[event-proxy] Issued GET.")
  //stash target response from target server
  //do something
  //write target response to my response
}

func main() {
  log_file,_ = os.Create("proxy.log")
  logger = log.New(log_file,"",LOG_FORMAT)
  logger.Printf("[event-proxy] Starting up ...")

  m := mux.NewRouter()
  m.HandleFunc("/",Relay).Methods("GET")
  http.ListenAndServe(":9091",m)
}