package article

import (
	"api-go-vue-admin/mysqli"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type Article struct {
	Id        int            `json:"id" form:"id"`
	Title     string         `json:"title" form:"title"`
	Img       string         `json:"img" form:"img"`
	Cid       int            `json:"cid" form:"cid"`
	Aid       int            `json:"aid" form:"aid"`
	Content   string         `json:"content" form:"content"`
	Status    int            `json:"status" form:"status"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" form:"deleted_at" gorm:"column:deleted_at"`
}

type Category struct {
	Id   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Pid  int    `form:"pid" json:"pid"`
}

type TreeList struct {
	Id   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	Pid  int    `form:"pid" json:"pid"`
	//子类下面可能也有子类，所以要用CategoryList的指针
	//Children []*TreeList  `json:"children"`
}

//绑定表名
func (a Article) TableName() string {
	return "article"
}

//获取栏目列表
func getCategory(c *gin.Context) {

	categoryList := getCategoryTree(0)
	//getCategoryTree(0)
	c.JSON(200, gin.H{
		"categoryList": categoryList,
		//"categoryList":"success",
	})

}

func getCategoryTree(pid int) []*TreeList {

	db := mysqli.GormConnect()

	//存放从数据库取出的所有栏目
	var category []Category

	//得到所有列表结果
	db.Table("category").Find(&category)

	var treeSlice []*TreeList

	for _, v := range category {

		if v.Pid != 0 {
			v.Name = "|-" + v.Name
		}

		node := &TreeList{

			Id:   v.Id,
			Name: v.Name,
			Pid:  v.Pid,
		}

		treeSlice = append(treeSlice, node)

	}

	return treeSlice

}

//文章删除 软删除
func delete(c *gin.Context) {

	id := c.Query("id")

	article := Article{}

	db := mysqli.GormConnect()

	db.Where("id = ?", id).Take(&article)

	db.Delete(&article)

	c.JSON(200, "success")

}

//文章新增
func add(c *gin.Context) {

	title := c.Query("title")
	img := c.Query("img")
	cid, _ := strconv.Atoi(c.Query("cid"))
	content := c.Query("content")
	aid, _ := strconv.Atoi(c.Query("aid"))

	fmt.Println(c.Query("aid"))

	article := Article{Title: title, Img: img, Cid: cid, Aid: aid, Content: content}

	db := mysqli.GormConnect()
	tx := db.Table("article")
	res := tx.Create(&article)

	if res.RowsAffected <= 0 {

		c.JSON(201, res.Error)

	}

	c.JSON(200, "success")

}

//文章更新
func update(c *gin.Context) {

	id := c.Query("id")
	title := c.Query("title")
	img := c.Query("img")
	content := c.Query("content")
	cid, _ := strconv.Atoi(c.Query("cid"))
	aid, _ := strconv.Atoi(c.Query("aid"))

	//通过结构体变量设置更新字段
	article := Article{
		Title:   title,
		Img:     img,
		Content: content,
		Cid:     cid,
		Aid:     aid,
	}

	db := mysqli.GormConnect()

	re := db.Model(&Article{}).Where("id = ?", id).Updates(&article)

	if re.Error != nil {

		fmt.Println(re.Error)

	}

	c.JSON(200, "success")
}

//通过id获取文章详情
func getArticleById(c *gin.Context) {

	id := c.Query("id")

	db := mysqli.GormConnect()

	article := Article{}

	res := db.Where("id = ?", id).Take(&article)

	if res.Error != nil {

		fmt.Println(res.Error)

	}

	c.JSON(200, gin.H{
		"articleList": article,
	})

}

//文章列表
func articleList(c *gin.Context) {

	db := mysqli.GormConnect()

	var article []Article

	//第几页
	page := c.DefaultQuery("page", "1")
	//每页显示多少
	pageSize := c.DefaultQuery("pageSize", "10")

	//字符串转成数字类型
	page1, _ := strconv.Atoi(page)
	pageSize1, _ := strconv.Atoi(pageSize)

	offsetNum := (page1 - 1) * pageSize1

	tx := db.Offset(offsetNum).Limit(pageSize1).Order("id DESC").Find(&article)

	var total int64 = 0

	db.Model(Article{}).Count(&total)

	//total := tx.RowsAffected

	if tx.Error != nil {
		fmt.Println("查询所有文章列表失败", tx.Error)
	}

	c.JSON(200, gin.H{
		"articleList": article,
		"total":       total,
	})

}
