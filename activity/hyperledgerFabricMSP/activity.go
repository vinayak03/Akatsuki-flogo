package hyperledgerFabricMSP

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	fabric "github.com/vinayak03/Akatsuki-flogo/hyperledgerFabric"
	clientmsp "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"fmt"
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
	iValueUserPassword   = "UserPassword"
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
		setOutput(context,"", err.Error())
		return true, nil
	}
	
	userOrg := context.GetInput(iValueAffiliation).(string)
	username := context.GetInput(iValueUsername).(string)
	password := context.GetInput(iValueUserPassword).(string)
		
    err = fabric.EnrollWithOrg(SDK, userOrg, username, password)
    if err != nil {
    	setOutput(context, "", err.Error())
		return true, nil
    }
    
	return true, nil
}

func RevokeUser(SDK *fabsdk.FabricSDK,NameRevoke string, SerialRevoke string, AKIRevoke string, ReasonRevoke string, CANameRevoke string , org string) (err error){
	log.Info("Revoked successfully")
	RevocationRequest := clientmsp.RevocationRequest{
		Name: NameRevoke, 
		Serial: SerialRevoke, 
		AKI: AKIRevoke, 
		Reason: ReasonRevoke, 
		CAName: CANameRevoke}
	localMSPRes, localMSPerr := fabric.GetLocalMSP(SDK, org)
	
	if localMSPerr !=nil {
		return fmt.Errorf("Error Getting MSP : %s",localMSPerr.Error())
	}
		revokeusr, err := localMSPRes.Revoke(&RevocationRequest)	
		
	if err!=nil || revokeusr==nil{		
		//setOutput(context, "", err.Error())
		return  nil
	}	
	log.Info("Revoked successfully")
	return nil
}

func setOutput(context activity.Context, status string, errorMessage string) {

}
