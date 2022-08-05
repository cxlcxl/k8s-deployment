> go-zero 项目 k8s 部署步骤
1. 进入服务目录：rbac
    1. 配置文件服务发现方式改成 target
2. 生成 dockerfile：goctl docker -go rbac.go
3. 切换到项目根路径制作镜像：docker build -f xxxxxxx/Dockerfile -t xx:xx .
4. 创建 k8s ServiceAccount 使用到的命名空间：kubectl create ns 命名空间名称
5. 创建 k8s ServiceAccount：kubectl apply -f sa.yaml [sa.yaml 在 k8s 目录中]
6. 进入服务目录生成 rpc K8S yaml 文件：
    - goctl kube deploy -replicas 3 -requestCpu 100 -requestMem 50 -limitCpu 200 -limitMem 100 -name rbac-rpc -namespace business -image rbac-rpc:v1.0.0 -o rbac-rpc.yaml -port 9001 --serviceAccount find-endpoints
7. 进入服务目录生成 api K8S yaml 文件：
    - goctl kube deploy --nodePort 30010 -replicas 3 -requestCpu 100 -requestMem 50 -limitCpu 200 -limitMem 100 -name rbac-api -namespace business -image rbac-api:v1.0.0 -o rbac-api.yaml -port 18001 --serviceAccount find-endpoints
    ***—nodePort K8S 限制端口在 30000-? 之间*** 