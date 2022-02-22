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

// CollectionOfDriveItems struct for CollectionOfDriveItems
type CollectionOfDriveItems struct {
	Value []DriveItem `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfDriveItems instantiates a new CollectionOfDriveItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDriveItems() *CollectionOfDriveItems {
	this := CollectionOfDriveItems{}
	return &this
}

// NewCollectionOfDriveItemsWithDefaults instantiates a new CollectionOfDriveItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDriveItemsWithDefaults() *CollectionOfDriveItems {
	this := CollectionOfDriveItems{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDriveItems) GetValue() []DriveItem {
	if o == nil || o.Value == nil {
		var ret []DriveItem
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDriveItems) GetValueOk() ([]DriveItem, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDriveItems) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []DriveItem and assigns it to the Value field.
func (o *CollectionOfDriveItems) SetValue(v []DriveItem) {
	o.Value = v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfDriveItems) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDriveItems) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfDriveItems) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfDriveItems) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfDriveItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDriveItems struct {
	value *CollectionOfDriveItems
	isSet bool
}

func (v NullableCollectionOfDriveItems) Get() *CollectionOfDriveItems {
	return v.value
}

func (v *NullableCollectionOfDriveItems) Set(val *CollectionOfDriveItems) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDriveItems) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDriveItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDriveItems(val *CollectionOfDriveItems) *NullableCollectionOfDriveItems {
	return &NullableCollectionOfDriveItems{value: val, isSet: true}
}

func (v NullableCollectionOfDriveItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDriveItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


