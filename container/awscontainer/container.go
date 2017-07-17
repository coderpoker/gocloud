package awscontainer

import (
	"bytes"
	"encoding/json"
	"fmt"
	awsauth "github.com/scorelab/gocloud-v2/awsauth"
	"io/ioutil"
	"net/http"
)

type Ecscontainer struct {
}

type Createservice struct {
	ServiceName              string
	TaskDefinition           string
	DesiredCount             int
	ClientToken              string
	Cluster                  string
	Role                     string
	DeploymentConfigurations DeploymentConfiguration
	LoadBalancers            []LoadBalancer
	PlacementConstraints     []Placementconstraint
	PlacementStrategys       []Placementstrategy
}

type Placementconstraint struct {
	Expression string
	Type       string
}

type Placementstrategy struct {
	Field string
	Type  string
}

type LoadBalancer struct {
	ContainerName    string
	ContainerPort    int
	LoadBalancerName string
	TargetGroupArn   string
}

type DeploymentConfiguration struct {
	MaximumPercent        int
	MinimumHealthyPercent int
}

type Runtask struct {
	PlacementConstraints []Placementconstraint
	PlacementStrategys   []Placementstrategy
	Cluster              string
	Count                int
	Group                string
	StartedBy            string
	TaskDefinition       string
	overrides            override
}

type Starttask struct {
	Cluster            string
	ContainerInstances []string
	Group              string
	StartedBy          string
	TaskDefinition     string
	overrides          override
}

type override struct {
	ContainerOverrides []ContainerOverride
	TaskRoleArn        string
}

type ContainerOverride struct {
	Name              string
	MemoryReservation string
	Memory            int
	Cpu               int
	Command           []string
	Environments      []Environment
}

type Environment struct {
	Name  string
	Value string
}

type Deleteservice struct {
	Cluster string
	Service string
}

type Stoptask struct {
	Cluster string
	Reason  string
	Task    string
}

func (ecscontainer *Ecscontainer) Stoptask(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var Region string
	var stoptask Stoptask
	for key, value := range param {
		switch key {
		case "cluster":
			clusterV, _ := value.(string)
			stoptask.Cluster = clusterV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "reason":
			ReasonV, _ := value.(string)
			stoptask.Reason = ReasonV

		case "task":
			taskV, _ := value.(string)
			stoptask.Task = taskV
		}
	}
	params := make(map[string]string)
	preparestoptaskparams(params, stoptask, Region)
	stoptaskjsonmap := make(map[string]interface{})
	preparestoptaskparamsdict(stoptaskjsonmap, stoptask)
	ecscontainer.PrepareSignatureV4query(params, stoptaskjsonmap)
	return
}

func preparestoptaskparamsdict(stoptaskjsonmap map[string]interface{}, stoptask Stoptask) {
	if stoptask.Cluster != "" {
		stoptaskjsonmap["cluster"] = stoptask.Cluster
	}

	if stoptask.Reason != "" {
		stoptaskjsonmap["reason"] = stoptask.Reason
	}

	if stoptask.Task != "" {
		stoptaskjsonmap["task"] = stoptask.Task
	}
}

func preparestoptaskparams(params map[string]string, stoptask Stoptask, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.StopTask"
}

func (ecscontainer *Ecscontainer) Deleteservice(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var Region string
	var deleteservice Deleteservice
	for key, value := range param {
		switch key {
		case "cluster":
			clusterV, _ := value.(string)
			deleteservice.Cluster = clusterV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "service":
			serviceV, _ := value.(string)
			deleteservice.Service = serviceV
		}
	}
	params := make(map[string]string)
	preparedeleteserviceparams(params, deleteservice, Region)
	deleteServicejsonmap := make(map[string]interface{})
	preparedeleteserviceparamsdict(deleteServicejsonmap, deleteservice)
	ecscontainer.PrepareSignatureV4query(params, deleteServicejsonmap)
	return

}

func preparedeleteserviceparams(params map[string]string, deleteservice Deleteservice, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.DeleteService"
}

func preparedeleteserviceparamsdict(deleteServicejsonmap map[string]interface{}, deleteservice Deleteservice) {
	if deleteservice.Cluster != "" {
		deleteServicejsonmap["cluster"] = deleteservice.Cluster
	}
	if deleteservice.Service != "" {
		deleteServicejsonmap["service"] = deleteservice.Service
	}
}

