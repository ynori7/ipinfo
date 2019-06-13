package model

import (
	"net"
	"net/http"
	"strings"
)

type IP struct {
	IpAddress    string
	ForwardedFor string
}

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

func GetIpFromRequest(r *http.Request) IP {
	ip := IP{ForwardedFor: r.Header.Get("X-Forwarded-For")}

	ipParts := strings.Split(r.RemoteAddr, ":")
	if len(ipParts) > 0 {
		ip.IpAddress = ipParts[0]
	}

	return ip
}
