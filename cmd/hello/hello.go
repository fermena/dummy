package hello

// #include <stdio.h>
// #include <stdlib.h>
//
// static void myprint(char* s) {
//   printf("%s\n", s);
// }
import "C"
import "unsafe"

func Hello() error {
	cs := C.CString("Hello, this id Dummy app :)")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
	return nil
}
