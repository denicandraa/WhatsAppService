package controllers

import (
	hdlr "../handlers"
	"../models"
	"github.com/gin-gonic/gin"
)

type GinCtrl struct {
}

func (u *GinCtrl) SendText(c *gin.Context) {

	response := models.Status{}
	param := models.ParamSendText{}

	err := c.ShouldBindJSON(&param)
	//data1:=models.DataSendText{
	//	Number: "85105765000",
	//	Message: "hehehe",
	//}
	//param.Data = append(param.Data, data1)

	//c.JSON(200,param)
	if err != nil {
		response.Status = false
		c.JSON(400, response)
		return
	}

	go hdlr.SendText(param)

	response.Status = true
	c.JSON(200, response)

}
func (u *GinCtrl) SendTextGroup(c *gin.Context) {

	response := models.Status{}
	param := models.ParamSendText{}

	err := c.ShouldBindJSON(&param)

	if err != nil {
		response.Status = false
		c.JSON(400, response)
		return
	}

	////data1:=models.DataSendText{
	////	Number: "85105765000",
	////	Message: "hehehe",
	////}
	////param.Data = append(param.Data, data1)
	//
	////c.JSON(200,param)
	//if err != nil {
	//	response.Status = false
	//	c.JSON(400, response)
	//	return
	//}

	//dataService := models.DataSendText{}
	//dataService.Number = "+6281224765832-1603640191" // mess
	//dataService.Number = "+6281224765832-1614752485"// simotor
	//dataService.Number = "+6282167899695-1553251552" // dna

	//sendMessageGroup := func() {
	//
	//	//for {
	//	//	//t := time.Now()
	//	//	//nextPantau := t.Add(5 * time.Hour)
	//	//
	//	//	//dataService.Message = "Daily Report Survey Kesehatan Head Office PT Digital Nayaka Abhinanya Senin, 22 November 2021\n\nSurvey terisi 100% \n57% WFO  total 8 Orang\n43% WFH  total 6 Orang\n0% Libur/Tgl Merah\n0% Tidak mengisi survey \n\nKesehatan karyawan\n100% karyawan dalam keadaan sehat.\n\nTerima kasih"
	//	//
	//	//	param.Data = append(param.Data, dataService)
	//	//
	//	//	go hdlr.SendTextGroup(param)
	//	//
	//	//	param = models.ParamSendText{}
	//	//
	//	//	time.Sleep(5 * time.Hour)
	//	//
	//	//}
	//
	//	go hdlr.SendTextGroup(param)
	//
	//	param = models.ParamSendText{}
	//}
	//
	//go sendMessageGroup()

	go hdlr.SendTextGroup(param)

	response.Status = true
	c.JSON(200, response)

}

//func (u *GinCtrl) ResetSession(c *gin.Context) {
//	err := hdlr.ResetSession()
//
//	if err != nil {
//		c.JSON(400, gin.H{
//			"message": err.Error(),
//		})
//		return
//	} else {
//		c.JSON(200, gin.H{
//			"message": "sukses",
//		})
//		return
//	}
//}
//
//func InitWhatsapp()  {
//
//	hdlr.InitWhatsapp()
//}
