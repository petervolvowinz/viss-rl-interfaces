/******** Peter Winzell (c), 11/9/23 *********************************************/

package broker

import (
	"github.com/petervolvowinz/viss-rl-interfaces/base"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func FindCertificates() string {
	var certFiles = []string{
		"/etc/ssl/certs/ca-certificates.crt",                // Debian/Ubuntu/Gentoo etc.
		"/etc/pki/tls/certs/ca-bundle.crt",                  // Fedora/RHEL 6
		"/etc/ssl/ca-bundle.pem",                            // OpenSUSE
		"/etc/pki/tls/cacert.pem",                           // OpenELEC
		"/etc/pki/ca-trust/extracted/pem/tls-ca-bundle.pem", // CentOS/RHEL 7
		"/etc/ssl/cert.pem",                                 // Alpine Linux
	}

	for _, file := range certFiles {
		if FileExists(file) {
			return file
		}
	}
	return GetGRPCBrokerSettingInstance().Conf.CertPathName
}

func GetSignalId(signalName string, namespaceName string) *base.SignalId {
	return &base.SignalId{
		Name: signalName,
		Namespace: &base.NameSpace{
			Name: namespaceName},
	}
}

func GetSubscriberConfig(cid string, signalids []*base.SignalId) *base.SubscriberConfig {
	return &base.SubscriberConfig{
		ClientId: &base.ClientId{
			Id: cid,
		},
		Signals: &base.SignalIds{
			SignalId: signalids,
		},
		OnChange: true,
	}
}

func GetPublisherConfig(cid string, sigs []*base.Signal, frequency int32) *base.PublisherConfig {
	return &base.PublisherConfig{
		ClientId: &base.ClientId{
			Id: cid,
		},
		Signals: &base.Signals{
			Signal: sigs,
		},
		Frequency: frequency,
	}
}