func (ecscontainer *Ecscontainer) Starttask(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var starttask Starttask
	var Region string
	for key, value := range param {
		switch key {
		case "cluster":
			clusterV, _ := value.(string)
			starttask.Cluster = clusterV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "containerInstances":
			ContainerInstancesV, _ := value.([]string)
			starttask.ContainerInstances = ContainerInstancesV

		case "group":
			GroupV, _ := value.(string)
			starttask.Group = GroupV

		case "startedBy":
			StartedByV, _ := value.(string)
			starttask.StartedBy = StartedByV

		case "taskDefinition":
			TaskDefinitionV, _ := value.(string)
			starttask.TaskDefinition = TaskDefinitionV

		case "overrides":
			overridesparam, _ := value.(map[string]interface{})
			for overridesparamkey, overridesparamvalue := range overridesparam {
				switch overridesparamkey {
				case "taskRoleArn":
					starttask.overrides.TaskRoleArn = overridesparamvalue.(string)
				case "containerOverrides":
					containerOverridesparam, _ := overridesparamvalue.([]map[string]interface{})
					for i := 0; i < len(containerOverridesparam); i++ {
						var containerOverride ContainerOverride
						for containerOverrideparamkey, containerOverrideparamvalue := range containerOverridesparam[i] {
							switch containerOverrideparamkey {
							case "name":
								containerOverride.Name = containerOverrideparamvalue.(string)

							case "memoryReservation":
								containerOverride.MemoryReservation = containerOverrideparamvalue.(string)

							case "memory":
								containerOverride.Memory = containerOverrideparamvalue.(int)

							case "cpu":
								containerOverride.Cpu = containerOverrideparamvalue.(int)

							case "command":
								containerOverride.Command = containerOverrideparamvalue.([]string)

							case "environment":
								Environmentsparam := containerOverrideparamvalue.([]map[string]string)
								for i := 0; i < len(Environmentsparam); i++ {
									var environment Environment
									for environmentsparamkey, environmentsparamvalue := range Environmentsparam[i] {
										switch environmentsparamkey {
										case "name":
											environment.Name = environmentsparamvalue
										case "value":
											environment.Value = environmentsparamvalue
										}
									}
									containerOverride.Environments = append(containerOverride.Environments, environment)
								}
							}
						}
						starttask.overrides.ContainerOverrides = append(starttask.overrides.ContainerOverrides, containerOverride)
					}
				}
			}
		}
	}
	params := make(map[string]string)
	preparestarttaskparams(params, starttask, Region)
	starttaskjsonmap := make(map[string]interface{})
	preparestarttaskparamsdict(starttaskjsonmap, starttask)
	ecscontainer.PrepareSignatureV4query(params, starttaskjsonmap)
	return
}

func preparestarttaskparams(params map[string]string, starttask Starttask, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.StartTask"
}

func preparestarttaskparamsdict(starttaskjsonmap map[string]interface{}, starttask Starttask) {
	if starttask.Cluster != "" {
		starttaskjsonmap["cluster"] = starttask.Cluster
	}
	if starttask.TaskDefinition != "" {
		starttaskjsonmap["taskDefinition"] = starttask.TaskDefinition
	}
	if len(starttask.ContainerInstances) != 0 {
		starttaskjsonmap["containerInstances"] = starttask.ContainerInstances
	}

	if starttask.Group != "" {
		starttaskjsonmap["group"] = starttask.Group
	}
	if starttask.StartedBy != "" {
		starttaskjsonmap["startedBy"] = starttask.StartedBy
	}

	preparestarttaskoverridesparams(starttaskjsonmap, starttask)
}

