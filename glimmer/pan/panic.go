package pan

import "fmt"

type PanErr struct {
	Err error
	Msg string
}

func (pe PanErr) Error() string {
	if pe.Msg != "" {
		return pe.Msg
	}
	return pe.Err.Error()
}

func Panic(err error) {
	if err != nil {
		panic(PanErr{Err: err})
	}
}

func Panicf(err error, format string, args ...interface{}) {
	if err != nil {
		panic(PanErr{Err: err, Msg: fmt.Sprintf(format, args...)})
	}
}

func Recover(errp *error) {
	i := recover()
	if i == nil {
		return
	}
	pe, ok := i.(PanErr)
	if !ok {
		panic(i)
	}
	*errp = pe.Err
}
