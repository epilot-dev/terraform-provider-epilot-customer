// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/epilot-dev/terraform-provider-epilot-customer/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *ContactDataSourceModel) RefreshFromSharedContact(resp *shared.Contact) {
	if resp.ACL.AdditionalProperties == nil {
		r.ACL.AdditionalProperties = types.StringNull()
	} else {
		additionalPropertiesResult, _ := json.Marshal(resp.ACL.AdditionalProperties)
		r.ACL.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
	}
	r.ACL.Delete = nil
	for _, v := range resp.ACL.Delete {
		r.ACL.Delete = append(r.ACL.Delete, types.StringValue(v))
	}
	r.ACL.Edit = nil
	for _, v := range resp.ACL.Edit {
		r.ACL.Edit = append(r.ACL.Edit, types.StringValue(v))
	}
	r.ACL.View = nil
	for _, v := range resp.ACL.View {
		r.ACL.View = append(r.ACL.View, types.StringValue(v))
	}
	r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
	r.Org = types.StringValue(resp.Org)
	if len(r.Owners) > len(resp.Owners) {
		r.Owners = r.Owners[:len(resp.Owners)]
	}
	for ownersCount, ownersItem := range resp.Owners {
		var owners1 BaseEntityOwner
		owners1.OrgID = types.StringValue(ownersItem.OrgID)
		owners1.UserID = types.StringPointerValue(ownersItem.UserID)
		if ownersCount+1 > len(r.Owners) {
			r.Owners = append(r.Owners, owners1)
		} else {
			r.Owners[ownersCount].OrgID = owners1.OrgID
			r.Owners[ownersCount].UserID = owners1.UserID
		}
	}
	r.Schema = types.StringValue(resp.Schema)
	r.Tags = nil
	for _, v := range resp.Tags {
		r.Tags = append(r.Tags, types.StringValue(v))
	}
	r.Title1 = types.StringValue(resp.Title1)
	r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	if resp.Account == nil {
		r.Account = nil
	} else {
		r.Account = &BaseRelation{}
		if len(r.Account.DollarRelation) > len(resp.Account.DollarRelation) {
			r.Account.DollarRelation = r.Account.DollarRelation[:len(resp.Account.DollarRelation)]
		}
		for dollarRelationCount, dollarRelationItem := range resp.Account.DollarRelation {
			var dollarRelation1 DollarRelation
			dollarRelation1.Tags = nil
			for _, v := range dollarRelationItem.Tags {
				dollarRelation1.Tags = append(dollarRelation1.Tags, types.StringValue(v))
			}
			dollarRelation1.EntityID = types.StringValue(dollarRelationItem.EntityID)
			if dollarRelationCount+1 > len(r.Account.DollarRelation) {
				r.Account.DollarRelation = append(r.Account.DollarRelation, dollarRelation1)
			} else {
				r.Account.DollarRelation[dollarRelationCount].Tags = dollarRelation1.Tags
				r.Account.DollarRelation[dollarRelationCount].EntityID = dollarRelation1.EntityID
			}
		}
		if resp.Account.AdditionalProperties == nil {
			r.Account.AdditionalProperties = types.StringNull()
		} else {
			additionalPropertiesResult1, _ := json.Marshal(resp.Account.AdditionalProperties)
			r.Account.AdditionalProperties = types.StringValue(string(additionalPropertiesResult1))
		}
	}
	if len(r.Address) > len(resp.Address) {
		r.Address = r.Address[:len(resp.Address)]
	}
	for addressCount, addressItem := range resp.Address {
		var address1 BaseAddress
		address1.ID = types.StringPointerValue(addressItem.ID)
		address1.Tags = nil
		for _, v := range addressItem.Tags {
			address1.Tags = append(address1.Tags, types.StringValue(v))
		}
		address1.AdditionalInfo = types.StringPointerValue(addressItem.AdditionalInfo)
		address1.City = types.StringPointerValue(addressItem.City)
		if addressItem.Country != nil {
			address1.Country = types.StringValue(string(*addressItem.Country))
		} else {
			address1.Country = types.StringNull()
		}
		address1.PostalCode = types.StringPointerValue(addressItem.PostalCode)
		address1.Street = types.StringPointerValue(addressItem.Street)
		address1.StreetNumber = types.StringPointerValue(addressItem.StreetNumber)
		if addressCount+1 > len(r.Address) {
			r.Address = append(r.Address, address1)
		} else {
			r.Address[addressCount].ID = address1.ID
			r.Address[addressCount].Tags = address1.Tags
			r.Address[addressCount].AdditionalInfo = address1.AdditionalInfo
			r.Address[addressCount].City = address1.City
			r.Address[addressCount].Country = address1.Country
			r.Address[addressCount].PostalCode = address1.PostalCode
			r.Address[addressCount].Street = address1.Street
			r.Address[addressCount].StreetNumber = address1.StreetNumber
		}
	}
	if resp.Birthdate != nil {
		r.Birthdate = types.StringValue(resp.Birthdate.String())
	} else {
		r.Birthdate = types.StringNull()
	}
	r.CustomerNumber = types.StringPointerValue(resp.CustomerNumber)
	if len(r.Email) > len(resp.Email) {
		r.Email = r.Email[:len(resp.Email)]
	}
	for emailCount, emailItem := range resp.Email {
		var email1 BaseEmail
		email1.ID = types.StringPointerValue(emailItem.ID)
		email1.Tags = nil
		for _, v := range emailItem.Tags {
			email1.Tags = append(email1.Tags, types.StringValue(v))
		}
		email1.Email = types.StringValue(emailItem.Email)
		if emailCount+1 > len(r.Email) {
			r.Email = append(r.Email, email1)
		} else {
			r.Email[emailCount].ID = email1.ID
			r.Email[emailCount].Tags = email1.Tags
			r.Email[emailCount].Email = email1.Email
		}
	}
	r.FirstName = types.StringValue(resp.FirstName)
	r.ID = types.StringValue(resp.ID)
	r.LastName = types.StringValue(resp.LastName)
	if len(r.Phone) > len(resp.Phone) {
		r.Phone = r.Phone[:len(resp.Phone)]
	}
	for phoneCount, phoneItem := range resp.Phone {
		var phone1 BasePhone
		phone1.ID = types.StringPointerValue(phoneItem.ID)
		phone1.Tags = nil
		for _, v := range phoneItem.Tags {
			phone1.Tags = append(phone1.Tags, types.StringValue(v))
		}
		phone1.Phone = types.StringValue(phoneItem.Phone)
		if phoneCount+1 > len(r.Phone) {
			r.Phone = append(r.Phone, phone1)
		} else {
			r.Phone[phoneCount].ID = phone1.ID
			r.Phone[phoneCount].Tags = phone1.Tags
			r.Phone[phoneCount].Phone = phone1.Phone
		}
	}
	if resp.Salutation != nil {
		r.Salutation = types.StringValue(string(*resp.Salutation))
	} else {
		r.Salutation = types.StringNull()
	}
	if resp.Title != nil {
		r.Title = types.StringValue(string(*resp.Title))
	} else {
		r.Title = types.StringNull()
	}
}
