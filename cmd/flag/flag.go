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
	Usage: "specify blob hash or blob preimage depending on context",
}

var BlobFileFlag = cli.StringFlag{
	Name:  "blob_file",
	Usage: "specify file containing a blob",
}

var TxFlag = cli.StringFlag{
	Name:  "tx",
	Usage: "specify tx hash",
}

var SpanFlag = cli.BoolFlag{
	Name:  "span",
	Usage: "specify span flag",
}

var TPSFlag = cli.IntFlag{
	Name:  "tps",
	Usage: "specify tps flag",
}

var DailyProposeTimesFlag = cli.IntFlag{
	Name:     "daily_propose",
	Usage:    "specify daily propose times flag",
	Required: true,
}

var BlobBaseFeeFlag = cli.IntFlag{
	Name:     "blob_base_fee",
	Usage:    "specify blob base fee flag",
	Required: true,
}

var ESInboxFlag = cli.BoolFlag{
	Name:  "es_inbox",
	Usage: "specify es inbox flag",
}
