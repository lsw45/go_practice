package ipld

import (
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/mount"
	flatfs "github.com/ipfs/go-ds-flatfs"
	leveldb "github.com/ipfs/go-ds-leveldb"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	exchange "github.com/ipfs/go-ipfs-exchange-interface"
	ipld "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	"os"
	"path"
)

/*
IPLD Node 数据的存储方案
分了两个库存储，其中一个是 leveldb 用来存储关系，也就是 ipld.Node 中的 cid 和 links 的映射，另一个 filedb 库存储数据，cid 和 rawdata 的映射，
将关系和数据分开存储，这样做的目的是为了重用 rawdata，
关系可以任意组合，但是每个 rawdata 就一份；
举个例子假设 rawdata 的 cid = a ，那么 a 这个 cid 可以出现在任意多个 links 集合中，但是 a 对应的 rawdata 就只有一份；
*/

type Adag struct {
	db   datastore.Batching
	bs   blockstore.Blockstore
	bsrv blockservice.BlockService
	dsrv ipld.DAGService
	rem  exchange.Interface
}

func NewAdag(homedir string, rem exchange.Interface) *Adag {
	a := new(Adag)
	a.rem = rem
	a.db = mountdb(homedir)
	a.bs = blockstore.NewBlockstore(a.db)
	a.bsrv = blockservice.New(a.bs, a.rem)
	a.dsrv = merkledag.NewDAGService(a.bsrv)
	return a
}

func (a *Adag) DAGService() ipld.DAGService { return a.dsrv }

func mountdb(homedir string) datastore.Batching {
	// db 是一个集合，根据前缀来区分使用 fdb 还是 leveldb
	// key 的前缀对应的 db 实例是在 .ipfs/config 中进行配置的 Datastore 项
	fp := path.Join(homedir, "blocks")
	os.MkdirAll(fp, 0755)
	fdb, err := flatfs.CreateOrOpen(fp, flatfs.NextToLast(2), true)
	if err != nil {
		panic(err)
	}

	ldb, err := leveldb.NewDatastore(path.Join(homedir, "datastore"), nil)
	if err != nil {
		panic(err)
	}
	mnt := []mount.Mount{
		{
			Prefix:    datastore.NewKey("/blocks"),
			Datastore: fdb,
			//Datastore: measure.New("flatfs.datastore", fdb),
		},
		{
			Prefix:    datastore.NewKey("/"),
			Datastore: ldb,
			//Datastore: measure.New("leveldb.datastore", ldb),
		},
	}
	return mount.New(mnt)
}
