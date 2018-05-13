package hybrisProduct

import (
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
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

func TestGetProduct(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, "GetProduct", "bvCSvnE0BAYcOLKFAD9GxnSToKIwUUWJ", "http://localhost:7778", "dummy", "product1", "")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)

	assert.Equal(t, "200", status, "Response Code does not match")
	t.Log("TestGetProduct Output\n Status:", status, "\nResponsePayload", payload)
}

func TestGetAllProducts(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, "GetAllProducts", "bvCSvnE0BAYcOLKFAD9GxnSToKIwUUWJ", "http://localhost:7778", "dummy", "", "")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)

	assert.Equal(t, "200", status, "Response Code does not match")
	t.Log("TestGetAllProducts Output\n Status:", status, "\nResponsePayload", payload)
}

func TestDeleteProduct(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, "DeleteProduct", "bvCSvnE0BAYcOLKFAD9GxnSToKIwUUWJ", "http://localhost:7778", "dummy", "product1", "")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)

	assert.Equal(t, "200", status, "Response Code does not match")
	t.Log("TestDeleteProduct Output\n Status:", status, "\nResponsePayload", payload)
}

func TestDeleteAllProducts(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, "DeleteAllProducts", "bvCSvnE0BAYcOLKFAD9GxnSToKIwUUWJ", "http://localhost:7778", "dummy", "", "")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)

	assert.Equal(t, "200", status, "Response Code does not match")
	t.Log("TestDeleteAllProducts Output\n Status:", status, "\nResponsePayload", payload)
}

func TestCreateProduct(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, "CreateProduct", "bvCSvnE0BAYcOLKFAD9GxnSToKIwUUWJ", "http://localhost:7778", "dummy", "1", "{\"metadata\":{\"additionalProperties\":true,\"type\":\"string\",\"properties\":{\"mixins\":{\"additionalProperties\":{\"format\":\"string\",\"type\":\"string\"},\"type\":\"string\"},\"variants\":{\"additionalProperties\":true,\"type\":\"string\",\"properties\":{\"options\":{\"additionalProperties\":{\"format\":\"string\",\"type\":\"string\"},\"type\":\"string\"}}}}},\"id\":\"string\",\"code\":\"string\",\"name\":\"string\",\"description\":\"string\",\"published\":true,\"mixins\":\"string\"}")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)

	assert.Equal(t, "200", status, "Response Code does not match")
	t.Log("TestCreateProduct Output\n Status:", status, "\nResponsePayload", payload)
}

func TestUpdateProduct(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, "UpdateProduct", "bvCSvnE0BAYcOLKFAD9GxnSToKIwUUWJ", "http://localhost:7778", "dummy", "1", "{\"metadata\": {\"mixins\": \"string\",\"variants\": {\"option\": \"string\"}},\"code\": \"string\",\"name\": \"string\",\"description\": \"string\",\"published\": true,\"mixins\": {\"type\": \"string\"}}")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)

	assert.Equal(t, "200", status, "Response Code does not match")
	t.Log("TestUpdateProduct Output\n Status:", status, "\nResponsePayload", payload)
}

func TestGetProductFail(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, "GetProduct", "bvCSvnE0BAYcOLKFAD9GxnSToKIwUUWJ", "http://localhost:7779", "dummy", "product1", "")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)

	assert.Equal(t, "500", status, "Response Code does not match")
	t.Log("TestGetProduct Output\n Status:", status, "\nResponsePayload", payload)
}
func TestWrongOption(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, "1", "bvCSvnE0BAYcOLKFAD9GxnSToKIwUUWJ", "http://localhost:7779", "dummy", "product1", "")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus).(string)
	payload := tc.GetOutput(oValueResponsePayload)

	assert.Equal(t, "500", status, "Response Code does not match")
	t.Log("TestGetProduct Output\n Status:", status, "\nResponsePayload", payload)
}

func setInput(tc *test.TestActivityContext, requestType string, APIKey string, url string, tenant string, productId string, body string) {
	tc.SetInput(iValueRequestType, requestType)
	tc.SetInput(iValueAPIKey, APIKey)
	tc.SetInput(iValueUrl, url)
	tc.SetInput(iValueTenant, tenant)
	tc.SetInput(iValueProductId, productId)
	tc.SetInput(iValueBody, body)
}
