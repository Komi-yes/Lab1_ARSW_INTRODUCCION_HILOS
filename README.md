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

### Solución:

El archivo CountThread.java se diseña para que se construya con el rango necesario para que en el metodo run con un for simple se recorra el rango y se impriman los valores.
Para ver el [repositorio Sofia](https://github.com/Sofia-ariza-783/ARSW_Lab_I.git).

Cuando se cambia el start por run, el hilo se ejecuta en el hilo principal, por lo que se imprimen los valores en el orden correcto.
- **Con start:**
![img.png](img.png)

- **Con run:**
![img_1.png](img_1.png)

---

## 🔍 Parte II – Ejercicio Black List Search

### 🎯 Contexto
Se desarrolla un componente de **seguridad informática** que valida direcciones IP en miles de listas negras y reporta aquellas presentes en al menos **5 listas**.

### 🧩 Componentes principales:
- **`HostBlackListsDataSourceFacade`** → Fachada para consultar listas negras y reportar hosts peligrosos. *(Thread-Safe, NO modificable)*
- **`HostBlackListsValidator`** → Método `checkHost` que valida un host y reporta si es confiable o no.

📊 Ejemplo de LOGs:
- INFO: HOST 205.24.34.55 Reported as trustworthy
- INFO: HOST 205.24.34.55 Reported as NOT trustworthy


### 🚀 Tareas:
1. Crear una clase **Thread** que busque en un segmento de servidores y registre ocurrencias.
2. Modificar `checkHost(N)` para:
	- Dividir espacio de búsqueda en **N hilos**.
	- Ejecutar en paralelo y esperar con `join()`.
	- Sumar ocurrencias y reportar confiabilidad.
	- Mantener LOGs verídicos sobre listas revisadas.

### Solución:

Consideramos que el método planteado en el archivo era ineficiente, ya que en el ejercicio anterior habíamos probado un enfoque similar. Por ello, decidimos diseñar una solución más limpia y eficiente, que aprovechara mejor el uso de los hilos y evitara depender de que todos finalizaran la búsqueda para poder reportar las coincidencias en las listas inseguras.

Para implementar nuestra propuesta, modificamos varios tipos de variables para que fueran Thread-Safe y pudieran ser compartidas directamente entre los hilos sin necesidad de usar la etiqueta synchronized. Este fue el caso de occurrencesCount, checkedListsCount y stopFlag. Con estas variables accesibles, incorporamos dos contadores: CountDownLatch stopLatch y completionLatch. El primero detiene el programa cuando se alcanzan las 5 ocurrencias, apoyándose en la variable stopFlag; el segundo controla el caso en que no se logren dichas ocurrencias mínimas.

El método checkHost inicializa completionLatch con el número de hilos y stopLatch con el número mínimo de ocurrencias requeridas. Luego, mediante un bucle, crea e inicia los hilos. Cada hilo recorre su segmento de la lista, verificando en cada iteración el estado de stopFlag. Si se alcanzan las 5 ocurrencias, todos los hilos se detienen; en caso contrario, se completa la búsqueda en toda la lista y el resultado es capturado por completionLatch.

De esta manera, se optimiza el tiempo de ejecución: no es necesario esperar a que todos los hilos terminen para reportar un host inseguro, pero se garantiza que, si no se encuentran las 5 coincidencias, el sistema lo registre correctamente como confiable.

---

## 💡 Parte II.I – Discusión (no implementar aún)
¿Cómo optimizar la búsqueda para detenerla cuando ya se alcanzan las ocurrencias mínimas? ➜ Introducir **mecanismos de sincronización** y **cancelación temprana**.

### Solución:

Aunque hay multiples soluciones que podrian ayudar a que la busqueda se detenga cuando se encuentra todas las coincidencias necesarias, la que nosotros consideramos mas interesante fue implementando una variable que funcionara como "luz roja" que indicara cuando tenian que detenerse los hilos, junto con los CountDownLatch que se encargan de controlar cuantas coincidencias se hicieron. Esta solucion en comparacion el join simple, agrega mas lineas de codigo, mas complejidad y aumenta la carga cognitiva del codigo. 

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
