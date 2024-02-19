package openstack

var C struct {
	API          openstackAPI
	CATALOG_TYPE catalogT
	CATALOG_NAME catalogN
}

type keystone struct {
	V3_AUTH_TOKENS string
	V3_PROJECTS    string
	V3_DOMAIN      string
	V3_GROUPS      string
	V3_REGIONS     string
	V3_USERS       string
}

type nova struct {
	V2_1_OSHYPERVISORS_DETAIL string
	V2_1_SERVERS_DETAIL       string
	V2_1_USERS                string
	V2_1_FLAVORS_DETAIL       string
}

type octavia struct {
	V2_LBAAS_LOADBALANCERS string
	V2_LBAAS_LISTENERS     string
}

type ironic struct {
	V1_NODES                     string
	XOPENSTACK_IRONICAPI_VERSION string
}

type neutron struct {
	V2_0_NETWORKS string
	V2_0_SUBNETS  string
	V2_0_ROUTERS  string
	V2_0_PORTS    string
}

type glance struct {
	V2_IMAGE string
}

type cinder struct {
	V2_VOLUMES_DETAIL string
}

type openstackAPI struct {
	KEYSTONE keystone
	NOVA     nova
	OCTAVIA  octavia
	IRONIC   ironic
	NEUTRON  neutron
	GLANCE   glance
	CINDER   cinder
}

type catalogT struct {
	KEYSTONE string
	NOVA     string
	OCTAVIA  string
	IRONIC   string
	NEUTRON  string
	GLANCE   string
	CINDER   string
}

type catalogN struct {
	KEYSTONE string
	NOVA     string
	OCTAVIA  string
	IRONIC   string
	NEUTRON  string
	GLANCE   string
	CINDER   string
}

func init() {
	C.API.KEYSTONE.V3_AUTH_TOKENS = "/v3/auth/tokens"
	C.API.KEYSTONE.V3_PROJECTS = "/v3/projects"
	C.API.KEYSTONE.V3_DOMAIN = "/v3/domains"
	C.API.KEYSTONE.V3_GROUPS = "/v3/groups"
	C.API.KEYSTONE.V3_REGIONS = "/v3/regions"
	C.API.KEYSTONE.V3_USERS = "/v3/users"

	C.API.NOVA.V2_1_OSHYPERVISORS_DETAIL = "/v2.1/os-hypervisors/detail"
	C.API.NOVA.V2_1_SERVERS_DETAIL = "/v2.1/servers/detail"
	C.API.NOVA.V2_1_USERS = "/v3/users"
	C.API.NOVA.V2_1_FLAVORS_DETAIL = "/v2.1/flavors/detail"

	C.API.OCTAVIA.V2_LBAAS_LOADBALANCERS = "/v2/lbaas/loadbalancers"
	C.API.OCTAVIA.V2_LBAAS_LISTENERS = "/v2/lbaas/listeners"

	C.API.IRONIC.XOPENSTACK_IRONICAPI_VERSION = ""
	C.API.IRONIC.V1_NODES = "/v1/nodes"

	C.API.NEUTRON.V2_0_NETWORKS = "/v2.0/networks"
	C.API.NEUTRON.V2_0_SUBNETS = "/v2.0/subnets"
	C.API.NEUTRON.V2_0_ROUTERS = "/v2.0/routers"
	C.API.NEUTRON.V2_0_PORTS = "/v2.0/ports"

	C.API.GLANCE.V2_IMAGE = "/v2/images"

	C.API.CINDER.V2_VOLUMES_DETAIL = "/v3/{project_id}/volumes/detail"
	C.CATALOG_TYPE.KEYSTONE = "identity"
	C.CATALOG_TYPE.NOVA = "compute"
	C.CATALOG_TYPE.OCTAVIA = "load-balancer"
	C.CATALOG_TYPE.IRONIC = "baremetal"
	C.CATALOG_TYPE.NEUTRON = "network"
	C.CATALOG_TYPE.GLANCE = "image"
	C.CATALOG_TYPE.CINDER = "volumev3"

	C.CATALOG_NAME.KEYSTONE = "keystone"
	C.CATALOG_NAME.NOVA = "nova"
	C.CATALOG_NAME.OCTAVIA = "octavia"
	C.CATALOG_NAME.IRONIC = "ironic"
	C.CATALOG_NAME.NEUTRON = "neutron"
	C.CATALOG_NAME.GLANCE = "glance"
	C.CATALOG_NAME.CINDER = "cinder"
}
