package offers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Index struct {
	FormatVersion   string           `json:"formatVersion"`
	Disclaimer      string           `json:"disclaimer"`
	PublicationDate string           `json:"publicationDate"`
	Offers          map[string]Offer `json:"offers"`
}

type Offer struct {
	OfferCode             string `json:"offerCode"`
	VersionIndexUrl       string `json:"versionIndexUrl"`
	CurrentVersionUrl     string `json:"currentVersionUrl"`
	CurrentRegionIndexUrl string `json:"currentRegionIndexUrl"`
}

type OffersApiClient struct {
	Region     string
	BaseUrl    string
	Version    string
	HttpClient http.Client
}

func NewOffersApiClient(region string) *OffersApiClient {
	c := new(OffersApiClient)
	c.Region = region
	c.Version = "v1.0"
	c.BaseUrl = fmt.Sprintf("https://pricing.%s.amazonaws.com", c.Region)
	c.HttpClient = http.Client{
		Timeout: time.Second * 2,
	}
	return c
}

func (c *OffersApiClient) newRequest(method string, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "aws_price_calc")
	return req, nil
}

func (c *OffersApiClient) doRequest(method string, url string) ([]byte, error) {
	req, err := c.newRequest(method, url)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func (c *OffersApiClient) GetIndex() (*Index, error) {
	url := fmt.Sprintf("%s/offers/%s/aws/index.json", c.BaseUrl, c.Version)

	body, err := c.doRequest(http.MethodGet, url)
	if err != nil {
		return nil, err
	}

	index := new(Index)
	err = json.Unmarshal(body, index)
	if err != nil {
		return nil, err
	}

	return index, nil
}
