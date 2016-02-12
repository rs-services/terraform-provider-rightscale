variable "account_id" { }
variable "api_host"   { }
variable "refresh_token" {}

provider "rightscale" {
  account_id    = "${var.account_id}"
  api_host      = "${var.api_host}"
  refresh_token = "${var.refresh_token}"
}

resource "rightscale_deployment" "stefhen" {
  name        = "stefhen_terraform_created_deployment"
  description = "Test deployment created with Terraform"
}

output "rightscale_deployment.href" {
  value = "${rightscale_deployment.stefhen.id}"
}
