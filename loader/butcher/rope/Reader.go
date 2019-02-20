package rope

import "io/ioutil"

func ReadClass(classFullName string)([]byte, error){
	return ioutil.ReadFile(classFullName)
}
