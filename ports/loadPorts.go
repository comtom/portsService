package ports

import (
	"encoding/json"
	"fmt"
	"io"
)

func LoadPorts(data io.Reader) (ports map[string]Port) {
	// TODO: instead of decoding the whole file, traverse it decoding one item at a time
	err := json.NewDecoder(data).Decode(&ports)
	if err != nil {
		fmt.Println(err)
	}

	return ports
}
