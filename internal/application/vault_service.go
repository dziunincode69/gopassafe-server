package application

type VaultService struct {
}

func (v *VaultService) GetVault(id uint32) string {
	return "0000000000-GET-VAULT-00000-00000"
}
func (v *VaultService) CreateVault(name, description, url, password string) string {
	return "0000000000-CREATE-VAULT-00000-00000"
}
func (v *VaultService) UpdateVault(id uint32, name, description, url, password string) string {
	return "0000000000-UPDATE-VAULT-00000-00000"
}
func (v *VaultService) DeleteVault(id uint32) string {
	return "0000000000-DELETE-VAULT-00000-00000"
}
