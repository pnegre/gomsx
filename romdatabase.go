package main

import "encoding/xml"

import "io/ioutil"
import "errors"

const XMLDATABASE = "softwaredb.xml"

func searchInRomDatabase(romToSearch string) (string, error) {
	type Softwaredb struct {
		XMLName xml.Name `xml:"softwaredb"`
		Soft    []struct {
			XMLName xml.Name `xml:"software"`
			Title   string   `xml:"title"`
			RomDump []struct {
				XMLName     xml.Name `xml:"dump"`
				Hash        string   `xml:"rom>hash"`
				HashMegaROM string   `xml:"megarom>hash"`
				TypeMegaROM string   `xml:"megarom>type"`
			} `xml:"dump"`
		} `xml:"software"`
	}

	bytes, err := ioutil.ReadFile(XMLDATABASE)
	if err != nil {
		return "", err
	}
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
