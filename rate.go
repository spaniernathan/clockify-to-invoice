package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const baseURL = "https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/"

// CompactData is the xml struct
type CompactData struct {
	XMLName        xml.Name `xml:"CompactData"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Header         struct {
		Text string `xml:",chardata"`
		ID   string `xml:"ID"`
		Test string `xml:"Test"`
		Name struct {
			Text string `xml:",chardata"`
			Lang string `xml:"lang,attr"`
		} `xml:"Name"`
		Prepared string `xml:"Prepared"`
		Sender   struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Name struct {
				Text string `xml:",chardata"`
				Lang string `xml:"lang,attr"`
			} `xml:"Name"`
			Contact struct {
				Text       string `xml:",chardata"`
				Department struct {
					Text string `xml:",chardata"`
					Lang string `xml:"lang,attr"`
				} `xml:"Department"`
				URI string `xml:"URI"`
			} `xml:"Contact"`
		} `xml:"Sender"`
		DataSetAgency string `xml:"DataSetAgency"`
		DataSetID     string `xml:"DataSetID"`
		Extracted     string `xml:"Extracted"`
	} `xml:"Header"`
	DataSet struct {
		Text           string `xml:",chardata"`
		Xmlns          string `xml:"xmlns,attr"`
		SchemaLocation string `xml:"schemaLocation,attr"`
		Group          struct {
			Text          string `xml:",chardata"`
			CURRENCY      string `xml:"CURRENCY,attr"`
			CURRENCYDENOM string `xml:"CURRENCY_DENOM,attr"`
			EXRTYPE       string `xml:"EXR_TYPE,attr"`
			EXRSUFFIX     string `xml:"EXR_SUFFIX,attr"`
			DECIMALS      string `xml:"DECIMALS,attr"`
			UNIT          string `xml:"UNIT,attr"`
			UNITMULT      string `xml:"UNIT_MULT,attr"`
			TITLECOMPL    string `xml:"TITLE_COMPL,attr"`
		} `xml:"Group"`
		Series struct {
			Text          string `xml:",chardata"`
			FREQ          string `xml:"FREQ,attr"`
			CURRENCY      string `xml:"CURRENCY,attr"`
			CURRENCYDENOM string `xml:"CURRENCY_DENOM,attr"`
			EXRTYPE       string `xml:"EXR_TYPE,attr"`
			EXRSUFFIX     string `xml:"EXR_SUFFIX,attr"`
			TIMEFORMAT    string `xml:"TIME_FORMAT,attr"`
			COLLECTION    string `xml:"COLLECTION,attr"`
			Obs           []struct {
				Text       string `xml:",chardata"`
				TIMEPERIOD string `xml:"TIME_PERIOD,attr"`
				OBSVALUE   string `xml:"OBS_VALUE,attr"`
				OBSSTATUS  string `xml:"OBS_STATUS,attr"`
				OBSCONF    string `xml:"OBS_CONF,attr"`
			} `xml:"Obs"`
		} `xml:"Series"`
	} `xml:"DataSet"`
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}
	return data, nil
}

// GetRate fetch EUR to currency exchangerate
func GetRate(currency string) (string, error) {
	URL := fmt.Sprintf("%s%s.xml", baseURL, strings.ToLower(currency))
	xmlBytes, err := getXML(URL)
	if err != nil {
		fmt.Printf("Failed to get XML: %v", err)
	}
	var data CompactData
	err = xml.Unmarshal(xmlBytes, &data)
	if err != nil {
		return "", fmt.Errorf("unmarshal data: %v", err)
	}
	now := time.Now()
	elem := data.DataSet.Series.Obs[len(data.DataSet.Series.Obs)-1]
	t, err := time.Parse("2006-01-02", elem.TIMEPERIOD)
	if err != nil {
		return "", fmt.Errorf("Failed to parse date: %v", err)
	}
	if t.Day() >= now.Day()-1 && t.Month() == now.Month() && t.Year() == now.Year() {
		floatRate, err := strconv.ParseFloat(elem.OBSVALUE, 64)
		if err != nil {
			return "", fmt.Errorf("Can't parse float: %v", err)
		}
		return fmt.Sprintf("%f", 1/floatRate), nil
	}
	return "", fmt.Errorf("No rate found")
}
