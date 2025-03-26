package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Train representa un tren con nombre, velocidad, ruta y prioridad.
type Train struct {
	Name     string    // Nombre del tren
	Speed    float64   // Velocidad en km/h
	Route    []float64 // Distancias entre paradas en kilómetros
	Priority bool      // Indica si el tren tiene prioridad
}

// Lista global de trenes
var trains []Train

func main() {
	// Crear una nueva aplicación Fyne
	app := app.New()
	win := app.NewWindow("Simulador de Trenes") // Crear una ventana
	win.Resize(fyne.NewSize(500, 400))          // Establecer tamaño de la ventana

	// Campos de entrada para capturar datos del usuario
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Nombre del tren")

	speedEntry := widget.NewEntry()
	speedEntry.SetPlaceHolder("Velocidad (km/h)")

	routeEntry := widget.NewEntry()
	routeEntry.SetPlaceHolder("Distancias (km, separadas por comas)")

	// Checkbox para marcar si el tren es prioritario
	priorityCheck := widget.NewCheck("Prioritario", nil)

	// Botón para añadir un nuevo tren a la lista
	addButton := widget.NewButton("Añadir Tren", func() {
		name := nameEntry.Text

		// Convertir la velocidad ingresada a float64
		speed, _ := strconv.ParseFloat(speedEntry.Text, 64)

		// Convertir la lista de distancias ingresadas a un slice de float64
		routeStr := strings.Split(routeEntry.Text, ",")
		var route []float64
		for _, r := range routeStr {
			dist, _ := strconv.ParseFloat(strings.TrimSpace(r), 64)
			route = append(route, dist)
		}

		// Agregar el nuevo tren a la lista
		trains = append(trains, Train{Name: name, Speed: speed, Route: route, Priority: priorityCheck.Checked})
		fmt.Println("Tren agregado:", name)
	})

	// Botón para iniciar la simulación
	simulateButton := widget.NewButton("Simular", func() {
		simulateTrains()
	})

	// Definir el contenido de la ventana
	win.SetContent(container.NewVBox(
		nameEntry,
		speedEntry,
		routeEntry,
		priorityCheck,
		addButton,
		simulateButton,
	))

	// Mostrar la ventana y ejecutar la aplicación
	win.ShowAndRun()
}

// simulateTrains organiza los trenes según prioridad y los simula
func simulateTrains() {
	// Verificar si hay trenes en la lista
	if len(trains) == 0 {
		fmt.Println("No hay trenes para simular")
		return
	}

	// Separar los trenes prioritarios de los normales
	priorityTrains := []Train{}
	normalTrains := []Train{}
	for _, t := range trains {
		if t.Priority {
			priorityTrains = append(priorityTrains, t)
		} else {
			normalTrains = append(normalTrains, t)
		}
	}

	// Mezclar aleatoriamente los trenes normales usando una nueva fuente de aleatoriedad
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rnd.Shuffle(len(normalTrains), func(i, j int) {
		normalTrains[i], normalTrains[j] = normalTrains[j], normalTrains[i]
	})

	// Concatenar primero los trenes prioritarios y luego los normales
	orderedTrains := append(priorityTrains, normalTrains...)
	fmt.Println("Orden de paso:")

	// Calcular y mostrar el tiempo total para cada tren
	for _, t := range orderedTrains {
		timeRequired := calculateTime(t)
		fmt.Printf("Tren %s: %0.2f horas\n", t.Name, timeRequired)
	}
}

// calculateTime calcula el tiempo total de viaje de un tren
func calculateTime(t Train) float64 {
	totalDistance := 0.0

	// Sumar todas las distancias de la ruta
	for _, d := range t.Route {
		totalDistance += d
	}

	// Retornar el tiempo total (distancia / velocidad)
	return totalDistance / t.Speed
}
