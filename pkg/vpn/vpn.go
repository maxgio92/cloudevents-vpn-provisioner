package vpn

import (
	"context"
	"log"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
)

const (
	Source = "knative/eventing/vpn/provisioner"
)

func Receive(ctx context.Context, event cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	// Here is where your code to process the event will go.
	// In this example we will log the event msg
	log.Printf("Event received. \n%s\n", event)

	switch event.Type() {
	case TypeVpnPending:
		eventData := &VPNPending{}
		if err := event.DataAs(eventData); err != nil {
			log.Printf("Error while extracting cloudevent Data: %s\n", err.Error())

			return nil, cloudevents.NewHTTPResult(400, "failed to convert data: %s", err)
		}
		log.Printf("A new VPN is ready, provided by: %s", eventData.Provider)
		log.Printf("A new VPN request is now ready to be satisfied")

		respEvent, err := doProvisioning(eventData.Provider)
		if err != nil {
			return nil, cloudevents.NewHTTPResult(500, "failed to set Crawling response data: %s", err)
		}
		log.Printf("Crawling done")

		return respEvent, nil
	}

	return nil, nil
}

func NewPendingEvent(provider VPNProvider, source string) (*event.Event, error) {
	data := VPNPending{
		Provider: provider,
	}

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(source)
	e.SetType(TypeVpnPending)
	if err := e.SetData(cloudevents.ApplicationJSON, data); err != nil {
		return nil, err
	}

	return &e, nil
}

func doProvisioning(provider VPNProvider) (*event.Event, error) {

	// Insert logic here.
	time.Sleep(time.Second * 5)

	return buildCrawlingDoneEvent(provider)
}

func buildCrawlingDoneEvent(provider VPNProvider) (*event.Event, error) {
	crawlingDone := VPNReady{
		Provider: provider,
	}

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(Source)
	e.SetType(TypeVpnReady)
	if err := e.SetData(cloudevents.ApplicationJSON, crawlingDone); err != nil {
		return nil, err
	}

	return &e, nil
}
