# Preparing Google Cloud Platform for Github Foundations

This document will walk you through what is required of your Google Cloud Platform (GCP) setup to run the Github Foundations bootstrap layer.

## Setup

**1. Install Gcloud CLI tool**
* Install the Google Cloud Platform tool according to your operating system's instructions ([https://cloud.google.com/sdk/docs/install](https://cloud.google.com/sdk/docs/install))

**2. Authenticate with required permissions**
* Login to the GCP CLI with the command `gcloud auth application-default login` under a service account with the following roles for your organization:
    * `roles/iam.workloadIdentityPoolAdmin`
    * `roles/iam.serviceAccountAdmin`
    * `roles/resourcemanager.projectMover`
    * `roles/resourcemanager.projectDeleter`
    * `roles/resourcemanager.folderEditor`
    * `roles/storage.admin`
    * `roles/storage.objectAdmin`

**Note:** This setup will grant extensive permissions within your GCP project and organization. Please ensure you understand the implications of assigning these roles and permissions before proceeding.

**Resources:**

* GCP CLI: [https://cloud.google.com/sdk/docs/install](https://cloud.google.com/sdk/docs/install)
* GCP IAM Roles: [https://cloud.google.com/iam/docs/understanding-roles](https://cloud.google.com/iam/docs/understanding-roles)