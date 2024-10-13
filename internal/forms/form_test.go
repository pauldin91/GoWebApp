package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	posted := url.Values{}

	form := New(posted)
	isValid := form.Valid()
	if !isValid {
		t.Error("invalid form")
	}
}

func TestForm_Required(t *testing.T) {
	posted := url.Values{}

	form := New(posted)
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("fields should missing")
	}
	posted = url.Values{}
	posted.Add("a", "a")
	posted.Add("b", "a")
	posted.Add("c", "a")
	form = New(posted)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestFrom_Has(t *testing.T) {
	posted := url.Values{}

	form := New(posted)
	has := form.Has("whatever")
	if has {
		t.Error("should have when it does not")
	}
	posted = url.Values{}
	posted.Add("a", "a")
	form = New(posted)
	has = form.Has("a")
	if !has {
		t.Error("Shows form does not have field when it should")
	}
}

func TestForm_MinLength(t *testing.T) {
	posted := url.Values{}

	form := New(posted)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("shows when it should not")
	}
	posted = url.Values{}
	posted.Add("a", "abb")
	form = New(posted)
	form.MinLength("a", 3)
	if !form.Valid() {
		t.Error("Should be valid")
	}
	posted = url.Values{}
	posted.Add("a", "abbqq")
	form = New(posted)
	form.MinLength("a", 6)
	if form.Valid() {
		t.Error("Should be vinvalid")
	}
}

func Test_IsEmail(t *testing.T) {

	posted := url.Values{}
	form := New(posted)
	form.IsEmail("x")
	if form.Valid() {
		t.Error("Should have been valid")
	}
	posted = url.Values{}
	posted.Add("email", "ao@joo.com")
	form = New(posted)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("Should have been valid")
	}
	posted = url.Values{}
	posted.Add("email", "ao.com")
	form = New(posted)
	form.IsEmail("email")
	if form.Valid() {
		t.Error("Should have been invalid")
	}

}
