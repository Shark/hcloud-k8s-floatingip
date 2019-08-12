package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"

	"github.com/hetznercloud/hcloud-go/hcloud"
)

func main() {
	hcloudToken, ok := os.LookupEnv("HCLOUD_TOKEN")
	if !ok {
		log.Fatal("HCLOUD_TOKEN must be set")
	}

	floatingIPIDStr, ok := os.LookupEnv("FLOATING_IP_ID")
	if !ok {
		log.Fatal("FLOATING_IP_ID must be set")
	}

	floatingIPID, err := strconv.Atoi(floatingIPIDStr)
	if err != nil {
		log.Fatalf("FLOATING_IP_ID must be an integer: %#v", err)
	}

	thisServerName, ok := os.LookupEnv("THIS_SERVER_NAME")
	if !ok {
		log.Fatal("THIS_SERVER_NAME must be set")
	}

	client := hcloud.NewClient(hcloud.WithToken(hcloudToken))
	ctx := context.Background()

	floatingIP, _, err := client.FloatingIP.GetByID(ctx, floatingIPID)
	if err != nil {
		log.Fatalf("Error getting floating ip: %#v", err)
	}

	server, _, err := client.Server.GetByName(ctx, thisServerName)
	if err != nil {
		log.Fatalf("Error getting server: %#v", err)
	}

	action, _, err := client.FloatingIP.Assign(ctx, floatingIP, server)
	if err != nil {
		log.Fatalf("Error assigning floating ip: %#v", err)
	}

	successChan, errChan := client.Action.WatchProgress(ctx, action)
	select {
	case _ = <-successChan:
		log.Printf("Successfully assigned %s to %s", floatingIP.IP, server.Name)
	case err = <-errChan:
		log.Printf("Error assigning floating ip: %#v", err)
	}

	// Wait for sigint
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	var signalChannel chan os.Signal
	signalChannel = make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	go func() {
		log.Println("Sleeping")
		<-signalChannel
		waitGroup.Done()
	}()
	waitGroup.Wait()
}
