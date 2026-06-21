# QuoteTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**TenantId** | **string** |  | 
**Name** | **string** |  | 
**IsDefault** | **bool** |  | 
**CompanyLegalName** | Pointer to **string** |  | [optional] 
**CompanyAddress** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**AccentColor** | Pointer to **string** |  | [optional] 
**TermsAndConditions** | Pointer to **string** |  | [optional] 
**FooterText** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewQuoteTemplateResponse

`func NewQuoteTemplateResponse(id string, tenantId string, name string, isDefault bool, ) *QuoteTemplateResponse`

NewQuoteTemplateResponse instantiates a new QuoteTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteTemplateResponseWithDefaults

`func NewQuoteTemplateResponseWithDefaults() *QuoteTemplateResponse`

NewQuoteTemplateResponseWithDefaults instantiates a new QuoteTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteTemplateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *QuoteTemplateResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *QuoteTemplateResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *QuoteTemplateResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetName

`func (o *QuoteTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuoteTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuoteTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *QuoteTemplateResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *QuoteTemplateResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *QuoteTemplateResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetCompanyLegalName

`func (o *QuoteTemplateResponse) GetCompanyLegalName() string`

GetCompanyLegalName returns the CompanyLegalName field if non-nil, zero value otherwise.

### GetCompanyLegalNameOk

`func (o *QuoteTemplateResponse) GetCompanyLegalNameOk() (*string, bool)`

GetCompanyLegalNameOk returns a tuple with the CompanyLegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLegalName

`func (o *QuoteTemplateResponse) SetCompanyLegalName(v string)`

SetCompanyLegalName sets CompanyLegalName field to given value.

### HasCompanyLegalName

`func (o *QuoteTemplateResponse) HasCompanyLegalName() bool`

HasCompanyLegalName returns a boolean if a field has been set.

### GetCompanyAddress

`func (o *QuoteTemplateResponse) GetCompanyAddress() string`

GetCompanyAddress returns the CompanyAddress field if non-nil, zero value otherwise.

### GetCompanyAddressOk

`func (o *QuoteTemplateResponse) GetCompanyAddressOk() (*string, bool)`

GetCompanyAddressOk returns a tuple with the CompanyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress

`func (o *QuoteTemplateResponse) SetCompanyAddress(v string)`

SetCompanyAddress sets CompanyAddress field to given value.

### HasCompanyAddress

`func (o *QuoteTemplateResponse) HasCompanyAddress() bool`

HasCompanyAddress returns a boolean if a field has been set.

### GetLogoUrl

`func (o *QuoteTemplateResponse) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *QuoteTemplateResponse) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *QuoteTemplateResponse) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *QuoteTemplateResponse) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetAccentColor

`func (o *QuoteTemplateResponse) GetAccentColor() string`

GetAccentColor returns the AccentColor field if non-nil, zero value otherwise.

### GetAccentColorOk

`func (o *QuoteTemplateResponse) GetAccentColorOk() (*string, bool)`

GetAccentColorOk returns a tuple with the AccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentColor

`func (o *QuoteTemplateResponse) SetAccentColor(v string)`

SetAccentColor sets AccentColor field to given value.

### HasAccentColor

`func (o *QuoteTemplateResponse) HasAccentColor() bool`

HasAccentColor returns a boolean if a field has been set.

### GetTermsAndConditions

`func (o *QuoteTemplateResponse) GetTermsAndConditions() string`

GetTermsAndConditions returns the TermsAndConditions field if non-nil, zero value otherwise.

### GetTermsAndConditionsOk

`func (o *QuoteTemplateResponse) GetTermsAndConditionsOk() (*string, bool)`

GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAndConditions

`func (o *QuoteTemplateResponse) SetTermsAndConditions(v string)`

SetTermsAndConditions sets TermsAndConditions field to given value.

### HasTermsAndConditions

`func (o *QuoteTemplateResponse) HasTermsAndConditions() bool`

HasTermsAndConditions returns a boolean if a field has been set.

### GetFooterText

`func (o *QuoteTemplateResponse) GetFooterText() string`

GetFooterText returns the FooterText field if non-nil, zero value otherwise.

### GetFooterTextOk

`func (o *QuoteTemplateResponse) GetFooterTextOk() (*string, bool)`

GetFooterTextOk returns a tuple with the FooterText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterText

`func (o *QuoteTemplateResponse) SetFooterText(v string)`

SetFooterText sets FooterText field to given value.

### HasFooterText

`func (o *QuoteTemplateResponse) HasFooterText() bool`

HasFooterText returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QuoteTemplateResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QuoteTemplateResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QuoteTemplateResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *QuoteTemplateResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *QuoteTemplateResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QuoteTemplateResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QuoteTemplateResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *QuoteTemplateResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


