package values

import "testing"

func TestUsageValues(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "learning net values",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UsageValues()
		})
	}
}
