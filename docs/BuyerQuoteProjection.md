# BuyerQuoteProjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | **string** |  | 
**Version** | **int32** |  | 
**Status** | **string** |  | 
**QuoteType** | Pointer to **string** |  | [optional] [default to "new_subscription"]
**BuyerContact** | **map[string]interface{}** |  | 
**LineItems** | [**[]BuyerQuoteLineItem**](BuyerQuoteLineItem.md) |  | 
**CommercialTerms** | **map[string]interface{}** |  | 
**CreditTerms** | **map[string]interface{}** |  | 
**LifecycleTerms** | Pointer to **map[string]interface{}** |  | [optional] 
**Pricing** | [**BuyerQuotePricing**](BuyerQuotePricing.md) |  | 
**SellerBranding** | Pointer to **map[string]interface{}** |  | [optional] 
**TermsAndConditions** | Pointer to **string** |  | [optional] 
**BuyerRequestStatus** | Pointer to **string** |  | [optional] [default to "none"]

## Methods

### NewBuyerQuoteProjection

`func NewBuyerQuoteProjection(quoteId string, version int32, status string, buyerContact map[string]interface{}, lineItems []BuyerQuoteLineItem, commercialTerms map[string]interface{}, creditTerms map[string]interface{}, pricing BuyerQuotePricing, ) *BuyerQuoteProjection`

NewBuyerQuoteProjection instantiates a new BuyerQuoteProjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyerQuoteProjectionWithDefaults

`func NewBuyerQuoteProjectionWithDefaults() *BuyerQuoteProjection`

NewBuyerQuoteProjectionWithDefaults instantiates a new BuyerQuoteProjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *BuyerQuoteProjection) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *BuyerQuoteProjection) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *BuyerQuoteProjection) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetVersion

`func (o *BuyerQuoteProjection) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BuyerQuoteProjection) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BuyerQuoteProjection) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetStatus

`func (o *BuyerQuoteProjection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BuyerQuoteProjection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BuyerQuoteProjection) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetQuoteType

`func (o *BuyerQuoteProjection) GetQuoteType() string`

GetQuoteType returns the QuoteType field if non-nil, zero value otherwise.

### GetQuoteTypeOk

`func (o *BuyerQuoteProjection) GetQuoteTypeOk() (*string, bool)`

GetQuoteTypeOk returns a tuple with the QuoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteType

`func (o *BuyerQuoteProjection) SetQuoteType(v string)`

SetQuoteType sets QuoteType field to given value.

### HasQuoteType

`func (o *BuyerQuoteProjection) HasQuoteType() bool`

HasQuoteType returns a boolean if a field has been set.

### GetBuyerContact

`func (o *BuyerQuoteProjection) GetBuyerContact() map[string]interface{}`

GetBuyerContact returns the BuyerContact field if non-nil, zero value otherwise.

### GetBuyerContactOk

`func (o *BuyerQuoteProjection) GetBuyerContactOk() (*map[string]interface{}, bool)`

GetBuyerContactOk returns a tuple with the BuyerContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerContact

`func (o *BuyerQuoteProjection) SetBuyerContact(v map[string]interface{})`

SetBuyerContact sets BuyerContact field to given value.


### GetLineItems

`func (o *BuyerQuoteProjection) GetLineItems() []BuyerQuoteLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *BuyerQuoteProjection) GetLineItemsOk() (*[]BuyerQuoteLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *BuyerQuoteProjection) SetLineItems(v []BuyerQuoteLineItem)`

SetLineItems sets LineItems field to given value.


### GetCommercialTerms

`func (o *BuyerQuoteProjection) GetCommercialTerms() map[string]interface{}`

GetCommercialTerms returns the CommercialTerms field if non-nil, zero value otherwise.

### GetCommercialTermsOk

`func (o *BuyerQuoteProjection) GetCommercialTermsOk() (*map[string]interface{}, bool)`

GetCommercialTermsOk returns a tuple with the CommercialTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialTerms

