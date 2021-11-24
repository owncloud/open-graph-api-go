/*
Open Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opengraph

import (
	"encoding/json"
)

// CollectionOfDrive struct for CollectionOfDrive
type CollectionOfDrive struct {
	Value *[]Drive `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfDrive instantiates a new CollectionOfDrive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDrive() *CollectionOfDrive {
	this := CollectionOfDrive{}
	return &this
}

// NewCollectionOfDriveWithDefaults instantiates a new CollectionOfDrive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDriveWithDefaults() *CollectionOfDrive {
	this := CollectionOfDrive{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDrive) GetValue() []Drive {
	if o == nil || o.Value == nil {
		var ret []Drive
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDrive) GetValueOk() (*[]Drive, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDrive) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []Drive and assigns it to the Value field.
func (o *CollectionOfDrive) SetValue(v []Drive) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfDrive) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDrive) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfDrive) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfDrive) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfDrive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDrive struct {
	value *CollectionOfDrive
	isSet bool
}

func (v NullableCollectionOfDrive) Get() *CollectionOfDrive {
	return v.value
}

func (v *NullableCollectionOfDrive) Set(val *CollectionOfDrive) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDrive) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDrive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDrive(val *CollectionOfDrive) *NullableCollectionOfDrive {
	return &NullableCollectionOfDrive{value: val, isSet: true}
}

func (v NullableCollectionOfDrive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDrive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


