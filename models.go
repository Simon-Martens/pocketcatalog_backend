package main

type DeletionEntry struct {
	Table  string `json:"table"`
	Record string `json:"record"`
	Field  string `json:"field"`
}

type DeletionDependencies struct {
	Table string   `json:"table"`
	Field []string `json:"field"`
}

type DependencyGraph struct {
	Table                string                 `json:"table"`
	OptionalDependencies []DeletionDependencies `json:"optional_dependencies"`
	RequiredDependencies []DeletionDependencies `json:"required_dependencies"`
}