func preparestarttaskoverridesparams(starttaskjsonmap map[string]interface{}, starttask Starttask) {
	overrides := make(map[string]interface{})
	if starttask.overrides.TaskRoleArn != "" {
		overrides["taskRoleArn"] = starttask.overrides.TaskRoleArn
	}
	if len(starttask.overrides.ContainerOverrides) != 0 {
		containerOverrides := make([]map[string]interface{}, 0)
		for i := 0; i < len(starttask.overrides.ContainerOverrides); i++ {
			containerOverride := make(map[string]interface{})
			if starttask.overrides.ContainerOverrides[i].Name != "" {
				containerOverride["name"] = starttask.overrides.ContainerOverrides[i].Name
			}
			if starttask.overrides.ContainerOverrides[i].MemoryReservation != "" {
				containerOverride["memoryReservation"] = starttask.overrides.ContainerOverrides[i].MemoryReservation
			}
			if starttask.overrides.ContainerOverrides[i].Memory != 0 {
				containerOverride["memory"] = starttask.overrides.ContainerOverrides[i].Memory
			}
			if starttask.overrides.ContainerOverrides[i].Cpu != 0 {
				containerOverride["cpu"] = starttask.overrides.ContainerOverrides[i].Cpu
			}
			if len(starttask.overrides.ContainerOverrides[i].Command) != 0 {
				containerOverride["command"] = starttask.overrides.ContainerOverrides[i].Command
			}

			if len(starttask.overrides.ContainerOverrides[i].Environments) != 0 {
				containerOverride["environment"] = starttask.overrides.ContainerOverrides[i].Environments
			}
			containerOverrides = append(containerOverrides, containerOverride)
		}
		overrides["containerOverrides"] = containerOverrides
	}
	if len(overrides) != 0 {
		starttaskjsonmap["overrides"] = overrides
	}
	fmt.Println(starttaskjsonmap)
}

func (ecscontainer *Ecscontainer) Runtask(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var runtask Runtask
	var Region string
	for key, value := range param {
		switch key {
		case "cluster":
			clusterV, _ := value.(string)
			runtask.Cluster = clusterV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "count":
			CountV, _ := value.(int)
			runtask.Count = CountV

		case "group":
			GroupV, _ := value.(string)
			runtask.Group = GroupV

		case "startedBy":
			StartedByV, _ := value.(string)
			runtask.StartedBy = StartedByV

		case "taskDefinition":
			TaskDefinitionV, _ := value.(string)
			runtask.TaskDefinition = TaskDefinitionV

		case "placementConstraints":
			placementconstraintsparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(placementconstraintsparam); i++ {
				var placementconstraint Placementconstraint
				for placementConstraintsparamkey, placementConstraintsparamvalue := range placementconstraintsparam[i] {
					switch placementConstraintsparamkey {
					case "Expression":
						placementconstraint.Expression = placementConstraintsparamvalue.(string)
					case "Type":
						placementconstraint.Type = placementConstraintsparamvalue.(string)
					}
				}
				runtask.PlacementConstraints = append(runtask.PlacementConstraints, placementconstraint)
			}

		case "placementStrategy":
			placementstrategyparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(placementstrategyparam); i++ {
				var placementstrategy Placementstrategy
				for placementstrategyparamkey, placementstrategyparamvalue := range placementstrategyparam[i] {
					switch placementstrategyparamkey {
					case "field":
						placementstrategy.Field = placementstrategyparamvalue.(string)
					case "Type":
						placementstrategy.Type = placementstrategyparamvalue.(string)
					}
				}
				runtask.PlacementStrategys = append(runtask.PlacementStrategys, placementstrategy)
			}

		case "overrides":
			overridesparam, _ := value.(map[string]interface{})
			fmt.Println(overridesparam)
			for overridesparamkey, overridesparamvalue := range overridesparam {
				switch overridesparamkey {
				case "taskRoleArn":
					runtask.overrides.TaskRoleArn = overridesparamvalue.(string)
				case "containerOverrides":
					containerOverridesparam, _ := overridesparamvalue.([]map[string]interface{})
					for i := 0; i < len(containerOverridesparam); i++ {
						var containerOverride ContainerOverride
						for containerOverrideparamkey, containerOverrideparamvalue := range containerOverridesparam[i] {
							switch containerOverrideparamkey {
							case "name":
								containerOverride.Name = containerOverrideparamvalue.(string)

							case "memoryReservation":
								containerOverride.MemoryReservation = containerOverrideparamvalue.(string)

							case "memory":
								containerOverride.Memory = containerOverrideparamvalue.(int)

							case "cpu":
								containerOverride.Cpu = containerOverrideparamvalue.(int)

							case "command":
								containerOverride.Command = containerOverrideparamvalue.([]string)

							case "environment":
								Environmentsparam := containerOverrideparamvalue.([]map[string]string)
								for i := 0; i < len(Environmentsparam); i++ {
									var environment Environment
									for environmentsparamkey, environmentsparamvalue := range Environmentsparam[i] {
										switch environmentsparamkey {
										case "name":
											environment.Name = environmentsparamvalue
										case "value":
											environment.Value = environmentsparamvalue
										}
									}
									containerOverride.Environments = append(containerOverride.Environments, environment)
								}
							}
						}
						runtask.overrides.ContainerOverrides = append(runtask.overrides.ContainerOverrides, containerOverride)
					}
				}
			}
		}
	}
	params := make(map[string]string)
	prepareruntaskparams(params, runtask, Region)
	runtaskjsonmap := make(map[string]interface{})
	prepareruntaskparamsdict(runtaskjsonmap, runtask)
	ecscontainer.PrepareSignatureV4query(params, runtaskjsonmap)
	return
}

