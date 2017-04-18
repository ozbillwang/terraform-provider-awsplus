### terraform-provider-awsplus
======================

[![Build Status](https://travis-ci.org/BWITS/terraform-provider-awsplus.svg?branch=master)](https://travis-ci.org/BWITS/terraform-provider-awsplus)

Missing feature in terraform aws provider. 

#### awsplus_vpc_peering_accept_all
Customized provider to accept VPC Peering connection 
requests from different AWS Account ID. At the moment,
aws_vpc_peering_connection resource only allow auto accept
requests if they belong to a same AWS Account ID.

### How to use this resource:

```
resource "awsplus_vpc_peering_accept_all" "MyPeering" {
  accepter        = "accepter_account_id"
  requester 	  = "requester_account_id"
  aws_region      = "us-west-2"
}

```

### Build the provider:

```
make release
```

## Contributing

1. Fork it 
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
