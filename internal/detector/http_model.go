package detector

type Input []string

type SimpleOutput []string

type ExtendedOutput map[string][]string

type PostTaskOutput struct {
	UUID  string `json:"uuid"`
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

type GetTaskOutput struct {
	UUID   string          `json:"uuid"`
	OK     bool            `json:"ok"`
	Error  string          `json:"error"`
	Result map[string]bool `json:"result,omitempty"`
}
