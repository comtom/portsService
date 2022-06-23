package ports

import "io"

func LoadPorts(io.Reader) map[string]Port {
	return map[string]Port{}
}
