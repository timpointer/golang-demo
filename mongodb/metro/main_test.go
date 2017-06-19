package main

import "testing"

func Test_parseChannel(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name        string
		args        args
		wantStore   string
		wantChannel string
	}{
		{"first", args{"170301001CP"}, "010", "CP"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStore, gotChannel, err := parseChannel(tt.args.c)
			if err != nil {
				t.Errorf("%v", err)
			}
			if gotStore != tt.wantStore {
				t.Errorf("parseChannel() gotStore = %v, want %v", gotStore, tt.wantStore)
			}
			if gotChannel != tt.wantChannel {
				t.Errorf("parseChannel() gotChannel = %v, want %v", gotChannel, tt.wantChannel)
			}
		})
	}
}
