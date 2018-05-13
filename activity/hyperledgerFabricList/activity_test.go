package hyperledgerFabricList

import (
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const (
	CONFIG = "D:/Projects/Flogo-Hackthon/go-ws/config.yaml"
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

func TestListChannel(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, CONFIG, "Channel", "Admin", "adminpw", "grpc://localhost:7051", "Org1", "")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)
	errorMessage := tc.GetOutput(oValueErrorMessage)

	t.Log("TestListChannel Output\n Status:", status, "\nResponsePayload", payload, "\nErrorMessage", errorMessage)
	assert.Equal(t, "200", status)
}
func TestListInstalledChainCode(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, CONFIG, "InstalledChainCode", "Admin", "adminpw", "grpc://localhost:7051", "Org1", "")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)
	errorMessage := tc.GetOutput(oValueErrorMessage)

	t.Log("TestListChannel Output\n Status:", status, "\nResponsePayload", payload, "\nErrorMessage", errorMessage)
	assert.Equal(t, "200", status)
}

func TestListInstantiatedChainCode(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, CONFIG, "InstantiatedChainCode", "Admin", "adminpw", "grpc://localhost:7051", "Org1", "mychannel")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)
	errorMessage := tc.GetOutput(oValueErrorMessage)

	t.Log("TestListChannel Output\n Status:", status, "\nResponsePayload", payload, "\nErrorMessage", errorMessage)
	assert.Equal(t, "200", status)
}

func setInput(tc *test.TestActivityContext, netconfig string, resourceType string, adminuser string, adminpw string, queryparam string, org string, channelId string) {
	tc.SetInput(iValueNetworkConfig, netconfig)
	tc.SetInput(iValueResourceType, resourceType)
	tc.SetInput(iValueAdminUser, adminuser)
	tc.SetInput(iValueAdminPasswd, adminpw)
	tc.SetInput(iValueQueryParam, queryparam)
	tc.SetInput(iValueOrg, org)
	tc.SetInput(iValueChannelID, channelId)
}
