# 复杂sql统计
 ## cloud-catalog
### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/ApprovalRuleDOMapper.xml
> sql统计:复杂0个，简单7个
>
> 使用的关键字：CJ_C2_APPROVAL_RULE,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/CatalogDOMapper.xml
> sql统计:复杂0个，简单16个
>
> 使用的关键字：and,count,CJ_C2_CATALOG,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/CatalogItemApprovalDOMapper.xml
> sql统计:复杂0个，简单9个
>
> 使用的关键字：CJ_C2_CATALOG_ITEM_APPROVAL,values,and

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/CatalogItemAuthDOMapper.xml
> sql统计:复杂0个，简单13个
>
> 使用的关键字：count,size,CJ_C2_CATALOG_ITEM_AUTH,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/CatalogItemContextDOMapper.xml
> sql统计:复杂0个，简单10个
>
> 使用的关键字：CJ_C2_CATALOG_ITEM_CONTEXT,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/CatalogItemDOMapper.xml
> sql统计:复杂10个，简单79个
>
> 使用的关键字：and,size,convert,or,now,DISTINCT,CAST,concat,values,isnull,CJ_C2_CATALOG_ITEM,count,join,MAX,in,from,max,COUNT
#### queryAdminPageList
```

        select <include refid="CatalogItem_With_Category_Column_List"/>
        from CJ_C2_CATALOG_ITEM item
        join ( select max(ID) ID, MEMBER_ID from CJ_C2_CATALOG_ITEM where IS_DELETED = 0
            <if test="category != null">and CATEGORY = #{category}</if>
            and IS_ACTIVE = 1 group by code, MEMBER_ID
            union
            select max(ID), MEMBER_ID ID from CJ_C2_CATALOG_ITEM where IS_DELETED = 0
            <if test="category != null">and CATEGORY = #{category}</if>
            group by code, MEMBER_ID having max(IS_ACTIVE) = 0
        ) fitered on item.ID = fitered.ID and item.MEMBER_ID = fitered.MEMBER_ID
        join CJ_C2_CATEGORY category on item.CATEGORY = category.ID
        left join CJ_C2_CATALOG_ITEM_CONTEXT context on item.ID = context.CATALOG_ITEM
        <if test="parentMemberId != null">
            <if test="isFiled == null or isFiled != 3">
                <!-- 子租户查询用 -->
                join
                CJ_C2_CATALOG_ITEM_AUTH auth
                on
                item.CODE = auth.CATALOG_ITEM
                and
                auth.TYPE = 'subtenant'
                and
                auth.VALUE = #{memberId}
                and
                auth.IS_DELETED = 0
                or
                item.SYSTEMED = 1
            </if>
        </if>
        where item.IS_DELETED = 0
        and category.IS_ACTIVE = 1
        <if test="shared != null">
            and item.SHARED = #{shared}
        </if>
        <if test="ids != null and ids.size() > 0">
            and
            item.ID in
            <foreach collection="ids" item="id" index="index" open="(" close=")" separator=",">
                #{id,jdbcType=BIGINT}
            </foreach>
        </if>
        <!--<if test="category != null">and CAST(item.CATEGORY as char) = CAST(#{category} as char)</if>-->
        <if test="versionType != null">
            <if test="versionType == 'draft'">and item.IS_ACTIVE = 0 and item.VERSION is NULL</if>
            <if test="versionType == 'offline'">and item.IS_ACTIVE = 0 and item.VERSION is NOT NULL</if>
            <if test="versionType == 'online'">and item.IS_ACTIVE = 1 and item.VERSION is NOT NULL</if>
        </if>
        <if test="name != null">and item.NAME like "%"#{name}"%"</if>
        <choose>
            <when test="parentMemberId != null">
                <!-- 子租户查询用 -->
                and (item.SYSTEMED = 1 OR item.MEMBER_ID = #{parentMemberId})
            </when>
            <otherwise>
                <choose>
                    <when test="isFiled != null and isFiled == 4">
                        and item.MEMBER_ID = #{memberId}
                        and item.SYSTEMED != 1
                    </when>
                    <otherwise>
                        <if test="memberId != null">and (item.SYSTEMED = 1 OR item.MEMBER_ID = #{memberId})</if>
                    </otherwise>
                </choose>
            </otherwise>
        </choose>
        <if test="catalog != null">and category.CATALOG = #{catalog} </if>
        <if test="catalogs != null and catalogs.size() > 0">
            and category.CATALOG IN
            <foreach collection="catalogs" item="id" index="index" open="(" close=")" separator=",">
                #{id,jdbcType=BIGINT}
            </foreach>
        </if>
        group by
        <include refid="CatalogItem_With_Category_Column"/>
        <if test="sortColumn != null and sortColumn != 'CATEGORY_TITLE' and sortColumn != 'NAME' and sortDirection != null">
            order by item.${sortColumn} ${sortDirection}
        </if>
        <if test="sortColumn == 'CATEGORY_TITLE' and sortDirection != null">order by convert(category.TITLE using gbk)
            ${sortDirection}
        </if>
        <if test="sortColumn == 'NAME' and sortDirection != null">order by convert(item.NAME using gbk)
            ${sortDirection}
        </if>
        <if test="sortColumn == null or sortDirection == null">order by item.LAST_MODIFIED_DATE desc</if>
        <if test="page != null and pageSize != null">limit #{page}, #{pageSize}</if>
    
```
#### queryUserPageList
```

        select
        <include refid="CatalogItem_With_Category_Column_List"/>
        from CJ_C2_CATALOG_ITEM item
        join (
        select max(ID) ID, MEMBER_ID from CJ_C2_CATALOG_ITEM where IS_DELETED = 0 and IS_ACTIVE = 1 group by MEMBER_ID,CODE
        ) fitered on item.ID = fitered.ID and item.MEMBER_ID = fitered.MEMBER_ID
        join CJ_C2_CATEGORY category on item.CATEGORY = category.ID
        left join CJ_C2_CATALOG_ITEM_CONTEXT context on item.ID = context.CATALOG_ITEM
        <if test="sortColumn == 'RECENT_USED'">
            left join (
            select MAX(reqItem.ID) as lastItem, reqItem.CAT_ITEM as CAT_ITEM from CJ_C2_REQUESTED_ITEM reqItem
            join CJ_C2_CATALOG_ITEM item on reqItem.CAT_ITEM = item.ID
            where item.IS_DELETED = 0
            and item.IS_ACTIVE = 1
            <if test="memberId != null">and item.MEMBER_ID = #{memberId}</if>
            group by reqItem.CAT_ITEM) b on item.ID = b.CAT_ITEM
        </if>
        <if test="sortColumn == 'OFTEN_USED'">
            left join (
            select count(reqItem.ID) as oftenItem, reqItem.CAT_ITEM as CAT_ITEM from CJ_C2_REQUESTED_ITEM reqItem
            join CJ_C2_CATALOG_ITEM item on reqItem.CAT_ITEM = item.ID
            where item.IS_DELETED = 0
            and item.IS_ACTIVE = 1
            <if test="memberId != null">and item.MEMBER_ID = #{memberId}</if>
            group by reqItem.CAT_ITEM) c on item.ID = c.CAT_ITEM
        </if>
        where item.IS_DELETED = 0
        and category.CATALOG > 0
        and category.IS_ACTIVE = 1
        <if test="ids != null and ids.size() > 0">
            and item.ID in
            <foreach collection="ids" item="id" index="index" open="(" close=")" separator=",">
                #{id,jdbcType=BIGINT}
            </foreach>
        </if>
        <choose>
            <!-- 无筛选条件,可查看内置服务项 -->
            <!-- 或者选择内置服务项类别带，也可以查看内置服务项 -->
            <when test="isFiled == 0">
                <if test="memberId == 0">
                    and item.VISIBLE = 1
                </if>
                <if test="memberId != 0">
                    AND item.VISIBLE = 1
                    and (item.MEMBER_ID = #{memberId} or item.SYSTEMED = 1)
                </if>
            </when>
            <!-- 有筛选条件 -->
            <when test="isFiled == 1">
                and (item.MEMBER_ID = #{memberId} or item.SYSTEMED = 1)
                AND item.VISIBLE = 1
                <if test="name != null">
                    and item.NAME like "%"#{name}"%" or (item.NAME like "%"#{name}"%" and item.SYSTEMED =1 and
                    item.VISIBLE = 1)
                </if>
                <if test="source != null">
                    and item.source = #{source}
                </if>
            </when>
            <!-- 带内置服务项类别带筛选条件，允许用户查看内置服务项信息 此时isFiled值为2 -->
            <otherwise>
                and item.SYSTEMED = 1
                and item.VISIBLE = 1
                <if test="name != null">AND item.NAME like "%"#{name}"%"</if>
                <if test="source != null">
                    and item.source = #{source}
                </if>
            </otherwise>
        </choose>
        <if test="catalog != null">and category.CATALOG = #{catalog} </if>
        <if test="catalogs != null and catalogs.size() > 0">
            and category.CATALOG IN
            <foreach collection="catalogs" item="id" index="index" open="(" close=")" separator=",">
                #{id,jdbcType=BIGINT}
            </foreach>
        </if>
        <if test="approvalMethod != null">AND item.APPROVAL_METHOD = #{approvalMethod}</if>
        <if test="isActive != null">and item.IS_ACTIVE = #{isActive}</if>
        <if test="sortColumn == 'RECENT_USED'">order by b.lastItem desc</if>
        <if test="sortColumn == 'OFTEN_USED'">order by c.oftenItem desc</if>
        <if test="sortColumn == 'LATEST_PUBLISHED'">order by item.LAST_MODIFIED_DATE desc</if>
        <if test="sortColumn == 'NORMAL'">order by convert(item.NAME using gbk) asc</if>
        <if test="sortColumn == 'CATEGORY_TITLE' and sortDirection != null">order by convert(category.TITLE using gbk)
            ${sortDirection}
        </if>
        <if test="sortColumn == 'NAME' and sortDirection != null">order by convert(item.NAME using gbk)
            ${sortDirection}
        </if>
        <if test="sortColumn == null">order by convert(item.NAME using gbk) asc</if>
        <if test="page != null and pageSize != null">limit #{page}, #{pageSize}</if>
    
```
#### querySimpleList
```

        select item.ID, item.NAME, item.STATE, item.SHORT_DESCRIPTION, item.DESCRIPTION, item.ICON, item.ICON_TYPE,
        context.OPERATION_TYPE
        from CJ_C2_CATALOG_ITEM item
        join (
        select max(ID) ID from CJ_C2_CATALOG_ITEM where IS_DELETED = 0 and IS_ACTIVE = 1 group by code
        ) fitered on item.ID = fitered.ID
        left join CJ_C2_CATALOG_ITEM_CONTEXT context on item.ID = context.CATALOG_ITEM
        <if test="sortColumn == 'RECENT_USED'">
            left join (
            select MAX(reqItem.ID) as lastItem, reqItem.CAT_ITEM as CAT_ITEM from CJ_C2_REQUESTED_ITEM reqItem
            join CJ_C2_CATALOG_ITEM item on reqItem.CAT_ITEM = item.ID
            where item.IS_DELETED = 0
            and item.IS_ACTIVE = 1
            <if test="memberId != null">and item.MEMBER_ID = #{memberId}</if>
            group by reqItem.CAT_ITEM) b on item.ID = b.CAT_ITEM
        </if>
        <if test="sortColumn == 'OFTEN_USED'">
            left join (
            select count(reqItem.ID) as oftenItem, reqItem.CAT_ITEM as CAT_ITEM from CJ_C2_REQUESTED_ITEM reqItem
            join CJ_C2_CATALOG_ITEM item on reqItem.CAT_ITEM = item.ID
            where item.IS_DELETED = 0
            and item.IS_ACTIVE = 1
            <if test="memberId != null">and item.MEMBER_ID = #{memberId}</if>
            group by reqItem.CAT_ITEM) c on item.ID = c.CAT_ITEM
        </if>
        where item.IS_DELETED = 0
        and ((item.SYSTEMED = 0
        <if test="memberId != null">and item.MEMBER_ID = #{memberId}</if>
        <if test="ids != null and ids.size() > 0">
            and item.ID in
            <foreach collection="ids" item="id" index="index" open="(" close=")" separator=",">
                #{id,jdbcType=BIGINT}
            </foreach>
        </if>
        ) or (item.SYSTEMED = 1 and item.VISIBLE = 1))
        <if test="tagsFilter != null and tagsFilter.size() > 0">
            and item.ID in
            <foreach collection="tagsFilter" item="tagId" index="index" open="(" close=")" separator=",">
                #{tagId,jdbcType=BIGINT}
            </foreach>
        </if>
        <if test="category != null">and item.CATEGORY = #{category}</if>
        <if test="isActive != null">and item.IS_ACTIVE = #{isActive}</if>
        <if test="name != null">and item.NAME like "%"#{name}"%"</if>
        <if test="shortDescription != null">and item.SHORT_DESCRIPTION like "%"#{shortDescription}"%"</if>
        <if test="description != null">and item.DESCRIPTION like "%"#{description}"%"</if>
        <if test="source != null">and item.SOURCE like "%"#{source}"%"</if>
        <if test="timeWindow != null">item.TIME_WINDOW like "%"#{timeWindow}"%"</if>
        <if test="contextId != null">and context.CONTEXT_ID = #{contextId}</if>
        <if test="createBy != null">and item.CREATE_BY like "%"#{createBy}"%"</if>
        <if test="createDate != null">and item.CREATE_DATE like "%"#{createDate}"%"</if>
        <if test="lastModifiedBy != null">and item.LAST_MODIFIED_BY like "%"#{lastModifiedBy}"%"</if>
        <if test="lastModifiedDate != null">and item.LAST_MODIFIED_DATE like "%"#{lastModifiedDate}"%"</if>
        <if test="operationType != null">and context.OPERATION_TYPE = #{operationType}</if>
        <if test="sortColumn == 'RECENT_USED'">order by b.lastItem desc</if>
        <if test="sortColumn == 'OFTEN_USED'">order by c.oftenItem desc</if>
        <if test="sortColumn == 'LATEST_PUBLISHED'">order by item.LAST_MODIFIED_DATE desc</if>
        <if test="sortColumn == 'NORMAL'">order by convert(item.NAME using gbk) asc</if>
        <if test="sortColumn == null">order by convert(item.NAME using gbk) asc</if>
        <if test="page != null and pageSize != null">limit #{page}, #{pageSize}</if>
    
```
#### countListThird
```

        select count(1) total
        from CJ_C2_CATALOG_ITEM item
        join (
        select max(ID) ID from CJ_C2_CATALOG_ITEM where IS_DELETED = 0 and IS_ACTIVE = 1 group by code
        ) fitered on item.ID = fitered.ID
        where item.IS_DELETED = 0 and item.IS_ACTIVE = 1 and item.MEMBER_ID = #{memberId}
        <if test="category != null">
            and item.CATEGORY = #{category}
        </if>
        <if test="name != null">
            and item.NAME like concat(concat('%',#{name}),'%')
        </if>
        <if test="ids != null and ids.size() > 0">
            and item.ID in
            <foreach collection="ids" item="id" index="index" open="(" close=")" separator=",">
                #{id,jdbcType=BIGINT}
            </foreach>
        </if>
        <if test="sortColumn != null and sortColumn != 'NAME' and sortDirection != null">
            order by item.${sortColumn} ${sortDirection}
        </if>
        <if test="sortColumn == 'NAME' and sortDirection != null">
            order by convert(item.NAME using gbk) ${sortDirection}
        </if>
        <if test="sortColumn == null or sortDirection == null">
            order by convert(item.NAME using gbk) asc
        </if>
    
```
#### queryListThird
```

        select item.ID, item.NAME, item.DESCRIPTION, item.ICON, item.ICON_TYPE, item.LAST_MODIFIED_DATE, item.CODE,
        item.VERSION
        from CJ_C2_CATALOG_ITEM item
        join (
        select max(ID) ID from CJ_C2_CATALOG_ITEM where IS_DELETED = 0 and IS_ACTIVE = 1 group by code
        ) fitered on item.ID = fitered.ID
        where item.IS_DELETED = 0 and item.IS_ACTIVE = 1 and item.MEMBER_ID = #{memberId}
        <if test="category != null">
            and item.CATEGORY = #{category}
        </if>
        <if test="name != null">
            and item.NAME like concat(concat('%',#{name}),'%')
        </if>
        <if test="ids != null and ids.size() > 0">
            and item.ID in
            <foreach collection="ids" item="id" index="index" open="(" close=")" separator=",">
                #{id,jdbcType=BIGINT}
            </foreach>
        </if>
        <if test="sortColumn != null and sortColumn != 'NAME' and sortDirection != null">
            order by item.${sortColumn} ${sortDirection}
        </if>
        <if test="sortColumn == 'NAME' and sortDirection != null">
            order by convert(item.NAME using gbk) ${sortDirection}
        </if>
        <if test="sortColumn == null or sortDirection == null">
            order by convert(item.NAME using gbk) asc
        </if>
        <if test="page == null or pageSize == null">
            limit 0, 10
        </if>
        <if test="page != null and pageSize != null">limit #{page}, #{pageSize}</if>
    
```
#### getIdByCode
```

        select ID from CJ_C2_CATALOG_ITEM
        where code = #{code} and IS_DELETED = 0
        order by isnull(VERSION) desc, convert(VERSION, unsigned integer) desc
        limit 1
    
```
#### getIdWillActive
```

        select ID
        from CJ_C2_CATALOG_ITEM
        where IS_DELETED = 0
        and CODE = #{code}
        and SYSTEMED = 0
        and IS_ACTIVE = 0
        and MEMBER_ID = #{memberId,jdbcType=BIGINT}
        and VERSION is NOT NULL
        order by convert(VERSION, unsigned integer) desc
        limit 1
    
```
#### queryHistoryList
```

        select
        <include refid="CatalogItem_With_Category_Column_List"/>
        from CJ_C2_CATALOG_ITEM item join CJ_C2_CATEGORY category on item.CATEGORY = category.ID and item.MEMBER_ID = category.MEMBER_ID
        left join CJ_C2_CATALOG_ITEM_CONTEXT context on item.ID = context.CATALOG_ITEM
        where item.IS_DELETED = 0
        and item.CODE = (select CODE from CJ_C2_CATALOG_ITEM where ID = #{id,jdbcType=BIGINT})
        <if test="sortColumn != null and sortDirection != null">order by item.${sortColumn} ${sortDirection}</if>
        <if test="sortColumn == null or sortDirection == null">order by isnull(item.VERSION) desc, convert(item.VERSION,
            unsigned integer) desc
        </if>
        <if test="page != null and pageSize != null">limit #{page}, #{pageSize}</if>
    
```
#### queryMaxVersion
```

        select max(convert(VERSION, unsigned integer))
        from CJ_C2_CATALOG_ITEM
        where CODE = (select CODE from CJ_C2_CATALOG_ITEM where ID = #{id,jdbcType=BIGINT})
    
```
#### getLatest
```

        select
        <include refid="Base_Column_List"/>
        ,
        <include refid="Blob_Column_List"/>
        from CJ_C2_CATALOG_ITEM
        where IS_DELETED = 0 and CODE = #{code} and (MEMBER_ID = #{memberId} or SYSTEMED = 1)
        order by isnull(VERSION) desc, convert(VERSION, unsigned integer) desc limit 1
    
```
### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/CatalogItemImportHistoryMapper.xml
> sql统计:复杂0个，简单2个
>
> 使用的关键字：values,CJ_C2_CATALOG_ITEM_IMPORT_HISTORY

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/CatalogItemNotifyDOMapper.xml
> sql统计:复杂0个，简单13个
>
> 使用的关键字：count,size,and,where,CJ_C2_CATALOG_ITEM_NOTIFY,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/CatalogItemPackageDOMapper.xml
> sql统计:复杂0个，简单18个
>
> 使用的关键字：count,size,CJ_C2_CATALOG_ITEM_PACKAGE,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/CatalogItemVariableApprovalEditDoMapper.xml
> sql统计:复杂0个，简单8个
>
> 使用的关键字：CJ_C2_CATALOG_ITEM_VARIABLE_APPROVAL_EDIT,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/CategoryDOMapper.xml
> sql统计:复杂3个，简单26个
>
> 使用的关键字：from,if,where,and,AND,COUNT,convert,count,in,values,join,size,max,JOIN,CJ_C2_CATEGORY
#### queryPageList
```

        select
        <include refid="Category_With_Catalog_Column_List"/>
        from
        CJ_C2_CATEGORY category
        <if test="parentMemberId != null">
            inner join
            (select CATEGORY, `CODE`
            from CJ_C2_CATALOG_ITEM
            where IS_DELETED = 0 )
            item
            on category.ID = item.CATEGORY
            inner join CJ_C2_CATALOG_ITEM_AUTH auth
            on item.`CODE` = auth.CATALOG_ITEM
        </if>
        where
        category.IS_DELETED = 0
        <if test="catalog != null">and CATALOG = #{catalog}</if>
        <if test="title != null">and category.TITLE like "%"#{title}"%"</if>
        <choose>
            <when test="parentMemberId != null">
                and auth.`VALUE` =  #{memberId}
                and auth.TYPE = 'subtenant'
                and auth.IS_DELETED = 0
                <!-- 子租户只有编辑权限没有禁、启用权限 默认查询启用 -->
                and category.IS_ACTIVE = 1
                and (category.MEMBER_ID = #{parentMemberId} or category.MEMBER_ID = 0)
            </when>
            <otherwise>
                <if test="isActive != null">and IS_ACTIVE = #{isActive}</if>
                and category.MEMBER_ID = #{memberId}
            </otherwise>
        </choose>

        <!-- 联合查询内置服务项类别信息 -->
        UNION
        select
        <include refid="Category_With_Catalog_Column_List"/>
        from
        CJ_C2_CATEGORY category
        where category.SYSTEMED = 1
        <!-- 内置服务项默认均为启用无需传入 -->
        and category.IS_ACTIVE = 1
        <choose>
            <when test="catalogs != null and catalogs.size() > 0">
                and TRUE
            </when>
            <otherwise>
                and false
            </otherwise>
        </choose>
        <if test="title != null">and category.TITLE like "%"#{title}"%"</if>

        group by ID
        <if test="sortColumn != null and sortColumn != 'TITLE' and sortDirection != null">order by ${sortColumn} ${sortDirection}</if>
        <if test="sortColumn == 'TITLE' and sortDirection != null">order by convert(TITLE using gbk) ${sortDirection}</if>
        <if test="sortColumn == null or sortDirection == null">order by `ORDER` asc</if>
        <if test="page != null and pageSize != null">limit #{page}, #{pageSize}</if>
    
```
#### queryUserPageCount
```

        select count(1)
        from CJ_C2_CATEGORY
        where IS_DELETED = 0
        and IS_ACTIVE = 1
        and (ID in (select distinct item.CATEGORY
                from CJ_C2_CATALOG_ITEM item
                join (select CODE, max(convert(VERSION, unsigned integer)) as VERSION from CJ_C2_CATALOG_ITEM where IS_ACTIVE = 1 group by CODE) vItem on item.CODE = vItem.CODE
                where item.IS_DELETED = 0
                and item.IS_ACTIVE = 1
                and item.VERSION = vItem.VERSION
                <if test="title != null">and item.NAME like "%"#{title}"%"</if>
                <if test="approvalMethod != null">and item.APPROVAL_METHOD = #{approvalMethod}</if>
                <if test="catItems != null and catItems.size() > 0">
                    and item.ID in
                    <foreach collection="catItems" item="id" index="index" open="(" close=")" separator=",">
                        #{id,jdbcType=BIGINT}
                    </foreach>
                </if>)
                and ((SYSTEMED = 1 and VISIBLE = 1)
                    <if test="memberId != null">or MEMBER_ID = #{memberId}</if>))
        <if test="catalogs != null and catalogs.size() > 0">
            and CATALOG IN
            <foreach collection="catalogs" item="id" index="index" open="(" close=")" separator=",">
                #{id,jdbcType=BIGINT}
            </foreach>
        </if>
        <if test="catalog != null">and CATALOG = #{catalog}</if>
        and ((SYSTEMED = 1 and VISIBLE = 1)
        <if test="memberId != null">or MEMBER_ID = #{memberId}</if>)
    
```
#### queryUserPageList
```

        select category.ID, category.DESCRIPTION, category.TITLE,
               count(if(catalog_item.itemVisible = 1,1,null)) as catalogItemTotal
        from CJ_C2_CATEGORY category
        JOIN
            (select item.VISIBLE as itemVisible,
                    item.CATEGORY as categoryId
            from CJ_C2_CATALOG_ITEM item
                left join (select CODE, max(convert(VERSION, unsigned integer)) as VERSION
                    from CJ_C2_CATALOG_ITEM
                    where IS_ACTIVE = 1
                    group by CODE) vItem on item.CODE = vItem.CODE
            where item.IS_DELETED = 0
            and item.IS_ACTIVE = 1
            and item.VERSION = vItem.VERSION
            <if test="title != null">and item.NAME like "%"#{title}"%"</if>
            <if test="approvalMethod != null">and item.APPROVAL_METHOD = #{approvalMethod}</if>
            <if test="catItems != null and catItems.size() > 0">
                and item.ID in
                <foreach collection="catItems" item="id" index="index" open="(" close=")" separator=",">
                    #{id,jdbcType=BIGINT}
                </foreach>
                </if> ) as catalog_item ON catalog_item.categoryId = category.ID
        where category.IS_DELETED = 0
            and category.CATALOG > 0
            and category.IS_ACTIVE = 1
            <if test="catalogs != null and catalogs.size() > 0">
                and category.CATALOG IN
                <foreach collection="catalogs" item="id" index="index" open="(" close=")" separator=",">
                    #{id,jdbcType=BIGINT}
                </foreach>
            </if>
            <if test="catalog != null">and category.catalog = #{catalog}</if>
            and ((category.SYSTEMED = 1 and category.VISIBLE = 1)
            <if test="memberId != null">or category.MEMBER_ID = #{memberId}</if>)
        group by category.ID, category.DESCRIPTION, category.TITLE
        order by category.`ORDER` asc
        <if test="page != null and pageSize != null">limit #{page}, #{pageSize}</if>
    
```
### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/CategoryManualNodeManageDOMapper.xml
> sql统计:复杂0个，简单11个
>
> 使用的关键字：count,CJ_C2_CATEGORY_MANUAL_NODE_MANAGE,values,and,IF

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/ItemOptionDOMapper.xml
> sql统计:复杂0个，简单21个
>
> 使用的关键字：CJ_C2_ITEM_OPTION,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/ItemVariableDOMapper.xml
> sql统计:复杂0个，简单56个
>
> 使用的关键字：values,CJ_C2_ITEM_VARIABLE,count,in,MAX,size,MIN,and

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/ItemVariableSetDOMapper.xml
> sql统计:复杂0个，简单24个
>
> 使用的关键字：COUNT,count,in,MAX,join,now,CJ_C2_ITEM_VARIABLE_SET,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/RequestApprovalDoMapper.xml
> sql统计:复杂0个，简单5个
>
> 使用的关键字：count,CJ_C2_REQUEST_APPROVAL,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/RequestDOMapper.xml
> sql统计:复杂3个，简单44个
>
> 使用的关键字：count,and,where,in,CURDATE,CJ_C2_REQUEST,size,from,join,date_format,sum,values,ifnull,curdate,IN,date_sub,DATE_SUB,date
#### queryProvisionSuccessRateByRecentDay
```

        select a.datetime         as time,
        ifnull(b.num, 0)          as num,
        ifnull(b.succeedTotal, 0) as succeedTotal
        from (
        SELECT curdate() as datetime
        union all
        SELECT date_sub(curdate(), interval 1 day) as click_date
        union all
        SELECT date_sub(curdate(), interval 2 day) as click_date
        union all
        SELECT date_sub(curdate(), interval 3 day) as click_date
        union all
        SELECT date_sub(curdate(), interval 4 day) as click_date
        union all
        SELECT date_sub(curdate(), interval 5 day) as click_date
        union all
        SELECT date_sub(curdate(), interval 6 day) as click_date
        ) a
        left join(
        select count(1)                      as num,
        date_format(CREATE_DATE, '%Y-%m-%d') as CREATE_DATE,
        sum(STATE = 'succeed')               as succeedTotal
        from CJ_C2_REQUEST
        where NUMBER in (select DISTINCT ritm.REQUEST
                     from CJ_C2_REQUESTED_ITEM ritm
                              join CJ_C2_CATALOG_ITEM item on item.ID = ritm.CAT_ITEM
                              join CJ_C2_CATALOG_ITEM_CONTEXT itemContex
                                   on item.ID = itemContex.CATALOG_ITEM
                     where (itemContex.METADATA -> '$.steps[*].operation' like '%"provision"%'
                        or itemContex.METADATA -> '$.steps[*].operation' like '%"create"%'
                        or itemContex.OPERATION_TYPE = 'provision')
                       and ritm.MEMBER_ID = #{memberId}
                       and ritm.IS_DELETED = 0
        )
        and MEMBER_ID = #{memberId}
        and IS_DELETED = 0
        and IS_DEBUG = 0
        and STAGE = 'implement'
        and STATE in ('failed', 'succeed', 'exception', 'succeedPartly', 'abort')
        and DATE_SUB(CURDATE(), INTERVAL 6 DAY) &lt;= date(CREATE_DATE)
        group by date_format(CREATE_DATE, '%Y-%m-%d')
        ) b on b.CREATE_DATE = a.datetime
        ORDER BY a.datetime ASC;
    
```
#### queryProvisionSuccessRateByRecentMonth
```

        select a.datetime         as time,
        ifnull(b.num, 0)          as num,
        ifnull(b.succeedTotal, 0) as succeedTotal
        from (
        SELECT date_format(curdate(), '%Y-%u') as datetime
        union all
        SELECT date_format(date_sub(curdate(), interval 1 week), '%Y-%u') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 2 week), '%Y-%u') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 3 week), '%Y-%u') as click_date
        ) a
        left join(
        select count(1)                      as num,
        date_format(CREATE_DATE, '%Y-%u') as datetime,
        sum(STATE = 'succeed')               as succeedTotal
        from CJ_C2_REQUEST
        where NUMBER in (select DISTINCT ritm.REQUEST
                     from CJ_C2_REQUESTED_ITEM ritm
                              join CJ_C2_CATALOG_ITEM item on item.ID = ritm.CAT_ITEM
                              join CJ_C2_CATALOG_ITEM_CONTEXT itemContex
                                   on item.ID = itemContex.CATALOG_ITEM
                     where (itemContex.METADATA -> '$.steps[*].operation' like '%"provision"%'
                        or itemContex.METADATA -> '$.steps[*].operation' like '%"create"%'
                        or itemContex.OPERATION_TYPE = 'provision')
                       and ritm.MEMBER_ID = #{memberId}
                       and ritm.IS_DELETED = 0
        )
        and MEMBER_ID = #{memberId}
        and IS_DELETED = 0
        and IS_DEBUG = 0
        and STAGE = 'implement'
        and STATE in ('failed', 'succeed', 'exception', 'succeedPartly', 'abort')
        and DATE_SUB(CURDATE(), INTERVAL 4 week) &lt;= date(CREATE_DATE)
        group by datetime
        ) b on b.datetime = a.datetime
        ORDER BY a.datetime ASC;
    
```
#### queryProvisionSuccessRateByRecentYear
```

        select a.datetime         as time,
        ifnull(b.num, 0)          as num,
        ifnull(b.succeedTotal, 0) as succeedTotal
        from (
        SELECT date_format(curdate(), '%Y-%m') as datetime
        union all
        SELECT date_format(date_sub(curdate(), interval 1 month), '%Y-%m') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 2 month), '%Y-%m') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 3 month), '%Y-%m') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 4 month), '%Y-%m') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 5 month), '%Y-%m') as click_date
        ) a
        left join(
        select count(1)                      as num,
        date_format(CREATE_DATE, '%Y-%m')    as datetime,
        sum(STATE = 'succeed')               as succeedTotal
        from CJ_C2_REQUEST
        where NUMBER in (select DISTINCT ritm.REQUEST
                     from CJ_C2_REQUESTED_ITEM ritm
                              join CJ_C2_CATALOG_ITEM item on item.ID = ritm.CAT_ITEM
                              join CJ_C2_CATALOG_ITEM_CONTEXT itemContex
                                   on item.ID = itemContex.CATALOG_ITEM
                     where (itemContex.METADATA -> '$.steps[*].operation' like '%"provision"%'
                        or itemContex.METADATA -> '$.steps[*].operation' like '%"create"%'
                        or itemContex.OPERATION_TYPE = 'provision')
                       and ritm.MEMBER_ID = #{memberId}
                       and ritm.IS_DELETED = 0
        )
        and MEMBER_ID = #{memberId}
        and IS_DELETED = 0
        and IS_DEBUG = 0
        and STAGE = 'implement'
        and STATE in ('failed', 'succeed', 'exception', 'succeedPartly', 'abort')
        and DATE_SUB(CURDATE(), INTERVAL 5 month ) &lt;= date(CREATE_DATE)
        group by datetime
        ) b on b.datetime = a.datetime
        ORDER BY a.datetime ASC;
    
```
### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/RequestOperateHistoryDOMapper.xml
> sql统计:复杂0个，简单4个
>
> 使用的关键字：CJ_C2_REQUEST_OPERATE_HISTORY,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/RequestedItemApprovalDoMapper.xml
> sql统计:复杂0个，简单5个
>
> 使用的关键字：CJ_C2_REQUESTED_ITEM_APPROVAL,values,count

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/RequestedItemDOMapper.xml
> sql统计:复杂7个，简单103个
>
> 使用的关键字：join,in,MAX,DATE_SUB,JSON_VALID,from,min,AVG,or,round,sum,AND,ifnull,where,COUNT,MIN,avg,concat,date_sub,date,week,CJ_C2_REQUESTED_ITEM,and,max,TIMESTAMPDIFF,curdate,ROUND,size,count,IF,JSON_CONTAINS,DATE,date_format,CURDATE,values,STRCMP,OR,substr,format
#### statisticsCatalogItem
```

        select a.ID catItem,
        a.NAME catItemName,
        IF(a.SOURCE = 'cloudTemplate', 'IaC环境', '传统服务') source,
        a.total num,
        format(a.total / c.totalAll * 100, 2) rate,
        b.totalSucceed firstSucceedNum,
        format(b.totalSucceed / a.total * 100, 2) firstSucceedRate,
        b.longestTime longTime,
        b.shortestTime shortTime,
        b.averageTime averageTime
        from (select count(1) total,
        item.ID,
        item.NAME,
        item.SOURCE
        from CJ_C2_CATALOG_ITEM item
        join CJ_C2_REQUESTED_ITEM ritm
        on item.ID = ritm.CAT_ITEM
        where ritm.IS_ACTIVE = 1
        and ritm.IS_DELETED = 0
        and ritm.IS_DEBUG = 0
        <if test="startDate != null and endDate != null">and DATE(ritm.CREATE_DATE) between #{startDate} and #{endDate}</if>
        <if test="memberId != null">and ritm.MEMBER_ID = #{memberId,jdbcType=BIGINT}</if>
        <if test="createBy != null">and ritm.CREATE_BY = #{createBy,jdbcType=VARCHAR}</if>
        group by item.ID, item.NAME) a
        left join
        (select t.CAT_ITEM,
        max(TIMESTAMPDIFF(MINUTE, t.startTime, t.endTime)) longestTime,
        min(TIMESTAMPDIFF(MINUTE, t.startTime, t.endTime)) shortestTime,
        round(avg(TIMESTAMPDIFF(MINUTE, t.startTime, t.endTime))) averageTime,
        count(1) totalSucceed
        from (select ritm.CAT_ITEM, ritm.NUMBER, startRecord.CREATE_DATE startTime, endRecord.CREATE_DATE endTime
        from (select REQUESTED_ITEM, min(CREATE_DATE) CREATE_DATE, count(1) num
        from CJ_C2_REQUESTED_ITEM_OPERATE_HISTORY
        where state = 'implement'
        group by REQUESTED_ITEM) startRecord
        join
        (select REQUESTED_ITEM, max(CREATE_DATE) CREATE_DATE, count(1) num
        from CJ_C2_REQUESTED_ITEM_OPERATE_HISTORY
        where state = 'succeed'
        group by REQUESTED_ITEM) endRecord
        on startRecord.REQUESTED_ITEM = endRecord.REQUESTED_ITEM
        join
        CJ_C2_REQUESTED_ITEM ritm
        on startRecord.REQUESTED_ITEM = ritm.ID
        join
        CJ_C2_CATALOG_ITEM item
        on item.ID = ritm.CAT_ITEM
        where ritm.state = 'succeed'
        and ritm.IS_ACTIVE = 1
        and ritm.IS_DELETED = 0
        and ritm.IS_DEBUG = 0
        <if test="startDate != null and endDate != null">and DATE(ritm.CREATE_DATE) between #{startDate} and #{endDate}</if>
        <if test="memberId != null">and ritm.MEMBER_ID = #{memberId,jdbcType=BIGINT}</if>
        <if test="createBy != null">and ritm.CREATE_BY = #{createBy,jdbcType=VARCHAR}</if>
        ) t
        group by CAT_ITEM
        ) b
        on a.ID = b.CAT_ITEM
        left join (select count(1) totalAll
        from CJ_C2_CATALOG_ITEM item
        join CJ_C2_REQUESTED_ITEM ritm
        on item.ID = ritm.CAT_ITEM
        where ritm.IS_ACTIVE = 1
        and ritm.IS_DELETED = 0
        and ritm.IS_DEBUG = 0
        <if test="startDate != null and endDate != null">and DATE(ritm.CREATE_DATE) between #{startDate} and #{endDate}</if>
        <if test="memberId != null">and ritm.MEMBER_ID = #{memberId,jdbcType=BIGINT}</if>
        <if test="createBy != null">and ritm.CREATE_BY = #{createBy,jdbcType=VARCHAR}</if>
        ) c
        on 1 = 1
        where 1 = 1
        <if test="catItemName != null">and a.NAME like concat("%",#{catItemName,jdbcType=VARCHAR},"%")</if>
        <if test="source != null and source.size() > 0">
            and a.SOURCE in
            <foreach collection="source" item="item" index="index" open="(" close=")" separator=",">
                #{item}
            </foreach>
        </if>
        order by num desc, firstSucceedNum desc
        <if test="page != null and pageSize != null">limit #{page}, #{pageSize}</if>
    
```
#### countStatisticsCatalogItem
```

        select
        count(1)
        from (select count(1) total,
        item.ID,
        item.NAME,
        item.SOURCE
        from CJ_C2_CATALOG_ITEM item
        join CJ_C2_REQUESTED_ITEM ritm
        on item.ID = ritm.CAT_ITEM
        where ritm.IS_ACTIVE = 1
        and ritm.IS_DELETED = 0
        and ritm.IS_DEBUG = 0
        <if test="startDate != null and endDate != null">and DATE(ritm.CREATE_DATE) between #{startDate} and #{endDate}</if>
        <if test="memberId != null">and ritm.MEMBER_ID = #{memberId,jdbcType=BIGINT}</if>
        <if test="createBy != null">and ritm.CREATE_BY = #{createBy,jdbcType=VARCHAR}</if>
        group by item.ID, item.NAME) a
        left join
        (select t.CAT_ITEM,
        max(TIMESTAMPDIFF(MINUTE, t.startTime, t.endTime)) longestTime,
        min(TIMESTAMPDIFF(MINUTE, t.startTime, t.endTime)) shortestTime,
        round(avg(TIMESTAMPDIFF(MINUTE, t.startTime, t.endTime))) averageTime,
        count(1) totalSucceed
        from (select ritm.CAT_ITEM, ritm.NUMBER, startRecord.CREATE_DATE startTime, endRecord.CREATE_DATE endTime
        from (select REQUESTED_ITEM, min(CREATE_DATE) CREATE_DATE, count(1) num
        from CJ_C2_REQUESTED_ITEM_OPERATE_HISTORY
        where state = 'implement'
        group by REQUESTED_ITEM) startRecord
        join
        (select REQUESTED_ITEM, max(CREATE_DATE) CREATE_DATE, count(1) num
        from CJ_C2_REQUESTED_ITEM_OPERATE_HISTORY
        where state = 'succeed'
        group by REQUESTED_ITEM) endRecord
        on startRecord.REQUESTED_ITEM = endRecord.REQUESTED_ITEM
        join
        CJ_C2_REQUESTED_ITEM ritm
        on startRecord.REQUESTED_ITEM = ritm.ID
        join
        CJ_C2_CATALOG_ITEM item
        on item.ID = ritm.CAT_ITEM
        where ritm.state = 'succeed'
        and ritm.IS_ACTIVE = 1
        and ritm.IS_DELETED = 0
        and ritm.IS_DEBUG = 0
        <if test="startDate != null and endDate != null">and DATE(ritm.CREATE_DATE) between #{startDate} and #{endDate}</if>
        <if test="memberId != null">and ritm.MEMBER_ID = #{memberId,jdbcType=BIGINT}</if>
        <if test="createBy != null">and ritm.CREATE_BY = #{createBy,jdbcType=VARCHAR}</if>
        ) t
        group by CAT_ITEM
        ) b
        on a.ID = b.CAT_ITEM
        left join (select count(1) totalAll
        from CJ_C2_CATALOG_ITEM item
        join CJ_C2_REQUESTED_ITEM ritm
        on item.ID = ritm.CAT_ITEM
        where ritm.IS_ACTIVE = 1
        and ritm.IS_DELETED = 0
        and ritm.IS_DEBUG = 0
        <if test="startDate != null and endDate != null">and DATE(ritm.CREATE_DATE) between #{startDate} and #{endDate}</if>
        <if test="memberId != null">and ritm.MEMBER_ID = #{memberId,jdbcType=BIGINT}</if>
        <if test="createBy != null">and ritm.CREATE_BY = #{createBy,jdbcType=VARCHAR}</if>
        ) c
        on 1 = 1
        where 1 = 1
        <if test="catItemName != null">and a.NAME like concat("%",#{catItemName,jdbcType=VARCHAR},"%")</if>
        <if test="source != null and source.size() > 0">
            and a.SOURCE in
            <foreach collection="source" item="item" index="index" open="(" close=")" separator=",">
                #{item}
            </foreach>
        </if>
    
```
#### queryProvisionSuccessRateByRecentDay
```

        select a.datetime as time,
        ifnull(b.num, 0) as num,
        ifnull(b.succeedTotal, 0) as succeedTotal
        from (
        SELECT curdate() as datetime
        union all
        SELECT date_sub(curdate(), interval 1 day) as click_date
        union all
        SELECT date_sub(curdate(), interval 2 day) as click_date
        union all
        SELECT date_sub(curdate(), interval 3 day) as click_date
        union all
        SELECT date_sub(curdate(), interval 4 day) as click_date
        union all
        SELECT date_sub(curdate(), interval 5 day) as click_date
        union all
        SELECT date_sub(curdate(), interval 6 day) as click_date
        ) a
        left join(
        select count(1) as num,
        date_format(ritm.CREATE_DATE, '%Y-%m-%d') as CREATE_DATE,
        sum(ritm.STATE = 'succeed') as succeedTotal
        from CJ_C2_REQUESTED_ITEM ritm
        join CJ_C2_CATALOG_ITEM item
        on item.ID = ritm.CAT_ITEM
        join CJ_C2_CATALOG_ITEM_CONTEXT itemContex
        on item.ID = itemContex.CATALOG_ITEM
        where (itemContex.METADATA -> '$.steps[*].operation' like '%"provision"%'
        or itemContex.METADATA -> '$.steps[*].operation' like '%"create"%'
        or itemContex.OPERATION_TYPE = 'provision')
        and ritm.MEMBER_ID = #{memberId}
        and ritm.IS_DELETED = 0
        and ritm.IS_DEBUG = 0
        and ritm.STAGE = 'implement'
        and ritm.STATE in ('failed', 'succeed', 'exception', 'abort')
        and DATE_SUB(CURDATE(), INTERVAL 6 DAY) &lt;= date(ritm.CREATE_DATE)
        group by date_format(ritm.CREATE_DATE, '%Y-%m-%d')
        ) b on b.CREATE_DATE = a.datetime
        ORDER BY a.datetime ASC;
    
```
#### queryProvisionSuccessRateByRecentMonth
```

        select a.datetime as time,
        ifnull(b.num, 0) as num,
        ifnull(b.succeedTotal, 0) as succeedTotal
        from (
        SELECT date_format(curdate(), '%Y-%u') as datetime
        union all
        SELECT date_format(date_sub(curdate(), interval 1 week), '%Y-%u') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 2 week), '%Y-%u') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 3 week), '%Y-%u') as click_date
        ) a
        left join(
        select count(1) as num,
        date_format(ritm.CREATE_DATE, '%Y-%u') as datetime,
        sum(ritm.STATE = 'succeed') as succeedTotal
        from CJ_C2_REQUESTED_ITEM ritm
        join CJ_C2_CATALOG_ITEM item
        on item.ID = ritm.CAT_ITEM
        join CJ_C2_CATALOG_ITEM_CONTEXT itemContex
        on item.ID = itemContex.CATALOG_ITEM
        where (itemContex.METADATA -> '$.steps[*].operation' like '%"provision"%'
        or itemContex.METADATA -> '$.steps[*].operation' like '%"create"%'
        or itemContex.OPERATION_TYPE = 'provision')
        and ritm.MEMBER_ID = #{memberId}
        and ritm.IS_DELETED = 0
        and ritm.IS_DEBUG = 0
        and ritm.STAGE = 'implement'
        and ritm.STATE in ('failed', 'succeed', 'exception', 'abort')
        and DATE_SUB(CURDATE(), INTERVAL 4 week) &lt;= date(ritm.CREATE_DATE)
        group by datetime
        ) b on b.datetime = a.datetime
        ORDER BY a.datetime ASC;
    
```
#### queryProvisionSuccessRateByRecentYear
```

        select a.datetime as time,
        ifnull(b.num, 0) as num,
        ifnull(b.succeedTotal, 0) as succeedTotal
        from (
        SELECT date_format(curdate(), '%Y-%m') as datetime
        union all
        SELECT date_format(date_sub(curdate(), interval 1 month), '%Y-%m') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 2 month), '%Y-%m') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 3 month), '%Y-%m') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 4 month), '%Y-%m') as click_date
        union all
        SELECT date_format(date_sub(curdate(), interval 5 month), '%Y-%m') as click_date
        ) a
        left join(
        select count(1) as num,
        date_format(ritm.CREATE_DATE, '%Y-%m') as datetime,
        sum(ritm.STATE = 'succeed') as succeedTotal
        from CJ_C2_REQUESTED_ITEM ritm
        join CJ_C2_CATALOG_ITEM item
        on item.ID = ritm.CAT_ITEM
        join CJ_C2_CATALOG_ITEM_CONTEXT itemContex
        on item.ID = itemContex.CATALOG_ITEM
        where (itemContex.METADATA -> '$.steps[*].operation' like '%"provision"%'
        or itemContex.METADATA -> '$.steps[*].operation' like '%"create"%'
        or itemContex.OPERATION_TYPE = 'provision')
        and ritm.MEMBER_ID = #{memberId}
        and ritm.IS_DELETED = 0
        and ritm.IS_DEBUG = 0
        and ritm.STAGE = 'implement'
        and ritm.STATE in ('failed', 'succeed', 'exception', 'abort')
        and DATE_SUB(CURDATE(), INTERVAL 5 month) &lt;= date(ritm.CREATE_DATE)
        group by datetime
        ) b on b.datetime = a.datetime
        ORDER BY a.datetime ASC;
    
```
#### summary
```

        select max(TIMESTAMPDIFF(MINUTE, t.startTime, t.endTime)) longTime,
        min(TIMESTAMPDIFF(MINUTE, t.startTime, t.endTime)) shortTime,
        round(avg(TIMESTAMPDIFF(MINUTE, t.startTime, t.endTime))) averageTime,
        count(1) firstSucceedNum
        from (select startRecord.CREATE_DATE startTime,
        endRecord.CREATE_DATE   endTime
        from (select REQUESTED_ITEM, min(CREATE_DATE) CREATE_DATE, count(1) num
        from CJ_C2_REQUESTED_ITEM_OPERATE_HISTORY
        where state = 'implement'
        group by REQUESTED_ITEM) startRecord
        join (select REQUESTED_ITEM, max(CREATE_DATE) CREATE_DATE, count(1) num
        from CJ_C2_REQUESTED_ITEM_OPERATE_HISTORY
        where state = 'succeed'
        group by REQUESTED_ITEM) endRecord
        on startRecord.REQUESTED_ITEM = endRecord.REQUESTED_ITEM
        join CJ_C2_REQUESTED_ITEM ritm on startRecord.REQUESTED_ITEM = ritm.ID
        where ritm.state = 'succeed'
        and ritm.IS_ACTIVE = 1
        and ritm.IS_DELETED = 0
        and ritm.IS_DEBUG = 0
        <if test="startDate != null and endDate != null">and DATE(ritm.CREATE_DATE) between #{startDate} and #{endDate}</if>
        <if test="memberId != null">and ritm.MEMBER_ID = #{memberId,jdbcType=BIGINT}</if>
        <if test="createBy != null">and ritm.CREATE_BY = #{createBy,jdbcType=VARCHAR}</if>) t
    
```
#### getTrend
```

        select sum(a.total)                                  as total,
            sum(IF(a.source = 'cloudTemplate', a.total, 0))  as iac,
            sum(IF(a.source != 'cloudTemplate', a.total, 0)) as ordinary,
            a.createDate                                     as `week`
            from (select COUNT(1) total, CONTEXT_SOURCE source, week(CREATE_DATE) createDate
        from CJ_C2_REQUESTED_ITEM
        where IS_ACTIVE = 1
            and IS_DELETED = 0
            and IS_DEBUG = 0
            <if test="startDate != null and endDate != null">and DATE(CREATE_DATE) between #{startDate} and #{endDate}</if>
            <if test="memberId != null">and MEMBER_ID = #{memberId,jdbcType=BIGINT}</if>
            <if test="createBy != null">and CREATE_BY = #{createBy,jdbcType=VARCHAR}</if>
        group by CONTEXT_SOURCE, week(CREATE_DATE)) a
        group by a.createDate
    
```
### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/RequestedItemManualDoMapper.xml
> sql统计:复杂0个，简单4个
>
> 使用的关键字：values,count,CJ_C2_REQUESTED_ITEM_MANUAL

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/RequestedItemOperateHistoryDOMapper.xml
> sql统计:复杂0个，简单3个
>
> 使用的关键字：CJ_C2_REQUESTED_ITEM_OPERATE_HISTORY,values

### /Users/jiangliuhong/develop/cvs/idcos/cloud-catalog/catalog-dal/src/main/resources/mapping/TaskDOMapper.xml
> sql统计:复杂0个，简单9个
>
> 使用的关键字：CJ_C2_TASK,values,NOW

