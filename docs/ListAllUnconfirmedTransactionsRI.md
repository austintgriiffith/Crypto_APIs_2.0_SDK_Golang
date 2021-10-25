# ListAllUnconfirmedTransactionsRI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | [**[]ListUnconfirmedTransactionsByAddressRIRecipients**](ListUnconfirmedTransactionsByAddressRIRecipients.md) | Represents a list of recipient addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Senders** | [**[]ListUnconfirmedTransactionsByAddressRISenders**](ListUnconfirmedTransactionsByAddressRISenders.md) | Represents a list of sender addresses with the respective amounts. In account-based protocols like Ethereum there is only one address in this list. | 
**Timestamp** | **int32** | Defines the exact date/time in Unix Timestamp when this transaction was mined, confirmed or first seen in Mempool, if it is unconfirmed. | 
**TransactionId** | **string** | Represents the unique identifier of a transaction, i.e. it could be &#x60;transactionId&#x60; in UTXO-based protocols like Bitcoin, and transaction &#x60;hash&#x60; in Ethereum blockchain. | 
**BlockchainSpecific** | [**ListAllUnconfirmedTransactionsRIBS**](ListAllUnconfirmedTransactionsRIBS.md) |  | 

## Methods

### NewListAllUnconfirmedTransactionsRI

`func NewListAllUnconfirmedTransactionsRI(recipients []ListUnconfirmedTransactionsByAddressRIRecipients, senders []ListUnconfirmedTransactionsByAddressRISenders, timestamp int32, transactionId string, blockchainSpecific ListAllUnconfirmedTransactionsRIBS, ) *ListAllUnconfirmedTransactionsRI`

NewListAllUnconfirmedTransactionsRI instantiates a new ListAllUnconfirmedTransactionsRI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllUnconfirmedTransactionsRIWithDefaults

`func NewListAllUnconfirmedTransactionsRIWithDefaults() *ListAllUnconfirmedTransactionsRI`

NewListAllUnconfirmedTransactionsRIWithDefaults instantiates a new ListAllUnconfirmedTransactionsRI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *ListAllUnconfirmedTransactionsRI) GetRecipients() []ListUnconfirmedTransactionsByAddressRIRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListAllUnconfirmedTransactionsRI) GetRecipientsOk() (*[]ListUnconfirmedTransactionsByAddressRIRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListAllUnconfirmedTransactionsRI) SetRecipients(v []ListUnconfirmedTransactionsByAddressRIRecipients)`

SetRecipients sets Recipients field to given value.


### GetSenders

`func (o *ListAllUnconfirmedTransactionsRI) GetSenders() []ListUnconfirmedTransactionsByAddressRISenders`

GetSenders returns the Senders field if non-nil, zero value otherwise.

### GetSendersOk

`func (o *ListAllUnconfirmedTransactionsRI) GetSendersOk() (*[]ListUnconfirmedTransactionsByAddressRISenders, bool)`

GetSendersOk returns a tuple with the Senders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenders

`func (o *ListAllUnconfirmedTransactionsRI) SetSenders(v []ListUnconfirmedTransactionsByAddressRISenders)`

SetSenders sets Senders field to given value.


### GetTimestamp

`func (o *ListAllUnconfirmedTransactionsRI) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListAllUnconfirmedTransactionsRI) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListAllUnconfirmedTransactionsRI) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactionId

`func (o *ListAllUnconfirmedTransactionsRI) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ListAllUnconfirmedTransactionsRI) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ListAllUnconfirmedTransactionsRI) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetBlockchainSpecific

`func (o *ListAllUnconfirmedTransactionsRI) GetBlockchainSpecific() ListAllUnconfirmedTransactionsRIBS`

GetBlockchainSpecific returns the BlockchainSpecific field if non-nil, zero value otherwise.

### GetBlockchainSpecificOk

`func (o *ListAllUnconfirmedTransactionsRI) GetBlockchainSpecificOk() (*ListAllUnconfirmedTransactionsRIBS, bool)`

GetBlockchainSpecificOk returns a tuple with the BlockchainSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockchainSpecific

`func (o *ListAllUnconfirmedTransactionsRI) SetBlockchainSpecific(v ListAllUnconfirmedTransactionsRIBS)`

SetBlockchainSpecific sets BlockchainSpecific field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

