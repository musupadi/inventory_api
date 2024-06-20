package Controller

import (
	"apiinventory/Config"
	"apiinventory/Model"
	"crypto/md5"
	"encoding/hex"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user Model.Logins
	c.Header("Content-Type", "application/json")
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menggunakan data username dan password dari permintaan POST
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Hash the password using MD5
	hashedPassword := md5.Sum([]byte(password))
	hashedPasswordStr := hex.EncodeToString(hashedPassword[:])

	// Query SQL dengan placeholder parameterized
	query := squirrel.Select(
		"a.id",
		"a.username",
		"a.name",
		"a.email",
		"a.jenis_kelamin",
		"a.kontak",
		"a.id_role",
		"a.photo",
		"a.id_shop",
		"b.level",
		"b.label as role_name",
		"c.label as shop_name",
	).From("user a").
		Join("role b ON a.id_role = b.id").
		Join("shop c ON a.id_shop = c.id").
		Where(squirrel.Eq{
			"a.username": username,
			"a.password": hashedPasswordStr,
		})
	// query := "SELECT a.username, a.name,a.email,a.jenis_kelamin,a.kontak,a.id_role,b.label as role_name,b.level,a.photo,a.id_shop,c.label as shop_name FROM user as a JOIN role as b ON a.id_role = b.id JOIN shop c ON a.id_shop = c.id WHERE username = ? AND password = ?"

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		// log.Printf("Error building SQL query: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to build SQL query"})
		return
	}

	// Execute the query
	rows, err := Config.DB.Query(sqlQuery, args...)
	if err != nil {
		// log.Printf("Database query error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return
	}
	defer rows.Close()

	var users []Model.UserData
	for rows.Next() {
		var userData Model.UserData
		if err := rows.Scan(
			&userData.Id,
			&userData.Username,
			&userData.Name,
			&userData.Email,
			&userData.Jenis_kelamin,
			&userData.Kontak,
			&userData.Id_role,
			&userData.Photo,
			&userData.Id_shop,
			&userData.Level,
			&userData.Role_name,
			&userData.Shop_name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, userData)
	}

	ResponseJSON(c, http.StatusOK, users)
}
