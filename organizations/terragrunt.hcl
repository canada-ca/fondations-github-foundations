locals {
  tf_state_bucket_project  = get_env("GCP_TF_STATE_BUCKET_PROJECT")
  tf_state_bucket_name     = get_env("GCP_TF_STATE_BUCKET_NAME")
  tf_state_bucket_location = get_env("GCP_TF_STATE_BUCKET_LOCATION")
}

remote_state {
  backend = "gcs"
  generate = {
    path      = "backend.tf"
    if_exists = "overwrite_terragrunt"
  }

  config = {
    project  = "${local.tf_state_bucket_project}"
    location = "${local.tf_state_bucket_location}"
    bucket   = "${local.tf_state_bucket_name}"
    prefix   = "terraform/github-foundations/organizations/${path_relative_to_include()}"
  }
}