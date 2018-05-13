package hyperledgerFabricMSP

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	fabric "github.com/vinayak03/Akatsuki-flogo/hyperledgerFabric"
)

var (
	log = logger.GetLogger("activity-fabric-msp")
)

const (
	iValueNetworkConfig  = "NetworkConfig"
	iValueRequestType    = "RequestType"
	iValueAdminUser      = "AdminUser"
	iValueAdminPassword  = "AdminPassword"
	iValueUsername       = "Username"
	iValueTypeOfUser     = "TypeOfUser"
	iValueMaxEnrollments = "MaxEnrollments"
	iValueAffiliation    = "Affiliation"
	iValueCAName         = "CAName"
	oValueStatus         = "status"
)

// FabricMSPActivity is a stub for your Activity implementation
type FabricMSPActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &FabricMSPActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *FabricMSPActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *FabricMSPActivity) Eval(context activity.Context) (done bool, err error) {
	//Initialize the sdk
	networkConfig := context.GetInput(iValueNetworkConfig).(string)
	SDK, err := fabric.GetSDK(networkConfig)

	if err != nil || SDK == nil {
		setOutput(context, "", "", err.Error(), "")
		return true, nil
	}

	return true, nil
}

func setOutput(context activity.Context, status string, errorMessage string) {

}
