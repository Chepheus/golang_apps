package forms

import (
	"net/url"
	"time"
)

type ReservationForm struct {
	City       string    `field:"city" validation:"not_empty"`
	HomeNumber int       `field:"home_number" validation:"not_empty"`
	DateFrom   time.Time `field:"date_from" validation:"not_empty"`
	DateTo     time.Time `field:"date_to" validation:"not_empty"`
	Name       string    `field:"name" validation:"not_empty"`
	Phone      string    `field:"phone" validation:"not_empty"`
	Email      string    `field:"email" validation:"not_empty"`
	errors     map[string][]string
}

func (f *ReservationForm) setErrors(errors map[string][]string) {
	f.errors = errors
}

func (f *ReservationForm) GetErrors() map[string][]string {
	return f.errors
}

func NewReservationForm(form url.Values) (*ReservationForm, error) {
	f := ReservationForm{}

	rf, err := createForm(&f, form)
	if err != nil {
		return nil, err
	}

	return (*rf).(*ReservationForm), nil
}
