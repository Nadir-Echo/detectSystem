package models

import (
	"fmt"
	"testing"
)

//test AddResult function
func TestAddResult(t *testing.T) {
	result := &Result{
		DefectResult: "crazing",
		Pid:          1,
		Conf:         0.855555,
		Position:     "11,22,33,55",
	}
	if err := AddResult(result); err != nil {
		t.Errorf("AddResult error: %v", err)
	}
}

// test GetResultByPid function
func TestGetResultByPid(t *testing.T) {
	pid := 30
	result, err := GetResultByPidString(pid)
	if err != nil {
		t.Errorf("GetResultByPid error: %v", err)
	}
	fmt.Println(result)
}

//test GetChartResult function
func TestGetChartResult(t *testing.T) {
	username := "admin"
	result, err := GetChartResult(username)
	if err != nil {
		t.Errorf("GetChartResult error: %v", err)
	}
	fmt.Println(result[1].DefectResult)
}