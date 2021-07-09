package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/asveg/order/domain/repository"
	"github.com/asveg/order/handler"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/asveg/order/proto/order"
	service2 "github.com/asveg/order/domain/service"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	"github.com/asveg/common"
	"strconv"
)
var QPS=1000

func main() {
	/*
	1.配置中心
	2.注册中心
	3.链路追踪
	4.数据库连接
	 */
	//1.配置中心
	consulConf, err := common.GetConsulConfig("192.168.10.168",8500, "/micro/config")
	if err !=nil {
		log.Error(err)
	}
	//2.注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.10.168:8500",
		}
	})
	//3.链路追踪
	t, io, err := common.NewTracer("go.micro.service.order","localhost:8045")
	if err !=nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//数据库连接
	mysqlInfo := common.GetMysqlFromConsul(consulConf,"mysql")
	mysqlInfoPort :=strconv.FormatInt(mysqlInfo.Port,10)
	//db,err := gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+"192.168.10.168:3306"+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	db,err := gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@("+mysqlInfo.Host+":"+mysqlInfoPort+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+":"+mysqlInfoPort+")/"+mysqlInfo.Database)
	if err !=nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//initialise database 第一次运行的时候创建表
	//err = repository.NewOrderRepostory(db).InitTable()
	//if err !=nil {
	//	log.Error(err)
	//}

	//connect database
	orderDataService :=service2.NewOrderService(repository.NewOrderRepostory(db))

	//expose prometheus monitor ip:port
	common.PrometheusBoot(9092)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Version("latest"),
		//暴露服务地址到外部
		micro.Address("0.0.0.0:9085"),
		//添加consul注册中心
		micro.Registry(consul),
		//添加链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(t)),
		//添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		// add monitor prometheus
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()


	// Register Handler
	order.RegisterOrderHandler(service.Server(), &handler.Order{orderDataService})


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
