package locator

import "goproject/service"

type Cache struct {
	services map[string]service.Service
}

func NewCache() *Cache {
	return &Cache{services: make(map[string]service.Service)}
}

func (c *Cache) GetService(serviceName string) service.Service {
	if service, ok := c.services[serviceName]; ok {
		return service
	}
	return nil
}

func (c *Cache) AddService(service service.Service) {
	c.services[service.GetName()] = service
}
