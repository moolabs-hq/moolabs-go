# AccountFilterOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceSystems** | **[]string** |  | 
**Subsidiaries** | Pointer to [**[]AccountSubsidiaryOption**](AccountSubsidiaryOption.md) |  | [optional] 

## Methods

### NewAccountFilterOptionsResponse

`func NewAccountFilterOptionsResponse(sourceSystems []string, ) *AccountFilterOptionsResponse`

NewAccountFilterOptionsResponse instantiates a new AccountFilterOptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFilterOptionsResponseWithDefaults

`func NewAccountFilterOptionsResponseWithDefaults() *AccountFilterOptionsResponse`

NewAccountFilterOptionsResponseWithDefaults instantiates a new AccountFilterOptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceSystems

`func (o *AccountFilterOptionsResponse) GetSourceSystems() []string`

GetSourceSystems returns the SourceSystems field if non-nil, zero value otherwise.

### GetSourceSystemsOk

`func (o *AccountFilterOptionsResponse) GetSourceSystemsOk() (*[]string, bool)`

GetSourceSystemsOk returns a tuple with the SourceSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystems

`func (o *AccountFilterOptionsResponse) SetSourceSystems(v []string)`

SetSourceSystems sets SourceSystems field to given value.


### GetSubsidiaries

`func (o *AccountFilterOptionsResponse) GetSubsidiaries() []AccountSubsidiaryOption`

GetSubsidiaries returns the Subsidiaries field if non-nil, zero value otherwise.

### GetSubsidiariesOk

`func (o *AccountFilterOptionsResponse) GetSubsidiariesOk() (*[]AccountSubsidiaryOption, bool)`

GetSubsidiariesOk returns a tuple with the Subsidiaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsidiaries

`func (o *AccountFilterOptionsResponse) SetSubsidiaries(v []AccountSubsidiaryOption)`

SetSubsidiaries sets Subsidiaries field to given value.

### HasSubsidiaries

`func (o *AccountFilterOptionsResponse) HasSubsidiaries() bool`

HasSubsidiaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


