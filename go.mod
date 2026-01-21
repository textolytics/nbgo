module github.com/textolytics/nbgo

go 1.25.5

require (
	github.com/textolytics/nbgo/cli v0.0.0
	github.com/textolytics/nbgo/conf v0.0.0
	github.com/textolytics/nbgo/core v0.0.0
	github.com/textolytics/nbgo/dw v0.0.0-20260121013637-a2a097c238fd
	github.com/textolytics/nbgo/gw v0.0.0-20260121013637-a2a097c238fd
	github.com/textolytics/nbgo/logs v0.0.0
	github.com/textolytics/nbgo/mb v0.0.0-20260121013637-a2a097c238fd
	github.com/textolytics/nbgo/mon v0.0.0-20260121013637-a2a097c238fd
	github.com/textolytics/nbgo/run v0.0.0-20260121013637-a2a097c238fd
	github.com/textolytics/nbgo/task v0.0.0-20260121013637-a2a097c238fd
)

require (
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/gateio/gatews/go v0.0.0-20250523113507-90357b11b694 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/textolytics/nbgo/cli => ./cli
	github.com/textolytics/nbgo/conf => ./conf
	github.com/textolytics/nbgo/core => ./core
	github.com/textolytics/nbgo/dw => ./dw
	github.com/textolytics/nbgo/gui => ./gui
	github.com/textolytics/nbgo/gw => ./gw
	github.com/textolytics/nbgo/logs => ./logs
	github.com/textolytics/nbgo/mb => ./mb
	github.com/textolytics/nbgo/mon => ./mon
	github.com/textolytics/nbgo/run => ./run
	github.com/textolytics/nbgo/task => ./task
)
