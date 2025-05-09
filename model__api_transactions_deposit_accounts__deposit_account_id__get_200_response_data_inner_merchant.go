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

// checks if the ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant{}

// ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant struct for ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant
type ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant struct {
	Name *string `json:"name,omitempty"`
	CategoryCode NullableString `json:"categoryCode,omitempty"`
	Category NullableString `json:"category,omitempty"`
	CountryCode NullableString `json:"countryCode,omitempty"`
	Address NullableString `json:"address,omitempty"`
	City NullableString `json:"city,omitempty"`
	State NullableString `json:"state,omitempty"`
	Country *string `json:"country,omitempty"`
	ZipCode NullableString `json:"zipCode,omitempty"`
	Website *string `json:"website,omitempty"`
	CustomerServicePageUrl *string `json:"customerServicePageUrl,omitempty"`
	FacebookPageUrl *string `json:"facebookPageUrl,omitempty"`
	TwitterPageUrl *string `json:"twitterPageUrl,omitempty"`
	MerchantImageRef *string `json:"merchantImageRef,omitempty"`
}

// NewApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant instantiates a new ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant() *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant {
	this := ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant{}
	return &this
}

// NewApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchantWithDefaults instantiates a new ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchantWithDefaults() *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant {
	this := ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetName(v string) {
	o.Name = &v
}

// GetCategoryCode returns the CategoryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCategoryCode() string {
	if o == nil || IsNil(o.CategoryCode.Get()) {
		var ret string
		return ret
	}
	return *o.CategoryCode.Get()
}

// GetCategoryCodeOk returns a tuple with the CategoryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCategoryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CategoryCode.Get(), o.CategoryCode.IsSet()
}

// HasCategoryCode returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasCategoryCode() bool {
	if o != nil && o.CategoryCode.IsSet() {
		return true
	}

	return false
}

// SetCategoryCode gets a reference to the given NullableString and assigns it to the CategoryCode field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetCategoryCode(v string) {
	o.CategoryCode.Set(&v)
}
// SetCategoryCodeNil sets the value for CategoryCode to be an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetCategoryCodeNil() {
	o.CategoryCode.Set(nil)
}

// UnsetCategoryCode ensures that no value is present for CategoryCode, not even an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) UnsetCategoryCode() {
	o.CategoryCode.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCategory() string {
	if o == nil || IsNil(o.Category.Get()) {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) UnsetCategory() {
	o.Category.Unset()
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}
// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) UnsetAddress() {
	o.Address.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) UnsetCity() {
	o.City.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetState() string {
	if o == nil || IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) UnsetState() {
	o.State.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetCountry(v string) {
	o.Country = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode.Get()) {
		var ret string
		return ret
	}
	return *o.ZipCode.Get()
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetZipCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZipCode.Get(), o.ZipCode.IsSet()
}

// HasZipCode returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasZipCode() bool {
	if o != nil && o.ZipCode.IsSet() {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given NullableString and assigns it to the ZipCode field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetZipCode(v string) {
	o.ZipCode.Set(&v)
}
// SetZipCodeNil sets the value for ZipCode to be an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetZipCodeNil() {
	o.ZipCode.Set(nil)
}

// UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) UnsetZipCode() {
	o.ZipCode.Unset()
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetWebsite(v string) {
	o.Website = &v
}

// GetCustomerServicePageUrl returns the CustomerServicePageUrl field value if set, zero value otherwise.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCustomerServicePageUrl() string {
	if o == nil || IsNil(o.CustomerServicePageUrl) {
		var ret string
		return ret
	}
	return *o.CustomerServicePageUrl
}

// GetCustomerServicePageUrlOk returns a tuple with the CustomerServicePageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetCustomerServicePageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerServicePageUrl) {
		return nil, false
	}
	return o.CustomerServicePageUrl, true
}

// HasCustomerServicePageUrl returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasCustomerServicePageUrl() bool {
	if o != nil && !IsNil(o.CustomerServicePageUrl) {
		return true
	}

	return false
}

