module gui

go 1.22.3

require (
	cli v0.0.0
	conf v0.0.0
	core v0.0.0
	logs v0.0.0
)

replace (
	cli => ../cli
	conf => ../conf
	core => ../core
	logs => ../logs
)
