/******** Peter Winzell (c), 8/31/23 *********************************************/

package viss_rl_interfaces

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

//VSS Signal Cache. uses a key value map to store values.
type SignalAPI interface {
	// returns true if the signal is stored in the cache
	GetSignalValue(id string) (bool, any)
	SetSignalValue(id string, val any)
	Start(quitSignal chan struct{}) error
}

var sigapiInstance *InternalConnection
var singletonlock = &sync.Mutex{}

/**
Return a singelton instance of the signal api
*/
func getSignalApiInstance() *InternalConnection {

	if sigapiInstance == nil {
		singletonlock.Lock()
		defer singletonlock.Unlock()
		if sigapiInstance == nil {
			log.Debug("Creating single instance now.")
			sigapiInstance = &InternalConnection{
				signalcache: make(map[string]any, 100),
			}
		} else {
			log.Debug("Single instance already created.")
		}
	} else {
		log.Debug("Single instance already created.")
	}

	return sigapiInstance
}

func GetSignalApi() (api SignalAPI) {
	return getSignalApiInstance()
}

//VSS broker relay API. channel reader writer api. Does not cache an\y values.
type SignalListenerAPI interface {
	WriterReader(quitSignal chan struct{}, writer chan ValueChannel, reader chan ValueChannel) error
}

var sigapiListenerInstance *InternalConnection_WR
var singletonListenerlock = &sync.Mutex{}

/**
Return a singelton instance of the signal api
*/
func getWriReadSignalApiInstance() *InternalConnection_WR {

	if sigapiListenerInstance == nil {
		singletonListenerlock.Lock()
		defer singletonListenerlock.Unlock()
		if sigapiListenerInstance == nil {
			log.Debug("Creating single instance now.")
			sigapiListenerInstance = &InternalConnection_WR{}
		} else {
			log.Debug("Single instance already created.")
		}
	} else {
		log.Debug("Single instance already created.")
	}

	return sigapiListenerInstance
}

func GetWriterReaderlApi() (api SignalListenerAPI) {
	return getWriReadSignalApiInstance()
}
