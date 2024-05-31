# GitHub Foundations

## Securing Your GitHub Environment at Scale

### Quick Start: [Get Started Now!](./README.md#getting-started)

Managing GitHub for multiple groups introduces complex security and consistency challenges. Misaligned permission levels, team structures, incomplete change rollouts across groups, and status reporting are just a few of the obstacles that can manifest with manual processes.  The GitHub Foundations Toolkit offers a secure and efficient way to manage your organization's GitHub environment through automation and centralized control.

## Features:

### Automate Secure Infrastructure throughout your organization using CI/CD:
- Implement DevSecOps using **Terraform and Terragrunt** to apply Infrastructure as Code (IaC) principles in managing multiple groups under a single GitHub Enterprise account.
- Establish and enforce security best practices by default.
- Employ drift detection to promptly identify and rectify unauthorized changes, guaranteeing continuous security in your configurations.

### Centralize Control:
- Gain a comprehensive overview for managing groups, repositories, and teams across your organization.
- Streamline updates and security policy enforcement across your entire organization, reducing the need for manual intervention.
- Simplify administration by consistently and efficiently managing teams and their memberships across your organization.
- Enhance your security with branch protection through customizable **[Rulesets](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-rulesets/about-rulesets#about-rulesets)**.

### Rapid Rollouts:
- Push policy updates across the entire organization in moments rather than waiting days for each group to schedule individual exercises.
- Streamline your workflow by automating the creation, deletion, and configuration of both private and public repositories.

### Reduce Security Risks:
- Employ GitHub Advanced Security (GHAS) for auto code and dependency scanning on public repositories or private repositories where GHAS has been purchased.
- Enforce consistent security policies across all groups to minimize vulnerabilities and protect against attacks.
- Protect your data by securely storing and managing sensitive secrets directly in your GitHub environment.
- Have a unified view of potential vulnerabilities to prevent gaps opening in one group or another.

### Accelerate your ITSG-33 Controls
- Out of the box, GitHub Foundations has a scaffolding to assist with meeting controls in the following areas:
  - Authentication
  - Cryptographic Protection
  - Information Monitoring and Protection
  - Incident Monitoring and Response
  - Configuration Change Control
  - Security Attributes

## Getting Started:
The repository is organized into two layers for ease of setup and management:

### [Bootstrap Layer](./bootstrap/README.md)

Initial setup of your state file backend, and creation of all organizations under your GitHub Enterprise account.

### [Organizations Layer](./organizations/README.md)

Management of organizations, repositories, and teams, Review results of drift detection, and execution of pull request plans for your organizations.

#### Included Tools:
- **Drift Detection:** Detects when someone makes a change to configuration, outside of the source-controlled configuration. Gives the ability to reapply the correct state.
- **Deletion Protection:** When a PR change requests resources be deleted, this tool forces the user to confirm the action
- **GitHub Advanced Security (GHAS) checks:** Checks the state of GHAS for the repos that have it enabled. Reports all of the GHAS scans in one report.
- **Assessment tool:** Used to assess the readiness of your repo, before importing it with the toolkit. Can be used to check whether toolkit guardrails are already in place in the repo.
- **Import tool:** Import repos not currently managed by the toolkit.


## How to Contribute

See [CONTRIBUTING.md](./CONTRIBUTING.md)

## License

Unless otherwise noted, the source code of this project is distributed under the [MIT License](./LICENSE.md).

The Canada wordmark and related graphics associated with this distribution are protected under trademark law and copyright law. No permission is granted to use them outside the parameters of the Government of Canada's corporate identity program. For more information, see [Federal identity requirements](https://www.canada.ca/en/treasury-board-secretariat/topics/government-communications/federal-identity-requirements.html).
