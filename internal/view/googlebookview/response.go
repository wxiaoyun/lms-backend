// nolint:tagliatelle
package googlebookview

type ResponseView struct {
	Kind       string     `json:"kind"`
	TotalItems int        `json:"totalItems"`
	Items      []ItemView `json:"items"`
}

type ItemView struct {
	Kind       string         `json:"kind"`
	ID         string         `json:"id"`
	Etag       string         `json:"etag"`
	SelfLink   string         `json:"selfLink"`
	VolumeInfo VolumeInfoView `json:"volumeInfo"`
	SaleInfo   SaleInfoView   `json:"saleInfo"`
	AccessInfo AccessInfoView `json:"accessInfo"`
	SearchInfo SearchInfoView `json:"searchInfo"`
}

type VolumeInfoView struct {
	Title               string                   `json:"title"`
	Authors             []string                 `json:"authors"`
	Publisher           string                   `json:"publisher"`
	PublishedDate       string                   `json:"publishedDate"`
	Description         string                   `json:"description"`
	IndustryIdentifiers []IndustryIdentifierView `json:"industryIdentifiers"`
	ReadingModes        ReadingModesView         `json:"readingModes"`
	PageCount           int                      `json:"pageCount"`
	PrintType           string                   `json:"printType"`
	Categories          []string                 `json:"categories"`
	MaturityRating      string                   `json:"maturityRating"`
	AllowAnonLogging    bool                     `json:"allowAnonLogging"`
	ContentVersion      string                   `json:"contentVersion"`
	PanelizationSummary PanelizationSummaryView  `json:"panelizationSummary"`
	ImageLinks          ImageLinksView           `json:"imageLinks"`
	Language            string                   `json:"language"`
	PreviewLink         string                   `json:"previewLink"`
	InfoLink            string                   `json:"infoLink"`
	CanonicalVolumeLink string                   `json:"canonicalVolumeLink"`
}

type IndustryIdentifierView struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type ReadingModesView struct {
	Text  bool `json:"text"`
	Image bool `json:"image"`
}

type PanelizationSummaryView struct {
	ContainsEpubBubbles  bool `json:"containsEpubBubbles"`
	ContainsImageBubbles bool `json:"containsImageBubbles"`
}

type ImageLinksView struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}

type SaleInfoView struct {
	Country     string `json:"country"`
	Saleability string `json:"saleability"`
	IsEbook     bool   `json:"isEbook"`
}

type AccessInfoView struct {
	Country                string   `json:"country"`
	Viewability            string   `json:"viewability"`
	Embeddable             bool     `json:"embeddable"`
	PublicDomain           bool     `json:"publicDomain"`
	TextToSpeechPermission string   `json:"textToSpeechPermission"`
	Epub                   EpubView `json:"epub"`
	Pdf                    PdfView  `json:"pdf"`
	WebReaderLink          string   `json:"webReaderLink"`
	AccessViewStatus       string   `json:"accessViewStatus"`
	QuoteSharingAllowed    bool     `json:"quoteSharingAllowed"`
}

type EpubView struct {
	IsAvailable bool `json:"isAvailable"`
}

type PdfView struct {
	IsAvailable bool `json:"isAvailable"`
}

type SearchInfoView struct {
	TextSnippet string `json:"textSnippet"`
}
