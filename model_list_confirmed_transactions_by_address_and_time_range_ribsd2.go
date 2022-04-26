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
)

// ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 Dash
type ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 struct {
	// Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid.
	Locktime int64 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents the transaction's version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListConfirmedTransactionsByAddressRIBSD2Vin `json:"vin"`
	// Represents the transaction outputs.
	Vout []ListConfirmedTransactionsByAddressRIBSD2Vout `json:"vout"`
}

// NewListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSD2(locktime int64, size int32, version int32, vin []ListConfirmedTransactionsByAddressRIBSD2Vin, vout []ListConfirmedTransactionsByAddressRIBSD2Vout) *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 {
	this := ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListConfirmedTransactionsByAddressAndTimeRangeRIBSD2WithDefaults instantiates a new ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTransactionsByAddressAndTimeRangeRIBSD2WithDefaults() *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 {
	this := ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) GetLocktime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) GetLocktimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) SetLocktime(v int64) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) GetVin() []ListConfirmedTransactionsByAddressRIBSD2Vin {
	if o == nil {
		var ret []ListConfirmedTransactionsByAddressRIBSD2Vin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) GetVinOk() ([]ListConfirmedTransactionsByAddressRIBSD2Vin, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vin, true
}

// SetVin sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) SetVin(v []ListConfirmedTransactionsByAddressRIBSD2Vin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) GetVout() []ListConfirmedTransactionsByAddressRIBSD2Vout {
	if o == nil {
		var ret []ListConfirmedTransactionsByAddressRIBSD2Vout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) GetVoutOk() ([]ListConfirmedTransactionsByAddressRIBSD2Vout, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vout, true
}

// SetVout sets field value
func (o *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) SetVout(v []ListConfirmedTransactionsByAddressRIBSD2Vout) {
	o.Vout = v
}

func (o ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locktime"] = o.Locktime
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["vin"] = o.Vin
	}
	if true {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 struct {
	value *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2
	isSet bool
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) Get() *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 {
	return v.value
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) Set(val *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTransactionsByAddressAndTimeRangeRIBSD2(val *ListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSD2 {
	return &NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSD2{value: val, isSet: true}
}

func (v NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTransactionsByAddressAndTimeRangeRIBSD2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

