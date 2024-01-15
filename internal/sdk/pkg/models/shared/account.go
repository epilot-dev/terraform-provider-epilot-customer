// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-customer/internal/sdk/pkg/utils"
	"time"
)

type Account struct {
	// Access control list (ACL) for an entity. Defines sharing access to external orgs or users.
	ACL       BaseEntityACL `json:"_acl"`
	CreatedAt time.Time     `json:"_created_at"`
	ID        string        `json:"_id"`
	// Organization Id the entity belongs to
	Org       string            `json:"_org"`
	Owners    []BaseEntityOwner `json:"_owners"`
	Schema    string            `json:"_schema"`
	Tags      []string          `json:"_tags"`
	Title     string            `json:"_title"`
	UpdatedAt time.Time         `json:"_updated_at"`
	// Addresses as a list of object, the element with index 0 is treated as the primary one.
	//
	Address        []BaseAddress `json:"address,omitempty"`
	CustomerNumber *string       `json:"customer_number,omitempty"`
	// Email addresses as a list of object, the element with index 0 is treated as the primary one.
	//
	Email []BaseEmail `json:"email,omitempty"`
	Name  string      `json:"name"`
	// Phone numbers as a list of object, the element with index 0 is treated as the primary one.
	//
	Phone   []BasePhone `json:"phone,omitempty"`
	Website *string     `json:"website,omitempty"`
}

func (a Account) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Account) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Account) GetACL() BaseEntityACL {
	if o == nil {
		return BaseEntityACL{}
	}
	return o.ACL
}

func (o *Account) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Account) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Account) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *Account) GetOwners() []BaseEntityOwner {
	if o == nil {
		return []BaseEntityOwner{}
	}
	return o.Owners
}

func (o *Account) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *Account) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Account) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}

func (o *Account) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *Account) GetAddress() []BaseAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *Account) GetCustomerNumber() *string {
	if o == nil {
		return nil
	}
	return o.CustomerNumber
}

func (o *Account) GetEmail() []BaseEmail {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *Account) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Account) GetPhone() []BasePhone {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *Account) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}
