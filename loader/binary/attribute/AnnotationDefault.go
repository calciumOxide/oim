package attribute

type AnnotationDefault struct {
	ElementValue ElementValue
}

func AllocAnnotationDefault(b []byte) (*AnnotationDefault, int) {
	e, s := AllocElementValue(b)
	v := AnnotationDefault {
		ElementValue: *e,
	}
	return &v, s
}

