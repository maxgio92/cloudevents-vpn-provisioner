package main

import (
	"context"
	"log"

	"github.com/maxgio92/cloudevents-vpn-provisioner/pkg/vpn"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func main() {
	log.Print("VPN provisioner started.")
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), vpn.Receive))
}
