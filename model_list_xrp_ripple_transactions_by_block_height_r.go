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

// ListXRPRippleTransactionsByBlockHeightR struct for ListXRPRippleTransactionsByBlockHeightR
type ListXRPRippleTransactionsByBlockHeightR struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data ListXRPRippleTransactionsByBlockHeightRData `json:"data"`
}

// NewListXRPRippleTransactionsByBlockHeightR instantiates a new ListXRPRippleTransactionsByBlockHeightR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListXRPRippleTransactionsByBlockHeightR(apiVersion string, requestId string, data ListXRPRippleTransactionsByBlockHeightRData) *ListXRPRippleTransactionsByBlockHeightR {
	this := ListXRPRippleTransactionsByBlockHeightR{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Data = data
	return &this
}

// NewListXRPRippleTransactionsByBlockHeightRWithDefaults instantiates a new ListXRPRippleTransactionsByBlockHeightR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListXRPRippleTransactionsByBlockHeightRWithDefaults() *ListXRPRippleTransactionsByBlockHeightR {
	this := ListXRPRippleTransactionsByBlockHeightR{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ListXRPRippleTransactionsByBlockHeightR) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightR) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ListXRPRippleTransactionsByBlockHeightR) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *ListXRPRippleTransactionsByBlockHeightR) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightR) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ListXRPRippleTransactionsByBlockHeightR) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ListXRPRippleTransactionsByBlockHeightR) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightR) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ListXRPRippleTransactionsByBlockHeightR) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *ListXRPRippleTransactionsByBlockHeightR) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *ListXRPRippleTransactionsByBlockHeightR) GetData() ListXRPRippleTransactionsByBlockHeightRData {
	if o == nil {
		var ret ListXRPRippleTransactionsByBlockHeightRData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListXRPRippleTransactionsByBlockHeightR) GetDataOk() (*ListXRPRippleTransactionsByBlockHeightRData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListXRPRippleTransactionsByBlockHeightR) SetData(v ListXRPRippleTransactionsByBlockHeightRData) {
	o.Data = v
}

func (o ListXRPRippleTransactionsByBlockHeightR) MarshalJSON() ([]byte, error) {
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

type NullableListXRPRippleTransactionsByBlockHeightR struct {
	value *ListXRPRippleTransactionsByBlockHeightR
	isSet bool
}

func (v NullableListXRPRippleTransactionsByBlockHeightR) Get() *ListXRPRippleTransactionsByBlockHeightR {
	return v.value
}

func (v *NullableListXRPRippleTransactionsByBlockHeightR) Set(val *ListXRPRippleTransactionsByBlockHeightR) {
	v.value = val
	v.isSet = true
}

func (v NullableListXRPRippleTransactionsByBlockHeightR) IsSet() bool {
	return v.isSet
}

func (v *NullableListXRPRippleTransactionsByBlockHeightR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListXRPRippleTransactionsByBlockHeightR(val *ListXRPRippleTransactionsByBlockHeightR) *NullableListXRPRippleTransactionsByBlockHeightR {
	return &NullableListXRPRippleTransactionsByBlockHeightR{value: val, isSet: true}
}

func (v NullableListXRPRippleTransactionsByBlockHeightR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListXRPRippleTransactionsByBlockHeightR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


