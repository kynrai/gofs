module github.com/atos-digital/ttz

go 1.22.4

require (
	golang.org/x/mod v0.18.0
	golang.org/x/tools v0.22.0
	module/placeholder v0.0.0
)

require (
	github.com/go-http-utils/etag v0.0.0-20161124023236-513ea8f21eb1 // indirect
	github.com/go-http-utils/fresh v0.0.0-20161124030543-7231e26a4b27 // indirect
	github.com/go-http-utils/headers v0.0.0-20181008091004-fed159eddc2a // indirect
)

replace module/placeholder => ./template
