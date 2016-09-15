package yahoo

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

// Price price information
type Price struct {
	Currency string `xml:"currency,attr,omitempty"`
	Value    string `xml:",chardata"`
}

// PriceLabel priceLabel information
type PriceLabel struct {
	TaxIncluded      string `xml:"taxIncluded,attr,omitempty"`
	FixedPrice       string `xml:",omitempty"`
	DefaultPrice     string `xml:",omitempty"`
	SalePrice        string `xml:",omitempty"`
	BaseFixedPrice   string `xml:",omitempty"`
	BaseDefaultPrice string `xml:",omitempty"`
	BaseSalePrice    string `xml:",omitempty"`
	PeriodStart      string `xml:",omitempty"`
	PeriodEnd        string `xml:",omitempty"`
}

// Review review information
type Review struct {
	Rate  string `xml:",omitempty"`
	Count string `xml:",omitempty"`
	URL   string `xml:"Url,omitempty"`
}

// Request request information
type Request struct {
	Query string
}

// Point point information
type Point struct {
	Amount          string `xml:",omitempty"`
	Times           string `xml:",omitempty"`
	PremiumAmount   string `xml:",omitempty"`
	PremiumTimes    string `xml:",omitempty"`
	PremiumCpAmount string `xml:",omitempty"`
	AppCpAmount     string `xml:",omitempty"`
	PreAppCpAmount  string `xml:",omitempty"`
	PremiumCpTimes  string `xml:",omitempty"`
	AppCpTimes      string `xml:",omitempty"`
	PreAppCpTimes   string `xml:",omitempty"`
	Grant           string `xml:",omitempty"`
	Accept          string `xml:",omitempty"`
}

// Shipping shipping information
type Shipping struct {
	Code string `xml:",omitempty"`
	Name string `xml:",omitempty"`
}

// Title title information
type Title struct {
	Short  string `xml:",omitempty"`
	Medium string `xml:",omitempty"`
	Long   string `xml:",omitempty"`
	Name   string `xml:",omitempty"`
}

// Current current information
type Current struct {
	ID       string `xml:"Id,omitempty"`
	Name     string `xml:",omitempty"`
	ParentID string `xml:"ParentId,omitempty"`
	URL      string `xml:"Url,omitempty"`
	Title    Title  `xml:",omitempty"`
	Path     Path   `xml:",omitempty"`
	IsAdult  string `xml:",omitempty"`
}

// Category category information
type Category struct {
	Depth    string `xml:"depth,attr,omitempty"`
	ID       string `xml:"Id,omitempty"`
	Name     string `xml:",omitempty"`
	ParentID string `xml:"ParentId,omitempty"`
	Title    Title  `xml:",omitempty"`
}

// BrandsPath brandsPath information
type BrandsPath struct {
	Brand []Brand `xml:",omitempty"`
}

// Brands brands information
type Brands struct {
	Name string     `xml:",omitempty"`
	Path BrandsPath `xml:",omitempty"`
}

// Brand brand information
type Brand struct {
	ID string `xml:"Id,omitempty"`
}

// Method method information
type Method struct {
	Code string `xml:",omitempty"`
	Name string `xml:",omitempty"`
}

// Ratings ratings information
type Ratings struct {
	Rate       string `xml:",omitempty"`
	Count      string `xml:",omitempty"`
	Total      string `xml:",omitempty"`
	DetailRate string `xml:",omitempty"`
	Average    string `xml:",omitempty"`
}

// Image image information
type Image struct {
	Index    string `xml:"index,attr,omitempty"`
	ID       string `xml:"Id,omitempty"`
	Medium   string `xml:",omitempty"`
	Small    string `xml:",omitempty"`
	Original string `xml:",omitempty"`
	Width    string `xml:",omitempty"`
	Height   string `xml:",omitempty"`
	URL      string `xml:"Url,omitempty"`
}

// ExImage exImage information
type ExImage struct {
	URL    string `xml:"Url,omitempty"`
	Width  string `xml:",omitempty"`
	Height string `xml:",omitempty"`
}

// RelatedImages relatedImages information
type RelatedImages struct {
	Image []Image `xml:",omitempty"`
}

// Prefectures prefectures information
type Prefectures struct {
	Prefecture []Prefecture `xml:",omitempty"`
}

