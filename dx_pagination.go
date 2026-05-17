// Page[T] pagination — Go port of _dx_pagination.py.
//
// Same contract as Python/TS: list methods return a Page[T] that exposes
// .Items / .NextCursor / .Total AND can be iterated across all pages with
// the All() method.
//
// Why a method rather than range-over-func: Go 1.23+ has range-over-func,
// but the generated SDK targets Go 1.21+ for compatibility. All() returns
// a function the customer can drive in a for loop, which is idiomatic for
// our Go version target. When the module bumps to Go 1.23+, swap All() for
// the iter.Seq[T] pattern.

package moolabs

import (
	"context"
	"fmt"
)

// FetchNext returns the next Page in the sequence or nil if no more pages.
// Called at most once per Page (single-use).
type FetchNext[T any] func(ctx context.Context) (*Page[T], error)

// Page holds one page of items plus a lazy way to advance through
// subsequent pages.
type Page[T any] struct {
	// Items on THIS page only (never auto-fetched).
	Items []T
	// NextCursor — opaque; nil-equivalent ("") means no more pages.
	NextCursor string
	// Total — total items across all pages if the API returned it, else -1.
	Total int

	fetchNext FetchNext[T]
	fetched   bool
}

// NewPage builds a Page. fetchNext may be nil for the last page.
func NewPage[T any](items []T, nextCursor string, total int, fetchNext FetchNext[T]) *Page[T] {
	itemsCopy := make([]T, len(items))
	copy(itemsCopy, items)
	return &Page[T]{
		Items:      itemsCopy,
		NextCursor: nextCursor,
		Total:      total,
		fetchNext:  fetchNext,
	}
}

// EmptyPage returns a terminal empty page. Useful for "no results" responses.
func EmptyPage[T any]() *Page[T] {
	return &Page[T]{Items: []T{}, NextCursor: "", Total: 0}
}

// Len returns CURRENT page size only (never triggers a fetch).
// Use Page.Total for the across-pages count.
func (p *Page[T]) Len() int {
	return len(p.Items)
}

// All yields items across ALL pages, fetching subsequent pages lazily.
// Returns a yield function the customer drives in a for loop:
//
//	for item, err := page.All(ctx); err == nil && item != nil; item, err = page.All(ctx) {
//	    // ...
//	}
//
// Closures over context for fetch calls. Subsequent fetches may return error;
// items already yielded before the error remain consumed.
//
// For cleaner consumer code, use Each() which calls a per-item callback.
func (p *Page[T]) All(ctx context.Context) func() (*T, error) {
	idx := 0
	current := p
	return func() (*T, error) {
		for {
			if idx < len(current.Items) {
				item := current.Items[idx]
				idx++
				return &item, nil
			}
			// Current page exhausted — advance
			if current.NextCursor == "" || current.fetchNext == nil || current.fetched {
				return nil, nil // end of stream
			}
			current.fetched = true
			next, err := current.fetchNext(ctx)
			current.fetchNext = nil
			if err != nil {
				return nil, fmt.Errorf("moolabs: page fetch failed: %w", err)
			}
			if next == nil {
				return nil, nil
			}
			current = next
			idx = 0
		}
	}
}

// Each invokes fn for each item across all pages. Stops at the first
// error from fn or from a page fetch. Items consumed before the error
// remain consumed.
func (p *Page[T]) Each(ctx context.Context, fn func(T) error) error {
	next := p.All(ctx)
	for {
		item, err := next()
		if err != nil {
			return err
		}
		if item == nil {
			return nil
		}
		if err := fn(*item); err != nil {
			return err
		}
	}
}
