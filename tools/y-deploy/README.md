# Deploy Tool



## Commands

### k8s
1. 创建 `deployment` 配置文件
    ```
   y-deploy k8s create deployment -n account-service
   ```
    > 同时创建 `service`
    ```
   y-deploy k8s create deployment -n account-service --with-service
   ```

2. 替换配置文件中的环境变量
    ```
   y-deploy k8s config replace -e qa
   ```

3. 替换配置文件中的环境变量
    ```
   y-deploy k8s deployment -e qa
   ```

4. 修改 deployment 中的镜像
   ```shell
   y-deploy k8s deployment update-image -n golang-account-serivce
   ```

### docker image
1. 构建镜像
    ```
   y-deploy image build -n golang-account-serivce -d ./deploy/build/account-service/Dockerfile
   ```
2. 发布镜像
    ```
   y-deploy image publish -n golang-account-serivce
   ```
   
