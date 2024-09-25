/*
 * @Descripttion: 日志工具 - Log Utils
 */
package LogUtils

import (
	"log"
	"os"
)

func init() {
	fileName, err := os.OpenFile("./logger.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Logger(err.Error())
	}

	log.SetOutput(fileName)
}

func Logger(anything string) {
	log.Println(anything)
}
