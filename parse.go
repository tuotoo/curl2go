package curl2go

import (
	"flag"
	"github.com/mattn/go-shellwords"
	"golang.org/x/xerrors"
	"log"
	"math/rand"
	"strings"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	rand.Seed(time.Now().UnixNano())
}

type arrayFlags []string

// Value ...
func (i *arrayFlags) String() string {
	return strings.Join(*i, " ")
}

// Set 方法是flag.Value接口, 设置flag Value的方法.
// 通过多个flag指定的值， 所以我们追加到最终的数组上.
func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func ParseCommand(command string) (string, error) {
	args, err := shellwords.Parse(command)
	if err != nil {
		return "", xerrors.Errorf("%w", err)
	}
	log.Printf("%#v\n", args)

	var headers = new(arrayFlags)
	var dataBinary = new(string)
	var compressed = new(bool)
	flagSet := flag.NewFlagSet(args[0], 1)
	flagSet.Var(headers, "H", "<header/@file> Pass custom header(s) to server")
	flagSet.Var(headers, "header", "<header/@file> Pass custom header(s) to server")
	flagSet.StringVar(dataBinary, "d", "", "HTTP POST data")
	flagSet.StringVar(dataBinary, "data", "", "HTTP POST data")
	flagSet.StringVar(dataBinary, "data-ascii", "", "HTTP POST ASCII data")
	flagSet.StringVar(dataBinary, "data-raw", "", "HTTP POST data, '@' allowed")
	flagSet.StringVar(dataBinary, "data-urlencode", "", "HTTP POST data url encoded")
	flagSet.StringVar(dataBinary, "delegation", "", "GSS-API delegation permission")

	flagSet.BoolVar(compressed, "compressed", false, "compressed")
	flagSet.Usage()
	err = flagSet.Parse(args[1:])
	if err != nil {
		return "", xerrors.Errorf("%w", err)
	}
	log.Printf("%#v\n", headers)
	log.Printf("%#v\n", flagSet)
	log.Printf("%#v\n", flagSet.Lookup("H"))
	log.Printf("%#v\n", flagSet.Lookup("data-binary"))
	return "", nil
}
