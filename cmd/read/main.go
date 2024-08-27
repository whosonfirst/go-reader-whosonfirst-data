package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/whosonfirst/go-reader-whosonfirst-data"

	"github.com/mitchellh/go-wordwrap"
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-whosonfirst-uri"
)

func main() {

	reader_uri := flag.String("reader-uri", "whosonfirst-data://", "A valid go-reader-whosonfirst-data URI.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Read one or more Who's On First records from the whosonfirst-data GitHub organization.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [options] [path1 path2 ... pathN]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nNotes:\n\n")
		fmt.Fprintf(os.Stderr, wordwrap.WrapString("pathN may be any valid Who's On First ID or URI that can be parsed by the go-whosonfirst-uri package.\n\n", 80))
		fmt.Fprintf(os.Stderr, wordwrap.WrapString("The default behaviour for resolving a Who's On First ID to its corresponding GitHub repository is to use the 'data.whosonfirst.org/findingaid' lookup service. If you know the repository that all of the pathN arguments are part of you can skip the lookup service by appending the repo name as a query parameter to the '-reader-uri' argument. For example: -reader-uri 'whosonfirst-data://?repo=whosonfirst-data-admin-ca'.\n\n", 80))
	}

	flag.Parse()

	ctx := context.Background()

	r, err := reader.NewReader(ctx, *reader_uri)

	if err != nil {
		log.Fatalf("Unable to create new reader, %v", err)
	}

	uris := flag.Args()

	for _, raw := range uris {

		id, _, err := uri.ParseURI(raw)

		if err != nil {
			log.Fatalf("Unable to parse URI '%s', %v", raw, err)
		}

		rel_path, err := uri.Id2RelPath(id)

		if err != nil {
			log.Fatalf("Unable to derive relative path from ID '%d', %v", id, err)
		}

		fh, err := r.Read(ctx, rel_path)

		if err != nil {
			log.Fatalf("Unable to read path '%s', %v", rel_path, err)
		}

		defer fh.Close()

		_, err = io.Copy(os.Stdout, fh)

		if err != nil {
			log.Fatalf("Unable to read contents of '%s', %v", rel_path, err)
		}
	}
}
