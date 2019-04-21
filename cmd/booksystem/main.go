package main

import (
	"booksystem/pkg/api"
	"booksystem/pkg/database"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowMethods = append(config.AllowMethods, "DELETE")
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type"}
	router.Use(cors.New(config))
	api.RunHTTPServer(router)
	engine, err := database.GetConnection()
	defer engine.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = database.InitTable(engine)

	if !Exists("./public_lock") {

		fileInfo, err := os.Open("./publisher.txt")
		defer fileInfo.Close()
		if err != nil {
			log.Fatal(err)
		}
		br := bufio.NewReader(fileInfo)

		for {
			line, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			content := strings.Split(string(line), ",")
			pb := database.Publisher{}
			pb.Name = content[1]
			pb.ISBN = content[0]
			pb.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
			_, err = engine.Insert(pb)
			if err != nil {
				log.Print(err)
			}
		}
		_, err = os.Create("./public_lock")
		if err != nil {
			log.Fatal(err)
		}

	}

	if err != nil {
		fmt.Println(err)
		return
	}
	router.Run()
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
