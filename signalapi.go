/******** Peter Winzell (c), 8/31/23 *********************************************/

package viss_rl_interfaces

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

//VSS Signal Cache
type SignalAPI interface {
	// returns true if the signal is stored in the cache
	GetSignalValue(id string) (bool, any)
	SetSignalValue(id string, val any)
	Start(quitSignal chan struct{}) error
}

var apiInstance *InternalConnection
var singletonlock = &sync.Mutex{}

/**
Return a singelton instance of the signal api
*/
func getSignalApiInstance() *InternalConnection {

	if apiInstance == nil {
		singletonlock.Lock()
		defer singletonlock.Unlock()
		if apiInstance == nil {
			log.Debug("Creating single instance now.")
			apiInstance = &InternalConnection{
				signalcache: make(map[string]any, 100),
			}
		} else {
			log.Debug("Single instance already created.")
		}
	} else {
		log.Debug("Single instance already created.")
	}

	return apiInstance
}

func GetSignalApi() (api SignalAPI) {
	return getSignalApiInstance()
}
