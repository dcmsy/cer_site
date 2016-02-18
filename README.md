# cer_site

在互联网上，实现扫描二维码获取材料证明书功能

# 业务需求

1.满足用通过扫描产品包装二维码，获取材料证明书相关数据；
2.提供word在线下载功能；
3.实现与ERP业务数据同步；

# 产品架构

1.nginx + go（beego\docx\）+ mysql + redis
  
  nginx -- 提供高效反向代理服务，负载均衡；<br>
  go    -- 提供web服务；<br>
  mysql -- 同步ERP相关数据；<br>
  redis -- 作为数据缓存服务，提高响应速度，降低数据库压力；<br>

2.前端 jquery\ajax\css\html

