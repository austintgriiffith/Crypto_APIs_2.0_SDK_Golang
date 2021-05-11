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

// ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin Bitcoin
type ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin struct {
	// Represents the locktime on the transaction on the specific blockchain, i.e. the blockheight at which the transaction is valid.
	Locktime int32 `json:"locktime"`
	// Represents the total size of this transaction.
	Size int32 `json:"size"`
	// Defines the transaction's virtual size.
	VSize int32 `json:"vSize"`
	// Defines the version of the transaction.
	Version int32 `json:"version"`
	// Represents the transaction inputs.
	Vin []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin `json:"vin"`
	// Represents the transaction outputs.
	Vout []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout `json:"vout"`
}

// NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoin instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoin(locktime int32, size int32, vSize int32, version int32, vin []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin, vout []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout) *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin {
	this := ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin{}
	this.Locktime = locktime
	this.Size = size
	this.VSize = vSize
	this.Version = version
	this.Vin = vin
	this.Vout = vout
	return &this
}

// NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinWithDefaults instantiates a new ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByAddressResponseItemBlockchainSpecificBitcoinWithDefaults() *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin {
	this := ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin{}
	return &this
}

// GetLocktime returns the Locktime field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetLocktime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Locktime
}

// GetLocktimeOk returns a tuple with the Locktime field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetLocktimeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Locktime, true
}

// SetLocktime sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetLocktime(v int32) {
	o.Locktime = v
}

// GetSize returns the Size field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetSize(v int32) {
	o.Size = v
}

// GetVSize returns the VSize field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VSize
}

// GetVSizeOk returns a tuple with the VSize field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VSize, true
}

// SetVSize sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetVSize(v int32) {
	o.VSize = v
}

// GetVersion returns the Version field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetVersion(v int32) {
	o.Version = v
}

// GetVin returns the Vin field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVin() []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin {
	if o == nil {
		var ret []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin
		return ret
	}

	return o.Vin
}

// GetVinOk returns a tuple with the Vin field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVinOk() (*[]ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vin, true
}

// SetVin sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetVin(v []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVin) {
	o.Vin = v
}

// GetVout returns the Vout field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVout() []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout {
	if o == nil {
		var ret []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) GetVoutOk() (*[]ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) SetVout(v []ListTransactionsByAddressResponseItemBlockchainSpecificBitcoinVout) {
	o.Vout = v
}

func (o ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoin struct {
	value *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin
	isSet bool
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) Get() *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin {
	return v.value
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) Set(val *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoin(val *ListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) *NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoin {
	return &NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoin{value: val, isSet: true}
}

func (v NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByAddressResponseItemBlockchainSpecificBitcoin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

