package resource

type Resource struct {
	ErrorMessage  string
	UserName      string
	Domains       []string
	Redirect      string
	AccessToken   string
	DisableLoader bool
	FromURL       string
}

func (rsc *Resource) SetErrorMessage(errorMessage string) {
	rsc.ErrorMessage = errorMessage
}

func (rsc *Resource) SetUserName(userName string) {
	rsc.UserName = userName
}

func (rsc *Resource) SetDomains(domains []string) {
	rsc.Domains = domains
}

func (rsc *Resource) IsSetCookies() bool {
	return len(rsc.Domains) > 0
}

func (rsc *Resource) SetRedirect(redirect string) {
	rsc.Redirect = redirect
}

func (rsc *Resource) SetAccessToken(accessToken string) {
	rsc.AccessToken = accessToken
}

func (rsc *Resource) SetDisableLoader(disableLoader bool) {
	rsc.DisableLoader = disableLoader
}

func (rsc *Resource) SetFromURL(fromUrl string) {
	rsc.FromURL = fromUrl
}
