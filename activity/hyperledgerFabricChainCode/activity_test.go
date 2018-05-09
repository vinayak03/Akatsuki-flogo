package fabricChainCode

import (
	"io/ioutil"
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil{
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
	os.Setenv("fabric-config","D:/Projects/Flogo-Hackthon/go-ws/config.yaml")
	
	tc.SetInput("NetworkConfig", "fabric-config")
	tc.SetInput("RequestType", "Query")
	tc.SetInput("ChannelID", "mychannel")
	tc.SetInput("User", "user1")
	tc.SetInput("UserPasswd", "user1pw")
	tc.SetInput("UserOrg", "Org1")
	tc.SetInput("ChainCodeID", "fabcar")
	tc.SetInput("FunctionName", "queryCar")
	params := "{\"params\" : [\"CAR1\"]}"
	tc.SetInput("Params", params)
	

	act.Eval(tc)

	status := tc.GetOutput("status").(string)
	payload := tc.GetOutput("ResponsePayload")
	errorMessage := tc.GetOutput("ErrorMessage")
	t.Log("TestQuery Output\n Status:",status,"\nResponsePayload",payload,"\nErrorMessage",errorMessage)
	
	assert.Equal(t, status, "200")
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
	os.Setenv("fabric-config","D:/Projects/Flogo-Hackthon/go-ws/config.yaml")
	
	tc.SetInput("NetworkConfig", "fabric-config")
	tc.SetInput("RequestType", "Execute")
	tc.SetInput("ChannelID", "mychannel")
	tc.SetInput("User", "user1")
	tc.SetInput("UserPasswd", "user1pw")
	tc.SetInput("UserOrg", "Org1")
	tc.SetInput("ChainCodeID", "fabcar")
	tc.SetInput("FunctionName", "createCar")
	params := "{\"params\" : [\"CAR12\", \"HONDA\",\"City\",\"Red\",\"Prashant\"]}"
	tc.SetInput("Params", params)
	

	act.Eval(tc)

	status := tc.GetOutput("status").(string)
	payload := tc.GetOutput("ResponsePayload")
	errorMessage := tc.GetOutput("ErrorMessage")
	
	t.Log("TestExecute Output\n Status:",status,"\nResponsePayload",payload,"\nErrorMessage",errorMessage)
	assert.Equal(t, status, "200")
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
	os.Setenv("fabric-config","D:/Projects/Flogo-Hackthon/go-ws/config.yaml")
	
	tc.SetInput("NetworkConfig", "fabric-config")
	tc.SetInput("RequestType", "Execute")
	tc.SetInput("ChannelID", "mychannel")
	tc.SetInput("User", "user2")
	tc.SetInput("UserPasswd", "user1pw")
	tc.SetInput("UserOrg", "Org1")
	tc.SetInput("ChainCodeID", "fabcar")
	tc.SetInput("FunctionName", "createCar")
	params := "{\"params\" : [\"CAR12\", \"HONDA\",\"City\",\"Red\",\"Prashant\"]}"
	tc.SetInput("Params", params)
	

	act.Eval(tc)

	status := tc.GetOutput("status").(string)
	payload := tc.GetOutput("ResponsePayload")
	errorMessage := tc.GetOutput("ErrorMessage")
	
	t.Log("TestInvalidUseExecuter Output\n Status:",status,"\nResponsePayload",payload,"\nErrorMessage",errorMessage)
	assert.Equal(t, status, "500")
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
	os.Setenv("fabric-config","D:/Projects/Flogo-Hackthon/go-ws/config.yaml")
	
	tc.SetInput("NetworkConfig", "fabric-config")
	tc.SetInput("RequestType", "Query")
	tc.SetInput("ChannelID", "mychannel")
	tc.SetInput("User", "user2")
	tc.SetInput("UserPasswd", "user1pw")
	tc.SetInput("UserOrg", "Org1")
	tc.SetInput("ChainCodeID", "fabcar")
	tc.SetInput("FunctionName", "queryCar")
	params := "{\"params\" : [\"CAR1\"]}"
	tc.SetInput("Params", params)
	

	act.Eval(tc)

	status := tc.GetOutput("status").(string)
	payload := tc.GetOutput("ResponsePayload")
	errorMessage := tc.GetOutput("ErrorMessage")
	
	t.Log("TestInvalidUserQuery Output\n Status:",status,"\nResponsePayload",payload,"\nErrorMessage",errorMessage)
	assert.Equal(t, status, "500")
}

func TestInvalidChannel(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
	os.Setenv("fabric-config","D:/Projects/Flogo-Hackthon/go-ws/config.yaml")
	
	tc.SetInput("NetworkConfig", "fabric-config")
	tc.SetInput("RequestType", "Query")
	tc.SetInput("ChannelID", "mychannel2")
	tc.SetInput("User", "user1")
	tc.SetInput("UserPasswd", "user1pw")
	tc.SetInput("UserOrg", "Org1")
	tc.SetInput("ChainCodeID", "fabcar")
	tc.SetInput("FunctionName", "queryCar")
	params := "{\"params\" : [\"CAR1\"]}"
	tc.SetInput("Params", params)
	

	act.Eval(tc)

	status := tc.GetOutput("status").(string)
	payload := tc.GetOutput("ResponsePayload")
	errorMessage := tc.GetOutput("ErrorMessage")
	t.Log("TestInvalidChannel Output\n Status:",status,"\nResponsePayload",payload,"\nErrorMessage",errorMessage)
	
	assert.Equal(t, status, "500")
}