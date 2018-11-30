package migrate

import (
	"{{projectName}}/app"
	"{{projectName}}/model"
)

func CreateTable() {
	app.DB.AutoMigrate(
        //!!do not delete the line, gen generate code at here
	)
}
