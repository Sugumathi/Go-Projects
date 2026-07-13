package worklist

type Entry struct {
	path string
}

type Worklist struct {
	jobs chan Entry
}

func New(buff int) Worklist {
	w := Worklist{
		make(chan Entry, buff),
	}
	return w
}

func (w *Worklist) Add(work Entry) {
	w.jobs <- work
}

func (w *Worklist) Next() Entry {
	jobs := <-w.jobs
	return jobs
}

func (w *Worklist) Finalize(numJobs int) {
	for i := 0; i < numJobs; i++ {
		w.Add(Entry{""})
	}
}
