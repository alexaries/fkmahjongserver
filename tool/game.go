package tool

import (
	"github.com/qipaiyouxi/fkmahjongserver/cache"
	"github.com/qipaiyouxi/fkmahjongserver/db"
	"github.com/qipaiyouxi/fkmahjongserver/db/dao"
	"github.com/qipaiyouxi/fkmahjongserver/def"
	"github.com/qipaiyouxi/fkmahjongserver/notice"

	log "github.com/Sirupsen/logrus"
)

func handleGamesManage(games []*dao.Game) error {
	for _, game := range games {
		daoGame := cache.GetGMTGame(game.IndexID)
		if daoGame == nil {
			createGame(game)
			continue
		}

		game.SetExist(true)
		updateGame(game)
	}

	notice.ToolInitGame()
	return nil
}

func createGame(game *dao.Game) {
	err := game.Insert(db.Pool)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error(def.ErrInsertGame)
	}
}

func updateGame(game *dao.Game) {
	err := game.Update(db.Pool)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error(def.ErrUpdateGame)
	}
}

func deleteGame(game *dao.Game) {
	err := game.Delete(db.Pool)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error(def.ErrDeleteGame)
	}
}
