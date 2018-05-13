package hyperledgerFabricChainCode

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

func TestQuery(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, CONFIG, REQUEST_TYPE_QUERY, "mychannel", "user1", "user1pw", "Org1", "fabcar", "queryCar", "{\"params\" : [\"CAR1\"]}")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)
	errorMessage := tc.GetOutput(oValueErrorMessage)

	t.Log("TestQuery Output\n Status:", status, "\nResponsePayload", payload, "\nErrorMessage", errorMessage)

	assert.Equal(t, "200", status)
}

func TestQueryInvalidEntry(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, CONFIG, REQUEST_TYPE_QUERY, "mychannel", "user1", "user1pw", "Org1", "fabcar", "queryCar", "{\"params\" : [\"CAR14\"]}")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)
	errorMessage := tc.GetOutput(oValueErrorMessage)

	t.Log("TestQuery Output\n Status:", status, "\nResponsePayload", payload, "\nErrorMessage", errorMessage)

	assert.Equal(t, "500", status)
}

func TestExecute(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
	setInput(tc, CONFIG, REQUEST_TYPE_EXECUTE, "mychannel", "user1", "user1pw", "Org1", "fabcar", "createCar", "{\"params\" : [\"CAR12\", \"HONDA\",\"City\",\"Red\",\"Seun\"]}")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)
	errorMessage := tc.GetOutput(oValueErrorMessage)
	transactionID := tc.GetOutput(oValueTransactionID)

	t.Log("TestExecute Output\n Status:", status, "\nResponsePayload", payload, "\nErrorMessage", errorMessage, "\n TransactionID ", transactionID)
	assert.Equal(t, "200", status)
	assert.NotNilf(t, transactionID, "Transaction ID is null after executing the transation on chaincode")
}

func TestInvalidUseExecuter(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
	setInput(tc, CONFIG, REQUEST_TYPE_EXECUTE, "mychannel", "user2", "user1pw", "Org1", "fabcar", "createCar", "{\"params\" : [\"CAR12\", \"HONDA\",\"City\",\"Red\",\"Seun\"]}")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)
	errorMessage := tc.GetOutput(oValueErrorMessage)

	t.Log("TestInvalidUseExecuter Output\n Status:", status, "\nResponsePayload", payload, "\nErrorMessage", errorMessage)
	assert.Equal(t, "500", status)
}

func TestInvalidUserQuery(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
	setInput(tc, CONFIG, REQUEST_TYPE_QUERY, "mychannel", "user2", "user1pw", "Org1", "fabcar", "queryCar", "{\"params\" : [\"CAR1\"]}")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)
	errorMessage := tc.GetOutput(oValueErrorMessage)

	t.Log("TestInvalidUserQuery Output\n Status:", status, "\nResponsePayload", payload, "\nErrorMessage", errorMessage)
	assert.Equal(t, "500", status)
}

func TestInvalidChannel(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, CONFIG, REQUEST_TYPE_QUERY, "mychannel2", "user1", "user1pw", "Org1", "", "queryCar", "{\"params\" : [\"CAR1\"]}")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)
	errorMessage := tc.GetOutput(oValueErrorMessage)
	t.Log("TestInvalidChannel Output\n Status:", status, "\nResponsePayload", payload, "\nErrorMessage", errorMessage)

	assert.Equal(t, "500", status)
}
func TestInvalidInputForArray(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, CONFIG, REQUEST_TYPE_QUERY, "mychannel2", "user1", "user1pw", "Org1", "", "queryCar", "{\"params\" : [\"CAR1\"}")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)
	errorMessage := tc.GetOutput(oValueErrorMessage)
	t.Log("TestInvalidChannel Output\n Status:", status, "\nResponsePayload", payload, "\nErrorMessage", errorMessage)

	assert.Equal(t, "500", status)
}

func setInput(tc *test.TestActivityContext, netconfig string, requestType string, channelid string, user string, userpasswd string, userorg string, chaincodeid string, fcnname string, params string) {
	tc.SetInput(iValueNetworkConfig, netconfig)
	tc.SetInput(iValueRequestType, requestType)
	tc.SetInput(iValueChannelID, channelid)
	tc.SetInput(iValueUser, user)
	tc.SetInput(iValueUserPasswd, userpasswd)
	tc.SetInput(iValueUserOrg, userorg)
	tc.SetInput(iValueChainCodeID, chaincodeid)
	tc.SetInput(iValueFunctionName, fcnname)
	tc.SetInput(iValueParams, params)
}