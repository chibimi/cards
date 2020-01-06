package main

import (
	"fmt"
	"net/http"

	"github.com/chibimi/cards/card"
	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	cfg := struct {
		Login       string `envconfig:"db_login"`
		Password    string `envconfig:"db_password"`
		Host        string `envconfig:"db_host"`
		DB          string `envconfig:"db"`
		Port        int    `envconfig:"port" default:"4203"`
		EditorFront string `envconfig:"editor_front"`
	}{}
	envconfig.Process("card_api", &cfg)

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", cfg.Login, cfg.Password, cfg.Host, cfg.DB))
	if err != nil {
		log15.Crit("Unable to access db", "err", err.Error())
	}
	defer db.Close()

	// s := api.NewService(card.NewService(db, log15.New()))
	ss := card.NewSService(db)

	router := httprouter.New()
	router.POST("/ref", ss.Ref.CreateRef)
	router.GET("/ref", ss.Ref.ListRef)
	router.GET("/status/ref", ss.Ref.ListRefByStatus)
	router.GET("/ref/:id", ss.Ref.GetRef)
	router.PUT("/ref/:id", ss.Ref.SaveRef)
	router.GET("/ref/:id/spell", ss.Spell.ListByRefEndpoint)
	router.PUT("/ref/:id/spell/:spell_id", ss.Spell.AddSpellRefEndpoint)
	router.DELETE("/ref/:id/spell/:spell_id", ss.Spell.DeleteSpellRefEndpoint)
	router.GET("/ref/:id/ability", ss.Ability.ListByRefEndpoint)
	router.PUT("/ref/:id/ability/:ability_id", ss.Ability.AddAbilityRefEndpoint)
	router.DELETE("/ref/:id/ability/:ability_id", ss.Ability.DeleteAbilityRefEndpoint)
	router.PUT("/ref/:id/feat", ss.Feat.SaveFeat)
	router.GET("/ref/:id/feat", ss.Feat.GetFeat)
	router.GET("/ref/:id/model", ss.Model.ListEndpoint)

	router.POST("/model", ss.Model.CreateEndpoint)
	router.PUT("/model/:id", ss.Model.SaveEndpoint)
	router.DELETE("/model/:id", ss.Model.DeleteEndpoint)
	router.GET("/model/:id/ability", ss.Ability.ListByModelEndpoint)
	router.PUT("/model/:id/ability/:ability_id", ss.Ability.AddAbilityModelEndpoint)
	router.DELETE("/model/:id/ability/:ability_id", ss.Ability.DeleteAbilityModelEndpoint)
	router.GET("/model/:id/weapon", ss.Weapon.ListEndpoint)

	router.POST("/weapon", ss.Weapon.CreateEndpoint)
	router.PUT("/weapon/:id", ss.Weapon.SaveEndpoint)
	router.DELETE("/weapon/:id", ss.Weapon.DeleteEndpoint)
	router.GET("/weapon/:id/ability", ss.Ability.ListByWeaponEndpoint)
	router.PUT("/weapon/:id/ability/:ability_id", ss.Ability.AddAbilityWeaponEndpoint)
	router.DELETE("/weapon/:id/ability/:ability_id", ss.Ability.DeleteAbilityWeaponEndpoint)

	router.GET("/spells/:id", ss.Spell.GetEndpoint)
	router.GET("/spells", ss.Spell.ListEndpoint)
	router.POST("/spells", ss.Spell.CreateEndpoint)
	router.PUT("/spells/:id", ss.Spell.SaveEndpoint)

	router.GET("/abilities/:id", ss.Ability.GetEndpoint)
	router.GET("/abilities", ss.Ability.ListEndpoint)
	router.POST("/abilities", ss.Ability.CreateEndpoint)
	router.PUT("/abilities/:id", ss.Ability.SaveEndpoint)
	//OLD
	// router.GET("/abilities", s.ListAbilities)
	// router.GET("/abilities/:id", s.GetAbility)
	// router.POST("/abilities", s.CreateAbility)
	// router.PUT("/abilities/:id", s.UpdateAbility)
	// router.DELETE("/abilities/:id", s.DeleteAbility)

	// // router.GET("/spells", s.ListSpells)
	// // router.POST("/spells", s.CreateSpell)
	// // router.PUT("/spells/:id", s.UpdateSpell)

	// // router.GET("/feats", s.ListFeats)
	// // router.POST("/feats", s.CreateFeat)
	// // router.PUT("/feats/:id", s.UpdateFeat)

	// router.GET("/cards", s.ListCards)
	// // router.GET("/cards/:id", s.GetCard)
	// router.GET("/cards/:id/related", s.GetRelatedCards)
	// router.GET("/cards/:id/abilities", s.GetCardAbilities)
	// // router.GET("/cards/:id/models", s.GetCardModels)
	// // router.GET("/cards/:id/spells", s.GetCardSpells)
	// // router.GET("/cards/:id/feats", s.GetCardFeat)
	// // router.POST("/cards", s.CreateCard)
	// // router.DELETE("/cards/:id", s.DeleteCard)

	// // router.GET("/models/", s.ListModels)
	// // router.GET("/models/:id", s.GetModel)
	// router.GET("/models/:id/abilities", s.GetModelAbilities)
	// // router.GET("/models/:id/weapons", s.GetModelWeapons)
	// // router.POST("/models", s.CreateModel)
	// // router.PUT("/models/:id", s.UpdateModel)
	// // router.DELETE("/models/:id", s.DeleteModel)

	// // router.GET("/weapons/", s.ListWeapons)
	// // router.GET("/weapons/:id", s.GetWeapon)
	// router.GET("/weapons/:id/abilities", s.GetWeaponAbilities)
	// // router.POST("/weapons", s.CreateWeapon)
	// // router.PUT("/weapons/:id", s.UpdateWeapon)
	// // router.DELETE("/weapons/:id", s.DeleteWeapon)

	// // router.PUT("/cards/:id/spells/:spell_id", s.AddCardSpell)
	// // router.DELETE("/cards/:id/spells/:spell_id", s.DeleteCardSpell)

	// router.PUT("/cards/:id/abilities/:ability_id", s.AddCardAbility)
	// router.DELETE("/cards/:id/abilities/:ability_id", s.DeleteCardAbility)

	// router.PUT("/models/:id/abilities/:ability_id", s.AddModelAbility)
	// router.DELETE("/models/:id/abilities/:ability_id", s.DeleteModelAbility)

	// router.PUT("/weapons/:id/abilities/:ability_id", s.AddWeaponAbility)
	// router.DELETE("/weapons/:id/abilities/:ability_id", s.DeleteWeaponAbility)

	router.ServeFiles("/editor/*filepath", http.Dir(cfg.EditorFront))

	stack := negroni.New()
	stack.Use(cors.AllowAll())
	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())
	stack.UseHandler(router)

	log15.Info("Listening...", "port", cfg.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), stack); err != nil {
		log15.Crit("Unable to start server", "err", err.Error())
	}
}
