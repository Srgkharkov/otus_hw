package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func wrap(in In, done In) Bi {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case tmp, isOpened := <-in:
				if !isOpened {
					return
				}
				out <- tmp
			}
		}
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	cur := in
	for _, s := range stages {
		wrap := wrap(cur, done)
		cur = s(wrap)
	}
	return cur
}
