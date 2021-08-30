#!/bin/bash

## kill redis进程 
ps -ef | grep redis-server | grep -v grep | awk '{print $2}' | xargs kill -9

echo "-----------------设置redis主从集群配置--------------------"

if [ -z $1 ];then
echo "节点角色不可为空"
exit
fi

if [ ! -e /usr/local/redis-4.0.5/redis.log ];then
echo "正在生成日志文件"
touch /usr/local/redis-5.0.5/redis.log
fi

if [ ! -d /usr/local/redis-5.0.5/redisworkplace ];then
echo "正在生成工作目录"
mkdir /usr/local/redis-5.0.5/redisworkplace
fi

sed -i 's/bind 127.0.0.1/# bind 127.0.0.1/g' /usr/local/redis-5.0.5/redis.conf
sed -i 's/daemonize no/daemonize yes/g' /usr/local/redis-5.0.5/redis.conf
sed -i 's/# requirepass foobared/requirepass 123456/g' /usr/local/redis-5.0.5/redis.conf
sed -i 's/logfile ""/logfile "\/usr\/local\/redis-5.0.5\/reids.log"/g' /usr/local/redis-5.0.5/redis.conf
sed -i 's/dir \.\//dir "\/usr\/local\/redis-5.0.5\/redisworkplace"/g' /usr/local/redis-5.0.5/redis.conf
sed -i 's/# masterauth <master-password>/masterauth 123456/g' /usr/local/redis-5.0.5/redis.conf

if [ "$1" = "master" ];then
echo "主节点设置完成"
elif [ "$1" = "slave" ];then
if [ -z  $2 ];then
echo "master ip不可为空"
exit
fi
sed -i 's/# replicaof <masterip> <masterport>/replicaof '$2' 6379/g' /usr/local/redis-5.0.5/redis.conf
echo "从节点设置完成, $2"
else
echo "参数错误"
fi

echo "正在启动redis"
rm -rf /var/run/redis_6379.pid
/usr/local/redis-5.0.5/bin/redis-server /usr/local/redis-5.0.5/redis.conf
echo "--------------启动成功--------------"
