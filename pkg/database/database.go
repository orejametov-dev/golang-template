package database

// DBConnection интерфейс абстракции для базы данных
type DBConnection interface {
	Connect() error
	Query(query string) ([]string, error)
	// Другие методы, которые вы хотите использовать
}
