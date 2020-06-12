package main

import(
	"context"
	"fmt"
	"github.com/wish/discovery"
	"github.com/wish/discovery/resolver"
)

type Discovery struct {
	c discovery.DiscoveryConfig
	r *resolver.DNSResolver
}

func main() {
	config, _ := discovery.ConfigFromEnv()
	d, _ := discovery.NewDiscovery(*config)

	addrs, _ := d.GetServiceAddresses(context.TODO(),"mp-config.01db.bjs.i.wish.com")
	fmt.Println(addrs)

	addr, _ := d.GetServiceAddresses(context.TODO(), "127.0.0.1:1234")
	fmt.Println(addr)
}
