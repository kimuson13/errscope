package a

import "log"

func f() {
	// The pattern can be written in regular expression.
	if err := e(); err != nil {
		log.Print(err)
	}

	hoge := h()
	if hoge {
		log.Print("hoge")
		err := e() //want "this error type can be narrowed in scope"
		if err != nil {
			log.Print(err)
		}
	}

	err, _ := e(), true //want "this error type can be narrowed in scope"
	if err != nil {
		log.Print(err)
	}

	v, err := t()
	if err != nil {
		log.Print(v)
	}

	err = e() //want "this error type can be narrowed in scope"
	if err != nil {
		log.Print(err)
	}

	err, piyo := e(), h() //want "this error type can be narrowed in scope"
	if err != nil && piyo {
		log.Print(err)
	}

	err, one, two := e(), h(), h() //want "this error type can be narrowed in scope"
	if err != nil {
		log.Print(one, two)
	}
}

func e() error {
	return nil
}

func h() bool {
	return true
}

func t() (int, error) {
	return 0, nil
}