`func (o *BuyerQuoteProjection) SetCommercialTerms(v map[string]interface{})`

SetCommercialTerms sets CommercialTerms field to given value.


### GetCreditTerms

`func (o *BuyerQuoteProjection) GetCreditTerms() map[string]interface{}`

GetCreditTerms returns the CreditTerms field if non-nil, zero value otherwise.

### GetCreditTermsOk

`func (o *BuyerQuoteProjection) GetCreditTermsOk() (*map[string]interface{}, bool)`

GetCreditTermsOk returns a tuple with the CreditTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditTerms

`func (o *BuyerQuoteProjection) SetCreditTerms(v map[string]interface{})`

SetCreditTerms sets CreditTerms field to given value.


### GetLifecycleTerms

`func (o *BuyerQuoteProjection) GetLifecycleTerms() map[string]interface{}`

GetLifecycleTerms returns the LifecycleTerms field if non-nil, zero value otherwise.

### GetLifecycleTermsOk

`func (o *BuyerQuoteProjection) GetLifecycleTermsOk() (*map[string]interface{}, bool)`

GetLifecycleTermsOk returns a tuple with the LifecycleTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleTerms

`func (o *BuyerQuoteProjection) SetLifecycleTerms(v map[string]interface{})`

SetLifecycleTerms sets LifecycleTerms field to given value.

### HasLifecycleTerms

`func (o *BuyerQuoteProjection) HasLifecycleTerms() bool`

HasLifecycleTerms returns a boolean if a field has been set.

### GetPricing

`func (o *BuyerQuoteProjection) GetPricing() BuyerQuotePricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *BuyerQuoteProjection) GetPricingOk() (*BuyerQuotePricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *BuyerQuoteProjection) SetPricing(v BuyerQuotePricing)`

SetPricing sets Pricing field to given value.


### GetSellerBranding

`func (o *BuyerQuoteProjection) GetSellerBranding() map[string]interface{}`

GetSellerBranding returns the SellerBranding field if non-nil, zero value otherwise.

### GetSellerBrandingOk

`func (o *BuyerQuoteProjection) GetSellerBrandingOk() (*map[string]interface{}, bool)`

GetSellerBrandingOk returns a tuple with the SellerBranding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerBranding

`func (o *BuyerQuoteProjection) SetSellerBranding(v map[string]interface{})`

SetSellerBranding sets SellerBranding field to given value.

### HasSellerBranding

`func (o *BuyerQuoteProjection) HasSellerBranding() bool`

HasSellerBranding returns a boolean if a field has been set.

### GetTermsAndConditions

`func (o *BuyerQuoteProjection) GetTermsAndConditions() string`

GetTermsAndConditions returns the TermsAndConditions field if non-nil, zero value otherwise.

### GetTermsAndConditionsOk

`func (o *BuyerQuoteProjection) GetTermsAndConditionsOk() (*string, bool)`

GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditions

`func (o *BuyerQuoteProjection) SetTermsAndConditions(v string)`

SetTermsAndConditions sets TermsAndConditions field to given value.

### HasTermsAndConditions

`func (o *BuyerQuoteProjection) HasTermsAndConditions() bool`

HasTermsAndConditions returns a boolean if a field has been set.

### GetBuyerRequestStatus

`func (o *BuyerQuoteProjection) GetBuyerRequestStatus() string`

GetBuyerRequestStatus returns the BuyerRequestStatus field if non-nil, zero value otherwise.

### GetBuyerRequestStatusOk

`func (o *BuyerQuoteProjection) GetBuyerRequestStatusOk() (*string, bool)`

GetBuyerRequestStatusOk returns a tuple with the BuyerRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerRequestStatus

`func (o *BuyerQuoteProjection) SetBuyerRequestStatus(v string)`

SetBuyerRequestStatus sets BuyerRequestStatus field to given value.

### HasBuyerRequestStatus

`func (o *BuyerQuoteProjection) HasBuyerRequestStatus() bool`

HasBuyerRequestStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


