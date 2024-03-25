package network

import (
	"golang-study/service"
	"golang-study/types"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	//service
	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

		router.registerGET("/", userRouterInstance.get)
		router.registerPOST("/", userRouterInstance.create)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)

		// 코드 가독성을 위해 위와 같이 바꿈 (network/root.go 에 함수 선언)
		// userRouterInstance.router.engin.POST("/", userRouterInstance.create)
		// userRouterInstance.router.engin.PUT("/", userRouterInstance.update)
		// userRouterInstance.router.engin.DELETE("/", userRouterInstance.delete)
		// userRouterInstance.router.engin.GET("/", userRouterInstance.get)
	})

	return userRouterInstance
}

// CRUD
func (u *userRouter) create(c *gin.Context) {
	var req types.CreateRequest

	// request에 대한 검증
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류 입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("Create 에러 입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil), // error가 없기 때문에 errCode = nil
		})
	}
}

func (u *userRouter) get(c *gin.Context) {
	// u.router.okResponse(c, &types.APIResponse{
	// 	Result:      1,
	// 	Description: "성공입니다.",
	// })

	// u.router.okResponse(c, &types.UserResponse{
	// 	ApiResponse: &types.ApiResponse{
	// 		Result:      1,
	// 		Description: "성공입니다.",
	// 	},
	// 	User: nil,
	// })

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		Users:       u.userService.Get(),
	})

}

func (u *userRouter) update(c *gin.Context) {
	var req types.UpdateUserRequest

	// request에 대한 검증
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류 입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Update(req.Name, req.UpdatedAge); err != nil {
		u.router.failedResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("Update 에러 입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteRequest

	// request에 대한 검증
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류 입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Delete(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("Delete 에러 입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil), // error가 없기 때문에 errCode = nil
		})
	}
}
