# 复杂sql统计
 ## cloud-kylin
### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/CustomerMapper.xml
> sql统计:复杂0个，简单14个
>
> 使用的关键字：where,now,uuid,KYLIN_CUSTOMER,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/CustomerSecretKeyMapper.xml
> sql统计:复杂0个，简单6个
>
> 使用的关键字：now,uuid,KYLIN_CUSTOMER_SECRET_KEY,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/NavigationBarMapper.xml
> sql统计:复杂0个，简单14个
>
> 使用的关键字：MAX,KY_NAVIGATION_BAR,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/NotifyChannelDOMapper.xml
> sql统计:复杂0个，简单19个
>
> 使用的关键字：values,count,CJ_NOTIFY_CHANNEL

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/NotifyChannelImplDOMapper.xml
> sql统计:复杂0个，简单6个
>
> 使用的关键字：CJ_NOTIFY_CHANNEL_IMPL,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/NotifyConfigDOMapper.xml
> sql统计:复杂0个，简单22个
>
> 使用的关键字：values,and,count,CJ_NOTIFY_CONFIG

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/NotifyConfigTemplateDOMapper.xml
> sql统计:复杂0个，简单16个
>
> 使用的关键字：CJ_NOTIFY_CONFIG_TEMPLATE,values,count

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/NotifyHelpCategoryDOMapper.xml
> sql统计:复杂0个，简单17个
>
> 使用的关键字：where,CJ_NOTIFY_HELP_CATEGORY,values,count,MAX,size

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/NotifyHelpDOMapper.xml
> sql统计:复杂0个，简单22个
>
> 使用的关键字：CJ_NOTIFY_HELP,values,count,size,MAX,and,now,where

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/NotifyHelpRecordDOMapper.xml
> sql统计:复杂0个，简单19个
>
> 使用的关键字：count,CJ_NOTIFY_HELP_RECORD,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/NotifyNoticeDoMapper.xml
> sql统计:复杂3个，简单16个
>
> 使用的关键字：STRCMP,and,now,CJ_NOTIFY_NOTICE,in,count,JSON_VALID,date_format,json_set,values
#### getIdsByPublishWindow
```

        select
        <include refid="Base_Column_List"/>
        from CJ_NOTIFY_NOTICE
        where IS_DELETED = 0
        and TIME_WINDOW_ACTIVE = 1
        and STATUS = 'draft'
        and JSON_VALID(TIME_WINDOW)
        and TIME_WINDOW  is not null
        and STRCMP(TIME_WINDOW -> '$.endTime', TIME_WINDOW -> '$.startTime') >= 0
        and (TIME_WINDOW -> '$.startTime' &lt;= date_format(now(),'%Y-%m-%d %H:%i:%s'))
        and (TIME_WINDOW -> '$.endTime' > date_format(now(),'%Y-%m-%d %H:%i:%s'))
        order by id;
    
```
#### getIdsByUnpublishedWindow
```

        select
        <include refid="Base_Column_List"/>
        from CJ_NOTIFY_NOTICE
        where IS_DELETED = 0
        and TIME_WINDOW_ACTIVE = 1
        and STATUS = 'published'
        and JSON_VALID(TIME_WINDOW)
        and TIME_WINDOW  is not null
        and STRCMP(TIME_WINDOW -> '$.endTime', TIME_WINDOW -> '$.startTime') >= 0
        and TIME_WINDOW -> '$.endTime' &lt; date_format(now(),'%Y-%m-%d %H:%i:%s')
        order by id;
    
```
#### updateTimeWindowById
```

        update CJ_NOTIFY_NOTICE
        set TIME_WINDOW = json_set(TIME_WINDOW,'$.startTime', date_format(now(),'%Y-%m-%d %H:%i:%s'))
        where ID = #{id}
    
```
### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/NotifyTaskDOMapper.xml
> sql统计:复杂0个，简单13个
>
> 使用的关键字：STRCMP,JSON_CONTAINS,OR,CJ_NOTIFY_TASK,values,count,JSON_VALID,AND

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/NotifyTaskRecordDOMapper.xml
> sql统计:复杂0个，简单18个
>
> 使用的关键字：count,JSON_VALID,and,CJ_NOTIFY_TASK_RECORD,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/OrderAttrMapper.xml
> sql统计:复杂0个，简单1个
>
> 使用的关键字：KYLIN_ORDER_ATTR,VALUES,uuid

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/OrderMapper.xml
> sql统计:复杂0个，简单11个
>
> 使用的关键字：KYLIN_ORDER,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/ServiceCustomerMapper.xml
> sql统计:复杂0个，简单8个
>
> 使用的关键字：and,now,uuid,KYLIN_SERVICE_CUSTOMER,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/ServiceHealthCheckApiMapper.xml
> sql统计:复杂0个，简单3个
>
> 使用的关键字：KYLIN_SERVICE_HEALTH_CHECK_API,uuid

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/ServiceMapper.xml
> sql统计:复杂0个，简单8个
>
> 使用的关键字：KYLIN_SERVICE,values,and,now,uuid

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/ServiceOperateRecordMapper.xml
> sql统计:复杂0个，简单3个
>
> 使用的关键字：KYLIN_SERVICE_OPERATE_RECORD,values,uuid,now

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/StorageAttachmentMapper.xml
> sql统计:复杂0个，简单4个
>
> 使用的关键字：CONCAT,now,KY_STORAGE_ATTACHMENT,VALUES

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/StorageBucketMapper.xml
> sql统计:复杂0个，简单3个
>
> 使用的关键字：now,size,and,CONCAT,AND,KY_STORAGE_BUCKET,VALUES

### /Users/jiangliuhong/develop/cvs/idcos/cloud-kylin/cloudkylin-dal/src/main/resources/mapping/StorageOperateMapper.xml
> sql统计:复杂0个，简单1个
>
> 使用的关键字：KY_STORAGE_OPERATE_LOG,VALUES,now

