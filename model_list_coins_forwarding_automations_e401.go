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

// ListCoinsForwardingAutomationsE401 - struct for ListCoinsForwardingAutomationsE401
type ListCoinsForwardingAutomationsE401 struct {
	InvalidApiKey *InvalidApiKey
	MissingApiKey *MissingApiKey
}

// InvalidApiKeyAsListCoinsForwardingAutomationsE401 is a convenience function that returns InvalidApiKey wrapped in ListCoinsForwardingAutomationsE401
func InvalidApiKeyAsListCoinsForwardingAutomationsE401(v *InvalidApiKey) ListCoinsForwardingAutomationsE401 {
	return ListCoinsForwardingAutomationsE401{ InvalidApiKey: v}
}

// MissingApiKeyAsListCoinsForwardingAutomationsE401 is a convenience function that returns MissingApiKey wrapped in ListCoinsForwardingAutomationsE401
func MissingApiKeyAsListCoinsForwardingAutomationsE401(v *MissingApiKey) ListCoinsForwardingAutomationsE401 {
	return ListCoinsForwardingAutomationsE401{ MissingApiKey: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListCoinsForwardingAutomationsE401) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(ListCoinsForwardingAutomationsE401)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListCoinsForwardingAutomationsE401)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListCoinsForwardingAutomationsE401) MarshalJSON() ([]byte, error) {
	if src.InvalidApiKey != nil {
		return json.Marshal(&src.InvalidApiKey)
	}

	if src.MissingApiKey != nil {
		return json.Marshal(&src.MissingApiKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListCoinsForwardingAutomationsE401) GetActualInstance() (interface{}) {
	if obj.InvalidApiKey != nil {
		return obj.InvalidApiKey
	}

	if obj.MissingApiKey != nil {
		return obj.MissingApiKey
	}

	// all schemas are nil
	return nil
}

type NullableListCoinsForwardingAutomationsE401 struct {
	value *ListCoinsForwardingAutomationsE401
	isSet bool
}

func (v NullableListCoinsForwardingAutomationsE401) Get() *ListCoinsForwardingAutomationsE401 {
	return v.value
}

func (v *NullableListCoinsForwardingAutomationsE401) Set(val *ListCoinsForwardingAutomationsE401) {
	v.value = val
	v.isSet = true
}

func (v NullableListCoinsForwardingAutomationsE401) IsSet() bool {
	return v.isSet
}

func (v *NullableListCoinsForwardingAutomationsE401) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCoinsForwardingAutomationsE401(val *ListCoinsForwardingAutomationsE401) *NullableListCoinsForwardingAutomationsE401 {
	return &NullableListCoinsForwardingAutomationsE401{value: val, isSet: true}
}

func (v NullableListCoinsForwardingAutomationsE401) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCoinsForwardingAutomationsE401) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

