package controllers

import (
	"fmt"
	"net/http"
	"time"

	"elsenova/config"
	"elsenova/models"
	"elsenova/query"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm/clause"
)

type DiscordController struct {
	dg     *discordgo.Session
	dgOnce sync.Once
}

func (d *DiscordController) Mount(baseGroup gin.IRouter) {
	us := baseGroup.Group("discord")

	us.GET("/user/:id", d.userRegistrationMiddleware, d.GetUser)
}

func (d *DiscordController) GetUser(c ctx) {
	id := c.Param("id")
	cached, _ := query.CachedUser.Where(query.CachedUser.ID.Eq(id)).First()
	if cached != nil {
		// If we pulled this over 5 minutes ago, refresh in the background
		if cached.UpdatedAt.Before(time.Now().Add(time.Minute * -5)) {
			go d.fetchUser(id)
		}

		c.JSON(http.StatusOK, cached)
		return
	}

	user, err := d.fetchUser(id)
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

func (d *DiscordController) fetchUser(id string) (*models.CachedUser, error) {
	cu := query.CachedUser

	dg := d.session()
	user, err := dg.User(id)
	if err != nil {
		return nil, err
	}

	model := &models.CachedUser{}
	copier.Copy(model, &user)

	err = cu.Clauses(clause.OnConflict{UpdateAll: true}).Create(model)
	return model, err
}

// Ensures that we're not just looking up arbitrary people
// We only want to fetch data for users that are on the boards
func (d *DiscordController) userRegistrationMiddleware(c ctx) {
	v := query.Vore
	id := c.Param("id")

	if matches, _ := v.Select(v.UserID.Eq(id)).Count(); matches == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("User %v is not registered!", id),
		})
		return
	}

	c.Next()
}
