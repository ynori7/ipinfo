package model

import (
	"net"
	"net/http"
	"net/url"
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

func GetHostnamesByIp(ipAddress string) []string {
	names, err := net.LookupAddr(ipAddress)
	if err != nil {
		return nil
	}
	return names
}

func GetHostnamesByIpList(ipAddresses []string) []string {
	var names []string
	for _, ip := range ipAddresses {
		hosts := GetHostnamesByIp(ip)
		if hosts != nil {
			names = append(names, hosts...)
		}
	}
	return names
}

func GetIpsByHostname(hostname string) []string {
	names, err := net.LookupHost(hostname)
	if err != nil {
		return nil
	}
	return names
}

func IsValidHostname(hostname string) bool {
	parsedUrl, err := url.Parse(hostname)
	if err != nil {
		return false
	}
	return hostname != parsedUrl.Hostname()
}

func GetIpFromRequest(r *http.Request) IP {
	ip := IP{ForwardedFor: r.Header.Get("X-Forwarded-For")}

	ipParts := strings.Split(r.RemoteAddr, ":")
	if len(ipParts) > 0 {
		ip.IpAddress = ipParts[0]
	}

	return ip
}
