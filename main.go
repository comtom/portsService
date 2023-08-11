package main

import (
	"os"

	"github.com/comtom/portsService/config"
	"github.com/comtom/portsService/logger"
	"github.com/comtom/portsService/ports"
	"github.com/comtom/portsService/signals"
	"github.com/comtom/portsService/store"
	"go.uber.org/zap"
)

type portsService struct {
	logger *logger.Logger
	store  store.Storage
}

func main() {
	// setup service
	conf := config.GetConfig()

	l := logger.NewLogger()
	s, err := store.NewDBStore(conf.DBHost, conf.DBPort, conf.DBUsername, conf.DBPassword, conf.DBname, l)
	service := portsService{
		logger: l,
		store:  s,
	}

	// signal handling
	go signals.InitSignals(service.logger)

	// load file
	f, err := os.Open(conf.IngestFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	portsMap := ports.LoadPorts(f)

	// present the data
	// for now we just print basic port data
	if len(portsMap) == 0 {
		panic("no ports loaded! shutting down service")
	}

	service.logger.Info("loaded ports")

	if err != nil {
		panic("connection to db failed. aborting ingestion")
	}

	for k, port := range portsMap {
		service.logger.Info("port upserted", zap.String("unilocs", k), zap.String("name", port.Name), zap.String("country", port.Country))
		err = service.store.Put(port)
		if err != nil {
			service.logger.Error("error while ingesting... halting the process", zap.Error(err))
			return
		}
	}

	service.logger.Info("ingestion finished succesfully")
}
