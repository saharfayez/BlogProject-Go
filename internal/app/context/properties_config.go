package context

type PropertiesConfig interface {
	GetProfile() string
	GetDatabaseUrl() string
}
