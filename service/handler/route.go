package handler

func (ox gatewayHandler) initRouteUsage() {
	router := ox.service.Group("/user")
	{
		router.POST("/create", ox.UserCreateRequest)
		router.PUT("/update", ox.UserUpdateRequest)
		router.GET("/list", ox.UserListRequest)
		router.GET("/view", ox.UserViewRequest)
		router.DELETE("/delete", ox.UserDeleteRequest)
		router.GET("/login", ox.UserLoginRequest)
	}

	router = ox.service.Group("/host")
	{
		router.POST("/create", ox.HostCreateRequest)
		router.PUT("/update", ox.HostUpdateRequest)
		router.GET("/list", ox.HostListRequest)
		router.GET("/view", ox.HostViewRequest)
		router.DELETE("/delete", ox.HostDeleteRequest)
	}
}
