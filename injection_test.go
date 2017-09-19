package gautumn_test

import (
	"reflect"

	"github.com/dugancathal/gautumn"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type foo struct {
	Name string
}

func (f *foo) GetStuff() {
}

type iFoo interface {
	GetStuff()
}

func constructFoo() *foo {
	return &foo{}
}

var _ = Describe("Dependency Injection", func() {
	BeforeEach(func() {
		gautumn.Clean()
	})
	It("can retrieve an object by name", func() {
		f := &foo{}
		gautumn.RegisterByName("item", f)

		Expect(gautumn.GetDep("item")).To(Equal(reflect.ValueOf(f)))
	})

	It("panics if two deps are added with the same name", func() {
		f := &foo{"Foo 1"}
		gautumn.RegisterByName("item", f)

		Expect(func() {gautumn.RegisterByName("item", &foo{"Foo 2"})}).To(Panic())
	})

	It("can add items by type", func() {
		f := &foo{}
		gautumn.RegisterByType(f)

		Expect(gautumn.GetDep("*gautumn_test.foo")).To(Equal(reflect.ValueOf(f)))
	})

	It("can add items by interface", func() {
		f := &foo{}
		gautumn.RegisterByInterface((*foo)(nil), f)

		Expect(gautumn.GetDep("*gautumn_test.foo")).To(Equal(reflect.ValueOf(f)))
	})

	It("can add items by constructor", func() {
		gautumn.RegisterByConstructor(constructFoo)

		Expect(reflect.TypeOf(gautumn.GetDep("*gautumn_test.foo").Interface()).String()).To(Equal("*gautumn_test.foo"))
	})
})
