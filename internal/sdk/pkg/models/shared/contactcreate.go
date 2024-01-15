// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-customer/internal/sdk/pkg/types"
	"github.com/epilot-dev/terraform-provider-epilot-customer/internal/sdk/pkg/utils"
)

type ContactCreateSalutation string

const (
	ContactCreateSalutationMr                   ContactCreateSalutation = "Mr."
	ContactCreateSalutationMsMrs                ContactCreateSalutation = "Ms. / Mrs."
	ContactCreateSalutationCompany              ContactCreateSalutation = "Company"
	ContactCreateSalutationContactPerson        ContactCreateSalutation = "Contact Person"
	ContactCreateSalutationCompanyContactPerson ContactCreateSalutation = "Company/Contact Person"
	ContactCreateSalutationSpouse               ContactCreateSalutation = "Spouse"
	ContactCreateSalutationFamily               ContactCreateSalutation = "Family"
	ContactCreateSalutationOwnership            ContactCreateSalutation = "Ownership"
	ContactCreateSalutationAssembly             ContactCreateSalutation = "Assembly"
	ContactCreateSalutationOther                ContactCreateSalutation = "Other"
)

func (e ContactCreateSalutation) ToPointer() *ContactCreateSalutation {
	return &e
}

func (e *ContactCreateSalutation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Mr.":
		fallthrough
	case "Ms. / Mrs.":
		fallthrough
	case "Company":
		fallthrough
	case "Contact Person":
		fallthrough
	case "Company/Contact Person":
		fallthrough
	case "Spouse":
		fallthrough
	case "Family":
		fallthrough
	case "Ownership":
		fallthrough
	case "Assembly":
		fallthrough
	case "Other":
		*e = ContactCreateSalutation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ContactCreateSalutation: %v", v)
	}
}

type ContactCreateTitle string

const (
	ContactCreateTitleDr     ContactCreateTitle = "Dr."
	ContactCreateTitleProf   ContactCreateTitle = "Prof."
	ContactCreateTitleProfDr ContactCreateTitle = "Prof. Dr."
)

func (e ContactCreateTitle) ToPointer() *ContactCreateTitle {
	return &e
}

func (e *ContactCreateTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Dr.":
		fallthrough
	case "Prof.":
		fallthrough
	case "Prof. Dr.":
		*e = ContactCreateTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ContactCreateTitle: %v", v)
	}
}

type ContactCreate struct {
	Account *BaseRelation `json:"account,omitempty"`
	// Addresses as a list of object, the element with index 0 is treated as the primary one.
	//
	Address        []BaseAddress `json:"address,omitempty"`
	Birthdate      *types.Date   `json:"birthdate,omitempty"`
	CustomerNumber *string       `json:"customer_number,omitempty"`
	// Email addresses as a list of object, the element with index 0 is treated as the primary one.
	//
	Email     []BaseEmail `json:"email,omitempty"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	// Phone numbers as a list of object, the element with index 0 is treated as the primary one.
	//
	Phone      []BasePhone              `json:"phone,omitempty"`
	Salutation *ContactCreateSalutation `json:"salutation,omitempty"`
	Title      *ContactCreateTitle      `json:"title,omitempty"`
}

func (c ContactCreate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ContactCreate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ContactCreate) GetAccount() *BaseRelation {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *ContactCreate) GetAddress() []BaseAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ContactCreate) GetBirthdate() *types.Date {
	if o == nil {
		return nil
	}
	return o.Birthdate
}

func (o *ContactCreate) GetCustomerNumber() *string {
	if o == nil {
		return nil
	}
	return o.CustomerNumber
}

func (o *ContactCreate) GetEmail() []BaseEmail {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ContactCreate) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *ContactCreate) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *ContactCreate) GetPhone() []BasePhone {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *ContactCreate) GetSalutation() *ContactCreateSalutation {
	if o == nil {
		return nil
	}
	return o.Salutation
}

func (o *ContactCreate) GetTitle() *ContactCreateTitle {
	if o == nil {
		return nil
	}
	return o.Title
}