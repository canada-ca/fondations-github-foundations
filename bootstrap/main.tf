module "github_gcloud_oidc" {
  source = "github.com/FociSolutions/github-foundations-modules//modules/github-gcloud-oidc?ref=v0.0.1"

  #Folder
  parent      = "organizations/${var.org_id}"
  folder_name = "fldr-github-foundations"

  #Project
  billing_account = var.billing_account
  project_name    = "github-foundations"
  prefix          = "prj-g"
  services = [
    "storage-api.googleapis.com",
    "iam.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "iamcredentials.googleapis.com",
    "sts.googleapis.com",
    "secretmanager.googleapis.com"
  ]

  #Bucket
  bucket_name = "github-tf-state-bucket"
  location    = "northamerica-northeast1"
  versioning  = true
  lifecycle_rules = {
    lr-0 = {
      action = {
        type = "Delete"
      }
      condition = {
        num_newer_versions = "5"
      }
    }
  }
  force_destroy = true

  #OIDC Setup
  github_foundations_organization_name = var.github_foundations_organization_name
}

# Github foundations setup
module "github_foundations_organization" {
  source = "github.com/FociSolutions/github-foundations-modules//modules/enterprise-organization?ref=v0.0.1"
  providers = {
    github = github.enterprise_scoped
  }

  count = local.no_enterprise_account ? 0 : 1

  enterprise_id = local.no_enterprise_account ? "" : data.github_enterprise.enterprise_account[0].id
  name          = var.github_foundations_organization_name
  display_name  = "Github Foundations"
  description   = "Organization created to host github foundation toolkit repositories"
  admin_logins  = var.github_organization_admin_logins
  billing_email = var.github_organization_billing_email
}

module "github_foundations" {
  source = "github.com/FociSolutions/github-foundations-modules//modules/github-foundations?ref=v0.0.1"
  providers = {
    github = github.foundation_org_scoped
  }

  oidc_configuration = {
    gcp = {
      workload_identity_provider_name   = module.github_gcloud_oidc.provider_name
      bootstrap_workload_identity_sa    = module.github_gcloud_oidc.bootstrap_sa
      organization_workload_identity_sa = module.github_gcloud_oidc.organizations_sa
      gcp_secret_manager_project_id     = module.github_gcloud_oidc.project_id

      gcp_tf_state_bucket_project_id = module.github_gcloud_oidc.project_id
      bucket_name                    = module.github_gcloud_oidc.bucket_name
      bucket_location                = module.github_gcloud_oidc.bucket_location
      readme_path                    = ""
    }
  }

  account_type = var.github_account_type
}

# Other organizations that should exist under your github enterprise account
module "organizations" {
  source = "github.com/FociSolutions/github-foundations-modules//modules/enterprise-organization?ref=v0.0.1"
  providers = {
    github = github.enterprise_scoped
  }

  for_each = local.no_enterprise_account ? {} : var.github_enterprise_organizations

  enterprise_id = local.no_enterprise_account ? "" : data.github_enterprise.enterprise_account[0].id

  name          = each.key
  display_name  = each.value.display_name
  description   = each.value.description
  billing_email = each.value.billing_email
  admin_logins  = each.value.admin_logins
}
