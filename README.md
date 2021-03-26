# Terraform provider for Github
This repository is a Terraform provider for Github.

The repo is meant to be educational for learning Go and to understand the
internals of Terraform. 

Are you looking for a [verified Terraform provider for
Github](https://registry.terraform.io/providers/integrations/github/latest)?
Then have a look
[here](https://github.com/integrations/terraform-provider-github).

Examples are heavily borrowed from the Hashicorp
[learn](https://learn.hashicorp.com/collections/terraform/providers) about
providers.

# What will you find here?
The provider is a Terraform wrapper around the [Github
API](https://docs.github.com/en/rest) written in Go.

- [provider for Github](github/provider.go)
- [organizations data source](github/data_source_organizations.go)

# Installation
First install [Go](https://golang.org/doc/install). Then install the provider:

``` bash
make install
```

Notice that a `hashicorp.com/edu/github` provider will be installed in your
Terraform plugins.

To validate the install worked, use the example.  Run Terraform `init`:

``` bash
terraform -chdir=terraform/ init
```

Run `plan`:

``` bash
terraform -chdir=terraform/ plan
```

# Development
To validate changes with the example.  Run Terraform `init`:

``` bash
terraform -chdir=terraform/ init -upgrade=true
```

Run `plan`:

``` bash
terraform -chdir=terraform/ plan
```
