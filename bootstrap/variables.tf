variable "org_id" {
  description = "The organization id for the associated services."
  type        = string
}

variable "billing_account" {
  description = "The ID of the billing account to associate this project with."
  type        = string
}

#Github variables
variable "github_enterprise_slug" {
  type        = string
  description = "The URL slug to identify the enterprise account to use."
}

variable "github_foundations_organization_name" {
  type        = string
  description = "The organization name to use for the github foundations organization."
  default     = "github-foundations"
}

variable "github_organization_admin_logins" {
  type        = list(string)
  description = "List of github foundation organization owner usernames."
  validation {
    condition     = length(var.github_organization_admin_logins) > 0
    error_message = "The github_organization_admin_logins value must be a list of atleast length 1."
  }
}

variable "github_organization_billing_email" {
  type        = string
  description = "The billing email to set in the github foundations organization."
}

variable "github_enterprise_organizations" {
  type = map(object({
    display_name  = string
    description   = string
    billing_email = string
    admin_logins  = list(string)
  }))

  description = "A map of organizations to create in the enterprise account."
  default     = {}
}

variable "github_account_type" {
  type        = string
  description = "The type of github account being used. Can be 'Personal', 'Organization', or 'Enterprise'."

  validation {
    condition     = var.github_account_type == "Personal" || var.github_account_type == "Organization" || var.github_account_type == "Enterprise"
    error_message = "The github_account_type value must be 'Personal', 'Organization', or 'Enterprise'."
  }
}

variable "tf_state_bucket_name" {
  type        = string
  description = "The name to use for the Cloud storage bucket for storing terraform state. Defaults to 'github-tf-state-bucket-{{ var.org_id }}'."
  default     = ""
}
