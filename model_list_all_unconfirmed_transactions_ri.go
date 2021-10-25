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

// ListAllUnconfirmedTransactionsRI struct for ListAllUnconfirmedTransactionsRI
type ListAllUnconfirmedTransactionsRI struct {
	// Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Recipients []ListUnconfirmedTransactionsByAddressRIRecipients `json:"recipients"`
	// Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list.
	Senders []ListUnconfirmedTransactionsByAddressRISenders `json:"senders"`
	// Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed.
	Timestamp int32 `json:"timestamp"`
	// Represents the unique identifier of a transaction, i.e. it could be `transactionId` in UTXO-based protocols like Bitcoin, and transaction `hash` in Ethereum blockchain.
	TransactionId string `json:"transactionId"`
	BlockchainSpecific ListAllUnconfirmedTransactionsRIBS `json:"blockchainSpecific"`
}

// NewListAllUnconfirmedTransactionsRI instantiates a new ListAllUnconfirmedTransactionsRI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllUnconfirmedTransactionsRI(recipients []ListUnconfirmedTransactionsByAddressRIRecipients, senders []ListUnconfirmedTransactionsByAddressRISenders, timestamp int32, transactionId string, blockchainSpecific ListAllUnconfirmedTransactionsRIBS) *ListAllUnconfirmedTransactionsRI {
	this := ListAllUnconfirmedTransactionsRI{}
	this.Recipients = recipients
	this.Senders = senders
	this.Timestamp = timestamp
	this.TransactionId = transactionId
	this.BlockchainSpecific = blockchainSpecific
	return &this
}

// NewListAllUnconfirmedTransactionsRIWithDefaults instantiates a new ListAllUnconfirmedTransactionsRI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllUnconfirmedTransactionsRIWithDefaults() *ListAllUnconfirmedTransactionsRI {
	this := ListAllUnconfirmedTransactionsRI{}
	return &this
}

// GetRecipients returns the Recipients field value
func (o *ListAllUnconfirmedTransactionsRI) GetRecipients() []ListUnconfirmedTransactionsByAddressRIRecipients {
	if o == nil {
		var ret []ListUnconfirmedTransactionsByAddressRIRecipients
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRI) GetRecipientsOk() (*[]ListUnconfirmedTransactionsByAddressRIRecipients, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value
func (o *ListAllUnconfirmedTransactionsRI) SetRecipients(v []ListUnconfirmedTransactionsByAddressRIRecipients) {
	o.Recipients = v
}

// GetSenders returns the Senders field value
func (o *ListAllUnconfirmedTransactionsRI) GetSenders() []ListUnconfirmedTransactionsByAddressRISenders {
	if o == nil {
		var ret []ListUnconfirmedTransactionsByAddressRISenders
		return ret
	}

	return o.Senders
}

// GetSendersOk returns a tuple with the Senders field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRI) GetSendersOk() (*[]ListUnconfirmedTransactionsByAddressRISenders, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Senders, true
}

// SetSenders sets field value
func (o *ListAllUnconfirmedTransactionsRI) SetSenders(v []ListUnconfirmedTransactionsByAddressRISenders) {
	o.Senders = v
}

// GetTimestamp returns the Timestamp field value
func (o *ListAllUnconfirmedTransactionsRI) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRI) GetTimestampOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ListAllUnconfirmedTransactionsRI) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetTransactionId returns the TransactionId field value
func (o *ListAllUnconfirmedTransactionsRI) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRI) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *ListAllUnconfirmedTransactionsRI) SetTransactionId(v string) {
	o.TransactionId = v
}

// GetBlockchainSpecific returns the BlockchainSpecific field value
func (o *ListAllUnconfirmedTransactionsRI) GetBlockchainSpecific() ListAllUnconfirmedTransactionsRIBS {
	if o == nil {
		var ret ListAllUnconfirmedTransactionsRIBS
		return ret
	}

	return o.BlockchainSpecific
}

// GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field value
// and a boolean to check if the value has been set.
func (o *ListAllUnconfirmedTransactionsRI) GetBlockchainSpecificOk() (*ListAllUnconfirmedTransactionsRIBS, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BlockchainSpecific, true
}

// SetBlockchainSpecific sets field value
func (o *ListAllUnconfirmedTransactionsRI) SetBlockchainSpecific(v ListAllUnconfirmedTransactionsRIBS) {
	o.BlockchainSpecific = v
}

func (o ListAllUnconfirmedTransactionsRI) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipients"] = o.Recipients
	}
	if true {
		toSerialize["senders"] = o.Senders
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["transactionId"] = o.TransactionId
	}
	if true {
		toSerialize["blockchainSpecific"] = o.BlockchainSpecific
	}
	return json.Marshal(toSerialize)
}

type NullableListAllUnconfirmedTransactionsRI struct {
	value *ListAllUnconfirmedTransactionsRI
	isSet bool
}

func (v NullableListAllUnconfirmedTransactionsRI) Get() *ListAllUnconfirmedTransactionsRI {
	return v.value
}

func (v *NullableListAllUnconfirmedTransactionsRI) Set(val *ListAllUnconfirmedTransactionsRI) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllUnconfirmedTransactionsRI) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllUnconfirmedTransactionsRI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAllUnconfirmedTransactionsRI(val *ListAllUnconfirmedTransactionsRI) *NullableListAllUnconfirmedTransactionsRI {
	return &NullableListAllUnconfirmedTransactionsRI{value: val, isSet: true}
}

func (v NullableListAllUnconfirmedTransactionsRI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllUnconfirmedTransactionsRI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


