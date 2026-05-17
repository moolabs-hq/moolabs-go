# AskCustomerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MissingField** | **string** |  | 
**Context** | Pointer to **string** |  | [optional] 

## Methods

### NewAskCustomerRequest

`func NewAskCustomerRequest(missingField string, ) *AskCustomerRequest`

NewAskCustomerRequest instantiates a new AskCustomerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAskCustomerRequestWithDefaults

`func NewAskCustomerRequestWithDefaults() *AskCustomerRequest`

NewAskCustomerRequestWithDefaults instantiates a new AskCustomerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMissingField

`func (o *AskCustomerRequest) GetMissingField() string`

GetMissingField returns the MissingField field if non-nil, zero value otherwise.

### GetMissingFieldOk

`func (o *AskCustomerRequest) GetMissingFieldOk() (*string, bool)`

GetMissingFieldOk returns a tuple with the MissingField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingField

`func (o *AskCustomerRequest) SetMissingField(v string)`

SetMissingField sets MissingField field to given value.


### GetContext

`func (o *AskCustomerRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AskCustomerRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AskCustomerRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *AskCustomerRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


