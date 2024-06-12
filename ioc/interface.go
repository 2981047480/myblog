package ioc

type Container interface {
	Get(name string) any
	Register(name string, obj Object)
	Init() error // 这里的init用来初始化Container
}

type Object interface {
	Init() error
} // 这里一开始做的时候没想到 为了后续方便加上init更方便 这里的init用来给container中的实例进行初始化
