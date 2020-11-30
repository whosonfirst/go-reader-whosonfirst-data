package reader

import (
	"context"
	"testing"
	"github.com/whosonfirst/go-reader"
	"io"
	"io/ioutil"
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
