package hatena

const (
	entryURL = "http://b.hatena.ne.jp/entry/json/"
)

type EntryInformation struct {
	Eid            int            `json:"eid"`
	Title          string         `json:"title"`
	Count          int            `json:"count"`
	Url            string         `json:"url"`
	EntryUrl       string         `json:"entry_url"`
	Screenshot     string         `json:"screenshot"`
	Bookmarks      []BookmarkUser `json:"bookmarks"`
	RelatedEntries []RelatedEntry `json:"related"`
}

type BookmarkUser struct {
	User      string   `json:"user"`
	Comment   string   `json:"comment"`
	Timestamp string   `json:"timestamp"`
	Tags      []string `json:"tags"`
}

type RelatedEntry struct {
	Eid      int    `json:"eid"`
	Title    string `json:"title"`
	Count    int    `json:"count"`
	Url      string `json:"url"`
	EntryUrl string `json:"entry_url"`
}

// はてなブックマークエントリー情報取得API
func EntryInfo(urlStr string) (*EntryInformation, error) {
	return DefaultClient.EntryInfo(urlStr)
}

func (c *Client) EntryInfo(urlStr string) (*EntryInformation, error) {

	req := entryURL + "?url=" + urlStr

	e := &EntryInformation{}
	err := c.get(req, e, "json")

	return e, err
}
