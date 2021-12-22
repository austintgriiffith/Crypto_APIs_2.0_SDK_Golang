/*
CryptoAPIs

Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.

API version: 2.0.0
Contact: developers@cryptoapis.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
	"fmt"
)

// NewConfirmedCoinsTransactionsE409 - struct for NewConfirmedCoinsTransactionsE409
type NewConfirmedCoinsTransactionsE409 struct {
	AlreadyExists *AlreadyExists
	InvalidData *InvalidData
}

// AlreadyExistsAsNewConfirmedCoinsTransactionsE409 is a convenience function that returns AlreadyExists wrapped in NewConfirmedCoinsTransactionsE409
func AlreadyExistsAsNewConfirmedCoinsTransactionsE409(v *AlreadyExists) NewConfirmedCoinsTransactionsE409 {
	return NewConfirmedCoinsTransactionsE409{ AlreadyExists: v}
}

// InvalidDataAsNewConfirmedCoinsTransactionsE409 is a convenience function that returns InvalidData wrapped in NewConfirmedCoinsTransactionsE409
func InvalidDataAsNewConfirmedCoinsTransactionsE409(v *InvalidData) NewConfirmedCoinsTransactionsE409 {
	return NewConfirmedCoinsTransactionsE409{ InvalidData: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NewConfirmedCoinsTransactionsE409) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AlreadyExists
	err = json.Unmarshal(data, &dst.AlreadyExists)
	if err == nil {
		jsonAlreadyExists, _ := json.Marshal(dst.AlreadyExists)
		if string(jsonAlreadyExists) == "{}" { // empty struct
			dst.AlreadyExists = nil
		} else {
			match++
		}
	} else {
		dst.AlreadyExists = nil
	}

	// try to unmarshal data into InvalidData
	err = json.Unmarshal(data, &dst.InvalidData)
	if err == nil {
		jsonInvalidData, _ := json.Marshal(dst.InvalidData)
		if string(jsonInvalidData) == "{}" { // empty struct
			dst.InvalidData = nil
		} else {
			match++
		}
	} else {
		dst.InvalidData = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AlreadyExists = nil
		dst.InvalidData = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(NewConfirmedCoinsTransactionsE409)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(NewConfirmedCoinsTransactionsE409)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NewConfirmedCoinsTransactionsE409) MarshalJSON() ([]byte, error) {
	if src.AlreadyExists != nil {
		return json.Marshal(&src.AlreadyExists)
	}

	if src.InvalidData != nil {
		return json.Marshal(&src.InvalidData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NewConfirmedCoinsTransactionsE409) GetActualInstance() (interface{}) {
	if obj.AlreadyExists != nil {
		return obj.AlreadyExists
	}

	if obj.InvalidData != nil {
		return obj.InvalidData
	}

	// all schemas are nil
	return nil
}

type NullableNewConfirmedCoinsTransactionsE409 struct {
	value *NewConfirmedCoinsTransactionsE409
	isSet bool
}

func (v NullableNewConfirmedCoinsTransactionsE409) Get() *NewConfirmedCoinsTransactionsE409 {
	return v.value
}

func (v *NullableNewConfirmedCoinsTransactionsE409) Set(val *NewConfirmedCoinsTransactionsE409) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedCoinsTransactionsE409) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedCoinsTransactionsE409) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedCoinsTransactionsE409(val *NewConfirmedCoinsTransactionsE409) *NullableNewConfirmedCoinsTransactionsE409 {
	return &NullableNewConfirmedCoinsTransactionsE409{value: val, isSet: true}
}

func (v NullableNewConfirmedCoinsTransactionsE409) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedCoinsTransactionsE409) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


