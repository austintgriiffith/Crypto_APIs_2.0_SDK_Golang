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

// ListTransactionsByBlockHeightRIBSL Litecoin
type ListTransactionsByBlockHeightRIBSL struct {
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int32 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents transaction version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListTransactionsByBlockHeightRIBSLVin `json:"vin"`
	// Represents the transaction outputs.
	Vout []ListTransactionsByBlockHeightRIBSLVout `json:"vout"`
	// Represents the virtual size of this transaction.
	Vsize int32 `json:"vsize"`
}

// NewListTransactionsByBlockHeightRIBSL instantiates a new ListTransactionsByBlockHeightRIBSL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByBlockHeightRIBSL(locktime int32, size int32, version int32, vin []ListTransactionsByBlockHeightRIBSLVin, vout []ListTransactionsByBlockHeightRIBSLVout, vsize int32) *ListTransactionsByBlockHeightRIBSL {
	this := ListTransactionsByBlockHeightRIBSL{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	this.Vsize = vsize
	return &this
}

// NewListTransactionsByBlockHeightRIBSLWithDefaults instantiates a new ListTransactionsByBlockHeightRIBSL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByBlockHeightRIBSLWithDefaults() *ListTransactionsByBlockHeightRIBSL {
	this := ListTransactionsByBlockHeightRIBSL{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListTransactionsByBlockHeightRIBSL) GetLocktime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSL) GetLocktimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListTransactionsByBlockHeightRIBSL) SetLocktime(v int32) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListTransactionsByBlockHeightRIBSL) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSL) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListTransactionsByBlockHeightRIBSL) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *ListTransactionsByBlockHeightRIBSL) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSL) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListTransactionsByBlockHeightRIBSL) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListTransactionsByBlockHeightRIBSL) GetVin() []ListTransactionsByBlockHeightRIBSLVin {
	if o == nil {
		var ret []ListTransactionsByBlockHeightRIBSLVin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSL) GetVinOk() (*[]ListTransactionsByBlockHeightRIBSLVin, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *ListTransactionsByBlockHeightRIBSL) SetVin(v []ListTransactionsByBlockHeightRIBSLVin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByBlockHeightRIBSL) GetVout() []ListTransactionsByBlockHeightRIBSLVout {
	if o == nil {
		var ret []ListTransactionsByBlockHeightRIBSLVout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSL) GetVoutOk() (*[]ListTransactionsByBlockHeightRIBSLVout, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByBlockHeightRIBSL) SetVout(v []ListTransactionsByBlockHeightRIBSLVout) {
	o.Vout = v
}

// GetVsize returns the Vsize field value
func (o *ListTransactionsByBlockHeightRIBSL) GetVsize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vsize
}

// GetVsizeOk returns a tuple with the Vsize field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByBlockHeightRIBSL) GetVsizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vsize, true
}

// SetVsize sets field value
func (o *ListTransactionsByBlockHeightRIBSL) SetVsize(v int32) {
	o.Vsize = v
}

func (o ListTransactionsByBlockHeightRIBSL) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["vsize"] = o.Vsize
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactionsByBlockHeightRIBSL struct {
	value *ListTransactionsByBlockHeightRIBSL
	isSet bool
}

func (v NullableListTransactionsByBlockHeightRIBSL) Get() *ListTransactionsByBlockHeightRIBSL {
	return v.value
}

func (v *NullableListTransactionsByBlockHeightRIBSL) Set(val *ListTransactionsByBlockHeightRIBSL) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByBlockHeightRIBSL) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByBlockHeightRIBSL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByBlockHeightRIBSL(val *ListTransactionsByBlockHeightRIBSL) *NullableListTransactionsByBlockHeightRIBSL {
	return &NullableListTransactionsByBlockHeightRIBSL{value: val, isSet: true}
}

func (v NullableListTransactionsByBlockHeightRIBSL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByBlockHeightRIBSL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


