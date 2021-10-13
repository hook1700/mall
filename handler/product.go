package handler

import (
	"github.com/gin-gonic/gin"
	"mall/enum"
	"mall/model"
	"mall/query"
	"mall/resp"
	"mall/service"
	"net/http"
)

type ProductHandler struct {
	ProductSrv service.ProductSrv

}

func (h *ProductHandler) GetEntity (result model.Product) resp.Product{
	return resp.Product{
		Id:                   result.ProductID,
		Key:                  result.ProductID,
		ProductId:            result.ProductID,
		ProductName:          result.ProductName,
		ProductIntro:         result.ProductIntro,
		CategoryId:           result.CategoryID,
		ProductCoverImg:      result.ProductCoverImg,
		ProductBanner:        result.ProductBanner,
		OriginalPrice:        result.OriginalPrice,
		SellingPrice:         result.SellingPrice,
		StockNum:             result.StockNum,
		Tag:                  result.Tag,
		SellStatus:           result.SellStatus,
		ProductDetailContent: result.ProductDetailContent,
		IsDeleted:            result.IsDeleted,
	}
}

func (h *ProductHandler) ProductInfoHandler (c *gin.Context) {
	entity := resp.Entity{
		Code: int(enum.OperateFail),
		Msg: enum.OperateFail.String(),
		Total: 0,
		TotalPage: 1,
		Data: nil,
	}

	productId := c.Param("id")
	if productId == ""{
		c.JSON(http.StatusInternalServerError,gin.H{"entity":entity})
		return
	}
	p := model.Product{ProductID: productId}
	result,err := h.ProductSrv.Get(p)
	if err !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"entity":entity})
		return
	}
	r := h.GetEntity(*result)
	entity = resp.Entity{
		Code: int(enum.OperateOk),
		Msg: "ok",
		Total: 0,
		TotalPage: 1,
		Data: r,
	}
	c.JSON(http.StatusOK,gin.H{"entity":entity})


}

func (h *ProductHandler) ProductListHandler (c *gin.Context) {
	var q query.ListQuery
	entity := resp.Entity{
		Code: int(enum.OperateFail),
		Msg: enum.OperateFail.String(),
		Total: 0,
		TotalPage: 1,
		Data: nil,
	}
	err := c.ShouldBindQuery(&q)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"entity":entity})
		return
	}
	list,err := h.ProductSrv.List(&q)
	total,err := h.ProductSrv.GetTotal(&q)

	if err!= nil{panic(err)}
	if q.PageSize == 0{
		q.PageSize = 5
	}
	ret := total % q.PageSize
	ret2 := total / q.PageSize
	totalPage := 0
	if ret == 0 {
		totalPage = ret2
	}else {
		totalPage = ret2 +1
	}
	var newList []*resp.Product
	for _,item := range list{
		r := h.GetEntity(*item)
		newList = append(newList,&r)
	}
	entity = resp.Entity{
		Code: http.StatusOK,
		Msg: "ok",
		Total: total,
		TotalPage: totalPage,
		Data: newList,
	}
	c.JSON(http.StatusOK,gin.H{"entity":entity})
}

func (h *ProductHandler) AddProductHandler (c *gin.Context)  {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	p := model.Product{}
	err := c.ShouldBindJSON(&p)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"entity":entity})
		return
	}
	result,err := h.ProductSrv.Add(p)
	if err != nil{
		entity.Msg = err.Error()
		return
	}
	if result.ProductID == ""{
		c.JSON(http.StatusOK,gin.H{"entity":entity})
		return
	}
	entity.Code = int(enum.OperateOk)
	entity.Msg = enum.OperateOk.String()
	c.JSON(http.StatusOK,gin.H{"entity":entity})
}

func (h *ProductHandler) EditProductHandler (c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	p := model.Product{}
	//绑定参数
	err := c.ShouldBindJSON(&p)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{"entity":entity})
	}
	//编辑
	b,err := h.ProductSrv.Edit(p)
	if err !=nil{
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b{
		entity.Code = int(enum.OperateOk)
		entity.Msg = enum.OperateOk.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}

func (h *ProductHandler) DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")
	b,err := h.ProductSrv.Delete(id)
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	if err!= nil{
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b{
		entity.Code = int(enum.OperateOk)
		entity.Msg = enum.OperateOk.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}
