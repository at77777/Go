package event

type Event struct {
	Id   int64  `db:"id,omitempty"`
	Name string `db:"name"`
}
