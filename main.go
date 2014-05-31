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

func Relay(response * ,request * ) {
  //forward request to target server
  //stash target response from target server
  //do something
  //write target response to my response
}

func main() {
  log_file,_ = os.Create("proxy.log")
  logger = log.New(log_file,"",LOG_FORMAT)
  logger.Printf("[event-proxy] Starting up ...")

  go func(){
    logger.Printf("[main] Launching main event listener")
    for {
      //do something in the background
    }
    logger.Printf("[main] Terminating main event listener ...")
  }();

  m := mux.NewRouter()
  m.HandleFunc("/records",Relay).Methods("GET")
  http.ListenAndServe(":9091",m)
}