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

// GetTransactionDetailsByTransactionIDFromCallbackRIBSL Litecoin
type GetTransactionDetailsByTransactionIDFromCallbackRIBSL struct {
	// Represents the time at which a particular transaction can be added to the blockchain.
	Locktime int32 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Represents the virtual size of this transaction.
	VSize int32 `json:"vSize"`
	// Represents transaction version number.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []GetTransactionDetailsByTransactionIDRIBSLVin `json:"vin"`
	// Represents the transaction outputs.
	Vout []GetTransactionDetailsByTransactionIDRIBSLVout `json:"vout"`
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSL instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSL(locktime int32, size int32, vSize int32, version int32, vin []GetTransactionDetailsByTransactionIDRIBSLVin, vout []GetTransactionDetailsByTransactionIDRIBSLVout) *GetTransactionDetailsByTransactionIDFromCallbackRIBSL {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSL{}
	this.Locktime = locktime
	this.Size = size
	this.VSize = vSize
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewGetTransactionDetailsByTransactionIDFromCallbackRIBSLWithDefaults instantiates a new GetTransactionDetailsByTransactionIDFromCallbackRIBSL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionDetailsByTransactionIDFromCallbackRIBSLWithDefaults() *GetTransactionDetailsByTransactionIDFromCallbackRIBSL {
	this := GetTransactionDetailsByTransactionIDFromCallbackRIBSL{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetLocktime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetLocktimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) SetLocktime(v int32) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) SetSize(v int32) {
	o.Size = v
}

// GetVSize returns the VSize field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetVSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VSize
}

// GetVSizeOk returns a tuple with the VSize field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetVSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VSize, true
}

// SetVSize sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) SetVSize(v int32) {
	o.VSize = v
}

// GetVersion returns the Version field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetVin() []GetTransactionDetailsByTransactionIDRIBSLVin {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSLVin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetVinOk() (*[]GetTransactionDetailsByTransactionIDRIBSLVin, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) SetVin(v []GetTransactionDetailsByTransactionIDRIBSLVin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetVout() []GetTransactionDetailsByTransactionIDRIBSLVout {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSLVout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSLVout, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) SetVout(v []GetTransactionDetailsByTransactionIDRIBSLVout) {
	o.Vout = v
}

func (o GetTransactionDetailsByTransactionIDFromCallbackRIBSL) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["locktime"] = o.Locktime
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["vSize"] = o.VSize
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

type NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSL struct {
	value *GetTransactionDetailsByTransactionIDFromCallbackRIBSL
	isSet bool
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSL) Get() *GetTransactionDetailsByTransactionIDFromCallbackRIBSL {
	return v.value
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSL) Set(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSL) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionDetailsByTransactionIDFromCallbackRIBSL(val *GetTransactionDetailsByTransactionIDFromCallbackRIBSL) *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSL {
	return &NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSL{value: val, isSet: true}
}

func (v NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionDetailsByTransactionIDFromCallbackRIBSL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


