terraform {
  required_providers {
    study = {
      source  = "test2/study"
      version = "0.0.1"
    }

  }
}

data "study_server_group" "d_test" {
  name = "no_empty"
}

locals {
  count = data.study_server_group.d_test.id == null ? 1 : 0
}
resource "study_server_group" "r_test" {
  count = local.count
  name  = "empty222"
}

output "sg_id" {
  value = data.study_server_group.d_test.id
}
output "name_d" {
  value = data.study_server_group.d_test.name
}

output "name_r" {
  value = study_server_group.r_test.*.name
}

resource "study_server_group" "r_one_test" {
  name = "empty2223333"
}

output "test_array_id" {
  value = study_server_group.r_one_test.test_array
}

