/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.15.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// CollectionOfDrives struct for CollectionOfDrives
type CollectionOfDrives struct {
	Value []Drive `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfDrives instantiates a new CollectionOfDrives object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDrives() *CollectionOfDrives {
	this := CollectionOfDrives{}
	return &this
}

// NewCollectionOfDrivesWithDefaults instantiates a new CollectionOfDrives object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDrivesWithDefaults() *CollectionOfDrives {
	this := CollectionOfDrives{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDrives) GetValue() []Drive {
	if o == nil || o.Value == nil {
		var ret []Drive
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDrives) GetValueOk() ([]Drive, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDrives) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []Drive and assigns it to the Value field.
func (o *CollectionOfDrives) SetValue(v []Drive) {
	o.Value = v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfDrives) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDrives) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfDrives) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfDrives) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfDrives) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDrives struct {
	value *CollectionOfDrives
	isSet bool
}

func (v NullableCollectionOfDrives) Get() *CollectionOfDrives {
	return v.value
}

func (v *NullableCollectionOfDrives) Set(val *CollectionOfDrives) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDrives) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDrives) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDrives(val *CollectionOfDrives) *NullableCollectionOfDrives {
	return &NullableCollectionOfDrives{value: val, isSet: true}
}

func (v NullableCollectionOfDrives) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDrives) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


