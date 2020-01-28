package monitor

type Monitor interface {
	Init()

	Close()
}

func GetMonitor() Monitor {

}
