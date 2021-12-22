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

// ListOmniTransactionsByAddressE403 - struct for ListOmniTransactionsByAddressE403
type ListOmniTransactionsByAddressE403 struct {
	BannedIpAddress *BannedIpAddress
	EndpointNotAllowedForApiKey *EndpointNotAllowedForApiKey
	EndpointNotAllowedForPlan *EndpointNotAllowedForPlan
	FeatureMainnetsNotAllowedForPlan *FeatureMainnetsNotAllowedForPlan
}

// BannedIpAddressAsListOmniTransactionsByAddressE403 is a convenience function that returns BannedIpAddress wrapped in ListOmniTransactionsByAddressE403
func BannedIpAddressAsListOmniTransactionsByAddressE403(v *BannedIpAddress) ListOmniTransactionsByAddressE403 {
	return ListOmniTransactionsByAddressE403{ BannedIpAddress: v}
}

// EndpointNotAllowedForApiKeyAsListOmniTransactionsByAddressE403 is a convenience function that returns EndpointNotAllowedForApiKey wrapped in ListOmniTransactionsByAddressE403
func EndpointNotAllowedForApiKeyAsListOmniTransactionsByAddressE403(v *EndpointNotAllowedForApiKey) ListOmniTransactionsByAddressE403 {
	return ListOmniTransactionsByAddressE403{ EndpointNotAllowedForApiKey: v}
}

// EndpointNotAllowedForPlanAsListOmniTransactionsByAddressE403 is a convenience function that returns EndpointNotAllowedForPlan wrapped in ListOmniTransactionsByAddressE403
func EndpointNotAllowedForPlanAsListOmniTransactionsByAddressE403(v *EndpointNotAllowedForPlan) ListOmniTransactionsByAddressE403 {
	return ListOmniTransactionsByAddressE403{ EndpointNotAllowedForPlan: v}
}

// FeatureMainnetsNotAllowedForPlanAsListOmniTransactionsByAddressE403 is a convenience function that returns FeatureMainnetsNotAllowedForPlan wrapped in ListOmniTransactionsByAddressE403
func FeatureMainnetsNotAllowedForPlanAsListOmniTransactionsByAddressE403(v *FeatureMainnetsNotAllowedForPlan) ListOmniTransactionsByAddressE403 {
	return ListOmniTransactionsByAddressE403{ FeatureMainnetsNotAllowedForPlan: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListOmniTransactionsByAddressE403) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(ListOmniTransactionsByAddressE403)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(ListOmniTransactionsByAddressE403)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListOmniTransactionsByAddressE403) MarshalJSON() ([]byte, error) {
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
func (obj *ListOmniTransactionsByAddressE403) GetActualInstance() (interface{}) {
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

type NullableListOmniTransactionsByAddressE403 struct {
	value *ListOmniTransactionsByAddressE403
	isSet bool
}

func (v NullableListOmniTransactionsByAddressE403) Get() *ListOmniTransactionsByAddressE403 {
	return v.value
}

func (v *NullableListOmniTransactionsByAddressE403) Set(val *ListOmniTransactionsByAddressE403) {
	v.value = val
	v.isSet = true
}

func (v NullableListOmniTransactionsByAddressE403) IsSet() bool {
	return v.isSet
}

func (v *NullableListOmniTransactionsByAddressE403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOmniTransactionsByAddressE403(val *ListOmniTransactionsByAddressE403) *NullableListOmniTransactionsByAddressE403 {
	return &NullableListOmniTransactionsByAddressE403{value: val, isSet: true}
}

func (v NullableListOmniTransactionsByAddressE403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOmniTransactionsByAddressE403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


