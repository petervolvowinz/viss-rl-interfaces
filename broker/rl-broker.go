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
	Tls                       string   `json:"tls"`
	CertPathName              string   `json:"cert_path_name"`
	NameSpaces                []string `json:"name_spaces"`
	BrokerUrl                 string   `json:"broker_url"`
	Port                      string   `json:"port"`
	ClientId                  string   `json:"client_id"`
	ApiKey                    string   `json:"api_key"`
	VssTreePath               string   `json:"vss_tree_path"`
	Signalfilter              []string `json:"signalfilter"`
	PublishSeparateConnection bool     `json:"publish-separate-connection"`
	PublishUrl                string   `json:"publish_url"`
	PublishApiKey             string   `json:"publish_api-key"`
}

type GRPCBrokerSettings struct {
	Creds     credentials.TransportCredentials
	Creds_pub credentials.TransportCredentials
	Md        metadata.MD
	Uri       string
	Conf      Configuration
	Uri_pub   string
	Md_pub    metadata.MD
}

type GrpcSetting interface {
	SetCredsApiMetadata()
}

var singletonlock = &sync.Mutex{}
var singleInstance *GRPCBrokerSettings

func GetGRPCBrokerSettingInstance() *GRPCBrokerSettings {
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

/*type frame struct {
	Frameid string   `json:frameid`
	Sigids  []string `json:sigids`
}

type spaces struct {
	Name   string  `json:name`
	Frames []frame `json:framee`
}

type settings struct {
	Namespaces []spaces `json:namespaces`
}*/

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	certfilename := FindCertificates()
	log.Info("gRPC tls file cert ", certfilename)

	pemServerCA, err := ioutil.ReadFile(certfilename)
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
func getSubSignals(conf *Configuration) *base.SubscriberConfig {
	var signalids []*base.SignalId
	namespacename := conf.NameSpaces[0]

	if conf.Signalfilter[0] == "*" {
		panic("vss tree config not implemented yet")
	}
	for cindex := 0; cindex < len(conf.Signalfilter); cindex++ {
		signalId := GetSignalId(conf.Signalfilter[cindex], namespacename)
		signalids = append(signalids, signalId)
	}
	// add selected signals to subscriber configuration
	return GetSubscriberConfig("pw-golang-client", signalids)
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
	var creds credentials.TransportCredentials = nil
	if conf.Tls == "yes" {
		creds, _ = loadTLSCredentials()
		if creds == nil {
			log.Info("could not create tls credentials, trying with insecure")
		}
	}
	if creds == nil || conf.Tls == "no" {
		creds = credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})
	}

	uri := net.JoinHostPort(conf.BrokerUrl, conf.Port)
	md := metadata.Pairs(
		"x-api-key", conf.ApiKey,
	)

	if conf.PublishSeparateConnection {
		uri_p := net.JoinHostPort(conf.PublishUrl, conf.Port)
		md_p := metadata.Pairs(
			"x-api-key", conf.PublishApiKey)
		settings.Uri_pub = uri_p
		settings.Md_pub = md_p
	}

	settings.Conf = conf
	settings.Creds = creds
	settings.Uri = uri
	settings.Md = md

	return settings
}

// prints current server config to the console
func SeverConfig() error {
	settings := GetGRPCBrokerSettingInstance().SetCredsApiMetadata()
	conn, err := grpc.Dial(settings.Uri, grpc.WithTransportCredentials(settings.Creds))
	if err != nil {
		log.Debug("did not connect ", err)
		return err
	}

	err = PrintSignalTree(conn, settings.Md)
	if err != nil {
		log.Debug("could not print config ", err)
		return err
	}
	conn.Close()
	return nil
}

func GetBrokerConnections() (base.NetworkServiceClient, base.NetworkServiceClient, *GRPCBrokerSettings, error) {
	c1, settings, err := getBrokerConnection()
	if err != nil {
		return nil, nil, nil, err
	}
	c2, err2 := checkSeparatePublishConnection(settings)
	if err2 != nil {
		return nil, nil, nil, err2
	}

	if c2 == nil {
		return c1, c1, settings, nil
	}

	return c1, c2, settings, nil
}

func getBrokerConnection() (base.NetworkServiceClient, *GRPCBrokerSettings, error) {
	grpc_settings := GetGRPCBrokerSettingInstance().SetCredsApiMetadata()
	conn, err := grpc.Dial(grpc_settings.Uri, grpc.WithTransportCredentials(grpc_settings.Creds))
	c := base.NewNetworkServiceClient(conn)
	if err != nil {
		log.Debug("did not connect to broker", err)
		return nil, nil, err
	}

	return c, grpc_settings, nil
}

func checkSeparatePublishConnection(grpc_settings *GRPCBrokerSettings) (base.NetworkServiceClient, error) {
	if grpc_settings.Conf.PublishSeparateConnection {
		conn, err := grpc.Dial(grpc_settings.Uri_pub, grpc.WithTransportCredentials(grpc_settings.Creds))
		c := base.NewNetworkServiceClient(conn)
		if err != nil {
			log.Debug("did not connect to broker", err)
			return nil, err
		}
		return c, nil
	}
	return nil, nil
}

func StartStreaming(connection base.NetworkServiceClient, grpc_settings *GRPCBrokerSettings) (base.NetworkService_SubscribeToSignalsClient, error) {
	// subscription
	subConfig := getSubSignals(&grpc_settings.Conf)
	ctx := metadata.NewOutgoingContext(context.Background(), grpc_settings.Md)
	response, err := connection.SubscribeToSignals(ctx, subConfig)
	if err != nil {
		log.Debug("did not connect ", err)
		return nil, err
	}
	return response, nil
}

func PublishSignals(signame string, sigvalue any, namespace string, serviceClient base.NetworkServiceClient) {

	log.Println("publishing: ", signame)

	signalId := GetSignalId(signame, namespace)
	pubsignal := &base.Signal{
		Id: signalId,
	}

	switch sigvalue.(type) {
	case int64:
		pubsignal.Payload = &base.Signal_Integer{
			Integer: sigvalue.(int64),
		}
		break
	case float64:
		pubsignal.Payload = &base.Signal_Double{
			Double: sigvalue.(float64),
		}
		break
	case bool:
		pubsignal.Payload = &base.Signal_Arbitration{
			Arbitration: sigvalue.(bool),
		}
		break
	}
	// pubsignal.Raw = sigvalue.([]byte) TODO discuss how to handle frame publishing

	var signals []*base.Signal
	signals = append(signals, pubsignal)

	pubconfig := GetPublisherConfig("pw-golang-client", signals, 0)
	ctx := context.Background()
	serviceClient.PublishSignals(ctx, pubconfig)

}
