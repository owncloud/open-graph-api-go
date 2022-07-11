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

// Quota Optional. Information about the drive's storage space quota. Read-only.
type Quota struct {
	// Total space consumed by files in the recycle bin, in bytes. Read-only.
	Deleted *int64 `json:"deleted,omitempty"`
	// Total space remaining before reaching the quota limit, in bytes. Read-only.
	Remaining *int64 `json:"remaining,omitempty"`
	// Enumeration value that indicates the state of the storage space. Either \"normal\", \"nearing\", \"critical\" or \"exceeded\". Read-only.
	State *string `json:"state,omitempty"`
	// Total allowed storage space, in bytes. Read-only.
	Total *int64 `json:"total,omitempty"`
	// Total space used, in bytes. Read-only.
	Used *int64 `json:"used,omitempty"`
}

// NewQuota instantiates a new Quota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuota() *Quota {
	this := Quota{}
	return &this
}

// NewQuotaWithDefaults instantiates a new Quota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaWithDefaults() *Quota {
	this := Quota{}
	return &this
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Quota) GetDeleted() int64 {
	if o == nil || o.Deleted == nil {
		var ret int64
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetDeletedOk() (*int64, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Quota) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given int64 and assigns it to the Deleted field.
func (o *Quota) SetDeleted(v int64) {
	o.Deleted = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *Quota) GetRemaining() int64 {
	if o == nil || o.Remaining == nil {
		var ret int64
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetRemainingOk() (*int64, bool) {
	if o == nil || o.Remaining == nil {
		return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *Quota) HasRemaining() bool {
	if o != nil && o.Remaining != nil {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int64 and assigns it to the Remaining field.
func (o *Quota) SetRemaining(v int64) {
	o.Remaining = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Quota) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Quota) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Quota) SetState(v string) {
	o.State = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Quota) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Quota) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *Quota) SetTotal(v int64) {
	o.Total = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *Quota) GetUsed() int64 {
	if o == nil || o.Used == nil {
		var ret int64
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetUsedOk() (*int64, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *Quota) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int64 and assigns it to the Used field.
func (o *Quota) SetUsed(v int64) {
	o.Used = &v
}

func (o Quota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Remaining != nil {
		toSerialize["remaining"] = o.Remaining
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Used != nil {
		toSerialize["used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableQuota struct {
	value *Quota
	isSet bool
}

func (v NullableQuota) Get() *Quota {
	return v.value
}

func (v *NullableQuota) Set(val *Quota) {
	v.value = val
	v.isSet = true
}

func (v NullableQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuota(val *Quota) *NullableQuota {
	return &NullableQuota{value: val, isSet: true}
}

func (v NullableQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


