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

// SyncHDWalletXPubYPubZPubE422 - struct for SyncHDWalletXPubYPubZPubE422
type SyncHDWalletXPubYPubZPubE422 struct {
	InvalidRequestBodyStructure *InvalidRequestBodyStructure
	XpubSyncInProgress *XpubSyncInProgress
}

// InvalidRequestBodyStructureAsSyncHDWalletXPubYPubZPubE422 is a convenience function that returns InvalidRequestBodyStructure wrapped in SyncHDWalletXPubYPubZPubE422
func InvalidRequestBodyStructureAsSyncHDWalletXPubYPubZPubE422(v *InvalidRequestBodyStructure) SyncHDWalletXPubYPubZPubE422 {
	return SyncHDWalletXPubYPubZPubE422{ InvalidRequestBodyStructure: v}
}

// XpubSyncInProgressAsSyncHDWalletXPubYPubZPubE422 is a convenience function that returns XpubSyncInProgress wrapped in SyncHDWalletXPubYPubZPubE422
func XpubSyncInProgressAsSyncHDWalletXPubYPubZPubE422(v *XpubSyncInProgress) SyncHDWalletXPubYPubZPubE422 {
	return SyncHDWalletXPubYPubZPubE422{ XpubSyncInProgress: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SyncHDWalletXPubYPubZPubE422) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InvalidRequestBodyStructure
	err = json.Unmarshal(data, &dst.InvalidRequestBodyStructure)
	if err == nil {
		jsonInvalidRequestBodyStructure, _ := json.Marshal(dst.InvalidRequestBodyStructure)
		if string(jsonInvalidRequestBodyStructure) == "{}" { // empty struct
			dst.InvalidRequestBodyStructure = nil
		} else {
			match++
		}
	} else {
		dst.InvalidRequestBodyStructure = nil
	}

	// try to unmarshal data into XpubSyncInProgress
	err = json.Unmarshal(data, &dst.XpubSyncInProgress)
	if err == nil {
		jsonXpubSyncInProgress, _ := json.Marshal(dst.XpubSyncInProgress)
		if string(jsonXpubSyncInProgress) == "{}" { // empty struct
			dst.XpubSyncInProgress = nil
		} else {
			match++
		}
	} else {
		dst.XpubSyncInProgress = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InvalidRequestBodyStructure = nil
		dst.XpubSyncInProgress = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SyncHDWalletXPubYPubZPubE422)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SyncHDWalletXPubYPubZPubE422)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SyncHDWalletXPubYPubZPubE422) MarshalJSON() ([]byte, error) {
	if src.InvalidRequestBodyStructure != nil {
		return json.Marshal(&src.InvalidRequestBodyStructure)
	}

	if src.XpubSyncInProgress != nil {
		return json.Marshal(&src.XpubSyncInProgress)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SyncHDWalletXPubYPubZPubE422) GetActualInstance() (interface{}) {
	if obj.InvalidRequestBodyStructure != nil {
		return obj.InvalidRequestBodyStructure
	}

	if obj.XpubSyncInProgress != nil {
		return obj.XpubSyncInProgress
	}

	// all schemas are nil
	return nil
}

type NullableSyncHDWalletXPubYPubZPubE422 struct {
	value *SyncHDWalletXPubYPubZPubE422
	isSet bool
}

func (v NullableSyncHDWalletXPubYPubZPubE422) Get() *SyncHDWalletXPubYPubZPubE422 {
	return v.value
}

func (v *NullableSyncHDWalletXPubYPubZPubE422) Set(val *SyncHDWalletXPubYPubZPubE422) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncHDWalletXPubYPubZPubE422) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncHDWalletXPubYPubZPubE422) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncHDWalletXPubYPubZPubE422(val *SyncHDWalletXPubYPubZPubE422) *NullableSyncHDWalletXPubYPubZPubE422 {
	return &NullableSyncHDWalletXPubYPubZPubE422{value: val, isSet: true}
}

func (v NullableSyncHDWalletXPubYPubZPubE422) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncHDWalletXPubYPubZPubE422) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


