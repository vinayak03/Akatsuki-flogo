package hybrisProduct

import (
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
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

	setInput(tc, "Get", "bvCSvnE0BAYcOLKFAD9GxnSToKIwUUWJ", "https://api.beta.yaas.io/hybris/product/v2/{tenant}/products", "dummy", "product1")

	act.Eval(tc)

	status := tc.GetOutput(oValueStatus)
	payload := tc.GetOutput(oValueResponsePayload)

	t.Log("TestQuery Output\n Status:", status, "\nResponsePayload", payload)
}

func setInput(tc *test.TestActivityContext, requestType string, apiKey string, url string, tenant string, productId string) {
	tc.SetInput(iValueRequestType, requestType)
	tc.SetInput(iValueApiKey, apiKey)
	tc.SetInput(iValueUrl, url)
	tc.SetInput(iValueTenant, tenant)
	tc.SetInput(iValueProductId, productId)
}
