# ListBlockchainEventsSubscriptionsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **int32** | The starting index of the response items, i.e. where the response should start listing the returned items. | 
**Limit** | **int32** | Defines how many items should be returned in the response per page basis. | 
**Total** | **int32** | Defines the total number of items returned in the response. | 
**Items** | [**[]ListBlockchainEventsSubscriptionsResponseItem**](ListBlockchainEventsSubscriptionsResponseItem.md) |  | 

## Methods

### NewListBlockchainEventsSubscriptionsResponseData

`func NewListBlockchainEventsSubscriptionsResponseData(offset int32, limit int32, total int32, items []ListBlockchainEventsSubscriptionsResponseItem, ) *ListBlockchainEventsSubscriptionsResponseData`

NewListBlockchainEventsSubscriptionsResponseData instantiates a new ListBlockchainEventsSubscriptionsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBlockchainEventsSubscriptionsResponseDataWithDefaults

`func NewListBlockchainEventsSubscriptionsResponseDataWithDefaults() *ListBlockchainEventsSubscriptionsResponseData`

NewListBlockchainEventsSubscriptionsResponseDataWithDefaults instantiates a new ListBlockchainEventsSubscriptionsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *ListBlockchainEventsSubscriptionsResponseData) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ListBlockchainEventsSubscriptionsResponseData) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ListBlockchainEventsSubscriptionsResponseData) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *ListBlockchainEventsSubscriptionsResponseData) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListBlockchainEventsSubscriptionsResponseData) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListBlockchainEventsSubscriptionsResponseData) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetTotal

`func (o *ListBlockchainEventsSubscriptionsResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListBlockchainEventsSubscriptionsResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListBlockchainEventsSubscriptionsResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *ListBlockchainEventsSubscriptionsResponseData) GetItems() []ListBlockchainEventsSubscriptionsResponseItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListBlockchainEventsSubscriptionsResponseData) GetItemsOk() (*[]ListBlockchainEventsSubscriptionsResponseItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListBlockchainEventsSubscriptionsResponseData) SetItems(v []ListBlockchainEventsSubscriptionsResponseItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

