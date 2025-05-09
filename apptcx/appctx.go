package apptcx

import "server/database"

var ConnectDB *AppContext

type AppContext struct {
	PostgresConnectors map[string]database.PostgresConnectorInterface
	MongoConnectors    map[string]database.MongoConnectorInterface
}
