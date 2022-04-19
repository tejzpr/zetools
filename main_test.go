package main

import (
	"testing"

	_ "github.com/tejzpr/zetools-base64"
	_ "github.com/tejzpr/zetools-checkport"
	_ "github.com/tejzpr/zetools-hmac"
	_ "github.com/tejzpr/zetools-ls"
	_ "github.com/tejzpr/zetools-ping"
	_ "github.com/tejzpr/zetools-tail"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
