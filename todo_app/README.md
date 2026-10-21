# Todo Project App

This is a simple web server application for the Kubernetes course project that serves an HTML page and logs its startup port.

## How to run

1. **Build and push the image:**

   ```bash
   docker build -t edgz082002/todo-app:v2 .
   docker push edgz082002/todo-app:v2
   ```

## Deploy to Kubernetes:

```Bash
kubectl apply -f manifests/deployment.yaml
```
## Check the application:

```
Bash
kubectl get pods
kubectl port-forward <your-todo-app-pod-name> 8080:8080
```
Then open http://localhost:8080 in your browser.