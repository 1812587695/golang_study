package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hytx_manager/middleware/auth/jwt"
	"hytx_manager/models"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Auth interface {
	Check(c *gin.Context) bool
	User(c *gin.Context) interface{}
	Login(r *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{}
	Logout(r *http.Request, w http.ResponseWriter) bool
}

/**
	登录检测
 */
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jwtAuth Auth
		jwtAuth = jwt.NewJwtAuth()
		if !jwtAuth.Check(c) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "请登录",
			})
			c.Abort()
			return
		}
		if User(c).IsEnabled == 2 {
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  "该账号禁止登录,请联系管理员!",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

/**
	权限检测
 */
func PermissionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := User(c)
		if user.ID == 1 {
			Log(c)
			c.Next()
			return
		}
		//查询用户权限
		url := c.Request.URL.Path
		method := c.Request.Method
		permission := models.GetPermissionByUser(user, url, method)
		if permission != nil && permission.FrontUrl == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  "你没有该操作的权限,请联系管理员",
			})
			c.Abort()
		}
		//记录日志
		Log(c)
		c.Next()
	}
}

func RegisterGlobalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("auth", jwt.NewJwtAuth())
		c.Next()
	}
}

func User(c *gin.Context) *models.AdminDetailsResult {
	u, isExist := c.Get("user")
	if isExist {
		return u.(*models.AdminDetailsResult)
	}
	user := jwt.NewJwtAuth().User(c).(map[string]interface{})
	adminId := int(user["id"].(float64))
	admin, _ := models.GetAdminById(adminId)
	c.Set("user", admin)
	return admin
}

func Log(c *gin.Context) {
	ua := c.GetHeader("User-Agent")

	var equipment string
	var os string
	var osVer string
	//windows
	if strings.ContainsAny(ua, "windows") {
		equipment = "PC"
		os = "windows"
		if r, _ := regexp.MatchString(`(?i)nt 6.0`, ua); r {
			osVer = "Vista"
		} else if r, _ := regexp.MatchString(`(?i)nt 10.0`, ua); r {
			osVer = "10"
		} else if r, _ := regexp.MatchString(`(?i)nt 6.3`, ua); r {
			osVer = "8.1"
		} else if r, _ := regexp.MatchString(`(?i)nt 6.2`, ua); r {
			osVer = "8"
		} else if r, _ := regexp.MatchString(`(?i)nt 6.1`, ua); r {
			osVer = "7"
		} else if r, _ := regexp.MatchString(`(?i)nt 5.1`, ua); r {
			osVer = "XP"
		} else if r, _ := regexp.MatchString(`(?i)nt 5`, ua); r {
			osVer = "2000"
		} else if r, _ := regexp.MatchString(`(?i)nt 98`, ua); r {
			osVer = "98"
		} else if r, _ := regexp.MatchString(`(?i)nt`, ua); r {
			osVer = "nt"
		} else {
			osVer = "Unknown"
		}

		if r, _ := regexp.MatchString(`(?i)X64`, ua); r {
			os += " (X64)"
		} else if r, _ := regexp.MatchString(`(?i)X32`, ua); r {
			os += " (X32)"
		}
		//Linux
	} else if strings.ContainsAny(ua, "linux") {
		if strings.ContainsAny(ua, "android") {
			equipment = "Mobile Phone"
			res := regexp.MustCompile(`(?i)android\s([\d\.]+)`).FindAllStringSubmatch(ua, -1)
			os = "android"
			osVer = res[0][1]
		} else {
			os = "Linux"
		}
		//Unix
	} else if strings.ContainsAny(ua, "unix") {
		os = "Unix"
	} else if r, _ := regexp.MatchString(`(?i)iPhone|iPad|iPod`, ua); r {
		os = "IOS"
		res := regexp.MustCompile(`(?i)OS\s([0-9_\.]+)`).FindAllStringSubmatch(ua, -1)
		osVer = strings.Replace(res[0][1], "_", ".", -1)
		if strings.ContainsAny(ua, "iPhone") {
			equipment = "iPhone"
		}
		if strings.ContainsAny(ua, "iPad") {
			equipment = "iPad"
		}
		if strings.ContainsAny(ua, "iPod") {
			equipment = "iPod"
		}
		//mac os
	}else if strings.ContainsAny(ua, "mac os") {
		res := regexp.MustCompile(`(?i)Mac OS X\s([0-9_\.])`).FindAllStringSubmatch(ua, -1)
		equipment = "PC"
		os = "Mac OS X"
		osVer = res[0][1]
	}else{
		os = "Unknown"
	}
	ip := c.GetHeader("X-Real-IP")
	url := c.Request.URL.Path
	uri := c.Request.RequestURI
	region := "unknown"
	user := User(c)
	var name string
	models.DB.Table("permissions").Select("display_name").Where("api_url=?", url).Row().Scan(&name)
	//归属地
	resp, err := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip)
	defer resp.Body.Close()
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			var r struct{
				Code int
				Data struct{
					IP string `json:"ip"`
					Country string `json:"country"`
					Area string `json:"area"`
					Region string `json:"region"`
					City string `json:"city"`
					County string `json:"county"`
					Isp string `json:"isp"`
					CountryId string `json:"country_id"`
					AreaId string `json:"area_id"`
					RegionId string `json:"region_id"`
					CityId string `json:"city_id"`
					CountyId string `json:"county_id"`
					IspId string `json:"isp_id"`
				}
			}
			json.Unmarshal(body, &r)
			if r.Code == 0 {
				region = r.Data.Region + " " + r.Data.Isp
			}
		}
	}
	models.DB.Table("operation_log").Create(&models.OperationLog{
		AdminId : user.ID,
		Ip : ip,
		Device : equipment + " " + os + " " + osVer,
		Url : url,
		Uri : uri,
		Name : name,
		Region : region,
		CreatedAt : time.Now(),
	})
}
