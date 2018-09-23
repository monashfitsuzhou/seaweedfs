package sink

import (
	"github.com/chrislusf/seaweedfs/weed/pb/filer_pb"
	"github.com/chrislusf/seaweedfs/weed/replication/source"
)

type ReplicationSink interface {
	DeleteEntry(key string, entry *filer_pb.Entry, deleteIncludeChunks bool) error
	CreateEntry(key string, entry *filer_pb.Entry) error
	UpdateEntry(key string, oldEntry, newEntry *filer_pb.Entry, deleteIncludeChunks bool) error
	GetSinkToDirectory() string
	SetSourceFiler(s *source.FilerSource)
}
