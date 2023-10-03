docker_containers=(
    "mysql"
    "redis"
    "mongo"
    "etcd"
    "rabbitmq"
    "minio"
    "elasticsearch"
    "grafana"
    "kibana"
    "prometheus"
    "jaeger"
    "nginx"
    "loki",
    "portainer"
)

for image in "${docker_images[@]}"; do
    sudo docker start "$image"
done