# PatchSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scratchpad** | Pointer to **map[string]interface{}** |  | [optional] 
**LineItems** | Pointer to [**[]QuoteLineItemInput**](QuoteLineItemInput.md) |  | [optional] 
**CommercialTerms** | Pointer to **map[string]interface{}** |  | [optional] 
**CreditTerms** | Pointer to **map[string]interface{}** |  | [optional] 

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

### GetCommercialTerms

`func (o *PatchSessionRequest) GetCommercialTerms() map[string]interface{}`

GetCommercialTerms returns the CommercialTerms field if non-nil, zero value otherwise.

### GetCommercialTermsOk

`func (o *PatchSessionRequest) GetCommercialTermsOk() (*map[string]interface{}, bool)`

GetCommercialTermsOk returns a tuple with the CommercialTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialTerms

`func (o *PatchSessionRequest) SetCommercialTerms(v map[string]interface{})`

SetCommercialTerms sets CommercialTerms field to given value.

### HasCommercialTerms

`func (o *PatchSessionRequest) HasCommercialTerms() bool`

HasCommercialTerms returns a boolean if a field has been set.

### GetCreditTerms

`func (o *PatchSessionRequest) GetCreditTerms() map[string]interface{}`

GetCreditTerms returns the CreditTerms field if non-nil, zero value otherwise.

### GetCreditTermsOk

`func (o *PatchSessionRequest) GetCreditTermsOk() (*map[string]interface{}, bool)`

GetCreditTermsOk returns a tuple with the CreditTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditTerms

`func (o *PatchSessionRequest) SetCreditTerms(v map[string]interface{})`

SetCreditTerms sets CreditTerms field to given value.

### HasCreditTerms

`func (o *PatchSessionRequest) HasCreditTerms() bool`

HasCreditTerms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