// Area area information
type Area struct {
	Code        string      `xml:",omitempty"`
	Name        string      `xml:",omitempty"`
	Prefectures Prefectures `xml:",omitempty"`
}

// Prefecture prefecture information
type Prefecture struct {
	Code string `xml:",omitempty"`
	Name string `xml:",omitempty"`
}

// Areas areas information
type Areas struct {
	Area []Area `xml:",omitempty"`
}

// Delivery delivery information
type Delivery struct {
	Areas      Areas  `xml:",omitempty"`
	Deadline   string `xml:",omitempty"`
	Conditions string `xml:",omitempty"`
}

// Payment payment information
type Payment struct {
	Method []Method `xml:",omitempty"`
}

// Store store information
type Store struct {
	ID               string   `xml:"Id,omitempty"`
	Name             string   `xml:",omitempty"`
	SellerType       string   `xml:",omitempty"`
	ToolType         string   `xml:",omitempty"`
	URL              string   `xml:"Url,omitempty"`
	Payment          Payment  `xml:",omitempty"`
	IsBestStore      string   `xml:",omitempty"`
	Ratings          Ratings  `xml:",omitempty"`
	Image            Image    `xml:",omitempty"`
	SameDayDelivery  Delivery `xml:",omitempty"`
	ExpressDelivery  Delivery `xml:",omitempty"`
	Point            Point    `xml:",omitempty"`
	InventoryMessage string   `xml:",omitempty"`
}

// Affiliate affiliate information
type Affiliate struct {
	Rate string `xml:",omitempty"`
}

// CategoryCurrent categoryCurrent information
type CategoryCurrent struct {
	Current Current `xml:",omitempty"`
}

// CategoryIDPath categoryIDPath information
type CategoryIDPath struct {
	Category []Category `xml:",omitempty"`
}

// ProductCategory productCategory information
type ProductCategory struct {
	ID string `xml:",omitempty"`
}

// Inventories inventories information
type Inventories struct {
	Inventory Inventory `xml:",omitempty"`
}

// Inventory inventory information
type Inventory struct {
	SubCode        string `xml:",omitempty"`
	Order          string `xml:",omitempty"`
	Availability   string `xml:",omitempty"`
	Quantity       string `xml:",omitempty"`
	AllowOverdraft string `xml:",omitempty"`
}

// EventTerm eventTerm information
type EventTerm struct {
	Start string `xml:",omitempty"`
	End   string `xml:",omitempty"`
}

// SubTitle SubTitle information
type SubTitle struct {
	Title string `xml:",omitempty"`
	URL   string `xml:"Url,omitempty"`
}

// SubTitles subTitles information
type SubTitles struct {
	SubTitle []SubTitle `xml:",omitempty"`
}

// Hit hit information
type Hit struct {
	Index                 string          `xml:"index,attr"`
	Title                 string          `xml:",omitempty"`
	EventTerm             EventTerm       `xml:",omitempty"`
	SubTitles             SubTitles       `xml:",omitempty"`
	Name                  string          `xml:",omitempty"`
	Description           string          `xml:",omitempty"`
	Headline              string          `xml:",omitempty"`
	URL                   string          `xml:"Url,omitempty"`
	ReleaseDate           string          `xml:",omitempty"`
	Availability          string          `xml:",omitempty"`
	Code                  string          `xml:",omitempty"`
	Condition             string          `xml:",omitempty"`
	PersonID              string          `xml:"PersonId,omitempty"`
	ProductID             string          `xml:"ProductId,omitempty"`
	Image                 Image           `xml:",omitempty"`
	ExImage               ExImage         `xml:",omitempty"`
	Review                Review          `xml:",omitempty"`
	Affiliate             Affiliate       `xml:",omitempty"`
	Price                 Price           `xml:",omitempty"`
	PriceLabel            PriceLabel      `xml:",omitempty"`
	Point                 Point           `xml:",omitempty"`
	Shipping              Shipping        `xml:",omitempty"`
	Category              CategoryCurrent `xml:",omitempty"`
	CategoryIDPath        CategoryIDPath  `xml:"CategoryIdPath,omitempty"`
	Brands                Brands          `xml:",omitempty"`
	JanCode               string          `xml:",omitempty"`
	Model                 string          `xml:",omitempty"`
	IsbnCode              string          `xml:",omitempty"`
	Store                 Store           `xml:",omitempty"`
	IsAdult               string          `xml:",omitempty"`
	Caption               string          `xml:",omitempty"`
	Abstract              string          `xml:",omitempty"`
	Additional1           string          `xml:",omitempty"`
	Additional2           string          `xml:",omitempty"`
	Additional3           string          `xml:",omitempty"`
	SpAdditional          string          `xml:",omitempty"`
	ProductCategory       ProductCategory `xml:",omitempty"`
	IsBargain             string          `xml:",omitempty"`
	OriginalPriceEvidence string          `xml:",omitempty"`
	RelatedImages         RelatedImages   `xml:",omitempty"`
	ShipWeight            string          `xml:",omitempty"`
	SaleLimit             string          `xml:",omitempty"`
	Inventories           Inventories     `xml:",omitempty"`
	Payment               Payment         `xml:",omitempty"`
	Order                 string          `xml:",omitempty"`
	IsCarBodySeller       string          `xml:",omitempty"`
}

