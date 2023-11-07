terraform {
  required_providers {
    vsphere = {
      source  = "hashicorp/vsphere"
      version = "2.0.2"
    }
  }
}

provider "vsphere" {
  user           = "administrator@vsphere.local"
  vsphere_server = "10.0.1.24"
  password       = "Yunjikeji#123"

  # If you have a self-signed cert
  allow_unverified_ssl = true
}

variable "name_regex" {
  type = string
  default = "Network"
}


data "vsphere_network" "test" {
  #name          = "VM Network"
  datacenter_id = "datacenter-1449"
  #name_regex    = "^b.*b$"1464
#   name_regex    = "^V.*k$" 1463
    name_regex    = ".*${var.name_regex}.*"
}


output "test_type" {
  value = data.vsphere_network.test
}
