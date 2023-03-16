package main

import (
    "fmt"
    "io/ioutil"
    "gopkg.in/yaml.v3"
)

// Require the YAML file have a top-level (I'm using 'slos').
// Idk how to get around this.
type SLOs struct {
	SLO					[]SLO		`slos`		// Why is this a warning??
}

type SLO struct {
    ApiVersion       	string 		`yaml:"apiVersion"` 	// Testing/remove
    Kind 				string 		`yaml:"kind"` 			// Testing/remove
	Metadata			 	        `yaml:"metadata"`
	Spec							`yaml:"spec"`
}

type Metadata struct {
	DisplayName			string		`yaml:"displayName"`
	Label struct {
		AnzxTechAsset 		[]string 	`yaml:"anz-x-tech-asset"`
		AnzxValueStream		[]string 	`yaml:"anz-x-value-stream"`
		Component			[]string 	`yaml:"component"`
		DeployedFrom		[]string 	`yaml:"deployed-from"`
		SloType				[]string 	`yaml:"slo-type"`
		TeamName			[]string 	`yaml:"team-name"`
		UserJourney			[]string 	`yaml:"user-journey"`
	} `yaml:"labels"`
	Name				string		`yaml:"name"`
	Project				string		`yaml:"project"`
}	

type Spec struct {
	AlertPolicies 		[]string 	`yaml:"alertPolicies"`
	Attachments   		[]struct {
		DisplayName 		string 		`yaml:"displayName"`
		URL         		string 		`yaml:"url"`
	} `yaml:"attachments"`
	Description			string		`yaml:"description"`
}

func main() {

	var slos SLOs

	data, err := ioutil.ReadFile("slos.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = yaml.Unmarshal(data, &slos)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(slos)

	for _, slo := range slos.SLO {
		fmt.Println("Name:", slo.Metadata.DisplayName)
		fmt.Println("Labels:")
		fmt.Println("    Tech Asset:", slo.Metadata.Label.AnzxTechAsset[0])
		fmt.Println("    Value Stream:", slo.Metadata.Label.AnzxValueStream[0])
		fmt.Println("    Component:", slo.Metadata.Label.Component[0])
		fmt.Println("    Deployed From:", slo.Metadata.Label.DeployedFrom[0])
		fmt.Println("    Slo Type:", slo.Metadata.Label.SloType[0])
		fmt.Println("    Team Name:", slo.Metadata.Label.TeamName[0])
		fmt.Println("    User Journey:", slo.Metadata.Label.UserJourney[0])
		fmt.Println("Alert Policies:")
		fmt.Println("    ", slo.Spec.AlertPolicies)
		fmt.Println("Attachments:")
		fmt.Println("    ", slo.Spec.Attachments[0].DisplayName, "-", slo.Spec.Attachments[0].URL)
		fmt.Println("Description:")
		fmt.Println("    ", slo.Spec.Description)
		fmt.Println("")
		fmt.Println("")
	}

}