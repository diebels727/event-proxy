package main

import (
  "log"
  "os"
)

func StartDataLog() {
  go func() {
    datalog_file,_ := os.Create("datalog.log")
    datalogger := log.New(datalog_file,"",LOG_FORMAT)
    logger.Printf("[datalog] Starting up ...")
    for {
      event := <- datalog
      datalogger.Printf("[datalog] [event: %s]",event.toString())
    }
  }()
  logger.Printf("[datalog] Shutting down ...")
}