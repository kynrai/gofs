module github.com/atos-digital/ttz

go 1.23

require (
	github.com/a-h/templ v0.2.747
	golang.org/x/mod v0.18.0
	golang.org/x/tools v0.22.0
	module/placeholder v0.0.0
)

require github.com/a-h/parse v0.0.0-20240121214402-3caf7543159a // indirect

replace module/placeholder => ./template
