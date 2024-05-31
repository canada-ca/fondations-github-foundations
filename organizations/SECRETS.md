# Secret Management

The github foundation terraform modules ways to manage organization-level and repository-level secrets in github.

The following sections will cover how to manage organization-level and repository-level secrets with these terraform modules.

## Usage

### Encryption

Before using any of the modules to create your github secrets it is important to encrypt them first. Github has documentation outlining how to encrypt secrets for their rest API that can be found here: https://docs.github.com/en/rest/guides/encrypting-secrets-for-the-rest-api

### Organization-level secrets

The organiztion and organization secrets modules have the following optional fields for organization-level secret management:

- `actions_secrets`: for creating and managing organization secrets accessible to github actions
- `codespaces_secrets`: for creating and managing organization secrets accessible to github codespaces
- `dependabot_secrets`: for creating and managing organization secrets accessible to dependabot

Each field takes a map where the key is the name of the secret and the value is an object with the following structure:
```hcl
{
    encrypted_value       = string
    visibility            = string
}
```

`encrypted_value` is the encrypted value of the secret you wish to store. See [encryption](#encryption) for instructions on how to encrypt your secret.

`visibility` configures what repositories can access your secret. It must be one of `all`, `private`, or `selected`.

If visibility is set to `selected` the secret will be created however no repositories will have access to it. To grant access to the secret make use of the repository set's `public_repositories` and `private_repositories` fields. Each of the fields are a map of object that contain the following optional fields:

- `organization_action_secrets`: for granting access to organization-level action secrets
- `organization_codespace_secrets`: for granting access to organization-level codespace secrets
- `organization_dependabot_secrets` : for granting access to organization-level dependabot secrets

Each of these fields take a list of strings that are the names of the secrets you want to give a repository access to. Example usage:
```hcl
inputs = {
  public_repositories = {
    "MyPublicRepo" = {
      description                          = "Repo1 description"
      default_branch                       = "main"
      repository_team_permissions_override = {}
      advance_security                     = false
      has_vulnerability_alerts             = true
      topics                               = []
      homepage                             = ""
      delete_head_on_merge                 = true
      allow_auto_merge                     = true
      dependabot_security_updates          = true
    
      organization_action_secrets    = ["SECRET_ONE"]
      organization_codespace_secrets = ["SECRET_THREE"]
    }

  }

  private_repositories = {
    "MyPrivateRepo" = {
      description                          = "Repo2 description"
      default_branch                       = "main"
      repository_team_permissions_override = {}
      advance_security                     = false
      has_vulnerability_alerts             = true
      topics                               = []
      homepage                             = ""
      delete_head_on_merge                 = true
      allow_auto_merge                     = true
      dependabot_security_updates          = true

      organization_action_secrets     = ["SECRET_ONE", "SECRET_TWO"] 
      organization_dependabot_secrets = ["SECRET_FOUR"]
    }
  }
}
```

### Repository-level secrets

The repository set module's `private_repositories` and `public_repositories` fields accepts the following optional fields for repository secret management:
- `action_secrets`: for creating repository secrets accessible to github actions
- `codespace_secrets`: for creating and managing repository secrets accessible to github codespaces
- `dependabot_secrets`: for creating and managing repository secrets accessible to dependabot
- `environments`: for creating and managing environments, including their action secrets.

`action_secrets`, `codespace_secrets`, and `dependabot_secrets` are maps where the keys are the name of the secrets and the values are the encrypted secret values to create and manage as secrets. See [encryption](#encryption) for instructions on how to encrypt your secret.

`environments` is a map where the keys are the names of the environments to create and manage. The values are objects that contain an `action_secrets` field which is a map where the keys are the name of the secrets and the values are the encrypted secret values to create and manage as environment action secrets