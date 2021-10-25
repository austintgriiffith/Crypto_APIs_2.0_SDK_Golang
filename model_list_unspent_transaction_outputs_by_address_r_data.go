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

// ListUnspentTransactionOutputsByAddressRData struct for ListUnspentTransactionOutputsByAddressRData
type ListUnspentTransactionOutputsByAddressRData struct {
	// The starting index of the response items, i.e. where the response should start listing the returned items.
	Offset int32 `json:"offset"`
	// Defines how many items should be returned in the response per page basis.
	Limit int32 `json:"limit"`
	// Defines the total number of items returned in the response.
	Total int32 `json:"total"`
	Items []ListUnspentTransactionOutputsByAddressRI `json:"items"`
}

// NewListUnspentTransactionOutputsByAddressRData instantiates a new ListUnspentTransactionOutputsByAddressRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUnspentTransactionOutputsByAddressRData(offset int32, limit int32, total int32, items []ListUnspentTransactionOutputsByAddressRI) *ListUnspentTransactionOutputsByAddressRData {
	this := ListUnspentTransactionOutputsByAddressRData{}
	this.Offset = offset
	this.Limit = limit
	this.Total = total
	this.Items = items
	return &this
}

// NewListUnspentTransactionOutputsByAddressRDataWithDefaults instantiates a new ListUnspentTransactionOutputsByAddressRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUnspentTransactionOutputsByAddressRDataWithDefaults() *ListUnspentTransactionOutputsByAddressRData {
	this := ListUnspentTransactionOutputsByAddressRData{}
	return &this
}

// GetOffset returns the Offset field value
func (o *ListUnspentTransactionOutputsByAddressRData) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ListUnspentTransactionOutputsByAddressRData) GetOffsetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ListUnspentTransactionOutputsByAddressRData) SetOffset(v int32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *ListUnspentTransactionOutputsByAddressRData) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListUnspentTransactionOutputsByAddressRData) GetLimitOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListUnspentTransactionOutputsByAddressRData) SetLimit(v int32) {
	o.Limit = v
}

// GetTotal returns the Total field value
func (o *ListUnspentTransactionOutputsByAddressRData) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListUnspentTransactionOutputsByAddressRData) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ListUnspentTransactionOutputsByAddressRData) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *ListUnspentTransactionOutputsByAddressRData) GetItems() []ListUnspentTransactionOutputsByAddressRI {
	if o == nil {
		var ret []ListUnspentTransactionOutputsByAddressRI
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListUnspentTransactionOutputsByAddressRData) GetItemsOk() (*[]ListUnspentTransactionOutputsByAddressRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *ListUnspentTransactionOutputsByAddressRData) SetItems(v []ListUnspentTransactionOutputsByAddressRI) {
	o.Items = v
}

func (o ListUnspentTransactionOutputsByAddressRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableListUnspentTransactionOutputsByAddressRData struct {
	value *ListUnspentTransactionOutputsByAddressRData
	isSet bool
}

func (v NullableListUnspentTransactionOutputsByAddressRData) Get() *ListUnspentTransactionOutputsByAddressRData {
	return v.value
}

func (v *NullableListUnspentTransactionOutputsByAddressRData) Set(val *ListUnspentTransactionOutputsByAddressRData) {
	v.value = val
	v.isSet = true
}

func (v NullableListUnspentTransactionOutputsByAddressRData) IsSet() bool {
	return v.isSet
}

func (v *NullableListUnspentTransactionOutputsByAddressRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUnspentTransactionOutputsByAddressRData(val *ListUnspentTransactionOutputsByAddressRData) *NullableListUnspentTransactionOutputsByAddressRData {
	return &NullableListUnspentTransactionOutputsByAddressRData{value: val, isSet: true}
}

func (v NullableListUnspentTransactionOutputsByAddressRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUnspentTransactionOutputsByAddressRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


