/******** Peter Winzell (c), 8/31/23 *********************************************/

package broker

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
	base "viss-rl-interfaces/proto_files"
)

func TestBrokerInfo(t *testing.T) {
	test := SeverConfig()
	if test != nil {
		fmt.Println(test)
		t.Error("test failed")
	}
}

func TestStreaming(t *testing.T) {
	resp, err := StartStreaming()
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
