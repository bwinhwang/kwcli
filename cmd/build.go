package cmd

var (
	describe = "Retrieve the list of builds for a project."
)

type buildsAction struct {
	Action  string `json:"action"` // Required
	Project string `json:"project,omitempty"`
}

func NewbuildsAction(project string) *buildsAction {
	return &buildsAction{
		Action:  "builds",
		Project: project,
	}
}
