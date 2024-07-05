# Preparing Amazon Web Services for Github Foundations

This document will walk you through what is required of your Amazon Web Services (AWS) setup to run the Github Foundations bootstrap layer.

## Setup

**1. Create an IAM Policy**
* Create an IAM Policy on the AWS account that the terraform state will be stored in. The policy requires the following minimum permissions:
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "dynamodb:CreateTable",
                "dynamodb:DeleteTable",
                "dynamodb:DescribeTimeToLive",
                "dynamodb:ListTables",
                "dynamodb:DescribeContinuousBackups",
                "dynamodb:DescribeTable",
                "dynamodb:ListTagsOfResource",
                "dynamodb:TagResource",
                "dynamodb:UntagResource",
                "dynamodb:UpdateContinuousBackups",
                "dynamodb:UpdateTable",
                "dynamodb:UpdateTimeToLive",
                "iam:AddClientIDToOpenIDConnectProvider",
                "iam:AttachRolePolicy",
                "iam:CreateRole",
                "iam:CreateOpenIDConnectProvider",
                "iam:DeleteOpenIDConnectProvider",
                "iam:DeleteRole",
                "iam:DeleteRolePermissionsBoundary",
                "iam:DeleteRolePolicy",
                "iam:DetachRolePolicy",
                "iam:GetRole",
                "iam:GetRolePolicy",
                "iam:GetOpenIDConnectProvider",
                "iam:ListInstanceProfilesForRole",
                "iam:ListAttachedRolePolicies",
                "iam:ListOpenIDConnectProviders",
                "iam:ListOpenIDConnectProviderTags",
                "iam:ListRolePolicies",
                "iam:ListRoles",
                "iam:ListRoleTags",
                "iam:PutRolePolicy",
                "iam:PutRolePermissionsBoundary",
                "iam:RemoveClientIDFromOpenIDConnectProvider",
                "iam:TagOpenIDConnectProvider",
                "iam:TagPolicy",
                "iam:TagRole",
                "iam:UntagOpenIDConnectProvider",
                "iam:UntagPolicy",
                "iam:UntagRole",
                "iam:UpdateOpenIDConnectProviderThumbprint",
                "iam:UpdateRoleDescription",
                "iam:UpdateRole",
                "iam:UpdateAssumeRolePolicy",
                "kms:CreateKey",
                "kms:DeleteAlias",
                "kms:DescribeKey",
                "kms:DisableKey",
                "kms:EnableKey",
                "kms:EnableKeyRotation",
                "kms:GetKeyPolicy",
                "kms:GetKeyRotationStatus",
                "kms:ListKeys",
                "kms:ListResourceTags",
                "kms:PutKeyPolicy",
                "kms:ScheduleKeyDeletion",
                "kms:TagResource",
                "kms:UntagResource",
                "resource-groups:CreateGroup",
                "resource-groups:DeleteGroup",
                "resource-groups:DeleteGroupPolicy",
                "resource-groups:GetGroup",
                "resource-groups:GetGroupConfiguration",
                "resource-groups:GetGroupQuery",
                "resource-groups:GetTags",
                "resource-groups:ListGroups",
                "resource-groups:PutGroupConfiguration",
                "resource-groups:PutGroupPolicy",
                "resource-groups:Tag",
                "resource-groups:Untag",
                "resource-groups:UpdateGroup",
                "resource-groups:GetGroupPolicy",
                "resource-groups:UpdateGroupQuery",
                "s3:CreateBucket",
                "s3:DeleteBucket",
                "s3:DeleteBucketPolicy",
                "s3:DeleteBucketWebsite",
                "s3:GetAccelerateConfiguration",
                "s3:GetBucketAcl",
                "s3:GetBucketCORS",
                "s3:GetBucketLogging",
                "s3:GetBucketObjectLockConfiguration",
                "s3:GetBucketPolicy",
                "s3:GetBucketPublicAccessBlock",
                "s3:GetBucketRequestPayment",
                "s3:GetBucketTagging",
                "s3:GetBucketVersioning",
                "s3:GetBucketWebsite",
                "s3:GetEncryptionConfiguration",
                "s3:GetLifecycleConfiguration",
                "s3:GetReplicationConfiguration",
                "s3:ListBucket",
                "s3:PutAccelerateConfiguration",
                "s3:PutBucketAcl",
                "s3:PutBucketCORS",
                "s3:PutBucketLogging",
                "s3:PutBucketObjectLockConfiguration",
                "s3:PutBucketPolicy",
                "s3:PutBucketPublicAccessBlock",
                "s3:PutBucketRequestPayment",
                "s3:PutBucketTagging",
                "s3:PutBucketVersioning",
                "s3:PutBucketWebsite",
                "s3:PutEncryptionConfiguration",
                "s3:PutLifecycleConfiguration",
                "s3:PutReplicationConfiguration",
                "s3:TagResource",
                "s3:UntagResource"
            ],
            "Resource": "*"
        }
    ]
}
```

**2. Create an IAM Role**
* Create an IAM Role that you can impersonate to run the bootstrap layer.
* Ensure that your account has permissions to assume the role. For example a trusted entity policy that allows users and roles of an account to impersonate a role (called `Custom trust policy` on the IAM console):
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": "<account-ID>"
            },
            "Action": "sts:AssumeRole",
            "Condition": {}
        }
    ]
}
```
* Attach the previously created IAM Policy to the Role

