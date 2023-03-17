package main

import (
    "fmt"
	"os"
	"sort"
    "gopkg.in/yaml.v3"
	"golang.org/x/exp/slices"
)

/* 
	gopkg.in/yaml.v3 seems to require the YAML file have a top-level.
	I'm manually adding 'slos' to the yaml provided by Nobl9. 
	Idk how to get around this.
*/

type SLOs struct {
    SLO[]				 SLO 		`yaml:"slos"`
}

type SLO struct {
    Metadata 						`yaml:"metadata"`
    Spec 							`yaml:"spec"`
}

// Nested under SLO
type Metadata struct {
    DisplayName 		string 		`yaml:"displayName"`
    Name 				string 		`yaml:"name"`
    Project 			string 		`yaml:"project"`
    Label 							`yaml:"labels"`
}

// Nested under SLO
type Spec struct {
    AlertPolicies[] 	string 		`yaml:"alertPolicies"`
    Description 		string 		`yaml:"description"`
    Attachments 					`yaml:"attachments"`
}

// Nested under Metadata
type Label struct {
    AnzxTechAsset[] 	string 		`yaml:"anz-x-tech-asset"`
    AnzxValueStream[] 	string 		`yaml:"anz-x-value-stream"`
    Component[] 		string 		`yaml:"component"`
    DeployedFrom[] 		string 		`yaml:"deployed-from"`
    SloType[] 			string 		`yaml:"slo-type"`
    TeamName[] 			string 		`yaml:"team-name"`
    UserJourney[] 		string 		`yaml:"user-journey"`
}

// Nested under Spec
type Attachments[] struct {
    DisplayName 		string 		`yaml:"displayName"`
    URL 				string 		`yaml:"url"`
}

// Used for created slices of components
type Component struct {
    ComponentName 		string
    ComponentCount		int
}

func readYamlFile (file_path string) ([]byte) {
	data, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	return data
}

func getUniqueComponentsList(slos SLOs) []string {
	var UniqueComponentsList []string
	for _, slo := range slos.SLO {
		if !(slices.Contains(UniqueComponentsList, slo.Metadata.Label.Component[0])) {
			UniqueComponentsList = append(UniqueComponentsList, slo.Metadata.Label.Component[0])
		}
	}
	sort.Strings(UniqueComponentsList)
	return UniqueComponentsList
}

func printAllSlos(slos SLOs) {
	// fmt.Println(slos)
	for _, slo := range slos.SLO {
		// fmt.Println(slo)
		fmt.Println("Name:", slo.Metadata.DisplayName)
		// fmt.Println("Labels:")
		// fmt.Println("    Tech Asset:", slo.Metadata.Label.AnzxTechAsset[0])
		// fmt.Println("    Value Stream:", slo.Metadata.Label.AnzxValueStream[0])
		// fmt.Println("    Component:", slo.Metadata.Label.Component[0])
		// fmt.Println("    Deployed From:", slo.Metadata.Label.DeployedFrom[0])
		// fmt.Println("    Slo Type:", slo.Metadata.Label.SloType[0])
		// fmt.Println("    Team Name:", slo.Metadata.Label.TeamName[0])
		// fmt.Println("    User Journey:", slo.Metadata.Label.UserJourney[0])
		// fmt.Println("Alert Policies:")
		// fmt.Println("    ", slo.Spec.AlertPolicies)
		// fmt.Println("Attachments:")
		// fmt.Println("    ", slo.Spec.Attachments[0].DisplayName, "-", slo.Spec.Attachments[0].URL)
		// fmt.Println("Description:")
		// fmt.Println("    ", slo.Spec.Description)
		// fmt.Println("")
	}
}

func printCountByComponent(slos SLOs, uniqueComponentsList []string) {
	fmt.Println("SLO Counts by Component")
	for _, uniqueComponent := range uniqueComponentsList {
		counter := 0
		// fmt.Println(uniqueComponent) 
		for _, slo := range slos.SLO {
			if uniqueComponent == slo.Metadata.Label.Component[0] {
				counter += 1
			}
		}
		fmt.Println(" ", uniqueComponent,":", counter)
	}
}

func main() {
	var slos SLOs
	data := readYamlFile("slos.yaml")
	unmarshalErr := yaml.Unmarshal(data, &slos)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	printAllSlos(slos)
	var uniqueComponentsList = getUniqueComponentsList(slos)
	printCountByComponent(slos, uniqueComponentsList)
}