# DebPackageReleaseComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to **string** | A URI of a repository the new content unit should be associated with. | [optional] 
**Package** | **string** | Package that is contained in release_comonent. | 
**ReleaseComponent** | **string** | ReleaseComponent this package is contained in. | 

## Methods

### NewDebPackageReleaseComponent

`func NewDebPackageReleaseComponent(package_ string, releaseComponent string, ) *DebPackageReleaseComponent`

NewDebPackageReleaseComponent instantiates a new DebPackageReleaseComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebPackageReleaseComponentWithDefaults

`func NewDebPackageReleaseComponentWithDefaults() *DebPackageReleaseComponent`

NewDebPackageReleaseComponentWithDefaults instantiates a new DebPackageReleaseComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *DebPackageReleaseComponent) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DebPackageReleaseComponent) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DebPackageReleaseComponent) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DebPackageReleaseComponent) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetPackage

`func (o *DebPackageReleaseComponent) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *DebPackageReleaseComponent) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *DebPackageReleaseComponent) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetReleaseComponent

`func (o *DebPackageReleaseComponent) GetReleaseComponent() string`

GetReleaseComponent returns the ReleaseComponent field if non-nil, zero value otherwise.

### GetReleaseComponentOk

`func (o *DebPackageReleaseComponent) GetReleaseComponentOk() (*string, bool)`

GetReleaseComponentOk returns a tuple with the ReleaseComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseComponent

`func (o *DebPackageReleaseComponent) SetReleaseComponent(v string)`

SetReleaseComponent sets ReleaseComponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


