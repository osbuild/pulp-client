# PatchedansibleAnsibleNamespaceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Required named, only accepts lowercase, numbers and underscores. | [optional] 
**Company** | Pointer to **string** | Optional namespace company owner. | [optional] 
**Email** | Pointer to **string** | Optional namespace contact email. | [optional] 
**Description** | Pointer to **string** | Optional short description. | [optional] 
**Resources** | Pointer to **string** | Optional resource page in markdown format. | [optional] 
**Links** | Pointer to [**[]NamespaceLink**](NamespaceLink.md) | Labeled related links. | [optional] 
**Avatar** | Pointer to ***os.File** | Optional avatar image for Namespace | [optional] 

## Methods

### NewPatchedansibleAnsibleNamespaceMetadata

`func NewPatchedansibleAnsibleNamespaceMetadata() *PatchedansibleAnsibleNamespaceMetadata`

NewPatchedansibleAnsibleNamespaceMetadata instantiates a new PatchedansibleAnsibleNamespaceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedansibleAnsibleNamespaceMetadataWithDefaults

`func NewPatchedansibleAnsibleNamespaceMetadataWithDefaults() *PatchedansibleAnsibleNamespaceMetadata`

NewPatchedansibleAnsibleNamespaceMetadataWithDefaults instantiates a new PatchedansibleAnsibleNamespaceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedansibleAnsibleNamespaceMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedansibleAnsibleNamespaceMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCompany

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PatchedansibleAnsibleNamespaceMetadata) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PatchedansibleAnsibleNamespaceMetadata) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEmail

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedansibleAnsibleNamespaceMetadata) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedansibleAnsibleNamespaceMetadata) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedansibleAnsibleNamespaceMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedansibleAnsibleNamespaceMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PatchedansibleAnsibleNamespaceMetadata) SetResources(v string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *PatchedansibleAnsibleNamespaceMetadata) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetLinks

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetLinks() []NamespaceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetLinksOk() (*[]NamespaceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PatchedansibleAnsibleNamespaceMetadata) SetLinks(v []NamespaceLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PatchedansibleAnsibleNamespaceMetadata) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAvatar

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetAvatar() *os.File`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *PatchedansibleAnsibleNamespaceMetadata) GetAvatarOk() (**os.File, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *PatchedansibleAnsibleNamespaceMetadata) SetAvatar(v *os.File)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *PatchedansibleAnsibleNamespaceMetadata) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


