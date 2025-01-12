package main

import (
	_ "github.com/coredns/coredns/core/plugin"
	_ "github.com/jwhited/wgsd"
	_ "github.com/machsix/coredns-addition"
	_ "github.com/openshift/coredns-mdns/v4"
	_ "github.com/relekang/coredns-blocklist"

	"fmt"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
)

func init() {
	// Function to insert an element after another
	insertAfter := func(list []string, after string, item string) []string {
		for i, v := range list {
			if v == after {
				return append(list[:i+1], append([]string{item}, list[i+1:]...)...)
			}
		}
		panic(fmt.Sprintf("failed to insert %q after %q: not found", item, after))
	}

	// Insert "blocklist" after "log"
	dnsserver.Directives = insertAfter(dnsserver.Directives, "log", "blocklist")

	// Insert "wgsd" after "file"
	dnsserver.Directives = insertAfter(dnsserver.Directives, "file", "wgsd")

	// Insert "mdns" after "cache"
	dnsserver.Directives = insertAfter(dnsserver.Directives, "cache", "mdns")

	// Insert "addition" after "hosts"
	dnsserver.Directives = insertAfter(dnsserver.Directives, "hosts", "addition")
}

func main() {
	coremain.Run()
}
