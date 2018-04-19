package webhooks

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/portofinolabs/recurly"
)

// Webhook notification constants.
const (
	// Account notifications.
	NewAccount         = "new_account_notification"
	UpdatedAccount     = "updated_account_notification"
	BillingInfoUpdated = "billing_info_updated_notification"
	ReactivedAccount   = "reactivated_account_notification"

	// Subscription notifications.
	NewSubscription         = "new_subscription_notification"
	UpdatedSubscription     = "updated_subscription_notification"
	RenewedSubscription     = "renewed_subscription_notification"
	ExpiredSubscription     = "expired_subscription_notification"
	CanceledSubscription    = "canceled_subscription_notification"
	ReactivatedSubscription = "reactivated_subcription_notification"

	// Invoice notifications.
	NewInvoice        = "new_invoice_notification"
	PastDueInvoice    = "past_due_invoice_notification"
	ClosedInvoice     = "closed_invoice_notification"
	ProcessingInvoice = "processing_invoice_notification"

	// Shipping Address notifications.
	NewShippingAddress     = "new_shipping_address_notification"
	DeletedShippingAddress = "deleted_shipping_address_notification"
	UpdatedShippingAddress = "updated_shipping_address_notification"

	// Payment notifications.
	SuccessfulPayment = "successful_payment_notification"
	FailedPayment     = "failed_payment_notification"
	VoidPayment       = "void_payment_notification"
	SuccessfulRefund  = "successful_refund_notification"

	// Dunning Event notifications.
	NewDunningEvent = "new_dunning_event_notification"
)

type notificationName struct {
	XMLName xml.Name
}

// Account represents the account object sent in webhooks.
type Account struct {
	XMLName     xml.Name `xml:"account" json:"-"`
	Code        string   `xml:"account_code,omitempty" json:"account_code"`
	Username    string   `xml:"username,omitempty" json:"username"`
	Email       string   `xml:"email,omitempty" json:"email"`
	FirstName   string   `xml:"first_name,omitempty" json:"first_name"`
	LastName    string   `xml:"last_name,omitempty" json:"last_name"`
	CompanyName string   `xml:"company_name,omitempty" json:"company"`
	Phone       string   `xml:"phone,omitempty" json:"phone"`
}

// Transaction represents the transaction object sent in webhooks.
type Transaction struct {
	XMLName           xml.Name         `xml:"transaction" json:"-"`
	UUID              string           `xml:"id,omitempty" json:"uuid"`
	InvoiceNumber     int              `xml:"invoice_number,omitempty" json:"invoice_number"`
	SubscriptionUUID  string           `xml:"subscription_id,omitempty" json:"subscription_uuid,omitempty"`
	Action            string           `xml:"action,omitempty" json:"action"`
	AmountInCents     int              `xml:"amount_in_cents,omitempty" json:"amount_in_cents"`
	Status            string           `xml:"status,omitempty" json:"status"`
	Message           string           `xml:"message,omitempty" json:"message"`
	GatewayErrorCodes string           `xml:"gateway_error_codes,omitempty" json:"gateway_error_codes"`
	FailureType       string           `xml:"failure_type,omitempty" json:"failure_type"`
	Reference         string           `xml:"reference,omitempty" json:"reference"`
	Source            string           `xml:"source,omitempty" json:"source"`
	Test              recurly.NullBool `xml:"test,omitempty" json:"test"`
	Voidable          recurly.NullBool `xml:"voidable,omitempty" json:"voidable"`
	Refundable        recurly.NullBool `xml:"refundable,omitempty" json:"refundable"`
}

// Invoice represents the invoice object sent in webhooks.
type Invoice struct {
	XMLName             xml.Name         `xml:"invoice,omitempty" json:"-"`
	SubscriptionUUID    string           `xml:"subscription_id,omitempty" json:"subscription_uuid,omitempty"`
	UUID                string           `xml:"uuid,omitempty" json:"uuid"`
	State               string           `xml:"state,omitempty" json:"state"`
	InvoiceNumberPrefix string           `xml:"invoice_number_prefix,omitempty" json:"invoice_number_prefix"`
	InvoiceNumber       int              `xml:"invoice_number,omitempty" json:"invoice_number"`
	PONumber            string           `xml:"po_number,omitempty" json:"vat_number"`
	VATNumber           string           `xml:"vat_number,omitempty" json:"vat_number"`
	TotalInCents        int              `xml:"total_in_cents,omitempty" json:"total_in_cents"`
	Currency            string           `xml:"currency,omitempty" json:"currency"`
	CreatedAt           recurly.NullTime `xml:"date,omitempty" json:"created_at"`
	ClosedAt            recurly.NullTime `xml:"closed_at,omitempty" json:"closed_at"`
	NetTerms            recurly.NullInt  `xml:"net_terms,omitempty" json:"net_terms"`
	CollectionMethod    string           `xml:"collection_method,omitempty" json:"collection_method"`
}

