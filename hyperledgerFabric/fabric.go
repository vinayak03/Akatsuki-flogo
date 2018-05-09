package hyperledgerFabric

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	clientmsp "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"os"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	SDK_INIT_ERROR= "SDK Not Initialized"
)

var (
	SDKStore map [string]*fabsdk.FabricSDK= make(map [string]*fabsdk.FabricSDK);
	log = logger.GetLogger("activity-fabric-chaincode")
)

func GetSDK(env string) (* fabsdk.FabricSDK, error) {
	//if SDK==nil {
		configLocation := os.Getenv(env)
		if configLocation==""{
			return nil,fmt.Errorf("Invalid Environment Variable")
		}
		//caching the sdk instances based on the config file name (need to correct the paths before using it as key)
		SDK := SDKStore[configLocation]
		if SDK ==nil {
			fmt.Println("SDK Instance not found creating a new one")
			SDK,err := fabsdk.New(config.FromFile(configLocation))
			
			if err!=nil {
				return nil,fmt.Errorf("Error in SDK Initialization : %s",err)
			}
			
			SDKStore[configLocation]=SDK
			return SDK,nil
		}
		//SDK Already Present for the config location
		return SDK,nil
}

func CloseSDK(SDK *fabsdk.FabricSDK){
	if SDK!=nil {
		SDK.Close()
	}	
	SDK = nil
}

func GetLocalMSP(SDK *fabsdk.FabricSDK,org string) (* clientmsp.Client,error) {
	//Get a connection to MSP
	localMSP,err := clientmsp.New(SDK.Context(),clientmsp.WithOrg(org))
	if err !=nil {
		return nil,fmt.Errorf("Error Getting MSP : %s",err)
	}
	return localMSP,nil	
}

func EnrollWithMSP(SDK *fabsdk.FabricSDK,localMSP * clientmsp.Client,username string, password string) error{
	identity,err := localMSP.GetSigningIdentity(username)
	
	if err!=nil || identity==nil{
		fmt.Println("Username: ",username,"\nPassword: ",password,clientmsp.WithSecret(password))
		err := localMSP.Enroll(username,clientmsp.WithSecret(password))
		if err!=nil{
			return fmt.Errorf("Error enrolling Admin : %s",err)
		}
	}
	
	return nil
}

func EnrollWithOrg(SDK *fabsdk.FabricSDK,org string,username string, password string) error {
	localMSP,err := GetLocalMSP(SDK,org)
	if err!=nil {
		return fmt.Errorf("Error Getting MSP for Org %s : %s", org,err)
	}
	return EnrollWithMSP(SDK,localMSP,username,password)
}

func GetChannelContext(SDK *fabsdk.FabricSDK,channelID string, username string ) (context.ChannelProvider,error){
	channelContext := SDK.ChannelContext(channelID, fabsdk.WithUser(username))
	if channelContext==nil {
		return nil,fmt.Errorf("Error Getting a  Channel Context for ChannelID : %s and Username:  %s", channelID,username)
	}
	
	return channelContext,nil
}

func GetUserContext(SDK *fabsdk.FabricSDK,user string, org string) (context.ClientProvider,error){
	userContext := SDK.Context(fabsdk.WithUser(user), fabsdk.WithOrg(org))
	
	if userContext == nil {
		return nil,fmt.Errorf("Error getting a Usercontext for User: %s and Org : %s ", user,org)	
	}
	return userContext,nil
}

func GetResourceMgmtClient(SDK *fabsdk.FabricSDK,context context.ClientProvider) (*resmgmt.Client,error){
	resMgmt, err := resmgmt.New(context)
	if err!=nil || resMgmt==nil {
		return nil,fmt.Errorf("Error Creating Resource Management Client")
	}
	return resMgmt,nil
}