/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.15.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
	"time"
)

// DirectoryObject Represents a Directory object. Read-only.
type DirectoryObject struct {
	// The unique identifier for the object. 12345678-9abc-def0-1234-56789abcde. The value of the ID property is often, but not exclusively, in the form of a GUID. The value should be treated as an opaque identifier and not based in being a GUID. Null values are not allowed. Read-only.
	Id *string `json:"id,omitempty"`
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
}

// NewDirectoryObject instantiates a new DirectoryObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectoryObject() *DirectoryObject {
	this := DirectoryObject{}
	return &this
}

// NewDirectoryObjectWithDefaults instantiates a new DirectoryObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectoryObjectWithDefaults() *DirectoryObject {
	this := DirectoryObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DirectoryObject) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectoryObject) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DirectoryObject) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DirectoryObject) SetId(v string) {
	o.Id = &v
}

// GetDeletedDateTime returns the DeletedDateTime field value if set, zero value otherwise.
func (o *DirectoryObject) GetDeletedDateTime() time.Time {
	if o == nil || o.DeletedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedDateTime
}

// GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectoryObject) GetDeletedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.DeletedDateTime == nil {
		return nil, false
	}
	return o.DeletedDateTime, true
}

// HasDeletedDateTime returns a boolean if a field has been set.
func (o *DirectoryObject) HasDeletedDateTime() bool {
	if o != nil && o.DeletedDateTime != nil {
		return true
	}

	return false
}

// SetDeletedDateTime gets a reference to the given time.Time and assigns it to the DeletedDateTime field.
func (o *DirectoryObject) SetDeletedDateTime(v time.Time) {
	o.DeletedDateTime = &v
}

func (o DirectoryObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DeletedDateTime != nil {
		toSerialize["deletedDateTime"] = o.DeletedDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableDirectoryObject struct {
	value *DirectoryObject
	isSet bool
}

func (v NullableDirectoryObject) Get() *DirectoryObject {
	return v.value
}

func (v *NullableDirectoryObject) Set(val *DirectoryObject) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectoryObject) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectoryObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectoryObject(val *DirectoryObject) *NullableDirectoryObject {
	return &NullableDirectoryObject{value: val, isSet: true}
}

func (v NullableDirectoryObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectoryObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


