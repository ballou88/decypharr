package link

import (
	"errors"
	"testing"

	"github.com/sirrobot01/decypharr/pkg/storage"
)

// A refresher that finds nothing returns (nil, nil). getPlacementFile must
// report a refetchable error, not dereference nil and crash the process.
func TestGetPlacementFileRefresherReturnsNil(t *testing.T) {
	s := &Service{entryRefresher: func(string) (*storage.Entry, error) { return nil, nil }}
	entry := &storage.Entry{
		InfoHash:       "abc",
		ActiveProvider: "torbox",
		Files:          map[string]*storage.File{"book.epub": {}},
		Providers: map[string]*storage.ProviderEntry{
			"torbox": {Files: map[string]*storage.ProviderFile{"book.epub": {}}},
		},
	}
	_, err := s.getPlacementFile(entry, "book.epub")
	var linkErr *Error
	if !errors.As(err, &linkErr) || !linkErr.ShouldRefetch() {
		t.Fatalf("want refetchable *Error, got %v", err)
	}
}
