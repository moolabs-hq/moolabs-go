# DomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | 
**DomainType** | Pointer to **string** |  | [optional] 
**VerificationStatus** | **string** |  | 
**LastCheckedAt** | **string** |  | 
**DnsRecords** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDomainResponse

`func NewDomainResponse(domain string, verificationStatus string, lastCheckedAt string, ) *DomainResponse`

NewDomainResponse instantiates a new DomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainResponseWithDefaults

`func NewDomainResponseWithDefaults() *DomainResponse`

NewDomainResponseWithDefaults instantiates a new DomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDomainType

`func (o *DomainResponse) GetDomainType() string`

GetDomainType returns the DomainType field if non-nil, zero value otherwise.

### GetDomainTypeOk

`func (o *DomainResponse) GetDomainTypeOk() (*string, bool)`

GetDomainTypeOk returns a tuple with the DomainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainType

`func (o *DomainResponse) SetDomainType(v string)`

SetDomainType sets DomainType field to given value.

### HasDomainType

`func (o *DomainResponse) HasDomainType() bool`

HasDomainType returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *DomainResponse) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *DomainResponse) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *DomainResponse) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.


### GetLastCheckedAt

`func (o *DomainResponse) GetLastCheckedAt() string`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *DomainResponse) GetLastCheckedAtOk() (*string, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *DomainResponse) SetLastCheckedAt(v string)`

SetLastCheckedAt sets LastCheckedAt field to given value.


### GetDnsRecords

`func (o *DomainResponse) GetDnsRecords() map[string]interface{}`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *DomainResponse) GetDnsRecordsOk() (*map[string]interface{}, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *DomainResponse) SetDnsRecords(v map[string]interface{})`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *DomainResponse) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


