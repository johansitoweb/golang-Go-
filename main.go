package main

import (
    "fmt"
)

type Proyecto struct {
    Nombre        string
    Descripcion   string
    FechaInicio   string
    FechaFin      string
}

var proyectos []Proyecto

func agregarProyecto(nombre, descripcion, fechaInicio, fechaFin string) {
    proyecto := Proyecto{
        Nombre:      nombre,
        Descripcion: descripcion,
        FechaInicio: fechaInicio,
        FechaFin:    fechaFin,
    }
    proyectos = append(proyectos, proyecto)
    fmt.Println("Proyecto agregado exitosamente.")
}

func listarProyectos() {
    fmt.Println("Lista de Proyectos:")
    for i, proyecto := range proyectos {
        fmt.Printf("%d. %s - %s (Inicio: %s, Fin: %s)\n", i+1, proyecto.Nombre, proyecto.Descripcion, proyecto.FechaInicio, proyecto.FechaFin)
    }
}

func main() {
    var opcion int
    for {
        fmt.Println("1. Agregar Proyecto")
        fmt.Println("2. Listar Proyectos")
        fmt.Println("3. Salir")
        fmt.Print("Seleccione una opción: ")
        fmt.Scan(&opcion)

        switch opcion {
        case 1:
            var nombre, descripcion, fechaInicio, fechaFin string
            fmt.Print("Nombre del Proyecto: ")
            fmt.Scan(&nombre)
            fmt.Print("Descripción del Proyecto: ")
            fmt.Scan(&descripcion)
            fmt.Print("Fecha de Inicio (YYYY-MM-DD): ")
            fmt.Scan(&fechaInicio)
            fmt.Print("Fecha de Fin (YYYY-MM-DD): ")
            fmt.Scan(&fechaFin)
            agregarProyecto(nombre, descripcion, fechaInicio, fechaFin)
        case 2:
            listarProyectos()
        case 3:
            fmt.Println("Saliendo...")
            return
        default:
            fmt.Println("Opción no válida, por favor intente de nuevo.")
        }
    }
}