// Range range information
type Range struct {
	From string `xml:",omitempty"`
	To   string `xml:",omitempty"`
	Hits string `xml:",omitempty"`
}

// Child child information
type Child struct {
	SortOrder string `xml:"sortOrder,attr,omitempty"`
	ID        string `xml:"Id,omitempty"`
	Name      string `xml:",omitempty"`
	Hits      string `xml:",omitempty"`
	URL       string `xml:"Url,omitempty"`
	Title     Title  `xml:",omitempty"`
	IsAdult   string `xml:",omitempty"`
}

// Path path information
type Path struct {
	Category []Category `xml:",omitempty"`
}

// Children children information
type Children struct {
	Child []Child `xml:",omitempty"`
}

// Subcategories subcategories information
type Subcategories struct {
	Path     Path     `xml:",omitempty"`
	Children Children `xml:",omitempty"`
}

// PriceRange priceRange information
type PriceRange struct {
	Range []Range `xml:",omitempty"`
}

// PriceRanges priceRanges information
type PriceRanges struct {
	Price PriceRange `xml:",omitempty"`
}

// Modules modules information
type Modules struct {
	PriceRanges   PriceRanges   `xml:",omitempty"`
	Subcategories Subcategories `xml:",omitempty"`
}

// ItemSearchResult result information
type ItemSearchResult struct {
	Request Request `xml:",omitempty"`
	Modules Modules `xml:",omitempty"`
	Hit     []Hit   `xml:",omitempty"`
}

// ItemSearchResultSet resultSet information
type ItemSearchResultSet struct {
	FirstResultPosition   string           `xml:"firstResultPosition,attr,omitempty"`
	TotalResultsAvailable string           `xml:"totalResultsAvailable,attr,omitempty"`
	TotalResultsReturned  string           `xml:"totalResultsReturned,attr,omitempty"`
	Result                ItemSearchResult `xml:",omitempty"`
}

// RankingInfo rankingInfo  information
type RankingInfo struct {
	LastModified string `xml:",omitempty"`
	StartDate    string `xml:",omitempty"`
	EndDate      string `xml:",omitempty"`
	CategoryID   string `xml:"CategoryId,omitempty"`
	Gender       string `xml:",omitempty"`
	Generation   string `xml:",omitempty"`
	Period       string `xml:",omitempty"`
	Type         string `xml:",omitempty"`
}

// RankingData rankingData information
type RankingData struct {
	Rank   string `xml:"rank,attr,omitempty"`
	Vector string `xml:"vector,attr,omitempty"`
	Type   string `xml:"type,attr,omitempty"`
	Name   string `xml:",omitempty"`
	Code   string `xml:",omitempty"`
	URL    string `xml:"Url,omitempty"`
	Image  Image  `xml:",omitempty"`
	Review Review `xml:",omitempty"`
	Store  Store  `xml:",omitempty"`
}

// Categories categories information
type Categories struct {
	Current  Current  `xml:",omitempty"`
	Children Children `xml:",omitempty"`
}

// Codes codes information
type Codes struct {
	Code string `xml:",omitempty"`
}

// ItemCode itemCode informaiton
type ItemCode struct {
	Codes Codes `xml:",omitempty"`
}

// RelationalTerm relationalTerm information
type RelationalTerm struct {
	Query string `xml:",omitempty"`
	URL   string `xml:"Url,omitempty"`
}

