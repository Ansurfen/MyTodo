sudo apt install net-tools
sudo apt install curl
curl -fsSL get.docker.com -o get-docker.sh
sudo sh get-docker.sh --mirror Aliyun

docker_images=(
    "mysql"
    "redis"
    "mongo"
    "bitnami/etcd"
    "rabbitmq:management"
    "minio/minio"
    "elasticsearch:8.10.2"
    "kibana:8.10.2"
    "grafana/grafana"
    "prom/prometheus"
    "jaegertracing/all-in-one:latest"
    "nginx"
    "portainer/portainer"
    "grafana/loki"
)

for image in "${docker_images[@]}"; do
    sudo docker pull "$image"
done

sudo docker network create es

sudo docker run -d --name grafana -p 3000:3000 grafana/grafana
sudo docker run --name mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql
sudo docker run -p 6379:6379 --name redis -d redis
sudo docker run --name mongo -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=123456 -d mongo
sudo docker run -id --name rabbitmq -p 15672:15672 -p 5672:5672 -e RABBITMQ_DEFAULT_USER=root -e RABBITMQ_DEFAULT_PASS=123456 rabbitmq:management
sudo docker run -d --name elasticsearch --net es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:8.10.2
sudo docker run -d --name kibana --net es -p 5601:5601 kibana:8.10.2
sudo docker run -d --name etcd -p 2379:2379 -p 2380:2380 -e ALLOW_NONE_AUTHENTICATION=yes -e ETCD_ADVERTISE_CLIENT_URLS=https://etcd-server:2379 bitnami/etcd
sudo mkdir -p /todo/minio/config
sudo mkdir -p /todo/minio/data
sudo docker run -p 9000:9000 -p 9090:9090 --name minio -d --restart=always -e "MINIO_ACCESS_KEY=root" -e "MINIO_SECRET_KEY=root123456" -v /home/minio/data:/data -v /home/minio/config:/root/.minio minio/minio server /data --console-address ":9090" -address ":9000"

sudo mkdir -p ~/todo/prom/{data,config,rules}
chmod -R 777 ~/todo/prom/data
chmod -R 777 ~/todo/prom/config
chmod -R 777 ~/todo/prom/rules
sudo docker run -d -p 9090:9090 -p 9100:9100 --name prometheus \
    -v /etc/localtime:/etc/localtime:ro \
    -v ~/todo/prom/data:/prometheus/data \
    -v ~/todo/prom/config:/etc/prometheus \
    -v ~/todo/prom/rules:/prometheus/rules \
    prom/prometheus --config.file=/etc/prometheus/prometheus.yml --web.enable-lifecycle

sudo mkdir -p ~/todo/es/{data,config,plugins}
chmod -R 777 ~/todo/es/data
chmod -R 777 ~/todo/es/config
chmod -R 777 ~/todo/es/plugins
sudo docker run -d --restart=always --name es \
 --network es -p 9200:9200 -p 9300:9300 --privileged \
 -v ~/todo/es/data:/usr/share/elasticsearch/data -v ~/todo/es/plugins:/usr/share/elasticsearch/plugins \
 -v ~/todo/es/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -e "discovery.type=singl
e-node" -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" elasticsearch:8.10.2

sudo mkdir -p ~/todo/kibana/{config,data}

sudo chmod -R 777 ~/todo/kibana/config
sudo chmod -R 777 ~/todo/kibana/data

sudo docker run -d \
    --restart=always \
    --name kibana \
    --network es \
    -p 5601:5601 \
    -e ELASTICSEARCH_HOSTS=http://es:9200 \
    kibana:8.10.2

sudo docker run -d --name jaeger  \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
    -p 5775:5775/udp     -p 6831:6831/udp  \
    -p 6832:6832/udp     -p 5778:5778     -p 16686:16686 \
        -p 14268:14268     -p 14250:14250     -p 9411:9411     jaegertracing/all-in-one

sudo mkdir -p ~/todo/nginx/{config,data}

sudo docker run -d -p 4030:80 --name nginx --restart=always -v ~/todo/nginx/nginx.conf:/etc/nginx/nginx.conf nginx

sudo docker run -d -p 6766:9000 -v /var/run/docker.sock:/var/run/docker.sock --restart=always --name portainer portainer/portainer

mkdir -p ~/todo/loki/index
mkdir -p ~/todo/loki/chunks
chmod -R 777 ~/todo/loki/index
chmod -R 777 ~/todo/loki/chunks

sudo docker run -d \
    --name loki \
    --privileged=true \
    -v ~/todo/loki:/mnt/config \
    -v ~/todo/loki/index:/opt/loki/index \
    -v ~/todo/loki/chunks:/opt/loki/chunks \
    -p 3100:3100 \
    grafana/loki -config.file=/mnt/config/loki-config.yaml
