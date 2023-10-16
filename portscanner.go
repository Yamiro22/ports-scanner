package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

func main() {
	const host = "127.0.0.1" // Change this to the target host IP
	const startPort = 1
	const endPort = 1024
	const timeout = 2 * time.Second

	var wg sync.WaitGroup
	openPorts := make([]int, 0)

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", host, p)
			conn, err := net.DialTimeout("tcp", address, timeout)
			if err == nil {
				openPorts = append(openPorts, p)
				conn.Close()
			}
		}(port)
	}

	wg.Wait()

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("Port %d is open\n", port)
	}
}
