package azureiot

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activit-azureiot")

const (
	ivdeviceID        = "deviceID"
	ivhostName        = "hostName"
	ivsharedAccessKey = "sharedAccessKey"

	ovresult = "output"
)

// AzureIot is a stub for your Activity implementation
type AzureIot struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &AzureIot{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *AzureIot) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *AzureIot) Eval(context activity.Context) (done bool, err error) {

	// do eval
	sharedAccessKey := context.GetInput(ivsharedAccessKey).(string)
	hostName := context.GetInput(ivhostName).(string)
	deviceID := context.GetInput(ivdeviceID).(string)

	log.Debug("The hostname is [%s]", hostName)
	log.Debug("The device is [%s]", deviceID)
	log.Debug("The shared access key is [%s]", sharedAccessKey)

	context.SetOutput(ovresult, "Trying to connect to device "+deviceID+" using hostname: "+hostName+" and sharedAccesskey as "+sharedAccessKey)
	return true, nil
}
