package xnet

import (
	"net"
	"net/http"
)

func RemoteIp(req *http.Request) string {
	remoteAddr := req.Header.Get("X-Forwarded-For")
	var ip string
	if remoteAddr == "" {
		remoteAddr = req.RemoteAddr
		ip, _, _ = net.SplitHostPort(remoteAddr)
	} else {
		ip = remoteAddr
	}
	return ip
}
