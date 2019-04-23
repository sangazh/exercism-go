package erratum

func Use(o ResourceOpener, input string) (err error) {
	var resource Resource
	resource, err = o()

	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		resource, err = o()
	}
	defer resource.Close()

	if err = frob(resource, input); err != nil {
		return err
	}
	return
}

//rewrite frob to handle the panic
func frob(rsc Resource, input string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if frobErr, ok := r.(FrobError); ok {
				rsc.Defrob(frobErr.defrobTag)
				err = frobErr.inner
			} else {
				err = r.(error)
			}
		}
	}()

	rsc.Frob(input)
	return
}
