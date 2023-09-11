/******** Peter Winzell (c), 9/4/23 *********************************************/

package viss_rl_interfaces

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"testing"
	"time"
)

func TestSignalApiStreaming(t *testing.T) {
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

// Test function that listens to the stream...
func listenToDataChannel(dataChannel chan ValueChannel) {
	for {
		dataValue := <-dataChannel
		switch dataValue.Value.(type) {
		case float64:
			str := strconv.FormatFloat(dataValue.Value.(float64), 'f', -1, 64)
			fmt.Println(dataValue.Name + " = " + str)
		case int64:
			str := strconv.FormatInt(dataValue.Value.(int64), 10)
			fmt.Println(dataValue.Name + " = " + str)
		case bool:
			str := strconv.FormatBool(dataValue.Value.(bool))
			fmt.Println(dataValue.Name + " = " + str)
		case []byte:
			fmt.Println(dataValue.Name + " = " + string(dataValue.Value.([]byte)))
		}
	}
}

func TestWriterReader(t *testing.T) {

	wr_api := GetWriterReaderlApi()

	readQuitSignal := make(chan struct{}, 1)
	streamQuitSignal := make(chan struct{}, 1)
	writerChannel := make(chan ValueChannel, 1)
	readerChannel := make(chan ValueChannel, 1)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM) // listen to OS interrupt, like ctrl-c
	go func() {
		<-sig
		close(streamQuitSignal)
		close(readQuitSignal)
	}()

	ch := make(chan int, 2)

	go func() {
		err := wr_api.WriterReader(streamQuitSignal, writerChannel, readerChannel)
		if err != nil {
			log.Println(err)
			ch <- 1
		}
		close(readQuitSignal)
		log.Println("subscribing is done")
		ch <- 0
	}()

	// writing values to the broker every 3 seconds...
	go func() {
		for {
			val := &ValueChannel{
				Name:  "Vehicle.Some.Signal",
				Value: int64(100),
			}
			writerChannel <- *val
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			listenToDataChannel(readerChannel)
		}
	}()

	os.Exit(<-ch)
}
