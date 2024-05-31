# Preparing Azure for Github Foundations

This document will walk you through what is required of your Azure (Az) setup to run the Github Foundations bootstrap layer.

## Setup

**1. Install Azure CLI tool**
* Install the Azure CLI tool according to your operating system's instructions ([https://learn.microsoft.com/en-us/cli/azure/install-azure-cli](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli))

**2. Setup an Azure Key Vault (optional)**
* Create an Azure Key Vault to store secrets that will be required by the Github Foundations organization layer to authenticate with Github. If you wish to bring your own secret manager you can skip this step.
* **Note: If you do bring a different solution to the toolkit ensure that it's secrets can either be accessed via Terraform or added to the environment of the github action runner so it can be read by Terraform**

**3. Authenticate with required permissions**
* At minimum your account will need the following roles scoped to the subscription or resource group that the Azure resources should be created in:
    * `Reader`
    * `Storage Account Contributor`
    * `Storage Blob Data Contributor`
    * `Managed Identity Contributor`
    * `Role Based Access Control Administrator`
* Additionally you will need the following roles if:
    * You want Github Foundations to create a new Resource Group for it's Azure resources:
        * `Contributor` or `Owner` scoped to the subscription the resource group will be created in.
        * **Note The previously listed roles should be scoped to the subscription since the resource group won't exist yet**
    * If you are using an Azure Key Vault to store secrets needed to authenticate with Github:
        * `Key Vault Reader` scoped to the resource group that the Azure Key Vault exists under.
        * `Role Based Access Control Administrator` scoped to the Azure Key Vault you want to use.
* Once you have confirmed your account has the required permissions authenticate using the Azure CLI: `az login`

**Resources:**

* Azure CLI: [https://learn.microsoft.com/en-us/cli/azure/install-azure-cli](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli)
* Azure Role-Based Access Control Built-in Roles: [https://learn.microsoft.com/en-us/azure/role-based-access-control/built-in-roles](https://learn.microsoft.com/en-us/azure/role-based-access-control/built-in-roles)
* Azure CLI Role Assignment: [https://learn.microsoft.com/en-us/azure/role-based-access-control/role-assignments-cli](https://learn.microsoft.com/en-us/azure/role-based-access-control/role-assignments-cli)