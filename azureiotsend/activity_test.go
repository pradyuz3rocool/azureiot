package azureiot

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	tc.SetInput("sharedAccessKey", "4X9XgcLcu1RjiP7kq7oWVffq+e1jXAcdrOKcC71DM8o=")
	tc.SetInput("hostName", "myhub.azure-devices.net")
	tc.SetInput("deviceID", "raspi-isteer")

	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("output")
	assert.Equal(t, result, "Trying to connect to device raspi-isteer using hostname: myhub.azure-devices.net and sharedAccesskey as 4X9XgcLcu1RjiP7kq7oWVffq+e1jXAcdrOKcC71DM8o=")

	//check result attr
}
