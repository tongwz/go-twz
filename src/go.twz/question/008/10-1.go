package main

type StudentA struct{}
type StudentB struct{}
type StudentC struct{}

func (stu *StudentA) AAA(think string) (talk string) {
	return think
}
func (stu *StudentA) BBB(think string) (talk string) {
	return think
}

func (stu StudentB) AAA(think string) (talk string) {
	return think
}
func (stu StudentB) BBB(think string) (talk string) {
	return think
}

func (stu StudentC) AAA(think string) (talk string) {
	return think
}
func (stu *StudentC) BBB(think string) (talk string) {
	return think
}

func main() {

}
