package store

import (
	"github.com/comtom/portsService/ports"
)

type Memory struct {
	state map[string]ports.Port
}

func (m *Memory) Put(p ports.Port) error {
	m.state[p.Unlocs[0]] = p

	return nil
}

func (m *Memory) Get(unloc string) (ports.Port, error) {
	return m.state[unloc], nil
}

func (m *Memory) Shutdown() {

}

func NewMemStore(string) Database {
	return Database{}
}
