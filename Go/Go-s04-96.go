type Worker struct {
	ctx context.Context
}

func New(ctx context.Context) *Worker {
	return &Worker{ctx: ctx}
}

func (w *Worker) Fetch() (*Work, error) {
	_ = w.ctx
}

func (w *Worker) Process(work *Work) error {
	_ = w.ctx
}
