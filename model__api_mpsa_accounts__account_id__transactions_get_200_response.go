/*
secure.scotiabank.com.har Mitmproxy2Swagger

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiMpsaAccountsAccountIdTransactionsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMpsaAccountsAccountIdTransactionsGet200Response{}

// ApiMpsaAccountsAccountIdTransactionsGet200Response struct for ApiMpsaAccountsAccountIdTransactionsGet200Response
type ApiMpsaAccountsAccountIdTransactionsGet200Response struct {
	Data []ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner `json:"data,omitempty"`
	Notifications map[string]interface{} `json:"notifications,omitempty"`
}

// NewApiMpsaAccountsAccountIdTransactionsGet200Response instantiates a new ApiMpsaAccountsAccountIdTransactionsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMpsaAccountsAccountIdTransactionsGet200Response() *ApiMpsaAccountsAccountIdTransactionsGet200Response {
	this := ApiMpsaAccountsAccountIdTransactionsGet200Response{}
	return &this
}

// NewApiMpsaAccountsAccountIdTransactionsGet200ResponseWithDefaults instantiates a new ApiMpsaAccountsAccountIdTransactionsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMpsaAccountsAccountIdTransactionsGet200ResponseWithDefaults() *ApiMpsaAccountsAccountIdTransactionsGet200Response {
	this := ApiMpsaAccountsAccountIdTransactionsGet200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApiMpsaAccountsAccountIdTransactionsGet200Response) GetData() []ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMpsaAccountsAccountIdTransactionsGet200Response) GetDataOk() ([]ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApiMpsaAccountsAccountIdTransactionsGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner and assigns it to the Data field.
func (o *ApiMpsaAccountsAccountIdTransactionsGet200Response) SetData(v []ApiMpsaAccountsAccountIdTransactionsGet200ResponseDataInner) {
	o.Data = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiMpsaAccountsAccountIdTransactionsGet200Response) GetNotifications() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiMpsaAccountsAccountIdTransactionsGet200Response) GetNotificationsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Notifications) {
		return map[string]interface{}{}, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *ApiMpsaAccountsAccountIdTransactionsGet200Response) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given map[string]interface{} and assigns it to the Notifications field.
func (o *ApiMpsaAccountsAccountIdTransactionsGet200Response) SetNotifications(v map[string]interface{}) {
	o.Notifications = v
}

func (o ApiMpsaAccountsAccountIdTransactionsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMpsaAccountsAccountIdTransactionsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	return toSerialize, nil
}

type NullableApiMpsaAccountsAccountIdTransactionsGet200Response struct {
	value *ApiMpsaAccountsAccountIdTransactionsGet200Response
	isSet bool
}

func (v NullableApiMpsaAccountsAccountIdTransactionsGet200Response) Get() *ApiMpsaAccountsAccountIdTransactionsGet200Response {
	return v.value
}

func (v *NullableApiMpsaAccountsAccountIdTransactionsGet200Response) Set(val *ApiMpsaAccountsAccountIdTransactionsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMpsaAccountsAccountIdTransactionsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMpsaAccountsAccountIdTransactionsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMpsaAccountsAccountIdTransactionsGet200Response(val *ApiMpsaAccountsAccountIdTransactionsGet200Response) *NullableApiMpsaAccountsAccountIdTransactionsGet200Response {
	return &NullableApiMpsaAccountsAccountIdTransactionsGet200Response{value: val, isSet: true}
}

func (v NullableApiMpsaAccountsAccountIdTransactionsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMpsaAccountsAccountIdTransactionsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


