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

// SyncHDWalletXPubYPubZPubE401 - struct for SyncHDWalletXPubYPubZPubE401
type SyncHDWalletXPubYPubZPubE401 struct {
	InvalidApiKey *InvalidApiKey
	MissingApiKey *MissingApiKey
}

// InvalidApiKeyAsSyncHDWalletXPubYPubZPubE401 is a convenience function that returns InvalidApiKey wrapped in SyncHDWalletXPubYPubZPubE401
func InvalidApiKeyAsSyncHDWalletXPubYPubZPubE401(v *InvalidApiKey) SyncHDWalletXPubYPubZPubE401 {
	return SyncHDWalletXPubYPubZPubE401{ InvalidApiKey: v}
}

// MissingApiKeyAsSyncHDWalletXPubYPubZPubE401 is a convenience function that returns MissingApiKey wrapped in SyncHDWalletXPubYPubZPubE401
func MissingApiKeyAsSyncHDWalletXPubYPubZPubE401(v *MissingApiKey) SyncHDWalletXPubYPubZPubE401 {
	return SyncHDWalletXPubYPubZPubE401{ MissingApiKey: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SyncHDWalletXPubYPubZPubE401) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(SyncHDWalletXPubYPubZPubE401)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SyncHDWalletXPubYPubZPubE401)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SyncHDWalletXPubYPubZPubE401) MarshalJSON() ([]byte, error) {
	if src.InvalidApiKey != nil {
		return json.Marshal(&src.InvalidApiKey)
	}

	if src.MissingApiKey != nil {
		return json.Marshal(&src.MissingApiKey)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SyncHDWalletXPubYPubZPubE401) GetActualInstance() (interface{}) {
	if obj.InvalidApiKey != nil {
		return obj.InvalidApiKey
	}

	if obj.MissingApiKey != nil {
		return obj.MissingApiKey
	}

	// all schemas are nil
	return nil
}

type NullableSyncHDWalletXPubYPubZPubE401 struct {
	value *SyncHDWalletXPubYPubZPubE401
	isSet bool
}

func (v NullableSyncHDWalletXPubYPubZPubE401) Get() *SyncHDWalletXPubYPubZPubE401 {
	return v.value
}

func (v *NullableSyncHDWalletXPubYPubZPubE401) Set(val *SyncHDWalletXPubYPubZPubE401) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncHDWalletXPubYPubZPubE401) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncHDWalletXPubYPubZPubE401) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncHDWalletXPubYPubZPubE401(val *SyncHDWalletXPubYPubZPubE401) *NullableSyncHDWalletXPubYPubZPubE401 {
	return &NullableSyncHDWalletXPubYPubZPubE401{value: val, isSet: true}
}

func (v NullableSyncHDWalletXPubYPubZPubE401) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncHDWalletXPubYPubZPubE401) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


