package yahoo

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

// ChildCategory childCategory information
type ChildCategory struct {
	CategoryID       string `xml:"CategoryId,omitempty"`
	CategoryName     string `xml:",omitempty"`
	CategoryPath     string `xml:",omitempty"`
	CategoryIDPath   string `xml:"CategoryIdPath,omitempty"`
	NumOfAuctions    string `xml:",omitempty"`
	ParentCategoryID string `xml:"ParentCategoryId,omitempty"`
	IsLeaf           string `xml:",omitempty"`
	Depth            string `xml:",omitempty"`
	Order            string `xml:",omitempty"`
	IsLink           string `xml:",omitempty"`
	IsAdult          string `xml:",omitempty"`
	IsLeafToLink     string `xml:",omitempty"`
}

// CategoryTreeResult result information
type CategoryTreeResult struct {
	CategoryID       string          `xml:"CategoryId,omitempty"`
	CategoryName     string          `xml:",omitempty"`
	CategoryPath     string          `xml:",omitempty"`
	CategoryIDPath   string          `xml:"CategoryIdPath,omitempty"`
	ParentCategoryID string          `xml:"ParentCategoryId,omitempty"`
	IsLeaf           string          `xml:",omitempty"`
	Depth            string          `xml:",omitempty"`
	Order            string          `xml:",omitempty"`
	IsLink           string          `xml:",omitempty"`
	IsAdult          string          `xml:",omitempty"`
	IsLeafToLink     string          `xml:",omitempty"`
	ChildCategoryNum string          `xml:",omitempty"`
	ChildCategory    []ChildCategory `xml:",omitempty"`
}

// CategoryTreeResultSet resultSet information
type CategoryTreeResultSet struct {
	FirstResultPosition   string             `xml:"firstResultPosition,attr,omitempty"`
	TotalResultsAvailable string             `xml:"totalResultsAvailable,attr,omitempty"`
	TotalResultsReturned  string             `xml:"totalResultsReturned,attr,omitempty"`
	Result                CategoryTreeResult `xml:",omitempty"`
}

// Rating rating information
type Rating struct {
	Point             string `xml:",omitempty"`
	IsSuspended       string `xml:",omitempty"`
	IsDeleted         string `xml:",omitempty"`
	TotalGoodRating   string `xml:",omitempty"`
	TotalNormalRating string `xml:",omitempty"`
	TotalBadRating    string `xml:",omitempty"`
}

// Seller Seller information
type Seller struct {
	ID          string `xml:"Id,omitempty"`
	ItemListURL string `xml:"ItemListUrl,omitempty"`
	RatingURL   string `xml:"RatingUrl,omitempty"`
	AboutURL    string `xml:"AboutUrl,omitempty"`
	Rating      Rating `xml:",omitempty"`
}

// AuctionImage auctionImage information
type AuctionImage struct {
	Width  string `xml:"width,attr,omitempty"`
	Height string `xml:"height,attr,omitempty"`
	Alt    string `xml:"alt,attr,omitempty"`
	Value  string `xml:",chardata"`
}

// CharityOption charityOption information
type CharityOption struct {
	Proportion string `xml:",omitempty"`
}

// Option Option information
type Option struct {
	FeaturedIcon         string `xml:",omitempty"`
	EasyPaymentIcon      string `xml:",omitempty"`
	IsBold               string `xml:",omitempty"`
	IsBackGroundColor    string `xml:",omitempty"`
	IsOffer              string `xml:",omitempty"`
	IsCharity            string `xml:",omitempty"`
	NewIcon              string `xml:",omitempty"`
	StoreIcon            string `xml:",omitempty"`
	CheckIcon            string `xml:",omitempty"`
	PublicIcon           string `xml:",omitempty"`
	FreeshippingIcon     string `xml:",omitempty"`
	NewItemIcon          string `xml:",omitempty"`
	WrappingIcon         string `xml:",omitempty"`
	BuynowIcon           string `xml:",omitempty"`
	GiftIcon             string `xml:",omitempty"`
	PointIcon            string `xml:",omitempty"`
	CharityOptionIcon    string `xml:",omitempty"`
	IsTradingNaviAuction string `xml:",omitempty"`
}

