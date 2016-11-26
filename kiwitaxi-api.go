package kiwitaxi

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
	"net/http"
	"net/url"
)

type KiwitaxiApi struct {
	token string
	log   LoggerInterface
}

type LoggerInterface interface {
	Debug(...interface{})
}

// NewKiwitaxiApi creates a new instance KiwitaxiApi.
func NewKiwitaxiApi(token string) *KiwitaxiApi {
	gocsv.SetCSVReader(func(in io.Reader) *csv.Reader {
		csvin := csv.NewReader(in)
		csvin.Comma = '\t'
		csvin.LazyQuotes = true
		return csvin
	})

	return &KiwitaxiApi{
		token: token,
	}
}

func (a *KiwitaxiApi) SetLogger(logger LoggerInterface) {
	a.log = logger
}

func (a *KiwitaxiApi) getCsvWithParams(path string, args map[string]string, v interface{}) error {
	apiUrl, err := url.Parse("https://kiwitaxi.com/" + path)
	if err != nil {
		return err
	}

	params := url.Values{}
	params.Add("security_token", a.token)
	for k, v := range args {
		if v == "" {
			continue
		}
		params.Add(k, v)
	}

	apiUrl.RawQuery = params.Encode()

	if a.log != nil {
		a.log.Debug("API Send: " + apiUrl.String())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl.String(), nil)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return gocsv.Unmarshal(NewFixerReader(res.Body), v)
}

func (a *KiwitaxiApi) getCsv(path string, v interface{}) error {
	return a.getCsvWithParams(path, map[string]string{}, v)
}
