package controllers

import (
	"go_server/database"
	"go_server/models"

	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	// "github.com/golang-jwt/jwt/v4"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {

var data map[string]string

err := c.BodyParser(&data)
if err != nil{
return err
}

password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
user := models.User{
    Name: data["name"],
    Email: data["email"],
    Password: password,
}

database.DB.Create(&user)

return c.JSON(user)
    }


    func Login(c *fiber.Ctx) error{


        var data map[string]string
        if err := c.BodyParser(&data); err != nil{
        return err
        }

        var user models.User

        database.DB.Where("email=?", data["email"]).First(&user)
        if user.Id == 0 {
            c.Status(fiber.StatusNotFound)
            return c.JSON(fiber.Map{
                "message":"user not found",
            })
        }

        if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
            c.Status(fiber.StatusBadRequest)
            return c.JSON(fiber.Map{
                "message":"Icorrect password",
            })
        }
        return c.JSON(user)

        // claims := jwt.RegisteredClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
        //     Issuer:    strconv.Itoa(int(user.Id)),
        //     ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
        // })
        claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
            Issuer:    strconv.Itoa(int(user.Id)),
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        })
    
        token, err := claims.SignedString([]byte(SecretKey))
    
        if err != nil {
            c.Status(fiber.StatusInternalServerError)
            return c.JSON(fiber.Map{
                "message": "could not login",
            })
        }
    
        cookie := fiber.Cookie{
            Name:     "jwt",
            Value:    token,
            Expires:  time.Now().Add(time.Hour * 24),
            HTTPOnly: true,
        }
    
        c.Cookie(&cookie)
    
        return c.JSON(fiber.Map{
            "message": "success",
        })
    }
    