package tutorial

import "testing"  

func FuzzBrokenMethod(f *testing.F) {
  f.Fuzz(func(t *testing.T, str string) {
    BrokenMethod(str)
  })
}
