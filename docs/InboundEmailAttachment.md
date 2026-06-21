# InboundEmailAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | **string** |  | 
**ContentType** | Pointer to **string** |  | [optional] [default to "application/octet-stream"]
**ContentBase64** | **string** |  | 

## Methods

### NewInboundEmailAttachment

`func NewInboundEmailAttachment(filename string, contentBase64 string, ) *InboundEmailAttachment`

NewInboundEmailAttachment instantiates a new InboundEmailAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundEmailAttachmentWithDefaults

`func NewInboundEmailAttachmentWithDefaults() *InboundEmailAttachment`

NewInboundEmailAttachmentWithDefaults instantiates a new InboundEmailAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *InboundEmailAttachment) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *InboundEmailAttachment) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *InboundEmailAttachment) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *InboundEmailAttachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *InboundEmailAttachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *InboundEmailAttachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *InboundEmailAttachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContentBase64

`func (o *InboundEmailAttachment) GetContentBase64() string`

GetContentBase64 returns the ContentBase64 field if non-nil, zero value otherwise.

### GetContentBase64Ok

`func (o *InboundEmailAttachment) GetContentBase64Ok() (*string, bool)`

GetContentBase64Ok returns a tuple with the ContentBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentBase64

`func (o *InboundEmailAttachment) SetContentBase64(v string)`

SetContentBase64 sets ContentBase64 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


