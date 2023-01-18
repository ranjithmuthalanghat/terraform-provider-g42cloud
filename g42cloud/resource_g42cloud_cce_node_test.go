package g42cloud

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/chnsz/golangsdk/openstack/cce/v3/nodes"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
)

func TestAccCCENodeV3_basic(t *testing.T) {
	var node nodes.Nodes

	rName := fmt.Sprintf("tf-acc-test-%s", acctest.RandString(5))
	updateName := rName + "update"
	resourceName := "g42cloud_cce_node.test"
	//clusterName here is used to provide the cluster id to fetch cce node.
	clusterName := "g42cloud_cce_cluster.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCCENodeV3Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCCENodeV3_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCCENodeV3Exists(resourceName, clusterName, &node),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
					resource.TestCheckResourceAttr(resourceName, "tags.key", "value"),
				),
			},
			{
				Config: testAccCCENodeV3_update(rName, updateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updateName),
					resource.TestCheckResourceAttr(resourceName, "tags.key", "value_update"),
				),
			},
			{
				Config: testAccCCENodeV3_auto_assign_eip(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestMatchResourceAttr(resourceName, "public_ip", regexp.MustCompile("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$")),
				),
			},
			{
				Config: testAccCCENodeV3_existing_eip(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestMatchResourceAttr(resourceName, "public_ip", regexp.MustCompile("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$")),
				),
			},
		},
	})
}

func TestAccCCENodeV3_volume_encryption(t *testing.T) {
	var node nodes.Nodes

	rName := fmt.Sprintf("tf-acc-test-%s", acctest.RandString(5))
	resourceName := "g42cloud_cce_node.test"
	//clusterName here is used to provide the cluster id to fetch cce node.
	clusterName := "g42cloud_cce_cluster.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckKms(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCCENodeV3Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCCENodeV3_volume_encryption(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCCENodeV3Exists(resourceName, clusterName, &node),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttrSet(resourceName, "root_volume.0.kms_key_id"),
					resource.TestCheckResourceAttrSet(resourceName, "data_volumes.0.kms_key_id"),
				),
			},
		},
	})
}

func testAccCheckCCENodeV3Destroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*config.Config)
	cceClient, err := config.CceV3Client(G42_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating G42Cloud CCE client: %s", err)
	}

	var clusterId string

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "g42cloud_cce_cluster" {
			clusterId = rs.Primary.ID
		}

		if rs.Type != "g42cloud_cce_node" {
			continue
		}

		_, err := nodes.Get(cceClient, clusterId, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Node still exists")
		}
	}

	return nil
}

func testAccCheckCCENodeV3Exists(n string, cluster string, node *nodes.Nodes) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		c, ok := s.RootModule().Resources[cluster]
		if !ok {
			return fmt.Errorf("Cluster not found: %s", c)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}
		if c.Primary.ID == "" {
			return fmt.Errorf("Cluster id is not set")
		}

		config := testAccProvider.Meta().(*config.Config)
		cceClient, err := config.CceV3Client(G42_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating G42Cloud CCE client: %s", err)
		}

		found, err := nodes.Get(cceClient, c.Primary.ID, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.Metadata.Id != rs.Primary.ID {
			return fmt.Errorf("Node not found")
		}

		*node = *found

		return nil
	}
}

func testAccCCENodeV3_Base(rName string) string {
	return fmt.Sprintf(`
%s

data "g42cloud_availability_zones" "test" {}

resource "g42cloud_compute_keypair" "test" {
  name = "%s"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAjpC1hwiOCCmKEWxJ4qzTTsJbKzndLo1BCz5PcwtUnflmU+gHJtWMZKpuEGVi29h0A/+ydKek1O18k10Ff+4tyFjiHDQAT9+OfgWf7+b1yK+qDip3X1C0UPMbwHlTfSGWLGZquwhvEFx9k3h/M+VtMvwR1lJ9LUyTAImnNjWG7TAIPmui30HvM2UiFEmqkr4ijq45MyX2+fLIePLRIFuu1p4whjHAQYufqyno3BS48icQb4p6iVEZPo4AE2o9oIyQvj2mx4dk5Y8CgSETOZTYDOR3rU2fZTRDRgPJDH9FWvQjF5tA0p3d9CoWWd2s6GKKbfoUIi8R/Db1BSPJwkqB jrp-hp-pc"
}

resource "g42cloud_cce_cluster" "test" {
  name                   = "%s"
  cluster_type           = "VirtualMachine"
  flavor_id              = "cce.s1.small"
  vpc_id                 = g42cloud_vpc.test.id
  subnet_id              = g42cloud_vpc_subnet.test.id
  container_network_type = "overlay_l2"
}
`, testAccCCEClusterV3_Base(rName), rName, rName)
}

