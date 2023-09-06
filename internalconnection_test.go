/******** Peter Winzell (c), 9/4/23 *********************************************/

package viss_rl_interfaces

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestStreaming(t *testing.T) {
	api := GetSignalApi()

	streamQuitSignal := make(chan struct{}, 1)
	readQuitSignal := make(chan struct{}, 1)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM) // listen to OS interrupt, like ctrl-c
	go func() {
		<-sig
		close(streamQuitSignal)
		close(readQuitSignal)
	}()

	ch := make(chan int, 2)
	go func() {
		err := api.Start(streamQuitSignal)
		if err != nil {
			log.Println(err)
			ch <- 1
		}
		close(readQuitSignal)
		log.Println("subscribing is done")
		ch <- 0
	}()

	// wait a second...
	time.Sleep(time.Second)

	// Check values and print em to console
	go func() {
		for {
			ok, VIN := api.GetSignalValue("Vehicle.VehicleIdentification.VIN")
			if !ok {
				t.Error("VIN not registrered")
			} else {
				vin_as_str := string(VIN.([]byte))
				log.Println("VIN IS", vin_as_str)
			}

			ok, value := api.GetSignalValue("Vehicle.Speed")
			if !ok {
				log.Println("speed not registered")
				close(streamQuitSignal)
				ch <- 1
			}
			time.Sleep(time.Second)
			switch value.(type) {
			case float64:
				{
					val := value.(float64)
					fmt.Println(val)
				}
			}
			select {
			case <-readQuitSignal: // do a nice quit...
				close(streamQuitSignal)
			default:
			}
		}
		ch <- 0
	}()

	os.Exit(<-ch | <-ch)

}
