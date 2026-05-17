# PortalPTPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | **string** |  | 
**PromisedAmountMicros** | **int32** |  | 
**PromisedDate** | **string** |  | 
**CurrencyCode** | Pointer to **string** |  | [optional] [default to "USD"]

## Methods

### NewPortalPTPRequest

`func NewPortalPTPRequest(caseId string, promisedAmountMicros int32, promisedDate string, ) *PortalPTPRequest`

NewPortalPTPRequest instantiates a new PortalPTPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalPTPRequestWithDefaults

`func NewPortalPTPRequestWithDefaults() *PortalPTPRequest`

NewPortalPTPRequestWithDefaults instantiates a new PortalPTPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseId

`func (o *PortalPTPRequest) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *PortalPTPRequest) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *PortalPTPRequest) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetPromisedAmountMicros

`func (o *PortalPTPRequest) GetPromisedAmountMicros() int32`

GetPromisedAmountMicros returns the PromisedAmountMicros field if non-nil, zero value otherwise.

### GetPromisedAmountMicrosOk

`func (o *PortalPTPRequest) GetPromisedAmountMicrosOk() (*int32, bool)`

GetPromisedAmountMicrosOk returns a tuple with the PromisedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedAmountMicros

`func (o *PortalPTPRequest) SetPromisedAmountMicros(v int32)`

SetPromisedAmountMicros sets PromisedAmountMicros field to given value.


### GetPromisedDate

`func (o *PortalPTPRequest) GetPromisedDate() string`

GetPromisedDate returns the PromisedDate field if non-nil, zero value otherwise.

### GetPromisedDateOk

`func (o *PortalPTPRequest) GetPromisedDateOk() (*string, bool)`

GetPromisedDateOk returns a tuple with the PromisedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedDate

`func (o *PortalPTPRequest) SetPromisedDate(v string)`

SetPromisedDate sets PromisedDate field to given value.


### GetCurrencyCode

`func (o *PortalPTPRequest) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PortalPTPRequest) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PortalPTPRequest) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PortalPTPRequest) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


