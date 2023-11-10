/******** Peter Winzell (c), 8/31/23 *********************************************/

package broker

import (
	"fmt"
	"github.com/petervolvowinz/viss-rl-interfaces/base"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestBrokerInfo(t *testing.T) {
	test := SeverConfig()
	if test != nil {
		fmt.Println(test)
		t.Error("test failed")
	}
}

func TestStreaming(t *testing.T) {

	subscriber, _, settings, error := GetBrokerConnections()
	if error != nil {
		t.Error(error)
	}

	resp, err := StartStreaming(subscriber, settings)
	if err != nil {
		t.Error(err)
	}

	for i := 1; i < 5; i++ {
		msg, err := resp.Recv() // wait for a subscription msg
		if err != nil {
			fmt.Println(err)
			log.Debug(" error ", err)
			// t.Error(err)
		} else {
			for _, asignal := range msg.Signal {
				switch asignal.Payload.(type) {
				case *base.Signal_Double:
					f64val := asignal.GetDouble()
					fmt.Println(asignal.Id.Name, " ", f64val)
				}
			}
		}

	}
}

type valueChannel struct {
	Name  string
	Value any
}

func TestPublishSignals(t *testing.T) {
	_, publisher, settings, error := GetBrokerConnections()
	if error != nil {
		t.Error(error)
	}

	writerChannel := make(chan valueChannel, 1)
	go func() {
		for {
			val := &valueChannel{
				Name:  "Vehicle.Body.Lights.IsLeftIndicatorOn",
				Value: true,
			}
			writerChannel <- *val
			time.Sleep(time.Second * 3)
		}
	}()

	for i := 0; i < 5; i++ {
		pub_val := <-writerChannel
		PublishSignals(pub_val.Name, pub_val.Value, settings.Conf.NameSpaces[0], publisher)
		log.Info("publishing ", pub_val.Name)
	}

}
