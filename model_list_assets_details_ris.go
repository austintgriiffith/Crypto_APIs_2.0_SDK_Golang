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

// ListAssetsDetailsRIS - Represents a specific asset's data depending on its type (whether it is \"crypto\" or \"fiat\").
type ListAssetsDetailsRIS struct {
	ListAssetsDetailsRISC *ListAssetsDetailsRISC
}

// ListAssetsDetailsRISCAsListAssetsDetailsRIS is a convenience function that returns ListAssetsDetailsRISC wrapped in ListAssetsDetailsRIS
func ListAssetsDetailsRISCAsListAssetsDetailsRIS(v *ListAssetsDetailsRISC) ListAssetsDetailsRIS {
	return ListAssetsDetailsRIS{ ListAssetsDetailsRISC: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListAssetsDetailsRIS) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListAssetsDetailsRISC
	err = json.Unmarshal(data, &dst.ListAssetsDetailsRISC)
	if err == nil {
		jsonListAssetsDetailsRISC, _ := json.Marshal(dst.ListAssetsDetailsRISC)
		if string(jsonListAssetsDetailsRISC) == "{}" { // empty struct
			dst.ListAssetsDetailsRISC = nil
		} else {
			match++
		}
	} else {
		dst.ListAssetsDetailsRISC = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListAssetsDetailsRISC = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ListAssetsDetailsRIS)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListAssetsDetailsRIS)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListAssetsDetailsRIS) MarshalJSON() ([]byte, error) {
	if src.ListAssetsDetailsRISC != nil {
		return json.Marshal(&src.ListAssetsDetailsRISC)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListAssetsDetailsRIS) GetActualInstance() (interface{}) {
	if obj.ListAssetsDetailsRISC != nil {
		return obj.ListAssetsDetailsRISC
	}

	// all schemas are nil
	return nil
}

type NullableListAssetsDetailsRIS struct {
	value *ListAssetsDetailsRIS
	isSet bool
}

func (v NullableListAssetsDetailsRIS) Get() *ListAssetsDetailsRIS {
	return v.value
}

func (v *NullableListAssetsDetailsRIS) Set(val *ListAssetsDetailsRIS) {
	v.value = val
	v.isSet = true
}

func (v NullableListAssetsDetailsRIS) IsSet() bool {
	return v.isSet
}

func (v *NullableListAssetsDetailsRIS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAssetsDetailsRIS(val *ListAssetsDetailsRIS) *NullableListAssetsDetailsRIS {
	return &NullableListAssetsDetailsRIS{value: val, isSet: true}
}

func (v NullableListAssetsDetailsRIS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAssetsDetailsRIS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

