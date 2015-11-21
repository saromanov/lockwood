package text

type Text interface {

	// PreRun provides preprocessing of the data
	PreRun(data string) string

	// Run provides main functional
	Run(data string)
}