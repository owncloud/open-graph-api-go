/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AppRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRole{}

// AppRole struct for AppRole
type AppRole struct {
	// Specifies whether this app role can be assigned to users and groups (by setting to ['User']), to other application's (by setting to ['Application'], or both (by setting to ['User', 'Application']). App roles supporting assignment to other applications' service principals are also known as application permissions. The 'Application' value is only supported for app roles defined on application entities.
	AllowedMemberTypes []string `json:"allowedMemberTypes,omitempty"`
	// The description for the app role. This is displayed when the app role is being assigned and, if the app role functions as an application permission, during  consent experiences.
	Description NullableString `json:"description,omitempty"`
	// Display name for the permission that appears in the app role assignment and consent experiences.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Unique role identifier inside the appRoles collection. When creating a new app role, a new GUID identifier must be provided.
	Id string `json:"id" validate:"regexp=^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"`
}

type _AppRole AppRole

// NewAppRole instantiates a new AppRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRole(id string) *AppRole {
	this := AppRole{}
	this.Id = id
	return &this
}

// NewAppRoleWithDefaults instantiates a new AppRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWithDefaults() *AppRole {
	this := AppRole{}
	return &this
}

// GetAllowedMemberTypes returns the AllowedMemberTypes field value if set, zero value otherwise.
func (o *AppRole) GetAllowedMemberTypes() []string {
	if o == nil || IsNil(o.AllowedMemberTypes) {
		var ret []string
		return ret
	}
	return o.AllowedMemberTypes
}

// GetAllowedMemberTypesOk returns a tuple with the AllowedMemberTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRole) GetAllowedMemberTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedMemberTypes) {
		return nil, false
	}
	return o.AllowedMemberTypes, true
}

// HasAllowedMemberTypes returns a boolean if a field has been set.
func (o *AppRole) HasAllowedMemberTypes() bool {
	if o != nil && !IsNil(o.AllowedMemberTypes) {
		return true
	}

	return false
}

// SetAllowedMemberTypes gets a reference to the given []string and assigns it to the AllowedMemberTypes field.
func (o *AppRole) SetAllowedMemberTypes(v []string) {
	o.AllowedMemberTypes = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppRole) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppRole) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AppRole) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AppRole) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AppRole) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AppRole) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppRole) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppRole) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AppRole) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *AppRole) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}

// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *AppRole) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *AppRole) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetId returns the Id field value
func (o *AppRole) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppRole) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppRole) SetId(v string) {
	o.Id = v
}

func (o AppRole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedMemberTypes) {
		toSerialize["allowedMemberTypes"] = o.AllowedMemberTypes
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *AppRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAppRole := _AppRole{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppRole)

	if err != nil {
		return err
	}

	*o = AppRole(varAppRole)

	return err
}

type NullableAppRole struct {
	value *AppRole
	isSet bool
}

func (v NullableAppRole) Get() *AppRole {
	return v.value
}

func (v *NullableAppRole) Set(val *AppRole) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRole) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRole(val *AppRole) *NullableAppRole {
	return &NullableAppRole{value: val, isSet: true}
}

func (v NullableAppRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
