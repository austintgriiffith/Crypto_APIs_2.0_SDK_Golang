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

// ListTransactionsByAddressRData struct for ListTransactionsByAddressRData
type ListTransactionsByAddressRData struct {
	// The starting index of the response items, i.e. where the response should start listing the returned items.
	Offset int32 `json:"offset"`
	// Defines how many items should be returned in the response per page basis.
	Limit int32 `json:"limit"`
	// Defines the total number of items returned in the response.
	Total int32 `json:"total"`
	Items []ListTransactionsByAddressRI `json:"items"`
}

// NewListTransactionsByAddressRData instantiates a new ListTransactionsByAddressRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactionsByAddressRData(offset int32, limit int32, total int32, items []ListTransactionsByAddressRI) *ListTransactionsByAddressRData {
	this := ListTransactionsByAddressRData{}
	this.Offset = offset
	this.Limit = limit
	this.Total = total
	this.Items = items
	return &this
}

// NewListTransactionsByAddressRDataWithDefaults instantiates a new ListTransactionsByAddressRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactionsByAddressRDataWithDefaults() *ListTransactionsByAddressRData {
	this := ListTransactionsByAddressRData{}
	return &this
}

// GetOffset returns the Offset field value
func (o *ListTransactionsByAddressRData) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressRData) GetOffsetOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ListTransactionsByAddressRData) SetOffset(v int32) {
	o.Offset = v
}

// GetLimit returns the Limit field value
func (o *ListTransactionsByAddressRData) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressRData) GetLimitOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListTransactionsByAddressRData) SetLimit(v int32) {
	o.Limit = v
}

// GetTotal returns the Total field value
func (o *ListTransactionsByAddressRData) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressRData) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ListTransactionsByAddressRData) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *ListTransactionsByAddressRData) GetItems() []ListTransactionsByAddressRI {
	if o == nil {
		var ret []ListTransactionsByAddressRI
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListTransactionsByAddressRData) GetItemsOk() (*[]ListTransactionsByAddressRI, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *ListTransactionsByAddressRData) SetItems(v []ListTransactionsByAddressRI) {
	o.Items = v
}

func (o ListTransactionsByAddressRData) MarshalJSON() ([]byte, error) {
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

type NullableListTransactionsByAddressRData struct {
	value *ListTransactionsByAddressRData
	isSet bool
}

func (v NullableListTransactionsByAddressRData) Get() *ListTransactionsByAddressRData {
	return v.value
}

func (v *NullableListTransactionsByAddressRData) Set(val *ListTransactionsByAddressRData) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactionsByAddressRData) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactionsByAddressRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactionsByAddressRData(val *ListTransactionsByAddressRData) *NullableListTransactionsByAddressRData {
	return &NullableListTransactionsByAddressRData{value: val, isSet: true}
}

func (v NullableListTransactionsByAddressRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactionsByAddressRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

