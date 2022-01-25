//dependencies
import (
	"log"
	"net/http"
	"context"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
)

//code
var sub = &common.Subscription{
	PubsubName: "order_pub_sub",
	Topic:      "orders",
	Route:      "/checkout",
}

func main() {
	s := daprd.NewService(":6002")
	//Subscribe to a topic
	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning: %v", err)
	}
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("Subscriber received: %s", e.Data)
	return false, nil
}

