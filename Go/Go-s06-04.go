func (page *Page) saveSourceAs(path string) {
	b := new(bytes.Buffer)
	b.Write(page.Source.Content)
	page.saveSource(b.Bytes(), path)
}

func (page *Page) saveSource(by []byte, inpath string) {
	WriteToDisk(inpath, bytes.NewReader(by))
}

func (page *Page) saveSourceAs(path string) {
	b := new(byts.Buffer)
	b.Write(page.Source.Content)
	page.saveSource(b, path)
}

func (page *Page) saveSource(b io.Reader, inpath string) {
	WriteToDisk(inpath, b)
}

type MyReader interface {
	Read(p []byte) (n int, err error)
}

type MyWriter interface {
	Write(p []byte) (n int, err error)
}

func (c *Command) SetOutput(o io.Writer) {
	c.output = o
}

func (v *Viper) ReadBufConfig(buf *bytes.Buffer) error {
	v.config = make(map[string]interface{})
	v.marshalReader(buf, v.config)

	return nil
}

func (v *Viper) ReadConfig(in io.Reader) error {
	v.config = make(map[string]interface{})
	v.marshalReader(in, v.config)

	return nil
}
