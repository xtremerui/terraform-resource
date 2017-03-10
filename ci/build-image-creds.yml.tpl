docker_email:             "ryang@pivotal.io"
docker_username:          "ruiyang"
docker_password:          "qmul1439"

resource_git_uri:         https://github.com/xtremerui/terraform-resource.git
resource_git_branch:      dev

terraform_git_uri:        https://github.com/hashicorp/terraform.git
terraform_git_branch:     master
# can be used to fetch RC builds - "v*-*"
terraform_git_tag_filter: "*"

docker_repository:        ruiyang/terraform-resource-gcs
docker_tag:               nightly

# optional, usually in format https://hooks.slack.com/services/XXXX
slack_url:                ""
