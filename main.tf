#providers {
#  awsplus="release/terraform-provider-awsplus-darwin-amd64"
#}

resource "awsplus_vpc_peering_accept_all" "MyPeering" {
  accepter        = "123456789"
  requester       = "987654321"
  aws_region      = "us-west-2"
}
