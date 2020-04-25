package main

import (
	
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var db *gorm.DB


type Libro struct {
	ID    	int
	Nombre  	string gorm:"type:varchar(100);"
	Desc 		string gorm:"type:varchar(450);"
	Autor  		string gorm:"type:varchar(200);"
	Edit   		string gorm:"type:varchar(255);"
	Fecha   	string gorm:"type:varchar(255);"
}


func main() {
	r := gin.Default()
	r.GET("/", Inicio)
	r.GET("/Libros", obtenerloslibros)
	r.GET("/Libros/:id", obtnerunsololibro)
	r.POST("/Libros", libronuevo)
	r.DELETE("/Libros/:id", eliminarlibro)
	r.PUT("/Libros/:id", actualizarlibro)
	r.Run(":8080")
	

	db, err = gorm.Open("mysql", "root:@(127.0.0.1)/bibloteca?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}


func Inicio(c *gin.Context) {
	var Libro []Libro
	Libro, _ =obtenerloslibros(Libro)
}

func obtenerloslibros(c *gin.Context) {
	var Libro[] Libro
	var err error
	books, err = obtenerloslibros(Libro)	
}

func obtnerunsololibro(c *gin.Context) {
	var Libro Libro
	var autor Autor
	var err error
	id := c.Param("id")
	book, err =obtenerloslibros(idInt, Libro)
}

func eliminarlibro(c *gin.Context) {
	var Libro Libro
	var err error
	errMessage := ""
	id := c.Param("id")

	if err == nil {
		if Libro.ID != 0 {
			err = eliminarlibro(Libro)
			if err == nil {
				c.JSON(200, gin.H{
					"message": "Eliminado correctamente.",
				})
			} else {
				errMessage = err.Error()
			}
		errMessage = err.Error()
	}
	if len(errMessage) != 0 {
		c.JSON(400, gin.H{
			"message": "Ah ocurrido un error: " + err.Error() + ".",
		})
	}
}

func actualizarlibro(c *gin.Context) {
	var Libro Libro
	var err error
	ID := c.Param("ID")
	Nombre := c.PostForm("Nombre")
	Descripcion := c.PostForm("Descripcion")
	Autor := c.PostForm("Autor")
	Editorial:= c.PostForm("Editorial")
	Fecha:= c.PostForm("Fecha")
	Libro, err = obtenerloslibros(IDint, Libro)
	if err == nil {
		if Libro.ID != 0 {
			if len(Nombre) != 0 {
				Libro.Nombre = Nombre
			}
			if len(Descripcion) != 0 {
				Libro.Descripcion = Descripcion
			}
			if len(autor) != 0 {
				Libro.Autor = Autor
			}
			if len(Editorial) != 0 {
				Libro.Editorial = Editorial
			}
			if len(Fecha) != 0 {
				Libro.Fecha= Fecha
			}
			fmt.Println(Libro)
			err = actualizarlibro(Libro)
			if err == nil {
				c.JSON(200, gin.H{
					"message": "Actualizado correctamente.",
				})
			} else {
				errMessage = err.Error()
			}
		}
	}
}

func libronuevo(c *gin.Context) {
	var err error
	Nombre := c.PostForm("Nombre")
	Descripcion := c.PostForm("Descripcion")
	Autor := c.PostForm("Autor")
	Editorial := c.PostForm("Editorial")
	Fecha := c.PostForm("Fecha")
	if len(Nombre) != 0 || len(Descripcion) != 0 || len(autor) != 0 || len(Editorial) != 0 || len(Fecha) != 0 {
		if !match {
			c.JSON(400, gin.H{
				"message": "Formato de fecha inválido."
			})
		} else {
			Libro := Libro{Nombre: Nombre, Descripcion: Descripcion, Autor: Autor, Editorial: Editorial, Fecha: Fecha}
			err = Crearlibro(Libro)
			if err == nil {
				c.JSON(200, gin.H{
					"message": "Registrado correctamente",
				})
			} else {
				c.JSON(400, gin.H{
					"message": "Ocurrió un error: " + err.Error(),
				})
			}
		}
	}
}



//CREATE DATABASE biblioteca;
//USE biblioteca;

//CREATE TABLE libros(
//ID int auto_increment primary key,
//Nombre varchar(100),
//Descripcion varchar(450),
//Autor varchar(200),
//Editoria; varchar(200),
//Fecha date 
//)