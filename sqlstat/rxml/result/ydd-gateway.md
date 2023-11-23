# 复杂sql统计
 ## ydd-gateway
### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/AdaptorInterfaceMapper.xml
> sql统计:复杂0个，简单5个
>
> 使用的关键字：now,GW_ADAPTOR_INTERFACE,values,concat

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/AdaptorMethodMapper.xml
> sql统计:复杂0个，简单6个
>
> 使用的关键字：now,GW_ADAPTOR_METHOD,values

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/AdaptorParameterMapper.xml
> sql统计:复杂0个，简单2个
>
> 使用的关键字：GW_ADAPTOR_PARAMETER,values,now

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/InvokeApiMapper.xml
> sql统计:复杂0个，简单7个
>
> 使用的关键字：values,CONCAT,now,GW_INVOKE_API

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/InvokeConfigMapper.xml
> sql统计:复杂1个，简单22个
>
> 使用的关键字：GW_INVOKE_CONFIG,and,max,size,concat,ifnull,now,CONCAT,group_concat,values
#### listCheckByApiId
```

		select c.ID,c.SERVICE,m.INTERFACE_ID,m.NAME method,c.ENDPOINT,c.ROUTER,c.TEMPLATE,c.INPUT_CONVERTER,c.OUTPUT_CONVERTER,c.VERSION,c.STATUS,group_concat(distinct  p.PARAM_NAME,'-',p.PARAM_TYPE,ifnull(p.DEFAULT_VALUE,''),p.REQUIRED,p.ESCAPE order by p.PARAM_NAME) as params,group_concat(distinct invokeover.PARAM_NAME order by invokeover.PARAM_NAME) as overrides
		from GW_INVOKE_CONFIG c
		  left outer join GW_INVOKE_OVERRIDE invokeover on invokeover.CONFIG_ID=c.ID
		  left outer join GW_ADAPTOR_METHOD m on m.ID=c.METHOD_ID and m.ACTIVE=1 left outer join GW_INVOKE_PARAMETER p on c.ID=p.CONFIG_ID and p.ACTIVE=1
		  left outer join GW_ADAPTOR_PARAMETER ap on ap.ID=p.PARAMETER_ID and ap.ACTIVE=1 where c.API_ID=#{apiId}
		group by c.ID,c.SERVICE,m.INTERFACE_ID,m.NAME,c.ENDPOINT,c.ROUTER,c.TEMPLATE,c.INPUT_CONVERTER,c.OUTPUT_CONVERTER,c.VERSION,c.STATUS;
	
```
### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/InvokeOverrideMapper.xml
> sql统计:复杂0个，简单8个
>
> 使用的关键字：IN,now,GW_INVOKE_OVERRIDE,values

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/InvokeParameterMapper.xml
> sql统计:复杂0个，简单8个
>
> 使用的关键字：IN,now,GW_INVOKE_PARAMETER,values

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/InvokeRecordMapper.xml
> sql统计:复杂0个，简单10个
>
> 使用的关键字：max,count,now,GW_INVOKE_RECORD,values

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/MidServerGroupAndNetworkMapper.xml
> sql统计:复杂0个，简单7个
>
> 使用的关键字：GW_MIDSERVER_GROUP_NETWORK,values,now

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/MsgMapper.xml
> sql统计:复杂0个，简单4个
>
> 使用的关键字：GW_MSG,values,now,max,count

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/PackRecordMapper.xml
> sql统计:复杂0个，简单3个
>
> 使用的关键字：values,now,where,concat,GW_PACK_RECORD

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/ProductMapper.xml
> sql统计:复杂0个，简单9个
>
> 使用的关键字：GW_PRODUCT,EXISTS,now

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/ProviderMapper.xml
> sql统计:复杂0个，简单6个
>
> 使用的关键字：concat,now,GW_PROVIDER,EXISTS

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/QosConfigMapper.xml
> sql统计:复杂0个，简单8个
>
> 使用的关键字：CONCAT,now,GW_QOS_CONFIG,values

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/QosTaskMapper.xml
> sql统计:复杂0个，简单9个
>
> 使用的关键字：in,max,now,GW_QOS_TASK,values

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/QosTaskRecordMapper.xml
> sql统计:复杂0个，简单10个
>
> 使用的关键字：count,NOW,max,values,in,and,date_add,now,GW_QOS_TASK_RECORD

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/RegionMapper.xml
> sql统计:复杂0个，简单2个
>
> 使用的关键字：

### /Users/jiangliuhong/develop/cvs/idcos/ydd-gateway/ydd-gateway-dal/src/main/resources/META-INF/mybatis/sqlmap/RouteConfigMapper.xml
> sql统计:复杂0个，简单5个
>
> 使用的关键字：now,GW_ROUTE_CONFIG,values,max

