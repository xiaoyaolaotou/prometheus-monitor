1、通过配置文件获取promethues监控数据

2、获取字段
    - ip
    - hostname
    - uptime
    - Total_Mem
    - Total_Cpu
    - CPU used  
    - Memory used 
    - /根分区使用率

3、可以自定义时间，默认都是7d 平均数据

4、执行方式
```
go build -o xxx
./xxx -f prometheus.yml
```

5、 QPL 公式
```
node_uname_info  查询主机名

sum(time() - node_boot_time_seconds)by(instance)  运行时间

node_memory_MemTotal_bytes   总内存

count(node_cpu_seconds_total) by (instance)  总核数

(1 - avg(rate(node_cpu_seconds_total{mode="idle"}[7d])) by (instance)) * 100  cpu 平均使用率

100 * (1 - ((avg_over_time(node_memory_MemFree_bytes[24h]) + avg_over_time(node_memory_Cached_bytes[24h]) + avg_over_time(node_memory_Buffers_bytes[24h])) / avg_over_time(node_memory_MemTotal_bytes[24h])))  mem使用率

max((node_filesystem_size_bytes{fstype=~"ext.?|xfs",mountpoint=~"/data.*|/web"}-node_filesystem_free_bytes{fstype=~"ext.?|xfs",mountpoint=~"/data.*|/web"}) *100/(node_filesystem_avail_bytes {fstype=~"ext.?|xfs"}+(node_filesystem_size_bytes{fstype=~"ext.?|xfs",mountpoint=~"/data.*|/web"}-node_filesystem_free_bytes{fstype=~"ext.?|xfs",mountpoint=~"/data.*|/web"})))by(instance,mountpoint)  磁盘使用率
```