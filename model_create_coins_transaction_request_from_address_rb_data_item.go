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

// CreateCoinsTransactionRequestFromAddressRBDataItem struct for CreateCoinsTransactionRequestFromAddressRBDataItem
type CreateCoinsTransactionRequestFromAddressRBDataItem struct {
	// Represents the specific amount of the transaction.
	Amount string `json:"amount"`
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security).
	CallbackSecretKey *string `json:"callbackSecretKey,omitempty"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs.
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// Represents the fee priority of the automation, whether it is \"slow\", \"standard\" or \"fast\".
	FeePriority string `json:"feePriority"`
	// Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request.
	Note *string `json:"note,omitempty"`
	// Defines the specific recipient address for the transaction.
	RecipientAddress string `json:"recipientAddress"`
}

// NewCreateCoinsTransactionRequestFromAddressRBDataItem instantiates a new CreateCoinsTransactionRequestFromAddressRBDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCoinsTransactionRequestFromAddressRBDataItem(amount string, feePriority string, recipientAddress string) *CreateCoinsTransactionRequestFromAddressRBDataItem {
	this := CreateCoinsTransactionRequestFromAddressRBDataItem{}
	this.Amount = amount
	this.FeePriority = feePriority
	this.RecipientAddress = recipientAddress
	return &this
}

// NewCreateCoinsTransactionRequestFromAddressRBDataItemWithDefaults instantiates a new CreateCoinsTransactionRequestFromAddressRBDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCoinsTransactionRequestFromAddressRBDataItemWithDefaults() *CreateCoinsTransactionRequestFromAddressRBDataItem {
	this := CreateCoinsTransactionRequestFromAddressRBDataItem{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetAmount(v string) {
	o.Amount = v
}

// GetCallbackSecretKey returns the CallbackSecretKey field value if set, zero value otherwise.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetCallbackSecretKey() string {
	if o == nil || o.CallbackSecretKey == nil {
		var ret string
		return ret
	}
	return *o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil || o.CallbackSecretKey == nil {
		return nil, false
	}
	return o.CallbackSecretKey, true
}

// HasCallbackSecretKey returns a boolean if a field has been set.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) HasCallbackSecretKey() bool {
	if o != nil && o.CallbackSecretKey != nil {
		return true
	}

	return false
}

// SetCallbackSecretKey gets a reference to the given string and assigns it to the CallbackSecretKey field.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = &v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetCallbackUrl() string {
	if o == nil || o.CallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetCallbackUrlOk() (*string, bool) {
	if o == nil || o.CallbackUrl == nil {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl != nil {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetFeePriority returns the FeePriority field value
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetFeePriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeePriority
}

// GetFeePriorityOk returns a tuple with the FeePriority field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetFeePriorityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FeePriority, true
}

// SetFeePriority sets field value
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetFeePriority(v string) {
	o.FeePriority = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetNote(v string) {
	o.Note = &v
}

// GetRecipientAddress returns the RecipientAddress field value
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetRecipientAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientAddress
}

// GetRecipientAddressOk returns a tuple with the RecipientAddress field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) GetRecipientAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientAddress, true
}

// SetRecipientAddress sets field value
func (o *CreateCoinsTransactionRequestFromAddressRBDataItem) SetRecipientAddress(v string) {
	o.RecipientAddress = v
}

func (o CreateCoinsTransactionRequestFromAddressRBDataItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.CallbackSecretKey != nil {
		toSerialize["callbackSecretKey"] = o.CallbackSecretKey
	}
	if o.CallbackUrl != nil {
		toSerialize["callbackUrl"] = o.CallbackUrl
	}
	if true {
		toSerialize["feePriority"] = o.FeePriority
	}
	if o.Note != nil {
		toSerialize["note"] = o.Note
	}
	if true {
		toSerialize["recipientAddress"] = o.RecipientAddress
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCoinsTransactionRequestFromAddressRBDataItem struct {
	value *CreateCoinsTransactionRequestFromAddressRBDataItem
	isSet bool
}

func (v NullableCreateCoinsTransactionRequestFromAddressRBDataItem) Get() *CreateCoinsTransactionRequestFromAddressRBDataItem {
	return v.value
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRBDataItem) Set(val *CreateCoinsTransactionRequestFromAddressRBDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCoinsTransactionRequestFromAddressRBDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRBDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCoinsTransactionRequestFromAddressRBDataItem(val *CreateCoinsTransactionRequestFromAddressRBDataItem) *NullableCreateCoinsTransactionRequestFromAddressRBDataItem {
	return &NullableCreateCoinsTransactionRequestFromAddressRBDataItem{value: val, isSet: true}
}

func (v NullableCreateCoinsTransactionRequestFromAddressRBDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCoinsTransactionRequestFromAddressRBDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


