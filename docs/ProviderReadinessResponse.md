# ProviderReadinessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ready** | **bool** |  | 
**FromAddress** | Pointer to **string** |  | [optional] 
**VerificationStatus** | Pointer to **string** |  | [optional] 
**PaymentInstructionsReady** | **bool** |  | 
**BlockingReasons** | **[]string** |  | 
**ReadinessHash** | **string** |  | 

## Methods

### NewProviderReadinessResponse

`func NewProviderReadinessResponse(ready bool, paymentInstructionsReady bool, blockingReasons []string, readinessHash string, ) *ProviderReadinessResponse`

NewProviderReadinessResponse instantiates a new ProviderReadinessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderReadinessResponseWithDefaults

`func NewProviderReadinessResponseWithDefaults() *ProviderReadinessResponse`

NewProviderReadinessResponseWithDefaults instantiates a new ProviderReadinessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReady

`func (o *ProviderReadinessResponse) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *ProviderReadinessResponse) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *ProviderReadinessResponse) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetFromAddress

`func (o *ProviderReadinessResponse) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ProviderReadinessResponse) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ProviderReadinessResponse) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *ProviderReadinessResponse) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *ProviderReadinessResponse) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *ProviderReadinessResponse) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *ProviderReadinessResponse) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *ProviderReadinessResponse) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetPaymentInstructionsReady

`func (o *ProviderReadinessResponse) GetPaymentInstructionsReady() bool`

GetPaymentInstructionsReady returns the PaymentInstructionsReady field if non-nil, zero value otherwise.

### GetPaymentInstructionsReadyOk

`func (o *ProviderReadinessResponse) GetPaymentInstructionsReadyOk() (*bool, bool)`

GetPaymentInstructionsReadyOk returns a tuple with the PaymentInstructionsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstructionsReady

`func (o *ProviderReadinessResponse) SetPaymentInstructionsReady(v bool)`

SetPaymentInstructionsReady sets PaymentInstructionsReady field to given value.


### GetBlockingReasons

`func (o *ProviderReadinessResponse) GetBlockingReasons() []string`

GetBlockingReasons returns the BlockingReasons field if non-nil, zero value otherwise.

### GetBlockingReasonsOk

`func (o *ProviderReadinessResponse) GetBlockingReasonsOk() (*[]string, bool)`

GetBlockingReasonsOk returns a tuple with the BlockingReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingReasons

`func (o *ProviderReadinessResponse) SetBlockingReasons(v []string)`

SetBlockingReasons sets BlockingReasons field to given value.


### GetReadinessHash

`func (o *ProviderReadinessResponse) GetReadinessHash() string`

GetReadinessHash returns the ReadinessHash field if non-nil, zero value otherwise.

### GetReadinessHashOk

`func (o *ProviderReadinessResponse) GetReadinessHashOk() (*string, bool)`

GetReadinessHashOk returns a tuple with the ReadinessHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessHash

`func (o *ProviderReadinessResponse) SetReadinessHash(v string)`

SetReadinessHash sets ReadinessHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


