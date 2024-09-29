package router

import "github.com/gin-gonic/gin"

func (r *Router) GET(path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.engine.GET(path, handlers...)
}

func (r *Router) POST(path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.engine.POST(path, handlers...)
}

func (r *Router) PUT(path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.engine.PUT(path, handlers...)
}

func (r *Router) DELETE(path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.engine.DELETE(path, handlers...)
}