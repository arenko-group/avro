// Code generated by generatetestcode.go; DO NOT EDIT.

package mapDefault

import (
	"testing"

	"github.com/arenko-group/avro/cmd/avrogo/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "R",
                "type": "record",
                "fields": [
                    {
                        "name": "_",
                        "type": "int",
                        "default": 0
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "_": 0
                    }`,
		OutDataJSON: `{
                        "mapOfInt": {
                            "a": 2,
                            "b": 5,
                            "c": 99
                        }
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
