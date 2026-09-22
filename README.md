# BlackChannel

sistema de mensajeria por terminal ====================================
el usuario "guest" se conecta al puerto de la direccion ip del host permitiendo el intercambio 
de mensajes entre estos dos usuarios 

(actualmente conecta al puerto 8081 de forma predeterminada)

para ejecutar como host: 
    go run main.go host
y como guest:
    go run main.go guest 


=========================================
cosas a implementar 
-parametro para elegir puerto  []
-eleccion de contrasena segura dsp de ingresar puerto []
    (los dos usuarios deberan tener la misma contrasena para poder iniciar la comunicacion)
-interfaz fachera con lipgloss y bubletea []
-menu para poder elegir todo []