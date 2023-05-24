package main

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show)
	s.Param = make(map[string]interface{})
	s.Param["RMB"] = 10000
}

// func main() {
// 	s := new(Show)
// 	fmt.Printf(" param: %#v, param:ptr: %p \n", s.Param, s.Param)
// 	s.Param = make(map[string]interface{}, 3)
// 	fmt.Printf("after param: %#v, param:ptr: %p \n", s.Param, s.Param)
// 	s.Param["RMB1"] = 10000
// 	s.Param["RMB2"] = 10000
// 	s.Param["RMB3"] = 10000
// 	s.Param["RMB4"] = 10000
// 	fmt.Printf("after2 param: %#v, param:ptr: %p \n", s.Param, s.Param)
// }
