/*
 * CryptoAPIs
 *
 * Crypto APIs 2.0 is a complex and innovative infrastructure layer that radically simplifies the development of any Blockchain and Crypto related applications. Organized around REST, Crypto APIs 2.0 can assist both novice Bitcoin/Ethereum enthusiasts and crypto experts with the development of their blockchain applications. Crypto APIs 2.0 provides unified endpoints and data, raw data, automatic tokens and coins forwardings, callback functionalities, and much more.
 *
 * API version: 2.0.0
 * Contact: developers@cryptoapis.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cryptoapis

import (
	"encoding/json"
)

// CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken Bitcoin Omni Token
type CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken struct {
	// Represents the specific `propertyId` of the token data that will be forwarded.
	PropertyId int32 `json:"propertyId"`
}

// NewCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken instantiates a new CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken(propertyId int32) *CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken {
	this := CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken{}
	this.PropertyId = propertyId
	return &this
}

// NewCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniTokenWithDefaults instantiates a new CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniTokenWithDefaults() *CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken {
	this := CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken{}
	return &this
}

// GetPropertyId returns the PropertyId field value
func (o *CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) GetPropertyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) GetPropertyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value
func (o *CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) SetPropertyId(v int32) {
	o.PropertyId = v
}

func (o CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["propertyId"] = o.PropertyId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken struct {
	value *CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken
	isSet bool
}

func (v NullableCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) Get() *CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken {
	return v.value
}

func (v *NullableCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) Set(val *CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken(val *CreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) *NullableCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken {
	return &NullableCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken{value: val, isSet: true}
}

func (v NullableCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAutomaticTokensForwardingRBTokenDataBitcoinOmniToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

