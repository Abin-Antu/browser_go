

# browser_go

`browser_go` es un paquete de Go ligero y multiplataforma que te permite abrir URLs en el navegador web predeterminado del usuario. Es compatible con Windows, macOS, Linux y Windows Subsystem for Linux (WSL).

## ¿Por qué usar browser_go?

Al desarrollar aplicaciones web o APIs (por ejemplo, con el framework web Gin), a menudo deseas que tu panel de control o página principal se abra automáticamente en el navegador en cuanto inicie tu servidor.  
`browser_go` hace que esto sea sencillo y mejora tu flujo de trabajo de desarrollo.

---

## Características

- **Multiplataforma**: Funciona sin problemas en Windows, macOS, Linux y WSL.
- **Con detección de WSL**: Detecta WSL y abre las URLs en el navegador nativo de Windows.
- **API sencilla**: Solo llama a `OpenURL` con tu URL.
- **Sin dependencias externas**: Utiliza únicamente la biblioteca estándar de Go y las utilidades del sistema integradas.

---

## Instalación

Instala el paquete con:

```bash
go get github.com/Abin-Antu/browser_go
```

---

## Uso

Puedes usar `browser_go` para abrir cualquier página web, como tu panel de control, interfaz de API o documentación, automáticamente cuando inicie tu servidor web de Go.

### Ejemplo: Abrir un panel de control o página principal cuando inicie el servidor

```go
package main

import (
    "github.com/Abin-Antu/browser_go/browser"
)

func main() {
    url := "http://localhost:8080/dashboard"
    browser.OpenURL(url) // This will open the dashboard in the default browser

    // Start your web server here (example with Gin or net/http)
    // router := gin.Default()
    // router.Run(":8080")
}
```

Solo llama a `browser.OpenURL(url)` con la URL que deseas abrir.  
Esto es especialmente útil para iniciar automáticamente tu panel de control web, página principal o documentación de la API durante el desarrollo local.

---

## Licencia

Licencia MIT. Consulta [LICENSE](LICENSE) para más detalles.
