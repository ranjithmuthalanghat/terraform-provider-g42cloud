# CHANGELOG

## 1.7.0 (December 29, 2022)

ENHANCEMENTS:

udpate docs ([#97](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/97))

## 1.6.1 (September 9, 2022)

FEATURES:

* **New Data Source:** `g42cloud_elb_flavors` ([#93](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/93))

## 1.6.0 (August 30, 2022)

FEATURES:

* **New Data Source:** `g42cloud_elb_certificate` ([#86](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/86))

* **New Resource:** `g42cloud_elb_certificate` ([#86](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/86))
* **New Resource:** `g42cloud_elb_ipgroup` ([#86](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/86))
* **New Resource:** `g42cloud_elb_l7policy` ([#86](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/86))
* **New Resource:** `g42cloud_elb_l7rule` ([#86](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/86))
* **New Resource:** `g42cloud_elb_listener` ([#86](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/86))
* **New Resource:** `g42cloud_elb_loadbalancer` ([#86](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/86))
* **New Resource:** `g42cloud_elb_member` ([#86](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/86))
* **New Resource:** `g42cloud_elb_monitor` ([#86](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/86))
* **New Resource:** `g42cloud_elb_pool` ([#86](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/86))
* **New Resource:** `g42cloud_networking_vip` ([#88](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/88))
* **New Resource:** `g42cloud_networking_vip_associate` ([#88](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/88))

## 1.5.0 (June 30, 2022)

FEATURES:

* **New Data Source:** `g42cloud_bms_flavors` ([#73](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/73))
* **New Data Source:** `g42cloud_css_flavors` ([#77](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/77))
* **New Data Source:** `g42cloud_modelarts_dataset_versions` ([#79](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/79))
* **New Data Source:** `g42cloud_modelarts_datasets` ([#79](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/79))
* **New Data Source:** `g42cloud_modelarts_notebook_images` ([#79](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/79))
* **New Data Source:** `g42cloud_apig_environments` ([#80](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/80))
* **New Data Source:** `g42cloud_servicestage_component_runtimes` ([#81](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/81))

* **New Resource:** `g42cloud_bms_instance` ([#73](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/73))
* **New Resource:** `g42cloud_cts_data_tracker` ([#74](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/74))
* **New Resource:** `g42cloud_cts_tracker` ([#74](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/74))
* **New Resource:** `g42cloud_lts_group` ([#75](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/75))
* **New Resource:** `g42cloud_lts_stream` ([#75](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/75))
* **New Resource:** `g42cloud_mapreduce_cluster` ([#76](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/76))
* **New Resource:** `g42cloud_mapreduce_job` ([#76](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/76))
* **New Resource:** `g42cloud_css_cluster` ([#77](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/77))
* **New Resource:** `g42cloud_css_snapshot` ([#77](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/77))
* **New Resource:** `g42cloud_css_thesaurus` ([#77](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/77))
* **New Resource:** `g42cloud_dws_cluster` ([#78](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/78))
* **New Resource:** `g42cloud_modelarts_dataset` ([#79](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/79))
* **New Resource:** `g42cloud_modelarts_dataset_version` ([#79](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/79))
* **New Resource:** `g42cloud_modelarts_notebook` ([#79](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/79))
* **New Resource:** `g42cloud_modelarts_notebook_mount_storage` ([#79](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/79))
* **New Resource:** `g42cloud_apig_api` ([#80](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/80))
* **New Resource:** `g42cloud_apig_api_publishment` ([#80](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/80))
* **New Resource:** `g42cloud_apig_application` ([#80](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/80))
* **New Resource:** `g42cloud_apig_custom_authorizer` ([#80](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/80))
* **New Resource:** `g42cloud_apig_environment` ([#80](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/80))
* **New Resource:** `g42cloud_apig_group` ([#80](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/80))
* **New Resource:** `g42cloud_apig_instance` ([#80](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/80))
* **New Resource:** `g42cloud_apig_response` ([#80](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/80))
* **New Resource:** `g42cloud_apig_throttling_policy` ([#80](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/80))
* **New Resource:** `g42cloud_apig_vpc_channel` ([#80](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/80))
* **New Resource:** `g42cloud_servicestage_application` ([#81](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/81))
* **New Resource:** `g42cloud_servicestage_component` ([#81](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/81))
* **New Resource:** `g42cloud_servicestage_component_instance` ([#81](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/81))
* **New Resource:** `g42cloud_servicestage_environment` ([#81](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/81))
* **New Resource:** `g42cloud_servicestage_repo_password_authorization` ([#81](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/81))
* **New Resource:** `g42cloud_servicestage_repo_token_authorization` ([#81](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/81))

ENHANCEMENTS:

* Support volume encryption in cce node and node pool ([#72](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/72))

## 1.4.0 (April 29, 2022)

FEATURES:

* **New Data Source:** `g42cloud_enterprise_project` ([#60](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/60))
* **New Data Source:** `g42cloud_vpcep_public_services` ([#63](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/63))
* **New Data Source:** `g42cloud_cbr_vaults` ([#67](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/67))
* **New Resource:** `g42cloud_enterprise_project` ([#60](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/60))
* **New Resource:** `g42cloud_swr_organization_permissions` ([#62](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/62))
* **New Resource:** `g42cloud_swr_organization` ([#62](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/62))
* **New Resource:** `g42cloud_swr_repository_sharing` ([#62](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/62))
* **New Resource:** `g42cloud_swr_repository` ([#62](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/62))
* **New Resource:** `g42cloud_vpcep_approval` ([#63](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/63))
* **New Resource:** `g42cloud_vpcep_endpoint` ([#63](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/63))
* **New Resource:** `g42cloud_vpcep_service` ([#63](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/63))
* **New Resource:** `g42cloud_cbr_policy` ([#67](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/67))
* **New Resource:** `g42cloud_cbr_vault` ([#67](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/67))
* **New Resource:** `g42cloud_tms_tags` ([#68](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/68))

ENHANCEMENTS:

* Support disk encryption in compute_instance ([#69](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/69))

## 1.3.1 (March 21, 2022)

FEATURES:

* **New Resource:** `g42cloud_vpc_route_table` ([#53](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/53))

ENHANCEMENTS:

* Update the reference huaweicloud ([#54](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/54))

## 1.3.0 (January 30, 2022)

FEATURES:

* **New Data Source:** `g42cloud_vpc_route_table` ([#48](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/48))
* **New Resource:** `g42cloud_vpc_route_table` ([#48](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/48))
* **New Resource:** `g42cloud_dns_ptrrecord` ([#49](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/49))

ENHANCEMENTS:

* Update the reference of sdk and huaweicloud ([#46](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/46))

## 1.2.1 (November 25, 2021)

ENHANCEMENTS:

* Support security_token ([#44](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/44))

## 1.2.0 (October 30, 2021)

FEATURES:

* **New Resource:** `g42cloud_networking_eip_associate` ([#42](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/42))

ENHANCEMENTS:

* Update the reference of sdk and huaweicloud ([#39](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/39))
* Update dcs resources ([#41](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/41))

## 1.1.0 (September 30, 2021)

FEATURES:

* **New Data Source:** `g42cloud_obs_bucket_object` ([#33](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/33))
* **New Data Source:** `g42cloud_rds_flavors` ([#34](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/34))
* **New Data Source:** `g42cloud_dms_az` ([#35](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/35))
* **New Data Source:** `g42cloud_dms_maintainwindow` ([#35](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/35))
* **New Data Source:** `g42cloud_dms_product` ([#35](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/35))
* **New Resource:** `g42cloud_obs_bucket_object` ([#33](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/33))
* **New Resource:** `g42cloud_obs_bucket_policy` ([#33](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/33))
* **New Resource:** `g42cloud_obs_bucket` ([#33](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/33))
* **New Resource:** `g42cloud_rds_configuration` ([#34](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/34))
* **New Resource:** `g42cloud_rds_instance` ([#34](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/34))
* **New Resource:** `g42cloud_rds_read_replica_instance` ([#34](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/34))
* **New Resource:** `g42cloud_dms_instance` ([#35](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/35))
* **New Resource:** `g42cloud_dli_queue` ([#36](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/36))
* **New Resource:** `g42cloud_ces_alarmrule` ([#37](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/37))

## 1.0.0 (July 8, 2021)

FEATURES:

* **New Data Source:** `g42cloud_kms_key` ([#26](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/26))
* **New Data Source:** `g42cloud_kms_data_key` ([#26](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/26))
* **New Data Source:** `g42cloud_dcs_az` ([#27](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/27))
* **New Data Source:** `g42cloud_dcs_maintainwindow` ([#27](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/27))
* **New Data Source:** `g42cloud_dcs_product` ([#27](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/27))
* **New Data Source:** `g42cloud_dds_flavors` ([#28](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/28))
* **New Resource:** `g42cloud_fgs_function` ([#25](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/25))
* **New Resource:** `g42cloud_kms_key` ([#26](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/26))
* **New Resource:** `g42cloud_dcs_instance` ([#27](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/27))
* **New Resource:** `g42cloud_dds_instance` ([#28](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/28))
* **New Resource:** `g42cloud_smn_topic` ([#29](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/29))
* **New Resource:** `g42cloud_smn_subscription` ([#29](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/29))

ENHANCEMENTS:

* provider: Make log message configurable ([#24](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/24))

## 0.4.0 (May 31, 2021)

FEATURES:

* **New Data Source:** `g42cloud_cce_cluster` ([#21](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/21))
* **New Data Source:** `g42cloud_cce_node` ([#21](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/21))
* **New Data Source:** `g42cloud_cce_node_pool` ([#21](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/21))
* **New Data Source:** `g42cloud_cce_addon_template` ([#21](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/21))
* **New Resource:** `g42cloud_identity_acl` ([#22](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/22))
* **New Resource:** `g42cloud_identity_agency` ([#22](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/22))
* **New Resource:** `g42cloud_identity_role` ([#22](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/22))
* **New Resource:** `g42cloud_cce_cluster` ([#21](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/21))
* **New Resource:** `g42cloud_cce_node` ([#21](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/21))
* **New Resource:** `g42cloud_cce_node_pool` ([#21](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/21))
* **New Resource:** `g42cloud_cce_addon` ([#21](https://github.com/g42cloud-terraform/terraform-provider-g42cloud/pull/21))

## 0.3.0 (March 2, 2021)

ENHANCEMENTS:

* provider: Add custom cloud and endpoints support (#10)

## 0.2.0 (February 9, 2021)

FEATURES:

* **New Resource:** `g42cloud_as_configuration`
* **New Resource:** `g42cloud_as_group`
* **New Resource:** `g42cloud_as_policy`
* **New Resource:** `g42cloud_network_acl`
* **New Resource:** `g42cloud_network_acl_rule`

## 0.1.0 (December 29, 2020)

FEATURES:

* **New Data Source:** `g42cloud_availability_zones`
* **New Data Source:** `g42cloud_compute_flavors`
* **New Data Source:** `g42cloud_identity_role`
* **New Data Source:** `g42cloud_images_image`
* **New Data Source:** `g42cloud_networking_port`
* **New Data Source:** `g42cloud_networking_secgroup`
* **New Data Source:** `g42cloud_vpc`
* **New Data Source:** `g42cloud_vpc_bandwidth`
* **New Data Source:** `g42cloud_vpc_route`
* **New Data Source:** `g42cloud_vpc_subnet`
* **New Data Source:** `g42cloud_vpc_subnet_ids`
* **New Resource:** `g42cloud_dns_recordset`
* **New Resource:** `g42cloud_dns_zone`
* **New Resource:** `g42cloud_identity_role_assignment`
* **New Resource:** `g42cloud_identity_user`
* **New Resource:** `g42cloud_identity_group`
* **New Resource:** `g42cloud_identity_group_membership`
* **New Resource:** `g42cloud_images_image`
* **New Resource:** `g42cloud_compute_instance`
* **New Resource:** `g42cloud_compute_interface_attach`
* **New Resource:** `g42cloud_compute_keypair`
* **New Resource:** `g42cloud_compute_servergroup`
* **New Resource:** `g42cloud_compute_eip_associate`
* **New Resource:** `g42cloud_compute_volume_attach`
* **New Resource:** `g42cloud_evs_snapshot`
* **New Resource:** `g42cloud_evs_volume`
* **New Resource:** `g42cloud_lb_certificate`
* **New Resource:** `g42cloud_lb_l7policy`
* **New Resource:** `g42cloud_lb_l7rule`
* **New Resource:** `g42cloud_lb_listener`
* **New Resource:** `g42cloud_lb_loadbalancer`
* **New Resource:** `g42cloud_lb_member`
* **New Resource:** `g42cloud_lb_monitor`
* **New Resource:** `g42cloud_lb_pool`
* **New Resource:** `g42cloud_lb_whitelist`
* **New Resource:** `g42cloud_networking_secgroup`
* **New Resource:** `g42cloud_networking_secgroup_rule`
* **New Resource:** `g42cloud_vpc`
* **New Resource:** `g42cloud_vpc_eip`
* **New Resource:** `g42cloud_vpc_subnet`
* **New Resource:** `g42cloud_vpc_route`
* **New Resource:** `g42cloud_vpc_peering_connection`
