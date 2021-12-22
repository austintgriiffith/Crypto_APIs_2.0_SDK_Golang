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

// NewConfirmedTokensTransactionsAndEachConfirmationE409 - struct for NewConfirmedTokensTransactionsAndEachConfirmationE409
type NewConfirmedTokensTransactionsAndEachConfirmationE409 struct {
	AlreadyExists *AlreadyExists
	InvalidData *InvalidData
}

// AlreadyExistsAsNewConfirmedTokensTransactionsAndEachConfirmationE409 is a convenience function that returns AlreadyExists wrapped in NewConfirmedTokensTransactionsAndEachConfirmationE409
func AlreadyExistsAsNewConfirmedTokensTransactionsAndEachConfirmationE409(v *AlreadyExists) NewConfirmedTokensTransactionsAndEachConfirmationE409 {
	return NewConfirmedTokensTransactionsAndEachConfirmationE409{ AlreadyExists: v}
}

// InvalidDataAsNewConfirmedTokensTransactionsAndEachConfirmationE409 is a convenience function that returns InvalidData wrapped in NewConfirmedTokensTransactionsAndEachConfirmationE409
func InvalidDataAsNewConfirmedTokensTransactionsAndEachConfirmationE409(v *InvalidData) NewConfirmedTokensTransactionsAndEachConfirmationE409 {
	return NewConfirmedTokensTransactionsAndEachConfirmationE409{ InvalidData: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NewConfirmedTokensTransactionsAndEachConfirmationE409) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(NewConfirmedTokensTransactionsAndEachConfirmationE409)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(NewConfirmedTokensTransactionsAndEachConfirmationE409)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NewConfirmedTokensTransactionsAndEachConfirmationE409) MarshalJSON() ([]byte, error) {
	if src.AlreadyExists != nil {
		return json.Marshal(&src.AlreadyExists)
	}

	if src.InvalidData != nil {
		return json.Marshal(&src.InvalidData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NewConfirmedTokensTransactionsAndEachConfirmationE409) GetActualInstance() (interface{}) {
	if obj.AlreadyExists != nil {
		return obj.AlreadyExists
	}

	if obj.InvalidData != nil {
		return obj.InvalidData
	}

	// all schemas are nil
	return nil
}

type NullableNewConfirmedTokensTransactionsAndEachConfirmationE409 struct {
	value *NewConfirmedTokensTransactionsAndEachConfirmationE409
	isSet bool
}

func (v NullableNewConfirmedTokensTransactionsAndEachConfirmationE409) Get() *NewConfirmedTokensTransactionsAndEachConfirmationE409 {
	return v.value
}

func (v *NullableNewConfirmedTokensTransactionsAndEachConfirmationE409) Set(val *NewConfirmedTokensTransactionsAndEachConfirmationE409) {
	v.value = val
	v.isSet = true
}

func (v NullableNewConfirmedTokensTransactionsAndEachConfirmationE409) IsSet() bool {
	return v.isSet
}

func (v *NullableNewConfirmedTokensTransactionsAndEachConfirmationE409) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewConfirmedTokensTransactionsAndEachConfirmationE409(val *NewConfirmedTokensTransactionsAndEachConfirmationE409) *NullableNewConfirmedTokensTransactionsAndEachConfirmationE409 {
	return &NullableNewConfirmedTokensTransactionsAndEachConfirmationE409{value: val, isSet: true}
}

func (v NullableNewConfirmedTokensTransactionsAndEachConfirmationE409) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewConfirmedTokensTransactionsAndEachConfirmationE409) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


