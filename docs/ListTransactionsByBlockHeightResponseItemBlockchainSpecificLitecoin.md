# ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locktime** | **int32** | Represents the time at which a particular transaction can be added to the blockchain. | 
**Size** | **int32** | Represents the total size of this transaction. | 
**Version** | **int32** | Represents transaction version number. | 
**Vin** | [**[]ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin**](ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin.md) | Represents the transaction inputs. | 
**Vout** | [**[]ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVout**](ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVout.md) | Represents the transaction outputs. | 
**Vsize** | **int32** | Represents the virtual size of this transaction. | 

## Methods

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin(locktime int32, size int32, version int32, vin []ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin, vout []ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVout, vsize int32, ) *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinWithDefaults

`func NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinWithDefaults() *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin`

NewListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinWithDefaults instantiates a new ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocktime

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetLocktime() int32`

GetLocktime returns the Locktime field if non-nil, zero value otherwise.

### GetLocktimeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetLocktimeOk() (*int32, bool)`

GetLocktimeOk returns a tuple with the Locktime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocktime

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) SetLocktime(v int32)`

SetLocktime sets Locktime field to given value.


### GetSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) SetSize(v int32)`

SetSize sets Size field to given value.


### GetVersion

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVin

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetVin() []ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetVinOk() (*[]ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) SetVin(v []ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVin)`

SetVin sets Vin field to given value.


### GetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetVout() []ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVout`

GetVout returns the Vout field if non-nil, zero value otherwise.

### GetVoutOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetVoutOk() (*[]ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVout, bool)`

GetVoutOk returns a tuple with the Vout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVout

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) SetVout(v []ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoinVout)`

SetVout sets Vout field to given value.


### GetVsize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetVsize() int32`

GetVsize returns the Vsize field if non-nil, zero value otherwise.

### GetVsizeOk

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) GetVsizeOk() (*int32, bool)`

GetVsizeOk returns a tuple with the Vsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsize

`func (o *ListTransactionsByBlockHeightResponseItemBlockchainSpecificLitecoin) SetVsize(v int32)`

SetVsize sets Vsize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