func prepareruntaskparamsdict(runtaskjsonmap map[string]interface{}, runtask Runtask) {
	if runtask.Cluster != "" {
		runtaskjsonmap["cluster"] = runtask.Cluster
	}
	if runtask.TaskDefinition != "" {
		runtaskjsonmap["taskDefinition"] = runtask.TaskDefinition
	}
	if runtask.Count != 0 {
		runtaskjsonmap["count"] = runtask.Count
	}

	if runtask.Group != "" {
		runtaskjsonmap["group"] = runtask.Group
	}
	if runtask.StartedBy != "" {
		runtaskjsonmap["startedBy"] = runtask.StartedBy
	}

	prepareruntaskoverridesparams(runtaskjsonmap, runtask)
	prepareruntaskplacementConstraintsparams(runtaskjsonmap, runtask)
	prepareruntaskplacementStrategyparams(runtaskjsonmap, runtask)
}

func prepareruntaskoverridesparams(runtaskjsonmap map[string]interface{}, runtask Runtask) {
	overrides := make(map[string]interface{})
	if runtask.overrides.TaskRoleArn != "" {
		overrides["taskRoleArn"] = runtask.overrides.TaskRoleArn
	}
	if len(runtask.overrides.ContainerOverrides) != 0 {
		containerOverrides := make([]map[string]interface{}, 0)
		for i := 0; i < len(runtask.overrides.ContainerOverrides); i++ {
			containerOverride := make(map[string]interface{})
			if runtask.overrides.ContainerOverrides[i].Name != "" {
				containerOverride["name"] = runtask.overrides.ContainerOverrides[i].Name
			}
			if runtask.overrides.ContainerOverrides[i].MemoryReservation != "" {
				containerOverride["memoryReservation"] = runtask.overrides.ContainerOverrides[i].MemoryReservation
			}
			if runtask.overrides.ContainerOverrides[i].Memory != 0 {
				containerOverride["memory"] = runtask.overrides.ContainerOverrides[i].Memory
			}
			if runtask.overrides.ContainerOverrides[i].Cpu != 0 {
				containerOverride["cpu"] = runtask.overrides.ContainerOverrides[i].Cpu
			}
			if len(runtask.overrides.ContainerOverrides[i].Command) != 0 {
				containerOverride["command"] = runtask.overrides.ContainerOverrides[i].Command
			}

			if len(runtask.overrides.ContainerOverrides[i].Environments) != 0 {
				containerOverride["environment"] = runtask.overrides.ContainerOverrides[i].Environments
			}
			containerOverrides = append(containerOverrides, containerOverride)
		}
		overrides["containerOverrides"] = containerOverrides
	}
	if len(overrides) != 0 {
		runtaskjsonmap["overrides"] = overrides
	}
	fmt.Println(runtaskjsonmap)
}

func prepareruntaskparams(params map[string]string, runtask Runtask, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.RunTask"
}

