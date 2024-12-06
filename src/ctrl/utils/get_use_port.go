package utils

import (
	"fmt"
	"net"
)

var Startport = 30000

var Endport = 60000

func GetUsePort() (int, error) {
	for {
		Startport++
		if Startport == Endport {
			Startport = 30000
		}
		address := fmt.Sprintf(":%d", Startport)
		listener, err := net.Listen("tcp", address)
		if err == nil {
			listener.Close()
			return Startport, nil
		}

	}
}
