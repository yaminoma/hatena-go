package hatena

import (
	"net/http"
	"os"
	"testing"
)

func TestEntryInfo(t *testing.T) {

	client := testClientFile(http.StatusOK, "test_data/entry.txt")
	result, err := client.EntryInfo("https://github.com/")

	if err != nil {
		t.Error(err)
	}
	if result.Eid != nil {
		t.Error("")
	}
	if result.Title != nil {
		t.Error("")
	}
	if result.Count != nil {
		t.Error("")
	}
	if result.Url != nil {
		t.Error("")
	}
	if result.EntryUrl != nil {
		t.Error("")
	}
	if result.Screenshot != nil {
		t.Error("")
	}
	if result.Bookmarks[0].User != nil {
		t.Error("")
	}
	if result.Bookmarks[0].Comment != nil {
		t.Error("")
	}
	if result.Bookmarks[0].Timestamp != nil {
		t.Error("")
	}
	if result.Bookmarks[0].Tags != nil {
		t.Error("")
	}
	if result.RelatedEntries[0].Eid != nil {
		t.Error("")
	}
	if result.RelatedEntries[0].Title != nil {
		t.Error("")
	}
	if result.RelatedEntries[0].Count != nil {
		t.Error("")
	}
	if result.RelatedEntries[0].Url != nil {
		t.Error("")
	}
	if result.RelatedEntries[0].EntryUrl != nil {
		t.Error("")
	}
}
