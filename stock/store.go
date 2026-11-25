package main

import "context"

type store struct {
	// DB connection
}

func NewStore(context.Context) *store {
	return &store{}
}

func (s *store) Create(context.Context) error {
	return nil
}



