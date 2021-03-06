package gli

import "fmt"

type Context struct{}

func New() (*Context, error) {
	if driver == nil {
		return nil, fmt.Errorf("No driver registered")
	}
	err := driver.Init()
	if err != nil {
		return nil, err
	}
	ctx := &Context{}
	return ctx, nil
}

func (ctx *Context) Driver() Driver {
	return driver
}
