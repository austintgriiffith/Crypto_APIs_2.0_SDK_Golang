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

// CreateCoinsTransactionFromAddressForWholeAmountRI struct for CreateCoinsTransactionFromAddressForWholeAmountRI
type CreateCoinsTransactionFromAddressForWholeAmountRI struct {
	// Represents the Secret Key value provided by the customer. This field is used for security purposes during the callback notification, in order to prove the sender of the callback as Crypto APIs. For more information please see our [Documentation](https://developers.cryptoapis.io/technical-documentation/general-information/callbacks#callback-security).
	CallbackSecretKey *string `json:"callbackSecretKey,omitempty"`
	// Represents the URL that is set by the customer where the callback will be received at. The callback notification will be received only if and when the event occurs.
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// Represents the fee priority of the automation, whether it is \"slow\", \"standard\" or \"fast\".
	FeePriority string `json:"feePriority"`
	// Represents an optional note to add a free text in, explaining or providing additional detail on the transaction request.Optional Transaction note with additional details
	Note *string `json:"note,omitempty"`
	// Defines the destination for the transaction, i.e. the recipient(s).
	Recipients []CreateCoinsTransactionFromAddressForWholeAmountRIRecipients `json:"recipients"`
	Senders CreateCoinsTransactionFromAddressForWholeAmountRISenders `json:"senders"`
	// Represents a unique identifier of the transaction request (the request sent to make a transaction), which helps in identifying which callback and which `referenceId` concern that specific transaction request.
	TransactionRequestId string `json:"transactionRequestId"`
	// Defines the status of the transaction, e.g. \"created, \"await_approval\", \"pending\", \"prepared\", \"signed\", \"broadcasted\", \"success\", \"failed\", \"rejected\", mined\".
	TransactionRequestStatus string `json:"transactionRequestStatus"`
}

// NewCreateCoinsTransactionFromAddressForWholeAmountRI instantiates a new CreateCoinsTransactionFromAddressForWholeAmountRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCoinsTransactionFromAddressForWholeAmountRI(feePriority string, recipients []CreateCoinsTransactionFromAddressForWholeAmountRIRecipients, senders CreateCoinsTransactionFromAddressForWholeAmountRISenders, transactionRequestId string, transactionRequestStatus string) *CreateCoinsTransactionFromAddressForWholeAmountRI {
	this := CreateCoinsTransactionFromAddressForWholeAmountRI{}
	this.FeePriority = feePriority
	this.Recipients = recipients
	this.Senders = senders
	this.TransactionRequestId = transactionRequestId
	this.TransactionRequestStatus = transactionRequestStatus
	return &this
}

// NewCreateCoinsTransactionFromAddressForWholeAmountRIWithDefaults instantiates a new CreateCoinsTransactionFromAddressForWholeAmountRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCoinsTransactionFromAddressForWholeAmountRIWithDefaults() *CreateCoinsTransactionFromAddressForWholeAmountRI {
	this := CreateCoinsTransactionFromAddressForWholeAmountRI{}
	return &this
}

// GetCallbackSecretKey returns the CallbackSecretKey field value if set, zero value otherwise.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetCallbackSecretKey() string {
	if o == nil || o.CallbackSecretKey == nil {
		var ret string
		return ret
	}
	return *o.CallbackSecretKey
}

// GetCallbackSecretKeyOk returns a tuple with the CallbackSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetCallbackSecretKeyOk() (*string, bool) {
	if o == nil || o.CallbackSecretKey == nil {
		return nil, false
	}
	return o.CallbackSecretKey, true
}

// HasCallbackSecretKey returns a boolean if a field has been set.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) HasCallbackSecretKey() bool {
	if o != nil && o.CallbackSecretKey != nil {
		return true
	}

	return false
}

