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

// ListUnconfirmedTransactionsByAddressRIBSD Dogecoin
type ListUnconfirmedTransactionsByAddressRIBSD struct {
	// Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid.
	Locktime int32 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Numeric representation of the transaction version
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListUnconfirmedTransactionsByAddressRIBSDVin `json:"vin"`
	// Represents the transaction outputs.
	Vout []GetTransactionDetailsByTransactionIDRIBSDVout `json:"vout"`
}

// NewListUnconfirmedTransactionsByAddressRIBSD instantiates a new ListUnconfirmedTransactionsByAddressRIBSD object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnconfirmedTransactionsByAddressRIBSD(locktime int32, size int32, version int32, vin []ListUnconfirmedTransactionsByAddressRIBSDVin, vout []GetTransactionDetailsByTransactionIDRIBSDVout) *ListUnconfirmedTransactionsByAddressRIBSD {
	this := ListUnconfirmedTransactionsByAddressRIBSD{}
	this.Locktime = locktime
	this.Size = size
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListUnconfirmedTransactionsByAddressRIBSDWithDefaults instantiates a new ListUnconfirmedTransactionsByAddressRIBSD object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnconfirmedTransactionsByAddressRIBSDWithDefaults() *ListUnconfirmedTransactionsByAddressRIBSD {
	this := ListUnconfirmedTransactionsByAddressRIBSD{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetLocktime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetLocktimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD) SetLocktime(v int32) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD) SetSize(v int32) {
	o.Size = v
}

// GetVersion returns the Version field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVin() []ListUnconfirmedTransactionsByAddressRIBSDVin {
	if o == nil {
		var ret []ListUnconfirmedTransactionsByAddressRIBSDVin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVinOk() (*[]ListUnconfirmedTransactionsByAddressRIBSDVin, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD) SetVin(v []ListUnconfirmedTransactionsByAddressRIBSDVin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVout() []GetTransactionDetailsByTransactionIDRIBSDVout {
	if o == nil {
		var ret []GetTransactionDetailsByTransactionIDRIBSDVout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListUnconfirmedTransactionsByAddressRIBSD) GetVoutOk() (*[]GetTransactionDetailsByTransactionIDRIBSDVout, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListUnconfirmedTransactionsByAddressRIBSD) SetVout(v []GetTransactionDetailsByTransactionIDRIBSDVout) {
	o.Vout = v
}

func (o ListUnconfirmedTransactionsByAddressRIBSD) MarshalJSON() ([]byte, error) {
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

type NullableListUnconfirmedTransactionsByAddressRIBSD struct {
	value *ListUnconfirmedTransactionsByAddressRIBSD
	isSet bool
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSD) Get() *ListUnconfirmedTransactionsByAddressRIBSD {
	return v.value
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSD) Set(val *ListUnconfirmedTransactionsByAddressRIBSD) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSD) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSD) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnconfirmedTransactionsByAddressRIBSD(val *ListUnconfirmedTransactionsByAddressRIBSD) *NullableListUnconfirmedTransactionsByAddressRIBSD {
	return &NullableListUnconfirmedTransactionsByAddressRIBSD{value: val, isSet: true}
}

func (v NullableListUnconfirmedTransactionsByAddressRIBSD) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnconfirmedTransactionsByAddressRIBSD) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


