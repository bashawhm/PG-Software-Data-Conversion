package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func splitToParts(l Locus) string {
	var final string
	for i := 0; i < 8; i++ {
		if i < len(l.Allele) {
			final += strconv.FormatFloat(l.Allele[i], 'f', -1, 64) + ","
		} else {
			final += ","
		}
	}
	for i := 0; i < 7; i++ {
		if i < len(l.Height) {
			final += strconv.FormatFloat(l.Height[i], 'f', -1, 64) + ","
		} else {
			final += ","
		}
	}
	return final
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./", os.Args[0], " <input_csv_list>...")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var sample Sample
	json.Unmarshal(data, &sample)

	fmt.Println("Sample Name,Marker,Allele 1,Allele 2,Allele 3,Allele 4,Allele 5,Allele 6,Allele 7,Allele 8,Height 1,Height 2,Height 3,Height 4,Height 5,Height 6,Height 7,Height 8")
	for i := 0; i < len(sample.Data); i++ {
		for j := 0; j < len(sample.Data[i].Replicates); j++ {
			fmt.Println(sample.Data[i].Name + ",D8S1179," + splitToParts(sample.Data[i].Replicates[j].D8S1179))
			fmt.Println(sample.Data[i].Name + ",D21S11," + splitToParts(sample.Data[i].Replicates[j].D21S11))
			fmt.Println(sample.Data[i].Name + ",D7S820," + splitToParts(sample.Data[i].Replicates[j].D7S820))
			fmt.Println(sample.Data[i].Name + ",CSF1PO," + splitToParts(sample.Data[i].Replicates[j].CSF1PO))
			fmt.Println(sample.Data[i].Name + ",D3S1358," + splitToParts(sample.Data[i].Replicates[j].D3S1358))
			fmt.Println(sample.Data[i].Name + ",TH01," + splitToParts(sample.Data[i].Replicates[j].TH01))
			fmt.Println(sample.Data[i].Name + ",D13S317," + splitToParts(sample.Data[i].Replicates[j].D13S317))
			fmt.Println(sample.Data[i].Name + ",D16S539," + splitToParts(sample.Data[i].Replicates[j].D16S539))
			fmt.Println(sample.Data[i].Name + ",D2S1338," + splitToParts(sample.Data[i].Replicates[j].D2S1338))
			fmt.Println(sample.Data[i].Name + ",D19S433," + splitToParts(sample.Data[i].Replicates[j].D19S433))
			fmt.Println(sample.Data[i].Name + ",vWA," + splitToParts(sample.Data[i].Replicates[j].VWA))
			fmt.Println(sample.Data[i].Name + ",TPOX," + splitToParts(sample.Data[i].Replicates[j].TPOX))
			fmt.Println(sample.Data[i].Name + ",D18S51," + splitToParts(sample.Data[i].Replicates[j].D18S51))
			fmt.Println(sample.Data[i].Name + ",D5S818," + splitToParts(sample.Data[i].Replicates[j].D5S818))
			fmt.Println(sample.Data[i].Name + ",FGA," + splitToParts(sample.Data[i].Replicates[j].FGA))
		}
	}
}
