package consumers

import (
	"fmt"
	"sync"

	"github.com/zalando-incubator/mate/pkg"
)

var params struct {
	domain        string
	project       string
	zone          string
	recordGroupID string
	awsAccountID  string
	awsRole       string
	awsHostedZone string
	awsGroupID    string
}

type Consumer interface {
	Sync([]*pkg.Endpoint) error
	Consume(<-chan *pkg.Endpoint, chan<- error, <-chan struct{}, *sync.WaitGroup)
	Process(*pkg.Endpoint) error
}

// Returns a Consumer implementation.
func New(name string) (Consumer, error) {
	var create func() (Consumer, error)
	switch name {
	case "google":
		create = NewGoogleDNS
	case "aws":
		create = NewAWSConsumer
	case "stdout":
		create = NewStdout
	default:
		return nil, fmt.Errorf("Unknown consumer '%s'.", name)
	}

	c, err := create()
	if err != nil {
		return nil, fmt.Errorf("error creating consumer '%s': %v", name, err)
	}

	return c, nil
}
