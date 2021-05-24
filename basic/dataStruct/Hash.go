package dataStruct

//Hi=Hash값(S)
//Str[i]=Si
type keyValue struct {
	Key   string
	Value string
}
type Map struct {
	keyArray [3571][]keyValue
}

func (m *Map) Add(key, value string) {
	h := DoHash(key)
	m.keyArray[h] = append(m.keyArray[h], keyValue{key, value})
}
func DoHash(str string) int {
	h := 0
	A := 256
	B := 3571
	for i := 0; i < len(str); i++ {
		h = ((h * A) + int(str[i])) % B
	}
	return h
}
func CreateMap() *Map {
	return &Map{}
}
func (m *Map) Get(key string) string {
	h := DoHash(key)
	for i := 0; i < len(m.keyArray[h]); i++ {
		if m.keyArray[h][i].Key == key {
			return m.keyArray[h][i].Value
		}
	}
	return ""
}
