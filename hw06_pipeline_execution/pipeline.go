package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, stage := range stages {
		out = func(in In) Out {
			inner := make(Bi)

			go func() {
				defer close(inner)

				for {
					select {
					case <-done:
						return
					case v, ok := <-in:
						if !ok {
							return
						}

						inner <- v
					}
				}
			}()

			return stage(inner)
		}(out)
	}

	return out
}
