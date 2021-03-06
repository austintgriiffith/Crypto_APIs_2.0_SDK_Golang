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

// ListOmniTransactionsByAddressRData struct for ListOmniTransactionsByAddressRData
type ListOmniTransactionsByAddressRData struct {
	// Defines how many items should be returned in the response per page basis.
	Limit int32 `json:"limit"`
	// The starting index of the response items, i.e. where the response should start listing the returned items.
	Offset int32 `json:"offset"`
	// Defines the total number of items returned in the response.
	Total int32 `json:"total"`
	Items []ListOmniTransactionsByAddressRI `json:"items"`
}

// NewListOmniTransactionsByAddressRData instantiates a new ListOmniTransactionsByAddressRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOmniTransactionsByAddressRData(limit int32, offset int32, total int32, items []ListOmniTransactionsByAddressRI) *ListOmniTransactionsByAddressRData {
	this := ListOmniTransactionsByAddressRData{}
	this.Limit = limit
	this.Offset = offset
	this.Total = total
	this.Items = items
	return &this
}

// NewListOmniTransactionsByAddressRDataWithDefaults instantiates a new ListOmniTransactionsByAddressRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOmniTransactionsByAddressRDataWithDefaults() *ListOmniTransactionsByAddressRData {
	this := ListOmniTransactionsByAddressRData{}
	return &this
}

// GetLimit returns the Limit field value
func (o *ListOmniTransactionsByAddressRData) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByAddressRData) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *ListOmniTransactionsByAddressRData) SetLimit(v int32) {
	o.Limit = v
}

// GetOffset returns the Offset field value
func (o *ListOmniTransactionsByAddressRData) GetOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByAddressRData) GetOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Offset, true
}

// SetOffset sets field value
func (o *ListOmniTransactionsByAddressRData) SetOffset(v int32) {
	o.Offset = v
}

// GetTotal returns the Total field value
func (o *ListOmniTransactionsByAddressRData) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByAddressRData) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ListOmniTransactionsByAddressRData) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *ListOmniTransactionsByAddressRData) GetItems() []ListOmniTransactionsByAddressRI {
	if o == nil {
		var ret []ListOmniTransactionsByAddressRI
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ListOmniTransactionsByAddressRData) GetItemsOk() ([]ListOmniTransactionsByAddressRI, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ListOmniTransactionsByAddressRData) SetItems(v []ListOmniTransactionsByAddressRI) {
	o.Items = v
}

func (o ListOmniTransactionsByAddressRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["offset"] = o.Offset
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableListOmniTransactionsByAddressRData struct {
	value *ListOmniTransactionsByAddressRData
	isSet bool
}

func (v NullableListOmniTransactionsByAddressRData) Get() *ListOmniTransactionsByAddressRData {
	return v.value
}

func (v *NullableListOmniTransactionsByAddressRData) Set(val *ListOmniTransactionsByAddressRData) {
	v.value = val
	v.isSet = true
}

func (v NullableListOmniTransactionsByAddressRData) IsSet() bool {
	return v.isSet
}

func (v *NullableListOmniTransactionsByAddressRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOmniTransactionsByAddressRData(val *ListOmniTransactionsByAddressRData) *NullableListOmniTransactionsByAddressRData {
	return &NullableListOmniTransactionsByAddressRData{value: val, isSet: true}
}

func (v NullableListOmniTransactionsByAddressRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOmniTransactionsByAddressRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


