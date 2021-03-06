package new

const GOMOD = `
module %s

go 1.15

replace (
	%s => ./
	github.com/commnerd/govel => ../../
)
`
