# 🏫 Escuela Colombiana de Ingeniería
## 📚 Arquitecturas de Software – ARSW
## ╰┈➤ - [ 🪼 ] | Ejercicio Introducción al paralelismo - Hilos - Caso BlackListSearch ┆⤿⌗


---

Nombres:
- Daniel Palacios Moreno 
- Sofia Nicolle Ariza Goenaga

---

## 📖 Dependencias
### 🔗 Lecturas recomendadas:
- [Threads in Java](http://beginnersbook.com/2013/03/java-threads/) *(hasta “Ending Threads”)*
- [Threads vs Processes](http://cs-fundamentals.com/tech-interview/java/differences-between-thread-and-process-in-java.php)

---

## 📝 Descripción
Este ejercicio introduce la **programación con hilos en Java** y su aplicación en un caso concreto de validación de direcciones IP en listas negras y un ejercicio inicial para aclimatar a los miembros del equipo que se presenta en los puntos siguientes.

---

## ⚙️ Parte I – Introducción a Hilos en Java
1. Completar la clase **`CountThread`** para definir el ciclo de vida de un hilo que imprima números entre A y B.
2. En el método **`main`** de **`CountMainThreads`**:
	- Crear 3 hilos con intervalos:
		- Hilo 1 → `[0..99]`
		- Hilo 2 → `[99..199]`
		- Hilo 3 → `[200..299]`
	- Iniciar con `start()`.
	- Revisar la salida.
	- Cambiar `start()` por `run()`. ➜ **Analizar diferencias y explicar.**

---

## 🔍 Parte II – Ejercicio Black List Search

### 🎯 Contexto
Se desarrolla un componente de **seguridad informática** que valida direcciones IP en miles de listas negras y reporta aquellas presentes en al menos **5 listas**.

### 🧩 Componentes principales:
- **`HostBlackListsDataSourceFacade`** → Fachada para consultar listas negras y reportar hosts peligrosos. *(Thread-Safe, NO modificable)*
- **`HostBlackListsValidator`** → Método `checkHost` que valida un host y reporta si es confiable o no.

📊 Ejemplo de LOGs:
INFO: HOST 205.24.34.55 Reported as trustworthy
INFO: HOST 205.24.34.55 Reported as NOT trustworthy


### 🚀 Tareas:
1. Crear una clase **Thread** que busque en un segmento de servidores y registre ocurrencias.
2. Modificar `checkHost(N)` para:
	- Dividir espacio de búsqueda en **N hilos**.
	- Ejecutar en paralelo y esperar con `join()`.
	- Sumar ocurrencias y reportar confiabilidad.
	- Mantener LOGs verídicos sobre listas revisadas.

---

## 💡 Parte II.I – Discusión (no implementar aún)
¿Cómo optimizar la búsqueda para detenerla cuando ya se alcanzan las ocurrencias mínimas? ➜ Introducir **mecanismos de sincronización** y **cancelación temprana**.

---

## 📊 Parte III – Evaluación de Desempeño

### 🔬 Experimentos:
1. 1 hilo.
2. Núm. de hilos = núm. de núcleos.
3. Núm. de hilos = 2 × núm. de núcleos.
4. 50 hilos.
5. 100 hilos.

📈 Monitorear con **jVisualVM**: consumo de CPU y memoria.  
➜ Graficar **tiempo de solución vs. número de hilos** y analizar.

---

## 📐 Parte IV – Análisis con Ley de Amdahls

- ¿Por qué el mejor desempeño no ocurre con 500 hilos?
- Comparar resultados con 200 hilos.
- Evaluar desempeño con núm. de hilos = núm. de núcleos vs. el doble.
- Reflexionar sobre escenarios distribuidos (100 máquinas vs. 1 CPU con 100 hilos).

---
