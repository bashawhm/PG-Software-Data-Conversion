package main

type Locus struct {
	Allele []float64 `json:"allele"`
	Height []float64 `json:"height"`
}

type Replicate struct {
	D8S1179 Locus `json:"D8S1179"`
	D21S11  Locus `json:"D21S11"`
	D7S820  Locus `json:"D7S820"`
	CSF1PO  Locus `json:"CSF1PO"`
	D3S1358 Locus `json:"D3S1358"`
	TH01    Locus `json:"TH01"`
	D13S317 Locus `json:"D13S317"`
	D16S539 Locus `json:"D16S539"`
	D2S1338 Locus `json:"D2S1338"`
	D19S433 Locus `json:"D19S433"`
	VWA     Locus `json:"vWA"`
	TPOX    Locus `json:"TPOX"`
	D18S51  Locus `json:"D18S51"`
	D5S818  Locus `json:"D5S818"`
	FGA     Locus `json:"FGA"`
}

type Comparison struct {
	D8S1179 []float64 `json:"D8S1179"`
	D21S11  []float64 `json:"D21S11"`
	D7S820  []float64 `json:"D7S820"`
	CSF1PO  []float64 `json:"CSF1PO"`
	D3S1358 []float64 `json:"D3S1358"`
	TH01    []float64 `json:"TH01"`
	D13S317 []float64 `json:"D13S317"`
	D16S539 []float64 `json:"D16S539"`
	D2S1338 []float64 `json:"D2S1338"`
	D19S433 []float64 `json:"D19S433"`
	VWA     []float64 `json:"vWA"`
	TPOX    []float64 `json:"TPOX"`
	D18S51  []float64 `json:"D18S51"`
	D5S818  []float64 `json:"D5S818"`
	FGA     []float64 `json:"FGA"`
}

type Data struct {
	Name         string       `json:"name"`
	Replicates   []Replicate  `json:"replicates"`
	Comparisons  []Comparison `json:"comparison"`
	Contributors int          `json:"contributors"`
	Deducable    bool         `json:"deducable"`
	Quantity     int          `json:"quantity"`
}

type Sample struct {
	SampleType string `json:"sampleType"`
	Data       []Data `json:"data"`
}
