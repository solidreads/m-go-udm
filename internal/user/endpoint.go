
type{
	Controller func(w http.ResponseWriter, r *http.Request){

	}
	Endpoints struct{
		Create Controller
		Get Controller
		GetAll Controller
		Update Controller
		Delete Controller
}

}


func MakeEndPoint() Endpoints{
	return Endpoints{
		Create: makeCreateEndPoint(),
		Get: makeGetEndPoint(),
		GetAll: makeGetAllEndPoint(),
		Update: makeUpdateEndPoint(),
		Delete: makeDeleteEndPoint(),
	}
}
}

func makeCreateEndPoint() Controller{
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Create")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeGetEndPoint() Controller{
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Get")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeGetAllEndPoint() Controller{
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Get All")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})// do something
	}
}
func makeUpdateEndPoint() Controller{
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Update")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeDeleteEndPoint() Controller{
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Delete user")
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}