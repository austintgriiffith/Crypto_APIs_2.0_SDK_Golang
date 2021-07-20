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

// DeleteAutomaticTokensForwardingRITSBOT Bitcoin Omni Token
type DeleteAutomaticTokensForwardingRITSBOT struct {
	// Defines the `propertyId` of the Omni Layer token.
	PropertyId int32 `json:"propertyId"`
}

// NewDeleteAutomaticTokensForwardingRITSBOT instantiates a new DeleteAutomaticTokensForwardingRITSBOT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAutomaticTokensForwardingRITSBOT(propertyId int32) *DeleteAutomaticTokensForwardingRITSBOT {
	this := DeleteAutomaticTokensForwardingRITSBOT{}
	this.PropertyId = propertyId
	return &this
}

// NewDeleteAutomaticTokensForwardingRITSBOTWithDefaults instantiates a new DeleteAutomaticTokensForwardingRITSBOT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAutomaticTokensForwardingRITSBOTWithDefaults() *DeleteAutomaticTokensForwardingRITSBOT {
	this := DeleteAutomaticTokensForwardingRITSBOT{}
	return &this
}

// GetPropertyId returns the PropertyId field value
func (o *DeleteAutomaticTokensForwardingRITSBOT) GetPropertyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *DeleteAutomaticTokensForwardingRITSBOT) GetPropertyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value
func (o *DeleteAutomaticTokensForwardingRITSBOT) SetPropertyId(v int32) {
	o.PropertyId = v
}

func (o DeleteAutomaticTokensForwardingRITSBOT) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["propertyId"] = o.PropertyId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteAutomaticTokensForwardingRITSBOT struct {
	value *DeleteAutomaticTokensForwardingRITSBOT
	isSet bool
}

func (v NullableDeleteAutomaticTokensForwardingRITSBOT) Get() *DeleteAutomaticTokensForwardingRITSBOT {
	return v.value
}

func (v *NullableDeleteAutomaticTokensForwardingRITSBOT) Set(val *DeleteAutomaticTokensForwardingRITSBOT) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAutomaticTokensForwardingRITSBOT) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAutomaticTokensForwardingRITSBOT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAutomaticTokensForwardingRITSBOT(val *DeleteAutomaticTokensForwardingRITSBOT) *NullableDeleteAutomaticTokensForwardingRITSBOT {
	return &NullableDeleteAutomaticTokensForwardingRITSBOT{value: val, isSet: true}
}

func (v NullableDeleteAutomaticTokensForwardingRITSBOT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAutomaticTokensForwardingRITSBOT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


