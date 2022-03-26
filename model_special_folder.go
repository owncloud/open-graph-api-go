/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// SpecialFolder If the current item is also available as a special folder, this facet is returned. Read-only
type SpecialFolder struct {
	// The unique identifier for this item in the /drive/special collection
	Name *string `json:"name,omitempty"`
}

// NewSpecialFolder instantiates a new SpecialFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecialFolder() *SpecialFolder {
	this := SpecialFolder{}
	return &this
}

// NewSpecialFolderWithDefaults instantiates a new SpecialFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialFolderWithDefaults() *SpecialFolder {
	this := SpecialFolder{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SpecialFolder) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialFolder) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SpecialFolder) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SpecialFolder) SetName(v string) {
	o.Name = &v
}

func (o SpecialFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSpecialFolder struct {
	value *SpecialFolder
	isSet bool
}

func (v NullableSpecialFolder) Get() *SpecialFolder {
	return v.value
}

func (v *NullableSpecialFolder) Set(val *SpecialFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecialFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecialFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecialFolder(val *SpecialFolder) *NullableSpecialFolder {
	return &NullableSpecialFolder{value: val, isSet: true}
}

func (v NullableSpecialFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecialFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