// Item item information
type Item struct {
	AuctionID        string        `xml:",omitempty"`
	Title            string        `xml:",omitempty"`
	CategoryID       string        `xml:"CategoryId,omitempty"`
	Seller           Seller        `xml:",omitempty"`
	ItemURL          string        `xml:"ItemUrl,omitempty"`
	AuctionItemURL   string        `xml:"AuctionItemUrl,omitempty"`
	Image            AuctionImage  `xml:",omitempty"`
	OriginalImageNum string        `xml:",omitempty"`
	CurrentPrice     string        `xml:",omitempty"`
	Bids             string        `xml:",omitempty"`
	EndTime          string        `xml:",omitempty"`
	BidOrBuy         string        `xml:",omitempty"`
	Affiliate        Affiliate     `xml:",omitempty"`
	IsReserved       string        `xml:",omitempty"`
	CharityOption    CharityOption `xml:",omitempty"`
	Option           Option        `xml:",omitempty"`
	IsAdult          string        `xml:",omitempty"`
}

// CategoryLeafResult result information
type CategoryLeafResult struct {
	CategoryPath   string `xml:",omitempty"`
	CategoryIDPath string `xml:"CategoryIdPath,omitempty"`
	Item           []Item `xml:",omitempty"`
}

// CategoryLeafResultSet resultSet information
type CategoryLeafResultSet struct {
	FirstResultPosition   string             `xml:"firstResultPosition,attr,omitempty"`
	TotalResultsAvailable string             `xml:"totalResultsAvailable,attr,omitempty"`
	TotalResultsReturned  string             `xml:"totalResultsReturned,attr,omitempty"`
	Result                CategoryLeafResult `xml:",omitempty"`
}

// SellingListResult result information
type SellingListResult struct {
	Seller Seller `xml:",omitempty"`
	Item   []Item `xml:",omitempty"`
}

// SellingListResultSet resultSet information
type SellingListResultSet struct {
	FirstResultPosition   string            `xml:"firstResultPosition,attr,omitempty"`
	TotalResultsAvailable string            `xml:"totalResultsAvailable,attr,omitempty"`
	TotalResultsReturned  string            `xml:"totalResultsReturned,attr,omitempty"`
	Result                SellingListResult `xml:",omitempty"`
}

// SearchResult result information
type SearchResult struct {
	UnitsWord []string `xml:",omitempty"`
	Item      []Item   `xml:",omitempty"`
}

// SearchResultSet resultSet information
type SearchResultSet struct {
	FirstResultPosition   string       `xml:"firstResultPosition,attr,omitempty"`
	TotalResultsAvailable string       `xml:"totalResultsAvailable,attr,omitempty"`
	TotalResultsReturned  string       `xml:"totalResultsReturned,attr,omitempty"`
	Result                SearchResult `xml:",omitempty"`
}

// Img img information
type Img struct {
	Image1  AuctionImage `xml:",omitempty"`
	Image2  AuctionImage `xml:",omitempty"`
	Image3  AuctionImage `xml:",omitempty"`
	Image4  AuctionImage `xml:",omitempty"`
	Image5  AuctionImage `xml:",omitempty"`
	Image6  AuctionImage `xml:",omitempty"`
	Image7  AuctionImage `xml:",omitempty"`
	Image8  AuctionImage `xml:",omitempty"`
	Image9  AuctionImage `xml:",omitempty"`
	Image10 AuctionImage `xml:",omitempty"`
}

// Bidder bidder information
type Bidder struct {
	ID          string `xml:"Id,omitempty"`
	Rating      Rating `xml:",omitempty"`
	ItemListURL string `xml:",omitempty"`
	RatingURL   string `xml:",omitempty"`
}

// HighestBidders highestBidders information
type HighestBidders struct {
	TotalHighestBidders string   `xml:"totalHighestBidders,attr,omitempty"`
	IsMore              string   `xml:",omitempty"`
	Bidder              []Bidder `xml:",omitempty"`
}

