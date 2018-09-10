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

)

func main() {
	//cwalker.Runner()
	args := &cwalker.ConfigArgs{}
	arg.MustParse(args)
	fmt.Println(args)
}
