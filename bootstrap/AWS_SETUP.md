# Preparing Amazon Web Services for Github Foundations

This document will walk you through what is required of your Amazon Web Services (AWS) setup to run the Github Foundations bootstrap layer.

## Setup

**1. Create an IAM Policy**
* Create an IAM Policy on the AWS account the terraform state will be stored in. The policy requires the following minimum permissions:
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "iam:ListRoleTags",
                "iam:UpdateOpenIDConnectProviderThumbprint",
                "iam:PutRolePolicy",
                "dynamodb:DeleteTable",
                "resource-groups:GetGroupConfiguration",
                "s3:PutLifecycleConfiguration",
                "dynamodb:DescribeContinuousBackups",
                "iam:ListRolePolicies",
                "iam:DeleteOpenIDConnectProvider",
                "iam:GetRole",
                "dynamodb:UpdateTimeToLive",
                "s3:GetBucketWebsite",
                "iam:RemoveClientIDFromOpenIDConnectProvider",
                "s3:PutReplicationConfiguration",
                "iam:DeleteRole",
                "iam:UpdateRoleDescription",
                "iam:TagPolicy",
                "s3:DeleteBucketPolicy",
                "kms:DisableKey",
                "s3:GetReplicationConfiguration",
                "dynamodb:CreateTable",
                "resource-groups:CreateGroup",
                "s3:PutBucketObjectLockConfiguration",
                "iam:GetOpenIDConnectProvider",
                "iam:GetRolePolicy",
                "dynamodb:UpdateTable",
                "kms:EnableKey",
                "s3:GetLifecycleConfiguration",
                "s3:GetBucketTagging",
                "s3:UntagResource",
                "kms:UntagResource",
                "iam:UntagRole",
                "dynamodb:ListTables",
                "kms:PutKeyPolicy",
                "iam:PutRolePermissionsBoundary",
                "iam:TagRole",
                "dynamodb:ListTagsOfResource",
                "resource-groups:GetTags",
                "s3:ListBucket",
                "kms:ListResourceTags",
                "iam:DeleteRolePermissionsBoundary",
                "resource-groups:DeleteGroupPolicy",
                "iam:ListInstanceProfilesForRole",
                "s3:PutBucketTagging",
                "iam:DeleteRolePolicy",
                "kms:CreateKey",
                "s3:DeleteBucket",
                "s3:PutBucketVersioning",
                "kms:EnableKeyRotation",
                "kms:GetKeyPolicy",
                "iam:ListRoles",
                "s3:GetBucketVersioning",
                "resource-groups:PutGroupConfiguration",
                "resource-groups:DeleteGroup",
                "s3:PutBucketWebsite",
                "s3:PutBucketRequestPayment",
                "s3:GetBucketCORS",
                "iam:UntagPolicy",
                "iam:UpdateRole",
                "iam:UntagOpenIDConnectProvider",
                "iam:AddClientIDToOpenIDConnectProvider",
                "iam:TagOpenIDConnectProvider",
                "iam:UpdateAssumeRolePolicy",
                "iam:CreateRole",
                "s3:CreateBucket",
                "iam:AttachRolePolicy",
                "resource-groups:Untag",
                "s3:GetBucketObjectLockConfiguration",
                "iam:DetachRolePolicy",
                "s3:DeleteBucketWebsite",
                "resource-groups:GetGroup",
                "dynamodb:DescribeTable",
                "kms:GetKeyRotationStatus",
                "iam:ListAttachedRolePolicies",
                "iam:ListOpenIDConnectProviderTags",
                "s3:PutBucketAcl",
                "resource-groups:Tag",
                "s3:TagResource",
                "resource-groups:PutGroupPolicy",
                "resource-groups:UpdateGroupQuery",
                "s3:PutBucketCORS",
                "s3:PutBucketLogging",
                "kms:DeleteAlias",
                "resource-groups:GetGroupQuery",
                "resource-groups:UpdateGroup",
                "s3:PutAccelerateConfiguration",
                "resource-groups:GetGroupPolicy",
                "s3:GetBucketLogging",
                "s3:GetAccelerateConfiguration",
                "s3:GetBucketPolicy",
                "s3:PutEncryptionConfiguration",
                "s3:GetEncryptionConfiguration",
                "kms:TagResource",
                "dynamodb:TagResource",
                "s3:GetBucketRequestPayment",
                "kms:ScheduleKeyDeletion",
                "kms:DescribeKey",
                "dynamodb:UntagResource",
                "resource-groups:ListGroups",
                "dynamodb:DescribeTimeToLive",
                "s3:GetBucketAcl",
                "iam:CreateOpenIDConnectProvider",
                "kms:ListKeys",
                "iam:ListOpenIDConnectProviders",
                "s3:PutBucketPolicy"
            ],
            "Resource": "*"
        }
    ]
}
```

**2. Create an IAM Role**
* Create an IAM Role that you can impersonate to run the bootstrap layer.
* Ensure that your account has permissions to assume the role. For example a trusted entity policy that allows users and roles of an account to impersonate a role:
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "AWS": "<account-number>"
            },
            "Action": "sts:AssumeRole",
            "Condition": {}
        }
    ]
}
``` 
* Attach the previously created IAM Policy to the Role

**3. Assume Role and Set Environment Variables**
* Set the environment variable `AWS_REGION` with the region you want to use for the bootstrap resources.
* Assume the role created to run the bootstrap layer and fetch credentials and output them into a `creds.json` file with the command: `aws sts assume-role --role-arn "<ROLE-ARN>" --role-session-name <SESSION-NAME> > creds.json`.
* Set the environment variables `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` to the fetched values in `creds.json`. 

**Resources:**

* AWS CLI: [https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
* Creating IAM Policy: [https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_create.html](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_create.html)
* Creating IAM Role: [https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user.html](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user.html)
* Assume Role Command Reference: [https://docs.aws.amazon.com/cli/latest/reference/sts/assume-role.html](https://docs.aws.amazon.com/cli/latest/reference/sts/assume-role.html)