package interpreter

type Script struct {
	Branches []Branch
}

type ScriptOutput struct {
	StoryID         uint
	HasResolutionID bool
	ResolutionID    uint
}

func (script Script) Execute() ScriptOutput {
	for _, branch := range script.Branches {
		if branch.condition.Evaluate() {
			return ScriptOutput{
				branch.storyID,
				branch.hasResolutionID,
				branch.resolutionID,
			}
		}
	}

	return ScriptOutput{}
}
