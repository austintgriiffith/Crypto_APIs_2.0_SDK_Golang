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

// GetZilliqaTransactionDetailsByTransactionIDRData struct for GetZilliqaTransactionDetailsByTransactionIDRData
type GetZilliqaTransactionDetailsByTransactionIDRData struct {
	Item GetZilliqaTransactionDetailsByTransactionIDRI `json:"item"`
}

// NewGetZilliqaTransactionDetailsByTransactionIDRData instantiates a new GetZilliqaTransactionDetailsByTransactionIDRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetZilliqaTransactionDetailsByTransactionIDRData(item GetZilliqaTransactionDetailsByTransactionIDRI) *GetZilliqaTransactionDetailsByTransactionIDRData {
	this := GetZilliqaTransactionDetailsByTransactionIDRData{}
	this.Item = item
	return &this
}

// NewGetZilliqaTransactionDetailsByTransactionIDRDataWithDefaults instantiates a new GetZilliqaTransactionDetailsByTransactionIDRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetZilliqaTransactionDetailsByTransactionIDRDataWithDefaults() *GetZilliqaTransactionDetailsByTransactionIDRData {
	this := GetZilliqaTransactionDetailsByTransactionIDRData{}
	return &this
}

// GetItem returns the Item field value
func (o *GetZilliqaTransactionDetailsByTransactionIDRData) GetItem() GetZilliqaTransactionDetailsByTransactionIDRI {
	if o == nil {
		var ret GetZilliqaTransactionDetailsByTransactionIDRI
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *GetZilliqaTransactionDetailsByTransactionIDRData) GetItemOk() (*GetZilliqaTransactionDetailsByTransactionIDRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *GetZilliqaTransactionDetailsByTransactionIDRData) SetItem(v GetZilliqaTransactionDetailsByTransactionIDRI) {
	o.Item = v
}

func (o GetZilliqaTransactionDetailsByTransactionIDRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item"] = o.Item
	}
	return json.Marshal(toSerialize)
}

type NullableGetZilliqaTransactionDetailsByTransactionIDRData struct {
	value *GetZilliqaTransactionDetailsByTransactionIDRData
	isSet bool
}

func (v NullableGetZilliqaTransactionDetailsByTransactionIDRData) Get() *GetZilliqaTransactionDetailsByTransactionIDRData {
	return v.value
}

func (v *NullableGetZilliqaTransactionDetailsByTransactionIDRData) Set(val *GetZilliqaTransactionDetailsByTransactionIDRData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetZilliqaTransactionDetailsByTransactionIDRData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetZilliqaTransactionDetailsByTransactionIDRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetZilliqaTransactionDetailsByTransactionIDRData(val *GetZilliqaTransactionDetailsByTransactionIDRData) *NullableGetZilliqaTransactionDetailsByTransactionIDRData {
	return &NullableGetZilliqaTransactionDetailsByTransactionIDRData{value: val, isSet: true}
}

func (v NullableGetZilliqaTransactionDetailsByTransactionIDRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetZilliqaTransactionDetailsByTransactionIDRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


