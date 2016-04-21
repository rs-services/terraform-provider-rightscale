variable "account_id" { }
variable "api_host"   { }
variable "refresh_token" {}

provider "rightscale" {
  account_id    = "${var.account_id}"
  api_host      = "${var.api_host}"
  refresh_token = "${var.refresh_token}"
}

resource "rightscale_deployment" "stefhen" {
  name        = "Stefhen Created from Terraform"
  description = "UPDATED !!! Description from Terraform"
}
resource "rightscale_ssh_key" "stefhen_test" {
  name = "stefhen_test"
  cloud_href = "/api/clouds/7/ssh_keys"
}

output "rightscale_ssh_key.stefhen_test.id" {
  value = "${rightscale_ssh_key.stefhen_test.id}"
}

output "rightscale_deployment.stefhen.href" {
  value = "${rightscale_deployment.stefhen.id}"
}

output "rightscale_deployment.stefhen.name" {
  value = "${rightscale_deployment.stefhen.name}"
}

output "rightscale_deployment.stefhen.description" {
  value = "${rightscale_deployment.stefhen.description}"
}
