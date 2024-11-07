# postgresql的操作文档

## 进入psql
docker run -d --name postgreYoga -p 5432:5432 -e      POSTGRES_PASSWORD=3411    postgres
创建完成postgresql之后进入容器
docker exec -it postgreYoga bash
进入了容器需要使用psql
psql -U postgres
**这里是一个初始的用户,拥有一切权限**,可以使用\du 查看用户列表和权限列表
CREATE USER yang WITH PASSWORD '3411';
这个可以创建用户
ALTER ROLE yang WITH SUPERUSER;
 这个可以添加用户到超级用户里

## 创建数据库
按下\q可以推出这个psql
重新使用yang用户来进入psql
psql -U yang postgres
\l 
查看所有的数据库
\c 
切换目前所在的数据库
create database yoga;
记得要;这个没有的话不会报错,这样就创建好了一个数据库
\c yoga
这样就能看到一些输出,能知道现在在以yang用户操作yoga数据库了.
\d teachers; 
可以查看数据库的结构

## 创建一个系统用户给程序使用
最关键的就是权限,肯定不能**superuser** 不然如果不小心被sql注入了或者别的意料之外的攻击或者意外就是超级大的损失   
所以给到的权限都应该是系统满足所有操作的最小权限  
这里给出所有的权限
1. SUPERUSER | NOSUPERUSE 这个是超级用户,基本不用考虑
2. CREATEDB | NOCREATEDB
3. INHERIT | NOINHERIT 有一个组的概念,这个就是是否能继承组的权限,没什么用,为什么不直接移出组呢?
4. BYPASSRLS | NOBYPASSRLS rls是一个防护策略,这个就是说是否可以绕过这个策略
5. CONNECTION LIMIT connlimit 这个角色的最大连接数
6. [ ENCRYPTED ] PASSWORD 'password' | PASSWORD NULL 这个是密码,关系不大
7. VALID UNTIL 'timestamp' 这个是一个有效期
8. IN ROLE role_name [, ...] 新版的用户组,可以理解为role是一个用户也可以是用户组
