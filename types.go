package gowikidata

// WikiDataGetEntitiesRequest stores entities request url
type WikiDataGetEntitiesRequest struct {
	URL string
}

// WikiDataGetClaimsRequest stores claims request url
type WikiDataGetClaimsRequest struct {
	URL string
}

// WikiDataSearchEntitiesRequest stores parameters for entities search
type WikiDataSearchEntitiesRequest struct {
	URL            string
	Limit          int
	Language       string
	Type           string
	Props          []string
	StrictLanguage bool
	Search         string
}

// Entity represents wikidata entities data
type Entity struct {
	ID           string                 `json:"id"`
	PageID       int                    `json:"pageid"`
	NS           int                    `json:"ns"`
	Title        string                 `json:"title"`
	LastRevID    int                    `json:"lastrevid"`
	Modified     string                 `json:"modified"`
	Type         string                 `json:"type"`
	Labels       map[string]Label       `json:"labels"`
	Descriptions map[string]Description `json:"descriptions"`
	Aliases      map[string][]Alias     `json:"aliases"`
	Claims       map[string][]Claim     `json:"claims"`
	SiteLinks    map[string]SiteLink    `json:"sitelinks"`
}

// Label represents wikidata labels data
type Label struct {
	Language    string `json:"language"`
	Value       string `json:"value"`
	ForLanguage string `json:"for-language"`
}

// Description represents wikidata descriptions data
type Description struct {
	Language    string `json:"language"`
	Value       string `json:"value"`
	ForLanguage string `json:"for-language"`
}

// Alias represents wikidata aliases data
type Alias struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// SiteLink represents wikidata site links data
type SiteLink struct {
	Site   string   `json:"site"`
	Title  string   `json:"title"`
	Badges []string `json:"badges"`
}

// Claim represents wikidata claims data
type Claim struct {
	ID              string            `json:"id"`
	Rank            string            `json:"rank"`
	Type            string            `json:"type"`
	MainSnak        Snak              `json:"mainsnak"`
	Qualifiers      map[string][]Snak `json:"qualifiers"`
	QualifiersOrder []string          `json:"qualifiers-order"`
}

// Snak represents wikidata snak values
type Snak struct {
	SnakType  string    `json:"snaktype"`
	Property  string    `json:"property"`
	Hash      string    `json:"hash"`
	DataType  string    `json:"datatype"`
	DataValue DataValue `json:"datavalue"`
}

// DataValue represents wikidata values
// Wikidata values can be either string or number
// It will store the data type so you can work with it accordingly
type DataValue struct {
	Type  string           `json:"type"`
	Value DynamicDataValue `json:"value"`
}

// DynamicDataValue represents wikidata values for DataValue struct
type DynamicDataValue struct {
	Data        interface{}
	S           string
	I           int
	ValueFields DataValueFields
	Type        string
}

// DataValueFields represents wikidata value fields
type DataValueFields struct {
	EntityType    string  `json:"entity-type"`
	NumericID     int     `json:"numeric-id"`
	ID            string  `json:"id"`
	Type          string  `json:"type"`
	Value         string  `json:"value"`
	Time          string  `json:"time"`
	Precision     float64 `json:"precision"`
	Before        int     `json:"before"`
	After         int     `json:"after"`
	TimeZone      int     `json:"timezone"`
	CalendarModel string  `json:"calendarmodel"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Globe         string  `json:"globe"`
	Amount        string  `json:"amount"`
	LowerBound    string  `json:"lowerbound"`
	UpperBound    string  `json:"upperbound"`
	Unit          string  `json:"unit"`
	Text          string  `json:"text"`
	Language      string  `json:"language"`
}

// Reference represents wikidata references
type Reference struct {
	Hash       string            `json:"hash"`
	Snaks      map[string][]Snak `json:"snaks"`
	SnaksOrder []string          `json:"snaks-order"`
}

// GetEntitiesResponse represents wikidata entities response
type GetEntitiesResponse struct {
	Entities map[string]Entity `json:"entities"`
	Success  uint              `json:"success"`
}

// GetClaimsResponse represents wikidata claims response
type GetClaimsResponse struct {
	Claims map[string][]Claim `json:"claims"`
}

// SearchEntity represents wikidata entities search
type SearchEntity struct {
	Repository  string      `json:"repository"`
	ID          string      `json:"id"`
	ConceptURI  string      `json:"concepturi"`
	Title       string      `json:"title"`
	PageID      int         `json:"pageid"`
	URL         string      `json:"url"`
	Label       string      `json:"label"`
	Description string      `json:"description"`
	Match       SearchMatch `json:"match"`
	DataType    string      `json:"datatype"`
}

// SearchMatch represents wikidata search match value
type SearchMatch struct {
	Type     string `json:"type"`
	Language string `json:"language"`
	Text     string `json:"text"`
}

// SearchInfo represents wikidata search info
type SearchInfo struct {
	Search string `json:"search"`
}

// SearchEntitiesResponse represents wikidata entities search response
type SearchEntitiesResponse struct {
	SearchInfo      SearchInfo     `json:"searchinfo"`
	SearchResult    []SearchEntity `json:"search"`
	SearchContinue  int            `json:"search-continue"`
	Success         uint           `json:"success"`
	CurrentContinue int
	SearchRequest   WikiDataSearchEntitiesRequest
}

// WikiPediaQuery represents wikipedia query
type WikiPediaQuery struct {
	BatchComplete string `json:"batchcomplete"`
	Query         struct {
		Pages map[string]struct {
			PageProps struct {
				WikiBaseItem string `json:"wikibase_item"`
			} `json:"pageprops"`
		} `json:"pages"`
	} `json:"query"`
}
