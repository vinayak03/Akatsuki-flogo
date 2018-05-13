package hyperledgerFabricMSP

import (
//	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
//	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	fabric "github.com/vinayak03/Akatsuki-flogo/hyperledgerFabric"
)

const (
	CONFIG = "C:/Users/rvadde/Akatsuki-flogo/fabric-setup/client/config.yaml"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}



func TestRevoke(t *testing.T) {
	
	SDK, err := fabric.GetSDK(CONFIG)
	if err==nil{
	revokeUsr := RevokeUser(SDK, "User", "1234", "0", "User not granted access", "CA","Affiliation")
	 if revokeUsr == nil {
		t.Log("Access revoked")
		return
	}
	}
	
	
}
