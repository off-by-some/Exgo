Just a boilerplate go application for learning purposes.. Some things may or may not be done appropriately

### Install Deps
[gopm](https://github.com/gpmgo/gopm):
`go get -u github.com/gpmgo/gopm`

[goose](https://bitbucket.org/liamstask/goose):
`go get bitbucket.org/liamstask/goose/cmd/goose`


### Migrations
Create migration using go:
`$ goose create ${MigrationName}`

Create migration using sql:
`$ goose create ${MigrationName} sql`

Migrate up:
`$ goose up`

... And down
`$ goose down`

### Bootstrap
`gopm get`

### Build
`gopm build`

### Run
`./output`
