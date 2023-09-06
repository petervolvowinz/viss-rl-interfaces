/******** Peter Winzell, Volvo Cars (c), 8/21/23 *********************************************/

package viss_rl_interfaces

import (
	"errors"
	"fmt"
	"github.com/petervolvowinz/viss-rl-interfaces/base"
	"github.com/petervolvowinz/viss-rl-interfaces/broker"
	log "github.com/sirupsen/logrus"
	"io"
	"sync"
)

type InternalConnection struct {
	instance    *SignalAPI
	signalcache map[string]any
	mx          sync.Mutex
}

type NoValue struct {
}

func (IC *InternalConnection) GetSignalValue(id string) (bool, any) {
	IC.mx.Lock()
	defer IC.mx.Unlock()
	_, ok := IC.signalcache[id]
	if !ok {
		log.Println("no values stored for:" + id)
		return ok, &NoValue{}
	}
	value := IC.signalcache[id]
	return true, value
}

func (IC *InternalConnection) SetSignalValue(id string, value any) {
	IC.mx.Lock()
	defer IC.mx.Unlock()
	if IC.signalcache == nil {
		IC.signalcache = make(map[string]any)
	}
	IC.signalcache[id] = value
}

// Interfaces RL signal broker, assumes data is VSS enabled.
// blocking on quitSignal
func (IC *InternalConnection) Start(quitSignal chan struct{}) error {
	resp, err := broker.StartStreaming()
	if err != nil {
		log.Debug(err)
		return err
	}

	for {
		sigs, err := resp.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Error on recv: " + err.Error())
				return err
			}
		} else {
			for _, signal := range sigs.Signal {

				switch signal.Payload.(type) {
				case *base.Signal_Double:
					IC.SetSignalValue(signal.Id.Name, signal.GetDouble())
				case *base.Signal_Integer:
					IC.SetSignalValue(signal.Id.Name, signal.GetInteger())
				case *base.Signal_Arbitration:
					IC.SetSignalValue(signal.Id.Name, signal.GetArbitration())
				case *base.Signal_Empty:
					IC.SetSignalValue(signal.Id.Name, signal.GetEmpty())
				default:
					IC.SetSignalValue(signal.Id.Name, signal.Raw)
					fmt.Errorf("unsupported new RL signal type %T", signal.Payload)
				}
			}
		}
		select {
		case <-quitSignal: // do a nice quit...
			return nil
		default:
		}
	}

	return nil
}
