/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
	"time"
)

// Permission The Permission resource provides information about a sharing permission granted for a DriveItem resource.
type Permission struct {
	// An optional expiration date which limits the permission in time.
	ExpirationDateTime  *time.Time    `json:"expirationDateTime,omitempty"`
	GrantedToIdentities []IdentitySet `json:"grantedToIdentities,omitempty"`
	Roles               []string      `json:"roles,omitempty"`
}

// NewPermission instantiates a new Permission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermission() *Permission {
	this := Permission{}
	return &this
}

// NewPermissionWithDefaults instantiates a new Permission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionWithDefaults() *Permission {
	this := Permission{}
	return &this
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise.
func (o *Permission) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permission) GetExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDateTime == nil {
		return nil, false
	}
	return o.ExpirationDateTime, true
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *Permission) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime != nil {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.
func (o *Permission) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime = &v
}

// GetGrantedToIdentities returns the GrantedToIdentities field value if set, zero value otherwise.
func (o *Permission) GetGrantedToIdentities() []IdentitySet {
	if o == nil || o.GrantedToIdentities == nil {
		var ret []IdentitySet
		return ret
	}
	return o.GrantedToIdentities
}

// GetGrantedToIdentitiesOk returns a tuple with the GrantedToIdentities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permission) GetGrantedToIdentitiesOk() ([]IdentitySet, bool) {
	if o == nil || o.GrantedToIdentities == nil {
		return nil, false
	}
	return o.GrantedToIdentities, true
}

// HasGrantedToIdentities returns a boolean if a field has been set.
func (o *Permission) HasGrantedToIdentities() bool {
	if o != nil && o.GrantedToIdentities != nil {
		return true
	}

	return false
}

// SetGrantedToIdentities gets a reference to the given []IdentitySet and assigns it to the GrantedToIdentities field.
func (o *Permission) SetGrantedToIdentities(v []IdentitySet) {
	o.GrantedToIdentities = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Permission) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permission) GetRolesOk() ([]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Permission) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *Permission) SetRoles(v []string) {
	o.Roles = v
}

func (o Permission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpirationDateTime != nil {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime
	}
	if o.GrantedToIdentities != nil {
		toSerialize["grantedToIdentities"] = o.GrantedToIdentities
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullablePermission struct {
	value *Permission
	isSet bool
}

func (v NullablePermission) Get() *Permission {
	return v.value
}

func (v *NullablePermission) Set(val *Permission) {
	v.value = val
	v.isSet = true
}

func (v NullablePermission) IsSet() bool {
	return v.isSet
}

func (v *NullablePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermission(val *Permission) *NullablePermission {
	return &NullablePermission{value: val, isSet: true}
}

func (v NullablePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