// ItemStatus itemStatus information
type ItemStatus struct {
	Condition string `xml:",omitempty"`
	Comment   string `xml:",omitempty"`
}

// ItemReturnable ItemReturnable information
type ItemReturnable struct {
	Allowed string `xml:",omitempty"`
	Comment string `xml:",omitempty"`
}

// EasyPayment easyPayment information
type EasyPayment struct {
	IsCreditCard       string `xml:",omitempty"`
	IsNetBank          string `xml:",omitempty"`
	SafeKeepingPayment string `xml:",omitempty"`
}

// Bank bank information
type Bank struct {
	TotalBankMethodAvailable string `xml:"totalBankMethodAvailable,attr,omitempty"`
	Method                   string `xml:",omitempty"`
}

// Other other information
type Other struct {
	TotalOtherMethodAvailable string `xml:"totalOtherMethodAvailable,attr,omitempty"`
	Method                    string `xml:",omitempty"`
}

// AuctionPayment auctionPayment information
type AuctionPayment struct {
	EasyPayment      EasyPayment `xml:",omitempty"`
	Bank             Bank        `xml:",omitempty"`
	CashRegistration string      `xml:",omitempty"`
	PostalTransfer   string      `xml:",omitempty"`
	PostalOrder      string      `xml:",omitempty"`
	CashOnDelivery   string      `xml:",omitempty"`
	CreditCard       string      `xml:",omitempty"`
	Loan             string      `xml:",omitempty"`
	Other            Other       `xml:",omitempty"`
}

// BaggageInfo baggageInfo information
type BaggageInfo struct {
	Size        string `xml:",omitempty"`
	SizeIndex   string `xml:",omitempty"`
	Weight      string `xml:",omitempty"`
	WeightIndex string `xml:",omitempty"`
}

// AuctionItemResult result information
type AuctionItemResult struct {
	AuctionID                 string         `xml:",omitempty"`
	CategoryID                string         `xml:",omitempty"`
	CategoryFarm              string         `xml:",omitempty"`
	CategoryIDPath            string         `xml:"CategoryIdPath,omitempty"`
	CategoryPath              string         `xml:",omitempty"`
	Title                     string         `xml:",omitempty"`
	Seller                    Seller         `xml:",omitempty"`
	AuctionItemURL            string         `xml:"AuctionItemUrl,omitempty"`
	Img                       Img            `xml:",omitempty"`
	Initprice                 string         `xml:",omitempty"`
	Price                     string         `xml:",omitempty"`
	Quantity                  string         `xml:",omitempty"`
	AvailableQuantity         string         `xml:",omitempty"`
	Bids                      string         `xml:",omitempty"`
	HighestBidders            HighestBidders `xml:",omitempty"`
	YPoint                    string         `xml:",omitempty"`
	ItemStatus                ItemStatus     `xml:",omitempty"`
	ItemReturnable            ItemReturnable `xml:",omitempty"`
	StartTime                 string         `xml:",omitempty"`
	EndTime                   string         `xml:",omitempty"`
	Bidorbuy                  string         `xml:",omitempty"`
	TaxRate                   string         `xml:",omitempty"`
	TaxinPrice                string         `xml:",omitempty"`
	TaxinBidorbuy             string         `xml:",omitempty"`
	Reserved                  string         `xml:",omitempty"`
	IsBidCreditRestrictions   string         `xml:",omitempty"`
	IsBidderRestrictions      string         `xml:",omitempty"`
	IsBidderRatioRestrictions string         `xml:",omitempty"`
	IsEarlyClosing            string         `xml:",omitempty"`
	IsAutomaticExtension      string         `xml:",omitempty"`
	IsOffer                   string         `xml:",omitempty"`
	HasOfferAccept            string         `xml:",omitempty"`
	IsCharity                 string         `xml:",omitempty"`
	SalesContract             string         `xml:",omitempty"`
	IsFleaMarket              string         `xml:",omitempty"`
	Option                    Option         `xml:",omitempty"`
	Description               string         `xml:",omitempty"`
	SeoKeywords               string         `xml:",omitempty"`
	Payment                   AuctionPayment `xml:",omitempty"`
	BlindBusiness             string         `xml:",omitempty"`
	SevenElevenReceive        string         `xml:",omitempty"`
	ChargeForShipping         string         `xml:",omitempty"`
	Location                  string         `xml:",omitempty"`
	IsWorldwide               string         `xml:",omitempty"`
	ShipTime                  string         `xml:",omitempty"`
	ShippingInput             string         `xml:",omitempty"`
	Shipping                  Shipping       `xml:",omitempty"`
	BaggageInfo               BaggageInfo    `xml:",omitempty"`
	IsAdult                   string         `xml:",omitempty"`
	IsCreature                string         `xml:",omitempty"`
	IsSpecificCategory        string         `xml:",omitempty"`
	IsCharityCategory         string         `xml:",omitempty"`
	CharityOption             CharityOption  `xml:",omitempty"`
	AnsweredQAndANum          string         `xml:",omitempty"`
	Status                    string         `xml:",omitempty"`
}

