package detector

//go:generate moq -out storage_moq_test.go . Storage

type Storage interface {
	Set(uuid string, result interface{}) error
	Get(uuid string) (interface{}, error)
}

func (s *Service) getScanInfo(uuid string) (ScanInfo, error) {
	data, err := s.output.Get(uuid)
	if err != nil {
		return ScanInfo{}, err
	}

	return data.(ScanInfo), nil
}
