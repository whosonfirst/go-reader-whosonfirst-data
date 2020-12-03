package reader

import (
	"context"
	"github.com/whosonfirst/go-reader"
	"io"
	"io/ioutil"
	"testing"
)

func TestWhosOnFirstDataReader(t *testing.T) {

	ctx := context.Background()

	r, err := reader.NewReader(ctx, "whosonfirst-data://")

	if err != nil {
		t.Fatal(err)
	}

	rel_path := "101/736/545/101736545.geojson"

	fh, err := r.Read(ctx, rel_path)

	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	_, err = io.Copy(ioutil.Discard, fh)

	if err != nil {
		t.Fatal(err)
	}

}

func TestSFOMuseumDataReader(t *testing.T) {

	ctx := context.Background()

	r, err := reader.NewReader(ctx, "whosonfirst-data://?organization=sfomuseum-data")

	if err != nil {
		t.Fatal(err)
	}

	rel_path := "172/956/253/7/1729562537.geojson"

	fh, err := r.Read(ctx, rel_path)

	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	_, err = io.Copy(ioutil.Discard, fh)

	if err != nil {
		t.Fatal(err)
	}

}
