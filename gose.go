package main

import (
	"flag"
	"fmt"
	"gose/entries"
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Println("Usage: vtool has many tools")
}

func main() {
	var e = flag.String("e", "", "which program")
	var a = flag.String("a", "", "which algorithm to use for sorting")
	var i = flag.String("i", "", "Specify input file")
	var o = flag.String("o", "", "Specify output file")
	flag.Parse()
	if *e == "sort" {
		err := entries.Run(*a, *i, *o)
		if nil != err {
			fmt.Println(err)
		}else {
			fmt.Println("done!")
		}
	}else {
		fmt.Println("nothing")
	}
}
