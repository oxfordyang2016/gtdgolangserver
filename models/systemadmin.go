package models
import(
  "fmt"
  "os/exec"
  "os"
  //"net/http"
	//"strconv"
//"github.com/jinzhu/gorm"
//"github.com/gin-contrib/sessions"
"github.com/gin-gonic/gin"
)








func Blogupdate (c *gin.Context){

var (
		cmdOut []byte
		err    error
	)
	cmdName := "git"
	cmdArgs := []string{"-C","/root/yangming/blog/source","pull","origin","master"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command: ", err)
		os.Exit(1)
	}
	sha := string(cmdOut)
	firstSix := sha[:6]
	fmt.Println("The first six chars of the SHA at HEAD in this repo are", firstSix)

  c.JSON(200, gin.H{
                        "status": "blog had updated",
                })

}




