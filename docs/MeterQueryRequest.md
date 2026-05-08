# MeterQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Client ID Useful to track progress of a query. | [optional] 
**From** | Pointer to **time.Time** | Start date-time in RFC 3339 format.  Inclusive. | [optional] 
**To** | Pointer to **time.Time** | End date-time in RFC 3339 format.  Inclusive. | [optional] 
**WindowSize** | Pointer to [**WindowSize**](WindowSize.md) | If not specified, a single usage aggregate will be returned for the entirety of the specified period for each subject and group. | [optional] 
**WindowTimeZone** | Pointer to **string** | The value is the name of the time zone as defined in the IANA Time Zone Database (http://www.iana.org/time-zones). If not specified, the UTC timezone will be used. | [optional] [default to "UTC"]
**Subject** | Pointer to **[]string** | Filtering by multiple subjects. | [optional] 
**FilterCustomerId** | Pointer to **[]string** | Filtering by multiple customers. | [optional] 
**FilterGroupBy** | Pointer to **map[string][]string** | Simple filter for group bys with exact match. | [optional] 
**AdvancedMeterGroupByFilters** | Pointer to [**map[string]FilterString**](FilterString.md) | Optional advanced meter group by filters. You can use this to filter for values of the meter groupBy fields. | [optional] 
**GroupBy** | Pointer to **[]string** | If not specified a single aggregate will be returned for each subject and time window. &#x60;subject&#x60; is a reserved group by value. | [optional] 

## Methods

### NewMeterQueryRequest

`func NewMeterQueryRequest() *MeterQueryRequest`

NewMeterQueryRequest instantiates a new MeterQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeterQueryRequestWithDefaults

`func NewMeterQueryRequestWithDefaults() *MeterQueryRequest`

NewMeterQueryRequestWithDefaults instantiates a new MeterQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *MeterQueryRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MeterQueryRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MeterQueryRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *MeterQueryRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetFrom

`func (o *MeterQueryRequest) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MeterQueryRequest) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MeterQueryRequest) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MeterQueryRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *MeterQueryRequest) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MeterQueryRequest) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MeterQueryRequest) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *MeterQueryRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetWindowSize

`func (o *MeterQueryRequest) GetWindowSize() WindowSize`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *MeterQueryRequest) GetWindowSizeOk() (*WindowSize, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *MeterQueryRequest) SetWindowSize(v WindowSize)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *MeterQueryRequest) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.

### GetWindowTimeZone

`func (o *MeterQueryRequest) GetWindowTimeZone() string`

GetWindowTimeZone returns the WindowTimeZone field if non-nil, zero value otherwise.

### GetWindowTimeZoneOk

`func (o *MeterQueryRequest) GetWindowTimeZoneOk() (*string, bool)`

GetWindowTimeZoneOk returns a tuple with the WindowTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowTimeZone

`func (o *MeterQueryRequest) SetWindowTimeZone(v string)`

SetWindowTimeZone sets WindowTimeZone field to given value.

### HasWindowTimeZone

`func (o *MeterQueryRequest) HasWindowTimeZone() bool`

HasWindowTimeZone returns a boolean if a field has been set.

### GetSubject

`func (o *MeterQueryRequest) GetSubject() []string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MeterQueryRequest) GetSubjectOk() (*[]string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MeterQueryRequest) SetSubject(v []string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *MeterQueryRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetFilterCustomerId

`func (o *MeterQueryRequest) GetFilterCustomerId() []string`

GetFilterCustomerId returns the FilterCustomerId field if non-nil, zero value otherwise.

### GetFilterCustomerIdOk

`func (o *MeterQueryRequest) GetFilterCustomerIdOk() (*[]string, bool)`

GetFilterCustomerIdOk returns a tuple with the FilterCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterCustomerId

`func (o *MeterQueryRequest) SetFilterCustomerId(v []string)`

SetFilterCustomerId sets FilterCustomerId field to given value.

### HasFilterCustomerId

`func (o *MeterQueryRequest) HasFilterCustomerId() bool`

HasFilterCustomerId returns a boolean if a field has been set.

### GetFilterGroupBy

`func (o *MeterQueryRequest) GetFilterGroupBy() map[string][]string`

GetFilterGroupBy returns the FilterGroupBy field if non-nil, zero value otherwise.

### GetFilterGroupByOk

`func (o *MeterQueryRequest) GetFilterGroupByOk() (*map[string][]string, bool)`

GetFilterGroupByOk returns a tuple with the FilterGroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroupBy

`func (o *MeterQueryRequest) SetFilterGroupBy(v map[string][]string)`

SetFilterGroupBy sets FilterGroupBy field to given value.

### HasFilterGroupBy

`func (o *MeterQueryRequest) HasFilterGroupBy() bool`

HasFilterGroupBy returns a boolean if a field has been set.

### GetAdvancedMeterGroupByFilters

`func (o *MeterQueryRequest) GetAdvancedMeterGroupByFilters() map[string]FilterString`

GetAdvancedMeterGroupByFilters returns the AdvancedMeterGroupByFilters field if non-nil, zero value otherwise.

### GetAdvancedMeterGroupByFiltersOk

`func (o *MeterQueryRequest) GetAdvancedMeterGroupByFiltersOk() (*map[string]FilterString, bool)`

GetAdvancedMeterGroupByFiltersOk returns a tuple with the AdvancedMeterGroupByFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedMeterGroupByFilters

`func (o *MeterQueryRequest) SetAdvancedMeterGroupByFilters(v map[string]FilterString)`

SetAdvancedMeterGroupByFilters sets AdvancedMeterGroupByFilters field to given value.

### HasAdvancedMeterGroupByFilters

`func (o *MeterQueryRequest) HasAdvancedMeterGroupByFilters() bool`

HasAdvancedMeterGroupByFilters returns a boolean if a field has been set.

### GetGroupBy

`func (o *MeterQueryRequest) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MeterQueryRequest) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MeterQueryRequest) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MeterQueryRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


