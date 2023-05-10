use little_bee;

set global time_zone='+8:00';

#name of super user must be admin
INSERT INTO users(user_name,user_nick,password,created_by) VALUES('admin','超级管理员','21232f297a57a5a743894a0e4a801fc3',NULL);
INSERT INTO roles(role_name,role_alias,role_note) VALUES('super_admin_role','超级管理员组','超级管理员组');

#### menus
# level 1
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(1,NULL,'home','首页',1,'/home','/home','Home.vue','HomeFilled','首页',NULL,NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(2,NULL,'monitor','实时监控',1,'/monitor','/monitor',NULL,'Monitor','实时监控',NULL,NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(3,NULL,'query','数据查询',1,'/query','/query',NULL,'DataAnalysis','数据查询',NULL,NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(4,NULL,'system','系统',1,'/system','/system',NULL,'Service','系统',NULL,NULL);

# level 2
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(201,2,'real-data','实时数据',2,'/real-data/:nodeType/:nodeId','/real-data/1/0','real/RealData.vue','DataLine','实时数据','layout/side/SignalTree.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(202,2,'real-alarm','实时告警',2,'/real-alarm','/real-alarm','real/RealAlarm.vue','AlarmClock','实时告警',NULL,NULL);

INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(301,3,'alarm-query','告警查询',2,'/alarm-query','/alarm-query','query/AlarmQuery.vue','Warning','告警查询',NULL,NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(302,3,'history-monitor-query','历史数据',2,'/history-monitor-query','/history-monitor-query','query/HistoryMonitorQuery.vue','Search','历史数据',NULL,NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(303,3,'log-query','日志查询',2,'/log-query','/log-query','query/LogQuery.vue','Comment','日志查询',NULL,NULL); 
 INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(304,3,'command-query','命令查询',2,'/command-query','/command-query','query/CommandQuery.vue','SetUp','命令查询',NULL,NULL); 

INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(401,4,'monitor-object','监控对象',2,'/logic/:logicObjectId/:operatedType/:objectType','/logic/0/2/1',
 'setting/logic/LogicObject.vue','Tickets','监控对象','layout/side/SignalTree.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(402,4,'agent-config','Agent配置',2,'/center','/center','setting/agent/Agent.vue','PriceTag','Agent配置','layout/side/FsuTree.vue', NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(403,4,'configuration-setting','组态配置',2,'/configuration-setting/:nodeType/:nodeId','/configuration-setting/1/0',
 'system/ConfigurationSetting.vue','Files','组态配置','layout/side/SignalTree.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404,4,'system-setting','设置',2,'/system-setting','/system-setting','system/SystemSetting.vue','Setting','系统设置','layout/side/Setting.vue',NULL);

# level 3
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404001,404,'area-type','区域类型',3,'/area-type','/area-type','setting/AreaType.vue','Position','区域类型','layout/side/Setting.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404002,404,'device-type','设备类型',3,'/device-type','/device-type','setting/DeviceType.vue','ForkSpoon','设备类型','layout/side/Setting.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404003,404,'alarm-level','告警级别',3,'/alarm-level','/alarm-level','setting/AlarmLevel.vue','TrendCharts','告警级别','layout/side/Setting.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404004,404,'metadata-config','元数据',3,'/metadata-config','/metadata-config','setting/MetadataConfig.vue','ScaleToOriginal','元数据','layout/side/Setting.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404005,404,'signal_unification','标准信号',3,'/signal_unification','/signal_unification','setting/SignalUnification.vue','Grid','标准信号','layout/side/Setting.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404006,404,'fsu-type','FSU类型',3,'/fsu-type','/fsu-type','setting/FsuType.vue','Link','FSU类型','layout/side/Setting.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404007,404,'driver-type','驱动类型',3,'/driver-type','/driver-type','setting/DriverType.vue','Switch','驱动类型','layout/side/Setting.vue',NULL);

INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404036,404,'department-config','部门配置',3,'/department-config','/department-config','setting/DepartmentConfig.vue','OfficeBuilding','部门配置','layout/side/Setting.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404037,404,'user-config','用户配置',3,'/user-config','/user-config','user/UserConfig.vue','UserFilled','用户配置','layout/side/Setting.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404038,404,'role-config','角色配置',3,'/role-config','/role-config','setting/RoleConfig.vue','User','角色配置','layout/side/Setting.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404039,404,'action-group','功能分组',3,'/action-group','/action-group','setting/ActionGroup.vue','SwitchFilled','功能分组','layout/side/Setting.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404040,404,'action-config','功能配置',3,'/action-config','/action-config','setting/ActionConfig.vue','Histogram','功能配置','layout/side/Setting.vue',NULL);
INSERT INTO menus(menu_id,parent_menu_id,menu_name,menu_alias,menu_level,path,url,component,icon,title, sidebar, created_by) 
 VALUES(404041,404,'menu-config','菜单配置',3,'/menu-config','/menu-config','setting/MenuConfig.vue','Menu','菜单配置','layout/side/Setting.vue',NULL);


#### action's groups
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(1,NULL,'monitor','实时监控',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(2,NULL,'query','数据查询',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(3,NULL,'system','系统',NULL);

INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(101,1,'real-data','实时数据',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(102,1,'real-alarm','实时告警',NULL);

INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(201,2,'alarm-query','告警查询',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(202,2,'history-monitor-query','监控数据',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(203,2,'log-query','日志查询',NULL); 
 INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(204,2,'command-query','命令查询',NULL); 

INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(301,3,'monitor-object','监控对象',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(302,3,'agent-config','Agent配置',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(303,3,'configuration-setting','组态配置',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304,3,'system-setting','设置',NULL);

INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304001,304,'area-type','区域类型',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304002,304,'device-type','设备类型',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304003,304,'alarm-level','告警级别',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304004,304,'metadata-config','元数据',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304005,304,'signal_unification','标准信号',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304006,304,'department-config','部门配置',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304007,304,'user-config','用户配置',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304008,304,'role-config','角色配置',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304009,304,'authorization-config','权限配置',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304010,304,'action-group','功能分组',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304011,304,'action-config','功能配置',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304012,304,'menu-config','菜单配置',NULL);
INSERT INTO action_groups(action_group_id,parent_action_group_id,action_group_name,action_group_alias,created_by) 
 VALUES(304013,304,'others','其他功能',NULL);


#### actions
# action_type
# CREATE   1
# DELETE   2
# UPDATE   3
# RETRIEVE 4

# affected_server
#8           4           2           1
#-------------------------------------  
#0           1           1           1
#NONE   Driverloader Fsuserver    RealData
 
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('assembled-menus','组织用户菜单',100000,'/user/assembled-menus','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('user-all-actions','获取用户所有授权的功能',100001,'/user/all-actions','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('request-realdata','请求实时数据',101,'/realdata','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('request-active-alarm','请求活动告警',102,'/active-alarm','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('control-command','控制命令',101,'/control','POST',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-action','增加功能',304011,'/config/action','POST',1,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-action','删除功能',304011,'/config/action/\\d{1,}$','DELETE',2,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-action','修改功能',304011,'/config/action','PUT',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-action','查找功能',304011,'/config/action/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-action','获取功能',304011,'/config/action/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-action','获取全部功能',304011,'/config/action/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-agent','增加Agent',302,'/config/agent','POST',1,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-agent','删除Agent',302,'/config/agent/\\d{1,}$','DELETE',2,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-agent','修改Agent',302,'/config/agent','PUT',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-agent','查找Agent',302,'/config/agent/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-agent','获取Agent',302,'/config/agent/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-agent','获取全部Agent',302,'/config/agent/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-alarm-level','增加告警级别',304003,'/config/alarm-level','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-alarm-level','删除告警级别',304003,'/config/alarm-level/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-alarm-level','修改告警级别',304003,'/config/alarm-level','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-alarm-level','查找告警级别',304003,'/config/alarm-level/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-alarm-level','获取告警级别',304003,'/config/alarm-level/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-alarm-level','获取全部告警级别',304003,'/config/alarm-level/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-alarm','查找告警',201,'/config/alarm/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-alarm','获取告警',201,'/config/area-type/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('ack-alarm','确认告警',201,'/config/alarm/ack','POST',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-area-type','增加区域类型',304001,'/config/area-type','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-area-type','删除区域类型',304001,'/config/area-type/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-area-type','修改区域类型',304001,'/config/area-type','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-area-type','查找区域类型',304001,'/config/area-type/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-area-type','获取区域类型',304001,'/config/area-type/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-area-type','获取全部区域类型',304001,'/config/area-type/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-configure','增加元数据实例',304004,'/config/configure','POST',1,7,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-configure','删除元数据实例',304004,'/config/configure/\\d{1,}$','DELETE',2,7,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-configure','修改元数据实例',304004,'/config/configure','PUT',3,7,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-configure','查找元数据实例',304004,'/config/configure/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-configure','获取元数据实例',304004,'/config/configure/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-configure','获取全部元数据实例',304004,'/config/configure/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-control-command-record','增加控制命令记录',204,'/config/control-command-record','POST',1,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-control-command-record','删除控制命令记录',204,'/config/control-command-record/\\d{1,}$','DELETE',2,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-control-command-record','修改控制命令记录',204,'/config/control-command-record','PUT',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-control-command-record','查找控制命令记录',204,'/config/control-command-record/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-control-command-record','获取控制命令记录',204,'/config/control-command-record/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-control-command-record','获取全部控制命令记录',204,'/config/control-command-record/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-department','增加部门',304006,'/config/department','POST',1,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-department','删除部门',304006,'/config/department/\\d{1,}$','DELETE',2,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-department','修改部门',304006,'/config/department','PUT',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-department','查找部门',304006,'/config/department/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-department','获取部门',304006,'/config/department/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-department','获取全部部门',304006,'/config/department/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-device-type','增加设备类型',304002,'/config/device-type','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-device-type','删除设备类型',304002,'/config/device-type/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-device-type','修改设备类型',304002,'/config/device-type','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-device-type','查找设备类型',304002,'/config/device-type/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-device-type','获取设备类型',304002,'/config/device-type/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-device-type','获取全部设备类型',304002,'/config/device-type/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-driver-type','增加驱动类型',302,'/config/driver-type','POST',1,4,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-driver-type','删除驱动类型',302,'/config/driver-type/\\d{1,}$','DELETE',2,4,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-driver-type','修改驱动类型',302,'/config/driver-type','PUT',3,4,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-driver-type','查找驱动类型',302,'/config/driver-type/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-driver-type','获取驱动类型',302,'/config/driver-type/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-driver-type','获取全部驱动类型',302,'/config/driver-type/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-driver','增加驱动',302,'/config/driver','POST',1,4,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-driver','删除驱动',302,'/config/driver/\\d{1,}$','DELETE',2,4,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-driver','修改驱动',302,'/config/driver','PUT',3,4,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-driver','查找驱动',302,'/config/driver/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-driver','获取驱动',302,'/config/driver/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-driver','获取全部驱动',302,'/config/driver/all','GET',4,0,NULL); 

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-fsu-type','增加FSU类型',302,'/config/fsu-type','POST',1,2,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-fsu-type','删除FSU类型',302,'/config/fsu-type/\\d{1,}$','DELETE',2,2,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-fsu-type','修改FSU类型',302,'/config/fsu-type','PUT',3,2,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-fsu-type','查找FSU类型',302,'/config/fsu-type/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-fsu-type','获取FSU类型',302,'/config/fsu-type/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-fsu-type','获取全部FSU类型',302,'/config/fsu-type/all','GET',4,0,NULL); 

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-fsu','增加FSU',302,'/config/fsu','POST',1,2,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-fsu','删除FSU',302,'/config/fsu/\\d{1,}$','DELETE',2,2,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-fsu','修改FSU',302,'/config/fsu','PUT',3,2,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-fsu','查找FSU',302,'/config/fsu/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-fsu','获取FSU',302,'/config/fsu/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-fsu','获取全部FSU',302,'/config/fsu/all','GET',4,0,NULL); 

 INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-logic-area','增加区域',301,'/config/logic-area','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-logic-area','删除区域',301,'/config/logic-area/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-logic-area','修改区域',301,'/config/logic-area','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-logic-area','查找区域',301,'/config/logic-area/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-logic-area','获取区域',301,'/config/logic-area/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-logic-area','获取全部区域',301,'/config/logic-area/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-logic-device','增加设备',301,'/config/logic-device','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-logic-device','删除设备',301,'/config/logic-device/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-logic-device','修改设备',301,'/config/logic-device','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-logic-device','查找设备',301,'/config/logic-device/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-logic-device','获取设备',301,'/config/logic-device/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-logic-device','获取全部设备',301,'/config/logic-device/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-logic-object','增加逻辑对象',301,'/config/logic-object','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-logic-object','删除逻辑对象',301,'/config/logic-object/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-logic-object','修改逻辑对象',301,'/config/logic-object','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-logic-object','查找逻辑对象',301,'/config/logic-object/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-logic-object','获取逻辑对象',301,'/config/logic-object/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-logic-object','获取全部逻辑对象',301,'/config/logic-object/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-menu','增加菜单',304012,'/config/menu','POST',1,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-menu','删除菜单',304012,'/config/menu/\\d{1,}$','DELETE',2,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-menu','修改菜单',304012,'/config/menu','PUT',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-menu','查找菜单',304012,'/config/menu/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-menu','获取菜单',304012,'/config/menu/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-menu','获取全部菜单',304012,'/config/menu/all','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-menu-action','增加菜单关联的功能',304012,'/config/menu-action','POST',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('batch-create-menu-action','批量增加菜单关联的功能',304012,'/config/menu-action/batch','POST',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-menu-action','删除菜单关联的功能',304012,'/config/menu-action/\\d{1,}$','DELETE',2,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-menu-action','修改菜单关联的功能',304012,'/config/menu-action','PUT',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-menu-action','查找菜单关联的功能',304012,'/config/menu-action/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-menu-action','获取菜单关联的功能',304012,'/config/menu-action/\\d{1,}$','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-metadata','增加元数据',304004,'/config/metadata','POST',1,7,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-metadata','删除元数据',304004,'/config/metadata/\\d{1,}$','DELETE',2,7,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-metadata','修改元数据',304004,'/config/metadata','PUT',3,7,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-metadata','查找元数据',304004,'/config/metadata/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-metadata','获取元数据',304004,'/config/metadata/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-metadata','获取全部元数据',304004,'/config/metadata/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-role','增加角色',304008,'/config/role','POST',1,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-role','删除角色',304008,'/config/role/\\d{1,}$','DELETE',2,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-role','修改角色',304008,'/config/role','PUT',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-role','查找角色',304008,'/config/role/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-role','获取角色',304008,'/config/role/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-role','获取全部角色',304008,'/config/role/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-signal-acquisition','增加信号采集',302,'/config/signal-acquisition','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-signal-acquisition','删除信号采集',302,'/config/signal-acquisition/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-signal-acquisition','修改信号采集',302,'/config/signal-acquisition','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-signal-acquisition','查找信号采集',302,'/config/signal-acquisition/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-signal-acquisition','获取信号采集',302,'/config/signal-acquisition/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-signal-acquisition','获取全部信号采集',302,'/config/signal-acquisition/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-signal-ai','增加AI信号',301,'/config/signal-ai','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-signal-ai','删除AI信号',301,'/config/signal-ai/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-signal-ai','修改AI信号',301,'/config/signal-ai','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-signal-ai','查找AI信号',301,'/config/signal-ai/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-signal-ai','获取AI信号',301,'/config/signal-ai/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-signal-ai','获取全部AI信号',301,'/config/signal-ai/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-signal-di','增加DI信号',301,'/config/signal-di','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-signal-di','删除DI信号',301,'/config/signal-di/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-signal-di','修改DI信号',301,'/config/signal-di','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-signal-di','查找DI信号',301,'/config/signal-di/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-signal-di','获取DI信号',301,'/config/signal-di/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-signal-di','获取全部DI信号',301,'/config/signal-di/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-signal-si','增加SI信号',301,'/config/signal-si','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-signal-si','删除SI信号',301,'/config/signal-si/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-signal-si','修改SI信号',301,'/config/signal-si','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-signal-si','查找SI信号',301,'/config/signal-si/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-signal-si','获取SI信号',301,'/config/signal-si/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-signal-si','获取全部SI信号',301,'/config/signal-si/all','GET',4,0,NULL);

 INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-signal-output','增加控制信号',301,'/config/signal-output','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-signal-output','删除控制信号',301,'/config/signal-output/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-signal-output','修改控制信号',301,'/config/signal-output','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-signal-output','查找控制信号',301,'/config/signal-output/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-signal-output','获取控制信号',301,'/config/signal-output/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-signal-output','获取全部控制信号',301,'/config/signal-output/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-signal-video','增加视频信号',301,'/config/signal-video','POST',1,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-signal-video','删除视频信号',301,'/config/signal-video/\\d{1,}$','DELETE',2,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-signal-video','修改视频信号',301,'/config/signal-video','PUT',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-signal-video','查找视频信号',301,'/config/signal-video/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-signal-video','获取视频信号',301,'/config/signal-video/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-signal-video','获取全部视频信号',301,'/config/signal-video/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-signal-threshold','增加信号门限',301,'/config/signal-threshhold','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-signal-threshold','删除信号门限',301,'/config/signal-threshhold/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-signal-threshold','修改信号门限',301,'/config/signal-threshhold','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-signal-threshold','查找信号门限',301,'/config/signal-threshhold/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-signal-threshold','获取信号门限',301,'/config/signal-threshhold/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-signal-threshold','获取全部信号门限',301,'/config/signal-threshhold/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-signal-unification','增加标准信号',304005,'/config/signal-unification','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-signal-unification','删除标准信号',304005,'/config/signal-unification/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-signal-unification','修改标准信号',304005,'/config/signal-unification','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-signal-unification','查找标准信号',304005,'/config/signal-unification/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-signal-unification','获取标准信号',304005,'/config/signal-unification/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-signal-unification','获取全部标准信号',304005,'/config/signal-unification/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-signal-virtual','增加虚拟信号',301,'/config/signal-virtual','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-signal-virtual','删除虚拟信号',301,'/config/signal-virtual/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-signal-virtual','修改虚拟信号',301,'/config/signal-virtual','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-signal-virtual','查找虚拟信号',301,'/config/signal-virtual/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-signal-virtual','获取虚拟信号',301,'/config/signal-virtual/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-signal-virtual','获取全部虚拟信号',301,'/config/signal-virtual/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-signal','增加信号',301,'/config/signal','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-signal','删除信号',301,'/config/signal/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-signal','修改信号',301,'/config/signal','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-signal','查找信号',301,'/config/signal/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-signal','获取信号',301,'/config/signal/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-signal','获取全部信号',301,'/config/signal/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-ai','增加AI',301,'/config/ai','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-ai','删除AI',301,'/config/ai/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-ai','修改AI',301,'/config/ai','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-ai','获取AI',301,'/config/ai/\\d{1,}$','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-di','增加DI',301,'/config/di','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-di','删除DI',301,'/config/di/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-di','修改DI',301,'/config/di','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-di','获取DI',301,'/config/di/\\d{1,}$','GET',4,0,NULL);

 INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-si','增加SI',301,'/config/si','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-si','删除SI',301,'/config/si/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-si','修改SI',301,'/config/si','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-si','获取SI',301,'/config/si/\\d{1,}$','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-vai','增加VAI',301,'/config/vai','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-vai','删除VAI',301,'/config/vai/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-vai','修改VAI',301,'/config/vai','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-vai','获取VAI',301,'/config/vai/\\d{1,}$','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-vdi','增加VDI',301,'/config/vdi','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-vdi','删除VDI',301,'/config/vdi/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-vdi','修改VDI',301,'/config/vdi','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-vdi','获取VDI',301,'/config/vdi/\\d{1,}$','GET',4,0,NULL);

 INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-vsi','增加VSI',301,'/config/vsi','POST',1,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-vsi','删除VSI',301,'/config/vsi/\\d{1,}$','DELETE',2,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-vsi','修改VSI',301,'/config/vsi','PUT',3,1,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-vsi','获取VSI',301,'/config/vsi/\\d{1,}$','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-ao','增加AO',301,'/config/ao','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-ao','删除AO',301,'/config/ao/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-ao','修改AO',301,'/config/ao','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-ao','获取AO',301,'/config/ao/\\d{1,}$','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-do','增加DO',301,'/config/do','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-do','删除DO',301,'/config/do/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-do','修改DO',301,'/config/do','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-do','获取DO',301,'/config/do/\\d{1,}$','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-so','增加SO',301,'/config/so','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-so','删除SO',301,'/config/so/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-so','修改SO',301,'/config/so','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-so','获取SO',301,'/config/so/\\d{1,}$','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-output','增加输出控制',301,'/config/output','POST',1,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-output','删除输出控制',301,'/config/output/\\d{1,}$','DELETE',2,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-output','修改输出控制',301,'/config/output','PUT',3,3,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-output','获取输出控制',301,'/config/output/\\d{1,}$','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-video','增加视频',301,'/config/video','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-video','删除视频',301,'/config/video/\\d{1,}$','DELETE',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-video','修改视频',301,'/config/video','PUT',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-video','获取视频',301,'/config/video/\\d{1,}$','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-user','增加用户',304007,'/user','POST',1,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-user','删除用户',304007,'/user/\\d{1,}$','DELETE',2,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-user-with-flag','标记删除用户',304007,'/user/\\d{1,}$','PUT',2,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-user','修改用户',304007,'/user','PUT',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('modify-user-password','修改用户密码',304007,'/user/password','PUT',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-user','查找用户',304007,'/user/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-user','获取用户',304007,'/user/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-user','获取全部用户',304007,'/user/all','GET',4,0,NULL);
 INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('verify-user','验证用户',304007,'/user/verify','POST',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('grant-user','用户授权',304009,'/user/grant-user','POST',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('grant-group','组授权',304009,'/user/grant-group','POST',3,0,NULL);
 INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('grant-user-batch','用户批量授权',304009,'/user/grant-user/batch','POST',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('grant-group-batch','组批量授权',304009,'/user/grant-group/batch','POST',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('get-join-group','获取用户加入组',304009,'/user/join-group/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('get-group-grant','获取组授权',304009,'/user/grant-group/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('get-user-grant','获取用户授权',304009,'/user/grant-user/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('get-grant-user-with-roles','获取用户组的授权',304009,
 '/user/grant-user-with-roles/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('join-group','用户加入组',304009,'/user/join-group','POST',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('remove-group','用户移除组',304009,'/user/remove-group','POST',3,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('create-action-group','增加功能分组',304010,'/config/action-group','POST',1,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('delete-action-group','删除功能分组',304010,'/config/action-group/\\d{1,}$','DELETE',2,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('update-action-group','修改功能分组',304010,'/config/action-group','PUT',3,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('query-action-group','查找功能分组',304010,'/config/action-group/query','POST',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('retrieve-action-group','获取功能分组',304010,'/config/action-group/\\d{1,}$','GET',4,0,NULL);
INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-action-group','获取全部功能分组',304010,'/config/action-group/all','GET',4,0,NULL);

INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('all-user-alias','获取全部用户的别名',304013,'/user/alias-all','GET',4,0,NULL);
 INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('logic-object-path','获取逻辑对象的父路径',304013,'/config/path/logic-object/\\d{1,}$','GET',4,0,NULL);
 INSERT INTO actions(action_name,action_alias,action_group_id,url,method,action_type,affected_server,created_by) 
 VALUES('signal-path','获取信号的父路径',304013,'/config/path/signal/\\d{1,}$','GET',4,0,NULL);