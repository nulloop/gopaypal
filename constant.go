package gopaypal

type Bool string

const (
	Yes Bool = "YES"
	No  Bool = "NO"
)

func (b Bool) String() string {
	return string(b)
}

type Mode string

func (m Mode) String() string {
	return string(m)
}

const (
	LiveMode    Mode = "api.paypal.com"
	SandboxMode Mode = "api.sandbox.paypal.com"
)

type BillingPlanType string

func (b BillingPlanType) String() string {
	return string(b)
}

const (
	FixedBillingPlanType    BillingPlanType = "FIXED"
	InfiniteBillingPlanType BillingPlanType = "INFINITE"
)

type PaymentDefinitionType string

func (p PaymentDefinitionType) String() string {
	return string(p)
}

const (
	TrialPaymentDefinitionType   PaymentDefinitionType = "TRIAL"
	RegularPaymentDefinitionType PaymentDefinitionType = "REGULAR"
)

type FrequencyType string

func (f FrequencyType) String() string {
	return string(f)
}

const (
	DayFrequencyType   = "DAY"
	WeekFrequencyType  = "WEEK"
	MonthFrequencyType = "MONTH"
	YearFrequencyType  = "YEAR"
)

type ChargeModelType string

func (c ChargeModelType) String() string {
	return string(c)
}

const (
	TaxChargeModelType      ChargeModelType = "TAX"
	ShippingChargeModelType ChargeModelType = "SHIPPING"
)

type Currency string

func (c Currency) String() string {
	return string(c)
}

const (
	CanadianDollar     Currency = "CAD"
	Euro               Currency = "EUR"
	UnitedStatesDollar Currency = "USD"
)

type InitialAction string

func (i InitialAction) String() string {
	return string(i)
}

const (
	ContinueInitialAction InitialAction = "CONTINUE"
	CancelInitialAction   InitialAction = "CANCEL"
)

type BillingPlanState string

func (b BillingPlanState) String() string {
	return string(b)
}

const (
	CreatedBillingPlanState  BillingPlanState = "CREATED"
	ActiveBillingPlanState   BillingPlanState = "ACTIVE"
	InactiveBillingPlanState BillingPlanState = "INACTIVE"
)

type BillingPlanTermType string

func (t BillingPlanTermType) String() string {
	return string(t)
}

const (
	WeeklyBillingPlanTermType  BillingPlanTermType = "WEEKLY"
	MonthlyBillingPlanTermType BillingPlanTermType = "MONTHLY"
	YearlyBillingPlanTermType  BillingPlanTermType = "YEARLY"
)
