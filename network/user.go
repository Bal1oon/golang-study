package network

import (
	"fmt"
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
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
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
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get입니다")

	// u.router.okResponse(c, &types.APIResponse{

	// 	Result:      1,
	// 	Description: "성공입니다.",
	// })

	u.router.okResponse(c, &types.UserResponse{
		APIResponse: &types.APIResponse{
			Result:      1,
			Description: "성공입니다.",
		},
		User: nil,
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update입니다")
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete입니다")
}
