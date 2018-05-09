package hybris

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-hybris")

// HybrisActivity is a stub for your Activity implementation
type HybrisActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &HybrisActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *HybrisActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *HybrisActivity) Eval(context activity.Context) (done bool, err error)  {

	// Get the activity data from the context
	name := context.GetInput("name").(string)
	salutation := context.GetInput("salutation").(string)

	// Use the log object to log the greeting
	log.Debugf("The Flogo engine says [%s] to [%s]", salutation, name)

	// Set the result as part of the context
	context.SetOutput("result", "The Flogo engine says "+salutation+" to "+name)

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}
