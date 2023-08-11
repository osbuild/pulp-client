# PackageUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | ***os.File** | A Python package release file to upload to the index. | 
**Action** | Pointer to **string** | Defaults to &#x60;file_upload&#x60;, don&#39;t change it or request will fail! | [optional] [default to "file_upload"]
**Sha256Digest** | **string** | SHA256 of package to validate upload integrity. | 

## Methods

### NewPackageUpload

`func NewPackageUpload(content *os.File, sha256Digest string, ) *PackageUpload`

NewPackageUpload instantiates a new PackageUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageUploadWithDefaults

`func NewPackageUploadWithDefaults() *PackageUpload`

NewPackageUploadWithDefaults instantiates a new PackageUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PackageUpload) GetContent() *os.File`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PackageUpload) GetContentOk() (**os.File, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PackageUpload) SetContent(v *os.File)`

SetContent sets Content field to given value.


### GetAction

`func (o *PackageUpload) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PackageUpload) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PackageUpload) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PackageUpload) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetSha256Digest

`func (o *PackageUpload) GetSha256Digest() string`

GetSha256Digest returns the Sha256Digest field if non-nil, zero value otherwise.

### GetSha256DigestOk

`func (o *PackageUpload) GetSha256DigestOk() (*string, bool)`

GetSha256DigestOk returns a tuple with the Sha256Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Digest

`func (o *PackageUpload) SetSha256Digest(v string)`

SetSha256Digest sets Sha256Digest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


