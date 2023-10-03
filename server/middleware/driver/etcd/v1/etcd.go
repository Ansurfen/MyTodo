package etcd

import (
	"MyTodo/conf"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	gresolver "google.golang.org/grpc/resolver"
)

type Endpoint struct {
	Group   string `json:"group"`
	Name    string `json:"name"`
	Host    string `json:"host"`
	Port    uint64    `json:"port"`
	Version string `json:"version"`
}

type EndpointUnit struct {
	Name    string `json:"name"`
	LeaseID clientv3.LeaseID
	E       Endpoint
}

type Naming struct {
	*clientv3.Client
	target     string
	manager    endpoints.Manager
	serviceMap map[string]*EndpointUnit
}

type KV struct {
	*clientv3.Client
	clientv3.KV
}

func (kv *KV) GetWithPrefix(key string) (*clientv3.GetResponse, error) {
	return kv.KV.Get(context.TODO(), key, clientv3.WithPrefix())
}

type Client struct {
	*clientv3.Client
}

func New(opt conf.ETCDOption) *Client {
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: time.Duration(opt.DialTimeout) * time.Second,
		Endpoints:   opt.EndPoints,
	})
	if err != nil {
		panic(err)
	}
	return &Client{Client: cli}
}

func (e *Client) NewNaming(scope string) *Naming {
	m, err := endpoints.NewManager(e.Client, scope)
	if err != nil {
		panic(err)
	}
	return &Naming{
		Client:     e.Client,
		manager:    m,
		target:     scope,
		serviceMap: make(map[string]*EndpointUnit),
	}
}

func (e *Client) NewKV() *KV {
	return &KV{
		Client: e.Client,
		KV:     clientv3.NewKV(e.Client),
	}
}

func (e *Client) ParseEndpoint(data []byte) (res Endpoint, err error) {
	eps := endpoints.Endpoint{}
	err = json.Unmarshal(data, &eps)
	if err != nil {
		return
	}
	if v, ok := eps.Metadata.(string); ok {
		err = json.Unmarshal([]byte(v), &res)
	} else {
		err = errors.New("invalid metadata")
	}
	return
}

func (e *Client) Put(k, v string) (*clientv3.PutResponse, error) {
	return e.Client.Put(context.TODO(), k, v)
}

func (e *Client) Get(k string) (*clientv3.GetResponse, error) {
	return e.Client.Get(context.TODO(), k)
}

func (e *Client) Watch(key string, callback func(*clientv3.Event)) {
	for wresp := range e.Client.Watch(context.Background(), key) {
		for _, eve := range wresp.Events {
			callback(eve)
		}
	}
}

func (e *Naming) RegisterEndpoint(ep Endpoint) error {
	data, err := json.Marshal(ep)
	if err != nil {
		return err
	}
	eps := endpoints.Endpoint{
		Addr:     fmt.Sprintf("%s:%d", ep.Host, ep.Port),
		Metadata: string(data),
	}
	lease, err := e.Grant(context.TODO(), 30)
	if err != nil {
		return err
	}
	key := e.getFullServiceName(ep.Name, lease.ID)
	err = e.manager.AddEndpoint(context.TODO(), key, eps, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	log.Println("RegisterEndpoint success:", key)
	ch, err := e.KeepAlive(context.TODO(), lease.ID)
	if err != nil {
		return err
	}
	go func() {
		for {
			ka := <-ch
			log.Println("ttl:", ka.ID, ka.TTL)
		}
	}()
	e.serviceMap[ep.Name] = &EndpointUnit{Name: ep.Name, LeaseID: lease.ID, E: ep}
	return nil
}

func (e *Naming) getFullServiceName(name string, leaseID clientv3.LeaseID) string {
	return fmt.Sprintf("%s/%s/%d", e.target, name, leaseID)
}

func (e *Naming) GetPathServerName(name string) string {
	return fmt.Sprintf("%s/%s", e.target, name)
}

func (e *Naming) ReleaseEndpoint(name string) error {
	eu := e.serviceMap[name]
	if eu == nil {
		return nil
	}
	err := e.manager.DeleteEndpoint(context.TODO(), e.getFullServiceName(name, eu.LeaseID))
	if err != nil {
		log.Fatalf("Delete endpoint error %v", err)
		return err
	}
	_, err = e.Client.Revoke(context.TODO(), eu.LeaseID)
	if err != nil {
		log.Fatalf("Revoke lease error %v", err)
		return err
	}
	delete(e.serviceMap, name)
	log.Printf("DeleteEndpoint [%s] success\n", name)
	return nil
}

func (e *Naming) ReleaseAllEndpoint() {
	for k := range e.serviceMap {
		err := e.ReleaseEndpoint(k)
		if err != nil {
			log.Fatalln("Ignore Failure Continue...")
		}
	}
}

func (e *Naming) NewEtcdResolver() (gresolver.Builder, error) {
	etcdResolver, err := resolver.NewBuilder(e.Client)
	if err != nil {
		log.Fatalf("Etcd resolver error %v", err)
		return nil, err
	}
	return etcdResolver, nil
}