// Relational relational information
type Relational struct {
	RelationalTerm []RelationalTerm `xml:",omitempty"`
}

// QueryRankingData queryRankingData information
type QueryRankingData struct {
	Rank       string     `xml:"rank,attr,omitempty"`
	Prerank    string     `xml:"prerank,attr,omitempty"`
	Vector     string     `xml:"vector,attr,omitempty"`
	Score      string     `xml:",omitempty"`
	Query      string     `xml:",omitempty"`
	Relational Relational `xml:",omitempty"`
	URL        string     `xml:"Url,omitempty"`
}

// Images images information
type Images struct {
	Image []Image `xml:",omitempty"`
}

// Content content information
type Content struct {
	Index       string `xml:"index,attr,omitempty"`
	Device      string `xml:",omitempty"`
	Description string `xml:",omitempty"`
	URL         string `xml:"Url,omitempty"`
	ArchivedURL string `xml:"ArchivedUrl,omitempty"`
	Carrier     string `xml:",omitempty"`
}

// Contents contents information
type Contents struct {
	Content []Content `xml:",omitempty"`
}

// Ratio ratio information
type Ratio struct {
	Index string `xml:"index,attr,omitempty"`
	Type  string `xml:",omitempty"`
	Num   string `xml:",omitempty"`
}

// Ratios ratios information
type Ratios struct {
	Ratio []Ratio `xml:",omitempty"`
}

// Rule rule information
type Rule struct {
	Index  string `xml:"index,attr,omitempty"`
	Target string `xml:",omitempty"`
	Ratios Ratios `xml:",omitempty"`
}

// Rules rules information
type Rules struct {
	Rule []Rule `xml:",omitempty"`
}

// Count count information
type Count struct {
	All string `xml:",omitempty"`
}

// TargetImage targetImage information
type TargetImage struct {
	ID     string `xml:"Id,omitempty"`
	Small  Image  `xml:",omitempty"`
	Medium Image  `xml:",omitempty"`
}

// Target Target information
type Target struct {
	Type      string      `xml:"type,attr,omitempty"`
	Name      string      `xml:",omitempty"`
	Code      string      `xml:",omitempty"`
	URL       string      `xml:"Url,omitempty"`
	Store     Store       `xml:",omitempty"`
	Image     TargetImage `xml:",omitempty"`
	Condition string      `xml:",omitempty"`
}

// ReviewSearchResult result information
type ReviewSearchResult struct {
	Index       string  `xml:"index,attr,omitempty"`
	ReviewTitle string  `xml:",omitempty"`
	Description string  `xml:",omitempty"`
	Ratings     Ratings `xml:",omitempty"`
	Count       Count   `xml:",omitempty"`
	Recommend   string  `xml:",omitempty"`
	Update      string  `xml:",omitempty"`
	URL         string  `xml:"Url,omitempty"`
	ReviewType  string  `xml:",omitempty"`
	Purpose     string  `xml:",omitempty"`
	SendTo      string  `xml:",omitempty"`
	Target      Target  `xml:",omitempty"`
}

// ReviewSearchResultSet resultSet information
type ReviewSearchResultSet struct {
	FirstResultPosition   string               `xml:"firstResultPosition,attr,omitempty"`
	TotalResultsAvailable string               `xml:"totalResultsAvailable,attr,omitempty"`
	TotalResultsReturned  string               `xml:"totalResultsReturned,attr,omitempty"`
	Result                []ReviewSearchResult `xml:",omitempty"`
	Categories            Categories           `xml:",omitempty"`
}

// ShopCampaignSearchResult result information
type ShopCampaignSearchResult struct {
	Index             string   `xml:"index,attr,omitempty"`
	Title             string   `xml:",omitempty"`
	ID                string   `xml:"Id,omitempty"`
	EventType         string   `xml:",omitempty"`
	IsDisplay         string   `xml:",omitempty"`
	IsCondition       string   `xml:",omitempty"`
	Sentence          string   `xml:",omitempty"`
	CampaignStartTime string   `xml:",omitempty"`
	CampaignEndTime   string   `xml:",omitempty"`
	EntryStartTime    string   `xml:",omitempty"`
	EntryEndTime      string   `xml:",omitempty"`
	DisplayStartTime  string   `xml:",omitempty"`
	DisplayEndTime    string   `xml:",omitempty"`
	PushStartTime     string   `xml:",omitempty"`
	PushEndTime       string   `xml:",omitempty"`
	IsAutoEntry       string   `xml:",omitempty"`
	Images            Images   `xml:",omitempty"`
	Contents          Contents `xml:",omitempty"`
	Rules             Rules    `xml:",omitempty"`
	MaxRatio          string   `xml:",omitempty"`
}

