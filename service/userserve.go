package service

import (
	"fmt"
	"goweb/models"
	"goweb/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAddressDetails
// @Tags 地址信息
// @param address query string false "地址"
// @Success 200 {string} json{"code","account"}
// @Router /details [get]
func GetAddressDetails(c *gin.Context) {
	account := &models.Account{}
	address := c.Query("address")
	account = models.GetAddressDetails(address)
	c.HTML(http.StatusOK, "templates/table.html", gin.H{
		"address": account,
	})
}

// GetAddressDetails
// @Summary 新增用户
// @Tags 用户
// @param email query string false "邮箱"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /AddUser [get]
func CreateUser(c *gin.Context) {
	user := models.User{}
	email := c.Query("email")
	err := models.CheckUserExist(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "邮箱已经被注册",
		})
		return
	}
	password := c.Query("password")
	repassword := c.Query("repassword")
	rand.Seed(time.Now().Unix())
	s := rand.Intn(1000000000)
	salt := strconv.Itoa(s)
	if password != repassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	user.Password = utils.EncryptoPassword(password, salt)
	//user.Password = password
	user.Salt = salt
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增用户成功",
	})
}

// @Summary 删除用户
// @Tags 用户
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /DeleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
	})
}

// @Summary 绑定钱包
// @Tags 用户
// @param email formData string false "邮箱"
// @param password formData string false "密码"
// @param address formData string false "地址"
// @Success 200 {string} json{"code","message"}
// @Router /UpdateUser [post]
func BindAddress(c *gin.Context) {
	user := models.User{}
	email := c.PostForm("email")
	password := c.PostForm("password")
	address := c.PostForm("address")
	user = models.FindUserByEmail(email)
	if user.Address != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "该账号已经绑定了其他钱包",
		})
		return
	}
	if user.Password == utils.EncryptoPassword(password, user.Salt) {
		models.BindAddress(user, address)
		c.JSON(http.StatusOK, gin.H{
			"message": "绑定地址成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "密码不正确",
		})
	}
	return
}

// @Summary 用户登录
// @Tags 用户
// @param email query string false "email"
// @param password query string false "password"
// @Success 200 {string} json{"code","message"}
// @Router /FindUserByEmailAndPwd [post]
func FindUserByEmailAndPwd(c *gin.Context) {
	data := models.User{}
	email := c.Query("email")
	password := c.Query("password")
	user := models.FindUserByEmail(email)
	if user.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "该用户不存在",
		})
		return
	}
	flag := utils.DecryptoPassword(password, user.Salt, user.Password)
	fmt.Println(user.Salt)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"message": "密码错误",
		})
		return
	}
	pwd := utils.EncryptoPassword(password, user.Salt)
	models.FindUserByEmailAndPwd(email, pwd)
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"data":    data,
	})
	return
}

// @Summary 查询钱包
// @Tags 用户
// @param address formData string false "地址"
// @Success 200 {string} json{"code","message"}
// @Router /BatchQuery [get]
func BatchQuery(c *gin.Context) {
	addresses := c.Query("address")
	accounts := models.BatchQuery(addresses)
	fmt.Println(accounts)
	c.JSON(http.StatusOK, gin.H{
		"Accounts": accounts,
	})
	return
}

// @Summary 查询钱包
// @Tags 用户
// @Param nonce query string false "Nonce值"
// @Param ethbalance query string false "以太坊余额"
// @Param usdcbalance query string false "USDC余额"
// @Param activeweeks query string false "活跃周数"
// @Param activedays query string false "活跃天数"
// @Param totalfee query string false "总手续费"
// @Param startday query string false "开始日期"
// @Param txvalue query string false "交易价值"
// @Success 200 {string} json{"code": 200, "message": "查询成功"}
// @Router /survey [get]
func Survey(c *gin.Context) {
	nonce := c.Query("nonce")
	ethbalance := c.Query("ethbalance")
	usdcbalance := c.Query("usdcbalance")
	activeweeks := c.Query("activeweeks")
	activedays := c.Query("activedays")
	totalfee := c.Query("totalfee")
	txvalue := c.Query("txvalue")
	zklite_nonce := c.Query("zklite_nonce")
	zklite_month := c.Query("zklite_month")
	zklite_week := c.Query("zklite_week")
	zklite_txvalue := c.Query("zklite_txvalue")
	zklite_eth := c.Query("zklite_eth")
	zklite_usdc := c.Query("zklite_usdc")
	res, countTotal := models.Survey(nonce, txvalue, totalfee, activedays, activeweeks, usdcbalance, ethbalance, zklite_nonce, zklite_month, zklite_week, zklite_txvalue, zklite_eth, zklite_usdc)
	fmt.Println(countTotal)
	c.JSON(http.StatusOK, gin.H{
		"code":       200,
		"message":    "success",
		"data":       res,
		"countTotal": countTotal,
	})
	return
}

// @Summary 更新
// @Tags 用户
// @param address formData string false "地址"
// @Success 200 {string} json{"code","message"}
// @update /update [get]
func Update(c *gin.Context) {
	addresses := c.Query("address")
	accountList, err := models.Update(addresses)
	if err != nil {
		fmt.Println("models.Update", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Accounts": accountList,
	})
	return
}
