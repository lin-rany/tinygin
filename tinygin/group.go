package tinygin

type Group struct {
	preStr      string
	engine      *Engine
	parent      *Group
	middlewares []HandlerFunc
}

func (g *Group) NewGroup(pre string) *Group {
	return &Group{
		preStr:      g.preStr + pre,
		engine:      g.engine,
		parent:      g,
		middlewares: make([]HandlerFunc, 0),
	}
}
func (g *Group) Use(methods ...HandlerFunc) {
	g.middlewares = append(g.middlewares, methods...)
}
func (g *Group) GET(path string, handler HandlerFunc) {
	g.AddRoute("GET", path, handler)
}
func (g *Group) POST(path string, handler HandlerFunc) {
	g.AddRoute("POST", path, handler)
}

func (g *Group) AddRoute(method, path string, handler HandlerFunc) {
	path = g.preStr + path
	g.engine.router.AddRoute(method, path, handler)
}
