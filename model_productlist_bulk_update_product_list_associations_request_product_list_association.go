/*
Collection Service

API for managing collection

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product-list

import (
	"encoding/json"
)

// checks if the ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation{}

// ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation struct for ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation
type ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation struct {
	Id *string `json:"id,omitempty"`
	Position *int32 `json:"position,omitempty"`
}

// NewProductlistBulkUpdateProductListAssociationsRequestProductListAssociation instantiates a new ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistBulkUpdateProductListAssociationsRequestProductListAssociation() *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation {
	this := ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation{}
	return &this
}

// NewProductlistBulkUpdateProductListAssociationsRequestProductListAssociationWithDefaults instantiates a new ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistBulkUpdateProductListAssociationsRequestProductListAssociationWithDefaults() *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation {
	this := ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) SetId(v string) {
	o.Id = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) SetPosition(v int32) {
	o.Position = &v
}

func (o ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	return toSerialize, nil
}

type NullableProductlistBulkUpdateProductListAssociationsRequestProductListAssociation struct {
	value *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation
	isSet bool
}

func (v NullableProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) Get() *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation {
	return v.value
}

func (v *NullableProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) Set(val *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistBulkUpdateProductListAssociationsRequestProductListAssociation(val *ProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) *NullableProductlistBulkUpdateProductListAssociationsRequestProductListAssociation {
	return &NullableProductlistBulkUpdateProductListAssociationsRequestProductListAssociation{value: val, isSet: true}
}

func (v NullableProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistBulkUpdateProductListAssociationsRequestProductListAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


