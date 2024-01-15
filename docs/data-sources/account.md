---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "epilot-customer_account Data Source - terraform-provider-epilot-customer"
subcategory: ""
description: |-
  Account DataSource
---

# epilot-customer_account (Data Source)

Account DataSource

## Example Usage

```terraform
data "epilot-customer_account" "my_account" {
  hydrate = false
  id      = "123e4567-e89b-12d3-a456-426614174000"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Account ID

### Optional

- `hydrate` (Boolean) Hydrates entities in relations when passed true

### Read-Only

- `acl` (Attributes) Access control list (ACL) for an entity. Defines sharing access to external orgs or users. (see [below for nested schema](#nestedatt--acl))
- `address` (Attributes List) Addresses as a list of object, the element with index 0 is treated as the primary one. (see [below for nested schema](#nestedatt--address))
- `created_at` (String)
- `customer_number` (String)
- `email` (Attributes List) Email addresses as a list of object, the element with index 0 is treated as the primary one. (see [below for nested schema](#nestedatt--email))
- `name` (String)
- `org` (String) Organization Id the entity belongs to
- `owners` (Attributes List) (see [below for nested schema](#nestedatt--owners))
- `phone` (Attributes List) Phone numbers as a list of object, the element with index 0 is treated as the primary one. (see [below for nested schema](#nestedatt--phone))
- `schema` (String)
- `tags` (List of String)
- `title` (String)
- `updated_at` (String)
- `website` (String)

<a id="nestedatt--acl"></a>
### Nested Schema for `acl`

Read-Only:

- `additional_properties` (String) Parsed as JSON.
- `delete` (List of String)
- `edit` (List of String)
- `view` (List of String)


<a id="nestedatt--address"></a>
### Nested Schema for `address`

Read-Only:

- `additional_info` (String)
- `city` (String)
- `country` (String) must be one of ["DE", "AT", "CH"]
- `id` (String)
- `postal_code` (String)
- `street` (String)
- `street_number` (String)
- `tags` (List of String)


<a id="nestedatt--email"></a>
### Nested Schema for `email`

Read-Only:

- `email` (String)
- `id` (String)
- `tags` (List of String)


<a id="nestedatt--owners"></a>
### Nested Schema for `owners`

Read-Only:

- `org_id` (String)
- `user_id` (String)


<a id="nestedatt--phone"></a>
### Nested Schema for `phone`

Read-Only:

- `id` (String)
- `phone` (String)
- `tags` (List of String)

