package services

import "gmmshops.go/app/providers/logs"

// Services Providers
type Services struct {
	_s map[string]interface{}
}

var s *Services

func init() {
	if s == nil {
		s = &Services{_s: make(map[string]interface{})}
	}

}

// Register : register service
func Register(name string, service interface{}) {
	s._s[name] = service
	logs.Logger().Info("Service: ", name, " ► [Registered]")
}

// GetServices :
func GetServices(name string) interface{} {
	return s._s[name]
}
