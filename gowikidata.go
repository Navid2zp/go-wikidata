package gowikidata

import (
	"errors"
	"fmt"
	"github.com/Navid2zp/easyreq"
)

var WikipediaQueryURL = "https://en.wikipedia.org/w/api.php?action=query&prop=pageprops&titles=%s&format=json"
var WikiDataAPIURL = "https://www.wikidata.org/w/api.php?action=%s&format=json"
var ImageResizerURL = "https://commons.wikimedia.org/w/thumb.php?width=%d&f=%s"

// Gets Wikipedia page item ID in wikidata
func GetPageItem(slug string) (string, error) {
	url := fmt.Sprintf(WikipediaQueryURL, slug)

	wikipediaQuery := WikiPediaQuery{}
	res, err := easyreq.Make("get", url, nil, "", "json", &wikipediaQuery, nil)
	if err != nil {
		return "", err
	}

	if res.StatusCode() != 200 {
		return "", errors.New("request failed with status code " + string(res.StatusCode()))
	}

	item := ""
	// Get the first item
	for _, value := range wikipediaQuery.Query.Pages {
		item = value.PageProps.WikiBaseItem
		break
	}
	return item, nil
}

// Use this function to initialize a request.
// It will return a pointer to WikiDataGetEntitiesRequest.
// You can make configurations on response.
func NewGetEntities(ids []string) (*WikiDataGetEntitiesRequest, error) {
	if len(ids) < 1 {
		return nil, errors.New("no ids provided")
	}

	req := WikiDataGetEntitiesRequest{
		URL: fmt.Sprintf(WikiDataAPIURL, "wbgetentities"),
	}
	req.setParam("ids", &ids)
	return &req, nil
}

// Param: sites
func (r *WikiDataGetEntitiesRequest) SetSites(sites []string) {
	r.setParam("sites", &sites)
}

// Param: titles
func (r *WikiDataGetEntitiesRequest) SetTitles(titles []string) {
	r.setParam("titles", &titles)
}

// Param: redirects
func (r *WikiDataGetEntitiesRequest) SetRedirects(redirect bool) {
	redirectString := "yes"
	if !redirect {
		redirectString = "no"
	}
	r.URL += "&redirects=" + redirectString
}

// Param: props
// Default: info|sitelinks|aliases|labels|descriptions|claims|datatype
func (r *WikiDataGetEntitiesRequest) SetProps(props []string) {
	r.setParam("props", &props)
}

// Param: languages
func (r *WikiDataGetEntitiesRequest) SetLanguages(languages []string) {
	r.setParam("languages", &languages)
}

// Param: languagefallback
func (r *WikiDataGetEntitiesRequest) SetLanguageFallback(fallback bool) {
	if fallback {
		r.URL += "&languagefallback="
	}
}

// Param: normalize
func (r *WikiDataGetEntitiesRequest) SetNormalize(normalize bool) {
	if normalize {
		r.URL += "&normalize="
	}
}

// Param: sitefilter
func (r *WikiDataGetEntitiesRequest) SetSiteFilter(sites []string) {
	r.setParam("sitefilter", &sites)
}

// Call this function after you finished configuring the request.
// It will send the request and unmarshales the response.
func (r *WikiDataGetEntitiesRequest) Get() (*map[string]Entity, error) {
	responseData := GetEntitiesResponse{}
	res, err := easyreq.Make("GET", r.URL, nil, "", "json", &responseData, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, errors.New("request failed with status code " + string(res.StatusCode()))
	}
	return &responseData.Entities, nil
}

// WikiData action: wbgetclaims
// WikiData API page: https://www.wikidata.org/w/api.php?action=help&modules=wbgetclaims
// Either entity or claim must be provided
func NewGetClaims(entity, claim string) (*WikiDataGetClaimsRequest, error) {
	if entity == "" && claim == "" {
		return nil, errors.New("either entity or claim must be provided")
	}

	req := WikiDataGetClaimsRequest{
		URL: fmt.Sprintf(WikiDataAPIURL, "wbgetclaims"),
	}
	if entity != "" {
		req.setParam("entity", &[]string{entity})
	} else {
		req.setParam("claim", &[]string{claim})
	}
	return &req, nil
}

func (e *Entity) NewGetClaims() (*WikiDataGetClaimsRequest, error) {
	return NewGetClaims(e.ID, "")
}

