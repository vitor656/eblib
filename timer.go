package eblib

type Timer struct {
	counterMax int
	counter    int
	started    bool
	repeat     bool
	onTrigger  func()
}

func NewTimer(count int, repeat bool, onTrigger func()) *Timer {
	timer := &Timer{
		counterMax: count,
		counter:    0,
		repeat:     repeat,
		onTrigger:  onTrigger,
	}

	return timer
}

func (t *Timer) Start() {
	t.started = true

}

func (t *Timer) Update() error {
	if !t.started {
		return nil
	}

	t.counter += 1

	if t.counter >= t.counterMax {
		t.counter = 0
		if t.onTrigger != nil {
			t.onTrigger()
		}
		if !t.repeat {
			t.started = false
		}
	}

	return nil
}
