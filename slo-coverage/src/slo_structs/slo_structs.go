package slo_structs

type SLOs struct {
    SLO[] SLO `yaml:"slos"`
}

type SLO struct {
    Metadata `yaml:"metadata"`
    Spec `yaml:"spec"`
}

// Nested under SLO
type Metadata struct {
    DisplayName string `yaml:"displayName"`
    Name string `yaml:"name"`
    Project string `yaml:"project"`
    Label `yaml:"labels"`
}

// Nested under SLO
type Spec struct {
    AlertPolicies[] string `yaml:"alertPolicies"`
    Description string `yaml:"description"`
    Attachments `yaml:"attachments"`
}

// Nested under Metadata
type Label struct {
    AnzxTechAsset[] string `yaml:"anz-x-tech-asset"`
    AnzxValueStream[] string `yaml:"anz-x-value-stream"`
    Component[] string `yaml:"component"`
    DeployedFrom[] string `yaml:"deployed-from"`
    SloType[] string `yaml:"slo-type"`
    TeamName[] string `yaml:"team-name"`
    UserJourney[] string `yaml:"user-journey"`
    XploreJobs[] string `yaml:"xplore-jobs"`
}

// Nested under Spec
type Attachments[] struct {
    DisplayName string `yaml:"displayName"`
    URL string `yaml:"url"`
}

// Used for created slices of components
// WHAT THIS FOR?
type Component struct {
    ComponentName string
    ComponentCount int
}
