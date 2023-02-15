package main

import (
	"reflect"
	"testing"
)

func TestCompareList(t *testing.T) {

	res := CompareLists(Users, UsersMaster)
	expeteted := []User{
		{
			id:          "1",
			isActivated: true,
		},
	}

	if !reflect.DeepEqual(res, expeteted) {
		t.Error("fail")
	}

}
