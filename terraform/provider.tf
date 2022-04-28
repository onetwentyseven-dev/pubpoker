provider "aws" {
  region  = "us-east-1"
  profile = "ots"
  default_tags {
    tags = {
      Project = "Pub Poker Championship"
    }
  }
}



provider "random" {
}


terraform {
  backend "s3" {
    bucket         = "ppc-terraform-state"
    region         = "us-east-1"
    key            = "terraform.tfstate"
    dynamodb_table = "ppc-terraform-state"
    profile        = "ots"
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.1.2"
    }
  }
}


