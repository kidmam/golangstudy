package main

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

func main() {
	net, err := ghw.Network()
	if err != nil {
		fmt.Printf("Error getting network info: %v", err)
	}

	fmt.Printf("%v\n", net)

	for _, nic := range net.NICs {
		fmt.Printf(" %v\n", nic)

		enabledCaps := make([]int, 0)
		for x, cap := range nic.Capabilities {
			if cap.IsEnabled {
				enabledCaps = append(enabledCaps, x)
			}
		}
		if len(enabledCaps) > 0 {
			fmt.Printf("  enabled capabilities:\n")
			for _, x := range enabledCaps {
				fmt.Printf("   - %s\n", nic.Capabilities[x].Name)
			}
		}
	}
}