// AuctionItemResultSet resultSet information
type AuctionItemResultSet struct {
	FirstResultPosition   string            `xml:"firstResultPosition,attr,omitempty"`
	TotalResultsAvailable string            `xml:"totalResultsAvailable,attr,omitempty"`
	TotalResultsReturned  string            `xml:"totalResultsReturned,attr,omitempty"`
	Result                AuctionItemResult `xml:",omitempty"`
}

// CategoryTreeParam parameters for CategoryTree
type CategoryTreeParam struct {
	Output   string `json:"output,omitempty"`
	Callback string `json:"callback,omitempty"`
	Category string `json:"category,omitempty"`
	Adf      string `json:"adf,omitempty"`
}

// CategoryLeafParam parameters for CategoryLeaf
type CategoryLeafParam struct {
	Output              string `json:"output,omitempty"`
	Callback            string `json:"callback,omitempty"`
	Category            string `json:"category,omitempty"`
	Page                string `json:"page,omitempty"`
	Sort                string `json:"sort,omitempty"`
	Order               string `json:"order,omitempty"`
	Store               string `json:"store,omitempty"`
	Aucminprice         string `json:"aucminprice,omitempty"`
	Aucmaxprice         string `json:"aucmaxprice,omitempty"`
	AucminBidorbuyPrice string `json:"aucmin_bidorbuy_price,omitempty"`
	AucmaxBidorbuyPrice string `json:"aucmax_bidorbuy_price,omitempty"`
	Easypayment         string `json:"easypayment,omitempty"`
	New                 string `json:"new,omitempty"`
	Freeshipping        string `json:"freeshipping,omitempty"`
	Wrappingicon        string `json:"wrappingicon,omitempty"`
	Buynow              string `json:"buynow,omitempty"`
	Thumbnail           string `json:"thumbnail,omitempty"`
	Attn                string `json:"attn,omitempty"`
	Point               string `json:"point,omitempty"`
	GiftIcon            string `json:"gift_icon,omitempty"`
	ItemStatus          string `json:"item_status,omitempty"`
	Offer               string `json:"offer,omitempty"`
	MinCharity          string `json:"min_charity,omitempty"`
	MaxCharity          string `json:"max_charity,omitempty"`
	MinAffiliate        string `json:"min_affiliate,omitempty"`
	MaxAffiliate        string `json:"max_affiliate,omitempty"`
	Timebuf             string `json:"timebuf,omitempty"`
	Ranking             string `json:"ranking,omitempty"`
	Seller              string `json:"seller,omitempty"`
}

