package console

type Router struct {
	kernel *Kernel
}

func NewRouter(kernel *Kernel) Router {
	return Router{
		kernel: kernel,
	}
}

func (r Router) Exec(name string, args []string) {
	c := r.kernel.commands()[name]()
	c.Exec(args)
}
