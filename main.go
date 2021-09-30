package main

import "fmt"

func main() {
	opcion_menu_principal := 0
	var sino string

	menu_principal := ` Menú Principal
	1. Cambiar su contraseña
	2. Editar perfil
	3. Foro
	4. Finalizar programa`

	menu_perfil := `Menú Perfil
	1. Cambiar foto de perfil
	2. Cambiar nombre de usuario
	3. Editar biografía
	4. Volver al menú principal`

	menu_foro := ` Menú Foro
	1. Publicar nueva entrada en el foro
	2. Editar una entrada del foro
	3. Eliminar una entrada del foro
	4. Volver al menú principal`

	for {

		if opcion_menu_principal == 0 {

			fmt.Println(menu_principal)

			_, err := fmt.Scanf("%d", &opcion_menu_principal)
			if err != nil {
				panic(err)
			}

		}

		if opcion_menu_principal == 1 {

			fmt.Println("La contraseña ha sido cambiada con éxito!!!\n¿Desea seguir navegando?")
			_, err := fmt.Scanf("%s", &sino)
			if err != nil {
				panic(err)
			}

			if sino == "NO" {
				break
			}
			opcion_menu_principal = 0

		} else if opcion_menu_principal == 2 {

			fmt.Println(menu_perfil)

			var opcion_menu_perfil int
			_, err := fmt.Scanf("%d", &opcion_menu_perfil)
			if err != nil {
				panic(err)
			}

			if opcion_menu_perfil == 1 {
				fmt.Println("Se ha cambiado la foto de perfil con éxito!!!")
			} else if opcion_menu_perfil == 2 {
				fmt.Println("Se ha cambiado el nombre de usuario con éxito!!!")
			} else if opcion_menu_perfil == 3 {
				fmt.Println("Se ha editado la biografía con éxito!!!")
			} else if opcion_menu_perfil == 4 {
				opcion_menu_principal = 0
			}

		} else if opcion_menu_principal == 3 {
			fmt.Println(menu_foro)

			var opcion_menu_foro int
			_, err := fmt.Scanf("%d", &opcion_menu_foro)
			if err != nil {
				panic(err)
			}

			if opcion_menu_foro == 1 {
				fmt.Println("Se ha publicado una nueva entrada con éxito!!!")
			} else if opcion_menu_foro == 2 {
				fmt.Println("Se ha editado la entrada con éxito!!!")
			} else if opcion_menu_foro == 3 {
				fmt.Println("Se ha eliminado la entreda con éxito!!!")
			} else if opcion_menu_foro == 4 {
				opcion_menu_principal = 0
			}
		} else if opcion_menu_principal == 4 {
			break
		}
	}
}
