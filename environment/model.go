package environment

type Environment struct {
	Name string `db:"name"`
}

var schema = `
CREATE TABLE environment (
    name text
)`
