# nginx 主要有两个用途
1. 作为一个转发和负载均衡的中转站
2. 给用户提供存储图片的途径

作为图片存储
用户的图片存储在服务器才能被访问，而oss却需要额外购买
基本原理，docker可以有一个地址映射功能，这样就能利用docker容器来在本地存储和管理图片，并且能被访问到
docker run -d -v /home/yang/temp/picture:/img -p 80:80 --name testNginx  nginx

作为转发
在修改完nginx的配置文件之后，可以做出如下规则定义
img开头的请求视为图片，从本地拿取
api开头的请求转发给golang 程序
其他请求，返回错误页面

