# ListXRPRippleTransactionsByAddress403Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | Specifies the version of the API that incorporates this endpoint. | 
**RequestId** | **string** | Defines the ID of the request. The &#x60;requestId&#x60; is generated by Crypto APIs and it&#39;s unique for every request. | 
**Context** | Pointer to **string** | In batch situations the user can use the context to correlate responses with requests. This property is present regardless of whether the response was successful or returned as an error. &#x60;context&#x60; is specified by the user. | [optional] 
**Error** | [**ListXRPRippleTransactionsByAddressE403**](ListXRPRippleTransactionsByAddressE403.md) |  | 

## Methods

### NewListXRPRippleTransactionsByAddress403Response

`func NewListXRPRippleTransactionsByAddress403Response(apiVersion string, requestId string, error_ ListXRPRippleTransactionsByAddressE403, ) *ListXRPRippleTransactionsByAddress403Response`

NewListXRPRippleTransactionsByAddress403Response instantiates a new ListXRPRippleTransactionsByAddress403Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListXRPRippleTransactionsByAddress403ResponseWithDefaults

`func NewListXRPRippleTransactionsByAddress403ResponseWithDefaults() *ListXRPRippleTransactionsByAddress403Response`

NewListXRPRippleTransactionsByAddress403ResponseWithDefaults instantiates a new ListXRPRippleTransactionsByAddress403Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ListXRPRippleTransactionsByAddress403Response) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ListXRPRippleTransactionsByAddress403Response) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ListXRPRippleTransactionsByAddress403Response) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetRequestId

`func (o *ListXRPRippleTransactionsByAddress403Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ListXRPRippleTransactionsByAddress403Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ListXRPRippleTransactionsByAddress403Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetContext

`func (o *ListXRPRippleTransactionsByAddress403Response) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ListXRPRippleTransactionsByAddress403Response) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ListXRPRippleTransactionsByAddress403Response) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ListXRPRippleTransactionsByAddress403Response) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetError

`func (o *ListXRPRippleTransactionsByAddress403Response) GetError() ListXRPRippleTransactionsByAddressE403`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListXRPRippleTransactionsByAddress403Response) GetErrorOk() (*ListXRPRippleTransactionsByAddressE403, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListXRPRippleTransactionsByAddress403Response) SetError(v ListXRPRippleTransactionsByAddressE403)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


