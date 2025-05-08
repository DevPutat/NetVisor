package network

import "net"

func NetInterfaces() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	res := []string{}
	for _, iface := range interfaces {
		res = append(res, iface.Name)
	}
	return res, nil
}
