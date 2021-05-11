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
	"time"
)

// BaseItem struct for BaseItem
type BaseItem struct {
	Entity
	// Identity of the user, device, or application which created the item. Read-only.
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Provides a user-visible description of the item. Optional.
	Description *string `json:"description,omitempty"`
	// ETag for the item. Read-only.
	ETag *string `json:"eTag,omitempty"`
	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The name of the item. Read-write.
	Name *string `json:"name,omitempty"`
	// Parent information, if the item has a parent. Read-write.
	ParentReference *ItemReference `json:"parentReference,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl *string `json:"webUrl,omitempty"`
	// Identity of the user who created the item. Read-only.
	CreatedByUser *User `json:"createdByUser,omitempty"`
	// Identity of the user who last modified the item. Read-only.
	LastModifiedByUser *User `json:"lastModifiedByUser,omitempty"`
}

// NewBaseItem instantiates a new BaseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseItem() *BaseItem {
	this := BaseItem{}
	return &this
}

// NewBaseItemWithDefaults instantiates a new BaseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseItemWithDefaults() *BaseItem {
	this := BaseItem{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BaseItem) GetCreatedBy() IdentitySet {
	if o == nil || o.CreatedBy == nil {
		var ret IdentitySet
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetCreatedByOk() (*IdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BaseItem) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given IdentitySet and assigns it to the CreatedBy field.
func (o *BaseItem) SetCreatedBy(v IdentitySet) {
	o.CreatedBy = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *BaseItem) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *BaseItem) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *BaseItem) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseItem) SetDescription(v string) {
	o.Description = &v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *BaseItem) GetETag() string {
	if o == nil || o.ETag == nil {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetETagOk() (*string, bool) {
	if o == nil || o.ETag == nil {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *BaseItem) HasETag() bool {
	if o != nil && o.ETag != nil {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *BaseItem) SetETag(v string) {
	o.ETag = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *BaseItem) GetLastModifiedBy() IdentitySet {
	if o == nil || o.LastModifiedBy == nil {
		var ret IdentitySet
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetLastModifiedByOk() (*IdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *BaseItem) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given IdentitySet and assigns it to the LastModifiedBy field.
func (o *BaseItem) SetLastModifiedBy(v IdentitySet) {
	o.LastModifiedBy = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *BaseItem) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *BaseItem) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *BaseItem) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseItem) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseItem) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseItem) SetName(v string) {
	o.Name = &v
}

// GetParentReference returns the ParentReference field value if set, zero value otherwise.
func (o *BaseItem) GetParentReference() ItemReference {
	if o == nil || o.ParentReference == nil {
		var ret ItemReference
		return ret
	}
	return *o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetParentReferenceOk() (*ItemReference, bool) {
	if o == nil || o.ParentReference == nil {
		return nil, false
	}
	return o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *BaseItem) HasParentReference() bool {
	if o != nil && o.ParentReference != nil {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given ItemReference and assigns it to the ParentReference field.
func (o *BaseItem) SetParentReference(v ItemReference) {
	o.ParentReference = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *BaseItem) GetWebUrl() string {
	if o == nil || o.WebUrl == nil {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetWebUrlOk() (*string, bool) {
	if o == nil || o.WebUrl == nil {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *BaseItem) HasWebUrl() bool {
	if o != nil && o.WebUrl != nil {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *BaseItem) SetWebUrl(v string) {
	o.WebUrl = &v
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *BaseItem) GetCreatedByUser() User {
	if o == nil || o.CreatedByUser == nil {
		var ret User
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetCreatedByUserOk() (*User, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *BaseItem) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given User and assigns it to the CreatedByUser field.
func (o *BaseItem) SetCreatedByUser(v User) {
	o.CreatedByUser = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *BaseItem) GetLastModifiedByUser() User {
	if o == nil || o.LastModifiedByUser == nil {
		var ret User
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetLastModifiedByUserOk() (*User, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *BaseItem) HasLastModifiedByUser() bool {
	if o != nil && o.LastModifiedByUser != nil {
		return true
	}

	return false
}

// SetLastModifiedByUser gets a reference to the given User and assigns it to the LastModifiedByUser field.
func (o *BaseItem) SetLastModifiedByUser(v User) {
	o.LastModifiedByUser = &v
}

func (o BaseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEntity, errEntity := json.Marshal(o.Entity)
	if errEntity != nil {
		return []byte{}, errEntity
	}
	errEntity = json.Unmarshal([]byte(serializedEntity), &toSerialize)
	if errEntity != nil {
		return []byte{}, errEntity
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ETag != nil {
		toSerialize["eTag"] = o.ETag
	}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ParentReference != nil {
		toSerialize["parentReference"] = o.ParentReference
	}
	if o.WebUrl != nil {
		toSerialize["webUrl"] = o.WebUrl
	}
	if o.CreatedByUser != nil {
		toSerialize["createdByUser"] = o.CreatedByUser
	}
	if o.LastModifiedByUser != nil {
		toSerialize["lastModifiedByUser"] = o.LastModifiedByUser
	}
	return json.Marshal(toSerialize)
}

type NullableBaseItem struct {
	value *BaseItem
	isSet bool
}

func (v NullableBaseItem) Get() *BaseItem {
	return v.value
}

func (v *NullableBaseItem) Set(val *BaseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseItem(val *BaseItem) *NullableBaseItem {
	return &NullableBaseItem{value: val, isSet: true}
}

func (v NullableBaseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


