package classpath

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {

}

func (self *Classpath) readClass(className string) ([]byte, Entry, error) {
	//TODO implement me
	panic("implement me")
}

func (self *Classpath) String() string {
	//TODO implement me
	panic("implement me")
}
