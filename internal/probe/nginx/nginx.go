package nginx

import "math/rand"

type NginxProbe struct {
}

func (n NginxProbe) Run(host string) bool {
	rn := rand.Intn(2)
	return rn == 1
}