**3. Assume Role and Set Environment Variables**

**Note:** For the following instructions, you will need to have setup a non-root user, and assigned an `Access Key ID` and `Secret Access Key`.
    See [here](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html) for more information.

* Set the environment variable `AWS_REGION` with the region you want to use for the bootstrap resources.
    * For example, `us-east-1`
* Make note of the Role `ARN`, for the role created above.
    * If you need to find the `ARN` in the console, it is under the `Summary` tab of the role.
* Make note of the `SESSION-NAME` you will use to assume the role.
    * This can be any value that you choose, for example: `ghf-bootstrap-session`
    * Click Show User Security Credentials
* Assume the role created to run the bootstrap layer and fetch credentials and output them into a `creds.json` file with the command: `aws sts assume-role --role-arn "<ROLE-ARN>" --role-session-name <SESSION-NAME> > creds.json`.
* Set the environment variables `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` to the fetched values in `creds.json`.

## Calculate Thumbprint of GitHub OIDC Provider

In order to run the AWS bootstrap, you will need to calculate the thumbprint of the GitHub OIDC provider. There is a [guide here](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc_verify-thumbprint.html) on how to do this.

There are two methods to obtain the thumbprint(s):

1. **Use the Value Provided**

    A sample value has been added to the `terraform.tfvars.template` file. The value is valid as of this writing, but may change in the future.

    If you are using the provided value, you can **skip the next method**.

2. **Calculate the Thumbprint**

    If you need to re-calculate the thumbprint, you can do so by running the following command in your terminal:

    ```bash
    HOST=$(curl https://vstoken.actions.githubusercontent.com/.well-known/openid-configuration \
    | jq -r '.jwks_uri | split("/")[2]')

    echo | openssl s_client -servername $HOST -showcerts -connect $HOST:443 2> /dev/null \
    | sed -n -e '/BEGIN/h' -e '/BEGIN/,/END/H' -e '$x' -e '$p' | tail +2 \
    | openssl x509 -fingerprint -noout \
    | sed -e "s/.*=//" -e "s/://g" \
    | tr "ABCDEF" "abcdef"
    ```

    The output will resemble: `1b511abead59c6ce207077c0bf0e0043b1382612`


**Resources:**

* AWS CLI: [https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
* Creating IAM Policy: [https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_create.html](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_create.html)
* Creating IAM Role: [https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user.html](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user.html)
* Assume Role Command Reference: [https://docs.aws.amazon.com/cli/latest/reference/sts/assume-role.html](https://docs.aws.amazon.com/cli/latest/reference/sts/assume-role.html)
* Calculate Thumbprint of OIDC Provider: [https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc_verify-thumbprint.html](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_providers_create_oidc_verify-thumbprint.html)
