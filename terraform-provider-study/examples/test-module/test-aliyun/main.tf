# terraform {
#   required_providers {

#     alicloud = {
#       source  = "aliyun/alicloud"
#       version = "1.211.2"
#     }
#   }
# }

# data "alicloud_vswitches" "default" {
#   ids = ["vpc-bp1bpdpuedya0dtg1gaag", "vsw-bp126oycrl9drrqy7202y"]
# }

variable "vswitchs" {
  type = list(object({
    availability_zone = string
    vswitch_id        = string
    security_groups   = list(string)
    image_id          = string
  }))
  default = [
    {
      zone_id = "cn-hangzhou-j"
      vswitch_id        = "vsw-bp1kwdeuw54dehjzmg9jm"
      security_groups   = ["sg-bp1dlrlw9omoq4r97j0f"]
      image_id          = "centos_7_7_x64_20G_alibase_20200426.vhd"
    },
    {
      availability_zone = "cn-hangzhou-k"
      vswitch_id        = "vsw-bp1ancu5ujklokvh58xp3"
      security_groups   = ["sg-bp1dlrlw9omoq4r97j0f"]
      image_id          = "centos_7_7_x64_20G_alibase_20200426.vhd"
    }
  ]
}


output "vswtiches" {
  value = data.alicloud_vswitches.default
}