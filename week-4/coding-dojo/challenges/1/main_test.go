package main

import (
	"testing"
	"time"
)

func TestGetInfos(t *testing.T) {
	timezone, _ := time.LoadLocation("America/New_York")

	cases := map[string]struct {
		inputBithday   time.Time
		inputName      string
		inputStatus    int
		expectedResult UserResponse
	}{
		"sucesso": {
			inputBithday:   time.Date(1996, 9, 02, 0, 0, 0, 0, timezone),
			inputName:      "leonardo miranda",
			inputStatus:    1,
			expectedResult: UserResponse{lastName: "miranda", age: 26, status: statusEnum[1]},
		},
		"sucesso inactivated": {
			inputBithday:   time.Date(1996, 9, 02, 0, 0, 0, 0, timezone),
			inputName:      "leonardo miranda",
			inputStatus:    1,
			expectedResult: UserResponse{lastName: "miranda", age: 26, status: statusEnum[1]},
		},
		"sucesso": {
			inputBithday:   time.Date(1996, 9, 02, 0, 0, 0, 0, timezone),
			inputName:      "leonardo miranda",
			inputStatus:    1,
			expectedResult: UserResponse{lastName: "miranda", age: 26, status: statusEnum[1]},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			if cs.expectedResult != PrintSomeInfos() {
				t.Error("Unexpected result")
			}
		})
	}
}