func testAccCCENodeV3_basic(rName string) string {
	return fmt.Sprintf(`
%s

resource "g42cloud_cce_node" "test" {
  cluster_id        = g42cloud_cce_cluster.test.id
  name              = "%s"
  flavor_id         = "m6.large.8"
  availability_zone = data.g42cloud_availability_zones.test.names[0]
  key_pair          = g42cloud_compute_keypair.test.name
  os                = "CentOS 7.6"

  root_volume {
    size       = 40
    volumetype = "SSD"
  }
  data_volumes {
    size       = 100
    volumetype = "SSD"
  }
  tags = {
    foo = "bar"
    key = "value"
  }
}
`, testAccCCENodeV3_Base(rName), rName)
}

func testAccCCENodeV3_update(rName, updateName string) string {
	return fmt.Sprintf(`
%s

resource "g42cloud_cce_node" "test" {
  cluster_id        = g42cloud_cce_cluster.test.id
  name              = "%s"
  flavor_id         = "m6.large.8"
  availability_zone = data.g42cloud_availability_zones.test.names[0]
  key_pair          = g42cloud_compute_keypair.test.name
  os                = "CentOS 7.6"

  root_volume {
    size       = 40
    volumetype = "SSD"
  }
  data_volumes {
    size       = 100
    volumetype = "SSD"
  }
  tags = {
    foo = "bar"
    key = "value_update"
  }
}
`, testAccCCENodeV3_Base(rName), updateName)
}

func testAccCCENodeV3_auto_assign_eip(rName string) string {
	return fmt.Sprintf(`
%s

resource "g42cloud_cce_node" "test" {
  cluster_id        = g42cloud_cce_cluster.test.id
  name              = "%s"
  flavor_id         = "m6.large.8"
  availability_zone = data.g42cloud_availability_zones.test.names[0]
  key_pair          = g42cloud_compute_keypair.test.name
  os                = "CentOS 7.6"

  root_volume {
    size       = 40
    volumetype = "SSD"
  }
  data_volumes {
    size       = 100
    volumetype = "SSD"
  }

  // Assign EIP
  iptype="5_bgp"
  bandwidth_charge_mode="traffic"
  sharetype= "PER"
  bandwidth_size= 100
}
`, testAccCCENodeV3_Base(rName), rName)
}

func testAccCCENodeV3_existing_eip(rName string) string {
	return fmt.Sprintf(`
%s

resource "g42cloud_vpc_eip" "test" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name        = "test"
    size        = 8
    share_type  = "PER"
    charge_mode = "traffic"
  }
}

resource "g42cloud_cce_node" "test" {
  cluster_id        = g42cloud_cce_cluster.test.id
  name              = "%s"
  flavor_id         = "m6.large.8"
  availability_zone = data.g42cloud_availability_zones.test.names[0]
  key_pair          = g42cloud_compute_keypair.test.name
  os                = "CentOS 7.6"

  root_volume {
    size       = 40
    volumetype = "SSD"
  }
  data_volumes {
    size       = 100
    volumetype = "SSD"
  }

  // Assign existing EIP
  eip_id = g42cloud_vpc_eip.test.id
}
`, testAccCCENodeV3_Base(rName), rName)
}

func testAccCCENodeV3_volume_encryption(rName string) string {
	return fmt.Sprintf(`
%s

resource "g42cloud_kms_key" "test" {
  key_alias    = "%s"
  pending_days = "7"
}

resource "g42cloud_cce_node" "test" {
  cluster_id        = g42cloud_cce_cluster.test.id
  name              = "%s"
  flavor_id         = "m6.large.8"
  availability_zone = data.g42cloud_availability_zones.test.names[0]
  key_pair          = g42cloud_compute_keypair.test.name
  os                = "CentOS 7.6"

  root_volume {
    size       = 40
    volumetype = "SSD"
	kms_key_id = g42cloud_kms_key.test.id
  }
  data_volumes {
    size       = 100
    volumetype = "SSD"
    kms_key_id = g42cloud_kms_key.test.id
  }
  tags = {
    foo = "bar"
    key = "value"
  }
}
`, testAccCCENodeV3_Base(rName), rName, rName)
}
