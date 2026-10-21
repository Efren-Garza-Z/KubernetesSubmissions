# Log Output App

This is a simple application for Exercise 1.1 that generates a random string on startup and prints it every 5 seconds along with a timestamp.

## How to run

1. **Build and push the image:**
   ```bash
   docker build -t edgz082002/log-output:v1 .
   docker push edgz082002/log-output:v1
   ```

## Deploy to Kubernetes:

```Bash
kubectl create deployment log-output-dep --image=edgz082002/log-output:v1
Check the logs:
```

```Bash
kubectl get pods
kubectl logs -f <your-pod-name>
```