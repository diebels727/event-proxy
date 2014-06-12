package main

import (
  "net/http"
  "net/http/httputil"
  "net/url"
  "os"
  "log"
)

const (
  LOG_FORMAT = 3
  TARGET = "http://www.reddit.com/"
  LISTEN_PORT = ":9091"
)

var log_file *os.File
var logger *log.Logger

var datalog chan LogEvent // main communication path to the datalog service

type LogEvent struct {
  Event
  Timestamp string
  Name string
  URI string
}

func (event *LogEvent) toString() string {
  return event.Timestamp + "," + event.Name + "," + event.URI
}

type Event interface {
  toString() string
}

func Proxy(request *http.Request) {
  logger.Printf("[event-proxy] Intercepting ...")
  //rewrite request
  request.Host = "www.reddit.com"
  request.URL.Host = request.Host
  request.URL.Scheme = "http"

  //create a new event

  //ambitious:  map URL to actions for domain specific handling
    //e.g. PUT,/comments/1/posts/2 (params) -> UpdateCommentPostEvent
    //e.g. POST,/authors/2 (params) -> CreateNewAuthorEvent
    //URL actions can be done in the background (e.g. the in-flight log can be writing, other logs can be rewritten with events in O(N) )


  //in REST, we follow this convention: /resource/:id/ resource/:id/ resource/:id
    //decompose the URL
    //strip resource,id
    //build a JSON event chain



  //send the event to the event logger

  if (request.Method == "POST") {
    // 1. Create a new event and assign the relevant information
    urlString := request.URL.String()
    event := LogEvent{Name:"CreateEvent",URI: urlString}
    // 2. Send the event to the datalog service
    datalog <- event
  }

  logger.Printf("[event-proxy] Received HTTP request: %s",request.Method)
}

func main() {
  log_file,_ = os.Create("proxy.log")
  logger = log.New(log_file,"",LOG_FORMAT)

  logger.Printf("[event-proxy] Starting up ...")
  targetURL,err := url.Parse(TARGET)
  if err != nil {
    logger.Fatal("[event-proxy] Cannot parse target URL.")
  }

  datalog = make(chan LogEvent,1) // buffered channel of 1; prevent deadlocks while I experiment

  proxy := httputil.NewSingleHostReverseProxy(&url.URL{
    Scheme: targetURL.Scheme,
    Host:   targetURL.Host,
  })

  StartDataLog()


  proxy.Director = Proxy
  http.Handle("/", proxy)
  http.ListenAndServe(LISTEN_PORT, nil)
}