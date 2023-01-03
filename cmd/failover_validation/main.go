package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	nodeID := flag.Int("node-id", 0, "the node id")
	nodeName := flag.String("node-name", "", "the name of the proposed leader")
	visibleNodes := flag.Int("visable-nodes", 0, "Total visible nodes from the perspective of the proposed leader")
	totalNodes := flag.Int("total-nodes", 0, "The total number of nodes registered")
	flag.Parse()

	fmt.Printf("Verifying failover candidate %d: %s\n", nodeID, *nodeName)

	// If there are no visible nodes, then we can't accept leadership as we are not able to
	// confirm a network partition.
	if *visibleNodes == 0 {
		fmt.Println("Zero visible nodes detected. Enabling read-only mode.")
		os.Exit(1)
	}

	// TODO - This will remove HA from a 2-node cluster setup until we can workout how to
	// differentiate between a down node and a network partition.

	// We have visible nodes, but not enough to meet quorum.
	if *visibleNodes < (*totalNodes/2 + 1) {
		fmt.Printf("Quorum not met. Total nodes: %d, Visible nodes: %d\n", *totalNodes, *visibleNodes)
		os.Exit(1)
	}

	os.Exit(0)
}
