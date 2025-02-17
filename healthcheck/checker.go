package healthcheck

type Checker interface {
	Name() string
	Check() map[string]interface{}
}
