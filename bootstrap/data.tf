locals {
  no_enterprise_account = var.github_enterprise_slug == ""
}

data "github_enterprise" "enterprise_account" {
  count = local.no_enterprise_account ? 0 : 1
  slug  = var.github_enterprise_slug
}
