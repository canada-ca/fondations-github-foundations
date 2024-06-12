locals {
  no_enterprise_account = var.github_enterprise_slug == ""
  tf_state_bucket_name  = var.tf_state_bucket_name == "" ? format("github-tf-state-bucket-%d", var.org_id) : var.tf_state_bucket_name
}

data "github_enterprise" "enterprise_account" {
  count = local.no_enterprise_account ? 0 : 1
  slug  = var.github_enterprise_slug
}
