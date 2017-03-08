package hatena

import (
	"net/http"
	"os"
	"testing"
)

func TestEntryInfo(t *testing.T) {

	res, err := hatena.EntryInfo("https://github.com/")

	if result.Eid != 10975646 {
		t.Error("Eid is invalid")
	}
	if result.Title != "GitHub" {
		t.Error("Title is invalid")
	}
	if result.Count > 974 {
		t.Error("Count is invalid")
	}
	if result.Url != "https://github.com/" {
		t.Error("")
	}
	if result.EntryUrl != "http://b.hatena.ne.jp/entry/s/github.com/" {
		t.Error("")
	}
	screenshot := "http://screenshot.hatena.ne.jp/images/200x150/f/d/e/b/0/3ba121c130cd7312d649e5f4fb308a2394c.jpg"
	if result.Screenshot != screenshot {
		t.Error("")
	}
	if result.Bookmarks[0].User == nil {
		t.Error("")
	}
	if result.Bookmarks[0].Comment == nil {
		t.Error("")
	}
	if result.Bookmarks[0].Timestamp == nil {
		t.Error("")
	}
	if result.Bookmarks[0].Tags == nil {
		t.Error("")
	}
	if result.RelatedEntries[0].Eid == nil {
		t.Error("")
	}
	if result.RelatedEntries[0].Title == nil {
		t.Error("")
	}
	if result.RelatedEntries[0].Count == nil {
		t.Error("")
	}
	if result.RelatedEntries[0].Url == nil {
		t.Error("")
	}
	if result.RelatedEntries[0].EntryUrl == nil {
		t.Error("")
	}
	if err != nil {
		t.Error(err)
	}
}

var getCategoryPlaylists = `
{
	related: [
		{
			count: 123,
			url: "http://www.infoq.com/jp/articles/9_Fallacies_Java_Performance#.UnBhsgQ0rDA.facebook",
			eid: 167393322,
			title: "Javaのパフォーマンスについての９つの誤信",
			entry_url: "http://b.hatena.ne.jp/entry/www.infoq.com/jp/articles/9_Fallacies_Java_Performance%23.UnBhsgQ0rDA.facebook"
		},
		{
			count: 76,
			url: "http://www.lifehacker.jp/2015/04/150408entrepreneur_must_outsource.html",
			eid: 246661713,
			title: "時間を有効に使うためにアウトソースすべき11のこと ｜ ライフハッカー［日本版］",
			entry_url: "http://b.hatena.ne.jp/entry/www.lifehacker.jp/2015/04/150408entrepreneur_must_outsource.html"
		}
	],
	count: 974,
	bookmarks: [
		{
			comment: "",
			timestamp: "2017/02/18 22:38:32",
			user: "pg4self",
			tags: [
				"github"
			]
		},
		{
			comment: "",
			timestamp: "2008/11/06 02:02:05",
			user: "d-_-b",
			tags: [
				"versioncontrol",
				"web",
				"tool"
			]
		}
	],
	url: "https://github.com/",
	eid: 10975646,
	title: "GitHub",
	screenshot: "http://screenshot.hatena.ne.jp/images/200x150/f/d/e/b/0/3ba121c130cd7312d649e5f4fb308a2394c.jpg",
	entry_url: "http://b.hatena.ne.jp/entry/s/github.com/"
}`