// SellingListParam parameters for SellingList
type SellingListParam struct {
	Output              string `json:"output,omitempty"`
	Callback            string `json:"callback,omitempty"`
	SellerID            string `json:"sellerID,omitempty"`
	Page                string `json:"page,omitempty"`
	Sort                string `json:"sort,omitempty"`
	Order               string `json:"order,omitempty"`
	Store               string `json:"store,omitempty"`
	Aucminprice         string `json:"aucminprice,omitempty"`
	Aucmaxprice         string `json:"aucmaxprice,omitempty"`
	AucminBidorbuyPrice string `json:"aucmin_bidorbuy_price,omitempty"`
	AucmaxBidorbuyPrice string `json:"aucmax_bidorbuy_price,omitempty"`
	Easypayment         string `json:"easypayment,omitempty"`
	New                 string `json:"new,omitempty"`
	Freeshipping        string `json:"freeshipping,omitempty"`
	Wrappingicon        string `json:"wrappingicon,omitempty"`
	Buynow              string `json:"buynow,omitempty"`
	Thumbnail           string `json:"thumbnail,omitempty"`
	Attn                string `json:"attn,omitempty"`
	Point               string `json:"point,omitempty"`
	GiftIcon            string `json:"gift_icon,omitempty"`
	ItemStatus          string `json:"item_status,omitempty"`
	Offer               string `json:"offer,omitempty"`
}

// SearchParam parameters for Search
type SearchParam struct {
	Output              string `json:"output,omitempty"`
	Callback            string `json:"callback,omitempty"`
	Query               string `json:"query,omitempty"`
	Type                string `json:"type,omitempty"`
	Category            string `json:"category,omitempty"`
	Page                string `json:"page,omitempty"`
	Sort                string `json:"sort,omitempty"`
	Order               string `json:"order,omitempty"`
	Store               string `json:"store,omitempty"`
	Aucminprice         string `json:"aucminprice,omitempty"`
	Aucmaxprice         string `json:"aucmaxprice,omitempty"`
	AucminBidorbuyPrice string `json:"aucmin_bidorbuy_price,omitempty"`
	AucmaxBidorbuyPrice string `json:"aucmax_bidorbuy_price,omitempty"`
	LocCd               string `json:"loc_cd,omitempty"`
	Easypayment         string `json:"easypayment,omitempty"`
	New                 string `json:"new,omitempty"`
	Freeshipping        string `json:"freeshipping,omitempty"`
	Wrappingicon        string `json:"wrappingicon,omitempty"`
	Buynow              string `json:"buynow,omitempty"`
	Thumbnail           string `json:"thumbnail,omitempty"`
	Attn                string `json:"attn,omitempty"`
	Point               string `json:"point,omitempty"`
	GiftIcon            string `json:"gift_icon,omitempty"`
	ItemStatus          string `json:"item_status,omitempty"`
	Offer               string `json:"offer,omitempty"`
	Adf                 string `json:"adf,omitempty"`
	MinCharity          string `json:"min_charity,omitempty"`
	MaxCharity          string `json:"max_charity,omitempty"`
	MinAffiliate        string `json:"min_affiliate,omitempty"`
	MaxAffiliate        string `json:"max_affiliate,omitempty"`
	Timebuf             string `json:"timebuf,omitempty"`
	Ranking             string `json:"ranking,omitempty"`
	Seller              string `json:"seller,omitempty"`
	F                   string `json:"f,omitempty"`
}

// AuctionItemParam parameters for AuctionItem
type AuctionItemParam struct {
	Output    string `json:"output,omitempty"`
	Callback  string `json:"callback,omitempty"`
	AuctionID string `json:"auctionID,omitempty"`
}

