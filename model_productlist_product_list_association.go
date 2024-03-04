/*
product-list/bulk_update_product_list_associations.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product-list

import (
	"encoding/json"
)

// checks if the ProductlistProductListAssociation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistProductListAssociation{}

// ProductlistProductListAssociation struct for ProductlistProductListAssociation
type ProductlistProductListAssociation struct {
	Id *string `json:"id,omitempty"`
	ListId *string `json:"listId,omitempty"`
	Position *int32 `json:"position,omitempty"`
	ProductGrn *string `json:"productGrn,omitempty"`
}

// NewProductlistProductListAssociation instantiates a new ProductlistProductListAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistProductListAssociation() *ProductlistProductListAssociation {
	this := ProductlistProductListAssociation{}
	return &this
}

// NewProductlistProductListAssociationWithDefaults instantiates a new ProductlistProductListAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistProductListAssociationWithDefaults() *ProductlistProductListAssociation {
	this := ProductlistProductListAssociation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductlistProductListAssociation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListAssociation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductlistProductListAssociation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductlistProductListAssociation) SetId(v string) {
	o.Id = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *ProductlistProductListAssociation) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListAssociation) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *ProductlistProductListAssociation) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *ProductlistProductListAssociation) SetListId(v string) {
	o.ListId = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductlistProductListAssociation) GetPosition() int32 {
	if o == nil || IsNil(o.Position) {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListAssociation) GetPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductlistProductListAssociation) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *ProductlistProductListAssociation) SetPosition(v int32) {
	o.Position = &v
}

// GetProductGrn returns the ProductGrn field value if set, zero value otherwise.
func (o *ProductlistProductListAssociation) GetProductGrn() string {
	if o == nil || IsNil(o.ProductGrn) {
		var ret string
		return ret
	}
	return *o.ProductGrn
}

// GetProductGrnOk returns a tuple with the ProductGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistProductListAssociation) GetProductGrnOk() (*string, bool) {
	if o == nil || IsNil(o.ProductGrn) {
		return nil, false
	}
	return o.ProductGrn, true
}

// HasProductGrn returns a boolean if a field has been set.
func (o *ProductlistProductListAssociation) HasProductGrn() bool {
	if o != nil && !IsNil(o.ProductGrn) {
		return true
	}

	return false
}

// SetProductGrn gets a reference to the given string and assigns it to the ProductGrn field.
func (o *ProductlistProductListAssociation) SetProductGrn(v string) {
	o.ProductGrn = &v
}

func (o ProductlistProductListAssociation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistProductListAssociation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ListId) {
		toSerialize["listId"] = o.ListId
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.ProductGrn) {
		toSerialize["productGrn"] = o.ProductGrn
	}
	return toSerialize, nil
}

type NullableProductlistProductListAssociation struct {
	value *ProductlistProductListAssociation
	isSet bool
}

func (v NullableProductlistProductListAssociation) Get() *ProductlistProductListAssociation {
	return v.value
}

func (v *NullableProductlistProductListAssociation) Set(val *ProductlistProductListAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistProductListAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistProductListAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistProductListAssociation(val *ProductlistProductListAssociation) *NullableProductlistProductListAssociation {
	return &NullableProductlistProductListAssociation{value: val, isSet: true}
}

func (v NullableProductlistProductListAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistProductListAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


