package recurly

import (
	"encoding/xml"
	"strings"
)

// SanitizeUUID returns the uuid without dashes.
func SanitizeUUID(id string) string {
	return strings.TrimSpace(strings.Replace(id, "-", "", -1))
}

const (
	// SubscriptionStateActive represents subscriptions that are valid for the
	// current time. This includes subscriptions in a trial period
	SubscriptionStateActive = "active"

	// SubscriptionStateCanceled are subscriptions that are valid for
	// the current time but will not renew because a cancelation was requested
	SubscriptionStateCanceled = "canceled"

	// SubscriptionStateExpired are subscriptions that have expired and are no longer valid
	SubscriptionStateExpired = "expired"

	// SubscriptionStateFuture are subscriptions that will start in the
	// future, they are not active yet
	SubscriptionStateFuture = "future"

	// SubscriptionStateInTrial are subscriptions that are active or canceled
	// and are in a trial period
	SubscriptionStateInTrial = "in_trial"

	// SubscriptionStateLive are all subscriptions that are not expired
	SubscriptionStateLive = "live"

	// SubscriptionStatePastDue are subscriptions that are active or canceled
	// and have a past-due invoice
	SubscriptionStatePastDue = "past_due"
)

// Subscription represents an individual subscription.
type Subscription struct {
	XMLName                xml.Name             `xml:"subscription" json:"-"`
	Plan                   NestedPlan           `xml:"plan,omitempty" json:"plan"`
	AccountCode            string               `xml:"-" json:"-"`
	InvoiceNumber          int                  `xml:"-" json:"-"`
	UUID                   string               `xml:"uuid,omitempty" json:"uuid"`
	State                  string               `xml:"state,omitempty" json:"state"`
	UnitAmountInCents      int                  `xml:"unit_amount_in_cents,omitempty" json:"unit_amount_in_cents"`
	Currency               string               `xml:"currency,omitempty" json:"currency"`
	Quantity               int                  `xml:"quantity,omitempty" json:"quantity"`
	TotalAmountInCents     int                  `xml:"total_amount_in_cents,omitempty" json:"total_amount_in_cents"`
	ActivatedAt            NullTime             `xml:"activated_at,omitempty" json:"activated_at"`
	CanceledAt             NullTime             `xml:"canceled_at,omitempty" json:"canceled_at"`
	ExpiresAt              NullTime             `xml:"expires_at,omitempty" json:"expires_at"`
	CurrentPeriodStartedAt NullTime             `xml:"current_period_started_at,omitempty" json:"current_period_started_at"`
	CurrentPeriodEndsAt    NullTime             `xml:"current_period_ends_at,omitempty" json:"current_period_ends_at"`
	TrialStartedAt         NullTime             `xml:"trial_started_at,omitempty" json:"trial_started_at"`
	TrialEndsAt            NullTime             `xml:"trial_ends_at,omitempty" json:"trial_ends_at"`
	TaxInCents             int                  `xml:"tax_in_cents,omitempty" json:"tax_in_cents"`
	TaxType                string               `xml:"tax_type,omitempty" json:"tax_type"`
	TaxRegion              string               `xml:"tax_region,omitempty" json:"tax_region"`
	TaxRate                float64              `xml:"tax_rate,omitempty" json:"tax_rate"`
	PONumber               string               `xml:"po_number,omitempty" json:"po_number"`
	NetTerms               NullInt              `xml:"net_terms,omitempty" json:"net_terms"`
	SubscriptionAddOns     []SubscriptionAddOn  `xml:"subscription_add_ons>subscription_add_on,omitempty" json:"-"`
	PendingSubscription    *PendingSubscription `xml:"pending_subscription,omitempty" json:"pending_subscription,omitempty"`
}

