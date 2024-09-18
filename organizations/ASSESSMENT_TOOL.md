([Back](../README.md#included-tools))

# Assessment Tool

The `Assessment Tool` allows you to assess the readiness of your pre-existing GitHub `repository`, before importing it with the toolkit.
It is used to check whether toolkit guardrails are already in place in the `repository`, as well as any settings required at the `organization` level.

## Running the Assessment Tool

The `Assessment Tool` is available as part of the `GitHub Foundations CLI` tool, found [here](https://github.com/canada-ca/fondations-github-foundations/releases)

To run an assessment check on a repository, run the following command:

```bash
    Usage:
    gh_foundations check <org-slug>
```

`<org-slug>` is the organization slug to check. See the [documentation](https://github.com/canada-ca/fondations-github-foundations/blob/main/README.md#check) for more information.
