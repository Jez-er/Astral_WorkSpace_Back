package google

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func Home(c *gin.Context) {
	tmpl, err := template.ParseFiles("backend/internal/google/index.html")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(404)
		return
	}
	err = tmpl.Execute(c.Writer, gin.H{})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func SignInWithProvider(c *gin.Context) {

	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func CallbackHandler(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	_, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/success")
}

func Success(c *gin.Context) {

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf(`
      <div style="
          background-color: #fff;
          padding: 40px;
          border-radius: 8px;
          box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
          text-align: center;
      ">
          <h1 style="
              color: #333;
              margin-bottom: 20px;
          ">You have Successfull signed in!</h1>
          
          </div>
      </div>
  `)))
}
