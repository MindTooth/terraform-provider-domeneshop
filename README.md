# Terraform Provider

Attempt at creating a Terraform-provider for [Domeneshop AS](https://domainname.shop/).

[API](https://api.domeneshop.no/)

## Getting started

Run `make install`.


```hcl
variable "ds_token" {
  type = string
}

variable "ds_secret" {
  type = string
}

terraform {
  required_providers {
    domeneshop = {
      source = "mindlab.no/terraform/domeneshop"
    }
  }
}

provider "domeneshop" {
  token  = var.ds_token
  secret = var.ds_secret
}

data "domeneshop_domain" "example" {
  name = "example.com"
}

output "example_id" {
  value = data.domeneshop_domain.example.domain_id
}

data "domeneshop_forward" "example_test" {
  domain_id = data.domeneshop_domain.example.domain_id
  host      = "test"
}

output "example_forward_test" {
  value = data.domeneshop_forward.example_test.url
}
```

### Data

* [Data](https://api.domeneshop.no/docs/index.html#id4)

### Resource

* [DNS](https://api.domeneshop.no/docs/index.html#dns)

## Guides

* [Official guide](https://www.terraform.io/docs/extend/writing-custom-providers.html)
* [Part #1](https://medium.com/spaceapetech/creating-a-terraform-provider-part-1-ed12884e06d7)
* [Part #2](https://medium.com/spaceapetech/creating-a-terraform-provider-part-2-1346f89f082c)
