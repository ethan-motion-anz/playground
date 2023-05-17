package main

import (
    "fmt"
	"os"
    "gopkg.in/yaml.v3"
	"src/goal_structs"
	"src/slo_structs"
)

var goals_file string = "../data/goals.json"
var slos_file string = "../data/slos.yaml"


func readFile (file_path string) ([]byte) {
	data, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	return data
}

func printArrayElements(array []string) {
	for _, elem := range array {
		fmt.Println(elem)
	}
}

func printAllGoals(goals goal_structs.Goals) {
	for _, goal := range goals.Data.Goals {
		fmt.Printf("GOAL: %v (%v)\n", goal.Name, goal.ID)
		for _, episode := range goal.Episodes {
			fmt.Printf("  EPISODE: %v (%v)\n", episode.Name, episode.ID)
			for _, job := range episode.Jobs {
				fmt.Printf("    JOB: %v (%v)\n", job.Name, job.ID)
			}
		}	
	}
}

func writeGoalJobsToArray(goals goal_structs.Goals, inputGoal string) []string {
	goalJobsArray := []string {}
	for _, goal := range goals.Data.Goals {
		if goal.Name == inputGoal && goal.ArchivedAt == nil {
			for _, episode := range goal.Episodes {
				if episode.ArchivedAt == nil {
					for _, job := range episode.Jobs {
						if job.ArchivedAt == nil {
							// fmt.Println(job.ID)
							goalJobsArray = append(goalJobsArray, job.ID)
						}
					}
				}
			}	
		}
	}
	return goalJobsArray
}

func writeFabricSloLabelsToArray(slos slo_structs.SLOs) []string {
	fabricSloLabelsArray := []string {}
	for _, slo := range slos.SLO {
		for _, xploreJob := range slo.Label.XploreJobs {
			// fmt.Println("Xplore Jobs:", xploreJob)
			fabricSloLabelsArray = append(fabricSloLabelsArray, xploreJob)
		}
	}
	return fabricSloLabelsArray
}

func doTheMatching(mmmGoalJobsArrayInput []string, fabricSloXploreJobsArrayInput []string) {
	var xploreJobsSlod int = 0
	for _, xploreJob := range mmmGoalJobsArrayInput {
		for _, sloJob := range fabricSloXploreJobsArrayInput {
			if xploreJob == sloJob {
				fmt.Println("MATCHED", xploreJob)
				xploreJobsSlod++
			}
		}
	}
	fmt.Printf("%v/%v MMM Xplore Jobs are referenced by SLOs", xploreJobsSlod, len(mmmGoalJobsArrayInput))
}


func main() {
	var goals goal_structs.Goals
	goals_data := readFile(goals_file)
	unmarshalGoalsErr := yaml.Unmarshal(goals_data, &goals) // this makes no sense
	if unmarshalGoalsErr != nil {
		panic(unmarshalGoalsErr)
	}

	mmmGoalJobsArray := writeGoalJobsToArray(goals, "Manage my money & Save")
	printAllGoals(goals)
	// printArrayElements(mmmGoalJobsArray)

	var slos slo_structs.SLOs
	slos_data := readFile(slos_file)
	unmarshalSlosErr := yaml.Unmarshal(slos_data, &slos)
	if unmarshalSlosErr != nil {
		panic(unmarshalSlosErr)
	}

	fabricSloXploreJobsArray := writeFabricSloLabelsToArray(slos)
	// printArrayElements(fabricSloXploreJobsArray)

	doTheMatching(mmmGoalJobsArray, fabricSloXploreJobsArray)
	
}