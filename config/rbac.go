// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package config

import (
	"sort"
	"strings"
)

// CreateRbacConfig return *[]KeyValuePair
func CreateRbacConfig() *RbacConfig {

	cfg := &RbacConfig{
		{
			Key:   "right.all",
			Value: 1,
		},
		{
			Key:   "right.edit",
			Value: 2,
		},
	}

	sort.Sort(cfg)

	return cfg
}

// RbacConfig struct
type RbacConfig []KeyValuePair

// KeyValuePair struct
type KeyValuePair struct {
	Key   string
	Value int64
}

// Len return len
func (o *RbacConfig) Len() int { return len(*o) }

// Swap swap i, j
func (o *RbacConfig) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *RbacConfig) Less(i, j int) bool { return (*o)[i].Key < (*o)[j].Key }

// Search use binary search to find and return the smallest index Value
func (o *RbacConfig) Search(key string) int64 {

	i := sort.Search(o.Len(), func(i int) bool { return (*o)[i].Key >= key })

	if i < o.Len() && (*o)[i].Key == key {
		return (*o)[i].Value
	}

	return 0
}

// Sum sum right.value
func (o *RbacConfig) Sum() int64 {
	var val int64 = 0
	for i := 0; i < o.Len(); i++ {
		val += (*o)[i].Value
	}
	return val
}

// Keys get keys by userRight
func (o *RbacConfig) Keys(userRight int64) string {
	var sb strings.Builder
	for i := 0; i < o.Len(); i++ {
		if (*o)[i].Value&userRight > 0 {
			if sb.Len() > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString((*o)[i].Key)
		}
	}
	return sb.String()
}
