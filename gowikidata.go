package gowikidata

import (
	"errors"
	"fmt"
	"github.com/Navid2zp/easyreq"
)

const (
	wikipediaQueryURL = "https://en.wikipedia.org/w/api.php?action=query&prop=pageprops&titles=%s&format=json"
	wikiDataAPIURL    = "https://www.wikidata.org/w/api.php?action=%s&format=json"
	imageResizerURL   = "https://commons.wikimedia.org/w/thumb.php?width=%d&f=%s"
)

// GetPageItem returns Wikipedia page item ID
func GetPageItem(slug string) (string, error) {
	url := fmt.Sprintf(wikipediaQueryURL, slug)

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

// NewGetEntities creates and returns a new WikiDataGetEntitiesRequest
// Use this function to initialize a request.
// It will return a pointer to WikiDataGetEntitiesRequest.
// You can make configurations on response.
func NewGetEntities(ids []string) (*WikiDataGetEntitiesRequest, error) {
	if len(ids) < 1 {
		return nil, errors.New("no ids provided")
	}

	req := WikiDataGetEntitiesRequest{
		URL: fmt.Sprintf(wikiDataAPIURL, "wbgetentities"),
	}
	req.setParam("ids", &ids)
	return &req, nil
}

// SetSites sets sites parameter for entities request
func (r *WikiDataGetEntitiesRequest) SetSites(sites []string) *WikiDataGetEntitiesRequest {
	r.setParam("sites", &sites)
	return r
}

// SetTitles sets titles parameter for entities request
func (r *WikiDataGetEntitiesRequest) SetTitles(titles []string) *WikiDataGetEntitiesRequest {
	r.setParam("titles", &titles)
	return r
}

// SetRedirects sets redirects parameter for entities request
func (r *WikiDataGetEntitiesRequest) SetRedirects(redirect bool) *WikiDataGetEntitiesRequest {
	redirectString := "yes"
	if !redirect {
		redirectString = "no"
	}
	r.URL += "&redirects=" + redirectString
	return r
}

// SetProps sets props parameter for entities request
// Default: info|sitelinks|aliases|labels|descriptions|claims|datatype
func (r *WikiDataGetEntitiesRequest) SetProps(props []string) *WikiDataGetEntitiesRequest {
	r.setParam("props", &props)
	return r
}

// SetLanguages sets languages parameter for entities request
func (r *WikiDataGetEntitiesRequest) SetLanguages(languages []string) *WikiDataGetEntitiesRequest {
	r.setParam("languages", &languages)
	return r
}

// SetLanguageFallback sets languagefallback parameter for entities request
func (r *WikiDataGetEntitiesRequest) SetLanguageFallback(fallback bool) *WikiDataGetEntitiesRequest {
	if fallback {
		r.URL += "&languagefallback="
	}
	return r
}

// SetNormalize sets normalize parameter for entities request
func (r *WikiDataGetEntitiesRequest) SetNormalize(normalize bool) *WikiDataGetEntitiesRequest {
	if normalize {
		r.URL += "&normalize="
	}
	return r
}

// SetSiteFilter sets sitefilter parameter for entities request
func (r *WikiDataGetEntitiesRequest) SetSiteFilter(sites []string) *WikiDataGetEntitiesRequest {
	r.setParam("sitefilter", &sites)
	return r
}

// Get makes a entities request and returns the response or an error
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

// NewGetClaims creates a new claim request for the given entity or claim
// WikiData action: wbgetclaims
// WikiData API page: https://www.wikidata.org/w/api.php?action=help&modules=wbgetclaims
// Either entity or claim must be provided
func NewGetClaims(entity, claim string) (*WikiDataGetClaimsRequest, error) {
	if entity == "" && claim == "" {
		return nil, errors.New("either entity or claim must be provided")
	}

	req := WikiDataGetClaimsRequest{
		URL: fmt.Sprintf(wikiDataAPIURL, "wbgetclaims"),
	}
	if entity != "" {
		req.setParam("entity", &[]string{entity})
	} else {
		req.setParam("claim", &[]string{claim})
	}
	return &req, nil
}

// NewGetClaims creates a new claim request for an entity
func (e *Entity) NewGetClaims() (*WikiDataGetClaimsRequest, error) {
	return NewGetClaims(e.ID, "")
}

// SetProperty sets property parameter for claims request
func (r *WikiDataGetClaimsRequest) SetProperty(property string) *WikiDataGetClaimsRequest {
	r.setParam("property", &[]string{property})
	return r
}

// SetRank sets rank parameter for claims request
// One of the following values: deprecated, normal, preferred
func (r *WikiDataGetClaimsRequest) SetRank(rank string) *WikiDataGetClaimsRequest {
	r.setParam("rank", &[]string{rank})
	return r
}

// SetProps sets props parameter for claims request
// Default: references
func (r *WikiDataGetClaimsRequest) SetProps(props []string) *WikiDataGetClaimsRequest {
	r.setParam("props", &props)
	return r
}

// Get creates a new request for claims and returns the response or an error
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

// GetAvailableBadges returns a pointer to a list of strings.
// WikiData action: wbavailablebadges
// WikiData API page: https://www.wikidata.org/w/api.php?action=help&modules=wbavailablebadges
func GetAvailableBadges() ([]string, error) {
	var data struct{Badges []string}
	url := fmt.Sprintf(wikiDataAPIURL, "wbavailablebadges")
	res, err := easyreq.Make("GET", url, nil, "", "json", &data, nil)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, errors.New("request failed with status code " + string(res.StatusCode()))
	}
	return data.Badges, nil
}

