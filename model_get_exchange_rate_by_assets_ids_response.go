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

// GetExchangeRateByAssetsIDsResponse struct for GetExchangeRateByAssetsIDsResponse
type GetExchangeRateByAssetsIDsResponse struct {
	// Specifies the version of the API that incorporates this endpoint.
	ApiVersion string `json:"apiVersion"`
	// Defines the ID of the request. The `requestId` is generated by Crypto APIs and it's unique for every request.
	RequestId string `json:"requestId"`
	// In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. `context` is specified by the user.
	Context *string `json:"context,omitempty"`
	Data GetExchangeRateByAssetsIDsResponseData `json:"data"`
}

// NewGetExchangeRateByAssetsIDsResponse instantiates a new GetExchangeRateByAssetsIDsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeRateByAssetsIDsResponse(apiVersion string, requestId string, data GetExchangeRateByAssetsIDsResponseData) *GetExchangeRateByAssetsIDsResponse {
	this := GetExchangeRateByAssetsIDsResponse{}
	this.ApiVersion = apiVersion
	this.RequestId = requestId
	this.Data = data
	return &this
}

// NewGetExchangeRateByAssetsIDsResponseWithDefaults instantiates a new GetExchangeRateByAssetsIDsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeRateByAssetsIDsResponseWithDefaults() *GetExchangeRateByAssetsIDsResponse {
	this := GetExchangeRateByAssetsIDsResponse{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *GetExchangeRateByAssetsIDsResponse) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *GetExchangeRateByAssetsIDsResponse) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *GetExchangeRateByAssetsIDsResponse) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetRequestId returns the RequestId field value
func (o *GetExchangeRateByAssetsIDsResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *GetExchangeRateByAssetsIDsResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *GetExchangeRateByAssetsIDsResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *GetExchangeRateByAssetsIDsResponse) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeRateByAssetsIDsResponse) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *GetExchangeRateByAssetsIDsResponse) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *GetExchangeRateByAssetsIDsResponse) SetContext(v string) {
	o.Context = &v
}

// GetData returns the Data field value
func (o *GetExchangeRateByAssetsIDsResponse) GetData() GetExchangeRateByAssetsIDsResponseData {
	if o == nil {
		var ret GetExchangeRateByAssetsIDsResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetExchangeRateByAssetsIDsResponse) GetDataOk() (*GetExchangeRateByAssetsIDsResponseData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetExchangeRateByAssetsIDsResponse) SetData(v GetExchangeRateByAssetsIDsResponseData) {
	o.Data = v
}

func (o GetExchangeRateByAssetsIDsResponse) MarshalJSON() ([]byte, error) {
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

type NullableGetExchangeRateByAssetsIDsResponse struct {
	value *GetExchangeRateByAssetsIDsResponse
	isSet bool
}

func (v NullableGetExchangeRateByAssetsIDsResponse) Get() *GetExchangeRateByAssetsIDsResponse {
	return v.value
}

func (v *NullableGetExchangeRateByAssetsIDsResponse) Set(val *GetExchangeRateByAssetsIDsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeRateByAssetsIDsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeRateByAssetsIDsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeRateByAssetsIDsResponse(val *GetExchangeRateByAssetsIDsResponse) *NullableGetExchangeRateByAssetsIDsResponse {
	return &NullableGetExchangeRateByAssetsIDsResponse{value: val, isSet: true}
}

func (v NullableGetExchangeRateByAssetsIDsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeRateByAssetsIDsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

