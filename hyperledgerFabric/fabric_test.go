package hyperledgerFabric

import (
	 "testing"
	 "os"
)

func TestGetSDK(t * testing.T){
	os.Setenv("fabric-config","D:/Projects/Flogo-Hackthon/go-ws/config.yaml")
	SDK,err := GetSDK("fabric-config")
	
	if err !=nil || SDK == nil {
		t.Log("SDK is not Initialized")
		t.FailNow()
	}
	
	newSDK,err := GetSDK("fabric-config")
	if err !=nil || newSDK == nil {
		t.Log("SDK is not Initialized")
		t.FailNow()
	}
	
	if newSDK!= SDK{
		t.Log("SDK Retrived is different Old:",SDK, ", New:",newSDK)
		t.FailNow()
	}
}

func TestGetSDKInvalidPath(t * testing.T){
	os.Setenv("fabric-config-2","D:/Projects/Flogo-Hackthon/go-ws/cony.yaml")
	SDK,err := GetSDK("fabric-config-2")
	if SDK != nil || err==nil {
		t.Log("SDK is not Initialized",err)
		t.Fail()
	}
}

func TestGetMSP(t *testing.T){
	os.Setenv("fabric-config-3","D:/Projects/Flogo-Hackthon/go-ws/config.yaml")
	SDK,err := GetSDK("fabric-config-3")
	if SDK == nil {
		t.Log("SDK is not Initialized")
		t.Fail()
	}
	msp,err := GetLocalMSP(SDK,"Org1")
	if msp==nil || err !=nil {
		t.Log("Did not get a MSP for Org1")
		t.Fail()
	}	
}

func TestGetResourceMgmt(t *testing.T){
	os.Setenv("fabric-config-4","D:/Projects/Flogo-Hackthon/go-ws/config.yaml")
	SDK,err := GetSDK("fabric-config-4")
	if SDK==nil{
		t.Log("SDK is not Initialized")
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
	os.Setenv("fabric-config-5","D:/Projects/Flogo-Hackthon/go-ws/config.yaml")
	SDK,err := GetSDK("fabric-config-5")
	if SDK==nil{
		t.Log("SDK is not Initialized")
		t.FailNow()
	}
	err = EnrollWithOrg(SDK,"Org1", "admin", "adminpw")
	if err!=nil{
		t.Log("Error Enrolling a Default User",err)
		t.Fail()
	}
	CloseSDK(SDK)
}
