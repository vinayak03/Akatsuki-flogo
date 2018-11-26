package sleep

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"time"
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

func TestGetProduct(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	setInput(tc, ConstSeconds, 1)

	var startTime = time.Now()
	act.Eval(tc)
	var endTime = time.Now()

	fmt.Print()

	assert.True(t, endTime.Sub(startTime) > 1, "Not Equal")

}

func setInput(tc *test.TestActivityContext, Unit string, Interval int) {
	tc.SetInput(iValueUnit, Unit)
	tc.SetInput(iValueInterval, Interval)

}
