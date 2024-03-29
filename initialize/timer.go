package initialize

import (
	"fmt"

	"server-fiber/config"
	"server-fiber/global"
	"server-fiber/utils"
)

func Timer() {
	if global.CONFIG.Timer.Start {
		for i := range global.CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				global.Timer.AddTaskByFunc("ClearDB", global.CONFIG.Timer.Spec, func() {
					err := utils.UpdateTable(global.DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
			}(global.CONFIG.Timer.Detail[i])
		}
	}
}
