package ogenapi

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
	return convertAssignRows(&oInt.Value, src)
}

