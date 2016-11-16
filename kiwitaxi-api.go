package kiwitaxi

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type KiwitaxiApi struct {
	token string
	log   LoggerInterface
}

type LoggerInterface interface {
	Debug(...interface{})
}

type Fixer struct {
	Reader io.Reader
	gbuf   *bytes.Buffer
}

func NewFixerReader(r io.Reader) *Fixer {
	return &Fixer{
		Reader: r,
		gbuf:   bytes.NewBuffer([]byte{}),
	}
}

func (f *Fixer) Read(p []byte) (n int, err error) {
	err = nil
	if f.gbuf.Len() == 0 {
		fmt.Println("START!")

		r := bufio.NewReader(f.Reader)
		data, _, err := r.ReadLine()
		if err == io.EOF {
			fmt.Println("!!!!!!!! EOF!", err, data)
			n = 0
			return n, err
		}
		if err != nil || len(data) == 0 {
			return 0, err
		}

		f.gbuf.Write(
			[]byte(strings.Replace(
				strings.Replace(string(data), "\"", "\\\"", -1),
				"\\N",
				"0",
				-1,
			)),
		)
		f.gbuf.Write([]byte{'\n'})
	}

	n, _ = f.gbuf.Read(p)
	return
}

// NewKiwitaxiApi creates a new instance KiwitaxiApi.
func NewKiwitaxiApi(token string) *KiwitaxiApi {
	gocsv.SetCSVReader(func(in io.Reader) *csv.Reader {
		csvin := csv.NewReader(in)
		csvin.Comma = '\t'
		//csvin.LazyQuotes = true
		return csvin
	})

	return &KiwitaxiApi{
		token: token,
	}
}

type Coordinates struct {
	Lon float64 `json:"lon" bson:"lon"`
	Lan float64 `json:"lat" bson:"lat"`
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

	tmp := NewFixerReader(res.Body)
	d, _ := ioutil.ReadAll(tmp)
	fmt.Print(string(d))
	return gocsv.Unmarshal(tmp, v)
}

func (a *KiwitaxiApi) getCsv(path string, v interface{}) error {
	return a.getCsvWithParams(path, map[string]string{}, v)
}
