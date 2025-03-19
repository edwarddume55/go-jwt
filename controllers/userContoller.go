package controllers

import(
	"context"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	helper "go-jwt-project/helpers"
	"go-jwt-project/models"
	"go-jwt-project/database"
	"golang.org/x/crypto/decrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user" ) 
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser()