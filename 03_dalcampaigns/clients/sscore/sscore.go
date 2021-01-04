package sscore

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type Deps struct{}

type Config struct {
	Address string
	Username string
	Password string
	DBName string
}

type SSCore interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Close() error
}

type impl struct {
	DB *sql.DB
}

func New(deps *Deps, config *Config) SSCore {
	dbConfig := mysql.Config{
		Addr: config.Address,
		User: config.Username,
		Passwd: config.Password,
		DBName: config.DBName,
	}
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE ss_core_dev")
	if err != nil {
		panic(err)
	}

	return &impl{DB: db}
}

func (impl *impl) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return impl.DB.Query(query, args...)
}

func (impl *impl) Close() error {
	return impl.DB.Close()
}
