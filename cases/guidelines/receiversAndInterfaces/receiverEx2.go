package receiversAndInterfaces

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

func testReceivers2() {
	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr

	// The following doesn't compile, since s2Val is a value, and there is no value receiver for f.
	//i = s2Val

	print(i, s2Val) //for var usage
}
