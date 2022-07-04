module github.com/eodnozerkin-ozon/allure-go-param

go 1.17

replace (
	github.com/eodnozerkin-ozon/allure-go-param/pkg/allure => ./pkg/allure
	github.com/eodnozerkin-ozon/allure-go-param/pkg/framework => ./pkg/framework
)

require (
	github.com/eodnozerkin-ozon/allure-go-param/pkg/allure v1.1.0
	github.com/eodnozerkin-ozon/allure-go-param/pkg/framework v1.3.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
