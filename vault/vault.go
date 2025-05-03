package vault

import "learngo/models"

type Vault map[string]models.Entry

func NewVault() Vault {
	return make(Vault)
}
func (v Vault) AddEntry(e models.Entry) {
	v[e.Site] = e
}
func (v Vault) GetEntry(site string) (models.Entry, bool) {
	e, found := v[site]
	return e, found
}
func (v Vault) DeleteEntry(site string) bool {
	_, exists := v[site]
	if exists {
		delete(v, site)
		return true
	}
	return false
}

func (v Vault) ListSites() []string {
	sites := make([]string, 0, len(v))
	for site := range v {
		sites = append(sites, site)
	}
	return sites
}
