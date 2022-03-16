package network

import (
	"net"

	"github.com/lovelock/toolkit/handlers"
)

func GetIpsFromHosts(hosts []string) []string {
	var res []string
	for _, host := range hosts {
		ips, err := net.LookupHost(host)
		if err != nil {
			handlers.LogErrorAndGoon(err)
			continue
		} else {
			res = append(res, ips...)
		}
	}

	return res
}
