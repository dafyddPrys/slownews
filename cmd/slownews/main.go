package main

import (
	"flag"
	"fmt"

	slownews "github.com/dafyddprys/slownews/lib"
	"github.com/dafyddprys/slownews/lib/sources"
)

func main() {

	key := flag.String("key", "", "Key for guardian")
	flag.Parse()

	sources := []slownews.Source{
		sources.New(*key),
	}

	for _, s := range sources {
		err, _ := s.GetNews()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("OK?")
	}

}
