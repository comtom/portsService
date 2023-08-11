package store

import "github.com/comtom/portsService/ports"

type Storage interface {
	Put(ports.Port) error
	Get(unloc string) (ports.Port, error)
	Shutdown()
}
