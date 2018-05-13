package hyperledgerFabricChainCode

import (
	"encoding/json"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	channel "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	retry "github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	fabric "github.com/vinayak03/Akatsuki-flogo/hyperledgerFabric"
)

var (
	log = logger.GetLogger("activity-fabric-chaincode")
)

const (
	oValueStatus          = "Status"
	oValueErrorMessage    = "ErrorMessage"
	oValueResponsePayload = "ResponsePayload"
	oValueTransactionID   = "TransactionID"
	eInternalError        = "500"
	iValueNetworkConfig   = "NetworkConfig"
	iValueRequestType     = "RequestType"
	iValueChainCodeID     = "ChainCodeID"
	iValueChannelID       = "ChannelID"
	iValueUser            = "User"
	iValueUserPasswd      = "UserPasswd"
	iValueUserOrg         = "UserOrg"
	iValueFunctionName    = "FunctionName"
	iValueParams          = "Params"

	REQUEST_TYPE_QUERY   = "Query"
	REQUEST_TYPE_EXECUTE = "Execute"
)

//structure for parsing the input array of strings
type Params struct {
	Params []string `json:"params"`
}

// ChainCodeActivity is a stub for your Activity implementation
type ChainCodeActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ChainCodeActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *ChainCodeActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *ChainCodeActivity) Eval(context activity.Context) (done bool, err error) {
	var status string
	var payload string
	var transactionID string
	//Initialize the sdk
	networkConfig := context.GetInput(iValueNetworkConfig).(string)
	SDK, err := fabric.GetSDK(networkConfig)

	if err != nil {
		log.Error(err.Error())
		setOutput(context, eInternalError, "", err.Error(), "")
		return true, err
	}

	//Extract the details from activity request.
	requestType := context.GetInput(iValueRequestType).(string)
	chainID := context.GetInput(iValueChainCodeID).(string)
	channelID := context.GetInput(iValueChannelID).(string)
	user := context.GetInput(iValueUser).(string)
	userPasswd := context.GetInput(iValueUserPasswd).(string)
	userOrg := context.GetInput(iValueUserOrg).(string)
	functionName := context.GetInput(iValueFunctionName).(string)
	inputParams := context.GetInput(iValueParams).(string)

	//check if the user is present in local trustsore if not enroll it in local truststore.
	fabric.EnrollWithOrg(SDK, userOrg, user, userPasswd)

	//Convert json array to string array
	params := Params{}
	if err := json.Unmarshal([]byte(inputParams), &params); err != nil {
		log.Error(err.Error())
		setOutput(context, eInternalError, "", err.Error(), "")
		return true, err
	}

	paramSize := len(params.Params)
	log.Debug("Param Size ", paramSize)
	var args [][]byte = make([][]byte, paramSize)

	for index, param := range params.Params {
		args[index] = []byte(param)
		log.Debug("Added param ", param, " at index ", index)
	}

	channelCtx, err := fabric.GetChannelContext(SDK, channelID, user)
	channelClient, err := channel.New(channelCtx)
	if err != nil {
		log.Error(err.Error())
		setOutput(context, eInternalError, "", err.Error(), "")
		return true, err
	}
	if requestType == REQUEST_TYPE_QUERY {
		resp, err := channelClient.Query(channel.Request{
			ChaincodeID: chainID,
			Fcn:         functionName,
			Args:        args},
			channel.WithRetry(retry.DefaultChannelOpts))
		if err != nil {
			log.Error(err.Error())
			setOutput(context, eInternalError, "", err.Error(), "")
			return true, err
		}
		//		if resp==nil {
		//			setOutput(context,"400","","Empty Response from Fabric Network","")
		//			return true,nil
		//		}

		status = strconv.FormatInt(int64(resp.ChaincodeStatus), 10)
		payload = string(resp.Payload)
		log.Info("Response Payload : ", payload)

	} else {
		resp, err := channelClient.Execute(channel.Request{
			ChaincodeID: chainID,
			Fcn:         functionName,
			Args:        args},
			channel.WithRetry(retry.DefaultChannelOpts))
		if err != nil {
			log.Error(err.Error())
			setOutput(context, eInternalError, "", err.Error(), "")
			return true, err
		}

		status = strconv.FormatInt(int64(resp.ChaincodeStatus), 10)
		payload = string(resp.Payload)
		transactionID = string(resp.TransactionID)

		log.Info("Response Payload : ", payload)

	}

	setOutput(context, status, payload, "", transactionID)
	return true, nil
}

//This will set all the response variable to be sent back.
func setOutput(context activity.Context, status string, payload string, errorMessage string, transactionID string) {
	context.SetOutput(oValueStatus, status)
	context.SetOutput(oValueErrorMessage, errorMessage)
	context.SetOutput(oValueResponsePayload, payload)
	context.SetOutput(oValueTransactionID, transactionID)
}
