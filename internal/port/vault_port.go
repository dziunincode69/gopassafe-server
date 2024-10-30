package port

type VaultServicePort interface {
	GetVault(id uint32) string
	CreateVault(name, description, url, password string) string
	UpdateVault(id uint32, name, description, url, password string) string
	DeleteVault(id uint32) string
}
