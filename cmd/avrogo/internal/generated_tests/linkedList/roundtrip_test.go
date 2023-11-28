// Code generated by generatetestcode.go; DO NOT EDIT.

package linkedList

import (
	"testing"

	"github.com/arenko-group/avro/cmd/avrogo/internal/testutil"
)

var tests = testutil.RoundTripTest{
	InSchema: `{
                "name": "List",
                "type": "record",
                "fields": [
                    {
                        "name": "Item",
                        "type": "int"
                    },
                    {
                        "name": "Next",
                        "type": [
                            "null",
                            "List"
                        ],
                        "default": null
                    }
                ]
            }`,
	GoType: new(List),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "Item": 1234,
                        "Next": {
                            "List": {
                                "Item": 9999,
                                "Next": null
                            }
                        }
                    }`,
		OutDataJSON: `{
                        "Item": 1234,
                        "Next": {
                            "List": {
                                "Item": 9999,
                                "Next": null
                            }
                        }
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
