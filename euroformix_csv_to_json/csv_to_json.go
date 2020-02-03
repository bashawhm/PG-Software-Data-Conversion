package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getAlleles(parts []string) []float64 {
	var alleles []float64
	for i := 2; i < 10; i++ {
		if parts[i] != "" {
			num, err := strconv.ParseFloat(parts[i], 64)
			if err != nil {
				panic(err)
			}
			alleles = append(alleles, num)
		}
	}
	return alleles
}

func getHeights(parts []string) []float64 {
	var heights []float64
	for i := 10; i < 18; i++ {
		if parts[i] != "" {
			num, err := strconv.ParseFloat(parts[i], 64)
			if err != nil {
				panic(err)
			}
			heights = append(heights, num)
		}
	}
	return heights
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

	var sample Sample
	sample.SampleType = "evidence"

	data := Data{}
	for i := 1; i < len(os.Args); i++ {
		fin := bufio.NewScanner(bufio.NewReader(file))
		fin.Split(bufio.ScanLines)
		fin.Scan()
		replicate := Replicate{}
		for i := 0; fin.Scan(); i++ {
			parts := strings.Split(fin.Text(), ",")
			data.Name = parts[0]
			//Appologies for this, will find better way, but in my sleep deprived state the only other way I can think to do this nicer is reflection
			switch i {
			case 0:
				replicate.D8S1179.Allele = getAlleles(parts)
				replicate.D8S1179.Height = getHeights(parts)
			case 1:
				replicate.D21S11.Allele = getAlleles(parts)
				replicate.D21S11.Height = getHeights(parts)
			case 2:
				replicate.D7S820.Allele = getAlleles(parts)
				replicate.D7S820.Height = getHeights(parts)
			case 3:
				replicate.CSF1PO.Allele = getAlleles(parts)
				replicate.CSF1PO.Height = getHeights(parts)
			case 4:
				replicate.D3S1358.Allele = getAlleles(parts)
				replicate.D3S1358.Height = getHeights(parts)
			case 5:
				replicate.TH01.Allele = getAlleles(parts)
				replicate.TH01.Height = getHeights(parts)
			case 6:
				replicate.D13S317.Allele = getAlleles(parts)
				replicate.D13S317.Height = getHeights(parts)
			case 7:
				replicate.D16S539.Allele = getAlleles(parts)
				replicate.D16S539.Height = getHeights(parts)
			case 8:
				replicate.D2S1338.Allele = getAlleles(parts)
				replicate.D2S1338.Height = getHeights(parts)
			case 9:
				replicate.D19S433.Allele = getAlleles(parts)
				replicate.D19S433.Height = getHeights(parts)
			case 10:
				replicate.VWA.Allele = getAlleles(parts)
				replicate.VWA.Height = getHeights(parts)
			case 11:
				replicate.TPOX.Allele = getAlleles(parts)
				replicate.TPOX.Height = getHeights(parts)
			case 12:
				replicate.D18S51.Allele = getAlleles(parts)
				replicate.D18S51.Height = getHeights(parts)
			case 13:
				replicate.D5S818.Allele = getAlleles(parts)
				replicate.D5S818.Height = getHeights(parts)
			case 14:
				replicate.FGA.Allele = getAlleles(parts)
				replicate.FGA.Height = getHeights(parts)
			}
		}
		data.Replicates = append(data.Replicates, replicate)
	}

	sample.Data = append(sample.Data, data)

	j, err := json.Marshal(sample)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}