func prepareruntaskplacementStrategyparams(Createservicejsonmap map[string]interface{}, runtask Runtask) {
	if len(runtask.PlacementStrategys) != 0 {
		placementstrategys := make([]map[string]interface{}, 0)
		for i := 0; i < len(runtask.PlacementStrategys); i++ {
			placementstrategy := make(map[string]interface{})

			if runtask.PlacementStrategys[i].Field != "" {
				placementstrategy["field"] = runtask.PlacementStrategys[i].Field
			}

			if runtask.PlacementStrategys[i].Type != "" {
				placementstrategy["type"] = runtask.PlacementStrategys[i].Type
			}

			placementstrategys = append(placementstrategys, placementstrategy)
		}

		Createservicejsonmap["placementstrategy"] = placementstrategys
	}
}

func prepareruntaskplacementConstraintsparams(runtaskjsonmap map[string]interface{}, runtask Runtask) {
	if len(runtask.PlacementConstraints) != 0 {
		placementConstraints := make([]map[string]interface{}, 0)
		for i := 0; i < len(runtask.PlacementConstraints); i++ {
			PlacementConstraint := make(map[string]interface{})

			if runtask.PlacementConstraints[i].Expression != "" {
				PlacementConstraint["expression"] = runtask.PlacementConstraints[i].Expression
			}

			if runtask.PlacementConstraints[i].Type != "" {
				PlacementConstraint["type"] = runtask.PlacementConstraints[i].Type
			}

			placementConstraints = append(placementConstraints, PlacementConstraint)
		}

		runtaskjsonmap["placementConstraints"] = placementConstraints
	}
}

func (ecscontainer *Ecscontainer) Createservice(request interface{}) (resp interface{}, err error) {
	param := request.(map[string]interface{})
	var createservice Createservice
	var Region string
	for key, value := range param {
		switch key {
		case "serviceName":
			serviceNameV, _ := value.(string)
			createservice.ServiceName = serviceNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV

		case "taskDefinition":
			TaskDefinitionV, _ := value.(string)
			createservice.TaskDefinition = TaskDefinitionV

		case "desiredCount":
			desiredCountV, _ := value.(int)
			createservice.DesiredCount = desiredCountV

		case "clientToken":
			clientTokenV, _ := value.(string)
			createservice.ClientToken = clientTokenV

		case "cluster":
			clusterV, _ := value.(string)
			createservice.Cluster = clusterV

		case "role":
			roleV, _ := value.(string)
			createservice.Role = roleV

		case "deploymentConfiguration":
			deploymentConfigurationV, _ := value.(map[string]int)
			createservice.DeploymentConfigurations.MaximumPercent = deploymentConfigurationV["maximumPercent"]
			createservice.DeploymentConfigurations.MinimumHealthyPercent = deploymentConfigurationV["minimumHealthyPercent"]

		case "LoadBalancers":
			LoadBalancersparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(LoadBalancersparam); i++ {
				var loadBalancer LoadBalancer
				for loadBalancersparamkey, loadBalancersparamparamvalue := range LoadBalancersparam[i] {
					switch loadBalancersparamkey {
					case "containerName":
						loadBalancer.ContainerName = loadBalancersparamparamvalue.(string)
					case "containerPort":
						loadBalancer.ContainerPort = loadBalancersparamparamvalue.(int)
					case "loadBalancerName":
						loadBalancer.LoadBalancerName = loadBalancersparamparamvalue.(string)
					case "targetGroupArn":
						loadBalancer.TargetGroupArn = loadBalancersparamparamvalue.(string)
					}
				}
				createservice.LoadBalancers = append(createservice.LoadBalancers, loadBalancer)
			}

		case "placementConstraints":
			placementconstraintsparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(placementconstraintsparam); i++ {
				var placementconstraint Placementconstraint
				for placementConstraintsparamkey, placementConstraintsparamvalue := range placementconstraintsparam[i] {
					switch placementConstraintsparamkey {
					case "Expression":
						placementconstraint.Expression = placementConstraintsparamvalue.(string)
					case "Type":
						placementconstraint.Type = placementConstraintsparamvalue.(string)
					}
				}
				createservice.PlacementConstraints = append(createservice.PlacementConstraints, placementconstraint)
			}

		case "placementStrategy":
			placementstrategyparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(placementstrategyparam); i++ {
				var placementstrategy Placementstrategy
				for placementstrategyparamkey, placementstrategyparamvalue := range placementstrategyparam[i] {
					switch placementstrategyparamkey {
					case "field":
						placementstrategy.Field = placementstrategyparamvalue.(string)
					case "Type":
						placementstrategy.Type = placementstrategyparamvalue.(string)
					}
				}
				createservice.PlacementStrategys = append(createservice.PlacementStrategys, placementstrategy)
			}

		}
	}
	params := make(map[string]string)
	preparecreateServiceparams(params, createservice, Region)
	Createservicejsonmap := make(map[string]interface{})
	preparecreateServiceparamsdict(Createservicejsonmap, createservice)
	ecscontainer.PrepareSignatureV4query(params, Createservicejsonmap)
	return
}

