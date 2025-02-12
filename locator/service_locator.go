package locator

import "goproject/service"

type ServiceLocator struct {
	cache *Cache
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{cache: NewCache()}
}

func (serviceLocator *ServiceLocator) Locate(serviceName string) service.Service {
	svc := serviceLocator.cache.GetService(serviceName)
	if svc != nil {
		return svc
	}

	context := InitialContext{}
	svc = context.Lookup(serviceName)
	if svc != nil {
		serviceLocator.cache.AddService(svc)
	}
	return svc
}
