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

// ListTransactionsByBlockHashRIBSD Dogecoin
type ListTransactionsByBlockHashRIBSD struct {
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int32 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents transaction version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListTransactionsByBlockHashRIBSDVin `json:"vin"`
	// Represents the transaction outputs.
	Vout []ListTransactionsByBlockHashRIBSDVout `json:"vout"`
}

// NewListTransactionsByBlockHashRIBSD instantiates a new ListTransactionsByBlockHashRIBSD object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHashRIBSD(locktime int32, size int32, version int32, vin []ListTransactionsByBlockHashRIBSDVin, vout []ListTransactionsByBlockHashRIBSDVout) *ListTransactionsByBlockHashRIBSD {
	this := ListTransactionsByBlockHashRIBSD{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListTransactionsByBlockHashRIBSDWithDefaults instantiates a new ListTransactionsByBlockHashRIBSD object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHashRIBSDWithDefaults() *ListTransactionsByBlockHashRIBSD {
	this := ListTransactionsByBlockHashRIBSD{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListTransactionsByBlockHashRIBSD) GetLocktime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD) GetLocktimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListTransactionsByBlockHashRIBSD) SetLocktime(v int32) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListTransactionsByBlockHashRIBSD) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListTransactionsByBlockHashRIBSD) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *ListTransactionsByBlockHashRIBSD) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListTransactionsByBlockHashRIBSD) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListTransactionsByBlockHashRIBSD) GetVin() []ListTransactionsByBlockHashRIBSDVin {
	if o == nil {
		var ret []ListTransactionsByBlockHashRIBSDVin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD) GetVinOk() (*[]ListTransactionsByBlockHashRIBSDVin, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *ListTransactionsByBlockHashRIBSD) SetVin(v []ListTransactionsByBlockHashRIBSDVin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByBlockHashRIBSD) GetVout() []ListTransactionsByBlockHashRIBSDVout {
	if o == nil {
		var ret []ListTransactionsByBlockHashRIBSDVout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHashRIBSD) GetVoutOk() (*[]ListTransactionsByBlockHashRIBSDVout, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByBlockHashRIBSD) SetVout(v []ListTransactionsByBlockHashRIBSDVout) {
	o.Vout = v
}

func (o ListTransactionsByBlockHashRIBSD) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByBlockHashRIBSD struct {
	value *ListTransactionsByBlockHashRIBSD
	isSet bool
}

func (v NullableListTransactionsByBlockHashRIBSD) Get() *ListTransactionsByBlockHashRIBSD {
	return v.value
}

func (v *NullableListTransactionsByBlockHashRIBSD) Set(val *ListTransactionsByBlockHashRIBSD) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHashRIBSD) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHashRIBSD) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHashRIBSD(val *ListTransactionsByBlockHashRIBSD) *NullableListTransactionsByBlockHashRIBSD {
	return &NullableListTransactionsByBlockHashRIBSD{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHashRIBSD) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHashRIBSD) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


