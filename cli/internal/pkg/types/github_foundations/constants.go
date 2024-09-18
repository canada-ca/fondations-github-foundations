package githubfoundations

type ResourceType int

const (
	None                   ResourceType = iota
	Repository                          = iota
	RepositoryCollaborator              = iota
)
