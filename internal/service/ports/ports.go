package ports

type Database interface {
	LoadSecret(service string) (string, error)
	SaveSecret(service string, secret string) error
}
