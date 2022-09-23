package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong"
		})
	})

	Connect()
	r.Run()
}


func Connect(){
	// endereço do banco de dados slqserver
	dsn:= "sqlserver://spisdb001:MSRC1012@10.59.162.248:1433?databse=PNDB001"
	db, err:= gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database.gorm ERROR:", err)
	} else {
		fmt.Println("database.gorm sucessfull", db)
	}

	rows, err:= db.Raw("select des_usr, tip_usr from TBPN003 where cod_usr = 'operacional'").Rows()
	defer rows.Close()
	var nome string
	var tipo string
	for rows.Next(){
		rows.Scan(&nome, &tipo)
		fmt.Println("usuario:", nome)
		fmt.Println("tipo:", tipo)
	}
}