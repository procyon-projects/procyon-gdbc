package gdbc

import (
	"database/sql"
	"github.com/go-gdbc/gdbc"
	"github.com/procyon-projects/procyon-configure"
	"github.com/procyon-projects/procyon-context"
)

type ConnectionProvider interface {
	GetDatabaseConnection() *sql.DB
}

type simpleConnectionProvider struct {
	logger             context.Logger
	databaseConnection *sql.DB
}

func newSimpleDatabaseConnectionProvider(logger context.Logger,
	dataSourceProperties configure.DataSourceProperties) *simpleConnectionProvider {
	provider := &simpleConnectionProvider{
		logger,
		nil,
	}
	provider.initDatabaseConnection(dataSourceProperties)
	return provider
}

func (connectionProvider *simpleConnectionProvider) initDatabaseConnection(properties configure.DataSourceProperties) {
	dataSource, err := gdbc.GetDataSource(properties.URL,
		gdbc.Username(properties.Username),
		gdbc.Password(properties.Password),
	)
	if err != nil {
		connectionProvider.logger.Fatal(nil, "Could not create the data source : "+err.Error())
		return
	}

	var databaseConnection *sql.DB
	databaseConnection, err = dataSource.GetConnection()
	if err != nil {
		connectionProvider.logger.Fatal(nil, "Could not create database connection : "+err.Error())
		return
	}
	connectionProvider.databaseConnection = databaseConnection
}

func (connectionProvider *simpleConnectionProvider) GetDatabaseConnection() *sql.DB {
	return connectionProvider.databaseConnection
}
