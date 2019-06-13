package model

import (
	"net"
)

func IsValidIpAddress(ipAddress string) bool {
	parsedIp := net.ParseIP(ipAddress)
	return parsedIp != nil
}

func ResolveHost(ipAddress string) []string {
	names, err := net.LookupAddr(ipAddress)
	if err != nil {
		return nil
	}
	return names
}