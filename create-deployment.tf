variable "api_host"   { }
variable "password"   { }
variable "account_id" { }
variable "email"      { }

provider "rightscale" {
  api_host   = "${var.api_host}"
  password   = "${var.password}"
  account_id = "${var.account_id}"
  email      = "${var.email}"
}

resource "rightscale_deployment" "stefhen" {
  name        = "stefhen_terraform_created_deployment"
  description = "Test deployment created with Terraform #2"
}

output "rightscale_deployment.href" {
  value = "${rightscale_deployment.stefhen.id}"
}
