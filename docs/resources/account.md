---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "epilot-customer_account Resource - terraform-provider-epilot-customer"
subcategory: ""
description: |-
  Account Resource
---

# epilot-customer_account (Resource)

Account Resource

## Example Usage

```terraform
resource "epilot-customer_account" "my_account" {
  customer_number = "...my_customer_number..."
  name            = "Mr. Todd Hills"
  website         = "https://illiterate-fan.info"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `address` (Attributes List) Addresses as a list of object, the element with index 0 is treated as the primary one. (see [below for nested schema](#nestedatt--address))
- `customer_number` (String)
- `email` (Attributes List) Email addresses as a list of object, the element with index 0 is treated as the primary one. (see [below for nested schema](#nestedatt--email))
- `phone` (Attributes List) Phone numbers as a list of object, the element with index 0 is treated as the primary one. (see [below for nested schema](#nestedatt--phone))
- `website` (String)

### Read-Only

- `acl` (Attributes) Access control list (ACL) for an entity. Defines sharing access to external orgs or users. (see [below for nested schema](#nestedatt--acl))
- `created_at` (String)
- `id` (String) Account ID
- `org` (String) Organization Id the entity belongs to
- `owners` (Attributes List) (see [below for nested schema](#nestedatt--owners))
- `schema` (String)
- `tags` (List of String)
- `title` (String)
- `updated_at` (String)

<a id="nestedatt--address"></a>
### Nested Schema for `address`

Optional:

- `additional_info` (String)
- `city` (String) Not Null
- `country` (String) Not Null; must be one of ["DE", "AT", "CH"]
- `id` (String)
- `postal_code` (String) Not Null
- `street` (String) Not Null
- `street_number` (String) Not Null
- `tags` (List of String)


<a id="nestedatt--email"></a>
### Nested Schema for `email`

Optional:

- `email` (String) Not Null
- `id` (String)
- `tags` (List of String)


<a id="nestedatt--phone"></a>
### Nested Schema for `phone`

Optional:

- `id` (String)
- `phone` (String) Not Null
- `tags` (List of String)


<a id="nestedatt--acl"></a>
### Nested Schema for `acl`

Read-Only:

- `additional_properties` (String) Parsed as JSON.
- `delete` (List of String)
- `edit` (List of String)
- `view` (List of String)


<a id="nestedatt--owners"></a>
### Nested Schema for `owners`

Read-Only:

- `org_id` (String)
- `user_id` (String)

