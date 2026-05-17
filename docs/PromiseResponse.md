# PromiseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CaseId** | **string** |  | 
**InvoiceId** | Pointer to **string** |  | [optional] 
**PromisedAmountMicros** | **int32** |  | 
**FulfilledAmountMicros** | Pointer to **int32** |  | [optional] [default to 0]
**CurrencyCode** | **string** |  | 
**PromisedDate** | **string** |  | 
**Status** | **string** |  | 
**BrokenAt** | Pointer to **time.Time** |  | [optional] 
**CapturedFromChannel** | **string** |  | 
**CapturedBy** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewPromiseResponse

`func NewPromiseResponse(id string, caseId string, promisedAmountMicros int32, currencyCode string, promisedDate string, status string, capturedFromChannel string, capturedBy string, createdAt time.Time, updatedAt time.Time, ) *PromiseResponse`

NewPromiseResponse instantiates a new PromiseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromiseResponseWithDefaults

`func NewPromiseResponseWithDefaults() *PromiseResponse`

NewPromiseResponseWithDefaults instantiates a new PromiseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromiseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromiseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromiseResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCaseId

`func (o *PromiseResponse) GetCaseId() string`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *PromiseResponse) GetCaseIdOk() (*string, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *PromiseResponse) SetCaseId(v string)`

SetCaseId sets CaseId field to given value.


### GetInvoiceId

`func (o *PromiseResponse) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *PromiseResponse) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *PromiseResponse) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *PromiseResponse) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetPromisedAmountMicros

`func (o *PromiseResponse) GetPromisedAmountMicros() int32`

GetPromisedAmountMicros returns the PromisedAmountMicros field if non-nil, zero value otherwise.

### GetPromisedAmountMicrosOk

`func (o *PromiseResponse) GetPromisedAmountMicrosOk() (*int32, bool)`

GetPromisedAmountMicrosOk returns a tuple with the PromisedAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedAmountMicros

`func (o *PromiseResponse) SetPromisedAmountMicros(v int32)`

SetPromisedAmountMicros sets PromisedAmountMicros field to given value.


### GetFulfilledAmountMicros

`func (o *PromiseResponse) GetFulfilledAmountMicros() int32`

GetFulfilledAmountMicros returns the FulfilledAmountMicros field if non-nil, zero value otherwise.

### GetFulfilledAmountMicrosOk

`func (o *PromiseResponse) GetFulfilledAmountMicrosOk() (*int32, bool)`

GetFulfilledAmountMicrosOk returns a tuple with the FulfilledAmountMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledAmountMicros

`func (o *PromiseResponse) SetFulfilledAmountMicros(v int32)`

SetFulfilledAmountMicros sets FulfilledAmountMicros field to given value.

### HasFulfilledAmountMicros

`func (o *PromiseResponse) HasFulfilledAmountMicros() bool`

HasFulfilledAmountMicros returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *PromiseResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PromiseResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PromiseResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetPromisedDate

`func (o *PromiseResponse) GetPromisedDate() string`

GetPromisedDate returns the PromisedDate field if non-nil, zero value otherwise.

### GetPromisedDateOk

`func (o *PromiseResponse) GetPromisedDateOk() (*string, bool)`

GetPromisedDateOk returns a tuple with the PromisedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromisedDate

`func (o *PromiseResponse) SetPromisedDate(v string)`

SetPromisedDate sets PromisedDate field to given value.


### GetStatus

`func (o *PromiseResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PromiseResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PromiseResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBrokenAt

`func (o *PromiseResponse) GetBrokenAt() time.Time`

GetBrokenAt returns the BrokenAt field if non-nil, zero value otherwise.

### GetBrokenAtOk

`func (o *PromiseResponse) GetBrokenAtOk() (*time.Time, bool)`

GetBrokenAtOk returns a tuple with the BrokenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokenAt

`func (o *PromiseResponse) SetBrokenAt(v time.Time)`

SetBrokenAt sets BrokenAt field to given value.

### HasBrokenAt

`func (o *PromiseResponse) HasBrokenAt() bool`

HasBrokenAt returns a boolean if a field has been set.

### GetCapturedFromChannel

`func (o *PromiseResponse) GetCapturedFromChannel() string`

GetCapturedFromChannel returns the CapturedFromChannel field if non-nil, zero value otherwise.

### GetCapturedFromChannelOk

`func (o *PromiseResponse) GetCapturedFromChannelOk() (*string, bool)`

GetCapturedFromChannelOk returns a tuple with the CapturedFromChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedFromChannel

`func (o *PromiseResponse) SetCapturedFromChannel(v string)`

SetCapturedFromChannel sets CapturedFromChannel field to given value.


### GetCapturedBy

`func (o *PromiseResponse) GetCapturedBy() string`

GetCapturedBy returns the CapturedBy field if non-nil, zero value otherwise.

### GetCapturedByOk

`func (o *PromiseResponse) GetCapturedByOk() (*string, bool)`

GetCapturedByOk returns a tuple with the CapturedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturedBy

`func (o *PromiseResponse) SetCapturedBy(v string)`

SetCapturedBy sets CapturedBy field to given value.


### GetCreatedAt

`func (o *PromiseResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PromiseResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PromiseResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PromiseResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PromiseResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PromiseResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


