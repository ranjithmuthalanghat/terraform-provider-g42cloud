# Terraform script for 3 Tier Application deployment on G42 Cloud

This is a sample terraform script to automate following Infrastructure components on G42 Cloud to enable provisioning of 3-tier application
 - VPC
 - Subnet
 - Security Group
 - ECS
 - RDS
 - ELB
 - EIP
The script contains 2 parts
 - Variable defintion: To define the varibles that will be used in the main terraform script
 - Main.tf: This is the core script which automate the provising of above G42 CLoud Infra structure


How to run the terraform script ?

1. Terraform installation. Download terraform file from https://developer.hashicorp.com/terraform/downloads to local machine or ECS 
2. Install terraform 
3. 
4. Copy the script to the terrafrom working directory
5. Initialize the working directory using the command "terraform init"
6. Verify the script using "terraform plan"
7. Deploy the Infra using "terraform apply"
8. To enable debug mode using the command "TF_LOG=DEBUG TF_LOG_PATH=./log terraform apply" 
