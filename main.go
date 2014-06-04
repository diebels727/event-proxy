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
  logger.Printf("[event-proxy] request.Body: %s",request.Body)
  logger.Printf("[event-proxy] request.Header: %s",request.Header)
  logger.Printf("[event-proxy] request.URL: %s",request.URL)
  logger.Printf("[event-proxy] request.Method: %s",request.Method)

  logger.Printf("[event-proxy] request.Host: %s",request.Host)

  //probably for IPv6?  Need to rewrite?
  logger.Printf("[event-proxy] request.RemoteAddr: %s",request.RemoteAddr)

  //cannot be set in a client request
  logger.Printf("[event-proxy] request.RequestURI: %s",request.RequestURI)

  //need to generate a new request

  //rewrite host
  request.Host = TARGET
  newRequest,err := http.NewRequest(request.Method,TARGET,nil)
  if err != nil {
    logger.Printf("[event-proxy] error generating new request")
  }

  client := &http.Client{}
  resp,err := client.Do(newRequest)
  if err != nil {
    logger.Printf("[event-proxy] error: %s",err)
  }
  logger.Printf("[event-proxy] response: %s",resp)
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