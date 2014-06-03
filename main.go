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
  logger.Printf("[event-proxy] Issuing GET.")
  //rewrite request
  //send request to target server

  resp,_ := http.Get(TARGET)
  defer resp.Body.Close()
  logger.Printf("[event-proxy] Issued GET.")
  //stash target response from target server
  //do something
  //write target response to my response
}

// func Proxy(response *http.ResponseWriter,request *http.Request) {

// }

type Proxy struct {}

func (p Proxy) ServeHTTP(response http.ResponseWriter,request *http.Request) {
  logger.Printf("[event-proxy] Intercepting ...")
  logger.Printf("[event-proxy] request: ",request)
}

func main() {
  log_file,_ = os.Create("proxy.log")
  logger = log.New(log_file,"",LOG_FORMAT)
  logger.Printf("[event-proxy] Starting up ...")

  router := mux.NewRouter()
  proxy := Proxy{}
  router.NotFoundHandler = proxy
  http.ListenAndServe(":9091",router)
}