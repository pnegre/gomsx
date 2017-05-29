package main

type Ports struct {
    a int;
}


func (self *Ports) ReadPort(address uint16) byte {
    return 0;
}

func (self *Ports) WritePort(address uint16, b byte) {

}

func (self *Ports) ReadPortInternal(address uint16, contend bool) byte {
    return 0;
}

func (self *Ports) WritePortInternal(address uint16, b byte, contend bool) {
}

func (self *Ports) ContendPortPreio(address uint16) {

}

func (self *Ports) ContendPortPostio(address uint16) {

}