// AuctionItem Auction Item
func (c *Client) AuctionItem(param *AuctionItemParam) (*AuctionItemResultSet, error) {
	v := url.Values{}
	if p := param.Output; p != "" {
		v.Set("output", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}
	if p := param.AuctionID; p != "" {
		v.Set("auctionID", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/AuctionWebService/V2/auctionItem").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data AuctionItemResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// Search Search
func (c *Client) Search(param *SearchParam) (*SearchResultSet, error) {
	v := url.Values{}
	if p := param.Output; p != "" {
		v.Set("output", p)
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
	if p := param.Category; p != "" {
		v.Set("category", p)
	}
	if p := param.Page; p != "" {
		v.Set("page", p)
	}
	if p := param.Sort; p != "" {
		v.Set("sort", p)
	}
	if p := param.Order; p != "" {
		v.Set("order", p)
	}
	if p := param.Store; p != "" {
		v.Set("store", p)
	}
	if p := param.Aucminprice; p != "" {
		v.Set("aucminprice", p)
	}
	if p := param.Aucmaxprice; p != "" {
		v.Set("aucmaxprice", p)
	}
	if p := param.AucminBidorbuyPrice; p != "" {
		v.Set("aucmin_bidorbuy_price", p)
	}
	if p := param.AucmaxBidorbuyPrice; p != "" {
		v.Set("aucmax_bidorbuy_price", p)
	}
	if p := param.LocCd; p != "" {
		v.Set("loc_cd", p)
	}
	if p := param.Easypayment; p != "" {
		v.Set("easypayment", p)
	}
	if p := param.New; p != "" {
		v.Set("new", p)
	}
	if p := param.Freeshipping; p != "" {
		v.Set("freeshipping", p)
	}
	if p := param.Wrappingicon; p != "" {
		v.Set("wrappingicon", p)
	}
	if p := param.Buynow; p != "" {
		v.Set("buynow", p)
	}
	if p := param.Thumbnail; p != "" {
		v.Set("thumbnail", p)
	}
	if p := param.Attn; p != "" {
		v.Set("attn", p)
	}
	if p := param.Point; p != "" {
		v.Set("point", p)
	}
	if p := param.GiftIcon; p != "" {
		v.Set("gift_icon", p)
	}
	if p := param.ItemStatus; p != "" {
		v.Set("item_status", p)
	}
	if p := param.Offer; p != "" {
		v.Set("offer", p)
	}
	if p := param.Adf; p != "" {
		v.Set("adf", p)
	}
	if p := param.MinCharity; p != "" {
		v.Set("min_charity", p)
	}
	if p := param.MaxCharity; p != "" {
		v.Set("max_charity", p)
	}
	if p := param.MinAffiliate; p != "" {
		v.Set("min_affiliate", p)
	}
	if p := param.MaxAffiliate; p != "" {
		v.Set("max_affiliate", p)
	}
	if p := param.Timebuf; p != "" {
		v.Set("timebuf", p)
	}
	if p := param.Ranking; p != "" {
		v.Set("ranking", p)
	}
	if p := param.Seller; p != "" {
		v.Set("seller", p)
	}
	if p := param.F; p != "" {
		v.Set("f", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/AuctionWebService/V2/search").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data SearchResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// SellingList Leaf Category
func (c *Client) SellingList(param *SellingListParam) (*SellingListResultSet, error) {
	v := url.Values{}
	if p := param.Output; p != "" {
		v.Set("output", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}
	if p := param.SellerID; p != "" {
		v.Set("sellerID", p)
	}
	if p := param.Page; p != "" {
		v.Set("page", p)
	}
	if p := param.Sort; p != "" {
		v.Set("sort", p)
	}
	if p := param.Order; p != "" {
		v.Set("order", p)
	}
	if p := param.Store; p != "" {
		v.Set("store", p)
	}
	if p := param.Aucminprice; p != "" {
		v.Set("aucminprice", p)
	}
	if p := param.Aucmaxprice; p != "" {
		v.Set("aucmaxprice", p)
	}
	if p := param.AucminBidorbuyPrice; p != "" {
		v.Set("aucmin_bidorbuy_price", p)
	}
	if p := param.AucmaxBidorbuyPrice; p != "" {
		v.Set("aucmax_bidorbuy_price", p)
	}
	if p := param.Easypayment; p != "" {
		v.Set("easypayment", p)
	}
	if p := param.New; p != "" {
		v.Set("new", p)
	}
	if p := param.Freeshipping; p != "" {
		v.Set("freeshipping", p)
	}
	if p := param.Wrappingicon; p != "" {
		v.Set("wrappingicon", p)
	}
	if p := param.Buynow; p != "" {
		v.Set("buynow", p)
	}
	if p := param.Thumbnail; p != "" {
		v.Set("thumbnail", p)
	}
	if p := param.Attn; p != "" {
		v.Set("attn", p)
	}
	if p := param.Point; p != "" {
		v.Set("point", p)
	}
	if p := param.GiftIcon; p != "" {
		v.Set("gift_icon", p)
	}
	if p := param.ItemStatus; p != "" {
		v.Set("item_status", p)
	}
	if p := param.Offer; p != "" {
		v.Set("offer", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/AuctionWebService/V2/sellingList").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data SellingListResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// CategoryLeaf Leaf Category
func (c *Client) CategoryLeaf(param *CategoryLeafParam) (*CategoryLeafResultSet, error) {
	v := url.Values{}
	if p := param.Output; p != "" {
		v.Set("output", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}
	if p := param.Category; p != "" {
		v.Set("category", p)
	}
	if p := param.Page; p != "" {
		v.Set("page", p)
	}
	if p := param.Sort; p != "" {
		v.Set("sort", p)
	}
	if p := param.Order; p != "" {
		v.Set("order", p)
	}
	if p := param.Store; p != "" {
		v.Set("store", p)
	}
	if p := param.Aucminprice; p != "" {
		v.Set("aucminprice", p)
	}
	if p := param.Aucmaxprice; p != "" {
		v.Set("aucmaxprice", p)
	}
	if p := param.AucminBidorbuyPrice; p != "" {
		v.Set("aucmin_bidorbuy_price", p)
	}
	if p := param.AucmaxBidorbuyPrice; p != "" {
		v.Set("aucmax_bidorbuy_price", p)
	}
	if p := param.Easypayment; p != "" {
		v.Set("easypayment", p)
	}
	if p := param.New; p != "" {
		v.Set("new", p)
	}
	if p := param.Freeshipping; p != "" {
		v.Set("freeshipping", p)
	}
	if p := param.Wrappingicon; p != "" {
		v.Set("wrappingicon", p)
	}
	if p := param.Buynow; p != "" {
		v.Set("buynow", p)
	}
	if p := param.Thumbnail; p != "" {
		v.Set("thumbnail", p)
	}
	if p := param.Attn; p != "" {
		v.Set("attn", p)
	}
	if p := param.Point; p != "" {
		v.Set("point", p)
	}
	if p := param.GiftIcon; p != "" {
		v.Set("gift_icon", p)
	}
	if p := param.ItemStatus; p != "" {
		v.Set("item_status", p)
	}
	if p := param.Offer; p != "" {
		v.Set("offer", p)
	}
	if p := param.MinCharity; p != "" {
		v.Set("min_charity", p)
	}
	if p := param.MaxCharity; p != "" {
		v.Set("max_charity", p)
	}
	if p := param.MinAffiliate; p != "" {
		v.Set("min_affiliate", p)
	}
	if p := param.MaxAffiliate; p != "" {
		v.Set("max_affiliate", p)
	}
	if p := param.Timebuf; p != "" {
		v.Set("timebuf", p)
	}
	if p := param.Ranking; p != "" {
		v.Set("ranking", p)
	}
	if p := param.Seller; p != "" {
		v.Set("seller", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/AuctionWebService/V2/categoryLeaf").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data CategoryLeafResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}

// CategoryTree Tree Category
func (c *Client) CategoryTree(param *CategoryTreeParam) (*CategoryTreeResultSet, error) {
	v := url.Values{}
	if p := param.Output; p != "" {
		v.Set("output", p)
	}
	if p := param.Callback; p != "" {
		v.Set("callback", p)
	}
	if p := param.Category; p != "" {
		v.Set("category", p)
	}
	if p := param.Adf; p != "" {
		v.Set("adf", p)
	}

	url := fmt.Sprintf("%s?%s", c.urlFor("/AuctionWebService/V2/categoryTree").String(), v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	var data CategoryTreeResultSet
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, err
}
