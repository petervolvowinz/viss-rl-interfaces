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

	subscriber, _, settings, error := broker.GetBrokerConnections()
	if error != nil {
		return error
	}

	resp, err := broker.StartStreaming(subscriber, settings)
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

/*
void GrpcConnection::publisher()
{

  // TODO we could use startvalue here as default
  // https://github.com/remotivelabs/remotivelabs-apis/blob/main/proto/common.proto#L34
  auto start_value = 12;

  auto N = 10;
  for (auto i = 0; i < N; i = ((i + 1) % N))
  {
    // std::cout << i << std::endl;
    auto signals = new Signals();
    {
      auto signal_id = new SignalId();
      signal_id->set_allocated_name(new std::string("SteeringAngle129"));
      signal_id->set_allocated_namespace_(new NameSpace(*name_space));
      auto handle = signals->add_signal();
      handle->set_allocated_id(signal_id);
      handle->set_integer(start_value + i);
    }
    {
      // append any number of signals here! (duplicate above code)
    }

    PublisherConfig pub_info;
    pub_info.set_allocated_clientid(new ClientId(*source));
    pub_info.set_allocated_signals(signals);
    pub_info.set_frequency(0);
    ClientContext ctx;
    Empty empty;
    stub->PublishSignals(&ctx, pub_info, &empty);

    // TODO we should derive this period from proto buffer,
    // https://github.com/remotivelabs/remotivelabs-apis/blob/main/proto/common.proto#L33
    usleep(30);
  }
}

*/
func writeToBroker(writer chan ValueChannel, serviceClient base.NetworkServiceClient, settings *broker.GRPCBrokerSettings) {

	for {
		dataValue := <-writer
		log.Debug("publishing signal ", dataValue.Name)
		broker.PublishSignals(dataValue.Name, dataValue.Value, settings.Conf.NameSpaces[0], serviceClient)
	}
}

func (IC *InternalConnection_WR) WriterReader(quitSignal chan struct{}, writer chan ValueChannel, reader chan ValueChannel) error {

	subscriber, publisher, settings, error := broker.GetBrokerConnections()
	if error != nil {
		return error
	}

	resp, err := broker.StartStreaming(subscriber, settings)
	if err != nil {
		log.Debug(err)
		return err
	}

	//publisher...
	go writeToBroker(writer, publisher, settings)

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
