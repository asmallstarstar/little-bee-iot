DROP DATABASE IF EXISTS little_bee;
CREATE SCHEMA `little_bee` DEFAULT CHARACTER SET utf8mb4 ;

SET character_set_client='utf8';
SET character_set_connection='utf8';
SET character_set_database='utf8';
SET character_set_filesystem='binary';
SET character_set_results='utf8';
SET character_set_server='utf8';

DROP TABLE IF EXISTS `little_bee`.`logic_objects`;
CREATE TABLE `little_bee`.`logic_objects` (
  `logic_object_id` INT NOT NULL AUTO_INCREMENT,
  `logic_object_name` VARCHAR(128) NOT NULL,
  `logic_object_alias` VARCHAR(128) DEFAULT NULL,
  `industrial_internet_id` VARCHAR(256) DEFAULT NULL,
  `parent_logic_object_id` INT NOT NULL DEFAULT 0,
  `logic_object_type_id` INT DEFAULT NULL,
  `position_among_brothers` INT DEFAULT NULL,
  `metadata_id` INT DEFAULT NULL,
  `configure_id` INT DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`logic_object_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`logic_areas`;  
CREATE TABLE `little_bee`.`logic_areas` (
  `logic_object_id` int(11) NOT NULL,
  `area_alias` varchar(128) DEFAULT NULL,
  `area_type` int(11) DEFAULT NULL,
  PRIMARY KEY (`logic_object_id`)) 
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`logic_devices`; 
CREATE TABLE `little_bee`.`logic_devices` (
  `logic_object_id` int(11) NOT NULL,
  `device_alias` varchar(128) DEFAULT NULL,
  `device_type` int(11) DEFAULT NULL,
  PRIMARY KEY (`logic_object_id`)) 
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4;
  
DROP TABLE IF EXISTS `little_bee`.`area_types`;
CREATE TABLE `little_bee`.`area_types` (
  `area_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `area_type_name` varchar(128) NOT NULL,
  `area_type_note` varchar(128) DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`area_type_id`)) 
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`device_types`;
CREATE TABLE `little_bee`.`device_types` (
  `device_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `device_type_name` varchar(128) NOT NULL,
  `device_type_note` varchar(128) DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`device_type_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`signals`;
CREATE TABLE `little_bee`.`signals` (
  `signal_id` INT NOT NULL AUTO_INCREMENT,
  `signal_name` VARCHAR(128) NOT NULL,
  `signal_alias` VARCHAR(128) DEFAULT NULL,
  `industrial_internet_id` VARCHAR(256) DEFAULT NULL,
  `signal_unification_id` INT NULL,
  `signal_type` INT NOT NULL,
  `parent_logic_id` INT NOT NULL,
  `position_among_brothers` INT NULL DEFAULT 0,
  `applied` TINYINT(1) NULL DEFAULT 1,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`signal_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;
 
DROP TABLE IF EXISTS `little_bee`.`signal_ais`;
CREATE TABLE `little_bee`.`signal_ais` (
  `signal_id` INT NOT NULL,
  `record_period` INT NULL DEFAULT 0,
  `multiple_rate` FLOAT NULL DEFAULT 1.0,
  `offset_value` FLOAT NULL DEFAULT 0.0,
  `decimal_precision` INT NULL DEFAULT 2,
  `value_unit` VARCHAR(32),
  PRIMARY KEY (`signal_id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`signal_dis`; 
CREATE TABLE `little_bee`.`signal_dis` (
  `signal_id` INT NOT NULL,
  `is_flip` TINYINT(1) NULL,
  `record_period` INT NULL DEFAULT 0,
  `high_desc` VARCHAR(128) NULL,
  `low_desc` VARCHAR(128) NULL,
  PRIMARY KEY (`signal_id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`signal_sis`;
CREATE TABLE `little_bee`.`signal_sis` (
  `signal_id` int(11) NOT NULL,
  `record_period` int(11) DEFAULT '0',
  PRIMARY KEY (`signal_id`))
  ENGINE=InnoDB 
  DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`signal_acquisitions`;
CREATE TABLE `little_bee`.`signal_acquisitions` (
  `signal_id` INT NOT NULL,
  `agent_id` INT NOT NULL,
  `fsu_id` INT NOT NULL,
  `driver_id` INT NOT NULL DEFAULT 0,
  `x` INT  NULL DEFAULT NULL,
  `y` INT  NULL DEFAULT NULL,
  `z` INT  NULL DEFAULT NULL,
  `signal_bind_metadata_id` INT  NULL DEFAULT NULL,
  `signal_bind_configure_id` INT  NULL DEFAULT NULL,
  PRIMARY KEY (`signal_id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`signal_virtuals`;
CREATE TABLE `little_bee`.`signal_virtuals` (
  `signal_id` INT NOT NULL,
  `signal_express` TEXT(6144) NOT NULL,
  PRIMARY KEY (`signal_id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`signal_videos`;
CREATE TABLE `little_bee`.`signal_videos` (
  `signal_id` INT NOT NULL,
  `video_play_type` INT NOT NULL,
  `video_bind_metadata_id` INT NOT NULL,
  `video_bind_configure_id` INT NOT NULL,
  PRIMARY KEY (`signal_id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`signal_thresholds`;
CREATE TABLE `little_bee`.`signal_thresholds` (
  `threshold_id` INT NOT NULL AUTO_INCREMENT,
  `signal_id` INT NOT NULL,
  `signal_type` INT NOT NULL,
  `threshold_number` INT NOT NULL,
  `threshold_value_ai` FLOAT NOT NULL DEFAULT 0.0,
  `threshold_value_di` INT NOT NULL DEFAULT 1,
  `threshold_direction` INT NOT NULL DEFAULT 0,
  `threshold_dead_area` FLOAT NULL DEFAULT 0.0,
  `is_notify` TINYINT(1) NOT NULL DEFAULT 1,
  `alarm_level_number` INT NOT NULL DEFAULT 0,
  `up_notify_note` VARCHAR(128) NULL,
  `down_notify_note` VARCHAR(128) NULL,
  `alarm_delay` INT NULL DEFAULT 0,
  `recover_alarm_delay` INT NULL DEFAULT 0,
  PRIMARY KEY (`threshold_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`alarm_levels`;
CREATE TABLE `little_bee`.`alarm_levels` (
  `alarm_level_id` int(11) NOT NULL AUTO_INCREMENT,
  `alarm_level_number` int(11) NOT NULL,
  `alarm_level_name` varchar(128) NOT NULL,
  `alarm_level_alias` varchar(128) DEFAULT NULL,
  `alarm_level_note` varchar(128) DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`alarm_level_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`signal_unifications`;
CREATE TABLE `little_bee`.`signal_unifications` (
  `signal_unification_id` int(11) NOT NULL AUTO_INCREMENT,
  `signal_unification_name` varchar(128) NOT NULL,
  `signal_unification_alias` varchar(128) DEFAULT NULL,
  `signal_unification_related_signal_id` int(11) DEFAULT NULL,
  `signal_unification_note` varchar(128) DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`signal_unification_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`alarms`; 
CREATE TABLE `little_bee`.`alarms` (
  `alarm_serial_number` VARCHAR(128) NOT NULL,
  `signal_id` INT NOT NULL,
  `signal_name` VARCHAR(128) NOT NULL,
  `signal_type` INT NOT NULL,
  `signal_unification_id` INT NULL,
  `alarm_begin_time` DATETIME(6) NOT NULL,
  `alarm_end_time` DATETIME(6) NULL,
  `device_location` VARCHAR(128) NULL,
  `device_name` VARCHAR(128) NULL,
  `device_type_id` INT NULL,
  `alarm_desc` VARCHAR(32) NULL,
  `alarm_level_number` INT NULL,
  `alarm_threshhold_number` INT NULL,
  `signal_value_type` INT NULL,
  `occurred_digital_value` INT NULL,
  `occurred_analog_value` FLOAT NULL,
  `value_unit` VARCHAR(32) NULL,
  `occurred_string_value` VARCHAR(128) NULL,
  `ack_state` TINYINT(1) NULL DEFAULT 0 ,
  `ack_time` DATETIME(6) NULL,
  `ack_user_id` INT NULL,
  `ack_user_name` VARCHAR(64) NULL,
  PRIMARY KEY (`alarm_serial_number`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`metadatas`;  
CREATE TABLE `little_bee`.`metadatas` (
  `metadata_id` INT NOT NULL AUTO_INCREMENT,
  `metadata_name` VARCHAR(128) NOT NULL,
  `metadata_alias` VARCHAR(128) NOT NULL,
  `metadata_attribute_json` JSON NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`metadata_id`),
  UNIQUE KEY `metadata_name_UNIQUE` (`metadata_name`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`configures`;
CREATE TABLE `little_bee`.`configures` (
  `configure_id` INT NOT NULL AUTO_INCREMENT,
  `metadata_id` INT NULL,
  `configure_json` JSON NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`configure_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`control_command_records`; 
CREATE TABLE `little_bee`.`control_command_records` (
  `control_command_serial_number` varchar(128) NOT NULL,
  `operator_user_id` INT NOT NULL,
  `operator_user_name` varchar(128) NOT NULL,
  `operate_time` datetime(6) NOT NULL,
  `signal_id` int(11) NOT NULL,
  `control_command_type` int(11) NOT NULL,
  `control_command_result` int(11) DEFAULT NULL,
  `delay` int(11) DEFAULT NULL,
  `digital_value` int(11) DEFAULT NULL,
  `analog_value` float(11,2) DEFAULT NULL,
  `string_value` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`control_command_serial_number`))
  ENGINE=InnoDB
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`agents`;
CREATE TABLE `little_bee`.`agents` (
  `agent_id` INT NOT NULL AUTO_INCREMENT,
  `agent_name` VARCHAR(128) NOT NULL,
  `agent_note` VARCHAR(128) NULL,
  `metadata_id` INT NULL,
  `configure_id` INT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`agent_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`fsu_types`;
CREATE TABLE `little_bee`.`fsu_types` (
  `fsu_type_id` INT NOT NULL AUTO_INCREMENT,
  `fsu_type_name` VARCHAR(128) NOT NULL,
  `data_analysis_mode` INT NOT NULL DEFAULT 1,
  `new_instance_name` VARCHAR(128) NOT NULL,
  `metadata_id` INT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`fsu_type_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`fsus`;
CREATE TABLE `little_bee`.`fsus` (
  `fsu_id` INT NOT NULL AUTO_INCREMENT,
  `fsu_type_id` INT NOT NULL,
  `fsu_name` VARCHAR(128) NOT NULL,
  `metadata_id` INT NULL,
  `configure_id` INT NULL,
  `agent_id` INT NOT NULL,
  `related_metadata_id` INT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`fsu_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`driver_types`;
CREATE TABLE `little_bee`.`driver_types` (
  `driver_type_id` INT NOT NULL AUTO_INCREMENT,
  `driver_type_name` VARCHAR(128) NOT NULL,
  `driver_type_file_name` VARCHAR(128) NOT NULL,
  `driver_type_parameter_metadata_id` INT NULL,
  `driver_type_note` VARCHAR(128) NULL,
  `driver_type_acquisite_period` INT NOT NULL DEFAULT 5000,
  `driver_type_timeout_count` INT NOT NULL DEFAULT 3,
  `driver_type_reset_count` INT NOT NULL DEFAULT 4,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`driver_type_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`drivers`;
CREATE TABLE `little_bee`.`drivers` (
  `driver_id` INT NOT NULL AUTO_INCREMENT,
  `driver_type_id` INT NOT NULL,
  `fsu_id` INT NOT NULL,
  `related_flag_configure_id` INT NULL,
  `driver_name` varchar(128) NOT NULL,
  `driver_parameter_metadata_id` INT NULL,
  `driver_parameter_configure_id` INT NULL,
  `driver_acquisite_period` INT NOT NULL DEFAULT 5000,
  `driver_timeout_count` INT NOT NULL DEFAULT 3,
  `driver_reset_count` INT NOT NULL DEFAULT 4,
  `is_write_log` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`driver_id`))
  ENGINE = InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`departments`;
CREATE TABLE `little_bee`.`departments` (
  `department_id` INT NOT NULL AUTO_INCREMENT,
  `parent_department_id` INT NOT NULL DEFAULT 0,
  `department_name` VARCHAR(128) NOT NULL,
  `position_among_brothers` INT DEFAULT NULL,
  `leader_name` varchar(128) DEFAULT NULL,
  `phone` varchar(32) DEFAULT NULL,
  `email` varchar(64) DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`department_id`) ) 
  ENGINE=InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`users`;
CREATE TABLE `little_bee`.`users` (
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `department_id` INT DEFAULT NULL,
  `user_name` VARCHAR(128) NOT NULL,
  `user_nick` VARCHAR(128) NOT NULL,
  `phone` VARCHAR(32) DEFAULT NULL,
  `email` VARCHAR(64) DEFAULT NULL,
  `sex` TINYINT(1)  DEFAULT 1 COMMENT '0:MALE,1:FEMALE,2:UNKNOWN',
  `avatar` VARCHAR(128) DEFAULT '',
  `password` VARCHAR(128) NOT NULL,
  `status` TINYINT(1) DEFAULT 0 COMMENT 'account status. 0:NORMAL,1:DEACTIVATED',
  `remark` VARCHAR(512) DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`user_id`) ) 
  ENGINE=InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`menus`;
CREATE TABLE `little_bee`.`menus` (
  `menu_id` INT NOT NULL,
  `parent_menu_id` INT DEFAULT NULL,
  `menu_name` VARCHAR(128) NOT NULL,
  `menu_alias` VARCHAR(128) NOT NULL,
  `menu_level` INT NOT NULL,
  `path` VARCHAR(256) DEFAULT NULL,
  `url` VARCHAR(256) DEFAULT NULL,
  `component` VARCHAR(64) DEFAULT NULL,
  `title` VARCHAR(64) DEFAULT NULL,
  `icon` VARCHAR(64) DEFAULT NULL,
  `sidebar` VARCHAR(64) DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`menu_id`) ) 
  ENGINE=InnoDB 
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`action_groups`;
CREATE TABLE `little_bee`.`action_groups` (
  `action_group_id` INT NOT NULL,
  `parent_action_group_id` INT DEFAULT NULL,
  `action_group_name` VARCHAR(128) NOT NULL,
  `action_group_alias` VARCHAR(128) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`action_group_id`) ) 
  ENGINE=InnoDB 
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`actions`;
CREATE TABLE `little_bee`.`actions` (
  `action_id` INT NOT NULL AUTO_INCREMENT,
  `action_name` VARCHAR(128) NOT NULL,
  `action_alias` VARCHAR(128) NOT NULL,
  `action_group_id` INT NOT NULL,
  `url` VARCHAR(256) NOT NULL,
  `method` VARCHAR(32) NOT NULL,
  `action_type` INT NOT NULL,
  `affected_server` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`action_id`) ) 
  ENGINE=InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`roles`;
CREATE TABLE `little_bee`.`roles` (
  `role_id` INT NOT NULL AUTO_INCREMENT,
  `role_name` VARCHAR(128) NOT NULL,
  `role_alias`VARCHAR(128) NOT NULL,
  `role_note` VARCHAR(256) DEFAULT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`role_id`) ) 
  ENGINE=InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`menu_actions`;
CREATE TABLE `little_bee`.`menu_actions` (
  `menu_action_id` INT NOT NULL AUTO_INCREMENT,
  `menu_name` VARCHAR(128) NOT NULL,
  `action_name` VARCHAR(128) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT DEFAULT NULL,
  `updated_at` DATETIME DEFAULT NULL,
  `updated_by` INT DEFAULT NULL,
  `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
  `deleted_at` DATETIME DEFAULT NULL,
  `deleted_by` INT DEFAULT NULL,
  PRIMARY KEY (`menu_action_id`) ) 
  ENGINE=InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `little_bee`.`operation_logs`;
CREATE TABLE `little_bee`.`operation_logs` (
  `log_id` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `user_nick` VARCHAR(128) NOT NULL,
  `action_id` INT NOT NULL,
  `action_name` VARCHAR(128) NOT NULL,
  `action_alias` VARCHAR(128) NOT NULL,
  `action_type` INT NOT NULL,
  `operation_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`log_id`) ) 
  ENGINE=InnoDB AUTO_INCREMENT 0
  DEFAULT CHARACTER SET = utf8mb4;