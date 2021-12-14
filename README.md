# grype-scan-aws-ecr

Use <https://github.com/anchore/grype/> to do Docker layer scanning on various AWS ECR repositories.

```bash
kubectl get pods -o custom-columns="IMAGE:.spec.containers[*].image" -A > images.txt
cat images.txt | grep '.ecr.' | sort -u > ~/Downloads/ecr_images.txt
go run main.go
```
