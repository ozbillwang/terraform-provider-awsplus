package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/hashicorp/terraform/plugin"
    "github.com/hashicorp/terraform/terraform"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() terraform.ResourceProvider {
            return Provider()
        },
    })
}

func Provider() *schema.Provider {
    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{
                "awsplus_vpc_peering_accept_all": resourceServer(),
        },
    }
}

func resourceServer() *schema.Resource {
    return &schema.Resource{
        Create: resourceServerCreate,
        Read:   resourceServerRead,
        Update: resourceServerUpdate,
        Delete: resourceServerDelete,

        Schema: map[string]*schema.Schema{
            "accepter": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "requester": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "aws_region": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
        accepter := d.Get("accepter").(string)
        requester := d.Get("requester").(string)
        aws_region := d.Get("aws_region").(string)
        err := query(accepter, requester, aws_region)
        if err != nil {
                return err
        }
        d.SetId(accepter + "-" + requester + " peering")
        return nil
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
    return nil
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
    return nil
}

func query(accepter string, requester string, aws_region string, ) error {
         svc := ec2.New(session.New(), &aws.Config{ Region: aws.String(aws_region) })

         params := &ec2.DescribeVpcPeeringConnectionsInput{}
         resp, err := svc.DescribeVpcPeeringConnections(params)
         if err != nil {
                 fmt.Println(err.Error())
                 return err
         }
         parseResponse(resp, svc, accepter, requester)
         return nil
}

func parseResponse(resp *ec2.DescribeVpcPeeringConnectionsOutput, svc *ec2.EC2, accepter string, requester string, ) {
        for _, v := range resp.VpcPeeringConnections {
                if isValidAccount(v.RequesterVpcInfo.OwnerId, requester) == true && isValidAccount(v.AccepterVpcInfo.OwnerId, accepter) == true {
                        err := acceptPeeringRequest(v, svc)
                        if err != nil {
                                fmt.Println(err.Error())
                        }
                }
        }
}

func isValidAccount(vpc *string, account_id string) bool {
         if account_id == *vpc {
                 return true
         }
         return false
}

func acceptPeeringRequest(v *ec2.VpcPeeringConnection, svc *ec2.EC2) error {
        if *v.Status.Code == "pending-acceptance" {
                params := &ec2.AcceptVpcPeeringConnectionInput{
                        VpcPeeringConnectionId: aws.String(*v.VpcPeeringConnectionId),
                }
                resp, err := svc.AcceptVpcPeeringConnection(params)
                if err != nil {
                        return err
                }
                fmt.Printf("VPC peering accepted")
                fmt.Printf("%s\n", resp)
        }
        return nil
}
