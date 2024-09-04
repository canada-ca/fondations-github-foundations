([Back](../README.md#included-tools)) / ([Retour](../README.md#outils-compris))

# Drift Detection

The `Drift Detection` tool is automatically installed on all GitHub Foundations organizations.

The `Drift Detection` tool is a tool that allows you to detect drift between your GitHub organization and your Terraform state. This tool is useful for detecting changes that have been made to your GitHub organization outside of Terraform.

## Table of Contents

  - [Drift Detection Jobs GitHub Action](#drift-detection-jobs-github-action)
    - [Features](#features)
  - [Usage](#usage)
    - [Scheduled Drift Detection](#scheduled-drift-detection)
    - [Issue Creation for Drifts](#issue-creation-for-drifts)
    - [Manual Re-Apply](#manual-re-apply)
    - [How to Manually Trigger a Drift Detection](#how-to-trigger-a-manual-drift-detection)
    - [How to Interpret the Drift Detection Results](#how-to-interpret-the-drift-detection-results)
  - [Configuration](#configuration)
  - [Setting Up Secrets](#setting-up-secrets)
  - [Contributing](#contributing)



## Drift Detection Jobs GitHub Action

This GitHub Action is designed to automate the detection of configuration drifts in your infrastructure as code (IaC) setup, specifically using Terraform and Terragrunt. It runs checks to identify any discrepancies between your IaC definitions and the actual state of your infrastructure, helping maintain consistency and compliance.

### Features

- **Scheduled Drift Detection**: Automatically runs `Drift Detection` jobs based on a schedule defined in cron format.
- **Issue Creation for Drifts**: Creates GitHub issues with detailed reports when drifts are detected.
- **Re-Apply Trigger**: Allows for manual re-triggering of the apply process by labeling issues with a specific label.
- **Support for Terraform and Terragrunt**: Utilizes both Terraform and Terragrunt for managing infrastructure as code.

## Prerequisites

Before you can use this GitHub Action, ensure you have:

- A GitHub repository with Terraform or Terragrunt configurations.
- Required secrets set up in your GitHub repository for GCP authentication and other necessary environment variables.

## Usage

### Scheduled Drift Detection

The `Drift Detection` jobs are scheduled to run automatically based on the cron schedule defined in the GitHub Action. By default, it's set to run at 00:10 every day. You can modify the cron schedule in the workflow file to meet your requirements.

### Issue Creation

When a drift is detected during the scheduled check and no other open issues exist with the `Drift` label, the action will create a GitHub issue in your repository. The issue will contain detailed information about the detected drift and steps for review. The issue will be labeled with `Drift` and `Action Required` labels for easy identification and filtering.

If a drift is detected and an open issue already exists with the `Drift` label, the action will update the GitHub issue's body with the newly detected changes in the GitHub infrastructure. This ensures only a single drift detection issue will be opened at a time and avoids a scenario where the repository's issue list is flooded with redundant drift detection issues.

### Manual Re-Apply

If after reviewing the drift report you decide to re-apply your Terraform or Terragrunt configurations, you can do so by adding a "Re-Apply" label to the created issue. This will trigger the re-apply job which attempts to resolve the drift by applying the latest configurations.

### How to Manually Trigger a Drift Detection

The `Drift Detection` tool can also be run manually. In order to run the `Drift Detection` tool manually, you will need to have the correct GCP permissions. To run the `Drift Detection` tool manually, follow these steps:

  1. Go to the `organizations` repository
  2. Click on the `Actions` tab
  3. Click on the `Drift Detection Jobs` workflow
  4. Click on the `Run Workflow` button
  5. Click on the `Run Workflow` button again to confirm that you want to run the workflow

### How to Interpret the Drift Detection Results

The `Drift Detection` tool raises an issue in the repository for each resource that has drifted, with the Terraform plan that has detected the drift.

Currently, `Terraform` will output the following change types:

* `+ create`
* `- destroy`
* `-/+ replace` (destroy and then create, or vice-versa if create-before-destroy is used)
* `~ update in-place`
* `<= read`

In order to determine the type of drift, and how to interpret the Terraform Plan in more detail, please refer to the [Terraform documentation](https://developer.hashicorp.com/terraform/cli), or try out the [Terraform Tutorials](https://www.terraform.io/docs/tutorials/)

## Configuration

To customize the behavior of the `Drift Detection` Jobs, you can modify the following environment variables in the GitHub Action workflow:

- `tf_version`: The version of Terraform to use.
- `tg_version`: The version of Terragrunt to use.
- `default_branch`: The default branch of your repository (usually `main`).
- `drift_label`: The label used to identify issues related to `Drift Detection`.
- `action_required_label`: An additional label used for issues requiring action.
- `working_dir`: The working directory for Terraform and Terragrunt commands.

## Setting Up Secrets

This action requires the following GitHub secrets to be set:

- `WORKLOAD_IDENTITY_PROVIDER`: Identifier for the GCP workload identity provider.
- `GCP_SERVICE_ACCOUNT`: The GCP service account to use for authentication.
- `GITHUB_TOKEN`: A GitHub token with permissions to create and manage issues and labels.

Ensure these secrets are correctly set in your repository's settings before using this action.

## Contributing

Contributions to the `Drift Detection` Jobs GitHub Action are welcome! Please feel free to submit pull requests or create issues for bugs, questions, or new features.

For more information, please see the [Contribution Guide](../CONTRIBUTING.md)
