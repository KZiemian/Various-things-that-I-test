func (f *File) Write(b []byte) (n int, err error)

func (f *File) WriteAt(b []byte, off int64) (n int, err error)

func (f *File) WriteString(s string) (n int, err error)

func Lstat(name string) (FileInfo, error)

func Stat(name string) (FileInfo, error)

buf := make([]byte, 100)
n, err := r.Read(buf)
buf = buf[:n]

if err == io.EOF {
	log.Fatal("read failed:", err)
}

func readfile(path string) error {
	err := openfile(path)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	// ...
}

func main() {
	err := readfile(".bashrc")
	if string.Contains(err.Error(), "not found") {
		// handle error
	}
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

type LinkError struct {
	Op  string
	Old string
	New string
	Err error
}

func (e *LinkError) Error() string

func (e *LinkError) Unwrap() error

type ProcAttr struct {
	Dir string
	Env []string
	Files []*File

	Sys *syscall.SysProcAttr
}

type Process struct {
	Pid int
	// contains filtered or unexported fields
}

func FindProcess(pid int) (*Process, error)

func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)

func (p *Process) Kill() error

func (p *Process) Release() error

func (p *Process) Signal(sig Signal) error

func (p *Process) Wait() (*ProcessState, error)

func (p *ProcessState) ExitCode() int

func (p *ProcessState) Exited() bool

func (p *ProcessState) Pid() int

func (p *ProcessState) String() string

func (p *ProcessState) Success() bool

func (p *ProcessState) Sys() interface{}

func (p *ProcessState) SysUsage() interface{}

func (p *ProcessState) SystemTime() time.Duration

func (p *ProcessState) UserTime() time.Duration

type Signal interface {
	String() string
	Singal() // to distinguis from other Stringers
}

type SyscallError struct {
	Syscall string
	Err     error
}
