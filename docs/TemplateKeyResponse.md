# TemplateKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CanonicalTemplateKey** | **string** |  | 
**DisplayName** | **string** |  | 
**Channel** | Pointer to **string** |  | [optional] [default to "email"]
**KeySource** | **string** |  | 
**VariablePolicy** | **map[string]interface{}** |  | 
**AllowedVariables** | **[]string** |  | 
**RequiredVariables** | **[]string** |  | 
**LegacyAliases** | Pointer to **[]string** |  | [optional] 
**SystemSeeded** | **bool** |  | 
**Selectable** | **bool** |  | 
**Archived** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewTemplateKeyResponse

`func NewTemplateKeyResponse(canonicalTemplateKey string, displayName string, keySource string, variablePolicy map[string]interface{}, allowedVariables []string, requiredVariables []string, systemSeeded bool, selectable bool, ) *TemplateKeyResponse`

NewTemplateKeyResponse instantiates a new TemplateKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateKeyResponseWithDefaults

`func NewTemplateKeyResponseWithDefaults() *TemplateKeyResponse`

NewTemplateKeyResponseWithDefaults instantiates a new TemplateKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TemplateKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateKeyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateKeyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCanonicalTemplateKey

`func (o *TemplateKeyResponse) GetCanonicalTemplateKey() string`

GetCanonicalTemplateKey returns the CanonicalTemplateKey field if non-nil, zero value otherwise.

### GetCanonicalTemplateKeyOk

`func (o *TemplateKeyResponse) GetCanonicalTemplateKeyOk() (*string, bool)`

GetCanonicalTemplateKeyOk returns a tuple with the CanonicalTemplateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalTemplateKey

`func (o *TemplateKeyResponse) SetCanonicalTemplateKey(v string)`

SetCanonicalTemplateKey sets CanonicalTemplateKey field to given value.


### GetDisplayName

`func (o *TemplateKeyResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TemplateKeyResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TemplateKeyResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetChannel

`func (o *TemplateKeyResponse) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *TemplateKeyResponse) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *TemplateKeyResponse) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *TemplateKeyResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetKeySource

`func (o *TemplateKeyResponse) GetKeySource() string`

GetKeySource returns the KeySource field if non-nil, zero value otherwise.

### GetKeySourceOk

`func (o *TemplateKeyResponse) GetKeySourceOk() (*string, bool)`

GetKeySourceOk returns a tuple with the KeySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySource

`func (o *TemplateKeyResponse) SetKeySource(v string)`

SetKeySource sets KeySource field to given value.


### GetVariablePolicy

`func (o *TemplateKeyResponse) GetVariablePolicy() map[string]interface{}`

GetVariablePolicy returns the VariablePolicy field if non-nil, zero value otherwise.

### GetVariablePolicyOk

`func (o *TemplateKeyResponse) GetVariablePolicyOk() (*map[string]interface{}, bool)`

GetVariablePolicyOk returns a tuple with the VariablePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablePolicy

`func (o *TemplateKeyResponse) SetVariablePolicy(v map[string]interface{})`

SetVariablePolicy sets VariablePolicy field to given value.


### GetAllowedVariables

`func (o *TemplateKeyResponse) GetAllowedVariables() []string`

GetAllowedVariables returns the AllowedVariables field if non-nil, zero value otherwise.

### GetAllowedVariablesOk

`func (o *TemplateKeyResponse) GetAllowedVariablesOk() (*[]string, bool)`

GetAllowedVariablesOk returns a tuple with the AllowedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVariables

`func (o *TemplateKeyResponse) SetAllowedVariables(v []string)`

SetAllowedVariables sets AllowedVariables field to given value.


### GetRequiredVariables

`func (o *TemplateKeyResponse) GetRequiredVariables() []string`

GetRequiredVariables returns the RequiredVariables field if non-nil, zero value otherwise.

### GetRequiredVariablesOk

`func (o *TemplateKeyResponse) GetRequiredVariablesOk() (*[]string, bool)`

GetRequiredVariablesOk returns a tuple with the RequiredVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredVariables

`func (o *TemplateKeyResponse) SetRequiredVariables(v []string)`

SetRequiredVariables sets RequiredVariables field to given value.


### GetLegacyAliases

`func (o *TemplateKeyResponse) GetLegacyAliases() []string`

GetLegacyAliases returns the LegacyAliases field if non-nil, zero value otherwise.

### GetLegacyAliasesOk

`func (o *TemplateKeyResponse) GetLegacyAliasesOk() (*[]string, bool)`

GetLegacyAliasesOk returns a tuple with the LegacyAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyAliases

`func (o *TemplateKeyResponse) SetLegacyAliases(v []string)`

SetLegacyAliases sets LegacyAliases field to given value.

### HasLegacyAliases

`func (o *TemplateKeyResponse) HasLegacyAliases() bool`

HasLegacyAliases returns a boolean if a field has been set.

### GetSystemSeeded

`func (o *TemplateKeyResponse) GetSystemSeeded() bool`

GetSystemSeeded returns the SystemSeeded field if non-nil, zero value otherwise.

### GetSystemSeededOk

`func (o *TemplateKeyResponse) GetSystemSeededOk() (*bool, bool)`

GetSystemSeededOk returns a tuple with the SystemSeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemSeeded

`func (o *TemplateKeyResponse) SetSystemSeeded(v bool)`

SetSystemSeeded sets SystemSeeded field to given value.


### GetSelectable

`func (o *TemplateKeyResponse) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *TemplateKeyResponse) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *TemplateKeyResponse) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.


### GetArchived

`func (o *TemplateKeyResponse) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *TemplateKeyResponse) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *TemplateKeyResponse) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *TemplateKeyResponse) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


