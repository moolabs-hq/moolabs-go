# ValueRecognitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JournalEntryId** | **string** |  | 
**BurnAllocationsProcessed** | **int32** |  | 
**TotalUsdMicros** | **int32** |  | 
**PostingsCreated** | **[]string** |  | 

## Methods

### NewValueRecognitionResponse

`func NewValueRecognitionResponse(journalEntryId string, burnAllocationsProcessed int32, totalUsdMicros int32, postingsCreated []string, ) *ValueRecognitionResponse`

NewValueRecognitionResponse instantiates a new ValueRecognitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueRecognitionResponseWithDefaults

`func NewValueRecognitionResponseWithDefaults() *ValueRecognitionResponse`

NewValueRecognitionResponseWithDefaults instantiates a new ValueRecognitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJournalEntryId

`func (o *ValueRecognitionResponse) GetJournalEntryId() string`

GetJournalEntryId returns the JournalEntryId field if non-nil, zero value otherwise.

### GetJournalEntryIdOk

`func (o *ValueRecognitionResponse) GetJournalEntryIdOk() (*string, bool)`

GetJournalEntryIdOk returns a tuple with the JournalEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJournalEntryId

`func (o *ValueRecognitionResponse) SetJournalEntryId(v string)`

SetJournalEntryId sets JournalEntryId field to given value.


### GetBurnAllocationsProcessed

`func (o *ValueRecognitionResponse) GetBurnAllocationsProcessed() int32`

GetBurnAllocationsProcessed returns the BurnAllocationsProcessed field if non-nil, zero value otherwise.

### GetBurnAllocationsProcessedOk

`func (o *ValueRecognitionResponse) GetBurnAllocationsProcessedOk() (*int32, bool)`

GetBurnAllocationsProcessedOk returns a tuple with the BurnAllocationsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnAllocationsProcessed

`func (o *ValueRecognitionResponse) SetBurnAllocationsProcessed(v int32)`

SetBurnAllocationsProcessed sets BurnAllocationsProcessed field to given value.


### GetTotalUsdMicros

`func (o *ValueRecognitionResponse) GetTotalUsdMicros() int32`

GetTotalUsdMicros returns the TotalUsdMicros field if non-nil, zero value otherwise.

### GetTotalUsdMicrosOk

`func (o *ValueRecognitionResponse) GetTotalUsdMicrosOk() (*int32, bool)`

GetTotalUsdMicrosOk returns a tuple with the TotalUsdMicros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsdMicros

`func (o *ValueRecognitionResponse) SetTotalUsdMicros(v int32)`

SetTotalUsdMicros sets TotalUsdMicros field to given value.


### GetPostingsCreated

`func (o *ValueRecognitionResponse) GetPostingsCreated() []string`

GetPostingsCreated returns the PostingsCreated field if non-nil, zero value otherwise.

### GetPostingsCreatedOk

`func (o *ValueRecognitionResponse) GetPostingsCreatedOk() (*[]string, bool)`

GetPostingsCreatedOk returns a tuple with the PostingsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingsCreated

`func (o *ValueRecognitionResponse) SetPostingsCreated(v []string)`

SetPostingsCreated sets PostingsCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


