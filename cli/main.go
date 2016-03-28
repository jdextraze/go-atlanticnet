package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/jdextraze/go-atlanticnet"
	"log"
)

func main() {
	var (
		apiKey string
		secret string
		debug  bool
		output interface{}
		err    error
	)

	flag.StringVar(&apiKey, "apiKey", "", "Api Key")
	flag.StringVar(&secret, "secret", "", "Secret")
	flag.BoolVar(&debug, "debug", false, "Debug")
	flag.Parse()

	if apiKey == "" || secret == "" {
		flag.Usage()
		return
	}

	client := atlanticnet.NewClient(apiKey, secret, debug)

	// TODO usage
	switch flag.Arg(0) {
	case "DescribePlan":
		output, err = client.DescribePlan(flag.Arg(1), flag.Arg(2))
	case "ListInstances":
		output, err = client.ListInstances()
	case "RunInstance":
		if flag.NArg() != 5 {
			log.Fatalln("Invalid number of arguments")
		}
		output, err = client.RunInstance(flag.Arg(1), flag.Arg(2), flag.Arg(3), flag.Arg(4))
	case "TerminateInstance":
		if flag.NArg() != 2 {
			log.Fatalln("Invalid number of arguments")
		}
		output, err = client.TerminateInstance(flag.Arg(1))
	case "DescribeInstance":
		if flag.NArg() != 2 {
			log.Fatalln("Invalid number of arguments")
		}
		output, err = client.DescribeInstance(flag.Arg(1))
	case "RebootInstance":
		if flag.NArg() < 2 || flag.NArg() > 3 {
			log.Fatalln("Invalid number of arguments")
		}
		output, err = client.RebootInstance(flag.Arg(1), atlanticnet.RebootType(flag.Arg(2)))
	case "DescribeImage":
		output, err = client.DescribeImage(flag.Arg(1))
	case "ListSshKeys":
		output, err = client.ListSshKeys()
	default:
		err = errors.New("Invalid action")
	}

	if err != nil {
		log.Fatalln(err)
	}

	bytes, _ := json.MarshalIndent(output, "", "  ")
	fmt.Println(string(bytes))
}
