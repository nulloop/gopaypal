package gopaypal

type Operation string

func (o Operation) String() string {
	return string(o)
}

const (
	AddOperation     Operation = "add"
	RemoveOperation  Operation = "remove"
	ReplaceOperation Operation = "replace"
	MoveOperation    Operation = "move"
	CopyOperation    Operation = "copy"
	TestOperation    Operation = "test"
)

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
	DeletedBillingPlanState  BillingPlanState = "DELETED"
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

type PaymentMethod string

func (p PaymentMethod) String() string {
	return string(p)
}

const (
	BankPaymentMethod   PaymentMethod = "BANK"
	PaypalPaymentMethod PaymentMethod = "PAYPAL"
)

type CreditCardState string

func (c CreditCardState) String() string {
	return string(c)
}

const (
	ExpiredCreditCardState CreditCardState = "EXPIRED"
	OKCreditCardState      CreditCardState = "OK"
)

type CountryCode string

func (c CountryCode) String() string {
	return string(c)
}

const (
	ALBANIA                CountryCode = "AL"
	ALGERIA                CountryCode = "DZ"
	ANDORRA                CountryCode = "AD"
	ANGOLA                 CountryCode = "AO"
	ANGUILLA               CountryCode = "AI"
	ANTIGUA_BARBUDA        CountryCode = "AG"
	ARGENTINA              CountryCode = "AR"
	ARMENIA                CountryCode = "AM"
	ARUBA                  CountryCode = "AW"
	AUSTRALIA              CountryCode = "AU"
	AUSTRIA                CountryCode = "AT"
	AZERBAIJAN             CountryCode = "AZ"
	BAHAMAS                CountryCode = "BS"
	BAHRAIN                CountryCode = "BH"
	BARBADOS               CountryCode = "BB"
	BELARUS                CountryCode = "BY"
	BELGIUM                CountryCode = "BE"
	BELIZE                 CountryCode = "BZ"
	BENIN                  CountryCode = "BJ"
	BERMUDA                CountryCode = "BM"
	BHUTAN                 CountryCode = "BT"
	BOLIVIA                CountryCode = "BO"
	BOSNIA_HERZEGOVINA     CountryCode = "BA"
	BOTSWANA               CountryCode = "BW"
	BRAZIL                 CountryCode = "BR"
	BRITISH_VIRGIN_ISLANDS CountryCode = "VG"
	BRUNEI                 CountryCode = "BN"
	BULGARIA               CountryCode = "BG"
	BURKINA_FASO           CountryCode = "BF"
	BURUNDI                CountryCode = "BI"
	CAMBODIA               CountryCode = "KH"
	CAMEROON               CountryCode = "CM"
	CANADA                 CountryCode = "CA"
	CAPE_VERDE             CountryCode = "CV"
	CAYMAN_ISLANDS         CountryCode = "KY"
	CHAD                   CountryCode = "TD"
	CHILE                  CountryCode = "CL"
	CHINA                  CountryCode = "C2"
	COLOMBIA               CountryCode = "CO"
	COMOROS                CountryCode = "KM"
	CONGO_BRAZZAVILLE      CountryCode = "CG"
	CONGO_KINSHASA         CountryCode = "CD"
	COOK_ISLANDS           CountryCode = "CK"
	COSTA_RICA             CountryCode = "CR"
	CÔTE_DIVOIRE           CountryCode = "CI"
	CROATIA                CountryCode = "HR"
	CYPRUS                 CountryCode = "CY"
	CZECH_REPUBLIC         CountryCode = "CZ"
	DENMARK                CountryCode = "DK"
	DJIBOUTI               CountryCode = "DJ"
	DOMINICA               CountryCode = "DM"
	DOMINICAN_REPUBLIC     CountryCode = "DO"
	ECUADOR                CountryCode = "EC"
	EGYPT                  CountryCode = "EG"
	EL_SALVADOR            CountryCode = "SV"
	ERITREA                CountryCode = "ER"
	ESTONIA                CountryCode = "EE"
	ETHIOPIA               CountryCode = "ET"
	FALKLAND_ISLANDS       CountryCode = "FK"
	FAROE_ISLANDS          CountryCode = "FO"
	FIJI                   CountryCode = "FJ"
	FINLAND                CountryCode = "FI"
	FRANCE                 CountryCode = "FR"
	FRENCH_GUIANA          CountryCode = "GF"
	FRENCH_POLYNESIA       CountryCode = "PF"
	GABON                  CountryCode = "GA"
	GAMBIA                 CountryCode = "GM"
	GEORGIA                CountryCode = "GE"
	GERMANY                CountryCode = "DE"
	GIBRALTAR              CountryCode = "GI"
	GREECE                 CountryCode = "GR"
	GREENLAND              CountryCode = "GL"
	GRENADA                CountryCode = "GD"
	GUADELOUPE             CountryCode = "GP"
	GUATEMALA              CountryCode = "GT"
	GUINEA                 CountryCode = "GN"
	GUINEA_BISSAU          CountryCode = "GW"
	GUYANA                 CountryCode = "GY"
	HONDURAS               CountryCode = "HN"
	HONG_KONG_SAR_CHINA    CountryCode = "HK"
	HUNGARY                CountryCode = "HU"
	ICELAND                CountryCode = "IS"
	INDIA                  CountryCode = "IN"
	INDONESIA              CountryCode = "ID"
	IRELAND                CountryCode = "IE"
	ISRAEL                 CountryCode = "IL"
	ITALY                  CountryCode = "IT"
	JAMAICA                CountryCode = "JM"
	JAPAN                  CountryCode = "JP"
	JORDAN                 CountryCode = "JO"
	KAZAKHSTAN             CountryCode = "KZ"
	KENYA                  CountryCode = "KE"
	KIRIBATI               CountryCode = "KI"
	KUWAIT                 CountryCode = "KW"
	KYRGYZSTAN             CountryCode = "KG"
	LAOS                   CountryCode = "LA"
	LATVIA                 CountryCode = "LV"
	LESOTHO                CountryCode = "LS"
	LIECHTENSTEIN          CountryCode = "LI"
	LITHUANIA              CountryCode = "LT"
	LUXEMBOURG             CountryCode = "LU"
	MACEDONIA              CountryCode = "MK"
	MADAGASCAR             CountryCode = "MG"
	MALAWI                 CountryCode = "MW"
	MALAYSIA               CountryCode = "MY"
	MALDIVES               CountryCode = "MV"
	MALI                   CountryCode = "ML"
	MALTA                  CountryCode = "MT"
	MARSHALL_ISLANDS       CountryCode = "MH"
	MARTINIQUE             CountryCode = "MQ"
	MAURITANIA             CountryCode = "MR"
	MAURITIUS              CountryCode = "MU"
	MAYOTTE                CountryCode = "YT"
	MEXICO                 CountryCode = "MX"
	MICRONESIA             CountryCode = "FM"
	MOLDOVA                CountryCode = "MD"
	MONACO                 CountryCode = "MC"
	MONGOLIA               CountryCode = "MN"
	MONTENEGRO             CountryCode = "ME"
	MONTSERRAT             CountryCode = "MS"
	MOROCCO                CountryCode = "MA"
	MOZAMBIQUE             CountryCode = "MZ"
	NAMIBIA                CountryCode = "NA"
	NAURU                  CountryCode = "NR"
	NEPAL                  CountryCode = "NP"
	NETHERLANDS            CountryCode = "NL"
	NEW_CALEDONIA          CountryCode = "NC"
	NEW_ZEALAND            CountryCode = "NZ"
	NICARAGUA              CountryCode = "NI"
	NIGER                  CountryCode = "NE"
	NIGERIA                CountryCode = "NG"
	NIUE                   CountryCode = "NU"
	NORFOLK_ISLAND         CountryCode = "NF"
	NORWAY                 CountryCode = "NO"
	OMAN                   CountryCode = "OM"
	PALAU                  CountryCode = "PW"
	PANAMA                 CountryCode = "PA"
	PAPUA_NEW_GUINEA       CountryCode = "PG"
	PARAGUAY               CountryCode = "PY"
	PERU                   CountryCode = "PE"
	PHILIPPINES            CountryCode = "PH"
	PITCAIRN_ISLANDS       CountryCode = "PN"
	POLAND                 CountryCode = "PL"
	PORTUGAL               CountryCode = "PT"
	QATAR                  CountryCode = "QA"
	RÉUNION                CountryCode = "RE"
	ROMANIA                CountryCode = "RO"
	RUSSIA                 CountryCode = "RU"
	RWANDA                 CountryCode = "RW"
	SAMOA                  CountryCode = "WS"
	SAN_MARINO             CountryCode = "SM"
	SÃO_TOMÉ_PRÍNCIPE      CountryCode = "ST"
	SAUDI_ARABIA           CountryCode = "SA"
	SENEGAL                CountryCode = "SN"
	SERBIA                 CountryCode = "RS"
	SEYCHELLES             CountryCode = "SC"
	SIERRA_LEONE           CountryCode = "SL"
	SINGAPORE              CountryCode = "SG"
	SLOVAKIA               CountryCode = "SK"
	SLOVENIA               CountryCode = "SI"
	SOLOMON_ISLANDS        CountryCode = "SB"
	SOMALIA                CountryCode = "SO"
	SOUTH_AFRICA           CountryCode = "ZA"
	SOUTH_KOREA            CountryCode = "KR"
	SPAIN                  CountryCode = "ES"
	SRI_LANKA              CountryCode = "LK"
	ST_HELENA              CountryCode = "SH"
	ST_KITTS_NEVIS         CountryCode = "KN"
	ST_LUCIA               CountryCode = "LC"
	ST_PIERRE_MIQUELON     CountryCode = "PM"
	ST_VINCENT_GRENADINES  CountryCode = "VC"
	SURINAME               CountryCode = "SR"
	SVALBARD_JAN_MAYEN     CountryCode = "SJ"
	SWAZILAND              CountryCode = "SZ"
	SWEDEN                 CountryCode = "SE"
	SWITZERLAND            CountryCode = "CH"
	TAIWAN                 CountryCode = "TW"
	TAJIKISTAN             CountryCode = "TJ"
	TANZANIA               CountryCode = "TZ"
	THAILAND               CountryCode = "TH"
	TOGO                   CountryCode = "TG"
	TONGA                  CountryCode = "TO"
	TRINIDAD_TOBAGO        CountryCode = "TT"
	TUNISIA                CountryCode = "TN"
	TURKMENISTAN           CountryCode = "TM"
	TURKS_CAICOS_ISLANDS   CountryCode = "TC"
	TUVALU                 CountryCode = "TV"
	UGANDA                 CountryCode = "UG"
	UKRAINE                CountryCode = "UA"
	UNITED_ARAB_EMIRATES   CountryCode = "AE"
	UNITED_KINGDOM         CountryCode = "GB"
	UNITED_STATES          CountryCode = "US"
	URUGUAY                CountryCode = "UY"
	VANUATU                CountryCode = "VU"
	VATICAN_CITY           CountryCode = "VA"
	VENEZUELA              CountryCode = "VE"
	VIETNAM                CountryCode = "VN"
	WALLIS_FUTUNA          CountryCode = "WF"
	YEMEN                  CountryCode = "YE"
	ZAMBIA                 CountryCode = "ZM"
	ZIMBABWE               CountryCode = "ZW"
)
