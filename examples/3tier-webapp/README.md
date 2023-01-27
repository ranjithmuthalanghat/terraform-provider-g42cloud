variable "eps_id" {
  description = "The id of the enterprise project"
  default = "fc04bc67-2b65-4898-ab71-ce0d63fbc267"
}

variable "security_group_name" {
  description = "The name of the security group"
  default = "high-available-web-demo"
}

###############################################
# ELB related variables
###############################################

variable "ecs_flavor_id" {
  description = "The flavor id used to create the ecs instance"
  default = "s6.medium.2"
}

variable "eip_bandwidth_size" {
  description = "The bandwidth size of the eip"
  default = 5
}

variable "elbweb_name" {
  description = "The name of web elb"
  default = "HA-web"
}

variable "elbapp_name" {
  description = "The name of app elb"
  default = "HA-App"
}

variable "elb_listen_port" {
  description = "The listen port of the elb"
  default = "80"
}

###############################################
# RDS related variables
###############################################

variable "rds_name" {
  description = "The name of the rds instance"
  default = "high-available-web-demo"
}

variable "rds_flavor_id" {
  description = "The flavor id used to create the rds instance"
  default = "rds.pg.m6.large.8.ha"
}

variable "rds_volume_size" {
  description = "The volume size of the rds instance"
  default = 100
}

variable "rds_volume_type" {
  description = "The volume type of the rds instance"
  default = "ULTRAHIGH"
}


