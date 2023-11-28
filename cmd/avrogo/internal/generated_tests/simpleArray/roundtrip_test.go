// Code generated by generatetestcode.go; DO NOT EDIT.

package simpleArray

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
                        "name": "A",
                        "type": {
                            "type": "array",
                            "items": "int"
                        }
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "empty",
		InDataJSON: `{
                        "A": []
                    }`,
		OutDataJSON: `{
                        "A": []
                    }`,
	}, {
		TestName: "non_empty",
		InDataJSON: `{
                        "A": [
                            23,
                            57,
                            444
                        ]
                    }`,
		OutDataJSON: `{
                        "A": [
                            23,
                            57,
                            444
                        ]
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
