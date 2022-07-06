package keystorage

type Service interface {
	Get(key string) string
	Set(key, value string)
	Err() error
	SetTableName(string)
}

type KeyStorage struct {
	Key   string `gorm:"primarykey"`
	Value string
}
