package dto

type PredictionDataOfCSVRequest struct {
	OperatingSystem      int    `json:"Operating System"`
	AnalysisOfAlgorithm  int    `json:"Analysis of Algorithm"`
	ProgrammingConcept   int    `json:"Programming Concept"`
	SoftwareEngineering  int    `json:"Software Engineering"`
	ComputerNetwork      int    `json:"Computer Network"`
	AppliedMathematics   int    `json:"Applied Mathematics"`
	ComputerSecurity     int    `json:"Computer Security"`
	HackathonsAttended   int    `json:"Hackathons attended"`
	TopmostCertification string `json:"Topmost Certification"`
	Personality          string `json:"Personality"`
	ManagementTechnical  string `json:"Management or technical"`
	Leadership           string `json:"Leadership"`
	Team                 string `json:"Team"`
	SelfAbility          string `json:"Self Ability"`
}
