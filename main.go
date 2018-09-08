package main

import (
	"github.com/jozn/cassandra_walker/cwalker"
    "github.com/alexflint/go-arg"
    "fmt"
)

func main() {
	//cwalker.Runner()
	args := &cwalker.ConfigArgs{}
	arg.MustParse(args)
	fmt.Println(args)""
}