// ShippingAdddress represents the shipping address object sent in webhooks.
type ShippingAdddress struct {
	XMLName     xml.Name `xml:"shipping_address,omitempty" json:"-"`
	ID          int      `xml:"id,omitempty" json:"id"`
	Nickname    string   `xml:"nickname,omitempty" json:"nickname"`
	FirstName   string   `xml:"first_name,omitempty" json:"first_name"`
	LastName    string   `xml:"last_name,omitempty" json:"last_name"`
	CompanyName string   `xml:"company_name,omitempty" json:"company"`
	VATNumber   string   `xml:"vat_number,omitempty" json:"vat_number"`
	Street1     string   `xml:"street1,omitempty" json:"address_1"`
	Street2     string   `xml:"street2,omitempty" json:"address_2"`
	City        string   `xml:"city,omitemtpy" json:"city"`
	State       string   `xml:"state,omitempty" json:"state"`
	Zip         string   `xml:"zip,omitempty" json:"zip"`
	Country     string   `xml:"country,omitempty" json:"country"`
	Email       string   `xml:"email,omitempty" json:"email"`
	Phone       string   `xml:"phone,omitempty" json:"phone"`
}

type DunningEvent struct {
}

// Transaction constants.
const (
	TransactionFailureTypeDeclined  = "declined"
	TransactionFailureTypeDuplicate = "duplicate_transaction"
)

// Account types.
type (
	NewAccountNotification struct {
		Account Account `xml:"account" json:"account"`
	}

	UpdatedAccountNotification struct {
		Account Account `xml:"account" json:"account"`
	}

	ReactivatedAccountNotification struct {
		Account Account `xml:"account" json:"account"`
	}
	// BillingInfoUpdatedNotification is sent when a customer updates or adds billing information.
	// https://dev.recurly.com/page/webhooks#section-updated-billing-information
	BillingInfoUpdatedNotification struct {
		Account Account `xml:"account" json:"account"`
	}
)

// Subscription types.
type (
	// NewSubscriptionNotification is sent when a new subscription is created.
	// https://dev.recurly.com/page/webhooks#section-new-subscription
	NewSubscriptionNotification struct {
		Account      Account              `xml:"account" json:"account"`
		Subscription recurly.Subscription `xml:"subscription" json:"subscription"`
	}

	// UpdatedSubscriptionNotification is sent when a subscription is upgraded or downgraded.
	// https://dev.recurly.com/page/webhooks#section-updated-subscription
	UpdatedSubscriptionNotification struct {
		Account      Account              `xml:"account" json:"account"`
		Subscription recurly.Subscription `xml:"subscription" json:"subscription"`
	}

	// RenewedSubscriptionNotification is sent when a subscription renew.
	// https://dev.recurly.com/page/webhooks#section-renewed-subscription
	RenewedSubscriptionNotification struct {
		Account      Account              `xml:"account" json:"account"`
		Subscription recurly.Subscription `xml:"subscription" json:"subscription"`
	}

	// ExpiredSubscriptionNotification is sent when a subscription is no longer valid.
	// https://dev.recurly.com/v2.4/page/webhooks#section-expired-subscription
	ExpiredSubscriptionNotification struct {
		Account      Account              `xml:"account" json:"account"`
		Subscription recurly.Subscription `xml:"subscription" json:"subscription"`
	}

	// CanceledSubscriptionNotification is sent when a subscription is canceled.
	// https://dev.recurly.com/page/webhooks#section-canceled-subscription
	CanceledSubscriptionNotification struct {
		Account      Account              `xml:"account" json:"account"`
		Subscription recurly.Subscription `xml:"subscription" json:"subscription"`
	}
)

// Invoice types.
type (
	// NewInvoiceNotification is sent when an invoice generated.
	// https://dev.recurly.com/page/webhooks#section-new-invoice
	NewInvoiceNotification struct {
		Account Account `xml:"account" json:"account"`
		Invoice Invoice `xml:"invoice" json:"invoice"`
	}

	// PastDueInvoiceNotification is sent when an invoice is past due.
	// https://dev.recurly.com/v2.4/page/webhooks#section-past-due-invoice
	PastDueInvoiceNotification struct {
		Account Account `xml:"account" json:"account"`
		Invoice Invoice `xml:"invoice" json:"invoice"`
	}

	ProcessingInvoiceNotification struct {
		Account Account `xml:"account" json:"account"`
		Invoice Invoice `xml:"invoice" json:"invoice"`
	}

	ClosedInvoiceNotification struct {
		Account Account `xml:"account" json:"account"`
		Invoice Invoice `xml:"invoice" json:"invoice"`
	}
)

