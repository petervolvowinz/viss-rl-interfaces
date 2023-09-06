/******** Peter Winzell (c), 8/31/23 *********************************************/

package broker

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	base "github.com/petervolvowinz/viss-rl-interfaces/base"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"net"
	"os"
	"sync"
)

type Configuration struct {
	CertPathName string   `json:"cert_path_name"`
	NameSpaces   []string `json:"name_spaces"`
	BrokerUrl    string   `json:"broker_url"`
	Port         string   `json:"port"`
	ClientId     string   `json:"client_id"`
	ApiKey       string   `json:"api_key"`
	VssTreePath  string   `json:"vss_tree_path"`
	Signalfilter []string `json:"signalfilter"`
}

type GRPCBrokerSettings struct {
	Creds credentials.TransportCredentials
	Md    metadata.MD
	Uri   string
	conf  Configuration
}

type GrpcSetting interface {
	SetCredsApiMetadata()
}

var singletonlock = &sync.Mutex{}

var singleInstance *GRPCBrokerSettings

func getGRPCBrokerSettingInstance() *GRPCBrokerSettings {
	if singleInstance == nil {
		singletonlock.Lock()
		defer singletonlock.Unlock()
		if singleInstance == nil {
			log.Debug("Creating single instance now.")
			singleInstance = &GRPCBrokerSettings{}
		} else {
			log.Debug("Single instance already created.")
		}
	} else {
		log.Debug("Single instance already created.")
	}

	return singleInstance
}

type frame struct {
	Frameid string   `json:frameid`
	Sigids  []string `json:sigids`
}

type spaces struct {
	Name   string  `json:name`
	Frames []frame `json:framee`
}

type settings struct {
	Namespaces []spaces `json:namespaces`
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("certificate.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

// set signal name and namespace to grpc generated data strcuture
func getSignalId(signalName string, namespaceName string) *base.SignalId {
	return &base.SignalId{
		Name: signalName,
		Namespace: &base.NameSpace{
			Name: namespaceName},
	}
}

func getSignals(conf *Configuration) *base.SubscriberConfig {
	var signalids []*base.SignalId
	namespacename := conf.NameSpaces[0]

	if conf.Signalfilter[0] == "*" {
		panic("vss tree config not implemented yet")
	}
	for cindex := 0; cindex < len(conf.Signalfilter); cindex++ {
		signalId := getSignalId(conf.Signalfilter[cindex], namespacename)
		signalids = append(signalids, signalId)
	}

	// add selected signals to subscriber configuration
	signals := &base.SubscriberConfig{
		ClientId: &base.ClientId{
			Id: "pw-golang-client",
		},
		Signals: &base.SignalIds{
			SignalId: signalids,
		},
		OnChange: true,
	}

	return signals
}

func readConfiguration() Configuration {
	file, err := os.Open("../config.json")
	if err != nil {
		log.Debug(err.Error())
		file.Close()
		file, err = os.Open("config.json")
		if err != nil {
			panic(err.Error())
		}
	}
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)

	var conf Configuration
	json.Unmarshal(bytes, &conf)

	return conf
}
func (settings *GRPCBrokerSettings) SetCredsApiMetadata() *GRPCBrokerSettings {

	conf := readConfiguration()
	creds, _ := loadTLSCredentials()
	uri := net.JoinHostPort(conf.BrokerUrl, conf.Port)
	md := metadata.Pairs(
		"x-api-key", conf.ApiKey,
	)

	settings.conf = conf
	settings.Creds = creds
	settings.Uri = uri
	settings.Md = md

	return settings
}

// prints current server config to the console
func SeverConfig() error {
	settings := getGRPCBrokerSettingInstance().SetCredsApiMetadata()
	conn, err := grpc.Dial(settings.Uri, grpc.WithTransportCredentials(settings.Creds))
	if err != nil {
		log.Debug("did not connect ", err)
		return err
	}
	defer conn.Close()

	return PrintSignalTree(conn)
}

func GetBrokerConnection() (base.NetworkServiceClient, *GRPCBrokerSettings, error) {
	grpc_settings := getGRPCBrokerSettingInstance().SetCredsApiMetadata()
	conn, err := grpc.Dial(grpc_settings.Uri, grpc.WithTransportCredentials(grpc_settings.Creds))
	c := base.NewNetworkServiceClient(conn)
	if err != nil {
		log.Debug("did not connect to broker", err)
		return nil, nil, err
	}

	return c, grpc_settings, nil
}

func StartStreaming() (base.NetworkService_SubscribeToSignalsClient, error) {
	connection, grpc_settings, err := GetBrokerConnection()
	if err == nil {
		// subscription
		subConfig := getSignals(&grpc_settings.conf)
		ctx := metadata.NewOutgoingContext(context.Background(), grpc_settings.Md)
		response, err := connection.SubscribeToSignals(ctx, subConfig)

		if err != nil {
			log.Debug("did not connect ", err)
			return nil, err
		}
		return response, nil
	}
	return nil, err
}
