Descripción de CommitNameGen
Nombre del Repositorio: CommitNameGen

Descripción:
CommitNameGen es una API desarrollada en Go que utiliza la inteligencia artificial para generar nombres únicos y creativos para cada commit en tu repositorio. Aprovechando el poder de la API de Google Cloud Platform (GCP), esta herramienta asigna automáticamente nombres descriptivos a los commits para facilitar el seguimiento de cambios y añadir un toque de creatividad al flujo de trabajo de desarrollo.

Con CommitNameGen, puedes mejorar la claridad y el contexto de tu historial de commits, haciendo que cada cambio cuente con una descripción generada automáticamente, ahorrando tiempo y estandarizando el proceso de etiquetado.

Características:

Generación automática de nombres para commits utilizando modelos de IA de GCP.
API REST desarrollada en Go para una integración rápida y eficiente.
Facilidad de personalización para diferentes estilos de nombres.
Documentación detallada para implementar en proyectos nuevos o existentes.
Objetivo:
Simplificar el proceso de etiquetado en commits y hacer que el historial de cambios sea más legible y organizado, con un toque de creatividad y uniformidad impulsado por IA.

```
Prompt para Commit:

Tipo: Empieza con el tipo de cambio 
[FEAT], [FIX], [REFACTOR], [DOCS], [STYLE], [TEST], [CHORE], seguido de un breve resumen. 
Ejemplo: 
[FEAT] agrega autenticación con Google

Resumen claro: Describe en pocas palabras el propósito principal del cambio.

Descripción opcional (si es necesario): En la segunda línea, explica el cambio en más detalle, especialmente si afecta funcionalidades clave o necesita contexto adicional.

Ejemplo de commit:

makefile
Copiar código
feat: agrega autenticación con Google

Referencia a tickets/issues (si aplica): Incluye Refs: o Fixes: seguido del ID de la tarea o issue. 
Ejemplo: 
[FIXES] #249 Arreglo de auth0 

```