// ShopCampaignSearchResultSet resultSet information
type ShopCampaignSearchResultSet struct {
	Result []ShopCampaignSearchResult `xml:",omitempty"`
}

// GetModuleResult result information
type GetModuleResult struct {
	CategoryID  string `xml:"CategoryId,omitempty"`
	Position    string `xml:",omitempty"`
	ModuleTitle string `xml:",omitempty"`
	Hit         []Hit  `xml:",omitempty"`
}

// GetModuleResultSet resultSet information
type GetModuleResultSet struct {
	TotalResultsReturned string          `xml:"totalResultsReturned,attr,omitempty"`
	Result               GetModuleResult `xml:",omitempty"`
}

// QueryRankingResult result information
type QueryRankingResult struct {
	RankingInfo      RankingInfo        `xml:",omitempty"`
	QueryRankingData []QueryRankingData `xml:",omitempty"`
	Categories       Categories         `xml:",omitempty"`
}

// QueryRankingResultSet resultSet information
type QueryRankingResultSet struct {
	FirstResultPosition   string             `xml:"firstResultPosition,attr,omitempty"`
	TotalResultsAvailable string             `xml:"totalResultsAvailable,attr,omitempty"`
	TotalResultsReturned  string             `xml:"totalResultsReturned,attr,omitempty"`
	Result                QueryRankingResult `xml:",omitempty"`
}

// ItemLookupResult result information
type ItemLookupResult struct {
	ItemCode ItemCode `xml:",omitempty"`
	Hit      Hit      `xml:",omitempty"`
}

// ItemLookupResultSet resultSet information
type ItemLookupResultSet struct {
	FirstResultPosition  string           `xml:"firstResultPosition,attr,omitempty"`
	TotalResultsReturned string           `xml:"totalResultsReturned,attr,omitempty"`
	Result               ItemLookupResult `xml:",omitempty"`
}

// CategoryRankingResult result information
type CategoryRankingResult struct {
	RankingInfo RankingInfo   `xml:",omitempty"`
	RankingData []RankingData `xml:",omitempty"`
	Categories  Categories    `xml:",omitempty"`
}

// CategoryRankingResultSet resultSet information
type CategoryRankingResultSet struct {
	FirstResultPosition   string                `xml:"firstResultPosition,attr,omitempty"`
	TotalResultsAvailable string                `xml:"totalResultsAvailable,attr,omitempty"`
	TotalResultsReturned  string                `xml:"totalResultsReturned,attr,omitempty"`
	Result                CategoryRankingResult `xml:",omitempty"`
}

// CategorySearchResult result information
type CategorySearchResult struct {
	Categories Categories `xml:",omitempty"`
}

// CategorySearchResultSet resultSet information
type CategorySearchResultSet struct {
	TotalResultsReturned string               `xml:"totalResultsReturned,attr,omitempty"`
	Result               CategorySearchResult `xml:",omitempty"`
}

// CategorySearchParam parameters for CategorySearch
type CategorySearchParam struct {
	Output        string `json:"output,omitempty"`
	AffiliateType string `json:"affiliate_type,omitempty"`
	AffiliateID   string `json:"affiliate_id,omitempty"`
	Callback      string `json:"callback,omitempty"`
	CategoryID    string `json:"category_id,omitempty"`
}

// CategoryRankingParam parameters for CategoryRanking
type CategoryRankingParam struct {
	Output        string `json:"output,omitempty"`
	AffiliateType string `json:"affiliate_type,omitempty"`
	AffiliateID   string `json:"affiliate_id,omitempty"`
	Callback      string `json:"callback,omitempty"`
	CategoryID    string `json:"category_id,omitempty"`
	Gender        string `json:"gender,omitempty"`
	Generation    string `json:"generation,omitempty"`
	Period        string `json:"period,omitempty"`
	Offset        string `json:"offset,omitempty"`
	Type          string `json:"type,omitempty"`
}

