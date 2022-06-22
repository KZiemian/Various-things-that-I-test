var ErrProcessDone = errors.New("os: process already finished")

func Chdir(dir string) error

func Chmod(name string, mode FileMode) error

func Chown(name string, uid, gid int) error

func Chtimes(name string, atime time.Time, mtime time.Time) error

func Clearenv()

func DirFS(dir string) fs.FS

func Environ() []string

errors.Is(err, fs.ErrExit)

errors.Is(err, fs.ErrNotExit)

errors.Is(err, fs.ErrPermission)

errors.Is(err, fs.ErrPErmission)

func Lchown(name string, uid, gid int) error

func Link(oldname, newname string) error

func Mkdir(name string, perm FileMode) error

func MkdirAll(path string, perm FileMode) error

func NewSyscallError(syscall string, err error) error

func Pipe() (r *File, w *File, err error)

func Readlink(name string) (string, error)

func Remove(name string) error

func RemoveAll(path string) error

func Rename(oldpath, newpath string) error

func SameFile(fi1, fi2 FileInfo) bool

func Setenv(key, value string) error

func Symlink(oldname, newname string) error

func Truncate(name string, size int64) error

type DirEntry = fs.DirEntry

func ReadDir(name string) ([]DirEntry, error)

func NewFile(fd uintptr, name string) *File

func Open(name string) (*File, error)

func (f *File) Chdir() error

func (f *File) Chmod(mode FileMode) error

func (f *File) Chown(uid, gid int) error

func (f *File) Close() error

func (f *File) Fd() uintptr

func (f *File) Name() string

func (f *File) Read(b []byte) (n int, err error)

func (f *File) ReadAt(b []byte, off int64) (n int, err error)

func (f *File) ReadDir(n int) ([]DirEntry, error)

func (f *File) ReadFrom(r io.Reader) (n int64, err error)

func (f *File) Readdir(n int) ([]FileInfo, error)

func (f *File) Readdirnames(n int) (names []string, err error)

func (f *File) Seek(offset int64, whence int) (ret int64, err error)

func (f *File) SetDeadline(t time.Time) error

error.Is(err, os.ErrDeadlineExceeded)

func (f *File) SetReadDeadline(t time.Time) error

func (f *File) SetWriteDeadline(t time.Time) error

func (f *File) Stat() (FileInfo, error)

func (f *File) Sync() error

func (f *File) SyscallConn() (syscall.RawConn, error)

func (f *File) Truncate(size int64) error
