package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "net/url"
  "os"
  "log"
  "io/ioutil"
  "fmt"
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
  logger.Printf("[event-proxy] request.Proto: %s",request.Proto)
  logger.Printf("[event-proxy] request.ProtoMajor: %d",request.ProtoMajor)
  logger.Printf("[event-proxy] request.ProtoMinor: %d",request.ProtoMinor)
  logger.Printf("[event-proxy] request.URL.Scheme: %s",request.URL.Scheme)

  //rewrite host
  request.Host = TARGET
  newRequest,err := http.NewRequest(request.Method,TARGET,nil)
  if err != nil {
    logger.Printf("[event-proxy] error generating new request")
  }

  logger.Printf("[event-proxy] newRequest: %s",newRequest)
  logger.Printf("[event-proxy] newRequest.Host: %s",newRequest.Host)
  logger.Printf("[event-proxy] newRequest.RequestURI: %s",newRequest.RequestURI)
  logger.Printf("[event-proxy] newRequest.RemoteAddr: %s",newRequest.RemoteAddr)
  logger.Printf("[event-proxy] newRequest.Proto: %s",newRequest.Proto)
  logger.Printf("[event-proxy] newRequest.ProtoMajor: %d",newRequest.ProtoMajor)
  logger.Printf("[event-proxy] newRequest.ProtoMinor: %d",newRequest.ProtoMinor)
  logger.Printf("[event-proxy] newRequest.URL: %s",newRequest.URL)
  logger.Printf("[event-proxy] newRequest.URL.Scheme: %s",newRequest.URL.Scheme)


  logger.Printf("[event-proxy] request.URL.Scheme %s",request.URL.Scheme)
  logger.Printf("[event-proxy] request.URL.Opaque %s",request.URL.Opaque)
  logger.Printf("[event-proxy] request.URL.User %s",request.URL.User)
  logger.Printf("[event-proxy] request.URL.Host %s",request.URL.Host)
  logger.Printf("[event-proxy] request.URL.Path %s",request.URL.Path)
  logger.Printf("[event-proxy] request.URL.RawQuery %s",request.URL.RawQuery)
  logger.Printf("[event-proxy] request.URL.Fragment %s",request.URL.Fragment)

  logger.Printf("[event-proxy] newRequest.URL.Scheme %s",newRequest.URL.Scheme)
  logger.Printf("[event-proxy] newRequest.URL.Opaque %s",newRequest.URL.Opaque)
  logger.Printf("[event-proxy] newRequest.URL.User %s",newRequest.URL.User)
  logger.Printf("[event-proxy] newRequest.URL.Host %s",newRequest.URL.Host)
  logger.Printf("[event-proxy] newRequest.URL.Path %s",newRequest.URL.Path)
  logger.Printf("[event-proxy] newRequest.URL.RawQuery %s",newRequest.URL.RawQuery)
  logger.Printf("[event-proxy] newRequest.URL.Fragment %s",newRequest.URL.Fragment)


  targetURL,err := url.Parse(TARGET)
  if err != nil {
    logger.Printf("[event-proxy] error parsing target URL.")
  }
  //rewrite request
  request.RemoteAddr = ""
  request.RequestURI = ""
  request.URL.Scheme = targetURL.Scheme
  request.URL.Host = targetURL.Host


  client := &http.Client{}
  // resp,err := client.Do(newRequest)
  targetResponse,err := client.Do(request)
  if err != nil {
    logger.Printf("[event-proxy] error: %s",err)
  }

  body,err := ioutil.ReadAll(targetResponse.Body)
  if err != nil {
    logger.Printf("[event-proxy] error: %s",err)
  }
  fmt.Fprintf(response,string(body))
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