// Payment types.
type (
	// SuccessfulPaymentNotification is sent when a payment is successful.
	// https://dev.recurly.com/v2.4/page/webhooks#section-successful-payment
	SuccessfulPaymentNotification struct {
		Account     Account     `xml:"account" json:"account"`
		Transaction Transaction `xml:"transaction" json:"transaction"`
	}

	// FailedPaymentNotification is sent when a payment fails.
	// https://dev.recurly.com/v2.4/page/webhooks#section-failed-payment
	FailedPaymentNotification struct {
		Account     Account     `xml:"account" json:"account"`
		Transaction Transaction `xml:"transaction" json:"transaction"`
	}

	// VoidPaymentNotification is sent when a successful payment is voided.
	// https://dev.recurly.com/page/webhooks#section-void-payment
	VoidPaymentNotification struct {
		Account     Account     `xml:"account" json:"account"`
		Transaction Transaction `xml:"transaction" json:"transaction"`
	}

	// SuccessfulRefundNotification is sent when an amount is refunded.
	// https://dev.recurly.com/page/webhooks#section-successful-refund
	SuccessfulRefundNotification struct {
		Account     Account     `xml:"account" json:"account"`
		Transaction Transaction `xml:"transaction" json:"transaction"`
	}
)

// Shipping Address types.
type (
	NewShippingAddressNotification struct {
		Account          Account          `xml:"account" json:"account"`
		ShippingAdddress ShippingAdddress `xml:"shipping_address" json:"shipping_address"`
	}

	UpdatedShippingAddressNotification struct {
		Account          Account          `xml:"account" json:"account"`
		ShippingAdddress ShippingAdddress `xml:"shipping_address" json:"shipping_address"`
	}

	DeletedShippingAddressNotification struct {
		Account          Account          `xml:"account" json:"account"`
		ShippingAdddress ShippingAdddress `xml:"shipping_address" json:"shipping_address"`
	}
)

type NewDunningEventNotification struct {
	Account     Account              `xml:"account" json:"account"`
	Invoice     Invoice              `xml:"invoice" json:"invoice"`
	Subsription recurly.Subscription `xml:"subscription" json:"subscription"`
	Transaction Transaction          `xml:"transaction" json:"transaction"`
}

// ErrUnknownNotification is used when the incoming webhook does not match a
// predefined notification type. It implements the error interface.
type ErrUnknownNotification struct {
	name string
}

// Error implements the error interface.
func (e ErrUnknownNotification) Error() string {
	return fmt.Sprintf("unknown notification: %s", e.name)
}

// Name returns the name of the unknown notification.
func (e ErrUnknownNotification) Name() string {
	return e.name
}

type ParseResponse struct {
	Message string
	Data    interface{}
}

// Parse parses an incoming webhook and returns the notification.
func Parse(r io.Reader) (*ParseResponse, error) {
	if closer, ok := r.(io.Closer); ok {
		defer closer.Close()
	}

	notification, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var n notificationName
	if err := xml.Unmarshal(notification, &n); err != nil {
		return nil, err
	}

	var dst interface{}
	switch n.XMLName.Local {
	case NewAccount:
		dst = &NewAccountNotification{}
	case UpdatedAccount:
		dst = &UpdatedAccountNotification{}
	case ReactivatedSubscription:
		dst = &ReactivatedAccountNotification{}
	case BillingInfoUpdated:
		dst = &BillingInfoUpdatedNotification{}
	case NewSubscription:
		dst = &NewSubscriptionNotification{}
	case UpdatedSubscription:
		dst = &UpdatedSubscriptionNotification{}
	case RenewedSubscription:
		dst = &RenewedSubscriptionNotification{}
	case ExpiredSubscription:
		dst = &ExpiredSubscriptionNotification{}
	case CanceledSubscription:
		dst = &CanceledSubscriptionNotification{}
	case NewInvoice:
		dst = &NewInvoiceNotification{}
	case PastDueInvoice:
		dst = &PastDueInvoiceNotification{}
	case ProcessingInvoice:
		dst = &ProcessingInvoiceNotification{}
	case ClosedInvoice:
		dst = &ClosedInvoiceNotification{}
	case SuccessfulPayment:
		dst = &SuccessfulPaymentNotification{}
	case FailedPayment:
		dst = &FailedPaymentNotification{}
	case VoidPayment:
		dst = &VoidPaymentNotification{}
	case SuccessfulRefund:
		dst = &SuccessfulRefundNotification{}
	case NewShippingAddress:
		dst = &NewShippingAddressNotification{}
	case UpdatedShippingAddress:
		dst = &UpdatedShippingAddressNotification{}
	case DeletedShippingAddress:
		dst = &DeletedShippingAddressNotification{}
	case NewDunningEvent:
		dst = &NewDunningEventNotification{}
	default:
		return nil, ErrUnknownNotification{name: n.XMLName.Local}
	}

	if err := xml.Unmarshal(notification, dst); err != nil {
		return nil, err
	}

	response := &ParseResponse{
		Message: n.XMLName.Local,
		Data:    dst,
	}
	return response, nil
}