func preparecreateServiceparams(params map[string]string, createservice Createservice, Region string) {
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.CreateService"
}

func preparecreateServiceparamsdict(Createservicejsonmap map[string]interface{}, createservice Createservice) {
	if createservice.ServiceName != "" {
		Createservicejsonmap["serviceName"] = createservice.ServiceName
	}
	if createservice.TaskDefinition != "" {
		Createservicejsonmap["taskDefinition"] = createservice.TaskDefinition
	}
	if createservice.DesiredCount != 0 {
		Createservicejsonmap["desiredCount"] = createservice.DesiredCount
	}

	if createservice.ClientToken != "" {
		Createservicejsonmap["clientToken"] = createservice.ClientToken
	}
	if createservice.Cluster != "" {
		Createservicejsonmap["cluster"] = createservice.Cluster
	}
	if createservice.Role != "" {
		Createservicejsonmap["role"] = createservice.Role
	}
	preparecreateServicedeploymentConfigurationparams(Createservicejsonmap, createservice)
	preparecreateServiceloadBalancersparams(Createservicejsonmap, createservice)

	preparecreateServiceplacementConstraintsparams(Createservicejsonmap, createservice)
	preparecreateServiceplacementStrategyparams(Createservicejsonmap, createservice)

}

func preparecreateServiceplacementStrategyparams(Createservicejsonmap map[string]interface{}, createservice Createservice) {
	if len(createservice.PlacementStrategys) != 0 {
		placementstrategys := make([]map[string]interface{}, 0)
		for i := 0; i < len(createservice.PlacementStrategys); i++ {
			placementstrategy := make(map[string]interface{})

			if createservice.PlacementStrategys[i].Field != "" {
				placementstrategy["field"] = createservice.PlacementStrategys[i].Field
			}

			if createservice.PlacementStrategys[i].Type != "" {
				placementstrategy["type"] = createservice.PlacementStrategys[i].Type
			}

			placementstrategys = append(placementstrategys, placementstrategy)
		}

		Createservicejsonmap["placementstrategy"] = placementstrategys
	}
}

func preparecreateServiceplacementConstraintsparams(Createservicejsonmap map[string]interface{}, createservice Createservice) {
	if len(createservice.PlacementConstraints) != 0 {
		placementConstraints := make([]map[string]interface{}, 0)
		for i := 0; i < len(createservice.PlacementConstraints); i++ {
			PlacementConstraint := make(map[string]interface{})

			if createservice.PlacementConstraints[i].Expression != "" {
				PlacementConstraint["expression"] = createservice.PlacementConstraints[i].Expression
			}

			if createservice.PlacementConstraints[i].Type != "" {
				PlacementConstraint["type"] = createservice.PlacementConstraints[i].Type
			}

			placementConstraints = append(placementConstraints, PlacementConstraint)
		}

		Createservicejsonmap["placementConstraints"] = placementConstraints
	}
}

func preparecreateServiceloadBalancersparams(Createservicejsonmap map[string]interface{}, createservice Createservice) {
	fmt.Println("len of createservice.LoadBalancers", len(createservice.LoadBalancers))
	if len(createservice.LoadBalancers) != 0 {
		loadBalancers := make([]map[string]interface{}, 0)
		fmt.Println("loadBalancers", loadBalancers)
		for i := 0; i < len(createservice.LoadBalancers); i++ {
			loadBalancer := make(map[string]interface{})

			if createservice.LoadBalancers[i].ContainerName != "" {
				loadBalancer["containerName"] = createservice.LoadBalancers[i].ContainerName
			}

			if createservice.LoadBalancers[i].LoadBalancerName != "" {
				loadBalancer["loadBalancerName"] = createservice.LoadBalancers[i].LoadBalancerName
			}

			if createservice.LoadBalancers[i].TargetGroupArn != "" {
				loadBalancer["targetGroupArn"] = createservice.LoadBalancers[i].TargetGroupArn
			}

			if createservice.LoadBalancers[i].ContainerPort != 0 {
				loadBalancer["containerPort"] = createservice.LoadBalancers[i].ContainerPort
			}

			loadBalancers = append(loadBalancers, loadBalancer)
		}

		Createservicejsonmap["loadBalancers"] = loadBalancers

		fmt.Println("Createservicejsonmap of loadBalancers", Createservicejsonmap["loadBalancers"])
	}
}

