package orm

import (
	"gorm.io/gorm"
)

const (
	// DefaultMaxInSize represents default variables number on IN () in SQL
	DefaultMaxInSize = 50
	DefaultPagingNum = 10
	MaxResponseItems = 100
)

// Paginator is the base for different ListOptions types
type Paginator interface {
	GetSkipTake() (skip, take int)
	GetStartEnd() (start, end int)
	IsListAll() bool
}

// SetSessionPagination sets pagination for a database session
func SetSessionPagination(sess *gorm.DB, p Paginator) *gorm.DB {
	skip, take := p.GetSkipTake()
	return sess.Offset(skip).Limit(take)
}

// ListOptions options to paginate results
type ListOptions struct {
	PageSize int
	Page     int  // start from 1
	ListAll  bool // if true, then PageSize and Page will not be taken
}

var _ Paginator = &ListOptions{}

// GetSkipTake returns the skip and take values
func (opts *ListOptions) GetSkipTake() (skip, take int) {
	opts.SetDefaultValues()
	return (opts.Page - 1) * opts.PageSize, opts.PageSize
}

// GetStartEnd returns the start and end of the ListOptions
func (opts *ListOptions) GetStartEnd() (start, end int) {
	start, take := opts.GetSkipTake()
	end = start + take
	return start, end
}

// IsListAll indicates PageSize and Page will be ignored
func (opts *ListOptions) IsListAll() bool {
	return opts.ListAll
}

// SetDefaultValues sets default values
func (opts *ListOptions) SetDefaultValues() {
	if opts.PageSize <= 0 {
		opts.PageSize = DefaultPagingNum
	}
	opts.PageSize = min(opts.PageSize, MaxResponseItems)
	opts.Page = max(opts.Page, 1)
}
