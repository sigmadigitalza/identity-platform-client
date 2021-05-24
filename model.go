package identity_platform

type ApiErrorResponse struct {
	Error *ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type AddProductRequest struct {
	Name    string `json:"name"`
	Service string `json:"service"`
}

type Config struct {
	Name              string                 `json:"name"`
	SignIn            *SignInConfig          `json:"signIn"`
	Notification      *NotificationConfig    `json:"notification"`
	Quota             *QuotaConfig           `json:"quota"`
	Monitoring        *MonitoringConfig      `json:"monitoring"`
	MultiTenant       *MultiTenantConfig     `json:"multiTenant"`
	Subtype           string                 `json:"subtype"`
	Client            *ClientConfig          `json:"client"`
	AuthorizedDomains []string               `json:"authorizedDomains"`
	Mfa               *MultiFactorAuthConfig `json:"mfa"`
}

type SignInConfig struct {
	Email                *Email       `json:"email"`
	PhoneNumber          *PhoneNumber `json:"phoneNumber"`
	Anonymous            *Anonymous   `json:"anonymous"`
	AllowDuplicateEmails bool         `json:"allowDuplicateEmails"`
	HashConfig           *HashConfig  `json:"hashConfig"`
}

type Email struct {
	Enabled          bool `json:"enabled"`
	PasswordRequired bool `json:"passwordRequired"`
}

type PhoneNumber struct {
	Enabled          bool               `json:"enabled"`
	TestPhoneNumbers *map[string]string `json:"testPhoneNumbers"`
}

type Anonymous struct {
	Enabled bool `json:"enabled"`
}

type HashConfig struct {
	Algorithm     string `json:"algorithm"`
	SignerKey     string `json:"signerKey"`
	SaltSeparator string `json:"saltSeparator"`
	Rounds        int    `json:"rounds"`
	MemoryCost    int    `json:"memoryCost"`
}

type NotificationConfig struct {
	SendEmail     *SendEmail `json:"sendEmail"`
	SendSms       *SendSms   `json:"sendSms"`
	DefaultLocale string     `json:"defaultLocale"`
}

type SendEmail struct {
	Method                             string         `json:"method"`
	ResetPasswordTemplate              *EmailTemplate `json:"resetPasswordTemplate"`
	VerifyEmailTemplate                *EmailTemplate `json:"verifyEmailTemplate"`
	ChangeEmailTemplate                *EmailTemplate `json:"changeEmailTemplate"`
	LegacyResetPasswordTemplate        *EmailTemplate `json:"legacyResetPasswordTemplate"`
	CallbackUri                        string         `json:"callbackUri"`
	DnsInfo                            *DnsInfo       `json:"dnsInfo"`
	RevertSecondFactorAdditionTemplate *EmailTemplate `json:"revertSecondFactorAdditionTemplate"`
	Smtp                               *Smtp          `json:"smtp"`
}

type EmailTemplate struct {
	SenderLocalPart   string `json:"senderLocalPart"`
	Subject           string `json:"subject"`
	SenderDisplayName string `json:"senderDisplayName"`
	Body              string `json:"body"`
	BodyFormat        string `json:"bodyFormat"`
	ReplyTo           string `json:"replyTo"`
	Customized        bool   `json:"customized"`
}

type DnsInfo struct {
	CustomDomain                  string `json:"customDomain"`
	UseCustomDomain               bool   `json:"useCustomDomain"`
	PendingCustomDomain           string `json:"pendingCustomDomain"`
	CustomDomainState             string `json:"customDomainState"`
	DomainVerificationRequestTime string `json:"domainVerificationRequestTime"`
}

type Smtp struct {
	SenderEmail  string `json:"senderEmail"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	SecurityMode string `json:"securityMode"`
}

type SendSms struct {
	UseDeviceLocale bool         `json:"useDeviceLocale"`
	SmsTemplate     *SmsTemplate `json:"smsTemplate"`
}

type SmsTemplate struct {
	Content string `json:"content"`
}

type QuotaConfig struct {
	SignUpQuotaConfig *TemporaryQuota `json:"signUpQuotaConfig"`
}

type TemporaryQuota struct {
	Quota             string `json:"quota"`
	StartTime         string `json:"startTime"`
	QuotaDistribution string `json:"quotaDistribution"`
}

type MonitoringConfig struct {
	RequestLogging *RequestLogging `json:"requestLogging"`
}

type RequestLogging struct {
	Enabled bool `json:"enabled"`
}

type MultiTenantConfig struct {
	AllowTenants          bool   `json:"allowTenants"`
	DefaultTenantLocation string `json:"defaultTenantLocation"`
}

type ClientConfig struct {
	ApiKey            string       `json:"apiKey"`
	Permissions       *Permissions `json:"permissions"`
	FirebaseSubdomain string       `json:"firebaseSubdomain"`
}

type Permissions struct {
	DisabledUserSignup   bool `json:"disabledUserSignup"`
	DisabledUserDeletion bool `json:"disabledUserDeletion"`
}

type MultiFactorAuthConfig struct {
	State            string   `json:"state"`
	EnabledProviders []string `json:"enabledProviders"`
}
