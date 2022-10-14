package routers

import (
	"encoding/json"
	"net/http"

	"github.com/IsmaEstrella/Twitor/bd"
	"github.com/IsmaEstrella/Twitor/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos Recibidos"+err.Error(), 500)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Contraseña mayor a 6 caracteres", 400)
		return
	}
	//Validacion contra los datos
	//Cuando solo me interesa encontrar un solo parametro se usa _,
	_, encontrado, _ := bd.UsuarioExiste(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe el usuario con ese email", 400)
	}
	_, status, _ := bd.InsertarRegistro(t)
	if err != nil {
		http.Error(w, "Registro duplicado"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se registro"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
