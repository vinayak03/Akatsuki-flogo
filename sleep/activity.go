package sleep

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"time"
)

var log = logger.GetLogger("activity-sleep")

const (
	oValueCode         = "code"
	oValueErrorMessage = "errorMessage"
	eInternalError     = "700"
	iValueUnit         = "Unit"
	iValueInterval     = "Interval"
	ConstMiliseconds   = "Miliseconds"
	ConstSeconds       = "Seconds"
	ConstMinutes       = "Minutes"
	ConstHours         = "Hours"
)

// SleepActivity is a stub for your Activity implementation
type SleepActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &SleepActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *SleepActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *SleepActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the activity data from the context
	Unit := context.GetInput(iValueUnit).(string)
	Interval := context.GetInput(iValueInterval).(int)

	switch Unit {
	case ConstMiliseconds:
		time.Sleep(time.Duration(Interval) * time.Millisecond)
	case ConstSeconds:
		time.Sleep(time.Duration(Interval) * time.Second)
	case ConstMinutes:
		time.Sleep(time.Duration(Interval) * time.Minute)
	case ConstHours:
		time.Sleep(time.Duration(Interval) * time.Hour)
	default:
		SetOutput(context, "700", "Incorrect Interval")
		return true, fmt.Errorf("Should choose atleast one option for requestType")
	}

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}

func SetOutput(context activity.Context, status string, errorMessage string) {
	context.SetOutput(oValueCode, status)
	context.SetOutput(oValueErrorMessage, errorMessage)
}
