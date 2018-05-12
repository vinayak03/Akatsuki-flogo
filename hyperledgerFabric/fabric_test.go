package hyperledgerFabric

import (
	 "testing"
)

const(
	GOOD_CONFIG = "../fabric-setup/client/config.yaml"
	BAD_CONFIG = "../fabric-setup/client/cony.yaml"
	SDK_ERROR = "SDK is Not Created"
)

func TestGetSDK(t * testing.T){
	SDK,err := GetSDK(GOOD_CONFIG)
	
	if err !=nil || SDK == nil {
		t.Log(SDK_ERROR,err)
		t.FailNow()
	}
	
	newSDK,err := GetSDK(GOOD_CONFIG)
	if err !=nil || newSDK == nil {
		t.Log(SDK_ERROR)
		t.FailNow()
	}
	
	if newSDK!= SDK{
		t.Log("SDK Retrived is different Old:",SDK, ", New:",newSDK)
		t.FailNow()
	}
}

func TestGetSDKInvalidPath(t * testing.T){
	SDK,err := GetSDK(BAD_CONFIG)
	if SDK != nil || err==nil {
		t.Log("SDK is not supposed to creare with invalid file",err)
		t.Fail()
	}
}

func TestGetMSP(t *testing.T){
	SDK,err := GetSDK(GOOD_CONFIG)
	
	if SDK == nil {
		t.Log(SDK_ERROR,err)
		t.Fail()
	}
	msp,err := GetLocalMSP(SDK,"Org1")
	if msp==nil || err !=nil {
		t.Log("Did not get a MSP for Org1")
		t.Fail()
	}	
}

func TestGetResourceMgmt(t *testing.T){
	SDK,err := GetSDK(GOOD_CONFIG)
	
	if SDK==nil{
		t.Log(SDK_ERROR,err)
		t.FailNow()
	}
	
	context,err :=GetUserContext(SDK,"Admin", "Org1")
	if err!=nil {
		t.Log("Error Getting Context for Admin User")
		t.FailNow()
	}
	
	resMgmtClient,err :=GetResourceMgmtClient(SDK,context)
	
	if err!=nil || resMgmtClient==nil {
		t.Log("Error Getting a Resource Management Client")
		t.Fail()
	}	
}

func TestEnroll(t * testing.T){
	SDK,err := GetSDK(GOOD_CONFIG)
	
	if SDK==nil{
		t.Log(SDK_ERROR)
		t.FailNow()
	}
	err = EnrollWithOrg(SDK,"Org1", "admin", "adminpw")
	if err!=nil{
		t.Log("Error Enrolling a Default User",err)
		t.Fail()
	}
	CloseSDK(SDK)
}
