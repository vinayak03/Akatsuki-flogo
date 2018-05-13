package hybrisProduct

import (
	"bytes"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"net/http"
)

var log = logger.GetLogger("activity-hybris-product")

const (
	oValueStatus           = "statusCode"
	oValueResponsePayload  = "responsePayload"
	eInternalError         = "500"
	iValueRequestType      = "requestType"
	iValueAPIKey           = "APIKey"
	iValueUrl              = "URL"
	iValueTenant           = "tenant"
	iValueProductId        = "productId"
	iValueBody             = "body"
	ConstGetAllProducts    = "GetAllProducts"
	ConstGetProduct        = "GetProduct"
	ConstCreateProduct     = "CreateProduct"
	ConstUpdateProduct     = "UpdateProduct"
	ConstDeleteProduct     = "DeleteProduct"
	ConstDeleteAllProducts = "DeleteAllProducts"
)

// HybrisProductActivity is a stub for your Activity implementation
type HybrisProductActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &HybrisProductActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *HybrisProductActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *HybrisProductActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the activity data from the context
	RequestType := context.GetInput(iValueRequestType).(string)
	APIKey := context.GetInput(iValueAPIKey).(string)
	BaseURL := context.GetInput(iValueUrl).(string)
	Tenant := context.GetInput(iValueTenant).(string)
	ProductId := context.GetInput(iValueProductId).(string)
	Body := context.GetInput(iValueBody).(string)

	var URL string
	var method string

	switch RequestType {
	case ConstGetAllProducts:
		URL = fmt.Sprintf("%s%s%s%s%s", BaseURL, "/", Tenant, "/", "products")
		method = "GET"
	case ConstCreateProduct:
		URL = fmt.Sprintf("%s%s%s%s%s", BaseURL, "/", Tenant, "/", "products")
		method = "POST"
	case ConstDeleteAllProducts:
		URL = fmt.Sprintf("%s%s%s%s%s", BaseURL, "/", Tenant, "/", "products")
		method = "DELETE"
	case ConstGetProduct:
		URL = fmt.Sprintf("%s%s%s%s%s%s%s", BaseURL, "/", Tenant, "/", "products", "/", ProductId)
		method = "GET"
	case ConstUpdateProduct:
		URL = fmt.Sprintf("%s%s%s%s%s%s%s", BaseURL, "/", Tenant, "/", "products", "/", ProductId)
		method = "PUT"
	case ConstDeleteProduct:
		URL = fmt.Sprintf("%s%s%s%s%s%s%s", BaseURL, "/", Tenant, "/", "products", "/", ProductId)
		method = "DELETE"
	}

	log.Info("Computed URL :", URL)
	log.Info("Computed Method :", method)

	var requestBuffer *bytes.Buffer = bytes.NewBuffer([]byte(""))
	if Body != "" {
		requestBuffer = bytes.NewBuffer([]byte(Body))
	}

	req, err := http.NewRequest(method, URL, requestBuffer)
	if err != nil {
		log.Errorf("Error Creatign HTTP Request %s", err.Error())
		SetOutput(context, "500")
		return true, fmt.Errorf("Error Creatign HTTP Request %s", err.Error())
	}

	client := &http.Client{}
	req.Header.Add("Authorization", fmt.Sprint("bearer ", APIKey))

	resp, err := client.Do(req)

	if err != nil {
		log.Errorf("Error HTTP Processing %s", err.Error())
		SetOutput(context, "500")
		return true, fmt.Errorf("Error HTTP Processing %s", err.Error())
	}

	if resp == nil {
		log.Error("Error HTTP Processing Empty Response from Server")
		SetOutput(context, "500")
		return false, fmt.Errorf("Error HTTP Processing Empty Response from Server")
	}

	buf := bytes.NewBuffer(make([]byte, 0, resp.ContentLength))
	_, err = buf.ReadFrom(resp.Body)

	if err != nil {
		log.Error("Error Reading Response")
		SetOutput(context, "500")
		return true, fmt.Errorf("Error Reading Response")
	}
	responseBody := buf.Bytes()
	log.Debug("Response Body: ", string(responseBody))
	log.Debug("Response Code: ", fmt.Sprintf("%d", resp.StatusCode))

	// Set the result as part of the context
	context.SetOutput(oValueStatus, fmt.Sprintf("%d", resp.StatusCode))
	context.SetOutput(oValueResponsePayload, string(responseBody))

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}

func SetOutput(context activity.Context, status string) {
	context.SetOutput(oValueStatus, status)
}
