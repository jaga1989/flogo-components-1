package splitpath

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"io/ioutil"
	"testing"
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

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	///////////////////

	fmt.Println("Splitting a path with normal length, should fill two parts and a fixed path")
	//setup attrs
	tc.SetInput("input", "flogo/device/DEV12345/distance")
	tc.SetInput("delimiter", "/")
	tc.SetInput("fixedpath", "flogo/device")

	act.Eval(tc)

	result := tc.GetOutput("result")
	fixedpath := tc.GetOutput("fixedpath")
	part1 := tc.GetOutput("part1")
	part2 := tc.GetOutput("part2")
	part3 := tc.GetOutput("part3")
	part4 := tc.GetOutput("part4")
	part5 := tc.GetOutput("part5")
	part6 := tc.GetOutput("part6")
	part7 := tc.GetOutput("part7")
	part8 := tc.GetOutput("part8")
	remainder := tc.GetOutput("remainder")

	fmt.Println("result: ", result)
	fmt.Println("fixedpath: ", fixedpath)
	fmt.Println("part1: ", part1)
	fmt.Println("part2: ", part2)
	fmt.Println("part3: ", part3)
	fmt.Println("part4: ", part4)
	fmt.Println("part5: ", part5)
	fmt.Println("part6: ", part6)
	fmt.Println("part7: ", part7)
	fmt.Println("part8: ", part8)
	fmt.Println("remainder: ", remainder)

	///////////////////

	fmt.Println("Splitting a path with high length, should fill all parts and a remainder")
	//setup attrs
	tc.SetInput("input", "flogo/device/DEV/1/2/3/4/5/6/7/8/9/distance")
	tc.SetInput("delimiter", "/")
	tc.SetInput("fixedpath", "flogo/device")

	act.Eval(tc)

	result = tc.GetOutput("result")
	fixedpath = tc.GetOutput("fixedpath")
	part1 = tc.GetOutput("part1")
	part2 = tc.GetOutput("part2")
	part3 = tc.GetOutput("part3")
	part4 = tc.GetOutput("part4")
	part5 = tc.GetOutput("part5")
	part6 = tc.GetOutput("part6")
	part7 = tc.GetOutput("part7")
	part8 = tc.GetOutput("part8")
	remainder = tc.GetOutput("remainder")

	fmt.Println("result: ", result)
	fmt.Println("fixedpath: ", fixedpath)
	fmt.Println("part1: ", part1)
	fmt.Println("part2: ", part2)
	fmt.Println("part3: ", part3)
	fmt.Println("part4: ", part4)
	fmt.Println("part5: ", part5)
	fmt.Println("part6: ", part6)
	fmt.Println("part7: ", part7)
	fmt.Println("part8: ", part8)
	fmt.Println("remainder: ", remainder)

	if result == nil {
		t.Fail()
	}
}
