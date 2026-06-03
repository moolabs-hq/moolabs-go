# ArcDunningPermissionRevokeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevokeReason** | **string** |  | 
**IdempotencyKey** | Pointer to **string** |  | [optional] 

## Methods

### NewArcDunningPermissionRevokeRequest

`func NewArcDunningPermissionRevokeRequest(revokeReason string, ) *ArcDunningPermissionRevokeRequest`

NewArcDunningPermissionRevokeRequest instantiates a new ArcDunningPermissionRevokeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcDunningPermissionRevokeRequestWithDefaults

`func NewArcDunningPermissionRevokeRequestWithDefaults() *ArcDunningPermissionRevokeRequest`

NewArcDunningPermissionRevokeRequestWithDefaults instantiates a new ArcDunningPermissionRevokeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevokeReason

`func (o *ArcDunningPermissionRevokeRequest) GetRevokeReason() string`

GetRevokeReason returns the RevokeReason field if non-nil, zero value otherwise.

### GetRevokeReasonOk

`func (o *ArcDunningPermissionRevokeRequest) GetRevokeReasonOk() (*string, bool)`

GetRevokeReasonOk returns a tuple with the RevokeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeReason

`func (o *ArcDunningPermissionRevokeRequest) SetRevokeReason(v string)`

SetRevokeReason sets RevokeReason field to given value.


### GetIdempotencyKey

`func (o *ArcDunningPermissionRevokeRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ArcDunningPermissionRevokeRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ArcDunningPermissionRevokeRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *ArcDunningPermissionRevokeRequest) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


