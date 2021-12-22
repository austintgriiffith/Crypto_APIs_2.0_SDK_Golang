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

// GetLastMinedBlockE401 - struct for GetLastMinedBlockE401
type GetLastMinedBlockE401 struct {
	InvalidApiKey *InvalidApiKey
	MissingApiKey *MissingApiKey
}

// InvalidApiKeyAsGetLastMinedBlockE401 is a convenience function that returns InvalidApiKey wrapped in GetLastMinedBlockE401
func InvalidApiKeyAsGetLastMinedBlockE401(v *InvalidApiKey) GetLastMinedBlockE401 {
	return GetLastMinedBlockE401{ InvalidApiKey: v}
}

// MissingApiKeyAsGetLastMinedBlockE401 is a convenience function that returns MissingApiKey wrapped in GetLastMinedBlockE401
func MissingApiKeyAsGetLastMinedBlockE401(v *MissingApiKey) GetLastMinedBlockE401 {
	return GetLastMinedBlockE401{ MissingApiKey: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetLastMinedBlockE401) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InvalidApiKey
	err = json.Unmarshal(data, &dst.InvalidApiKey)
	if err == nil {
		jsonInvalidApiKey, _ := json.Marshal(dst.InvalidApiKey)
		if string(jsonInvalidApiKey) == "{}" { // empty struct
			dst.InvalidApiKey = nil
		} else {
			match++
		}
	} else {
		dst.InvalidApiKey = nil
	}

	// try to unmarshal data into MissingApiKey
	err = json.Unmarshal(data, &dst.MissingApiKey)
	if err == nil {
		jsonMissingApiKey, _ := json.Marshal(dst.MissingApiKey)
		if string(jsonMissingApiKey) == "{}" { // empty struct
			dst.MissingApiKey = nil
		} else {
			match++
		}
	} else {
		dst.MissingApiKey = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InvalidApiKey = nil
		dst.MissingApiKey = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GetLastMinedBlockE401)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetLastMinedBlockE401)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetLastMinedBlockE401) MarshalJSON() ([]byte, error) {
	if src.InvalidApiKey != nil {
		return json.Marshal(&src.InvalidApiKey)
	}

	if src.MissingApiKey != nil {
		return json.Marshal(&src.MissingApiKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetLastMinedBlockE401) GetActualInstance() (interface{}) {
	if obj.InvalidApiKey != nil {
		return obj.InvalidApiKey
	}

	if obj.MissingApiKey != nil {
		return obj.MissingApiKey
	}

	// all schemas are nil
	return nil
}

type NullableGetLastMinedBlockE401 struct {
	value *GetLastMinedBlockE401
	isSet bool
}

func (v NullableGetLastMinedBlockE401) Get() *GetLastMinedBlockE401 {
	return v.value
}

func (v *NullableGetLastMinedBlockE401) Set(val *GetLastMinedBlockE401) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLastMinedBlockE401) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLastMinedBlockE401) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLastMinedBlockE401(val *GetLastMinedBlockE401) *NullableGetLastMinedBlockE401 {
	return &NullableGetLastMinedBlockE401{value: val, isSet: true}
}

func (v NullableGetLastMinedBlockE401) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLastMinedBlockE401) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

