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

// ListOmniTransactionsByBlockHeightE403 - struct for ListOmniTransactionsByBlockHeightE403
type ListOmniTransactionsByBlockHeightE403 struct {
	BannedIpAddress *BannedIpAddress
	EndpointNotAllowedForApiKey *EndpointNotAllowedForApiKey
	EndpointNotAllowedForPlan *EndpointNotAllowedForPlan
	FeatureMainnetsNotAllowedForPlan *FeatureMainnetsNotAllowedForPlan
}

// BannedIpAddressAsListOmniTransactionsByBlockHeightE403 is a convenience function that returns BannedIpAddress wrapped in ListOmniTransactionsByBlockHeightE403
func BannedIpAddressAsListOmniTransactionsByBlockHeightE403(v *BannedIpAddress) ListOmniTransactionsByBlockHeightE403 {
	return ListOmniTransactionsByBlockHeightE403{ BannedIpAddress: v}
}

// EndpointNotAllowedForApiKeyAsListOmniTransactionsByBlockHeightE403 is a convenience function that returns EndpointNotAllowedForApiKey wrapped in ListOmniTransactionsByBlockHeightE403
func EndpointNotAllowedForApiKeyAsListOmniTransactionsByBlockHeightE403(v *EndpointNotAllowedForApiKey) ListOmniTransactionsByBlockHeightE403 {
	return ListOmniTransactionsByBlockHeightE403{ EndpointNotAllowedForApiKey: v}
}

// EndpointNotAllowedForPlanAsListOmniTransactionsByBlockHeightE403 is a convenience function that returns EndpointNotAllowedForPlan wrapped in ListOmniTransactionsByBlockHeightE403
func EndpointNotAllowedForPlanAsListOmniTransactionsByBlockHeightE403(v *EndpointNotAllowedForPlan) ListOmniTransactionsByBlockHeightE403 {
	return ListOmniTransactionsByBlockHeightE403{ EndpointNotAllowedForPlan: v}
}

// FeatureMainnetsNotAllowedForPlanAsListOmniTransactionsByBlockHeightE403 is a convenience function that returns FeatureMainnetsNotAllowedForPlan wrapped in ListOmniTransactionsByBlockHeightE403
func FeatureMainnetsNotAllowedForPlanAsListOmniTransactionsByBlockHeightE403(v *FeatureMainnetsNotAllowedForPlan) ListOmniTransactionsByBlockHeightE403 {
	return ListOmniTransactionsByBlockHeightE403{ FeatureMainnetsNotAllowedForPlan: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListOmniTransactionsByBlockHeightE403) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BannedIpAddress
	err = json.Unmarshal(data, &dst.BannedIpAddress)
	if err == nil {
		jsonBannedIpAddress, _ := json.Marshal(dst.BannedIpAddress)
		if string(jsonBannedIpAddress) == "{}" { // empty struct
			dst.BannedIpAddress = nil
		} else {
			match++
		}
	} else {
		dst.BannedIpAddress = nil
	}

	// try to unmarshal data into EndpointNotAllowedForApiKey
	err = json.Unmarshal(data, &dst.EndpointNotAllowedForApiKey)
	if err == nil {
		jsonEndpointNotAllowedForApiKey, _ := json.Marshal(dst.EndpointNotAllowedForApiKey)
		if string(jsonEndpointNotAllowedForApiKey) == "{}" { // empty struct
			dst.EndpointNotAllowedForApiKey = nil
		} else {
			match++
		}
	} else {
		dst.EndpointNotAllowedForApiKey = nil
	}

	// try to unmarshal data into EndpointNotAllowedForPlan
	err = json.Unmarshal(data, &dst.EndpointNotAllowedForPlan)
	if err == nil {
		jsonEndpointNotAllowedForPlan, _ := json.Marshal(dst.EndpointNotAllowedForPlan)
		if string(jsonEndpointNotAllowedForPlan) == "{}" { // empty struct
			dst.EndpointNotAllowedForPlan = nil
		} else {
			match++
		}
	} else {
		dst.EndpointNotAllowedForPlan = nil
	}

	// try to unmarshal data into FeatureMainnetsNotAllowedForPlan
	err = json.Unmarshal(data, &dst.FeatureMainnetsNotAllowedForPlan)
	if err == nil {
		jsonFeatureMainnetsNotAllowedForPlan, _ := json.Marshal(dst.FeatureMainnetsNotAllowedForPlan)
		if string(jsonFeatureMainnetsNotAllowedForPlan) == "{}" { // empty struct
			dst.FeatureMainnetsNotAllowedForPlan = nil
		} else {
			match++
		}
	} else {
		dst.FeatureMainnetsNotAllowedForPlan = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BannedIpAddress = nil
		dst.EndpointNotAllowedForApiKey = nil
		dst.EndpointNotAllowedForPlan = nil
		dst.FeatureMainnetsNotAllowedForPlan = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(ListOmniTransactionsByBlockHeightE403)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListOmniTransactionsByBlockHeightE403)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListOmniTransactionsByBlockHeightE403) MarshalJSON() ([]byte, error) {
	if src.BannedIpAddress != nil {
		return json.Marshal(&src.BannedIpAddress)
	}

	if src.EndpointNotAllowedForApiKey != nil {
		return json.Marshal(&src.EndpointNotAllowedForApiKey)
	}

	if src.EndpointNotAllowedForPlan != nil {
		return json.Marshal(&src.EndpointNotAllowedForPlan)
	}

	if src.FeatureMainnetsNotAllowedForPlan != nil {
		return json.Marshal(&src.FeatureMainnetsNotAllowedForPlan)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListOmniTransactionsByBlockHeightE403) GetActualInstance() (interface{}) {
	if obj.BannedIpAddress != nil {
		return obj.BannedIpAddress
	}

	if obj.EndpointNotAllowedForApiKey != nil {
		return obj.EndpointNotAllowedForApiKey
	}

	if obj.EndpointNotAllowedForPlan != nil {
		return obj.EndpointNotAllowedForPlan
	}

	if obj.FeatureMainnetsNotAllowedForPlan != nil {
		return obj.FeatureMainnetsNotAllowedForPlan
	}

	// all schemas are nil
	return nil
}

type NullableListOmniTransactionsByBlockHeightE403 struct {
	value *ListOmniTransactionsByBlockHeightE403
	isSet bool
}

func (v NullableListOmniTransactionsByBlockHeightE403) Get() *ListOmniTransactionsByBlockHeightE403 {
	return v.value
}

func (v *NullableListOmniTransactionsByBlockHeightE403) Set(val *ListOmniTransactionsByBlockHeightE403) {
	v.value = val
	v.isSet = true
}

func (v NullableListOmniTransactionsByBlockHeightE403) IsSet() bool {
	return v.isSet
}

func (v *NullableListOmniTransactionsByBlockHeightE403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOmniTransactionsByBlockHeightE403(val *ListOmniTransactionsByBlockHeightE403) *NullableListOmniTransactionsByBlockHeightE403 {
	return &NullableListOmniTransactionsByBlockHeightE403{value: val, isSet: true}
}

func (v NullableListOmniTransactionsByBlockHeightE403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOmniTransactionsByBlockHeightE403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


