package detector

//go:generate moq -out probe_moq_test.go . Probe

type Probe interface {
	Run(host string) bool
}
