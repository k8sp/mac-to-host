package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	flagHttp := flag.String("http", "", "Listening address if run an HTTP server")
	flag.Parse()

	if len(*flagHttp) == 0 {
		fmt.Println(Hostname(flag.Arg(0)))
	}
}

func Hostname(mac string) string {
	ret := ""

	fs := strings.Split(strings.ToLower(mac), "-")
	if len(fs) != 6 {
		fs = strings.Split(strings.ToLower(mac), ":")
	}
	if len(fs) != 6 {
		log.Panicf("MAC address %s not separated by : or -", mac)
	}

	for i := 0; i < len(fs); i++ {
		if len(fs[i]) != 2 {
			log.Panicf("Invalid the %d-th segment %s of MAC address %s", i, fs[i], mac)
		}

		v, e := strconv.ParseInt(fs[i], 16, 0)
		if e != nil {
			log.Panicf("Cannot parse segment %s of MAC address %s", fs[i], mac)
		}

		if i == 0 {
			ret += adverbs[v] + "-"
		} else if i == len(fs)-1 {
			ret += nouns[v]
		} else {
			ret += adjectives[v] + "-"
		}
	}

	return ret
}
