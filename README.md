Just a boilerplate go application for learning purposes.. Some things may or may not be done appropriately

### Install Deps
[gopm](https://github.com/gpmgo/gopm):
`$ go get -u github.com/gpmgo/gopm`

[godo](https://github.com/go-godo/godo):
`$ go get -u gopkg.in/godo.v1/cmd/godo`

[goose](https://bitbucket.org/liamstask/goose):
`$ go get bitbucket.org/liamstask/goose/cmd/goose`

### Bootstrap
`godo bootstrap`

### Run
One command to build, run and watch
`$ godo --watch`

### Migrations
Create migration using go:
`$ goose create ${MigrationName}`

Create migration using sql:
`$ goose create ${MigrationName} sql`

Migrate up:
`$ goose up`

... And down
`$ goose down`
