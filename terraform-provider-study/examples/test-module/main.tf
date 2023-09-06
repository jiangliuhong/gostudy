terraform {
  required_providers {
    study = {
      source  = "test/study"
      version = "0.0.1"
    }
  }
}

data "study_server_group" "d_test" {
  name = "empty2"
}
resource "study_server_group" "r_test" {
  count = data.study_server_group.d_test.id == null ? 1 : 0
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