package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	proxy_port := os.Getenv("envoy_proxy_port")
	if proxy_port == "" {
		fmt.Println("envoy_proxy_port data err")
		return
	}
	branch := os.Getenv("envoy_branch")
	if branch == "" {
		fmt.Println("envoy_branch data err")
		return
	}
	service_name := os.Getenv("envoy_service_name")
	if service_name == "" {
		fmt.Println("envoy_service_name data err")
		return
	}
	service_port := os.Getenv("envoy_service_port")
	if service_port == "" {
		fmt.Println("envoy_service_port data err")
		return
	}
	ssl := os.Getenv("envoy_ssl")
	protocol := os.Getenv("envoy_protocol")
	if protocol != "http" && protocol != "grpc" {
		fmt.Println("envoy_protocol data err")
		return
	}
	certificate_chain := os.Getenv("envoy_certificate_chain")
	if branch == "" && ssl == "1" {
		fmt.Println("envoy_certificate_chain data err")
		return
	}
	private_key := os.Getenv("envoy_private_key")
	if branch == "" && ssl == "1" {
		fmt.Println("envoy_private_key data err")
		return
	}

	envoy_head := `---
admin:
  access_log_path: /tmp/admin_access.log
  address:
    socket_address:
      protocol: TCP
      address: 127.0.0.1
      port_value: 9901
static_resources:
  listeners:
  - name: listener_0
    address:
      socket_address:
        protocol: TCP
        address: 0.0.0.0
        port_value: {proxy_port}
    filter_chains:
    - filters:
      - name: envoy.filters.network.http_connection_manager
        typed_config:
          "@type": type.googleapis.com/envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager
          stat_prefix: ingress_http
          route_config:
            name: local_route
            virtual_hosts:
            - name: local_service
              domains: ["*"]
              routes:`
	envoy_head = strings.Replace(envoy_head, "{proxy_port}", proxy_port, 1) //替换监听端口

	branch_tpl := `
              - match:
                  prefix: "/"
                  headers:
                  - name: trace-branch
                    exact_match: {branch}
                route:
                  cluster: {service_branch}`
	master_branch_tpl := `
              - match:
                  prefix: "/"
                route:
                  cluster: {service_branch}`
	branchSlice := strings.Split(branch, ",")
	if len(branchSlice) == 0 {
		fmt.Println("err: branch info err")
		return
	}
	for _, v := range branchSlice { //替换分支-路由
		if v == "master" {
			envoy_head += strings.Replace(master_branch_tpl, "{service_branch}", service_name+"-"+v, 1)
		} else {
			tmp := strings.Replace(branch_tpl, "{service_branch}", service_name+"-"+v, 1)
			envoy_head += strings.Replace(tmp, "{branch}", v, 1)
		}
	}

	//判断是否加密
	if ssl == "1" {
		envoy_head += `
          http_filters:
          - name: envoy.filters.http.router
      transport_socket:
        name: envoy.transport_sockets.tls
        typed_config:
          "@type": type.googleapis.com/envoy.api.v2.auth.DownstreamTlsContext
          common_tls_context:
            tls_certificate_sds_secret_configs:
            - name: server_cert
  clusters:`
	} else {
		envoy_head += `
          http_filters:
          - name: envoy.filters.http.router
  clusters:`
	}

	cluster_http_tpl := `
  - name: {service_cluster}
    connect_timeout: 0.25s
    type: LOGICAL_DNS
    # Comment out the following line to test on v6 networks
    dns_lookup_family: V4_ONLY
    lb_policy: ROUND_ROBIN
    load_assignment:
      cluster_name: {service_cluster}
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: {service_cluster}
                port_value: {service_port}`
	cluster_grpc_tpl := `
  - name: {service_cluster}
    connect_timeout: 0.25s
    type: LOGICAL_DNS
    # Comment out the following line to test on v6 networks
    dns_lookup_family: V4_ONLY
    lb_policy: ROUND_ROBIN
    http2_protocol_options: {}
    load_assignment:
      cluster_name: {service_cluster}
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: {service_cluster}
                port_value: {service_port}`
	cluster_ssl_tpl := `
    transport_socket:
      name: envoy.transport_sockets.tls
      typed_config:
        "@type": type.googleapis.com/envoy.api.v2.auth.UpstreamTlsContext
        common_tls_context:
          tls_certificate_sds_secret_configs:
          - name: server_cert`
	for _, v := range branchSlice {
		if protocol == "grpc" {
			tmp := strings.Replace(cluster_grpc_tpl, "{service_cluster}", service_name+"-"+v, 3)
			envoy_head += strings.Replace(tmp, "{service_port}", service_port, 1)
		} else if protocol == "http" {
			tmp := strings.Replace(cluster_http_tpl, "{service_cluster}", service_name+"-"+v, 3)
			envoy_head += strings.Replace(tmp, "{service_port}", service_port, 1)
		} else {
			fmt.Println("err: protocol date err")
		}
		if ssl == "1" {
			envoy_head += cluster_ssl_tpl
		}
	}

	secrets_tpl := `
  secrets:
    - name: server_cert
      tls_certificate:
        certificate_chain:
          filename: {secrets_certificate_chain}
        private_key:
          filename: {secrets_private_key}`
	if ssl == "1" {
		tmp := strings.Replace(secrets_tpl, "{secrets_certificate_chain}", certificate_chain, 1)
		envoy_head += strings.Replace(tmp, "{secrets_private_key}", private_key, 1)
	}

	err := ioutil.WriteFile("/etc/envoy.yaml", []byte(envoy_head), 0666)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
}
