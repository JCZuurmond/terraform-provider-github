terraform {
  required_providers {
    github = {
      version = "0.2"
      source  = "hashicorp.com/edu/github"
    }
  }
}

variable "login" {
  type    = string
  default = "The login of the org"
}

data "github_organizations" "all" {}

output "all_organizations" {
  value = data.github_organizations.all.organizations
}

output "organization" {
  value = {
    for org in data.github_organizations.all.organizations :
    org.id => org
    if org.login == var.login
  }
}
