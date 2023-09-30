package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	httpietocurl "github.com/maruware/http-to-curl"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
			if cCtx.Args().Len() == 0 {
				return fmt.Errorf("usage: httpie-to-curl [...httpie args]")
			}

			r := httpietocurl.ParseHttpie(cCtx.Args().Slice())

			if args, err := httpietocurl.MakeCurlArgs(r); err != nil {
				return fmt.Errorf("invalid httpie args")
			} else {
				fmt.Printf("curl %s\n", strings.Join(args, " "))
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
