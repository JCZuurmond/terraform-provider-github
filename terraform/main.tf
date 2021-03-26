terraform {
  required_providers {
    github = {
      version = "0.2"
      source  = "hashicorp.com/edu/github"
    }
  }
}

provider "github" {}

module "orgs" {
  source = "./organizations"

  login = "gumgum"
}

output "organization" {
  value = module.orgs.organization
}

output "all_organizations" {
  value = module.orgs.all_organizations
}
