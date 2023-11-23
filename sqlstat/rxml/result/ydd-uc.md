# 复杂sql统计
 ## ydd-uc
### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcCompanyMapper.xml
> sql统计:复杂0个，简单5个
>
> 使用的关键字：now,UC_COMPANY,VALUES

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcConfigMapper.xml
> sql统计:复杂0个，简单4个
>
> 使用的关键字：now,UC_CONFIG,values

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcDepartmentMapper.xml
> sql统计:复杂0个，简单11个
>
> 使用的关键字：from,NOW,now,UC_DEPARTMENT,VALUES

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcDicMapper.xml
> sql统计:复杂0个，简单6个
>
> 使用的关键字：NOW,now,UC_DIC,VALUES

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcEmailLogMapper.xml
> sql统计:复杂0个，简单4个
>
> 使用的关键字：UC_EMAIL_LOG,VALUES,NOW,now

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcGroupMapper.xml
> sql统计:复杂0个，简单12个
>
> 使用的关键字：UC_GROUP,VALUES,and,count,NOW,now

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcGroupXroleMapper.xml
> sql统计:复杂0个，简单5个
>
> 使用的关键字：now,UC_GROUP_X_ROLE,VALUES,NOW

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcGroupXuserMapper.xml
> sql统计:复杂0个，简单7个
>
> 使用的关键字：NOW,now,UC_GROUP_X_USER,VALUES

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcHighRiskOperationControlMapper.xml
> sql统计:复杂0个，简单5个
>
> 使用的关键字：VALUES,where,and,NOW,now,UC_HIGH_RISK_OPERATION_CONTROL

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcLoginCaptchaMapper.xml
> sql统计:复杂0个，简单3个
>
> 使用的关键字：now,UC_LOGIN_CAPTCHA,VALUES

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcLoginLogMapper.xml
> sql统计:复杂2个，简单3个
>
> 使用的关键字：UC_LOGIN_LOG,VALUES,now,ifnull,concat,and,where,count
#### queryLoginList
```

        select l.ID, l.LOGIN_TYPE,ifnull(l.USER_NAME,u.USER_NAME) as USER_NAME, l.LOGIN_IP, l.CREATE_DATE ,d.NAME as DEPARTMENT_NAME,l.STATUS,m.CODE as MEMBER_CODE,
        case when ifnull(u.ACTIVE,0)=0 then concat(u.NAME,' (已删除)') when ifnull(u.LOCKED_OUT,0)=0 then u.NAME else concat(u.NAME,' (已冻结)') end as NAME
        from UC_LOGIN_LOG l left outer join UC_MEMBER m on l.MEMBER=m.ID or l.MEMBER_CODE=m.CODE
        left outer join UC_USER u on l.USER_ID=u.ID or l.USER_NAME=u.USER_NAME and u.MEMBER=m.ID
        left outer join UC_DEPARTMENT d on d.ID=u.DEPARTMENT
        where 1=1
        <if test="member > 0" >
            AND u.MEMBER = #{member}
        </if>
        <if test="searchKey!=null and searchKey!=''">
           and (l.USER_NAME like #{searchKey} or u.NAME like #{searchKey} or l.LOGIN_IP like #{searchKey} )
        </if>
        <if test="userName!=null and userName!=''">
           and l.USER_NAME like #{userName}
        </if>
        <if test="name!=null and name!=''">
           and u.NAME like #{name}
        </if>
        <if test="loginIp!=null and loginIp!=''">
           and l.LOGIN_IP like #{loginIp}
        </if>
        <if test="createDateStart!=null and createDateEnd!=null">
           and l.CREATE_DATE between #{createDateStart} and #{createDateEnd}
        </if>
        <if test="createDateStart!=null and createDateEnd==null">
           and l.CREATE_DATE &gt;= #{createDateStart}
        </if>
        <if test="createDateStart==null and createDateEnd!=null">
           and l.CREATE_DATE &lt;= #{createDateEnd}
        </if>
         order by ID desc
    
```
#### queryUserLoginList
```

        select l.ID, ifnull(l.USER_NAME,u.USER_NAME) as USER_NAME,l.LOGIN_IP, l.CREATE_DATE ,l.LOGIN_TYPE,l.STATUS,#{memberCode} as MEMBER_CODE,
        case when ifnull(u.ACTIVE,0)=0 then concat(u.NAME,' (已删除)') when ifnull(u.LOCKED_OUT,0)=0 then u.NAME else concat(u.NAME,' (已冻结)') end as NAME
        from UC_LOGIN_LOG l left outer join UC_USER u on u.ID=#{userId}
        where (l.MEMBER_CODE=#{memberCode} and l.USER_NAME=#{userName} or l.USER_ID=u.ID)
        <if test="member > 0" >
            AND l.MEMBER = #{member}
        </if>
        <if test="createDateStart!=null and createDateEnd!=null">
           and l.CREATE_DATE between #{createDateStart} and #{createDateEnd}
        </if>
        <if test="createDateStart!=null and createDateEnd==null">
           and l.CREATE_DATE &gt;= #{createDateStart}
        </if>
        <if test="createDateStart==null and createDateEnd!=null">
           and l.CREATE_DATE &lt;= #{createDateEnd}
        </if>
         order by l.ID desc
    
```
### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcLoginTicketMapper.xml
> sql统计:复杂0个，简单3个
>
> 使用的关键字：VALUES,NOW,now,UC_LOGIN_TICKET

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcMemberMapper.xml
> sql统计:复杂0个，简单23个
>
> 使用的关键字：size,and,distinct,VALUES,CONCAT,COUNT,count,NOW,now,UC_MEMBER

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcMemberTypeMapper.xml
> sql统计:复杂0个，简单7个
>
> 使用的关键字：NOW,now,UC_MEMBER_TYPE,VALUES

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcRoleGroupMapper.xml
> sql统计:复杂0个，简单5个
>
> 使用的关键字：NOW,now,UC_ROLE_GROUP,VALUES,and

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcRoleMapper.xml
> sql统计:复杂0个，简单9个
>
> 使用的关键字：NOW,now,UC_ROLE,VALUES,count

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcSmsLogMapper.xml
> sql统计:复杂0个，简单1个
>
> 使用的关键字：UC_SMS_LOG,VALUES,now

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcUserFindPwdMapper.xml
> sql统计:复杂0个，简单6个
>
> 使用的关键字：NOW,now,UC_USER_FIND_PWD,VALUES

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcUserMapper.xml
> sql统计:复杂3个，简单31个
>
> 使用的关键字：size,CONVERT,exists,date,NOW,DATEDIFF,VALUES,count,and,join,ifnull,group_concat,now,UC_USER,where
#### findUserVo
```

        select
        <include refid="UserVO_Query"/> ,ug.USERGROUP_NAME
        from UC_USER u
        inner join UC_MEMBER m on u.MEMBER=m.ID
        left outer join UC_DEPARTMENT dept on dept.ID=u.DEPARTMENT and dept.ACTIVE = 1
        left outer join (select tx.USER_ID,group_concat(tg.NAME) USERGROUP_NAME from UC_GROUP_X_USER tx,UC_GROUP tg where tx.GROUP_ID=tg.ID and tg.ACTIVE=1 <if test="member>0">and tg.`MEMBER`= #{member}</if> group by tx.USER_ID) ug on ug.USER_ID=u.ID
        <if test="groupId !=null  and groupId>0" >
        inner join UC_GROUP_X_USER gx on gx.USER_ID=u.ID and gx.GROUP_ID =#{groupId}
        </if>
        where m.ACTIVE = 1 and u.ACTIVE = 1
        <if test="member>0">and m.ID=#{member}</if>
        <if test="userName != null and userName != ''">and u.USER_NAME = #{userName}</if>
        <if test="userNameLike != null and userNameLike != ''">and u.USER_NAME like #{userNameLike}</if>
        <if test="name != null and name != ''">and u.NAME = #{name}</if>
        <if test="nameLike != null and nameLike != ''">and u.NAME like #{nameLike}</if>
        <if test="email != null and email != ''">and u.EMAIL = #{email}</if>
        <if test="departments != null and departments.size()>0"> 
            and u.DEPARTMENT IN <foreach collection="departments" item="id" open="(" separator="," close=")">#{id}</foreach>
        </if>
        <if test="mobilePhone != null and mobilePhone != ''">and u.MOBILE_PHONE = #{mobilePhone}</if>
        <if test="searchKey != null and searchKey != ''">
            and (u.NAME like #{searchKey} or u.USER_NAME like #{searchKey} or u.EMAIL like #{searchKey} or
            u.MOBILE_PHONE like #{searchKey})
        </if>
        <if test="lockedOut!=null"> and u.LOCKED_OUT=#{lockedOut}</if>
        <if test="online!=null and online"> and u.ONLINE_SESSION &gt; 0</if>
        <if test="online!=null and !online"> and ifnull(u.ONLINE_SESSION,0) = 0</if>
        <if test="orderBy == null">order by u.CREATE_DATE DESC</if>
        <if test="orderBy != null">order by ${orderBy} <if test="direction != null">${direction}</if></if>
    
```
#### findSubUserVo
```

        select
        <include refid="UserVO_Query"/>
        from UC_USER u
        inner join UC_MEMBER m on u.MEMBER=m.ID and m.ACTIVE = 1 and u.ACTIVE = 1
        left outer join UC_DEPARTMENT dept on dept.ID=u.DEPARTMENT and dept.ACTIVE = 1
        <if test="groupId !=null  and groupId>0" >
        inner join UC_GROUP_X_USER gx on gx.USER_ID=u.ID and gx.GROUP_ID =#{groupId}
        </if>
        where 1=1
        <if test="member>0">and (m.ID=#{member} or m.PARENT=#{member})</if>
        <if test="userName != null and userName != ''">and u.USER_NAME = #{userName}</if>
        <if test="userNameLike != null and userNameLike != ''">and u.USER_NAME like #{userNameLike}</if>
        <if test="name != null and name != ''">and u.NAME = #{name}</if>
        <if test="nameLike != null and nameLike != ''">and u.NAME like #{nameLike}</if>
        <if test="email != null and email != ''">and u.EMAIL = #{email}</if>
        <if test="departments != null and departments.size()>0"> 
            and u.DEPARTMENT IN <foreach collection="departments" item="id" open="(" separator="," close=")">#{id}</foreach>
        </if>
        <if test="mobilePhone != null and mobilePhone != ''">and u.MOBILE_PHONE = #{mobilePhone}</if>
        <if test="searchKey != null and searchKey != ''">
            and (u.NAME like #{searchKey} or u.USER_NAME like #{searchKey} or u.EMAIL like #{searchKey} or
            u.MOBILE_PHONE like #{searchKey} or dept.NAME like #{searchKey} or m.NAME like #{searchKey})
        </if>
        <if test="lockedOut!=null"> and u.LOCKED_OUT=#{lockedOut}</if>
        <if test="online!=null and online"> and u.ONLINE_SESSION &gt; 0</if>
        <if test="online!=null and !online"> and ifnull(u.ONLINE_SESSION,0) = 0</if>
        order by CONVERT(u.NAME USING gbk )
    
```
#### findExpReminder
```

        select <include refid="Base_Column_List"/> from UC_USER u inner join UC_MEMBER m on u.MEMBER=m.ID and u.ACTIVE = 1 and m.ACTIVE = 1
        where m.STATUS='ENABLED' and date(u.LAST_MODIFIED_PWD) =#{expReminder} order by u.LAST_MODIFIED_PWD,u.ID
    
```
### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcUserPwdChangelog.xml
> sql统计:复杂0个，简单2个
>
> 使用的关键字：UC_USER_PWD_CHANGELOG,VALUES,now

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcUserSessionMapper.xml
> sql统计:复杂0个，简单8个
>
> 使用的关键字：now,and,NOW,UC_USER_SESSION,VALUES

### /Users/jiangliuhong/develop/cvs/idcos/ydd-uc/ydd-uc-dal/src/main/resources/META-INF/mybatis/sqlmap/UcUserXroleMapper.xml
> sql统计:复杂0个，简单7个
>
> 使用的关键字：and,count,NOW,now,UC_USER_X_ROLE,VALUES