// Param: property
func (r *WikiDataGetClaimsRequest) SetProperty(property string) {
	r.setParam("property", &[]string{property})
}

// Param: rank
// One of the following values: deprecated, normal, preferred
func (r *WikiDataGetClaimsRequest) SetRank(rank string) {
	r.setParam("rank", &[]string{rank})
}

// Param: props
// Default: references
func (r *WikiDataGetClaimsRequest) SetProps(props []string) {
	r.setParam("props", &props)
}

func (r *WikiDataGetClaimsRequest) Get() (*map[string][]Claim, error) {
	responseData := GetClaimsResponse{}
	res, err := easyreq.Make("GET", r.URL, nil, "", "json", &responseData, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, errors.New("request failed with status code " + string(res.StatusCode()))
	}
	return &responseData.Claims, nil
}

// Returns a pointer to a list of strings.
// WikiData action: wbavailablebadges
// WikiData API page: https://www.wikidata.org/w/api.php?action=help&modules=wbavailablebadges
func GetAvailableBadges() (*[]string, error) {
	var data []string
	url := fmt.Sprintf(WikiDataAPIURL, "wbavailablebadges")
	res, err := easyreq.Make("GET", url, nil, "", "json", &data, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, errors.New("request failed with status code " + string(res.StatusCode()))
	}
	return &data, nil
}

// Create a request
// Both search and language are required
// WikiData action: wbsearchentities
// WikiData API page: https://www.wikidata.org/w/api.php?action=help&modules=wbsearchentities
func NewSearch(search, language string) (*WikiDataSearchEntitiesRequest, error) {
	req := WikiDataSearchEntitiesRequest{
		URL: fmt.Sprintf(WikiDataAPIURL, "wbsearchentities"),
		Language: language,
		// default api value
		Limit: 7,
	}
	req.setParam("search", &[]string{search})
	req.setParam("language", &[]string{language})
	return &req, nil
}

// Param: limit
// Default: 7
func (r *WikiDataSearchEntitiesRequest) SetLimit(limit int) {
	r.Limit = limit
	r.setParam("limit", &[]string{string(limit)})
}

// Param: strictlanguage
func (r *WikiDataSearchEntitiesRequest) SetStrictLanguage(strictLanguage bool) {
	if strictLanguage {
		r.URL += "&strictlanguage="
	}
}

// Param: type
// One of the following values: item, property, lexeme, form, sense
// Default: item
func (r *WikiDataSearchEntitiesRequest) SetType(t string) {
	r.setParam("type", &[]string{t})
}

// Param: props
// Default: url
func (r *WikiDataSearchEntitiesRequest) SetProps(props []string) {
	r.setParam("props", &props)
}

// Param: continue
// Default: 0
func (r *WikiDataSearchEntitiesRequest) SetContinue(c int) {
	r.setParam("continue", &[]string{string(c)})
}

func (r *WikiDataSearchEntitiesRequest) Get() (*SearchEntitiesResponse, error) {
	responseData := SearchEntitiesResponse{}
	res, err := easyreq.Make("GET", r.URL, nil, "", "json", &responseData, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, errors.New("request failed with status code " + string(res.StatusCode()))
	}

	responseData.SearchRequest = *r
	return &responseData, nil
}

// Next page of results for search
// Will use the same configurations as previous request
func (r *SearchEntitiesResponse) Next() (*SearchEntitiesResponse, error) {
	res, err := NewSearch(r.SearchRequest.Search, r.SearchRequest.Language)
	if err != nil {
		return nil, err
	}
	if r.SearchRequest.Limit > 0 {
		res.SetLimit(r.SearchRequest.Limit)
	}
	if r.SearchRequest.Type != "" {
		res.SetType(r.SearchRequest.Type)
	}
	if len(r.SearchRequest.Props) > 0 {
		res.SetProps(r.SearchRequest.Props)
	}
	if r.SearchRequest.StrictLanguage {
		res.SetStrictLanguage(true)
	}

	res.SetContinue(r.CurrentContinue + r.SearchRequest.Limit)
	response, err := res.Get()
	return response, err
}


func ImageResizer(imageName string, size int) string {
	return fmt.Sprintf(ImageResizerURL, size, imageName)
}
