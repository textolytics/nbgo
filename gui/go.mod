module gui

go 1.22.3

require (
	cli v0.0.0
	conf v0.0.0
	core v0.0.0
	logs v0.0.0
)

require gopkg.in/yaml.v2 v2.4.0 // indirect

replace (
	cli => ../cli
	conf => ../conf
	core => ../core
	logs => ../logs
)
