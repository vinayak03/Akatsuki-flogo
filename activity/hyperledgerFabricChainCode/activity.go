package hyperledgerFabricChainCode

import (
	"encoding/json"
	"strconv"
	
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	fabric "github.com/vinayak03/Akatsuki-flogo/hyperledgerFabric"
)
import(
	channel "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	retry "github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
)

var (
	log = logger.GetLogger("activity-fabric-chaincode")
)

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
	
	//Initialize the sdk 
	networkConfig := context.GetInput("NetworkConfig").(string)
	SDK, err := fabric.GetSDK(networkConfig)

	if err != nil {
		setOutput(context,"500","",err.Error())
		return true, nil
	}
	
	//Extract the details from activity request.
	requestType := context.GetInput("RequestType").(string)
	chainID := context.GetInput("ChainCodeID").(string)
	channelID := context.GetInput("ChannelID").(string)
	user := context.GetInput("User").(string)
	userPasswd := context.GetInput("UserPasswd").(string)
	userOrg := context.GetInput("UserOrg").(string)
	functionName := context.GetInput("FunctionName").(string)
	inputParams := context.GetInput("Params").(string)
	
	//check if the user is present in local trustsore if not enroll it in local truststore.
	fabric.EnrollWithOrg(SDK, userOrg, user, userPasswd)
	
	//Convert json array to string array
	params := Params{}
	if resp := json.Unmarshal([]byte(inputParams), &params);  resp!=nil {
		setOutput(context,"500","",err.Error())
		return true,nil
	}
	
	paramSize := len(params.Params)
	log.Debug("Param Size ", paramSize)
	var args [][]byte = make([][]byte, paramSize) 
	
	for index,param := range params.Params {
		args[index] = []byte(param)
		log.Debug("Added param ",param," at index ", index)
	}

	channelCtx,err := fabric.GetChannelContext(SDK, channelID, user)
	channelClient,err := channel.New(channelCtx)
	if err!=nil {
		setOutput(context,"500","",err.Error())
		return true, nil
	}
	if requestType == "Query" {
		resp, err := channelClient.Query(channel.Request{
			ChaincodeID: chainID,
			Fcn:         functionName,
			Args:        args},
			channel.WithRetry(retry.DefaultChannelOpts))
		if err != nil {
			setOutput(context,"500","",err.Error())
			return true,nil
		}
		
		status = strconv.FormatInt(int64(resp.ChaincodeStatus),10)
		payload= string(resp.Payload)
		log.Info("Response Payload : ",payload)
		
	} else {
		resp, err := channelClient.Execute(channel.Request{
			ChaincodeID: chainID,
			Fcn:         functionName,
			Args:        args},
			channel.WithRetry(retry.DefaultChannelOpts))
		if err != nil {
			setOutput(context,"500","",err.Error())
			return true,nil
		}
		
		status = strconv.FormatInt(int64(resp.ChaincodeStatus),10)
		payload= string(resp.Payload)
		log.Info("Response Payload : ",payload)
	}
	
	setOutput(context,status,payload,"")
	return true, nil
}

//This will set all the response variable to be sent back.
func setOutput(context activity.Context, status string, payload string, errorMessage string){
		context.SetOutput("status", status)
		context.SetOutput("ErrorMessage", errorMessage)
		context.SetOutput("ResponsePayload", payload)
}
