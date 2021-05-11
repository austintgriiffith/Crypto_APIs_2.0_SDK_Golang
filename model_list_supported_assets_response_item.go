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

// ListSupportedAssetsResponseItem struct for ListSupportedAssetsResponseItem
type ListSupportedAssetsResponseItem struct {
	// Defines the unique ID of the specific asset.
	AssetId string `json:"assetId"`
	// Specifies the name of the asset in question.
	AssetName string `json:"assetName"`
	// Specifies the asset's unique symbol in the Crypto APIs listings.
	AssetSymbol string `json:"assetSymbol"`
	// Defines the type of the supported asset. This could be either \"crypto\" or \"fiat\".
	AssetType string `json:"assetType"`
	// Specifies the asset's original symbol as introduced by its founders.
	OriginalSymbol string `json:"originalSymbol"`
}

// NewListSupportedAssetsResponseItem instantiates a new ListSupportedAssetsResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSupportedAssetsResponseItem(assetId string, assetName string, assetSymbol string, assetType string, originalSymbol string) *ListSupportedAssetsResponseItem {
	this := ListSupportedAssetsResponseItem{}
	this.AssetId = assetId
	this.AssetName = assetName
	this.AssetSymbol = assetSymbol
	this.AssetType = assetType
	this.OriginalSymbol = originalSymbol
	return &this
}

// NewListSupportedAssetsResponseItemWithDefaults instantiates a new ListSupportedAssetsResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSupportedAssetsResponseItemWithDefaults() *ListSupportedAssetsResponseItem {
	this := ListSupportedAssetsResponseItem{}
	return &this
}

// GetAssetId returns the AssetId field value
func (o *ListSupportedAssetsResponseItem) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *ListSupportedAssetsResponseItem) GetAssetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *ListSupportedAssetsResponseItem) SetAssetId(v string) {
	o.AssetId = v
}

// GetAssetName returns the AssetName field value
func (o *ListSupportedAssetsResponseItem) GetAssetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetName
}

// GetAssetNameOk returns a tuple with the AssetName field value
// and a boolean to check if the value has been set.
func (o *ListSupportedAssetsResponseItem) GetAssetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetName, true
}

// SetAssetName sets field value
func (o *ListSupportedAssetsResponseItem) SetAssetName(v string) {
	o.AssetName = v
}

// GetAssetSymbol returns the AssetSymbol field value
func (o *ListSupportedAssetsResponseItem) GetAssetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetSymbol
}

// GetAssetSymbolOk returns a tuple with the AssetSymbol field value
// and a boolean to check if the value has been set.
func (o *ListSupportedAssetsResponseItem) GetAssetSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetSymbol, true
}

// SetAssetSymbol sets field value
func (o *ListSupportedAssetsResponseItem) SetAssetSymbol(v string) {
	o.AssetSymbol = v
}

// GetAssetType returns the AssetType field value
func (o *ListSupportedAssetsResponseItem) GetAssetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value
// and a boolean to check if the value has been set.
func (o *ListSupportedAssetsResponseItem) GetAssetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AssetType, true
}

// SetAssetType sets field value
func (o *ListSupportedAssetsResponseItem) SetAssetType(v string) {
	o.AssetType = v
}

// GetOriginalSymbol returns the OriginalSymbol field value
func (o *ListSupportedAssetsResponseItem) GetOriginalSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalSymbol
}

// GetOriginalSymbolOk returns a tuple with the OriginalSymbol field value
// and a boolean to check if the value has been set.
func (o *ListSupportedAssetsResponseItem) GetOriginalSymbolOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OriginalSymbol, true
}

// SetOriginalSymbol sets field value
func (o *ListSupportedAssetsResponseItem) SetOriginalSymbol(v string) {
	o.OriginalSymbol = v
}

func (o ListSupportedAssetsResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["assetId"] = o.AssetId
	}
	if true {
		toSerialize["assetName"] = o.AssetName
	}
	if true {
		toSerialize["assetSymbol"] = o.AssetSymbol
	}
	if true {
		toSerialize["assetType"] = o.AssetType
	}
	if true {
		toSerialize["originalSymbol"] = o.OriginalSymbol
	}
	return json.Marshal(toSerialize)
}

type NullableListSupportedAssetsResponseItem struct {
	value *ListSupportedAssetsResponseItem
	isSet bool
}

func (v NullableListSupportedAssetsResponseItem) Get() *ListSupportedAssetsResponseItem {
	return v.value
}

func (v *NullableListSupportedAssetsResponseItem) Set(val *ListSupportedAssetsResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableListSupportedAssetsResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableListSupportedAssetsResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSupportedAssetsResponseItem(val *ListSupportedAssetsResponseItem) *NullableListSupportedAssetsResponseItem {
	return &NullableListSupportedAssetsResponseItem{value: val, isSet: true}
}

func (v NullableListSupportedAssetsResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSupportedAssetsResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