func preparecreateServicedeploymentConfigurationparams(Createservicejsonmap map[string]interface{}, createservice Createservice) {
	if (createservice.DeploymentConfigurations != DeploymentConfiguration{}) {
		deploymentConfiguration := make(map[string]interface{})
		if createservice.DeploymentConfigurations.MaximumPercent != 0 {
			deploymentConfiguration["maximumPercent"] = createservice.DeploymentConfigurations.MaximumPercent
		}
		if createservice.DeploymentConfigurations.MinimumHealthyPercent != 0 {
			deploymentConfiguration["minimumHealthyPercent"] = createservice.DeploymentConfigurations.MinimumHealthyPercent
		}
		Createservicejsonmap["deploymentConfiguration"] = deploymentConfiguration
	}
}

func (ecscontainer *Ecscontainer) Createcontainer(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var clusterName, Region string
	for key, value := range param {
		switch key {
		case "clusterName":
			clusterNameV, _ := value.(string)
			clusterName = clusterNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}
	params := make(map[string]string)
	prepareCreatcontainerparams(params, clusterName, Region)

	Creatcontainerjsonmap := map[string]interface{}{
		"clusterName": params["clusterName"],
	}
	ecscontainer.PrepareSignatureV4query(params, Creatcontainerjsonmap)
	return
}

func (ecscontainer *Ecscontainer) Deletecontainer(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var clusterName, Region string
	for key, value := range param {
		switch key {
		case "clusterName":
			clusterNameV, _ := value.(string)
			clusterName = clusterNameV

		case "Region":
			RegionV, _ := value.(string)
			Region = RegionV
		}
	}
	params := make(map[string]string)
	prepareDeletecontainer(params, clusterName, Region)

	Creatcontainerjsonmap := map[string]interface{}{
		"clusterName": params["clusterName"],
	}
	ecscontainer.PrepareSignatureV4query(params, Creatcontainerjsonmap)
	return
}

func prepareDeletecontainer(params map[string]string, clusterName string, Region string) {
	if clusterName != "" {
		params["clusterName"] = clusterName
	}
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.DeleteCluster"
}

func prepareCreatcontainerparams(params map[string]string, clusterName string, Region string) {
	if clusterName != "" {
		params["clusterName"] = clusterName
	}
	if Region != "" {
		params["Region"] = Region
	}
	params["amztarget"] = "AmazonEC2ContainerServiceV20141113.CreateCluster"
}

func (ecscontainer *Ecscontainer) PrepareSignatureV4query(params map[string]string, paramsmap map[string]interface{}) {
	fmt.Println(paramsmap)
	ECSEndpoint := "https://ecs." + params["Region"] + ".amazonaws.com"
	service := "ecs"
	method := "POST"
	host := service + "." + params["Region"] + ".amazonaws.com"
	ContentType := "application/x-amz-json-1.1"
	signedheaders := "content-type;host;x-amz-date;x-amz-target"
	amztarget := params["amztarget"]

	requestparametersjson, _ := json.Marshal(paramsmap)
	requestparametersjsonstring := string(requestparametersjson)
	requestparametersjsonstringbyte := []byte(requestparametersjsonstring)
	fmt.Println("requestparametersjsonstring", requestparametersjsonstring)
	client := new(http.Client)
	request, _ := http.NewRequest("POST", ECSEndpoint, bytes.NewBuffer(requestparametersjsonstringbyte))
	request = awsauth.SignatureV4(request, requestparametersjsonstringbyte, amztarget, method, params["Region"], service, host, ContentType, signedheaders)
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("resp Body:", string(body))
}
