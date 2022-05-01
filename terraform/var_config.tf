variable "enable_bastion" {
  default = true
}

variable "azs" {
  type    = list(string)
  default = ["us-east-1a", "us-east-1b"]
}

variable "vpc_cidr" {
  default = "10.0.0.0/24"
}

variable "az_count" {
  default = 2
}
