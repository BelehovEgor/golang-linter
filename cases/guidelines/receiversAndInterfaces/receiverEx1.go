package receiversAndInterfaces

type S struct {
	data string
}

func (s S) Read() string {
	return s.data
}

func (s *S) Write(str string) {
	s.data = str
}

func testReceivers1() {
	// We cannot get pointers to values stored in maps, because they are not
	// addressable values.
	var sVals = map[int]S{1: {"A"}}

	// We can call Read on values stored in the map because Read
	// has a value receiver, which does not require the value to
	// be addressable.
	sVals[1].Read()

	// We cannot call Write on values stored in the map because Write
	// has a pointer receiver, and it's not possible to get a pointer
	// to a value stored in a map.
	//sVals[1].Write("test")            it doesn't compile, comment is ok

	sPtrs := map[int]*S{1: {"A"}}

	// You can call both Read and Write if the map stores pointers,
	// because pointers are intrinsically addressable.
	sPtrs[1].Read()
	sPtrs[1].Write("test")
}
