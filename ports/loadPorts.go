package ports

import (
	"encoding/json"
	"fmt"
	"io"
)

func LoadPorts(data io.Reader) (ports map[string]Port) {
	err := json.NewDecoder(data).Decode(&ports)
	if err != nil {
		fmt.Println(err)
	}

	return ports
}