// NewSearch creates a new request for entities search and returns response or an error
// Create a request
// Both search and language are required
// WikiData action: wbsearchentities
// WikiData API page: https://www.wikidata.org/w/api.php?action=help&modules=wbsearchentities
func NewSearch(search, language string) (*WikiDataSearchEntitiesRequest, error) {
	req := WikiDataSearchEntitiesRequest{
		URL:      fmt.Sprintf(wikiDataAPIURL, "wbsearchentities"),
		Language: language,
		// default api value
		Limit: 7,
	}
	req.setParam("search", &[]string{search})
	req.setParam("language", &[]string{language})
	return &req, nil
}

// SetLimit sets limit parameter
// Default: 7
func (r *WikiDataSearchEntitiesRequest) SetLimit(limit int) *WikiDataSearchEntitiesRequest {
	r.Limit = limit
	r.setParam("limit", &[]string{string(limit)})
	return r
}

// SetStrictLanguage sets strictlanguage parameter
func (r *WikiDataSearchEntitiesRequest) SetStrictLanguage(strictLanguage bool) *WikiDataSearchEntitiesRequest {
	if strictLanguage {
		r.URL += "&strictlanguage="
	}
	return r
}

// SetType sets type parameter
// One of the following values: item, property, lexeme, form, sense
// Default: item
func (r *WikiDataSearchEntitiesRequest) SetType(t string) *WikiDataSearchEntitiesRequest {
	r.setParam("type", &[]string{t})
	return r
}

// SetProps sets props parameter
// Default: url
func (r *WikiDataSearchEntitiesRequest) SetProps(props []string) *WikiDataSearchEntitiesRequest {
	r.setParam("props", &props)
	return r
}

// SetContinue sets continue parameter
// Default: 0
func (r *WikiDataSearchEntitiesRequest) SetContinue(c int) *WikiDataSearchEntitiesRequest {
	r.setParam("continue", &[]string{string(c)})
	return r
}

// Get makes a entity search request and returns the response or an error
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

// ImageResizer returns the url for resizing a wikimedia image to the given size
func ImageResizer(imageName string, size int) string {
	return fmt.Sprintf(imageResizerURL, size, imageName)
}
