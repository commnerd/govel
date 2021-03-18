package new

const GOMOD = `module %s

go 1.15

replace (
	govel => ../../
	%s => ./
)

require (
	github.com/stretchr/testify v1.7.0
)

`
