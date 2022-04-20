module github.com/funcgql/cli

go 1.18

require (
	github.com/neakor/cloud-functions-gql v0.0.0-20220406024313-da28c0ad8d76
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
)

exclude (
	github.com/funcgql/cli/template/fixtures/mod_1 v0.0.0
	github.com/funcgql/cli/template/fixtures/mod_2 v0.0.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
