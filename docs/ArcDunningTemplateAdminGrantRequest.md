# ArcDunningTemplateAdminGrantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserSubject** | **string** |  | 
**GrantReason** | Pointer to **string** |  | [optional] [default to "Granted from ARC Communications Settings"]

## Methods

### NewArcDunningTemplateAdminGrantRequest

`func NewArcDunningTemplateAdminGrantRequest(userSubject string, ) *ArcDunningTemplateAdminGrantRequest`

NewArcDunningTemplateAdminGrantRequest instantiates a new ArcDunningTemplateAdminGrantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcDunningTemplateAdminGrantRequestWithDefaults

`func NewArcDunningTemplateAdminGrantRequestWithDefaults() *ArcDunningTemplateAdminGrantRequest`

NewArcDunningTemplateAdminGrantRequestWithDefaults instantiates a new ArcDunningTemplateAdminGrantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserSubject

`func (o *ArcDunningTemplateAdminGrantRequest) GetUserSubject() string`

GetUserSubject returns the UserSubject field if non-nil, zero value otherwise.

### GetUserSubjectOk

`func (o *ArcDunningTemplateAdminGrantRequest) GetUserSubjectOk() (*string, bool)`

GetUserSubjectOk returns a tuple with the UserSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSubject

`func (o *ArcDunningTemplateAdminGrantRequest) SetUserSubject(v string)`

SetUserSubject sets UserSubject field to given value.


### GetGrantReason

`func (o *ArcDunningTemplateAdminGrantRequest) GetGrantReason() string`

GetGrantReason returns the GrantReason field if non-nil, zero value otherwise.

### GetGrantReasonOk

`func (o *ArcDunningTemplateAdminGrantRequest) GetGrantReasonOk() (*string, bool)`

GetGrantReasonOk returns a tuple with the GrantReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantReason

`func (o *ArcDunningTemplateAdminGrantRequest) SetGrantReason(v string)`

SetGrantReason sets GrantReason field to given value.

### HasGrantReason

`func (o *ArcDunningTemplateAdminGrantRequest) HasGrantReason() bool`

HasGrantReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


