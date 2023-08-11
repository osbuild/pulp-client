# CollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Get href. | [optional] [readonly] 
**Namespace** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Deprecated** | **bool** |  | 
**VersionsUrl** | Pointer to **string** | Get a link to a collection versions list. | [optional] [readonly] 
**HighestVersion** | Pointer to **map[string]interface{}** | Get a highest version and its link. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Get the timestamp of the lowest version CollectionVersion&#39;s created timestamp. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Get the timestamp of the latest version CollectionVersion&#39;s created timestamp. | [optional] [readonly] 
**DownloadCount** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewCollectionResponse

`func NewCollectionResponse(deprecated bool, ) *CollectionResponse`

NewCollectionResponse instantiates a new CollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResponseWithDefaults

`func NewCollectionResponseWithDefaults() *CollectionResponse`

NewCollectionResponseWithDefaults instantiates a new CollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CollectionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CollectionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CollectionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CollectionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetNamespace

`func (o *CollectionResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CollectionResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CollectionResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CollectionResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetName

`func (o *CollectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CollectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CollectionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CollectionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeprecated

`func (o *CollectionResponse) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *CollectionResponse) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *CollectionResponse) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.


### GetVersionsUrl

`func (o *CollectionResponse) GetVersionsUrl() string`

GetVersionsUrl returns the VersionsUrl field if non-nil, zero value otherwise.

### GetVersionsUrlOk

`func (o *CollectionResponse) GetVersionsUrlOk() (*string, bool)`

GetVersionsUrlOk returns a tuple with the VersionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionsUrl

`func (o *CollectionResponse) SetVersionsUrl(v string)`

SetVersionsUrl sets VersionsUrl field to given value.

### HasVersionsUrl

`func (o *CollectionResponse) HasVersionsUrl() bool`

HasVersionsUrl returns a boolean if a field has been set.

### GetHighestVersion

`func (o *CollectionResponse) GetHighestVersion() map[string]interface{}`

GetHighestVersion returns the HighestVersion field if non-nil, zero value otherwise.

### GetHighestVersionOk

`func (o *CollectionResponse) GetHighestVersionOk() (*map[string]interface{}, bool)`

GetHighestVersionOk returns a tuple with the HighestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestVersion

`func (o *CollectionResponse) SetHighestVersion(v map[string]interface{})`

SetHighestVersion sets HighestVersion field to given value.

### HasHighestVersion

`func (o *CollectionResponse) HasHighestVersion() bool`

HasHighestVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CollectionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CollectionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CollectionResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CollectionResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CollectionResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CollectionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDownloadCount

`func (o *CollectionResponse) GetDownloadCount() string`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *CollectionResponse) GetDownloadCountOk() (*string, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *CollectionResponse) SetDownloadCount(v string)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *CollectionResponse) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