// SetCallbackSecretKey gets a reference to the given string and assigns it to the CallbackSecretKey field.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) SetCallbackSecretKey(v string) {
	o.CallbackSecretKey = &v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetCallbackUrl() string {
	if o == nil || o.CallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetCallbackUrlOk() (*string, bool) {
	if o == nil || o.CallbackUrl == nil {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl != nil {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetFeePriority returns the FeePriority field value
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetFeePriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeePriority
}

// GetFeePriorityOk returns a tuple with the FeePriority field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetFeePriorityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FeePriority, true
}

// SetFeePriority sets field value
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) SetFeePriority(v string) {
	o.FeePriority = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetNote() string {
	if o == nil || o.Note == nil {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetNoteOk() (*string, bool) {
	if o == nil || o.Note == nil {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) HasNote() bool {
	if o != nil && o.Note != nil {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) SetNote(v string) {
	o.Note = &v
}

// GetRecipients returns the Recipients field value
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetRecipients() []CreateCoinsTransactionFromAddressForWholeAmountRIRecipients {
	if o == nil {
		var ret []CreateCoinsTransactionFromAddressForWholeAmountRIRecipients
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetRecipientsOk() (*[]CreateCoinsTransactionFromAddressForWholeAmountRIRecipients, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) SetRecipients(v []CreateCoinsTransactionFromAddressForWholeAmountRIRecipients) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetSenders() CreateCoinsTransactionFromAddressForWholeAmountRISenders {
	if o == nil {
		var ret CreateCoinsTransactionFromAddressForWholeAmountRISenders
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetSendersOk() (*CreateCoinsTransactionFromAddressForWholeAmountRISenders, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Senders, true
}

// SetSenders sets field value
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) SetSenders(v CreateCoinsTransactionFromAddressForWholeAmountRISenders) {
	o.Senders = v
}

// GetTransactionRequestId returns the TransactionRequestId field value
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetTransactionRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionRequestId
}

// GetTransactionRequestIdOk returns a tuple with the TransactionRequestId field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetTransactionRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionRequestId, true
}

// SetTransactionRequestId sets field value
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) SetTransactionRequestId(v string) {
	o.TransactionRequestId = v
}

// GetTransactionRequestStatus returns the TransactionRequestStatus field value
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetTransactionRequestStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionRequestStatus
}

// GetTransactionRequestStatusOk returns a tuple with the TransactionRequestStatus field value
// and a boolean to check if the value has been set.
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) GetTransactionRequestStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionRequestStatus, true
}

// SetTransactionRequestStatus sets field value
func (o *CreateCoinsTransactionFromAddressForWholeAmountRI) SetTransactionRequestStatus(v string) {
	o.TransactionRequestStatus = v
}

func (o CreateCoinsTransactionFromAddressForWholeAmountRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["senders"] = o.Senders
	}
	if true {
		toSerialize["transactionRequestId"] = o.TransactionRequestId
	}
	if true {
		toSerialize["transactionRequestStatus"] = o.TransactionRequestStatus
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCoinsTransactionFromAddressForWholeAmountRI struct {
	value *CreateCoinsTransactionFromAddressForWholeAmountRI
	isSet bool
}

func (v NullableCreateCoinsTransactionFromAddressForWholeAmountRI) Get() *CreateCoinsTransactionFromAddressForWholeAmountRI {
	return v.value
}

func (v *NullableCreateCoinsTransactionFromAddressForWholeAmountRI) Set(val *CreateCoinsTransactionFromAddressForWholeAmountRI) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCoinsTransactionFromAddressForWholeAmountRI) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCoinsTransactionFromAddressForWholeAmountRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCoinsTransactionFromAddressForWholeAmountRI(val *CreateCoinsTransactionFromAddressForWholeAmountRI) *NullableCreateCoinsTransactionFromAddressForWholeAmountRI {
	return &NullableCreateCoinsTransactionFromAddressForWholeAmountRI{value: val, isSet: true}
}

func (v NullableCreateCoinsTransactionFromAddressForWholeAmountRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCoinsTransactionFromAddressForWholeAmountRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


