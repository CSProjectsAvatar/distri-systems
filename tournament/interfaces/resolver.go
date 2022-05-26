package interfaces

// Responsible for Save and Retrieve Data
type resolver interface {
	SaveFiles(tour_name string, files map[string]string) error
	SaveFile(tour_name string, file_name string, file_content string) error
	GetFile(tour_name string, file_name string) string
}
