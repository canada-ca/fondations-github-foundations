# Setting Up the Toolkit for Different Cloud Providers

Currently Github Foundations supports using the following cloud providers:
- Google Cloud Platform (GCP)
- Microsoft Azure (Az)

However it is possible to use other cloud providers to store Terraform state files remotely, the setup process is just a little different and requires some file modifications.

This document outlines the steps for configuring our toolkit to work with cloud providers that we don't currently support out of the box.

**Prerequisites:**

* Familiarity with Github Workflows, Terraform and Terragrunt
* Access to your chosen cloud provider with appropriate permissions

**Steps:**

1. **Setup Method of Authentication:**

   * Establish an authentication method for your cloud provider. This can be OIDC, a key file, a token, or any method supported by your platform.

2. **Bootstrap Layer Configuration:**

   * Update the bootstrap layer with the necessary secrets and action variables for GitHub Actions runners to authenticate with your cloud environment. Utilize the `custom` input variable with the following schema:

   ```
   object({
       organization_secrets = map(string)
       organization_variables = map(string)
       repository_secrets = map(map(string))
       repository_variables = map(map(string))
   })
   ```
   * In the `main.tf` file in the bootstrap layer remove the `github_gcloud_oidc` module. Your method of authentication for your cloud environment setup should have been done in the previous step.
   * Store the required secrets and variables required for Github Workflow runners to authenticate with your cloud provider within the appropriate sections (`organization_secrets` for secrets that runners in all Github Foundation repositories will need, `repository_secrets` for secrets that runners in only a select amount of repositories need, etc.).

3. **Terraform Backend Setup:**

   1. Execute `terraform apply` to apply the changes. Ensure you're authenticated to both GitHub and your cloud provider with the required permissions. Required scopes for Github can be found on the [main page readme](README.md)
   2. Within the bootstrap layer, modify the `backend.tf` file to leverage the appropriate backend configuration for your chosen cloud provider. Refer to the Terraform website for a list of supported backends [https://developer.hashicorp.com/terraform/language/settings/backends/remote](https://developer.hashicorp.com/terraform/language/settings/backends/configuration#available-backends).
   3. Run `terraform init -migrate-state` to migrate the local Terraform state created in the previous step to the configured remote backend storage in your cloud environment.

4. **Workflow Authentication Update:**

   * In the organization layer, modify all Github Workflow files to switch authentication away from GCP and towards your chosen cloud provider.
   * To do this replace all instances of the `GCP Auth` step with one that authenticates against your specific platform.

5. **Terragrunt Configuration:**

   * Update the root `terragrunt.hcl` file within the organization layer to define the appropriate backend configuration for your chosen cloud provider.

   **Example Configuration (AWS S3 Backend):**

   ```hcl
   locals {
     tf_state_bucket_name     = get_env("AWS_TF_STATE_BUCKET_NAME")
     tf_state_bucket_location = get_env("AWS_TF_STATE_BUCKET_LOCATION")
   }

   remote_state {
     backend = "s3"
     generate {
       path     = "backend.tf"
       if_exists = "overwrite_terragrunt"
     }

     config = {
       key            = "terraform/github-foundations/organizations/${path_relative_to_include()}"
       region         = "${local.tf_state_bucket_location}"
       bucket         = "${local.tf_state_bucket_name}"
       encrypt        = true
     }
   }
   ```

Once complete you should be able to push terragrunt.hcl configurations to the organization layer and the workflows should be able to plan and apply your infrastructure changes.
