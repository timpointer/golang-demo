package goldenfile

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

func TestAdd(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected string
	}{
		{"foo", "1", "1", "11"},
		{"bar", "1", "-1", "1-1"},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			actual := tc.A + tc.B
			golden := filepath.Join("test-fixtures", tc.Name+".golden")
			if *update {
				ioutil.WriteFile(golden, []byte(actual), 0644)
			}
			Expected, _ := ioutil.ReadFile(golden)
			if !bytes.Equal([]byte(actual), Expected) {
				t.Fatalf("failure actual %v want %v", actual, string(Expected))
			}
		})
	}
}
