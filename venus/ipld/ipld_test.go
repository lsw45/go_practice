package ipld

import (
	"context"
	"fmt"
	"github.com/ipfs/go-ipfs-chunker"
	"github.com/ipfs/go-unixfs/importer/balanced"
	h "github.com/ipfs/go-unixfs/importer/helpers"
	"strings"
	"testing"
)

var homedir = "/Users/carey/workspace/helloipfs"

func TestBalanceLayout(t *testing.T) {
	ds := NewAdag(homedir, nil).DAGService()
	blocksize := 256 * 1024
	data := homedir + "233333"
	t.Log(len([]byte(data)) / blocksize)
	r := strings.NewReader(data)
	spl := chunk.NewSizeSplitter(r, int64(blocksize))

	dbp := h.DagBuilderParams{
		Dagserv:  ds,
		Maxlinks: h.DefaultLinksPerBlock,
	}

	dd, err := dbp.New(spl)
	if err != nil {
		panic(err)
	}

	nd, err := balanced.Layout(dd)
	if err != nil {
		panic(err)
	}
	ds.Add(context.Background(), nd)
	size, _ := nd.Size()
	t.Log(nd.String(), blocksize, size, len(nd.Links()))
	for i, l := range nd.Links() {
		n, err := ds.Get(context.Background(), l.Cid)
		fmt.Println(i, "-->", err, n, len(n.Links()))
	}
}
