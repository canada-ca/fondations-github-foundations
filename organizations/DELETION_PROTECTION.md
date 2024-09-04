([Back](../README.md#included-tools))

# Deletion Protection

The `Deletion Protection` tool is automatically installed on all GitHub Foundations organizations.

`Deletion Protection` is a feature that helps to prevent accidental deletion of resources. When a resource is removed from the Terraform configuration, a mechanism kicks-in to ask for confirmation before the resource is deleted. It is implemented as a GitHub Action that is invoked when a Pull Request (PR) into the default branch is opened, and works by checking for any resources that are about to be deleted and asking for confirmation before proceeding with the deletion.

## Confirmation

When a PR is opened that deletes resources, the `Deletion Protection` tool will comment on the PR with a message telling the user:

```bash
#### ⚠️ The Terraform Plan contains ${process.env.DELETIONS} Deletion(s) ⚠️
            Please review the plan and ensure that the deletions are expected.

            If the deletions are expected, you must:

              1. Create a new comment on this PR.
              2. Set the contents to 'delete' (no quotes)
              3. Press the comment button.

            before you can merge.
```

The status of the PR Review is then set to `Changes Requested` until the user follows the instructions and confirms the deletions.

The user should review the plan, which is also a comment on the PR, and if the deletions are expected, they should create a new comment on the PR with the contents: `delete`.

Once the user has confirmed the deletions, the status of the PR Review is set to `Approved`.
