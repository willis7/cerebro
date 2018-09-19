package environment

type Environment struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

var schema = `
CREATE TABLE environment (
	id int
    name text
)`
