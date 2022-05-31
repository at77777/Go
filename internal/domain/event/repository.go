package event

import (
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Create(id int64, name string) error
	Update(id int64, name string) error
	Delete(id int64, name string) error
}

const EventsCount int64 = 10

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

var settings = postgresql.ConnectionURL{
	Database: `training`,
	Host:     `localhost:54322`,
	User:     `postgres`,
	Password: `086Teq2`,
}

func (r *repository) FindAll() ([]Event, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	var events []Event

	q, err := sess.SQL().Query(`SELECT * FROM events ORDER BY id`)
	i := sess.SQL().NewIterator(q)
	if err := i.All(&events); err != nil {
		log.Fatal("Query: ", err)
	}

	return events, nil
}

func (r *repository) FindOne(id int64) (*Event, error) {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	var event *Event

	q, err := sess.SQL().Query(`SELECT * FROM events WHERE id = ($1)`, id)
	i := sess.SQL().NewIterator(q)
	if err := i.One(&event); err != nil {
		log.Fatal("Query: ", err)
	}

	return event, nil

}

func (r *repository) Create(id int64, name string) error {

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	_, err = sess.SQL().Exec(`INSERT INTO events (id, name) VALUES ($1, $2)`, id, name)
	if err != nil {
		log.Fatal("Query: ", err)
	}

	return nil
}

func (r *repository) Update(id int64, name string) error {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	_, err = sess.SQL().Exec(`UPDATE events SET name = ($1) WHERE id = ($2)`, name, id)
	if err != nil {
		log.Fatal("Query: ", err)
	}

	return nil

}

func (r *repository) Delete(id int64, name string) error {
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	_, err = sess.SQL().Exec(`DELETE FROM events WHERE id = ($1) AND name = ($2)`, id, name)
	if err != nil {
		log.Fatal("Query: ", err)
	}

	return nil
}
