# SimuladorDeTrenes

Descargar y extraer el archivo ZIP

Descarga el archivo ZIP y extraelo en cualquier carpeta de tu computadora.

Instalar Go (si no lo tiene instalado)

Descarga Go desde https://go.dev/dl/

Instala Go siguiendo las instrucciones de la página oficial. Si usa Windows, puede abrir una terminal (cmd o PowerShell) y navegar hasta la carpeta usando cd ruta/de/proyecto

En Linux/macOS, debe asegurarse de que Go esté en su PATH


Para comprobar que está instalado correctamente, hay que abrir una terminal y ejecuta:


#go version


Si muestra algo como go1.20.5 o cualquier otra versión, significa que está listo.


Instala dependencias en una terminal, en la carpeta del proyecto

#go mod tidy


Ejecutar el programa

En la misma terminal, navega hasta la carpeta donde está el código y ejecuta:


#go run main.go

Si todo está bien, se abrirá la ventana del simulador de trenes.
