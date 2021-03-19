package data

import (
	"database/sql"
	"fmt"

	// MySQL driver
	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nkarpenko/playlog-test/api/conf"
	"github.com/nkarpenko/playlog-test/api/data/comment"
	"github.com/nkarpenko/playlog-test/api/data/user"
)

// Data interface
type Data interface {
	Close() error

	// Data subservices
	Comment() comment.Service
	User() user.Service
}

type data struct {
	storage *sql.DB
	comment comment.Service
	user    user.Service
}

// Close any open db connections.
func (d *data) Close() error {
	// If mysql db connecton is open, close it.
	err := d.storage.Close()
	if err != nil {
		log.Infof("Error closing MySQL connection: %s", err.Error())
	}
	return nil
}

func (d *data) Comment() comment.Service {
	return d.comment
}

func (d *data) User() user.Service {
	return d.user
}

// New database instance
func New(msc *conf.DBConfig) (Data, error) {
	// Establish mysql db connection.
	msConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", msc.User, msc.Password, msc.Host, msc.Port, msc.Database)
	st, err := sql.Open("mysql", msConnString)
	if err != nil {
		log.Warnf("MySQL - Unable to connect: %s", err.Error())
		return &data{}, err
	}

	data := &data{
		storage: st,
		comment: comment.New(st),
		user:    user.New(st),
	}

	return data, nil
}
