package hyperledgerFabricList

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	fabric "github.com/vinayak03/Akatsuki-flogo/hyperledgerFabric"
)

var (
	log = logger.GetLogger("activity-fabric-list")
)

const (
	iValueNetworkConfig        = "NetworkConfig"
	iValueResourceType         = "ResourceType"
	iValueAdminUser            = "AdminUser"
	iValueAdminPasswd          = "AdminPasswd"
	iValueQueryParam           = "QueryParam"
	iValueOrg                  = "Org"
	oValueStatus               = "Status"
	oValueErrorMessage         = "ErrorMessage"
	oValueResponsePayload      = "ResponsePayload"
	ConstChannel               = "Channel"
	ConstInstalledChainCode    = "InstalledChainCode"
	ConstInstantiatedChainCode = "InstantiatedChainCode"
	iValueChannelID            = "ChannelID"
)

// FabricListActivity is a stub for your Activity implementation
type FabricListActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &FabricListActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *FabricListActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *FabricListActivity) Eval(context activity.Context) (done bool, err error) {
	//Initialize the sdk
	networkConfig := context.GetInput(iValueNetworkConfig).(string)
	SDK, err := fabric.GetSDK(networkConfig)

	if err != nil || SDK == nil {
		setOutput(context, "", "", err.Error())
		return true, err
	}
	adminUser := context.GetInput(iValueAdminUser).(string)
	adminPasswd := context.GetInput(iValueAdminPasswd).(string)
	org := context.GetInput(iValueOrg).(string)

	//Get a Channel List to confirm channel exists
	err = fabric.EnrollWithOrg(SDK, org, adminUser, adminPasswd)
	if err != nil {
		log.Error(err.Error())
		return true, err
	}

	adminContext, err := fabric.GetUserContext(SDK, adminUser, org)
	if err != nil {
		log.Error(err.Error())
		return true, err
	}

	orgResMgmt, err := fabric.GetResourceMgmtClient(SDK, adminContext)
	if err != nil {
		log.Error(err.Error())
		return true, err
	}
	var output string
	queryParam := context.GetInput(iValueQueryParam).(string)
	resourceType := context.GetInput(iValueResourceType).(string)

	switch resourceType {
	case ConstChannel:
		channelList, err := orgResMgmt.QueryChannels(resmgmt.WithTargetEndpoints(queryParam))
		if err != nil {
			log.Error(err.Error())
			return true, err
		}
		output = channelList.String()
		log.Info("List of Channels :", channelList.String())
	case ConstInstalledChainCode:
		chainCodeQueryResponse, err := orgResMgmt.QueryInstalledChaincodes(resmgmt.WithTargetEndpoints(queryParam))
		if err != nil {
			log.Error(err.Error())
			return true, err
		}
		output = chainCodeQueryResponse.String()
		log.Info("List of Installed Chaincodes :", chainCodeQueryResponse.String())
	case ConstInstantiatedChainCode:
		channelID := context.GetInput(iValueChannelID).(string)

		chainCodeQueryResponse, err := orgResMgmt.QueryInstantiatedChaincodes(channelID, resmgmt.WithTargetEndpoints(queryParam))
		if err != nil {
			log.Error(err.Error())
			return true, err
		}
		output = chainCodeQueryResponse.String()
		log.Info("List of Instantiated Chaincodes :", chainCodeQueryResponse.String())
	}

	setOutput(context, "200", output, "")
	return true, nil
}

func setOutput(context activity.Context, status string, payload string, errorMessage string) {
	context.SetOutput(oValueStatus, status)
	context.SetOutput(oValueErrorMessage, errorMessage)
	context.SetOutput(oValueResponsePayload, payload)
}
