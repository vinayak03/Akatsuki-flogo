package hybrisProduct

import (
	"bytes"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"net/http"
)

var log = logger.GetLogger("activity-hybris")

const (
	oValueStatus          = "status"
	oValueResponsePayload = "responsePayload"
	eInternalError        = "500"
	iValueRequestType     = "requestType"
	iValueApiKey          = "apiKey"
	iValueUrl             = "url"
	iValueTenant          = "tenant"
	iValueProductId       = "productId"
)

// HybrisActivity is a stub for your Activity implementation
type HybrisActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &HybrisActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *HybrisActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *HybrisActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the activity data from the context
	//RequestType := context.GetInput(iValueRequestType).(string)
	ApiKey := context.GetInput(iValueApiKey).(string)
	//Url := context.GetInput(iValueUrl).(string)
	//Tenant := context.GetInput(iValueTenant).(string)
	//ProductId := context.GetInput(iValueProductId).(string)

	resp, err := http.Get("http://vkadam-t420.apac.tibco.com:7778/1/products")
	if err != nil {
		fmt.Println("Got Error in the call ", err)
	}
	buf := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error in reading data")
	}
	body := buf.Bytes()
	fmt.Println("Response Body: ", string(body))
	fmt.Println("Response Code: ", resp.StatusCode)
	fmt.Println("ApiKey: ", ApiKey)

	// Set the result as part of the context
	context.SetOutput(oValueStatus, string(resp.StatusCode))
	context.SetOutput(oValueResponsePayload, string(body))

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}
