package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/role", Get)

	router.GET("/role/:id", GetOne)

	router.POST("/role", Post)

	router.PUT("/role/:id", Put)

	router.DELETE("/role/:id", Delete)

	router.Run(":8080")
}

type Response struct {
	StatusCode int
	Msg        interface{}
}

// 取得全部資料
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, Response{http.StatusOK, Data})
}

// 取得單一筆資料
func GetOne(c *gin.Context) {
	intID, _ := strconv.Atoi(c.Param("id"))
	key, ok := exsistID(Data, intID)
	if ok {
		c.JSON(http.StatusOK, Response{http.StatusOK, Data[key]})
	} else {
		err := errors.New("沒有該筆資料")
		log.Printf("err: %#+v\n", err)
		c.JSON(http.StatusForbidden, Response{http.StatusForbidden, "沒有該筆資料"})
	}
}

// 新增資料
func Post(c *gin.Context) {
	decoder := json.NewDecoder(c.Request.Body)
	var addRole Role
	err := decoder.Decode(&addRole)
	if err != nil {
		panic(err)
	}
	addRole.ID = Data[len(Data)-1].ID + 1
	Data = append(Data, addRole)
	c.JSON(http.StatusOK, Response{http.StatusOK, addRole})
}

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

// 更新資料, 更新角色名稱與介紹
func Put(c *gin.Context) {
	updateID, _ := strconv.Atoi(c.Param("id"))

	decoder := json.NewDecoder(c.Request.Body)
	var modfyRole Role
	err := decoder.Decode(&modfyRole)
	if err != nil {
		panic(err)
	}

	key, ok := exsistID(Data, updateID)
	if ok {
		Data[key].Name = modfyRole.Name
		Data[key].Summary = modfyRole.Summary
		c.JSON(http.StatusOK, Response{http.StatusOK, Data[key]})
	} else {
		err := errors.New("沒有該筆資料")
		log.Printf("err: %#+v\n", err)
		c.JSON(http.StatusForbidden, Response{http.StatusForbidden, "沒有該筆資料"})
	}
}

// 刪除資料
func Delete(c *gin.Context) {
	delID, _ := strconv.Atoi(c.Param("id"))

	key, ok := exsistID(Data, delID)
	if ok {
		Data = append(Data[:key], Data[key+1:]...)
		c.JSON(http.StatusOK, Response{http.StatusOK, Data})
	} else {
		err := errors.New("沒有該筆資料")
		log.Printf("err: %#+v\n", err)
		c.JSON(http.StatusForbidden, Response{http.StatusForbidden, "沒有該筆資料"})
	}
}

func exsistID(slice []Role, id int) (key int, bl bool) {
	for key, value := range slice {
		if value.ID == uint(id) {
			return key, true
		}
	}
	return -1, false
}
