module github.com/micro/services/portfolio/trades

go 1.12

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.10
	github.com/micro/services/portfolio/helpers/iex-cloud v1.0.0
	github.com/micro/services/portfolio/helpers/microgorm v1.0.0
	github.com/micro/services/portfolio/helpers/microtime v0.0.0-00010101000000-000000000000
	github.com/micro/services/portfolio/ledger v1.0.0
	github.com/micro/services/portfolio/stocks v1.0.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	github.com/pkg/errors v0.8.1
	github.com/satori/go.uuid v1.2.0
)

replace github.com/micro/services/portfolio/helpers/microtime => ../helpers/microtime

replace github.com/micro/services/portfolio/helpers/microgorm => ../helpers/microgorm

replace github.com/micro/services/portfolio/helpers/iex-cloud => ../helpers/iex-cloud

replace github.com/micro/services/portfolio/ledger => ../ledger

replace github.com/micro/services/portfolio/stocks => ../stocks

replace github.com/micro/services/portfolio/portfolios => ../portfolios