// ItemSearchParam parameters for ItemSearch
type ItemSearchParam struct {
	AffiliateType       string `json:"affiliate_type,omitempty"`
	AffiliateID         string `json:"affiliate_id,omitempty"`
	Callback            string `json:"callback,omitempty"`
	Query               string `json:"query,omitempty"`
	Type                string `json:"type,omitempty"`
	Jan                 string `json:"jan,omitempty"`
	Isbn                string `json:"isbn,omitempty"`
	ImageSize           string `json:"image_size,omitempty"`
	CategoryID          string `json:"category_id,omitempty"`
	ProductID           string `json:"product_id,omitempty"`
	PersonID            string `json:"person_id,omitempty"`
	BrandID             string `json:"brand_id,omitempty"`
	StoreID             string `json:"store_id,omitempty"`
	PriceFrom           string `json:"price_from,omitempty"`
	PriceTo             string `json:"price_to,omitempty"`
	AffiliateFrom       string `json:"affiliate_from,omitempty"`
	AffiliateTo         string `json:"affiliate_to,omitempty"`
	Preorder            string `json:"preorder,omitempty"`
	Hits                string `json:"hits,omitempty"`
	Offset              string `json:"offset,omitempty"`
	Module              string `json:"module,omitempty"`
	Availability        string `json:"availability,omitempty"`
	Discount            string `json:"discount,omitempty"`
	Shipping            string `json:"shipping,omitempty"`
	Payment             string `json:"payment,omitempty"`
	License             string `json:"license,omitempty"`
	SalestartFrom       string `json:"salestart_from,omitempty"`
	SalestartTo         string `json:"salestart_to,omitempty"`
	SaleendFrom         string `json:"saleend_from,omitempty"`
	SaleendTo           string `json:"saleend_to,omitempty"`
	ExpArea             string `json:"exp_area,omitempty"`
	ExpDeadlineFrom     string `json:"exp_deadline_from,omitempty"`
	ExpDeadlineTo       string `json:"exp_deadline_to,omitempty"`
	SameDayArea         string `json:"same_day_area,omitempty"`
	SameDayDeadlineFrom string `json:"same_day_deadline_from,omitempty"`
	SameDayDeadlineTo   string `json:"same_day_deadline_to,omitempty"`
	Seller              string `json:"seller,omitempty"`
}

// ItemLookupParam parameters for ItemLookup
type ItemLookupParam struct {
	AffiliateType string `json:"affiliate_type,omitempty"`
	AffiliateID   string `json:"affiliate_id,omitempty"`
	Callback      string `json:"callback,omitempty"`
	ItemCode      string `json:"itemcode,omitempty"`
	ResponseGroup string `json:"responsegroup,omitempty"`
	ImageSize     string `json:"image_size,omitempty"`
	License       string `json:"license,omitempty"`
}

// QueryRankingParam parameters for QueryRanking
type QueryRankingParam struct {
	AffiliateType string `json:"affiliate_type,omitempty"`
	AffiliateID   string `json:"affiliate_id,omitempty"`
	Callback      string `json:"callback,omitempty"`
	Type          string `json:"type,omitempty"`
	Hits          string `json:"hits,omitempty"`
	Offset        string `json:"offset,omitempty"`
	CategoryID    string `json:"category_id,omitempty"`
}

// GetModuleParam parameters for GetModule
type GetModuleParam struct {
	AffiliateType string `json:"affiliate_type,omitempty"`
	AffiliateID   string `json:"affiliate_id,omitempty"`
	Callback      string `json:"callback,omitempty"`
	CategoryID    string `json:"category_id,omitempty"`
	Position      string `json:"type,omitempty"`
}

// ShopCampaignSearchParam parameters for ShopCampaignSearch
type ShopCampaignSearchParam struct {
	AffiliateType string `json:"affiliate_type,omitempty"`
	AffiliateID   string `json:"affiliate_id,omitempty"`
	Callback      string `json:"callback,omitempty"`
}

// ReviewSearchParam parameters for ReviewSearch
type ReviewSearchParam struct {
	AffiliateType string `json:"affiliate_type,omitempty"`
	AffiliateID   string `json:"affiliate_id,omitempty"`
	Callback      string `json:"callback,omitempty"`
	Jan           string `json:"jan,omitempty"`
	CategoryID    string `json:"category_id,omitempty"`
	ProductID     string `json:"product_id,omitempty"`
	PersonID      string `json:"person_id,omitempty"`
	StoreID       string `json:"store_id,omitempty"`
	Results       string `json:"results,omitempty"`
	Start         string `json:"start,omitempty"`
	Sort          string `json:"sort,omitempty"`
}

