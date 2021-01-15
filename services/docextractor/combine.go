// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package docextractor

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/mattermost/mattermost-server/v5/mlog"
)

type combineExtractor struct {
	SubExtractors []Extractor
}

func (ce *combineExtractor) Add(extractor Extractor) {
	ce.SubExtractors = append(ce.SubExtractors, extractor)
}

func (ce *combineExtractor) Match(filename string) bool {
	for _, extractor := range ce.SubExtractors {
		if extractor.Match(filename) {
			return true
		}
	}
	return false
}

func (ce *combineExtractor) Extract(filename string, r io.Reader) (string, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		mlog.Warn("unable to extract file content", mlog.Err(err))
		return "", nil
	}

	for _, extractor := range ce.SubExtractors {
		if extractor.Match(filename) {
			text, err := extractor.Extract(filename, bytes.NewReader(data))
			if err != nil {
				mlog.Warn("unable to extract file content", mlog.Err(err))
				continue
			}
			return text, nil
		}
	}
	return "", nil
}
