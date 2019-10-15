#环境
coredns主机 172.20.52.36
etcd主机 172.20.57.58

#运行coredns
./coredns -config corefile

#运行etcd
#注意etcd的版本3.4.2 API版本3.4 
./etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 \
  --initial-cluster s1=http://0.0.0.0:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --log-level info \
  --logger zap \
  --log-outputs stder

#put上传etcd数据
./etcdctl put /skydns/local/skydns/ '{"host":"1.1.1.1","ttl":60}'

#测试
dig skydns.local  @172.20.52.36  +short

