/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// MemberReference struct for MemberReference
type MemberReference struct {
	OdataId *string `json:"@odata.id,omitempty"`
}

// NewMemberReference instantiates a new MemberReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberReference() *MemberReference {
	this := MemberReference{}
	return &this
}

// NewMemberReferenceWithDefaults instantiates a new MemberReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberReferenceWithDefaults() *MemberReference {
	this := MemberReference{}
	return &this
}

// GetOdataId returns the OdataId field value if set, zero value otherwise.
func (o *MemberReference) GetOdataId() string {
	if o == nil || o.OdataId == nil {
		var ret string
		return ret
	}
	return *o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberReference) GetOdataIdOk() (*string, bool) {
	if o == nil || o.OdataId == nil {
		return nil, false
	}
	return o.OdataId, true
}

// HasOdataId returns a boolean if a field has been set.
func (o *MemberReference) HasOdataId() bool {
	if o != nil && o.OdataId != nil {
		return true
	}

	return false
}

// SetOdataId gets a reference to the given string and assigns it to the OdataId field.
func (o *MemberReference) SetOdataId(v string) {
	o.OdataId = &v
}

func (o MemberReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OdataId != nil {
		toSerialize["@odata.id"] = o.OdataId
	}
	return json.Marshal(toSerialize)
}

type NullableMemberReference struct {
	value *MemberReference
	isSet bool
}

func (v NullableMemberReference) Get() *MemberReference {
	return v.value
}

func (v *NullableMemberReference) Set(val *MemberReference) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberReference) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberReference(val *MemberReference) *NullableMemberReference {
	return &NullableMemberReference{value: val, isSet: true}
}

func (v NullableMemberReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


