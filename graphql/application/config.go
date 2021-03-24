package application

type Config struct {
	Port             int
	ProductTableName string
	DB               struct {
		DriverName     string
		DataSourceName string
	}
	GatewayEndpoint string
	EmailSender     string
}
