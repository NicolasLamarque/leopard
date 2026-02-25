package importer

// Pipeline est la structure centrale
type Pipeline struct {
	FilePath string
	Config   ImportConfig
}

type ImportConfig struct {
	DateFormat string
	SkipHeader bool
}

// NewPipeline initialise l'objet
func NewPipeline(path string) *Pipeline {
	return &Pipeline{
		FilePath: path,
		Config: ImportConfig{
			DateFormat: "2006-01-02",
			SkipHeader: true,
		},
	}
}
