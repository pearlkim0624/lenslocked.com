package controllers

import (
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

func parseURLParams(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	return parseValue(r.Form, dst)
}

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	return parseValue(r.PostForm, dst)
}

func parseValue(values url.Values, dst interface{}) error {
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	if err := dec.Decode(dst, values); err != nil {
		return err
	}
	return nil
}
