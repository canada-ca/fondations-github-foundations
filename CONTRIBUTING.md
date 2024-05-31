# Contributing to the The GitHub Foundations Toolkit

We welcome contributions to the GitHub Foundations Toolkit. Please read these Contributing Guidelines for more information on how to get started.

## How to contribute

We welcome contributions from the community and are happy to have you as a contributor. Here are a few ways you can contribute:

- [Reporting issues](#reporting-issues)
- [Conventional Commits](#conventional-commits)
- [Submitting pull requests](#submitting-pull-requests)
- [Providing feedback](#providing-feedback)

### Requirements

You will need the following installed:
- [Terraform](https://developer.hashicorp.com/terraform/install)
- [Terragrunt](https://terragrunt.gruntwork.io/docs/getting-started/install/)
- [terraform-docs](https://github.com/terraform-docs/terraform-docs#installation)
- [Pre-commit](#pre-commit-checks)
- [TFLint](https://github.com/terraform-linters/tflint)
- [Trivy](https://aquasecurity.github.io/trivy/v0.49/getting-started/installation/)

#### Installing using Homebrew

To install the above using `homebrew`, do the following:

```bash
$ brew install terraform terragrunt pre-commit tflint trivy terraform-docs
```

#### Running the pre-commit checks manually

You can run the pre-commit checks manually. See [here](#pre-commit-checks) for more information.


## Reporting issues

If you find a bug or have a feature request, please report it in the [Issues](https://github.com/FociSolutions/github-foundations/issues)

## Conventional Commits

We use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) to ensure that our commit messages are easy to read and understand. This makes it easier to understand the history of the project and to generate changelogs.

When you submit a pull request, please use conventional commits for your commit messages. This will make it easier for us to understand your changes and to generate changelogs.

The format is:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Where type is one of:
 - build
 - ci
 - docs
 - feat
 - fix
 - perf
 - refactor
 - style
 - test
 - chore
 - revert
 - bump

## Submitting pull requests

If you would like to contribute to the project, please submit a pull request. We will review your pull request and work with you to get it merged.

A pull request description must include:
* A description of the problem you are solving
* A description of the solution
* A description of the testing you have done to verify the solution
* A description of any documentation updates that were needed

Your pull request must pass all existing tests, and you should add tests for any new functionality you are adding.

### Pre-commit checks

Your changes should also pass the [pre-commit](pre-commit.com) checks. To install pre-commit, see [the installation instructions](https://pre-commit.com/index.html#install).

You can install the pre-commit checks by running

```
$ pre-commit install
```

in the root of the repository.

Then, whenever you run `git commit`, the pre-commit checks will run. If any of the checks fail, the commit will be aborted.

If you'd like to run pre-commit on your `staged` files, you can run:

```
$ pre-commit run
```

If you would like to run pre-commit across your entire repo, you can run:

```
$ pre-commit run --all-files
```
## Providing feedback

We are always looking for feedback on how we can improve the project. If you have any feedback, please let us know in the [Issues](https://github.com/FociSolutions/github-foundations/issues) or [Discussions](https://github.com/FociSolutions/github-foundations/discussions).
