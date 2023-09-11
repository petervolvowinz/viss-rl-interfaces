/******** Peter Winzell, Volvo Cars (c), 8/21/23 *********************************************/

package viss_rl_interfaces

import (
	"errors"
	"github.com/petervolvowinz/viss-rl-interfaces/base"
	"github.com/petervolvowinz/viss-rl-interfaces/broker"
	log "github.com/sirupsen/logrus"
	"io"
	"sync"
)

// If some one wants to listen to the signal stream, go ahead and declare a Value Channel.
type ValueChannel struct {
	Name  string
	Value any
}

type InternalConnection struct {
	instance    *SignalAPI
	signalcache map[string]any
	mx          sync.Mutex
}

type InternalConnection_WR struct {
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
			// data := &ValueChannel{}
			for _, signal := range sigs.Signal {
				//data.Name = signal.Id.Name
				switch signal.Payload.(type) {
				case *base.Signal_Double:
					IC.SetSignalValue(signal.Id.Name, signal.GetDouble())
					//data.Value = signal.GetDouble()
				case *base.Signal_Integer:
					IC.SetSignalValue(signal.Id.Name, signal.GetInteger())
					//data.Value = signal.GetInteger()
				case *base.Signal_Arbitration:
					IC.SetSignalValue(signal.Id.Name, signal.GetArbitration())
					//data.Value = signal.GetArbitration()
				case *base.Signal_Empty:
					IC.SetSignalValue(signal.Id.Name, signal.GetEmpty())
					//data.Value = signal.GetEmpty()
				default:
					IC.SetSignalValue(signal.Id.Name, signal.Raw)
					//data.Value = signal.Raw
					// fmt.Errorf("unsupported new RL signal type %T", signal.Payload)
				}
			}

			// value <- *data // send to channel...
		}
		select {
		case <-quitSignal: // do a nice quit...
			return nil
		default:
		}
	}

	return nil
}

func writeToBroker(writer chan ValueChannel) {
	for {
		dataValue := <-writer
		log.Println("***** TODO writing signals to broker *****")
		log.Println(dataValue.Name)
	}
}

func (IC *InternalConnection_WR) WriterReader(quitSignal chan struct{}, writer chan ValueChannel, reader chan ValueChannel) error {
	resp, err := broker.StartStreaming()
	if err != nil {
		log.Debug(err)
		return err
	}

	go writeToBroker(writer)

	for {
		sigs, err := resp.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Println("Error on recv: " + err.Error())
				return err
			}
		} else {
			data := &ValueChannel{}
			for _, signal := range sigs.Signal {
				data.Name = signal.Id.Name
				switch signal.Payload.(type) {
				case *base.Signal_Double:
					data.Value = signal.GetDouble()
				case *base.Signal_Integer:
					data.Value = signal.GetInteger()
				case *base.Signal_Arbitration:
					data.Value = signal.GetArbitration()
				case *base.Signal_Empty:
					data.Value = signal.GetEmpty()
				default:
					data.Value = signal.Raw
				}
			}

			reader <- *data // send to read channel...
		}
		select {
		case <-quitSignal: // do a nice quit...
			return nil
		default:
		}
	}

	return nil
}
