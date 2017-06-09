package main

// import "fmt"
import "encoding/xml"
import "os"
import "io/ioutil"
import "errors"

type Softwaredb struct {
	XMLName xml.Name   `xml:"softwaredb"`
	Soft    []Software `xml:"software"`
}

type Software struct {
	XMLName xml.Name `xml:"software"`
	Title   string   `xml:"title"`
	RomDump []Dump   `xml:"dump"`
}

type Dump struct {
	XMLName     xml.Name `xml:"dump"`
	Hash        string   `xml:"rom>hash"`
	HashMegaROM string   `xml:"megarom>hash"`
	TypeMegaROM string   `xml:"megarom>type"`
}

func searchInRomDatabase(romToSearch string) (string, error) {
	f, err := os.Open("softwaredb.xml")
	if err != nil {
		return "", err
	}
	defer f.Close()
	bytes, _ := ioutil.ReadAll(f)
	var swdb Softwaredb
	xml.Unmarshal(bytes, &swdb)

	for i := 0; i < len(swdb.Soft); i++ {
		soft := swdb.Soft[i]
		for j := 0; j < len(soft.RomDump); j++ {
			rd := soft.RomDump[j]
			if romToSearch == rd.HashMegaROM {
				return rd.TypeMegaROM, nil
			} else if romToSearch == rd.Hash {
				return "NORMAL", nil
			}
		}
	}

	return "", errors.New("Not found")
}
