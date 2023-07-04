package clihelper

import (
	"testing"
)

func Test_yesNoValidator(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
	}{
		{
			"yes",
			true,
		},
		{
			"y",
			false,
		},
		{
			"Y",
			false,
		},
		{
			"n",
			false,
		},
		{
			"N",
			false,
		},
		{
			"yn",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if err := yesNoValidator(tt.input); (err != nil) != tt.wantErr {
				t.Errorf("yesNoValidator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
