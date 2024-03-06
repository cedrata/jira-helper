package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStruct(t *testing.T) {
	type TestStruct1 struct {
		Foo string `validation:"required"`
		Bar string `validation:"required"`
	}

	type TestStruct2 struct {
		Boo string
		Far string `validation:"required"`
	}

	type tester struct {
		name          string
		expectedError bool
		toValidate    interface{}
	}

	tests := []tester{
		{
			name:          "all fields required first missing",
			expectedError: true,
			toValidate: TestStruct1{
				Bar: "not empty string",
			},
		},
		{
			name:          "all fields required second missing",
			expectedError: true,
			toValidate: TestStruct1{
				Foo: "not empty string",
			},
		},
		{
			name:          "all fields required all provided",
			expectedError: false,
			toValidate: TestStruct1{
				Foo: "not empty string",
				Bar: "not empty string",
			},
		},
		{
			name:          "one field required provided",
			expectedError: false,
			toValidate: TestStruct2{
				Far: "not empty string",
			},
		},
		{
			name:          "one field required provided",
			expectedError: false,
			toValidate: TestStruct2{
				Boo: "not empty string",
				Far: "not empty string",
			},
		},
		{
			name:          "one field required not provided",
			expectedError: true,
			toValidate: TestStruct2{
				Boo: "not empty string",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateStruct(test.toValidate)
            t.Log(err)
			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
