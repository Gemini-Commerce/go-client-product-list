/*
Collection Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productlist

import (
	"encoding/json"
)

// checks if the ProductlistOrderBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistOrderBy{}

// ProductlistOrderBy struct for ProductlistOrderBy
type ProductlistOrderBy struct {
	Field                *string           `json:"field,omitempty"`
	Direction            *OrderByDirection `json:"direction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistOrderBy ProductlistOrderBy

// NewProductlistOrderBy instantiates a new ProductlistOrderBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistOrderBy() *ProductlistOrderBy {
	this := ProductlistOrderBy{}
	var direction OrderByDirection = ORDERBYDIRECTION_DEFAULT
	this.Direction = &direction
	return &this
}

// NewProductlistOrderByWithDefaults instantiates a new ProductlistOrderBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistOrderByWithDefaults() *ProductlistOrderBy {
	this := ProductlistOrderBy{}
	var direction OrderByDirection = ORDERBYDIRECTION_DEFAULT
	this.Direction = &direction
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ProductlistOrderBy) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistOrderBy) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ProductlistOrderBy) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ProductlistOrderBy) SetField(v string) {
	o.Field = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *ProductlistOrderBy) GetDirection() OrderByDirection {
	if o == nil || IsNil(o.Direction) {
		var ret OrderByDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistOrderBy) GetDirectionOk() (*OrderByDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *ProductlistOrderBy) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given OrderByDirection and assigns it to the Direction field.
func (o *ProductlistOrderBy) SetDirection(v OrderByDirection) {
	o.Direction = &v
}

func (o ProductlistOrderBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistOrderBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistOrderBy) UnmarshalJSON(data []byte) (err error) {
	varProductlistOrderBy := _ProductlistOrderBy{}

	err = json.Unmarshal(data, &varProductlistOrderBy)

	if err != nil {
		return err
	}

	*o = ProductlistOrderBy(varProductlistOrderBy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "field")
		delete(additionalProperties, "direction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistOrderBy) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductlistOrderBy) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductlistOrderBy struct {
	value *ProductlistOrderBy
	isSet bool
}

func (v NullableProductlistOrderBy) Get() *ProductlistOrderBy {
	return v.value
}

func (v *NullableProductlistOrderBy) Set(val *ProductlistOrderBy) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistOrderBy) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistOrderBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistOrderBy(val *ProductlistOrderBy) *NullableProductlistOrderBy {
	return &NullableProductlistOrderBy{value: val, isSet: true}
}

func (v NullableProductlistOrderBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistOrderBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
