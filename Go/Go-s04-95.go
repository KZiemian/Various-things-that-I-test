type Worker struct { /* ... */ }

type Work   struct { /* ... */ }

func New() *Worker {
	return &Worker{}
}

func (w *Worker) Fetch(ctx context.Context) (*Work, error) {
	_ = ctx
}

func (w *Worker) Process(ctx context.Context, work *Work) error {
	_ = ctx
}
