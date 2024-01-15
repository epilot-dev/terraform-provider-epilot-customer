// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-customer/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteContactRequest struct {
	// Contact ID
	ContactID string `pathParam:"style=simple,explode=false,name=contactId"`
}

func (o *DeleteContactRequest) GetContactID() string {
	if o == nil {
		return ""
	}
	return o.ContactID
}

type DeleteContactResponse struct {
	// Any error based on client data errors
	ClientError *shared.ClientError
	// Response body for a contact entity
	Contact *shared.Contact
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Any error based on the server-side
	ServerError *shared.ServerError
}

func (o *DeleteContactResponse) GetClientError() *shared.ClientError {
	if o == nil {
		return nil
	}
	return o.ClientError
}

func (o *DeleteContactResponse) GetContact() *shared.Contact {
	if o == nil {
		return nil
	}
	return o.Contact
}

func (o *DeleteContactResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteContactResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteContactResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteContactResponse) GetServerError() *shared.ServerError {
	if o == nil {
		return nil
	}
	return o.ServerError
}
