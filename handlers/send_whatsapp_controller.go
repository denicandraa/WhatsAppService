package handlers

import (
	"../middleware/def"
	"../models"
	"fmt"
	"github.com/gogo/protobuf/proto"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	waTypes "go.mau.fi/whatsmeow/types"
	"strings"
)

type WaGorout struct {
}

func SendText(param models.ParamSendText){

	defer def.DefFoo()

	for i := range param.Data {

		//6285105765000@s.whatsapp.net

		//go func() {
		//
		//}()

		param.Data[i].Number = strings.ReplaceAll(param.Data[i].Number, "+", "")

		targetJID := waTypes.NewJID(param.Data[i].Number,"s.whatsapp.net")

		_ , errx := client.SendMessage(targetJID, "", &waProto.Message{
			Conversation: proto.String(param.Data[i].Message),
		})

		if errx != nil{
			fmt.Println("send message errx. ",errx.Error())
		}else{
			fmt.Println("sukses send message")
		}

	}

}
func SendTextGroup(param models.ParamSendText){

	defer def.DefFoo()

	for i := range param.Data {

		//6285105765000@s.whatsapp.net

		//go func() {
		//
		//}()

		param.Data[i].Number = strings.ReplaceAll(param.Data[i].Number, "+", "")

		targetJID := waTypes.NewJID(param.Data[i].Number,"g.us")

		_ , errx := client.SendMessage(targetJID, "", &waProto.Message{
			Conversation: proto.String(param.Data[i].Message),
		})

		if errx != nil{
			fmt.Println("send message errx. ",errx.Error())
		}else{
			fmt.Println("sukses send message")
		}

	}

}

