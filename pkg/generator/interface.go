package generator

type Server interface {
	Run(port string) error
}
