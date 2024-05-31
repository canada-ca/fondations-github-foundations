# Projects

Projects in the organization are the main way to organize work. They are the primary way to group organizations and repos, under one umbrella.

**NOTE:** The term project will be renamed in the future to avoid confusion with GitHub's concept of a project.

## Table of Contents

   * [Projects](#projects)
      * [Configuring Projects](#configuring-projects)
      * [Configuring Repositories](#configuring-repositories)
      * [Configuring Teams](#configuring-teams)


## Configuring Projects

To configure a project for your organization:

1. Create a folder under `projects` with the name of your project.
2. Under this folder, create a folder for each of the organizations that will be involved in the project.
3. Under each of these folders, create a `repositories` and `teams` folder.
4. Under each of these folders, you will need to create a `terragrunt.hcl` file in order to configure the repositories and teams for the project.

### Example Directory Structure for Projects

Your directory structure will end up looking like this:

```
- projects
  - <YOUR PROJECT NAME>
    - <ORG 1>
      - repositories
        - terragrunt.hcl
      - teams
        - terragrunt.hcl
      ...
    ...
    - <ORG N>
      - repositories
        - terragrunt.hcl
      - teams
        - terragrunt.hcl
      ...
```

See the following sections for more information on how to configure repositories and teams for your project.

## Configuring Repositories

To configure a repository for your project:

1. Create a `terragrunt.hcl` file under the `repositories` folder for the organization that the repository will belong to.
2. The `terragrunt.hcl` file should make use of the `repository_set` module to manage all repositories related to this project that belong to this organization.

### Example `terragrunt.hcl` for a repository

```hcl
include "root" {
  path   = "${find_in_parent_folders()}"
  expose = true
}

include "providers" {
  path   = "${get_repo_root()}/providers/${basename(dirname(get_terragrunt_dir()))}/providers.hcl"
  expose = true
}

inputs = {
  public_repositories = {
    "<YOUR PUBLIC REPO NAME>" = {
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
    }

  }

  private_repositories = {
    "<YOUR PRIVATE REPO NAME>" = {
      description                          = "Repo2 description"
      default_branch                       = "main"
      repository_team_permissions_override = {}
      protected_branches                   = []       # Explicitly disable branch protection
      advance_security                     = false
      has_vulnerability_alerts             = true
      topics                               = []
      homepage                             = ""
      delete_head_on_merge                 = true
      allow_auto_merge                     = true
      dependabot_security_updates          = true
    }
  }

  default_repository_team_permissions = {
    "${dependency.teams.outputs.team_slugs["<YOUR TEAM SLUG>"]}" = "<DEFAULT PERMISSION>"
  }
}

terraform {
  source = "github.com/FociSolutions/github-foundations-modules//modules/repository_set"
}

dependency "teams" {
  config_path = "../teams"
}
```

You will need to configure the values for:
- `<YOUR PUBLIC REPO NAME>` and `<YOUR PRIVATE REPO NAME>` with the names of your public and private repositories.
- `<YOUR TEAM SLUG>` with the slug of the team that will have access to the repository.
- `<DEFAULT PERMISSION>` with the default permission that the team will have for the repository. It can be one of `pull`, `push`, or `admin`.

### Directory Structure for Repositories

Your directory structure should look like this:

```
- projects
  - <YOUR PROJECT NAME>
    - <ORG 1>
      - repositories
        - terragrunt.hcl
      ...
    ...
    - <ORG N>
      - repositories
        - terragrunt.hcl
        ...
```

## Configuring Teams

To configure a team for your project:

1. Create a `terragrunt.hcl` file under the `teams` folder for the organization that the team will belong to.
2. The `terragrunt.hcl` file should make use of the `team_set` module to manage all teams related to this project that belong to this organization.

### Example `terragrunt.hcl` for a team

```hcl
include "root" {
  path   = "${find_in_parent_folders()}"
  expose = true
}

include "providers" {
  path   = "${get_repo_root()}/providers/${basename(dirname(get_terragrunt_dir()))}/providers.hcl"
  expose = true
}

inputs = {
  teams = {
    "GhFoundationsAdmins" = {
      description = "The Admin team for ..."
      privacy     = "closed"
      members     = ["Member1", "Member2", ...]
      maintainers = ["Admin1", "Admin2", ...]
    }
    "GhFoundationsDevelopers" = {
      description     = "The development team for ..."
      privacy         = "closed"
      members         = ["Member1", "Member2", ...]
      maintainers     = ["Admin1", "Admin2", ...]
      parent_team_id  = <Optional parent team id>
    }
  }
}

terraform {
  source = "github.com/FociSolutions/github-foundations-modules//modules/team_set"
}
```

### Directory Structure for Teams

Your directory structure should look like this:

```
- projects
  - <YOUR PROJECT NAME>
    - <ORG 1>
      - teams
        - terragrunt.hcl
      ...
    ...
    - <ORG N>
      - teams
        - terragrunt.hcl
        ...
```