// UnmarshalXML unmarshals transactions and handles intermediary state during unmarshaling
// for types like href.
func (s *Subscription) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v struct {
		XMLName                xml.Name             `xml:"subscription"`
		Plan                   NestedPlan           `xml:"plan,omitempty"`
		AccountCode            hrefString           `xml:"account"`
		InvoiceNumber          hrefInt              `xml:"invoice"`
		UUID                   string               `xml:"uuid,omitempty"`
		State                  string               `xml:"state,omitempty"`
		UnitAmountInCents      int                  `xml:"unit_amount_in_cents,omitempty"`
		Currency               string               `xml:"currency,omitempty"`
		Quantity               int                  `xml:"quantity,omitempty"`
		TotalAmountInCents     int                  `xml:"total_amount_in_cents,omitempty"`
		ActivatedAt            NullTime             `xml:"activated_at,omitempty"`
		CanceledAt             NullTime             `xml:"canceled_at,omitempty"`
		ExpiresAt              NullTime             `xml:"expires_at,omitempty"`
		CurrentPeriodStartedAt NullTime             `xml:"current_period_started_at,omitempty"`
		CurrentPeriodEndsAt    NullTime             `xml:"current_period_ends_at,omitempty"`
		TrialStartedAt         NullTime             `xml:"trial_started_at,omitempty"`
		TrialEndsAt            NullTime             `xml:"trial_ends_at,omitempty"`
		TaxInCents             int                  `xml:"tax_in_cents,omitempty"`
		TaxType                string               `xml:"tax_type,omitempty"`
		TaxRegion              string               `xml:"tax_region,omitempty"`
		TaxRate                float64              `xml:"tax_rate,omitempty"`
		PONumber               string               `xml:"po_number,omitempty"`
		NetTerms               NullInt              `xml:"net_terms,omitempty"`
		SubscriptionAddOns     []SubscriptionAddOn  `xml:"subscription_add_ons>subscription_add_on,omitempty"`
		PendingSubscription    *PendingSubscription `xml:"pending_subscription,omitempty"`
	}
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}
	*s = Subscription{
		XMLName:                v.XMLName,
		Plan:                   v.Plan,
		AccountCode:            string(v.AccountCode),
		InvoiceNumber:          int(v.InvoiceNumber),
		UUID:                   v.UUID,
		State:                  v.State,
		UnitAmountInCents:      v.UnitAmountInCents,
		Currency:               v.Currency,
		Quantity:               v.Quantity,
		TotalAmountInCents:     v.TotalAmountInCents,
		ActivatedAt:            v.ActivatedAt,
		CanceledAt:             v.CanceledAt,
		ExpiresAt:              v.ExpiresAt,
		CurrentPeriodStartedAt: v.CurrentPeriodStartedAt,
		CurrentPeriodEndsAt:    v.CurrentPeriodEndsAt,
		TrialStartedAt:         v.TrialStartedAt,
		TrialEndsAt:            v.TrialEndsAt,
		TaxInCents:             v.TaxInCents,
		TaxType:                v.TaxType,
		TaxRegion:              v.TaxRegion,
		TaxRate:                v.TaxRate,
		PONumber:               v.PONumber,
		NetTerms:               v.NetTerms,
		SubscriptionAddOns:     v.SubscriptionAddOns,
		PendingSubscription:    v.PendingSubscription,
	}

	return nil
}

// MakeUpdate creates an UpdateSubscription with values that need to be passed
// on update to be retained (meaning nil/zero values will delete that value).
// After calling MakeUpdate you should modify the struct with your updates.
// Once you're ready you can call client.Subscriptions.Update
func (s Subscription) MakeUpdate() UpdateSubscription {
	return UpdateSubscription{
		// NetTerms need to be copied over because on update they default to 0.
		// This ensures the NetTerms don't get overridden.
		NetTerms:           s.NetTerms,
		SubscriptionAddOns: &s.SubscriptionAddOns,
	}
}

type NestedPlan struct {
	Code string `xml:"plan_code,omitempty" json:"plan_code"`
	Name string `xml:"name,omitempty" json:"name"`
}

