package main
import (
	"fmt"
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)
func obtenerConexiones()  [5]*sql.DB {
	var conexiones [5]*sql.DB
	var ips = [5]string{"127.0.0.1","127.0.0.1","127.0.0.1","127.0.0.1","127.0.0.1"};
	for i:=0; i < 5; i++ {
		db, err := sql.Open("mysql", "root@tcp(" + ips[i] +":3306)/pruebas_inv")
		if  err != nil {
			panic("Fallo una de las conexiones");
		} else {
			conexiones[i] = db
		}
	}
	return conexiones;
}

func main() {
	dat, err := ioutil.ReadFile("post_busqueda.txt")
	if err != nil {
		panic("No se pudo abrir el archivo de busquedas")
	}
	tiempos, err := os.Create("tiempos.txt")
	if err != nil {
		panic("No se pudo abrir el archivo de tiempos")
	}
	conexiones := obtenerConexiones()
	texto := string(dat)
	idsParaBuscar := strings.Split(texto, "\n")
	limit := len(idsParaBuscar)
	fmt.Println(limit)
	for i:= 0; i < limit; i++ {
		idLength := len(idsParaBuscar[i])
		nodoParaBuscar, _ := strconv.Atoi(idsParaBuscar[i][0:2])
		idParaBuscar, _ := strconv.Atoi(idsParaBuscar[i][3:idLength])
		nodoParaBuscar = nodoParaBuscar - 1
		start := time.Now()
		row, err := conexiones[nodoParaBuscar].Query("SELECT * FROM  post where id = ?", idParaBuscar)
		tiempos.WriteString("" + fmt.Sprintf("%f", float64(time.Since(start) / time.Millisecond)) + "\n")
		fmt.Println(row)
		if err != nil {
			panic(err.Error())
		}
		row.Close()
	}
	defer tiempos.Close()
	fmt.Println("Terminado")
}