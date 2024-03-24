package network

import (
	"fmt"
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
	fmt.Println("create입니다")

	u.router.okResponse(c, &types.CreateUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get입니다")

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
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
		User:        nil,
	})

}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update입니다")

	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete입니다")

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
	})
}
