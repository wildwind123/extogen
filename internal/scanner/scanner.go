package scanner

func (oS *OptString) Scan(src any) error {
	if src == nil {
		return nil
	}
	oS.Set = true
	return convertAssignRows(&oS.Value, src)
}

func (oInt *OptInt) Scan(src any) error {
	if src == nil {
		return nil
	}
	oInt.Set = true
	return convertAssignRows(&oInt.Value, src)
}

func (oInt *OptInt64) Scan(src any) error {
	if src == nil {
		return nil
	}
	oInt.Set = true
	return convertAssignRows(&oInt.Value, src)
}

func (oFloat64 *OptFloat64) Scan(src any) error {
	if src == nil {
		return nil
	}
	oFloat64.Set = true
	return convertAssignRows(&oFloat64.Value, src)
}
