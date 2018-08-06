package helper

import (
	"fmt"
	"os"
	"reflect"

	"github.com/zyxar/image2ascii/ascii"
)

// ShowLogo : show logo in console
func ShowLogo(img string, width int, heigth int) {
	opt := ascii.Options{
		Width:  width,
		Height: heigth,
		Color:  false,
		Invert: false,
		Flipx:  false,
		Flipy:  false}

	r, err := os.Open(img)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer r.Close()

	a, err := ascii.Decode(r, opt)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if _, err = a.WriteTo(os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

// inArray : fidding element existing and indexs
func inArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
