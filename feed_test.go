package hatena

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_HotEntryAll(t *testing.T) {
	tests := []struct {
		name    string
		want    *Entries
		wantErr bool
	}{
		{
			name: "TestClient_HotEntryAll",
			want: &Entries{
				Title:       "はてなブックマーク - 人気エントリー",
				Link:        "http://b.hatena.ne.jp/hotentry",
				Description: "最近の人気エントリー",
				Entry: []Entry{
					{
						Title:         "TopNews1",
						Link:          "https://example.com/news1/",
						Description:   "TopNews1 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
					{
						Title:         "TopNews2",
						Link:          "https://example.com/news2/",
						Description:   "TopNews2 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientFile(http.StatusOK, "test_data/hotentry_all.txt")
			got, err := c.HotEntryAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.HotEntryAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.HotEntryAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_HotEntry(t *testing.T) {
	type fields struct {
		http *http.Client
	}
	type args struct {
		category string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Entries
		wantErr bool
	}{
		{
			name: "TestClient_HotEntry",
			want: &Entries{
				Title:       "はてなブックマーク - 人気エントリー - テクノロジー",
				Link:        "http://b.hatena.ne.jp/hotentry/it",
				Description: "最近の人気エントリー - テクノロジー",
				Entry: []Entry{
					{
						Title:         "TopNews1",
						Link:          "https://example.com/news1/",
						Description:   "TopNews1 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
					{
						Title:         "TopNews2",
						Link:          "https://example.com/news2/",
						Description:   "TopNews2 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientFile(http.StatusOK, "test_data/hotentry.txt")
			got, err := c.HotEntry(tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.HotEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.HotEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewEntryAll(t *testing.T) {
	type fields struct {
		http *http.Client
	}
	type args struct {
		options map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Entries
		wantErr bool
	}{
		{
			name: "TestClient_NewEntryAll",
			want: &Entries{
				Title:       "はてなブックマーク - 新着エントリー",
				Link:        "http://b.hatena.ne.jp/entrylist",
				Description: "新着エントリー",
				Entry: []Entry{
					{
						Title:         "TopNews1",
						Link:          "https://example.com/news1/",
						Description:   "TopNews1 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
					{
						Title:         "TopNews2",
						Link:          "https://example.com/news2/",
						Description:   "TopNews2 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientFile(http.StatusOK, "test_data/newentry_all.txt")
			got, err := c.NewEntryAll(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.NewEntryAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewEntryAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewEntry(t *testing.T) {
	type fields struct {
		http *http.Client
	}
	type args struct {
		category string
		options  map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Entries
		wantErr bool
	}{
		{
			name: "TestClient_NewEntry",
			want: &Entries{
				Title:       "はてなブックマーク - 新着エントリー - テクノロジー",
				Link:        "http://b.hatena.ne.jp/entrylist/it",
				Description: "新着エントリー - テクノロジー",
				Entry: []Entry{
					{
						Title:         "TopNews1",
						Link:          "https://example.com/news1/",
						Description:   "TopNews1 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
					{
						Title:         "TopNews2",
						Link:          "https://example.com/news2/",
						Description:   "TopNews2 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientFile(http.StatusOK, "test_data/newEntry.txt")
			got, err := c.NewEntry(tt.args.category, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.NewEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SearchEntry(t *testing.T) {
	type args struct {
		searchWord string
		searchType string
		options    map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *Entries
		wantErr bool
	}{
		{
			name: "TestClient_SearchEntry",
			args: args{
				searchWord: "hatena",
				searchType: "text",
				options:    map[string]string{"sort": "recent", "date_begin": "2017-06-01"},
			},
			want: &Entries{
				Title:       "本文「hatena」を検索 - はてなブックマーク",
				Link:        "http://b.hatena.ne.jp/search/text?safe=on&sort=recent&q=hatena&date_begin=2017-06-10",
				Description: "本文「hatena」を検索 - はてなブックマーク",
				Entry: []Entry{
					{
						Title:         "TopNews1",
						Link:          "https://example.com/news1/",
						Description:   "TopNews1 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
					{
						Title:         "TopNews2",
						Link:          "https://example.com/news2/",
						Description:   "TopNews2 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientFile(http.StatusOK, "test_data/searchEntry.txt")
			got, err := c.SearchEntry(tt.args.searchWord, tt.args.searchType, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SearchEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SearchEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SearchUrlEntry(t *testing.T) {
	type fields struct {
		http *http.Client
	}
	type args struct {
		searchUrl string
		options   map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Entries
		wantErr bool
	}{
		{
			name: "TestClient_SearchUrlEntry",
			args: args{
				searchUrl: "https://yahoo.co.jp",
				options:   map[string]string{"sort": "recent", "date_begin": "2017-06-01"},
			},
			want: &Entries{
				Title:       "はてなブックマーク - 人気エントリー - 『yahoo.co.jp』",
				Link:        "http://b.hatena.ne.jp/entrylist?sort=count&url=https%3A%2F%2Fyahoo.co.jp",
				Description: "『yahoo.co.jp』 の人気エントリー",
				Entry: []Entry{
					{
						Title:         "TopNews1",
						Link:          "https://example.com/news1/",
						Description:   "TopNews1 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
					{
						Title:         "TopNews2",
						Link:          "https://example.com/news2/",
						Description:   "TopNews2 Description",
						Date:          "2017-07-19T02:06:07+09:00",
						Subject:       "テクノロジー",
						Bookmarkcount: 10,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testClientFile(http.StatusOK, "test_data/searchUrlEntry.txt")
			got, err := c.SearchUrlEntry(tt.args.searchUrl, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SearchUrlEntry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SearchUrlEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}
