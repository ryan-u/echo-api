package controller
import (
    "fmt"
    "log"
    "net/http"
  
    "github.com/labstack/echo/v4"
)

// company is a structure that contains the stock ticker from the client's HTTP request
type Dog struct {
    Name string `json:"name" form:"name" query:"name"`
    Type string `json:"type" from:"type" query:"type"`
}
    
// GrabPrice - handler method for binding JSON body and scraping for stock price
func GetDogs(c echo.Context) (err error) {
    dogName := c.QueryParam("name")
    dogType := c.QueryParam("type")
    dataType := c.Param("data")
    
    if dataType == "string" {
        return c.String(http.StatusOK, fmt.Sprintf("Your dog's name is: %s\nYour dog's type is: %s\n", dogName, dogType))
    } else if dataType == "json" {
        return c.JSON(http.StatusOK, map[string]string{
        "name": dogName,
        "type": dogType,})
    } else {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Please specify the data as String or JSON",})
    }
    
    return c.String(http.StatusOK, fmt.Sprintf("Your dog's name is: %s\nYour dog's type is: %s\n", dogName, dogType))
}

func AddDog(c echo.Context) (err error) {
    
    d := new(Dog)
    if err := c.Bind(d); err != nil {
        return err
    }
    
    log.Printf("This is your dog %#v", d)
    return c.JSON(http.StatusOK, d)
    
    

}