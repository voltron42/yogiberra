package yogi

import "errors"

func Catch(fn func()) error {
	var err error
	catch(fn, &err)
	return err
}

func catch(fn func(), err *error) {
	defer func() {
		if r := recover(); r != nil {
			var er error
			switch r.(type) {
			case string:
				str, _ := r.(string)
				er = errors.New(str)
			case error:
				value, _ := r.(error)
				er = value
			default:
				er = errors.New("Unknown Error")
			}
			*err = er
		}
	}()
	fn()
}
