package main

import (
	"JvmCreatedByGolang/JvmGo/classpath"
	"fmt"
	"strings"
)

func main() {
	//fmt.Println("jvm starting...")
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("qiu 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%s class:%s args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	class, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("could not find or load main class %s\n", class)
		return
	}
	fmt.Printf("class data:%v\n", class)
}
