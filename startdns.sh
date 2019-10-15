etcd --listen-peer-urls=http://172.20.52.36:4380 --listen-client-urls=http://172.20.52.36:4379  --data-dir=/tmp/etcd  --advertise-client-urls=http://172.20.52.36:4379 --initial-cluster-state=new --wal-dir=/tmp/wal

etcdctl  --endpoints http://172.20.52.36:4379 ls

docker run -p 53:53 -v /tmp:/tmp 172.20.58.169:81/kubernetes/coredns:1.2.6  -conf /tmp/corefile


etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:4379 \
  --advertise-client-urls http://0.0.0.0:4379 \
  --listen-peer-urls http://0.0.0.0:4380 \
  --initial-advertise-peer-urls http://0.0.0.0:4380 \
  --initial-cluster s1=http://0.0.0.0:4380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --log-level info \
  --logger zap \
  --log-outputs stderr



etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:4379 \
  --advertise-client-urls http://0.0.0.0:4379 \
  --listen-peer-urls http://0.0.0.0:4380 \
  --initial-advertise-peer-urls http://0.0.0.0:4380 \
  --initial-cluster s1=http://0.0.0.0:4380 \
  --initial-cluster-state new \
  --log-level info \
  --logger zap \
  --log-outputs stderr
