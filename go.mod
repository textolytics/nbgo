module github.com/textolytics/nbgo

go 1.22.3

require (
	github.com/gansidui/skiplist v0.0.0-20141121051332-c6a909ce563b
	github.com/gateio/gatews/go v0.0.0-20240430073619-0a9f3ff9f49e
	github.com/rs/zerolog v1.33.0
	github.com/shopspring/decimal v1.4.0
	github.com/stretchr/testify v1.9.0
	github.com/yireyun/go-queue v0.0.0-20220725040158-a4dd64810e1e
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/textolytics/nbgo/cli => ./cli
	github.com/textolytics/nbgo/conf => ./conf
	github.com/textolytics/nbgo/core => ./core
	github.com/textolytics/nbgo/dw => ./dw
	github.com/textolytics/nbgo/gw => ./gw
	github.com/textolytics/nbgo/logs => ./logs
	github.com/textolytics/nbgo/mb => ./mb
	github.com/textolytics/nbgo/mon => ./mon
	github.com/textolytics/nbgo/run => ./run
	github.com/textolytics/nbgo/task => ./task
)

require (
	github.com/textolytics/nbgo/cli v0.0.0-20260120222905-8cbb55391d93 // indirect
	github.com/textolytics/nbgo/core v0.0.0-20260120222905-8cbb55391d93 // indirect
	github.com/textolytics/nbgo/gw v0.0.0-20260120222905-8cbb55391d93 // indirect
	github.com/textolytics/nbgo/logs v0.0.0-20260120222905-8cbb55391d93 // indirect
	github.com/textolytics/nbgo/mb v0.0.0-20260120222905-8cbb55391d93 // indirect
	github.com/textolytics/nbgo/mon v0.0.0-20260120222905-8cbb55391d93 // indirect
	github.com/textolytics/nbgo/run v0.0.0-20260120222905-8cbb55391d93 // indirect
	github.com/textolytics/nbgo/task v0.0.0-20260120222905-8cbb55391d93 // indirect
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/textolytics/nbgo/conf v0.0.0-20260120222905-8cbb55391d93
	github.com/textolytics/nbgo/dw v0.0.0-20260120222905-8cbb55391d93
	golang.org/x/sys v0.12.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