// SubscriptionAddOn are add ons to subscriptions.
// https://docs.com/api/subscriptions/subscription-add-ons
type SubscriptionAddOn struct {
	XMLName           xml.Name `xml:"subscription_add_on"`
	Type              string   `xml:"add_on_type,omitempty"`
	Code              string   `xml:"add_on_code"`
	UnitAmountInCents int      `xml:"unit_amount_in_cents"`
	Quantity          int      `xml:"quantity,omitempty"`
}

// PendingSubscription are updates to the subscription or subscription add ons that
// will be made on the next renewal.
type PendingSubscription struct {
	XMLName            xml.Name            `xml:"pending_subscription" json:"-"`
	Plan               NestedPlan          `xml:"plan,omitempty" json:"plan,omitempty"`
	Quantity           int                 `xml:"quantity,omitempty" json:"quantity,omitempty"` // Quantity of subscriptions
	Price              int                 `xml:"unit_amount_in_cents,omitempty" json:"unit_amount_in_cents,omitempty"`
	SubscriptionAddOns []SubscriptionAddOn `xml:"subscription_add_ons>subscription_add_on,omitempty"`
}

// NewSubscription is used to create new subscriptions.
type NewSubscription struct {
	XMLName                 xml.Name             `xml:"subscription"`
	PlanCode                string               `xml:"plan_code"`
	Account                 Account              `xml:"account"`
	SubscriptionAddOns      *[]SubscriptionAddOn `xml:"subscription_add_ons>subscription_add_on,omitempty"`
	CouponCode              string               `xml:"coupon_code,omitempty"`
	UnitAmountInCents       int                  `xml:"unit_amount_in_cents,omitempty"`
	Currency                string               `xml:"currency"`
	Quantity                int                  `xml:"quantity,omitempty"`
	TrialEndsAt             NullTime             `xml:"trial_ends_at,omitempty"`
	StartsAt                NullTime             `xml:"starts_at,omitempty"`
	TotalBillingCycles      int                  `xml:"total_billing_cycles,omitempty"`
	FirstRenewalDate        NullTime             `xml:"first_renewal_date,omitempty"`
	CollectionMethod        string               `xml:"collection_method,omitempty"`
	NetTerms                NullInt              `xml:"net_terms,omitempty"`
	PONumber                string               `xml:"po_number,omitempty"`
	Bulk                    bool                 `xml:"bulk,omitempty"`
	TermsAndConditions      string               `xml:"terms_and_conditions,omitempty"`
	CustomerNotes           string               `xml:"customer_notes,omitempty"`
	VATReverseChargeNotes   string               `xml:"vat_reverse_charge_notes,omitempty"`
	BankAccountAuthorizedAt NullTime             `xml:"bank_account_authorized_at,omitempty"`
}

// NewSubscriptionResponse is used to unmarshal either the subscription or the transaction.
type NewSubscriptionResponse struct {
	Subscription *Subscription
	Transaction  *Transaction // UnprocessableEntity errors return only the transaction
}

// UpdateSubscription is used to update subscriptions
type UpdateSubscription struct {
	XMLName            xml.Name             `xml:"subscription"`
	Timeframe          string               `xml:"timeframe,omitempty"`
	PlanCode           string               `xml:"plan_code,omitempty"`
	Quantity           int                  `xml:"quantity,omitempty"`
	UnitAmountInCents  int                  `xml:"unit_amount_in_cents,omitempty"`
	CollectionMethod   string               `xml:"collection_method,omitempty"`
	NetTerms           NullInt              `xml:"net_terms,omitempty"`
	PONumber           string               `xml:"po_number,omitempty"`
	SubscriptionAddOns *[]SubscriptionAddOn `xml:"subscription_add_ons>subscription_add_on,omitempty"`
}

// SubscriptionNotes is used to update a subscription's notes.
type SubscriptionNotes struct {
	XMLName               xml.Name `xml:"subscription"`
	TermsAndConditions    string   `xml:"terms_and_conditions,omitempty"`
	CustomerNotes         string   `xml:"customer_notes,omitempty"`
	VATReverseChargeNotes string   `xml:"vat_reverse_charge_notes,omitempty"`
}
