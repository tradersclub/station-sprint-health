package main

import (
	"testing"
	"time"
)

func TestGetInfos(t *testing.T) {
	timezone, _ := time.LoadLocation("America/New_York")

	var cases = map[string]struct {
		ageOutput      int
		lastnameOutput string
		statusOutput   string
		input          User
	}{
		"case1": {
			ageOutput:      26,
			lastnameOutput: "miranda",
			statusOutput:   "activated",
			input: User{
				name:     "leonardo miranda",
				birthday: time.Date(1996, 9, 02, 0, 0, 0, 0, timezone),
				status:   1,
			},
		},
		"case2": {
			ageOutput:      26,
			lastnameOutput: "miranda",
			statusOutput:   "inactivated",
			input: User{
				name:     "leonardo miranda",
				birthday: time.Date(1996, 9, 02, 0, 0, 0, 0, timezone),
				status:   0,
			},
		},
		"case3": {
			ageOutput:      26,
			lastnameOutput: "miranda",
			statusOutput:   "status not found",
			input: User{
				name:     "leonardo miranda",
				birthday: time.Date(1996, 9, 02, 0, 0, 0, 0, timezone),
				status:   33,
			},
		},
	}

	for _, c := range cases {
		age, lastname, status := getInfos(c.input)

		if age != c.ageOutput {
			t.Error("unexpeted age")
		}
		if lastname != c.lastnameOutput {
			t.Error("unexpeted lastname")
		}
		if status != c.statusOutput {
			t.Error("unexpeted status")
		}
	}

}

// BenchmarkGetInfos-16    	 5445838	       213.2 ns/op	      32 B/op	       1 allocs/op
// BenchmarkGetInfos-16    	 5080663	       208.7 ns/op	      32 B/op	       1 allocs/op
func BenchmarkGetInfos(b *testing.B) {
	timezone, _ := time.LoadLocation("America/New_York")
	for i := 0; i < b.N; i++ {
		getInfos(User{
			name:     "leonardo miranda",
			birthday: time.Date(1996, 9, 02, 0, 0, 0, 0, timezone),
			status:   1,
		})
	}
}
