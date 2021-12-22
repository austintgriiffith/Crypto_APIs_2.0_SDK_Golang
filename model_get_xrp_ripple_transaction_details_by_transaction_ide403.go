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

// GetXRPRippleTransactionDetailsByTransactionIDE403 - struct for GetXRPRippleTransactionDetailsByTransactionIDE403
type GetXRPRippleTransactionDetailsByTransactionIDE403 struct {
	BannedIpAddress *BannedIpAddress
	EndpointNotAllowedForApiKey *EndpointNotAllowedForApiKey
	EndpointNotAllowedForPlan *EndpointNotAllowedForPlan
	FeatureMainnetsNotAllowedForPlan *FeatureMainnetsNotAllowedForPlan
}

// BannedIpAddressAsGetXRPRippleTransactionDetailsByTransactionIDE403 is a convenience function that returns BannedIpAddress wrapped in GetXRPRippleTransactionDetailsByTransactionIDE403
func BannedIpAddressAsGetXRPRippleTransactionDetailsByTransactionIDE403(v *BannedIpAddress) GetXRPRippleTransactionDetailsByTransactionIDE403 {
	return GetXRPRippleTransactionDetailsByTransactionIDE403{ BannedIpAddress: v}
}

// EndpointNotAllowedForApiKeyAsGetXRPRippleTransactionDetailsByTransactionIDE403 is a convenience function that returns EndpointNotAllowedForApiKey wrapped in GetXRPRippleTransactionDetailsByTransactionIDE403
func EndpointNotAllowedForApiKeyAsGetXRPRippleTransactionDetailsByTransactionIDE403(v *EndpointNotAllowedForApiKey) GetXRPRippleTransactionDetailsByTransactionIDE403 {
	return GetXRPRippleTransactionDetailsByTransactionIDE403{ EndpointNotAllowedForApiKey: v}
}

// EndpointNotAllowedForPlanAsGetXRPRippleTransactionDetailsByTransactionIDE403 is a convenience function that returns EndpointNotAllowedForPlan wrapped in GetXRPRippleTransactionDetailsByTransactionIDE403
func EndpointNotAllowedForPlanAsGetXRPRippleTransactionDetailsByTransactionIDE403(v *EndpointNotAllowedForPlan) GetXRPRippleTransactionDetailsByTransactionIDE403 {
	return GetXRPRippleTransactionDetailsByTransactionIDE403{ EndpointNotAllowedForPlan: v}
}

// FeatureMainnetsNotAllowedForPlanAsGetXRPRippleTransactionDetailsByTransactionIDE403 is a convenience function that returns FeatureMainnetsNotAllowedForPlan wrapped in GetXRPRippleTransactionDetailsByTransactionIDE403
func FeatureMainnetsNotAllowedForPlanAsGetXRPRippleTransactionDetailsByTransactionIDE403(v *FeatureMainnetsNotAllowedForPlan) GetXRPRippleTransactionDetailsByTransactionIDE403 {
	return GetXRPRippleTransactionDetailsByTransactionIDE403{ FeatureMainnetsNotAllowedForPlan: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetXRPRippleTransactionDetailsByTransactionIDE403) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("Data matches more than one schema in oneOf(GetXRPRippleTransactionDetailsByTransactionIDE403)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GetXRPRippleTransactionDetailsByTransactionIDE403)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetXRPRippleTransactionDetailsByTransactionIDE403) MarshalJSON() ([]byte, error) {
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
func (obj *GetXRPRippleTransactionDetailsByTransactionIDE403) GetActualInstance() (interface{}) {
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

type NullableGetXRPRippleTransactionDetailsByTransactionIDE403 struct {
	value *GetXRPRippleTransactionDetailsByTransactionIDE403
	isSet bool
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDE403) Get() *GetXRPRippleTransactionDetailsByTransactionIDE403 {
	return v.value
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDE403) Set(val *GetXRPRippleTransactionDetailsByTransactionIDE403) {
	v.value = val
	v.isSet = true
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDE403) IsSet() bool {
	return v.isSet
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDE403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetXRPRippleTransactionDetailsByTransactionIDE403(val *GetXRPRippleTransactionDetailsByTransactionIDE403) *NullableGetXRPRippleTransactionDetailsByTransactionIDE403 {
	return &NullableGetXRPRippleTransactionDetailsByTransactionIDE403{value: val, isSet: true}
}

func (v NullableGetXRPRippleTransactionDetailsByTransactionIDE403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetXRPRippleTransactionDetailsByTransactionIDE403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


