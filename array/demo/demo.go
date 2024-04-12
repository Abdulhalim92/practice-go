package demo

const NewValue = "changedValue"

func UpdateArray1(array [2]string) {
	array[0] = NewValue
}

func UpdateArray2(array *[2]string) {
	array[0] = NewValue
}
