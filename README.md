## CLI para Interactuar con Kubernetes

Este es un programa escrito en Go que actúa como una interfaz de línea de comandos (CLI) para interactuar con un clúster de Kubernetes. Permite listar nodos, servicios y pods del clúster, así como también vigilar los cambios en los pods en tiempo real.


El programa permite ejecutar comandos para interactuar con el clúster de Kubernetes. Los siguientes son algunos ejemplos de cómo usarlo:

### Listar nodos:

```bash
    ./agaKube ls nodes --kubeconfig=path/to/kubeconfig
```    

### Listar servicios:

```bash
    ./agaKube ls srv --kubeconfig=path/to/kubeconfig
```

### Listar pods:

```bash
   ./agaKube ls pods --kubeconfig=path/to/kubeconfig
```

### Vigilar los pods:

```bash
    ./agaKube ls pods -w --kubeconfig=path/to/kubeconfig
```

Los comandos admiten la bandera --kubeconfig para especificar la ruta al archivo kubeconfig.
Contribución

¡Siéntete libre de contribuir al proyecto! Si encuentras problemas, puedes crear una issue en este repositorio o enviar un pull request con tus cambios.
Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.