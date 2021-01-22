package gdbc

import (
	"database/sql"
	"github.com/go-gdbc/gdbc"
	"github.com/procyon-projects/procyon-configure"
)

type ConnectionProvider interface {
	GetDatabaseConnection() *sql.DB
}

type simpleConnectionProvider struct {
	databaseConnection *sql.DB
}

func newSimpleDatabaseConnectionProvider(dataSourceProperties configure.DataSourceProperties) *simpleConnectionProvider {
	provider := &simpleConnectionProvider{}
	provider.initDatabaseConnection(dataSourceProperties)
	return provider
}

func (connectionProvider *simpleConnectionProvider) initDatabaseConnection(properties configure.DataSourceProperties) {
	dataSource, err := gdbc.GetDataSource(properties.URL,
		gdbc.Username(properties.Username),
		gdbc.Password(properties.Password),
	)
	if err != nil {
		panic("Could not get the data source : " + err.Error())
	}

	var databaseConnection *sql.DB
	databaseConnection, err = dataSource.GetConnection()
	if err != nil {
		panic("Could not create the database connection : " + err.Error())
	}
	connectionProvider.databaseConnection = databaseConnection
}

func (connectionProvider *simpleConnectionProvider) GetDatabaseConnection() *sql.DB {
	return connectionProvider.databaseConnection
}
