# PatchSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scratchpad** | Pointer to **map[string]interface{}** |  | [optional] 
**LineItems** | Pointer to [**[]QuoteLineItemInput**](QuoteLineItemInput.md) |  | [optional] 

## Methods

### NewPatchSessionRequest

`func NewPatchSessionRequest() *PatchSessionRequest`

NewPatchSessionRequest instantiates a new PatchSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchSessionRequestWithDefaults

`func NewPatchSessionRequestWithDefaults() *PatchSessionRequest`

NewPatchSessionRequestWithDefaults instantiates a new PatchSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScratchpad

`func (o *PatchSessionRequest) GetScratchpad() map[string]interface{}`

GetScratchpad returns the Scratchpad field if non-nil, zero value otherwise.

### GetScratchpadOk

`func (o *PatchSessionRequest) GetScratchpadOk() (*map[string]interface{}, bool)`

GetScratchpadOk returns a tuple with the Scratchpad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScratchpad

`func (o *PatchSessionRequest) SetScratchpad(v map[string]interface{})`

SetScratchpad sets Scratchpad field to given value.

### HasScratchpad

`func (o *PatchSessionRequest) HasScratchpad() bool`

HasScratchpad returns a boolean if a field has been set.

### GetLineItems

`func (o *PatchSessionRequest) GetLineItems() []QuoteLineItemInput`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *PatchSessionRequest) GetLineItemsOk() (*[]QuoteLineItemInput, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *PatchSessionRequest) SetLineItems(v []QuoteLineItemInput)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *PatchSessionRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


