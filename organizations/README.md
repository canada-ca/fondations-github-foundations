# Organizations

## Table of Contents

   * [Organizations](#organizations)
      * [Introduction](#introduction)
      * [Github App Installation](#github-app-installation)
      * [File and Folder Structure](#file-and-folder-structure)
      * [Configuring Providers](#configuring-providers)
      * [Configuring Repositories](#configuring-repositories)
      * [Configuring Teams](#configuring-teams)
      * [Secret Management](#secret-management)
   * [Running the Organizations Layer locally](#running-the-organizations-layer-locally)
      * [Prerequisites](#prerequisites)
   * [Pre-installed tools](#pre-installed-tools)


## Introduction

This layer contains all the terragrunt & terraform needed to:

1. Manage organization settings.
2. Create and manage all teams and repositories needed for a project across multiple organizations

This layer comes with two workflows:

1. Performs drift detection on a once per day basis, checking deployed github resources against the terraform state and raising issues when any drift is detected.
2. A workflow that runs plans on pull requests and performs a terraform apply after merge into `main`.

With the correct gcp permissions it can also be run locally.

## Github App Installation

Before this layer can be run to manage github resources under your organization(s) you will need to perform the following steps per organization:

**1. Create new github app with the organization**

* Register and install a new application with the github organization.

**2. Generate a new private key**

* In the github application's settings, generate a new private key. Do not delete the key from your local disk until it's stored in GCP secret manager.

**3. Create a secret in GCP secret manager**

* Create a secret that contains the private key contents generated in the previous step and the following annotations:
    * `appId` - The id of the github app registered and installed in your organization
    * `installationId` - The installation id of the github app registered and installed in your organization

**Note:** You don't have to include those annotations if you don't want but they are required for authenticating the terraform github provider. See [configuring providers](#configuring-providers) for more details.

**Note:** After this step you can delete the private key from your local disk.

## File and Folder Structure

The recommended file and folder structure for this layer is as follows:

* **organizations**
    * **terragrunt.hcl** - Terragrunt configuration that makes use of the `organization_settings` module to manage an organization
* **projects**
    * **PROJECT_NAME**
        * **Org1**
            * **repositories**
                * **terragrunt.hcl** - Terragrunt configuration that makes use of the `repository_set` module to manage all repositories related to this project that belong to this org (`Org1`)
            * **teams**
                * **terragrunt.hcl** - Terragrunt configuration that makes use of the `team_set` module to manage all teams related to this project that belong to this org (`Org1`)
        * **Org2**
            * **repositories**
                * **terragrunt.hcl**- Terragrunt configuration that makes use of the `repository_set` module to manage all repositories related to this project that belong to this org (`Org2`)
            * **teams**
                * **terragrunt.hcl** - Terragrunt configuration that makes use of the `team_set` module to manage all teams related to this project that belong to this org (`Org2`)
* **providers**
    * **Org1**
        * **providers.hcl** - Terragrunt configuration that sets up the github terraform provider with credentials for this org (`Org1`)
    * **Org2**
        * **providers.hcl** - Terragrunt configuration that sets up the github terraform provider with credentials for this org (`Org2`)
* **terragrunt.hcl** - Terragrunt configuration that sets up the backend to be used by all other configurations.

The following sections describing how to configure providers, organizations, project teams, and project repositories will all assume you are using this file and folder structure.

### Configuring Providers

To configure a provider for you organization name create a folder under `providers` with the name of your organization. Under this folder create a `providers.hcl` file with the following contents.

```
locals {
  organization_name      = "ORGANIZATION_NAME"
  secret_manager_project = get_env("GCP_SECRET_MANAGER_PROJECT")
}

generate "github_provider" {
  path      = "provider.tf"
  if_exists = "overwrite"
  contents  = <<EOF


provider "google" {
}

data "google_secret_manager_secret_version_access" "pem_file" {
  project = "${local.secret_manager_project}"
  secret = "GCP_SECRET_NAME"
}

data "google_secret_manager_secret" "pem_file_metadata" {
  project = "${local.secret_manager_project}"
  secret_id = "GCP_SECRET_NAME"
}

provider "github" {
    owner = "${local.organization_name}"
    app_auth {
      id              = data.google_secret_manager_secret.pem_file_metadata.annotations.appId
      installation_id = data.google_secret_manager_secret.pem_file_metadata.annotations.installationId
      pem_file        = data.google_secret_manager_secret_version_access.pem_file.secret_data
    }
    }
EOF
}
```

Change `ORGANIZATION_NAME` to match the organization id of the organization this provider should manage.

Change `GCP_SECRET_NAME` to match the name of the secret you created during [app installation](#github-app-installation). If you didn't create any secret please read over that section before continuing.

**Note:** If you didn't store installation id or app id in the secret annotations then you will need to either add those in for this file to work. Or you need to change the `installation_id` and `id` fields under the github provider's `app_auth` block.

### Configuring Repositories

See the documentation [here](./TEAMS_REPOS.md#configuring-repositories)

### Configuring Teams

See the documentation [here](./TEAMS_REPOS.md#configuring-teams)

### Secret Management

See the documentation [here](./SECRETS.md)

## Running the Organizations Layer locally

### Prerequisites

Before running the organizations layer, please ensure you have the following prerequisites met:

**1. Google Cloud Platform (GCP) CLI:**

* Install the GCP CLI according to your operating system's instructions ([https://cloud.google.com/sdk/docs/install](https://cloud.google.com/sdk/docs/install)).
* Authenticate to your GCP instance with the following command `gcloud auth application-default login`.

**2. GCP Service Account and Permissions:**

* Your GCP instance needs a service account with the following roles:
    * `roles/secretmanager.viewer`
    * `roles/secretmanager.secretAccessor`
    * `roles/iam.workloadIdentityUser`
    * `roles/storage.objectAdmin`
    * `roles/storage.admin`
* Additionally, you either: Need permission to impersonate the service account, or; you own the service account.

**3. Environment Variables:**
* You will need to set the following environment variables
    * `GCP_SECRET_MANAGER_PROJECT` - The project id where the secrets containing the github applications credentials are stored.
    * `GCP_TF_STATE_BUCKET_PROJECT` - The id of the gcp project with the bucket storing the terraform state files.
    * `GCP_TF_STATE_BUCKET_NAME` - The name of the bucket where the terraform state files will be stored.
    * `GCP_TF_STATE_BUCKET_LOCATION` - The location of the bucket where the terraform state files are stored.

**Note:** These prerequisites grant extensive permissions within your GCP project and organization. Please ensure you understand the implications of assigning these roles and permissions before proceeding.

**Resources:**

* GCP CLI: [https://cloud.google.com/sdk/docs/install](https://cloud.google.com/sdk/docs/install)
* GCP IAM Roles: [https://cloud.google.com/iam/docs/understanding-roles](https://cloud.google.com/iam/docs/understanding-roles)

## Pre-installed tools

The GitHub Foundations toolkit comes with:
* A [Drift Detection](./DRIFT_DETECTION.md) tool
* An [Interactive HCL generation](./GEN_INTERACTIVE.md) tool
* A [Deletion Protection](./DELETION_PROTECTION.md) tool
* [GitHub Advanced Security (GHAS) checks](./GH_ADVANCED_SECURITY.md)
* An [Assessment](./ASSESSMENT_TOOL.md) tool
* An [Import](./IMPORT_TOOL.md) tool
