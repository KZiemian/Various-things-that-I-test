func Compile(expr string) (*Regexp, error)

func MustCompile(str string) *Regexp

func MustCompilePOSIX(str string) *Regexp

func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)

func Compile(expr string) (*Regexp, error)

func MustCompile(str string) *Regexp

func MustCompilePOSIX(str string) *Regexp

func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)

func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int

func (re *Regexp) FindStringSubmatchIndex(s string) []int

func (re *Regexp) LiteralPrefix() (prefix string, complete bool)

func (re *Regexp) MatchReader(r io.RuneReader) bool

func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)

func Compile(expr string) (*Regexp, error)

func CompilePOSIX(expr string) (*Regexp, error)

func MustCompile(str string) *Regexp

func MustCompilePOSIX(str string) *Regexp

func (re *Regexp) FindAllStringIndex(s string, n int) [][]int

func (re *Regexp) FindStringSubmatchIndex(s string) []int

func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)

func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int

func (re *Regexp) LiteralPrefix() (prefix string, complete bool)

func (re *Regexp) MatchReader(r io.RuneReader) bool

func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte

func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
