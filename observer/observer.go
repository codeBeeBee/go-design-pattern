package observer

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

type IObserver interface {
	Update(msg string)
}

type Subject struct {
	observers []IObserver
}

func (subject *Subject) Register(observer IObserver) {
	subject.observers = append(subject.observers, observer)
}

func (subject *Subject) Remove(observer IObserver) {
	for i, ob := range subject.observers {
		if ob == observer {
			subject.observers = append(subject.observers[:i], subject.observers[i+1:]...)
		}
	}
}

func (subject *Subject) Notify(msg string) {
	for _, ob := range subject.observers {
		ob.Update(msg)
	}
}
