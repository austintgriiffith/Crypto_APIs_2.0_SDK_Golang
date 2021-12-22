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

// GetXRPRippleAddressDetailsE401 - struct for GetXRPRippleAddressDetailsE401
type GetXRPRippleAddressDetailsE401 struct {
	InvalidApiKey *InvalidApiKey
	MissingApiKey *MissingApiKey
}

// InvalidApiKeyAsGetXRPRippleAddressDetailsE401 is a convenience function that returns InvalidApiKey wrapped in GetXRPRippleAddressDetailsE401
func InvalidApiKeyAsGetXRPRippleAddressDetailsE401(v *InvalidApiKey) GetXRPRippleAddressDetailsE401 {
	return GetXRPRippleAddressDetailsE401{ InvalidApiKey: v}
}

// MissingApiKeyAsGetXRPRippleAddressDetailsE401 is a convenience function that returns MissingApiKey wrapped in GetXRPRippleAddressDetailsE401
func MissingApiKeyAsGetXRPRippleAddressDetailsE401(v *MissingApiKey) GetXRPRippleAddressDetailsE401 {
	return GetXRPRippleAddressDetailsE401{ MissingApiKey: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetXRPRippleAddressDetailsE401) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(GetXRPRippleAddressDetailsE401)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetXRPRippleAddressDetailsE401)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetXRPRippleAddressDetailsE401) MarshalJSON() ([]byte, error) {
	if src.InvalidApiKey != nil {
		return json.Marshal(&src.InvalidApiKey)
	}

	if src.MissingApiKey != nil {
		return json.Marshal(&src.MissingApiKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetXRPRippleAddressDetailsE401) GetActualInstance() (interface{}) {
	if obj.InvalidApiKey != nil {
		return obj.InvalidApiKey
	}

	if obj.MissingApiKey != nil {
		return obj.MissingApiKey
	}

	// all schemas are nil
	return nil
}

type NullableGetXRPRippleAddressDetailsE401 struct {
	value *GetXRPRippleAddressDetailsE401
	isSet bool
}

func (v NullableGetXRPRippleAddressDetailsE401) Get() *GetXRPRippleAddressDetailsE401 {
	return v.value
}

func (v *NullableGetXRPRippleAddressDetailsE401) Set(val *GetXRPRippleAddressDetailsE401) {
	v.value = val
	v.isSet = true
}

func (v NullableGetXRPRippleAddressDetailsE401) IsSet() bool {
	return v.isSet
}

func (v *NullableGetXRPRippleAddressDetailsE401) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetXRPRippleAddressDetailsE401(val *GetXRPRippleAddressDetailsE401) *NullableGetXRPRippleAddressDetailsE401 {
	return &NullableGetXRPRippleAddressDetailsE401{value: val, isSet: true}
}

func (v NullableGetXRPRippleAddressDetailsE401) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetXRPRippleAddressDetailsE401) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


