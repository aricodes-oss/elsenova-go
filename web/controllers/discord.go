package controllers

import (
	"fmt"
	"net/http"
	"slices"
	"sync"

	"elsenova/config"
	"elsenova/models"
	"elsenova/query"
	"elsenova/util"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

type DiscordController struct {
	dg     *discordgo.Session
	dgOnce sync.Once

	userCache map[string]*discordgo.User
}

func (d *DiscordController) Mount(baseGroup gin.IRouter) {
	d.userCache = make(map[string]*discordgo.User)
	us := baseGroup.Group("discord")

	us.GET("/user/:id", d.userRegistrationMiddleware, d.GetUser)
}

func (d *DiscordController) GetUser(c ctx) {
	id := c.Param("id")
	user, inCache := d.userCache[id]
	if inCache {
		c.JSON(http.StatusOK, user)
		return
	}

	dg := d.session()
	user, err := dg.User(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

// -- Internal --

func (d *DiscordController) session() *discordgo.Session {
	d.dgOnce.Do(func() {
		conf := config.Load()
		sess, _ := discordgo.New("Bot " + conf.Token)
		sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
		sess.Open()

		d.dg = sess
	})

	return d.dg
}

// Ensures that we're not just looking up arbitrary people
// We only want to fetch data for users that are on the boards
func (d *DiscordController) userRegistrationMiddleware(c ctx) {
	v := query.Vore
	id := c.Param("id")

	knownUsers, _ := v.Select(v.UserID).Find()
	ids := util.Map(knownUsers, func(record *models.Vore, idx int) string {
		return record.UserID
	})

	if !slices.Contains(ids, id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("User %v is not registered!", id),
		})
		return
	}
	c.Next()
}
