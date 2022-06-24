package main

import (
	"errors"
	"fmt"
	"io"

	xerrors "github.com/pkg/errors"
)

func doSomething() error {
	return errors.New("111")
}

func main() {
	err := errors.New("bad")
	fmt.Printf("%+v \n\n", err)
	err = xerrors.Wrap(err, "111")
	fmt.Printf("%+v \n\n", err)
	fmt.Printf("%+v \n\n", errors.Unwrap(err))
	fmt.Printf("%+v \n\n", xerrors.Cause(err))
	err = fmt.Errorf("this is error：%s", "bad news")
	fmt.Printf("%+v \n\n", err)

	err = io.EOF
	err = xerrors.Wrap(err, "111")

	if errors.Is(xerrors.Cause(err), io.EOF) {
		fmt.Print("true \n")
	}

	err = doSomething()

	if err != nil {
		err = xerrors.WithMessage(err, "dosomething")
		fmt.Printf("%+v", err)
		return
	}
}
