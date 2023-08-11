# GalaxyCollectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Namespace** | Pointer to **map[string]interface{}** | Create a namespace dict. | [optional] [readonly] 
**Href** | Pointer to **string** | Get href. | [optional] [readonly] 
**VersionsUrl** | Pointer to **string** | Get versions_url. | [optional] [readonly] 
**Created** | **time.Time** |  | 
**Modified** | **time.Time** |  | 
**LatestVersion** | Pointer to **map[string]interface{}** | Get latest version. | [optional] [readonly] 

## Methods

### NewGalaxyCollectionResponse

`func NewGalaxyCollectionResponse(id string, name string, created time.Time, modified time.Time, ) *GalaxyCollectionResponse`

NewGalaxyCollectionResponse instantiates a new GalaxyCollectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGalaxyCollectionResponseWithDefaults

`func NewGalaxyCollectionResponseWithDefaults() *GalaxyCollectionResponse`

NewGalaxyCollectionResponseWithDefaults instantiates a new GalaxyCollectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GalaxyCollectionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GalaxyCollectionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GalaxyCollectionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GalaxyCollectionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GalaxyCollectionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GalaxyCollectionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *GalaxyCollectionResponse) GetNamespace() map[string]interface{}`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GalaxyCollectionResponse) GetNamespaceOk() (*map[string]interface{}, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GalaxyCollectionResponse) SetNamespace(v map[string]interface{})`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GalaxyCollectionResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetHref

`func (o *GalaxyCollectionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GalaxyCollectionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GalaxyCollectionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *GalaxyCollectionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetVersionsUrl

`func (o *GalaxyCollectionResponse) GetVersionsUrl() string`

GetVersionsUrl returns the VersionsUrl field if non-nil, zero value otherwise.

### GetVersionsUrlOk

`func (o *GalaxyCollectionResponse) GetVersionsUrlOk() (*string, bool)`

GetVersionsUrlOk returns a tuple with the VersionsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionsUrl

`func (o *GalaxyCollectionResponse) SetVersionsUrl(v string)`

SetVersionsUrl sets VersionsUrl field to given value.

### HasVersionsUrl

`func (o *GalaxyCollectionResponse) HasVersionsUrl() bool`

HasVersionsUrl returns a boolean if a field has been set.

### GetCreated

`func (o *GalaxyCollectionResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GalaxyCollectionResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GalaxyCollectionResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *GalaxyCollectionResponse) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *GalaxyCollectionResponse) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *GalaxyCollectionResponse) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetLatestVersion

`func (o *GalaxyCollectionResponse) GetLatestVersion() map[string]interface{}`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *GalaxyCollectionResponse) GetLatestVersionOk() (*map[string]interface{}, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *GalaxyCollectionResponse) SetLatestVersion(v map[string]interface{})`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *GalaxyCollectionResponse) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


