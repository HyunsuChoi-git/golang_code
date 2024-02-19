package test

import (
	"github.com/rs/zerolog/log"
	"golang_code/openstack"
	"golang_code/openstack/os_api_model"
)

func OpenstackTest() {

	provId := "9"
	err := openstack.GetAuthToken(provId)
	if err != nil {
		return
	}

	projects := &os_api_model.Projects{}
	openstack.GetMetrics(provId, "identity", "keystone", "/projects", projects)

	for _, p := range projects.Projects {
		log.Printf(p.ID)
		//volumes := &os_api_model.VolumesDetail{}
		//err := openstack.GetMetrics(provId, "volume3", "cinder", p.ID+"/volumes", volumes)
		//
		//if err != nil {
		//	log.Printf("openstack api error | project ID : %s | error msg : %s", p.ID, err.Error())
		//}
	}
}
