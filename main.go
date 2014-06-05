package main

import (
  // "github.com/gorilla/mux"

  "net/http"
  "net/http/httputil"
  "net/url"

  "os"
  "log"
  // "io/ioutil"
  // "fmt"
)

const (
  LOG_FORMAT = 3
  TARGET = "http://www.reddit.com/"
  LISTEN_PORT = ":9091"
)

var log_file *os.File
var logger *log.Logger

func Proxy(request *http.Request) {
  request.Host = "www.reddit.com"
  request.URL.Host = request.Host
  request.URL.Scheme = "http"
}

func main() {
  log_file,_ = os.Create("proxy.log")
  logger = log.New(log_file,"",LOG_FORMAT)
  logger.Printf("[event-proxy] Starting up ...")

  targetURL,err := url.Parse(TARGET)
  if err != nil {
    logger.Printf("[event-proxy] Cannot parse target URL.")
  }

  proxy := httputil.NewSingleHostReverseProxy(&url.URL{
    Scheme: targetURL.Scheme,
    Host:   targetURL.Host,
  })

  proxy.Director = Proxy
  http.Handle("/", proxy)
  http.ListenAndServe(LISTEN_PORT, nil)
}