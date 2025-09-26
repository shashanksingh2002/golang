package main
import "fmt"


type name struct {
    nameSlice [] string   
}

func Slice() {
    listOfNames:= name{
        nameSlice: []string {"shashank", "mohit"},
    }
    
    fmt.Println(listOfNames)
    listOfNames.updateName("mohit")
    fmt.Println(listOfNames)
}

func (n name) updateName(newName string) {
    n.nameSlice[0] = newName
}