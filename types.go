package gowikidata

// Request type for wikidata entities
type WikiDataGetEntitiesRequest struct {
	URL string
}

// Request type for wikidata claims
type WikiDataGetClaimsRequest struct {
	URL string
}

// Request type for wikidata entities search
type WikiDataSearchEntitiesRequest struct {
	URL            string
	Limit          int
	Language       string
	Type           string
	Props          []string
	StrictLanguage bool
	Search         string
}

// wikidata entities struct
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

// wikidata labels struct
type Label struct {
	Language    string `json:"language"`
	Value       string `json:"value"`
	ForLanguage string `json:"for-language"`
}

// Description struct for wikidata descriptions
type Description struct {
	Language    string `json:"language"`
	Value       string `json:"value"`
	ForLanguage string `json:"for-language"`
}

// wikidata aliases struct
type Alias struct {
	Language string `json:"language"`
	Value    string `json:"value"`
}

// wikidata site links struct
type SiteLink struct {
	Site   string   `json:"site"`
	Title  string   `json:"title"`
	Badges []string `json:"badges"`
}

// wikidata claims struct
type Claim struct {
	ID              string            `json:"id"`
	Rank            string            `json:"rank"`
	Type            string            `json:"type"`
	MainSnak        Snak              `json:"mainsnak"`
	Qualifiers      map[string][]Snak `json:"qualifiers"`
	QualifiersOrder []string          `json:"qualifiers-order"`
}

// wikidata snak struct
type Snak struct {
	SnakType  string    `json:"snaktype"`
	Property  string    `json:"property"`
	Hash      string    `json:"hash"`
	DataType  string    `json:"datatype"`
	DataValue DataValue `json:"datavalue"`
}

// wikidata values can be either string or a number
// this struct will hold the value and the data
type DataValue struct {
	Type  string           `json:"type"`
	Value DynamicDataValue `json:"value"`
}

// the actual struct for wikidata values
type DynamicDataValue struct {
	Data        interface{}
	S           string
	I           int
	ValueFields DataValueFields
	Type        string
}

// wikidata value fields struct
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

// wikidata references struct
type Reference struct {
	Hash       string            `json:"hash"`
	Snaks      map[string][]Snak `json:"snaks"`
	SnaksOrder []string          `json:"snaks-order"`
}

// entities response struct
type GetEntitiesResponse struct {
	Entities map[string]Entity `json:"entities"`
	Success  uint              `json:"success"`
}

// claims response struct
type GetClaimsResponse struct {
	Claims map[string][]Claim `json:"claims"`
}

// entities search struct
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

// search match struct
type SearchMatch struct {
	Type     string `json:"type"`
	Language string `json:"language"`
	Text     string `json:"text"`
}

// search info struct
type SearchInfo struct {
	Search string `json:"search"`
}

// entities search response struct
type SearchEntitiesResponse struct {
	SearchInfo      SearchInfo     `json:"searchinfo"`
	SearchResult    []SearchEntity `json:"search"`
	SearchContinue  int            `json:"search-continue"`
	Success         uint           `json:"success"`
	CurrentContinue int
	SearchRequest   WikiDataSearchEntitiesRequest
}

// wikipedia query struct
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
