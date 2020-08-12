package db

import (
	"applib/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "go.mongodb.org/mongo-driver/mongo"
)

const (
	GORMDB_DIALECT_MSSQL = "mssql"
	GORMDB_DIALECT_MYSQL = "mysql"
	GORMDB_DIALECT_POSTGRES = "postgres"
)

type GormDb struct {
	dbConfs map[string]conf.DbConf
	connections map[string]*gorm.DB
}

func(db *GormDb) Init(dbConfs map[string]conf.DbConf) {
	db.dbConfs = dbConfs
	db.connections = map[string]*gorm.DB{}
	for connectionID, conf := range db.dbConfs {
		conn, err := db.CreateConnection(conf)
		if err!=nil {
			panic(err.Error())
		}
		db.connections[connectionID] = conn
	}
}

func(db *GormDb) SetConnection(connectionID string,db2 *gorm.DB) {
	db.connections[connectionID] = db2
}

func(db GormDb) GetConnection(connectionID string) *gorm.DB {
	conn, e := db.connections[connectionID]
	if !e {
		panic("cannot_find_connection_id=" + connectionID)
	}
	return conn
}

func (db GormDb) createConnectionString(dialect string, dbConf conf.DbConf) string {

	if dialect == GORMDB_DIALECT_MSSQL {
		connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", dbConf.User, dbConf.Passwd,
			dbConf.Server, dbConf.Port, dbConf.DbName)
		return connString

	} else if dialect == GORMDB_DIALECT_MYSQL {

		connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			dbConf.User, dbConf.Passwd, dbConf.Server, dbConf.Port, dbConf.DbName)

		return connString

	} else if dialect == GORMDB_DIALECT_POSTGRES {
		connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			dbConf.Server, dbConf.Port, dbConf.User, dbConf.DbName, dbConf.Passwd)

		return connString

	} else {
		panic("Unknown DbType=" + dbConf.DbType)
	}
}

func (db GormDb) parseDialect(dbType string) string {
	dialect := strings.ToLower(dbType)
	if ! (dialect == GORMDB_DIALECT_MSSQL ||
		dialect == GORMDB_DIALECT_MYSQL ||
		dialect == GORMDB_DIALECT_POSTGRES)  {
		panic("invalid DbType. DbType must be in" +
			GORMDB_DIALECT_MSSQL + " or " +
			GORMDB_DIALECT_MYSQL + " or " +
			GORMDB_DIALECT_POSTGRES)
	}
	return dialect
}

func (db GormDb) CreateConnection(dbConf conf.DbConf) (*gorm.DB, error){
	dialect := db.parseDialect(dbConf.DbType)
	connection, err := gorm.Open(dialect, db.createConnectionString(dialect, dbConf))
	if err != nil {
		return nil, err
	}
	if dbConf.Debug {
		connection.LogMode(true)
	}
	connection.SingularTable(dbConf.Singular)
	return connection, nil
}


