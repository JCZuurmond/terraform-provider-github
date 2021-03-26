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

We start with the two most commonly used endpoints:
- [repository](https://docs.github.com/en/rest/reference/repos)
- [users](https://docs.github.com/en/rest/reference/users)
