#!/bin/bash

echo "--------------配置Sentinel------------------"
if [ ! -d /usr/local/redis-5.0.5/sentinel ];then
echo "正在创建工作目录"
mkdir /usr/local/redis-5.0.5/sentinel
fi

if [ ! -e /usr/local/redis-5.0.5/sentinel.log ];then
echo "正在创建日志文件"
touch /usr/local/redis-5.0.5/sentinel.log
fi

if [ -z $1 ];then
echo "erro: 主节点ip不可为空"
exit
fi

sed -i 's/daemonize no/daemonize yes/g' /usr/local/redis-5.0.5/sentinel.conf
sed -i 's/logfile ""/logfile "\/usr\/local\/redis-5.0.5\/sentinel.log"/g' /usr/local/redis-5.0.5/sentinel.conf
sed -i 's/dir \/tmp/dir \/usr\/local\/redis-5.0.5\/sentinel/g' /usr/local/redis-5.0.5/sentinel.conf
sed -i 's/sentinel monitor mymaster 127.0.0.1 6379 2/sentinel monitor mymaster '$1' 6379 2/g' /usr/local/redis-5.0.5/sentinel.conf
sed -i 's/# sentinel auth-pass <master-name> <password>/sentinel auth-pass mymaster 123456/g' /usr/local/redis-5.0.5/sentinel.conf
sed -i 's/sentinel down-after-milliseconds mymaster 30000/sentinel down-after-milliseconds mymaster 3000/g' /usr/local/redis-5.0.5/sentinel.conf
sed -i 's/# sentinel failover-timeout <master-name> <milliseconds>/sentinel failover-timeout mymaster 5000/g' /usr/local/redis-5.0.5/sentinel.conf


echo "Sentinel配置完毕"
echo "正在启动Sentinel"
rm -rf /var/run/redis-sentinel.pid
/usr/local/redis-5.0.5/bin/redis-sentinel /usr/local/redis-5.0.5/sentinel.conf
echo "-------------启动成功---------------"

