/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// V1AnyValue - AnyValue is used to represent any type of attribute value. AnyValue may contain a primitive value such as a string or integer or it may contain an arbitrary nested object containing arrays, key-value lists and primitives.
type V1AnyValue struct {
	StringValue string `json:"stringValue,omitempty"`

	BoolValue bool `json:"boolValue,omitempty"`

	IntValue string `json:"intValue,omitempty"`

	DoubleValue float64 `json:"doubleValue,omitempty"`

	ArrayValue V1ArrayValue `json:"arrayValue,omitempty"`

	KvlistValue V1KeyValueList `json:"kvlistValue,omitempty"`

	BytesValue string `json:"bytesValue,omitempty"`
}

// AssertV1AnyValueRequired checks if the required fields are not zero-ed
func AssertV1AnyValueRequired(obj V1AnyValue) error {
	if err := AssertV1ArrayValueRequired(obj.ArrayValue); err != nil {
		return err
	}
	if err := AssertV1KeyValueListRequired(obj.KvlistValue); err != nil {
		return err
	}
	return nil
}

// AssertRecurseV1AnyValueRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of V1AnyValue (e.g. [][]V1AnyValue), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseV1AnyValueRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aV1AnyValue, ok := obj.(V1AnyValue)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertV1AnyValueRequired(aV1AnyValue)
	})
}