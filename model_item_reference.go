/*
 * Open Graph API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opengraph

import (
	"encoding/json"
)

// ItemReference struct for ItemReference
type ItemReference struct {
	// Unique identifier of the drive instance that contains the item. Read-only.
	DriveId *string `json:"driveId,omitempty"`
	// Identifies the type of drive. See [drive][] resource for values. Read-only.
	DriveType *string `json:"driveType,omitempty"`
	// Unique identifier of the item in the drive. Read-only.
	Id *string `json:"id,omitempty"`
	// The name of the item being referenced. Read-only.
	Name *string `json:"name,omitempty"`
	// Path that can be used to navigate to the item. Read-only.
	Path *string `json:"path,omitempty"`
	// A unique identifier for a shared resource that can be accessed via the [Shares][] API.
	ShareId *string `json:"shareId,omitempty"`
}

// NewItemReference instantiates a new ItemReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemReference() *ItemReference {
	this := ItemReference{}
	return &this
}

// NewItemReferenceWithDefaults instantiates a new ItemReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemReferenceWithDefaults() *ItemReference {
	this := ItemReference{}
	return &this
}

// GetDriveId returns the DriveId field value if set, zero value otherwise.
func (o *ItemReference) GetDriveId() string {
	if o == nil || o.DriveId == nil {
		var ret string
		return ret
	}
	return *o.DriveId
}

// GetDriveIdOk returns a tuple with the DriveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemReference) GetDriveIdOk() (*string, bool) {
	if o == nil || o.DriveId == nil {
		return nil, false
	}
	return o.DriveId, true
}

// HasDriveId returns a boolean if a field has been set.
func (o *ItemReference) HasDriveId() bool {
	if o != nil && o.DriveId != nil {
		return true
	}

	return false
}

// SetDriveId gets a reference to the given string and assigns it to the DriveId field.
func (o *ItemReference) SetDriveId(v string) {
	o.DriveId = &v
}

// GetDriveType returns the DriveType field value if set, zero value otherwise.
func (o *ItemReference) GetDriveType() string {
	if o == nil || o.DriveType == nil {
		var ret string
		return ret
	}
	return *o.DriveType
}

// GetDriveTypeOk returns a tuple with the DriveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemReference) GetDriveTypeOk() (*string, bool) {
	if o == nil || o.DriveType == nil {
		return nil, false
	}
	return o.DriveType, true
}

// HasDriveType returns a boolean if a field has been set.
func (o *ItemReference) HasDriveType() bool {
	if o != nil && o.DriveType != nil {
		return true
	}

	return false
}

// SetDriveType gets a reference to the given string and assigns it to the DriveType field.
func (o *ItemReference) SetDriveType(v string) {
	o.DriveType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ItemReference) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemReference) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ItemReference) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ItemReference) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ItemReference) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemReference) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ItemReference) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ItemReference) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ItemReference) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemReference) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ItemReference) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ItemReference) SetPath(v string) {
	o.Path = &v
}

// GetShareId returns the ShareId field value if set, zero value otherwise.
func (o *ItemReference) GetShareId() string {
	if o == nil || o.ShareId == nil {
		var ret string
		return ret
	}
	return *o.ShareId
}

// GetShareIdOk returns a tuple with the ShareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemReference) GetShareIdOk() (*string, bool) {
	if o == nil || o.ShareId == nil {
		return nil, false
	}
	return o.ShareId, true
}

// HasShareId returns a boolean if a field has been set.
func (o *ItemReference) HasShareId() bool {
	if o != nil && o.ShareId != nil {
		return true
	}

	return false
}

// SetShareId gets a reference to the given string and assigns it to the ShareId field.
func (o *ItemReference) SetShareId(v string) {
	o.ShareId = &v
}

func (o ItemReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DriveId != nil {
		toSerialize["driveId"] = o.DriveId
	}
	if o.DriveType != nil {
		toSerialize["driveType"] = o.DriveType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.ShareId != nil {
		toSerialize["shareId"] = o.ShareId
	}
	return json.Marshal(toSerialize)
}

type NullableItemReference struct {
	value *ItemReference
	isSet bool
}

func (v NullableItemReference) Get() *ItemReference {
	return v.value
}

func (v *NullableItemReference) Set(val *ItemReference) {
	v.value = val
	v.isSet = true
}

func (v NullableItemReference) IsSet() bool {
	return v.isSet
}

func (v *NullableItemReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemReference(val *ItemReference) *NullableItemReference {
	return &NullableItemReference{value: val, isSet: true}
}

func (v NullableItemReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