// ReviewSearch search review
func (c *Client) ReviewSearch(param *ReviewSearchParam) (*ReviewSearchResultSet, error) {
	v := url.Values{}
	if p := param.AffiliateType; p != "" {
		v.Set("affiliate_type", p)
	}
	if p := param.AffiliateID; p != "" {
		v.Set("affiliate_id", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}
	if p := param.Jan; p != "" {
		v.Set("jan", p)
	}
	if p := param.CategoryID; p != "" {
		v.Set("category_id", p)
	}
	if p := param.ProductID; p != "" {
		v.Set("product_id", p)
	}
	if p := param.PersonID; p != "" {
		v.Set("person_id", p)
	}
	if p := param.StoreID; p != "" {
		v.Set("store_id", p)
	}
	if p := param.Results; p != "" {
		v.Set("results", p)
	}
	if p := param.Start; p != "" {
		v.Set("start", p)
	}
	if p := param.Sort; p != "" {
		v.Set("sort", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/ShoppingWebService/V1/reviewSearch").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data ReviewSearchResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// ShopCampaignSearch search shop campaign
func (c *Client) ShopCampaignSearch(param *ShopCampaignSearchParam) (*ShopCampaignSearchResultSet, error) {
	v := url.Values{}
	if p := param.AffiliateType; p != "" {
		v.Set("affiliate_type", p)
	}
	if p := param.AffiliateID; p != "" {
		v.Set("affiliate_id", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/ShoppingWebService/V1/shopCampaignSearch").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data ShopCampaignSearchResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// GetModule get module
func (c *Client) GetModule(param *GetModuleParam) (*GetModuleResultSet, error) {
	v := url.Values{}
	if p := param.AffiliateType; p != "" {
		v.Set("affiliate_type", p)
	}
	if p := param.AffiliateID; p != "" {
		v.Set("affiliate_id", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}
	if p := param.CategoryID; p != "" {
		v.Set("category_id", p)
	}
	if p := param.Position; p != "" {
		v.Set("position", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/ShoppingWebService/V1/getModule").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data GetModuleResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// QueryRanking ranking query
func (c *Client) QueryRanking(param *QueryRankingParam) (*QueryRankingResultSet, error) {
	v := url.Values{}
	if p := param.AffiliateType; p != "" {
		v.Set("affiliate_type", p)
	}
	if p := param.AffiliateID; p != "" {
		v.Set("affiliate_id", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}
	if p := param.Type; p != "" {
		v.Set("type", p)
	}
	if p := param.Hits; p != "" {
		v.Set("hits", p)
	}
	if p := param.Offset; p != "" {
		v.Set("offset", p)
	}
	if p := param.CategoryID; p != "" {
		v.Set("category_id", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/ShoppingWebService/V1/queryRanking").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data QueryRankingResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// ItemLookup lookup item
func (c *Client) ItemLookup(param *ItemLookupParam) (*ItemLookupResultSet, error) {
	v := url.Values{}
	if p := param.AffiliateType; p != "" {
		v.Set("affiliate_type", p)
	}
	if p := param.AffiliateID; p != "" {
		v.Set("affiliate_id", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}
	if p := param.ItemCode; p != "" {
		v.Set("itemcode", p)
	}
	if p := param.ResponseGroup; p != "" {
		v.Set("responsegroup", p)
	}
	if p := param.ImageSize; p != "" {
		v.Set("image_size", p)
	}
	if p := param.License; p != "" {
		v.Set("license", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/ShoppingWebService/V1/itemLookup").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data ItemLookupResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// CategorySearch search category
func (c *Client) CategorySearch(param *CategorySearchParam) (*CategorySearchResultSet, error) {
	v := url.Values{}
	if p := param.Output; p != "" {
		v.Set("output", p)
	}
	if p := param.AffiliateType; p != "" {
		v.Set("affiliate_type", p)
	}
	if p := param.AffiliateID; p != "" {
		v.Set("affiliate_id", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}
	if p := param.CategoryID; p != "" {
		v.Set("category_id", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/ShoppingWebService/V1/categorySearch").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data CategorySearchResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// CategoryRanking ranking category
func (c *Client) CategoryRanking(param *CategoryRankingParam) (*CategoryRankingResultSet, error) {
	v := url.Values{}
	if p := param.Output; p != "" {
		v.Set("output", p)
	}
	if p := param.AffiliateType; p != "" {
		v.Set("affiliate_type", p)
	}
	if p := param.AffiliateID; p != "" {
		v.Set("affiliate_id", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}
	if p := param.CategoryID; p != "" {
		v.Set("category_id", p)
	}
	if p := param.Gender; p != "" {
		v.Set("gender", p)
	}
	if p := param.Generation; p != "" {
		v.Set("generation", p)
	}
	if p := param.Period; p != "" {
		v.Set("period", p)
	}
	if p := param.Offset; p != "" {
		v.Set("offset", p)
	}
	if p := param.Type; p != "" {
		v.Set("type", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/ShoppingWebService/V1/categoryRanking").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data CategoryRankingResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// ItemSearch search item
func (c *Client) ItemSearch(param *ItemSearchParam) (*ItemSearchResultSet, error) {
	v := url.Values{}
	if p := param.AffiliateType; p != "" {
		v.Set("affiliate_type", p)
	}
	if p := param.AffiliateID; p != "" {
		v.Set("affiliate_id", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}
	if p := param.Query; p != "" {
		v.Set("query", p)
	}
	if p := param.Type; p != "" {
		v.Set("type", p)
	}
	if p := param.Jan; p != "" {
		v.Set("jan", p)
	}
	if p := param.Isbn; p != "" {
		v.Set("isbn", p)
	}
	if p := param.ImageSize; p != "" {
		v.Set("image_size", p)
	}
	if p := param.CategoryID; p != "" {
		v.Set("category_id", p)
	}
	if p := param.ProductID; p != "" {
		v.Set("product_id", p)
	}
	if p := param.PersonID; p != "" {
		v.Set("person_id", p)
	}
	if p := param.BrandID; p != "" {
		v.Set("brand_id", p)
	}
	if p := param.StoreID; p != "" {
		v.Set("store_id", p)
	}
	if p := param.PriceFrom; p != "" {
		v.Set("price_from", p)
	}
	if p := param.PriceTo; p != "" {
		v.Set("price_to", p)
	}
	if p := param.AffiliateFrom; p != "" {
		v.Set("affiliate_from", p)
	}
	if p := param.AffiliateTo; p != "" {
		v.Set("affiliate_to", p)
	}
	if p := param.Preorder; p != "" {
		v.Set("preorder", p)
	}
	if p := param.Hits; p != "" {
		v.Set("hits", p)
	}
	if p := param.Offset; p != "" {
		v.Set("offset", p)
	}
	if p := param.Module; p != "" {
		v.Set("module", p)
	}
	if p := param.Availability; p != "" {
		v.Set("availability", p)
	}
	if p := param.Discount; p != "" {
		v.Set("discount", p)
	}
	if p := param.Shipping; p != "" {
		v.Set("shipping", p)
	}
	if p := param.Payment; p != "" {
		v.Set("payment", p)
	}
	if p := param.License; p != "" {
		v.Set("license", p)
	}
	if p := param.SalestartFrom; p != "" {
		v.Set("salestart_from", p)
	}
	if p := param.SalestartTo; p != "" {
		v.Set("salestart_to", p)
	}
	if p := param.SaleendFrom; p != "" {
		v.Set("saleend_from", p)
	}
	if p := param.SaleendTo; p != "" {
		v.Set("saleend_to", p)
	}
	if p := param.ExpArea; p != "" {
		v.Set("exp_area", p)
	}
	if p := param.ExpDeadlineFrom; p != "" {
		v.Set("exp_deadline_from", p)
	}
	if p := param.ExpDeadlineTo; p != "" {
		v.Set("exp_deadline_to", p)
	}
	if p := param.SameDayArea; p != "" {
		v.Set("same_day_area", p)
	}
	if p := param.SameDayDeadlineFrom; p != "" {
		v.Set("same_day_deadline_from", p)
	}
	if p := param.SameDayDeadlineTo; p != "" {
		v.Set("same_day_deadline_to", p)
	}
	if p := param.Seller; p != "" {
		v.Set("seller", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/ShoppingWebService/V1/itemSearch").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data ItemSearchResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}
