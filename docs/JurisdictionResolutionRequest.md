# JurisdictionResolutionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | **string** |  | 
**RegionCode** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewJurisdictionResolutionRequest

`func NewJurisdictionResolutionRequest(countryCode string, ) *JurisdictionResolutionRequest`

NewJurisdictionResolutionRequest instantiates a new JurisdictionResolutionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJurisdictionResolutionRequestWithDefaults

`func NewJurisdictionResolutionRequestWithDefaults() *JurisdictionResolutionRequest`

NewJurisdictionResolutionRequestWithDefaults instantiates a new JurisdictionResolutionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *JurisdictionResolutionRequest) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *JurisdictionResolutionRequest) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *JurisdictionResolutionRequest) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetRegionCode

`func (o *JurisdictionResolutionRequest) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *JurisdictionResolutionRequest) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *JurisdictionResolutionRequest) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *JurisdictionResolutionRequest) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetTimezone

`func (o *JurisdictionResolutionRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *JurisdictionResolutionRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *JurisdictionResolutionRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *JurisdictionResolutionRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


