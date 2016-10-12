package samfs

import (
	"sync"
	//"syscall"
	"time"

	"github.com/golang/glog"
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type SamFsOptions struct{}

type SamFs struct {
	pathfs.FileSystem
	Mount     string
	cacheLock sync.RWMutex
	fileCache map[string]*SamFsFileData
	options   *SamFsOptions
}

func NewSamFs(opts *SamFsOptions) *SamFs {
	return &SamFs{
		options:   opts,
		fileCache: make(map[string]*SamFsFileData),
	}
}

func (c *SamFs) SetDebug(debug bool) {
}

// Attributes.  This function is the main entry point, through
// which FUSE discovers which files and directories exist.
//
// If the filesystem wants to implement hard-links, it should
// return consistent non-zero FileInfo.Ino data.  Using
// hardlinks incurs a performance hit.
func (c *SamFs) GetAttr(name string, context *fuse.Context) (*fuse.Attr,
	fuse.Status) {

	return nil, fuse.OK
}

func (c *SamFs) Truncate(path string, size uint64,
	context *fuse.Context) fuse.Status {

	return fuse.EINVAL
}

func (c *SamFs) Utimens(name string, atime *time.Time, mtime *time.Time,
	context *fuse.Context) fuse.Status {

	return fuse.OK
}

func (c *SamFs) Chown(name string, uid uint32, gid uint32,
	context *fuse.Context) fuse.Status {

	return fuse.OK
}

func (c *SamFs) Chmod(name string, mode uint32,
	context *fuse.Context) fuse.Status {

	return fuse.OK
}

func (c *SamFs) Access(name string, mode uint32,
	context *fuse.Context) fuse.Status {

	return fuse.OK
}

func (c *SamFs) Link(orig string, newName string,
	context *fuse.Context) fuse.Status {

	return fuse.OK
}

func (c *SamFs) Rmdir(path string, context *fuse.Context) fuse.Status {
	return fuse.OK
}

func (c *SamFs) Mkdir(path string, mode uint32,
	context *fuse.Context) fuse.Status {

	return fuse.OK
}

func (c *SamFs) Rename(oldName string, newName string,
	context *fuse.Context) fuse.Status {

	return fuse.OK
}

func (c *SamFs) Unlink(name string, context *fuse.Context) fuse.Status {
	return fuse.OK
}

func (c *SamFs) GetXAttr(name string, attribute string,
	context *fuse.Context) ([]byte, fuse.Status) {

	return []byte{}, fuse.OK
}

func (c *SamFs) RemoveXAttr(name string, attr string,
	context *fuse.Context) fuse.Status {

	return fuse.OK
}

func (c *SamFs) SetXAttr(name string, attr string, data []byte, flags int,
	context *fuse.Context) fuse.Status {

	return fuse.OK
}

func (c *SamFs) ListXAttr(name string, context *fuse.Context) ([]string,
	fuse.Status) {

	return []string{}, fuse.OK
}

func (c *SamFs) OnMount(nodefs *pathfs.PathNodeFs) {
	glog.Info("mount ok")
}

func (c *SamFs) OnUnmount() {
}

func (c *SamFs) Open(name string, flags uint32,
	context *fuse.Context) (nodefs.File, fuse.Status) {

	return nil, fuse.OK
}

func (c *SamFs) OpenDir(name string, context *fuse.Context) ([]fuse.DirEntry,
	fuse.Status) {

	return nil, fuse.OK
}

func (c *SamFs) Create(name string, flags uint32, mode uint32,
	context *fuse.Context) (nodefs.File, fuse.Status) {

	return nil, fuse.OK
}

func (c *SamFs) Symlink(pointedTo string, linkName string,
	context *fuse.Context) fuse.Status {

	return fuse.OK
}

func (c *SamFs) Readlink(name string, context *fuse.Context) (string,
	fuse.Status) {

	return "", fuse.OK
}

func (c *SamFs) StatFs(name string) *fuse.StatfsOut {
	return &fuse.StatfsOut{}
}