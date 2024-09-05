package lister_test

import (
	"testing"

	"github.com/gomig/lister"
)

func TestRecordResolver(t *testing.T) {
	lister1 := lister.New()
	lister1.SetSorts("_id", "title")
	if err := lister.RecordResolver(lister1, lister.ListerRequest{
		Page:   10,
		Limit:  100,
		Sort:   "title",
		Order:  "desc",
		Search: "John",
		Filters: map[string]any{
			"username": "JackMa",
		},
	}); err != nil {
		t.Fatal(err.Error())
	} else if lister1.Page() != 10 ||
		lister1.Limit() != 100 ||
		lister1.Sort() != "title" ||
		lister1.Order() != "desc" ||
		lister1.Search() != "John" ||
		lister1.CastFilter("username").StringSafe("ss") != "JackMa" {
		t.Fail()
	}

	lister2 := lister.New()
	lister2.SetSorts("id", "clients.name")
	if err := lister.JsonMapperResolver(
		lister2,
		`{"sort": "client"}`,
		map[string]string{"client": "name"},
		map[string][]string{"clients": {"name"}},
	); err != nil {
		t.Fatal(err.Error())
	} else if lister2.Sort() != "clients.name" {
		t.Fail()
	}
}
