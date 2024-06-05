/**
 * Copyright 2021 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

# After the first successful "terraform apply":
#
# - uncomment this block
# - replace the bucket name with the one you set, or the default: github-tf-state-bucket-{{ you_org_id }}
# - (Optional) If using Azure: Replace the Storage Account name
# - run "terraform init"

### GCP ###
# terraform {
#   backend "gcs" {
#     bucket = "github-tf-state-bucket-1234567890"
#     prefix = "terraform/github-foundations/bootstrap"
#   }
# }

### AZURE ###
# terraform {
#  backend "azurerm" {
#    resource_group_name  = "StorageAccount-ResourceGroup"
#    storage_account_name = "replace-me-with-your-storage-account-name"
#    container_name       = "github-tf-state-bucket-1234567890"
#    key                  = "prod.terraform.tfstate"
#  }
# }
