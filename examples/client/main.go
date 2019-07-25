package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/iocn-io/mdns"
)

func main() {
	//serviceTag := "_iocn._tcp"
	serviceTag := "iocn"
	if len(os.Args) > 1 {
		serviceTag = os.Args[1]
	}

	// Make a channel for results and start listening
	entriesCh := make(chan *mdns.ServiceEntry, 8)
	defer close(entriesCh)

	go func() {
		/*
			for entry := range entriesCh {
				if len(entry.InfoFields) == 0 || entry.InfoFields[0] != serviceTag {
					continue
				}
				fmt.Printf("Got new entry: %v\n", entry)
			}
		*/
		for {
			var entry *mdns.ServiceEntry
			var ok bool
			select {
			case entry, ok = <-entriesCh:
				if !ok {
					return
				}
				if len(entry.InfoFields) == 0 || entry.InfoFields[0] != serviceTag {
					continue
				}
				fmt.Printf("Got new entry: %v\n", entry)
			}
		}
	}()

	// Start the lookups
	/*
		err := mdns.Lookup(serviceTag, entriesCh)
		if err != nil {
			fmt.Println(err)
		}
	*/
	/*
		var waitCtx context.CancelFunc
		params := mdns.DefaultParams(serviceTag)
		params.Entries = entriesCh
		for {
			params.Context, waitCtx = context.WithTimeout(context.Background(), time.Second*10)
			if err := mdns.Query(params); err != nil {
				fmt.Println("Query service error:", err)
			}
			waitCtx()
		}
	*/
	/*
		exitCh := make(chan struct{}, 1)
		if err := mdns.ListenExitChan(entriesCh, exitCh); err != nil {
			fmt.Println("Listen service error:", err)
		}
	*/
	ctx, _ := context.WithCancel(context.Background())
	if err := mdns.ListenContext(ctx, entriesCh); err != nil {
		fmt.Println("Listen Service error:", err)
	}

	wait()
}

func wait() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
