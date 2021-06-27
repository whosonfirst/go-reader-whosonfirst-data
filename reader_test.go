package reader

import (
	"context"
	"github.com/whosonfirst/go-reader"
	"io"
	"io/ioutil"
	"testing"
)

const PATH_SFO string = "102/527/513/102527513.geojson"

func TestWhosOnFirstDataReader(t *testing.T) {

	ctx := context.Background()

	r, err := reader.NewReader(ctx, "whosonfirst-data://")

	if err != nil {
		t.Fatal(err)
	}

	fh, err := r.Read(ctx, PATH_SFO)

	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	_, err = io.Copy(ioutil.Discard, fh)

	if err != nil {
		t.Fatal(err)
	}

}