// SetCustomerServicePageUrl gets a reference to the given string and assigns it to the CustomerServicePageUrl field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetCustomerServicePageUrl(v string) {
	o.CustomerServicePageUrl = &v
}

// GetFacebookPageUrl returns the FacebookPageUrl field value if set, zero value otherwise.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetFacebookPageUrl() string {
	if o == nil || IsNil(o.FacebookPageUrl) {
		var ret string
		return ret
	}
	return *o.FacebookPageUrl
}

// GetFacebookPageUrlOk returns a tuple with the FacebookPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetFacebookPageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FacebookPageUrl) {
		return nil, false
	}
	return o.FacebookPageUrl, true
}

// HasFacebookPageUrl returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasFacebookPageUrl() bool {
	if o != nil && !IsNil(o.FacebookPageUrl) {
		return true
	}

	return false
}

// SetFacebookPageUrl gets a reference to the given string and assigns it to the FacebookPageUrl field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetFacebookPageUrl(v string) {
	o.FacebookPageUrl = &v
}

// GetTwitterPageUrl returns the TwitterPageUrl field value if set, zero value otherwise.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetTwitterPageUrl() string {
	if o == nil || IsNil(o.TwitterPageUrl) {
		var ret string
		return ret
	}
	return *o.TwitterPageUrl
}

// GetTwitterPageUrlOk returns a tuple with the TwitterPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetTwitterPageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TwitterPageUrl) {
		return nil, false
	}
	return o.TwitterPageUrl, true
}

// HasTwitterPageUrl returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasTwitterPageUrl() bool {
	if o != nil && !IsNil(o.TwitterPageUrl) {
		return true
	}

	return false
}

// SetTwitterPageUrl gets a reference to the given string and assigns it to the TwitterPageUrl field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetTwitterPageUrl(v string) {
	o.TwitterPageUrl = &v
}

// GetMerchantImageRef returns the MerchantImageRef field value if set, zero value otherwise.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetMerchantImageRef() string {
	if o == nil || IsNil(o.MerchantImageRef) {
		var ret string
		return ret
	}
	return *o.MerchantImageRef
}

// GetMerchantImageRefOk returns a tuple with the MerchantImageRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) GetMerchantImageRefOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantImageRef) {
		return nil, false
	}
	return o.MerchantImageRef, true
}

// HasMerchantImageRef returns a boolean if a field has been set.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) HasMerchantImageRef() bool {
	if o != nil && !IsNil(o.MerchantImageRef) {
		return true
	}

	return false
}

// SetMerchantImageRef gets a reference to the given string and assigns it to the MerchantImageRef field.
func (o *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) SetMerchantImageRef(v string) {
	o.MerchantImageRef = &v
}

func (o ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.CategoryCode.IsSet() {
		toSerialize["categoryCode"] = o.CategoryCode.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.CountryCode.IsSet() {
		toSerialize["countryCode"] = o.CountryCode.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if o.ZipCode.IsSet() {
		toSerialize["zipCode"] = o.ZipCode.Get()
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.CustomerServicePageUrl) {
		toSerialize["customerServicePageUrl"] = o.CustomerServicePageUrl
	}
	if !IsNil(o.FacebookPageUrl) {
		toSerialize["facebookPageUrl"] = o.FacebookPageUrl
	}
	if !IsNil(o.TwitterPageUrl) {
		toSerialize["twitterPageUrl"] = o.TwitterPageUrl
	}
	if !IsNil(o.MerchantImageRef) {
		toSerialize["merchantImageRef"] = o.MerchantImageRef
	}
	return toSerialize, nil
}

type NullableApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant struct {
	value *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant
	isSet bool
}

func (v NullableApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) Get() *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant {
	return v.value
}

func (v *NullableApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) Set(val *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant(val *ApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) *NullableApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant {
	return &NullableApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant{value: val, isSet: true}
}

func (v NullableApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTransactionsDepositAccountsDepositAccountIdGet200ResponseDataInnerMerchant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


