# 📓 Diario de aprendizaje de Go

Cada entrada registra **el error que encontré** y **cómo lo resolví**.

---

## 05/10/2026 — Primeros pasos

### `expected 'package', found 'EOF'`
- **Qué pasaba:** el archivo `.go` estaba vacío.
- **Regla:** Go siempre espera `package` en la primera línea.
- **Aprendido:** `found 'EOF'` = *"llegué al final del archivo y no había nada"*.

### `imported and not used`
- **Qué pasaba:** importé un paquete y no lo usé.
- **Regla:** Go obliga a usar **todos** los imports.

### `declared and not used`
- **Qué pasaba:** creé una variable local y nunca la leí.
- **Regla:** aplica a variables **dentro** de funciones. Las de nivel paquete sí se pueden dejar sin usar.

---

## 06/10/2026 — Variables y ámbitos

### `:=` vs `=` vs `var`
| Forma | Sirve para | Dónde |
|-------|-----------|-------|
| `x := 1` | **crear** (una sola vez) | solo dentro de funciones |
| `x = 1` | **modificar** | donde `x` sea visible |
| `var x = 1` | **crear** | paquete o función |

- Error típico: `var x := 1` → no se puede mezclar `var` y `:=`.

### Scope
- **Nivel paquete** (fuera de `func`): visible en todo el archivo, vive toda la ejecución, puede quedar sin usar.
- **Nivel función**: visible solo ahí, se crea al llegar a esa línea y **debe** usarse.

### Tipos enteros
- `int8` / `int16` / `int` → aceptan negativos.
- `uint8` / `uint16` / `uint` → **solo positivos**.
- `uint8` llega a 255 y **da la vuelta** si se pasa (overflow silencioso).

---

## 06/10/2026 — Paquetes

### `found packages main (...) and calcular (...)`
- **Qué pasaba:** dos packages distintos en la **misma** carpeta.
- **Regla:** **una carpeta = un package**. Para otro package, otra subcarpeta.

### `no required module provides package ...`
- **Qué pasaba:** importaba `gopath.local/sintaxis/calc` pero mi módulo se llama `pruebago`.
- **Regla:** el import debe empezar con el nombre del módulo (`go.mod`) + la ruta de la carpeta.

### Mayúscula vs minúscula
- `func Sumar(...)` → **exportada**, visible desde otros paquetes.
- `func sumar(...)` → **privada**, solo dentro de su paquete.
- Error: `cannot refer to unexported name calcular.sumar`

---

## 06/10/2026 — Funciones y tipos

### `not enough arguments in call`
- **Qué pasaba:** la función pedía `(a, b)` y le pasaba `(x + y)`.

### `mismatched types uint8 and uint16`
- **Qué pasaba:** `x + y` con tipos distintos.
- **Regla:** Go **nunca** convierte tipos solo. Hay que ser explícito: `int16(x)`.

### `a + b evaluated but not used`
- **Qué pasaba:** calculaba un valor y no lo guardaba ni lo devolvía.
- **Regla:** una expresión suelta sin efecto secundario no es válida como sentencia.

---

## 06/10/2026 — Generics

```go
type Entero interface {
    int | int8 | int16 | uint | uint8 | uint16
}

func Sumar[T Entero](a, b T) T {
    return a + b
}
```

### El error que costó entender
```go
var x int8 = 10
var y int16 = 20
calcular.Sumar(x, y)   // ❌ cannot use x ... as type int16
```
- **Por qué:** `T` es **un solo hueco** para los dos argumentos. No puede ser `int8` **y** `int16` a la vez.
- **Soluciones:**
  1. Convertir: `Sumar(int16(x), y)`
  2. Mismo tipo en las variables
  3. Dos parámetros de tipo: `func SumarMixto[T, U Entero](a T, b U) int64`

---

## 06/10/2026 — Entorno

### La terminal de VS Code
- Las credenciales **no viven en la terminal**, viven en los archivos de config del shell (`~/.config/fish/config.fish`, `~/.ssh/...`).
- Warp es un **emulador de terminal** (app), no un shell → no puede incrustarse en VS Code.
- Solución: VS Code → `terminal.integrated.defaultProfile.linux: "fish"` → mismo shell, mismas credenciales.
- `~/go` es el **GOPATH** (ahí vive `gopls` y la caché de módulos) → **nunca** crear proyectos ahí.
