package scanner

//Secret
type Secret struct {
	Name        string
	Description string
	Regex       string
	Poc         string
}

var regexes = []Secret{
	{
		"AWS Access Key",
		"AWS Access Key",
		"(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}",
		"?",
	},
	{
		"AWS Secret Key",
		"AWS Secret Key",
		`(?i)aws(.{0,20})?(?-i)['\"][0-9a-zA-Z\/+]{40}['\"]`,
		"?",
	},
	{
		"AWS MWS Key",
		"AWS MWS Key",
		`amzn\.mws\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`,
		"?",
	},
	{
		"Facebook Secret Key",
		"Facebook Secret Key",
		"(?i)(facebook|fb)(.{0,20})?(?-i)['\"][0-9a-f]{32}['\"]",
		"?",
	},
	{
		"Facebook Client ID",
		"Facebook Client ID",
		"(?i)(facebook|fb)(.{0,20})?['\"][0-9]{13,17}['\"]",
		"?",
	},
	{
		"Twitter Secret Key",
		"Twitter Secret Key",
		"(?i)twitter(.{0,20})?[0-9a-z]{35,44}",
		"?",
	},
	{
		"Twitter Client ID",
		"Twitter Client ID",
		"(?i)twitter(.{0,20})?[0-9a-z]{18,25}",
		"?",
	},
	{
		"Github Personal Access Token",
		"Github Personal Access Token",
		"ghp_[0-9a-zA-Z]{36}",
		"?",
	},
	{
		"Github OAuth Access Token",
		"Github OAuth Access Token",
		"gho_[0-9a-zA-Z]{36}",
		"?",
	},
	{
		"Github App Token",
		"Github App Token",
		"(ghu|ghs)_[0-9a-zA-Z]{36}",
		"?",
	},
	{
		"Github Refresh Token",
		"Github Refresh Token",
		"ghr_[0-9a-zA-Z]{76}",
		"?",
	},
	{
		"LinkedIn Client ID",
		"LinkedIn Client ID",
		"(?i)linkedin(.{0,20})?(?-i)[0-9a-z]{12}",
		"?",
	},
	{
		"LinkedIn Secret Key",
		"LinkedIn Secret Key",
		"(?i)linkedin(.{0,20})?[0-9a-z]{16}",
		"?",
	},
	{
		"Slack",
		"Slack",
		"xox[baprs]-([0-9a-zA-Z]{10,48})?",
		"?",
	},
	{
		"Asymmetric Private Key",
		"Asymmetric Private Key",
		"-----BEGIN ((EC|PGP|DSA|RSA|OPENSSH) )?PRIVATE KEY( BLOCK)?-----",
		"?",
	},
	{
		"Google API key",
		"Google API key",
		"AIza[0-9A-Za-z\\-_]{35}",
		"?",
	},
	{
		"Google (GCP) Service Account",
		"Slack",
		`"type": "service_account"`,
		"?",
	},
	{
		"Heroku API key",
		"Heroku API key",
		"(?i)heroku(.{0,20})?[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
		"?",
	},
	{
		"MailChimp API key",
		"MailChimp API key",
		"(?i)(mailchimp|mc)(.{0,20})?[0-9a-f]{32}-us[0-9]{1,2}",
		"?",
	},
	{
		"Mailgun API key",
		"Mailgun API key",
		"((?i)(mailgun|mg)(.{0,20})?)?key-[0-9a-z]{32}",
		"?",
	},
	{
		"PayPal Braintree access token",
		"PayPal Braintree access token",
		`access_token\$production\$[0-9a-z]{16}\$[0-9a-f]{32}`,
		"?",
	},
	{
		"Picatic API key",
		"Picatic API key",
		"sk_live_[0-9a-z]{32}",
		"?",
	},
	{
		"SendGrid API Key",
		"SendGrid API Key",
		`SG\.[\w_]{16,32}\.[\w_]{16,64}`,
		"?",
	},
	{
		"Slack Webhook",
		"Slack Webhook",
		"https://hooks.slack.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8,12}/[a-zA-Z0-9_]{24}",
		"?",
	},
	{
		"Stripe API key",
		"Stripe API key",
		"(?i)stripe(.{0,20})?[sr]k_live_[0-9a-zA-Z]{24}",
		"?",
	},
	{
		"Square access token",
		"Square access token",
		`sq0atp-[0-9A-Za-z\-_]{22}`,
		"?",
	},
	{
		"Square OAuth secret",
		"Square OAuth secret",
		"sq0csp-[0-9A-Za-z\\-_]{43}",
		"?",
	},
	{
		"Twilio API key",
		"Twilio API key",
		"(?i)twilio(.{0,20})?SK[0-9a-f]{32}",
		"?",
	},
	{
		"Dynatrace token",
		"Dynatrace token",
		`dt0[a-zA-Z]{1}[0-9]{2}\.[A-Z0-9]{24}\.[A-Z0-9]{64}`,
		"?",
	},
	{
		"Shopify shared secret",
		"Shopify shared secret",
		"shpss_[a-fA-F0-9]{32}",
		"?",
	},
	{
		"Shopify access token",
		"Shopify access token",
		"shpat_[a-fA-F0-9]{32}",
		"?",
	},
	{
		"Shopify custom app access token",
		"Shopify custom app access token",
		"shpca_[a-fA-F0-9]{32}",
		"?",
	},
	{
		"Shopify private app access token",
		"Shopify private app access token",
		"shppa_[a-fA-F0-9]{32}",
		"?",
	},
	{
		"PyPI upload token",
		"PyPI upload token",
		"pypi-AgEIcHlwaS5vcmc[A-Za-z0-9-_]{50,1000}",
		"?",
	},
}
