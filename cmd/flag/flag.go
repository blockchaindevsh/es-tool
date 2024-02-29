package flag

import "github.com/urfave/cli"

var NetworkFlag = cli.StringFlag{
	Name:  "network",
	Usage: "specify network",
}

var PKFlag = cli.StringFlag{
	Name:  "pk",
	Usage: "specify private key",
}

var ContractFlag = cli.StringFlag{
	Name:  "contract",
	Usage: "specify contract address",
}

var BlobFlag = cli.StringFlag{
	Name:  "blob",
	Usage: "specify blob hash",
}

var SpanFlag = cli.BoolFlag{
	Name:  "span",
	Usage: "specify span flag",
}
