package main

import (
	"github.com/knowit/gogo-gopher/cruddy/internal/api"
	"github.com/knowit/gogo-gopher/cruddy/internal/logging"
	"github.com/knowit/gogo-gopher/cruddy/internal/signals"
	"go.uber.org/zap"
)

func main() {
	logger, _ := logging.InitZap()
	defer logger.Sync()
	stdLog := zap.RedirectStdLog(logger)
	defer stdLog()

	logger.Info("Starting cruddy",
		zap.String("version", "1.0.0"),
		zap.String("revision", "1"),
	)

	srv, err := api.NewServer(logger)
	if err != nil {
		panic("SERVER CRASH!")
	}

	httpServer, httpsServer, healthy, ready := srv.ListenAndServe()

	var srvConfig api.Config
	stopCh := signals.SetupSignalHandler()
	sd, _ := signals.NewShutdown(srvConfig.ServerShutdownTimeout, logger)
	sd.Graceful(stopCh, httpServer, httpsServer, healthy, ready)
}
