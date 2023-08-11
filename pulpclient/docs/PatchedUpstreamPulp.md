# PatchedUpstreamPulp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name for this Pulp server. | [optional] 
**BaseUrl** | Pointer to **string** | The transport, hostname, and an optional port of the Pulp server. e.g. https://example.com | [optional] 
**ApiRoot** | Pointer to **string** | The API root. Defaults to &#39;/pulp/&#39;. | [optional] 
**Domain** | Pointer to **NullableString** | The domain of the Pulp server if enabled. | [optional] 
**CaCert** | Pointer to **NullableString** | A PEM encoded CA certificate used to validate the server certificate presented by the remote server. | [optional] 
**ClientCert** | Pointer to **NullableString** | A PEM encoded client certificate used for authentication. | [optional] 
**ClientKey** | Pointer to **NullableString** | A PEM encoded private key used for authentication. | [optional] 
**TlsValidation** | Pointer to **bool** | If True, TLS peer validation must be performed. | [optional] 
**Username** | Pointer to **NullableString** | The username to be used for authentication when syncing. | [optional] 
**Password** | Pointer to **NullableString** | The password to be used for authentication when syncing. Extra leading and trailing whitespace characters are not trimmed. | [optional] 
**PulpLabelSelect** | Pointer to **NullableString** | One or more comma separated labels that will be used to filter distributions on the upstream Pulp. E.g. \&quot;foo&#x3D;bar,key&#x3D;val\&quot; or \&quot;foo,key\&quot; | [optional] 

## Methods

### NewPatchedUpstreamPulp

`func NewPatchedUpstreamPulp() *PatchedUpstreamPulp`

NewPatchedUpstreamPulp instantiates a new PatchedUpstreamPulp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUpstreamPulpWithDefaults

`func NewPatchedUpstreamPulpWithDefaults() *PatchedUpstreamPulp`

NewPatchedUpstreamPulpWithDefaults instantiates a new PatchedUpstreamPulp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedUpstreamPulp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedUpstreamPulp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedUpstreamPulp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedUpstreamPulp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBaseUrl

`func (o *PatchedUpstreamPulp) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *PatchedUpstreamPulp) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *PatchedUpstreamPulp) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *PatchedUpstreamPulp) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### GetApiRoot

`func (o *PatchedUpstreamPulp) GetApiRoot() string`

GetApiRoot returns the ApiRoot field if non-nil, zero value otherwise.

### GetApiRootOk

`func (o *PatchedUpstreamPulp) GetApiRootOk() (*string, bool)`

GetApiRootOk returns a tuple with the ApiRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRoot

`func (o *PatchedUpstreamPulp) SetApiRoot(v string)`

SetApiRoot sets ApiRoot field to given value.

### HasApiRoot

`func (o *PatchedUpstreamPulp) HasApiRoot() bool`

HasApiRoot returns a boolean if a field has been set.

### GetDomain

`func (o *PatchedUpstreamPulp) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PatchedUpstreamPulp) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PatchedUpstreamPulp) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PatchedUpstreamPulp) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *PatchedUpstreamPulp) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *PatchedUpstreamPulp) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetCaCert

`func (o *PatchedUpstreamPulp) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *PatchedUpstreamPulp) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *PatchedUpstreamPulp) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *PatchedUpstreamPulp) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### SetCaCertNil

`func (o *PatchedUpstreamPulp) SetCaCertNil(b bool)`

 SetCaCertNil sets the value for CaCert to be an explicit nil

### UnsetCaCert
`func (o *PatchedUpstreamPulp) UnsetCaCert()`

UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
### GetClientCert

`func (o *PatchedUpstreamPulp) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *PatchedUpstreamPulp) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *PatchedUpstreamPulp) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *PatchedUpstreamPulp) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### SetClientCertNil

`func (o *PatchedUpstreamPulp) SetClientCertNil(b bool)`

 SetClientCertNil sets the value for ClientCert to be an explicit nil

### UnsetClientCert
`func (o *PatchedUpstreamPulp) UnsetClientCert()`

UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
### GetClientKey

`func (o *PatchedUpstreamPulp) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *PatchedUpstreamPulp) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *PatchedUpstreamPulp) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.

### HasClientKey

`func (o *PatchedUpstreamPulp) HasClientKey() bool`

HasClientKey returns a boolean if a field has been set.

### SetClientKeyNil

`func (o *PatchedUpstreamPulp) SetClientKeyNil(b bool)`

 SetClientKeyNil sets the value for ClientKey to be an explicit nil

### UnsetClientKey
`func (o *PatchedUpstreamPulp) UnsetClientKey()`

UnsetClientKey ensures that no value is present for ClientKey, not even an explicit nil
### GetTlsValidation

`func (o *PatchedUpstreamPulp) GetTlsValidation() bool`

GetTlsValidation returns the TlsValidation field if non-nil, zero value otherwise.

### GetTlsValidationOk

`func (o *PatchedUpstreamPulp) GetTlsValidationOk() (*bool, bool)`

GetTlsValidationOk returns a tuple with the TlsValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsValidation

`func (o *PatchedUpstreamPulp) SetTlsValidation(v bool)`

SetTlsValidation sets TlsValidation field to given value.

### HasTlsValidation

`func (o *PatchedUpstreamPulp) HasTlsValidation() bool`

HasTlsValidation returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedUpstreamPulp) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedUpstreamPulp) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedUpstreamPulp) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedUpstreamPulp) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *PatchedUpstreamPulp) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *PatchedUpstreamPulp) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *PatchedUpstreamPulp) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedUpstreamPulp) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedUpstreamPulp) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedUpstreamPulp) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PatchedUpstreamPulp) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PatchedUpstreamPulp) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPulpLabelSelect

`func (o *PatchedUpstreamPulp) GetPulpLabelSelect() string`

GetPulpLabelSelect returns the PulpLabelSelect field if non-nil, zero value otherwise.

### GetPulpLabelSelectOk

`func (o *PatchedUpstreamPulp) GetPulpLabelSelectOk() (*string, bool)`

GetPulpLabelSelectOk returns a tuple with the PulpLabelSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPulpLabelSelect

`func (o *PatchedUpstreamPulp) SetPulpLabelSelect(v string)`

SetPulpLabelSelect sets PulpLabelSelect field to given value.

### HasPulpLabelSelect

`func (o *PatchedUpstreamPulp) HasPulpLabelSelect() bool`

HasPulpLabelSelect returns a boolean if a field has been set.

### SetPulpLabelSelectNil

`func (o *PatchedUpstreamPulp) SetPulpLabelSelectNil(b bool)`

 SetPulpLabelSelectNil sets the value for PulpLabelSelect to be an explicit nil

### UnsetPulpLabelSelect
`func (o *PatchedUpstreamPulp) UnsetPulpLabelSelect()`

UnsetPulpLabelSelect ensures that no value is present for PulpLabelSelect, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


