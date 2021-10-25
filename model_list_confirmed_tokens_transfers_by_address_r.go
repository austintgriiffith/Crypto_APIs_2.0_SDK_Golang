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

// ListConfirmedTokensTransfersByAddressR struct for ListConfirmedTokensTransfersByAddressR
type ListConfirmedTokensTransfersByAddressR struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data ListConfirmedTokensTransfersByAddressRData `json:"data"`
}

// NewListConfirmedTokensTransfersByAddressR instantiates a new ListConfirmedTokensTransfersByAddressR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConfirmedTokensTransfersByAddressR(apiVersion string, requestId string, data ListConfirmedTokensTransfersByAddressRData) *ListConfirmedTokensTransfersByAddressR {
	this := ListConfirmedTokensTransfersByAddressR{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Data = data
	return &this
}

// NewListConfirmedTokensTransfersByAddressRWithDefaults instantiates a new ListConfirmedTokensTransfersByAddressR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConfirmedTokensTransfersByAddressRWithDefaults() *ListConfirmedTokensTransfersByAddressR {
	this := ListConfirmedTokensTransfersByAddressR{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ListConfirmedTokensTransfersByAddressR) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressR) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ListConfirmedTokensTransfersByAddressR) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *ListConfirmedTokensTransfersByAddressR) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressR) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ListConfirmedTokensTransfersByAddressR) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ListConfirmedTokensTransfersByAddressR) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressR) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ListConfirmedTokensTransfersByAddressR) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *ListConfirmedTokensTransfersByAddressR) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *ListConfirmedTokensTransfersByAddressR) GetData() ListConfirmedTokensTransfersByAddressRData {
	if o == nil {
		var ret ListConfirmedTokensTransfersByAddressRData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListConfirmedTokensTransfersByAddressR) GetDataOk() (*ListConfirmedTokensTransfersByAddressRData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListConfirmedTokensTransfersByAddressR) SetData(v ListConfirmedTokensTransfersByAddressRData) {
	o.Data = v
}

func (o ListConfirmedTokensTransfersByAddressR) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if true {
		toSerialize["requestId"] = o.RequestId
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListConfirmedTokensTransfersByAddressR struct {
	value *ListConfirmedTokensTransfersByAddressR
	isSet bool
}

func (v NullableListConfirmedTokensTransfersByAddressR) Get() *ListConfirmedTokensTransfersByAddressR {
	return v.value
}

func (v *NullableListConfirmedTokensTransfersByAddressR) Set(val *ListConfirmedTokensTransfersByAddressR) {
	v.value = val
	v.isSet = true
}

func (v NullableListConfirmedTokensTransfersByAddressR) IsSet() bool {
	return v.isSet
}

func (v *NullableListConfirmedTokensTransfersByAddressR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConfirmedTokensTransfersByAddressR(val *ListConfirmedTokensTransfersByAddressR) *NullableListConfirmedTokensTransfersByAddressR {
	return &NullableListConfirmedTokensTransfersByAddressR{value: val, isSet: true}
}

func (v NullableListConfirmedTokensTransfersByAddressR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConfirmedTokensTransfersByAddressR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


