package helper

import (
	"fmt"
	"os"

	"github.com/imrenagi/go-payment/invoice"
	"github.com/xendit/xendit-go/ewallet"
)

// NewDana is factory for Dana payment with xendit latest charge API
func NewDana(inv *invoice.Invoice) (*ewallet.CreateEWalletChargeParams, error) {

	successRedirectURL := os.Getenv("DANA_SUCCESS_REDIRECT_URL")
	if inv.SuccessRedirectURL != "" {
		successRedirectURL = inv.SuccessRedirectURL
	}

	props := map[string]string{
		"success_redirect_url": successRedirectURL,
	}

	return newBuilder(inv).
		SetPaymentMethod(EWalletIDDana).
		SetChannelProperties(props).
		Build()
}

// NewLinkAja is factory for LinkAja payment with xendit latest charge API
func NewLinkAja(inv *invoice.Invoice) (*ewallet.CreateEWalletChargeParams, error) {

	successRedirectURL := os.Getenv("LINKAJA_SUCCESS_REDIRECT_URL")
	if inv.SuccessRedirectURL != "" {
		successRedirectURL = inv.SuccessRedirectURL
	}

	props := map[string]string{
		"success_redirect_url": successRedirectURL,
	}

	return newBuilder(inv).
		SetPaymentMethod(EWalletIDLinkAja).
		SetChannelProperties(props).
		Build()
}

// NewOVO is factory for OVO payment with xendit latest charge API
func NewOVO(inv *invoice.Invoice) (*ewallet.CreateEWalletChargeParams, error) {

	if inv.BillingAddress == nil {
		return nil, fmt.Errorf("customer phone number is required")
	}

	if !OvoChargePhoneValidator.IsValid(inv.BillingAddress.PhoneNumber) {
		return nil, fmt.Errorf("invalid phone format. must be in +628xxxxxx format")
	}

	props := map[string]string{
		"mobile_number": inv.BillingAddress.PhoneNumber,
	}

	return newBuilder(inv).
		SetPaymentMethod(EWalletIDOVO).
		SetChannelProperties(props).
		Build()
}
