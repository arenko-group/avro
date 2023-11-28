// Code generated by generatetestcode.go; DO NOT EDIT.

package multiSchemaExternalType

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
                        "name": "F",
                        "type": {
                            "name": "com.heetch.Message",
                            "type": "record",
                            "fields": [
                                {
                                    "name": "Metadata",
                                    "type": {
                                        "name": "Metadata",
                                        "type": "record",
                                        "fields": [
                                            {
                                                "name": "CloudEvent",
                                                "type": {
                                                    "name": "CloudEvent",
                                                    "type": "record",
                                                    "fields": [
                                                        {
                                                            "name": "id",
                                                            "type": "string"
                                                        },
                                                        {
                                                            "name": "source",
                                                            "type": "string"
                                                        },
                                                        {
                                                            "name": "specversion",
                                                            "type": "string"
                                                        },
                                                        {
                                                            "name": "time",
                                                            "type": {
                                                                "type": "long",
                                                                "logicalType": "timestamp-micros"
                                                            }
                                                        }
                                                    ]
                                                }
                                            }
                                        ]
                                    }
                                }
                            ],
                            "go.package": "github.com/arenko-group/avro/internal/testtypes"
                        }
                    },
                    {
                        "name": "G",
                        "type": "com.heetch.CloudEvent"
                    },
                    {
                        "name": "H",
                        "type": "string"
                    }
                ]
            }`,
	GoType: new(R),
	Subtests: []testutil.RoundTripSubtest{{
		TestName: "main",
		InDataJSON: `{
                        "F": {
                            "Metadata": {
                                "CloudEvent": {
                                    "time": 1580486871000000,
                                    "id": "xid",
                                    "source": "xsource",
                                    "specversion": "xspecversion"
                                }
                            }
                        },
                        "G": {
                            "time": 1580495933000000,
                            "id": "yd",
                            "source": "ysource",
                            "specversion": "yspecversion"
                        },
                        "H": "xh"
                    }`,
		OutDataJSON: `{
                        "F": {
                            "Metadata": {
                                "CloudEvent": {
                                    "time": 1580486871000000,
                                    "id": "xid",
                                    "source": "xsource",
                                    "specversion": "xspecversion"
                                }
                            }
                        },
                        "G": {
                            "time": 1580495933000000,
                            "id": "yd",
                            "source": "ysource",
                            "specversion": "yspecversion"
                        },
                        "H": "xh"
                    }`,
	}},
}

func TestGeneratedCode(t *testing.T) {
	tests.Test(t)
}
