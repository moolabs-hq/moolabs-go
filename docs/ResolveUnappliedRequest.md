# ResolveUnappliedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disposition** | **string** | One of: applied_as_credit, refunded, abandoned_property, operational_void, de_minimis_writeoff, reversal_return | 

## Methods

### NewResolveUnappliedRequest

`func NewResolveUnappliedRequest(disposition string, ) *ResolveUnappliedRequest`

NewResolveUnappliedRequest instantiates a new ResolveUnappliedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolveUnappliedRequestWithDefaults

`func NewResolveUnappliedRequestWithDefaults() *ResolveUnappliedRequest`

NewResolveUnappliedRequestWithDefaults instantiates a new ResolveUnappliedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisposition

`func (o *ResolveUnappliedRequest) GetDisposition() string`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *ResolveUnappliedRequest) GetDispositionOk() (*string, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *ResolveUnappliedRequest) SetDisposition(v string)`

SetDisposition sets Disposition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


