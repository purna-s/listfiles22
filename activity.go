package listfiles22

import (

	"fmt"
    "os"
	"strings"
	"time"
    "path/filepath"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-listfiles22")

// MyActivity is a stub for your Activity implementation
type listfiles22 struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &listfiles22{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *listfiles22) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *listfiles22) Eval(ctx activity.Context) (done bool, err error) {
	
	
		loc := ctx.GetInput("Path").(string)
		subs := ctx.GetInput("SubDirectories[Y/N]").(string)
	
		dt := time.Now()
	
	// the function that handles each file or dir
	err = filepath.Walk(loc, func(pathX string, infoX os.FileInfo, errX error) error {

		if errX != nil {
			fmt.Println("error at a path \n", errX, pathX)
			return errX
		}

		if infoX.IsDir() {
			fmt.Println("\n'", pathX, "'", " is a directory.\n")
		} else if subs == "Y" {
				fmt.Println("FileName", infoX.Name())
				fmt.Println("Directory", filepath.Dir(pathX))
				fmt.Println("Extension", filepath.Ext(pathX))
				fmt.Println("Size", infoX.Size())
				fmt.Println("ModTime", infoX.ModTime())
				
				diff := dt.Sub(infoX.ModTime())
				mins := int(diff.Minutes())
					fmt.Println("MinutesDiff", mins)
			} else if filepath.Dir(pathX) == strings.Replace(loc, "/", "\\", -1) {
					fmt.Println("FileName", infoX.Name())
					fmt.Println("Directory", filepath.Dir(pathX))
					fmt.Println("Extension", filepath.Ext(pathX))
					fmt.Println("Size", infoX.Size())
					fmt.Println("ModTime", infoX.ModTime())
					
					diff := dt.Sub(infoX.ModTime())
					mins := int(diff.Minutes())
						fmt.Println("MinutesDiff", mins)
				}
	return nil
   })

	if err != nil {
		fmt.Println("error walking the path : \n", loc, err)
	}

	activityLog.Debugf("Activity has listed out the files Successfully")
	fmt.Println("Activity has listed out the files Successfully")
	
	return true, nil
}

