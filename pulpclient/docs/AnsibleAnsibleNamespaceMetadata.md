# AnsibleAnsibleNamespaceMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Required named, only accepts lowercase, numbers and underscores. | 
**Company** | Pointer to **string** | Optional namespace company owner. | [optional] 
**Email** | Pointer to **string** | Optional namespace contact email. | [optional] 
**Description** | Pointer to **string** | Optional short description. | [optional] 
**Resources** | Pointer to **string** | Optional resource page in markdown format. | [optional] 
**Links** | Pointer to [**[]NamespaceLink**](NamespaceLink.md) | Labeled related links. | [optional] 
**Avatar** | Pointer to ***os.File** | Optional avatar image for Namespace | [optional] 

## Methods

### NewAnsibleAnsibleNamespaceMetadata

`func NewAnsibleAnsibleNamespaceMetadata(name string, ) *AnsibleAnsibleNamespaceMetadata`

NewAnsibleAnsibleNamespaceMetadata instantiates a new AnsibleAnsibleNamespaceMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleAnsibleNamespaceMetadataWithDefaults

`func NewAnsibleAnsibleNamespaceMetadataWithDefaults() *AnsibleAnsibleNamespaceMetadata`

NewAnsibleAnsibleNamespaceMetadataWithDefaults instantiates a new AnsibleAnsibleNamespaceMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AnsibleAnsibleNamespaceMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnsibleAnsibleNamespaceMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnsibleAnsibleNamespaceMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetCompany

`func (o *AnsibleAnsibleNamespaceMetadata) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AnsibleAnsibleNamespaceMetadata) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AnsibleAnsibleNamespaceMetadata) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AnsibleAnsibleNamespaceMetadata) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEmail

`func (o *AnsibleAnsibleNamespaceMetadata) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AnsibleAnsibleNamespaceMetadata) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AnsibleAnsibleNamespaceMetadata) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AnsibleAnsibleNamespaceMetadata) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDescription

`func (o *AnsibleAnsibleNamespaceMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnsibleAnsibleNamespaceMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnsibleAnsibleNamespaceMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AnsibleAnsibleNamespaceMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResources

`func (o *AnsibleAnsibleNamespaceMetadata) GetResources() string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *AnsibleAnsibleNamespaceMetadata) GetResourcesOk() (*string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *AnsibleAnsibleNamespaceMetadata) SetResources(v string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *AnsibleAnsibleNamespaceMetadata) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetLinks

`func (o *AnsibleAnsibleNamespaceMetadata) GetLinks() []NamespaceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AnsibleAnsibleNamespaceMetadata) GetLinksOk() (*[]NamespaceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AnsibleAnsibleNamespaceMetadata) SetLinks(v []NamespaceLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AnsibleAnsibleNamespaceMetadata) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAvatar

`func (o *AnsibleAnsibleNamespaceMetadata) GetAvatar() *os.File`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *AnsibleAnsibleNamespaceMetadata) GetAvatarOk() (**os.File, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *AnsibleAnsibleNamespaceMetadata) SetAvatar(v *os.File)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *AnsibleAnsibleNamespaceMetadata) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


