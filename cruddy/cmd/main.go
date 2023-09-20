package main

import (
  "go.uber.org/zap"
  "github.com/knowit/gogo-gopher/cruddy/internal/api"
  "github.com/knowit/gogo-gopher/cruddy/internal/logging"
)

func main() {
  logger, _ := logging.InitZap()
  defer logger.Sync()
  stdLog := zap.RedirectStdLog(logger)
  defer stdLog()

  srv, err := api.NewServer(logger)
  if err != nil {
    panic("SERVER CRASH!")
  }
  srv.ListenAndServe()
}
