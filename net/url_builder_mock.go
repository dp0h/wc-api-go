package net

import (
	"github.com/dp0h/wc-api-go/request"
)

// URLBuilderMock ...
type URLBuilderMock struct {
	url         string
	isBasicAuth bool
}

//GetURL ...
func (b *URLBuilderMock) GetURL(req request.Request) string {
	return b.url
}

// IsBasicAuth ...
func (b *URLBuilderMock) IsBasicAuth() bool {
	return b.isBasicAuth
}
