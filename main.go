package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/jozn/cassandra_walker/cwalker"
)

const (
	DEFAULT_CLUSTER_ADDRESS = "127.0.0.1"
	DEFAULT_PORT            = 9042
	DEFAULT_GO_PACKAGE_NAEM = "x"
	DEFAULT_OUTPUT          = "./walker_out"
)

func main() {
	args := &cwalker.ConfigArgs{}
	arg.MustParse(args)
	if args.Host == "" {
		args.Host = DEFAULT_CLUSTER_ADDRESS
	}

	if args.Port == 0 {
		args.Port = DEFAULT_PORT
	}

	if args.Package == "" {
		args.Package = DEFAULT_GO_PACKAGE_NAEM
	}

	if args.Dir == "" {
		args.Dir = DEFAULT_OUTPUT
	}

	cwalker.Runner(args)

	fmt.Println(args)
}
