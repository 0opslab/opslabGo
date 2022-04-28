CREATE  TABLE app_table_httpquery (
    it_name varchar(30) not null COMMENT '接口名称' PRIMARY KEY,
    cname varchar(200) not null COMMENT '接口描述',
    ctime datetime default now() COMMENT '接口创建时间',
    query_type varchar(20) COMMENT '接口查询类型',
    row_type int(2) COMMENT '接口返回结果集1为单一对象大于1位list对象',
    cache_time int(20) default 0 comment '接口缓存时间单位秒',
    f_status int(2) not null COMMENT '接口状态',
    query_string varchar(3000) comment '接口查询语句',
    remark varchar(2000) comment '备注字段'
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


insert into app_table_httpquery(it_name,cname,query_type,row_type,cache_time,f_status,query_string) values('db_time','数据库时间','query',1,0,1,'select now() times from dual');
insert into app_table_httpquery(it_name,cname,query_type,row_type,cache_time,f_status,query_string) values('test_1','测试接口','name_query',1,0,1,"SELECT 'country', 'city', 11 as telcode FROM dual where 1=:fn");
insert into app_table_httpquery(it_name,cname,query_type,row_type,cache_time,f_status,query_string) values('test_3','测试接口','name_query',2,100,1,"select * from t_stock_info limit 2");

truncate table app_table_httpquery;
select it_name as Name,query_type as QueryType,row_type as RowType,cache_time as CacheTime,query_string as QueryString from app_table_httpquery where f_status=1;