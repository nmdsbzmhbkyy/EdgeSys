-- MySQL dump 10.13  Distrib 8.0.29, for Win64 (x86_64)
--
-- Host: localhost    Database: edge_sys
-- ------------------------------------------------------
-- Server version	8.0.29

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `edge_sys`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `edge_sys` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `edge_sys`;

--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `casbin_rule` (
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10604 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule` DISABLE KEYS */;
INSERT INTO `casbin_rule` VALUES ('p','test','/system/api/list','GET','','','',8679),('p','test','/system/api/all','GET','','','',8680),('p','test','/system/api/getPolicyPathByRoleId','GET','','','',8681),('p','test','/system/api/:id','GET','','','',8682),('p','test','/system/api','POST','','','',8683),('p','test','/system/api','PUT','','','',8684),('p','test','/system/api/:id','DELETE','','','',8685),('p','manage','/system/api/list','GET','','','',10162),('p','manage','/system/api/all','GET','','','',10163),('p','manage','/system/api/getPolicyPathByRoleId','GET','','','',10164),('p','manage','/system/api/:id','GET','','','',10165),('p','manage','/system/config/list','GET','','','',10166),('p','manage','/system/config/configKey','GET','','','',10167),('p','manage','/system/config/:configId','GET','','','',10168),('p','manage','/system/dict/type/list','GET','','','',10183),('p','manage','/system/dict/type/:dictId','GET','','','',10184),('p','manage','/system/dict/data/list','GET','','','',10185),('p','manage','/system/dict/data/type','GET','','','',10186),('p','manage','/system/dict/data/:dictCode','GET','','','',10187),('p','manage','/develop/code/table/db/list','GET','','','',10188),('p','manage','/develop/code/table/list','GET','','','',10189),('p','manage','/develop/code/table/info/:tableId','GET','','','',10190),('p','manage','/develop/code/table/info/tableName','GET','','','',10191),('p','manage','/develop/code/table/tableTree','GET','','','',10192),('p','manage','/develop/code/gen/preview/:tableId','GET','','','',10193),('p','manage','/job/list','GET','','','',10194),('p','manage','/job/:jobId','GET','','','',10195),('p','manage','/log/logLogin/list','GET','','','',10196),('p','manage','/log/logOper/list','GET','','','',10197),('p','manage','/system/menu/menuTreeSelect','GET','','','',10198),('p','manage','/system/menu/menuRole','GET','','','',10199),('p','manage','/system/menu/roleMenuTreeSelect/:roleId','GET','','','',10200),('p','manage','/system/menu/menuPaths','GET','','','',10201),('p','manage','/system/menu/list','GET','','','',10202),('p','manage','/system/menu/:menuId','GET','','','',10203),('p','manage','/system/notice/list','GET','','','',10204),('p','manage','/system/organization/list','GET','','','',10205),('p','manage','/system/organization/:organizationId','GET','','','',10206),('p','manage','/system/organization/roleOrganizationTreeSelect/:roleId','GET','','','',10207),('p','manage','/system/organization/organizationTree','GET','','','',10208),('p','manage','/system/post/list','GET','','','',10211),('p','manage','/system/post/:postId','GET','','','',10212),('p','manage','/system/role/list','GET','','','',10222),('p','manage','/system/role/:roleId','GET','','','',10223),('p','manage','/system/user/list','GET','','','',10237),('p','manage','/system/user/getById/:userId','GET','','','',10238),('p','manage','/system/user/getInit','GET','','','',10239),('p','manage','/system/user/getRoPo','GET','','','',10240),('p','admin','/system/api/list','GET','','','',10425),('p','admin','/system/api/all','GET','','','',10426),('p','admin','/system/api/getPolicyPathByRoleId','GET','','','',10427),('p','admin','/system/api/:id','GET','','','',10428),('p','admin','/system/api','POST','','','',10429),('p','admin','/system/api','PUT','','','',10430),('p','admin','/system/api/:id','DELETE','','','',10431),('p','admin','/system/config/list','GET','','','',10432),('p','admin','/system/config/configKey','GET','','','',10433),('p','admin','/system/config/:configId','GET','','','',10434),('p','admin','/system/config','POST','','','',10435),('p','admin','/system/config','PUT','','','',10436),('p','admin','/system/config/:configId','DELETE','','','',10437),('p','admin','/system/dict/type/list','GET','','','',10464),('p','admin','/system/dict/type/:dictId','GET','','','',10465),('p','admin','/system/dict/type','POST','','','',10466),('p','admin','/system/dict/type','PUT','','','',10467),('p','admin','/system/dict/type/:dictId','DELETE','','','',10468),('p','admin','/system/dict/type/export','GET','','','',10469),('p','admin','/system/dict/data/list','GET','','','',10470),('p','admin','/system/dict/data/type','GET','','','',10471),('p','admin','/system/dict/data/:dictCode','GET','','','',10472),('p','admin','/system/dict/data','POST','','','',10473),('p','admin','/system/dict/data','PUT','','','',10474),('p','admin','/system/dict/data/:dictCode','DELETE','','','',10475),('p','admin','/develop/code/table/db/list','GET','','','',10476),('p','admin','/develop/code/table/list','GET','','','',10477),('p','admin','/develop/code/table/info/:tableId','GET','','','',10478),('p','admin','/develop/code/table/info/tableName','GET','','','',10479),('p','admin','/develop/code/table/tableTree','GET','','','',10480),('p','admin','/develop/code/table','POST','','','',10481),('p','admin','/develop/code/table','PUT','','','',10482),('p','admin','/develop/code/table/:tableId','DELETE','','','',10483),('p','admin','/develop/code/gen/preview/:tableId','GET','','','',10484),('p','admin','/develop/code/gen/code/:tableId','GET','','','',10485),('p','admin','/develop/code/gen/configure/:tableId','GET','','','',10486),('p','admin','/job/list','GET','','','',10487),('p','admin','/job','POST','','','',10488),('p','admin','/job','PUT','','','',10489),('p','admin','/job/:jobId','GET','','','',10490),('p','admin','/job/:jobId','DELETE','','','',10491),('p','admin','/job/stop/:jobId','GET','','','',10492),('p','admin','/job/start/:jobId','GET','','','',10493),('p','admin','/job/log/list','GET','','','',10494),('p','admin','/job/log/all','DELETE','','','',10495),('p','admin','/job/log/:logId','DELETE','','','',10496),('p','admin','/job/changeStatus','PUT','','','',10497),('p','admin','/log/logLogin/list','GET','','','',10498),('p','admin','/log/logLogin/:infoId','DELETE','','','',10499),('p','admin','/log/logLogin/all','DELETE','','','',10500),('p','admin','/log/logOper/list','GET','','','',10501),('p','admin','/log/logOper/:operId','DELETE','','','',10502),('p','admin','/log/logOper/all','DELETE','','','',10503),('p','admin','/system/menu/menuTreeSelect','GET','','','',10504),('p','admin','/system/menu/menuRole','GET','','','',10505),('p','admin','/system/menu/roleMenuTreeSelect/:roleId','GET','','','',10506),('p','admin','/system/menu/menuPaths','GET','','','',10507),('p','admin','/system/menu/list','GET','','','',10508),('p','admin','/system/menu/:menuId','GET','','','',10509),('p','admin','/system/menu','POST','','','',10510),('p','admin','/system/menu','PUT','','','',10511),('p','admin','/system/menu/:menuId','DELETE','','','',10512),('p','admin','/system/notice/list','GET','','','',10513),('p','admin','/system/notice','POST','','','',10514),('p','admin','/system/notice','PUT','','','',10515),('p','admin','/system/notice/:noticeId','DELETE','','','',10516),('p','admin','/system/organization/list','GET','','','',10517),('p','admin','/system/organization/:organizationId','GET','','','',10518),('p','admin','/system/organization/roleOrganizationTreeSelect/:roleId','GET','','','',10519),('p','admin','/system/organization/organizationTree','GET','','','',10520),('p','admin','/system/organization','POST','','','',10521),('p','admin','/system/organization','PUT','','','',10522),('p','admin','/system/organization/:organizationId','DELETE','','','',10523),('p','admin','/system/post/list','GET','','','',10529),('p','admin','/system/post/:postId','GET','','','',10530),('p','admin','/system/post','POST','','','',10531),('p','admin','/system/post','PUT','','','',10532),('p','admin','/system/post/:postId','DELETE','','','',10533),('p','admin','/system/role/list','GET','','','',10549),('p','admin','/system/role/:roleId','GET','','','',10550),('p','admin','/system/role','POST','','','',10551),('p','admin','/system/role','PUT','','','',10552),('p','admin','/system/role/:roleId','DELETE','','','',10553),('p','admin','/system/role/changeStatus','PUT','','','',10554),('p','admin','/system/role/dataScope','PUT','','','',10555),('p','admin','/system/role/export','GET','','','',10556),('p','admin','/upload/up','POST','','','',10587),('p','admin','/system/user/list','GET','','','',10588),('p','admin','/system/user/changeStatus','PUT','','','',10589),('p','admin','/system/user/:userId','DELETE','','','',10590),('p','admin','/system/user/avatar','POST','','','',10591),('p','admin','/system/user/pwd','PUT','','','',10592),('p','admin','/system/user/getById/:userId','GET','','','',10593),('p','admin','/system/user/getInit','GET','','','',10594),('p','admin','/system/user/getRoPo','GET','','','',10595),('p','admin','/system/user','POST','','','',10596),('p','admin','/system/user','PUT','','','',10597),('p','admin','/system/user/export','GET','','','',10598);
/*!40000 ALTER TABLE `casbin_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dev_gen_table_columns`
--

DROP TABLE IF EXISTS `dev_gen_table_columns`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dev_gen_table_columns` (
  `column_id` bigint NOT NULL AUTO_INCREMENT,
  `org_id` bigint DEFAULT NULL COMMENT '机构ID',
  `owner` varchar(64) DEFAULT NULL COMMENT '创建者,所有者',
  `table_id` bigint DEFAULT NULL,
  `table_name` varchar(191) DEFAULT NULL,
  `column_name` varchar(191) DEFAULT NULL,
  `column_comment` varchar(191) DEFAULT NULL,
  `column_type` varchar(191) DEFAULT NULL,
  `column_key` varchar(191) DEFAULT NULL,
  `go_type` varchar(191) DEFAULT NULL,
  `go_field` varchar(191) DEFAULT NULL,
  `json_field` varchar(191) DEFAULT NULL,
  `html_field` varchar(191) DEFAULT NULL,
  `is_pk` varchar(191) DEFAULT NULL,
  `is_increment` varchar(191) DEFAULT NULL,
  `is_required` varchar(191) DEFAULT NULL,
  `is_insert` varchar(191) DEFAULT NULL,
  `is_edit` varchar(191) DEFAULT NULL,
  `is_list` varchar(191) DEFAULT NULL,
  `is_query` varchar(191) DEFAULT NULL,
  `query_type` varchar(191) DEFAULT NULL,
  `html_type` varchar(191) DEFAULT NULL,
  `dict_type` varchar(191) DEFAULT NULL,
  `sort` bigint DEFAULT NULL,
  `link_table_name` varchar(191) DEFAULT NULL,
  `link_table_class` varchar(191) DEFAULT NULL,
  `link_table_package` varchar(191) DEFAULT NULL,
  `link_label_id` varchar(191) DEFAULT NULL,
  `link_label_name` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`column_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dev_gen_table_columns`
--

LOCK TABLES `dev_gen_table_columns` WRITE;
/*!40000 ALTER TABLE `dev_gen_table_columns` DISABLE KEYS */;
/*!40000 ALTER TABLE `dev_gen_table_columns` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `dev_gen_tables`
--

DROP TABLE IF EXISTS `dev_gen_tables`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `dev_gen_tables` (
  `table_id` bigint NOT NULL AUTO_INCREMENT,
  `org_id` bigint DEFAULT NULL COMMENT '机构ID',
  `owner` varchar(64) DEFAULT NULL COMMENT '创建者,所有者',
  `table_name` varchar(191) DEFAULT NULL,
  `table_comment` varchar(191) DEFAULT NULL,
  `class_name` varchar(191) DEFAULT NULL,
  `tpl_category` varchar(191) DEFAULT NULL,
  `package_name` varchar(191) DEFAULT NULL,
  `module_name` varchar(191) DEFAULT NULL,
  `business_name` varchar(191) DEFAULT NULL,
  `function_name` varchar(191) DEFAULT NULL,
  `function_author` varchar(191) DEFAULT NULL,
  `options` varchar(191) DEFAULT NULL,
  `remark` varchar(191) DEFAULT NULL,
  `pk_column` varchar(191) DEFAULT NULL,
  `pk_go_field` varchar(191) DEFAULT NULL,
  `pk_go_type` varchar(191) DEFAULT NULL,
  `pk_json_field` varchar(191) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`table_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dev_gen_tables`
--

LOCK TABLES `dev_gen_tables` WRITE;
/*!40000 ALTER TABLE `dev_gen_tables` DISABLE KEYS */;
/*!40000 ALTER TABLE `dev_gen_tables` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `job_logs`
--

DROP TABLE IF EXISTS `job_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `job_logs` (
  `id` varchar(64) NOT NULL,
  `owner` varchar(64) DEFAULT NULL COMMENT '创建者,所有者',
  `org_id` varchar(64) DEFAULT NULL COMMENT '机构ID',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL COMMENT '任务名称',
  `entry_id` int DEFAULT NULL COMMENT '任务id',
  `target_invoke` varchar(64) DEFAULT NULL COMMENT '调用方法',
  `log_info` varchar(255) DEFAULT NULL COMMENT '日志信息',
  `status` varchar(1) DEFAULT NULL COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `job_logs`
--

LOCK TABLES `job_logs` WRITE;
/*!40000 ALTER TABLE `job_logs` DISABLE KEYS */;
INSERT INTO `job_logs` VALUES ('74GZAOIGuA','','2','2023-11-11 00:00:00','2023-11-11 00:00:00','作业日志001',1,'cronJobAdown','任务运行成功，总耗时 0.001264','0');
/*!40000 ALTER TABLE `job_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `jobs`
--

DROP TABLE IF EXISTS `jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `jobs` (
  `id` varchar(64) NOT NULL,
  `owner` varchar(64) DEFAULT NULL COMMENT '创建者,所有者',
  `org_id` int DEFAULT NULL COMMENT '机构ID',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `job_name` varchar(64) DEFAULT NULL COMMENT '名称',
  `target_invoke` varchar(64) DEFAULT NULL COMMENT '调用目标',
  `target_args` varchar(64) DEFAULT NULL COMMENT '目标传参',
  `job_content` varchar(255) DEFAULT NULL COMMENT '目标传参 要执行的内容',
  `cron_expression` varchar(64) DEFAULT NULL COMMENT 'cron表达式',
  `misfire_policy` varchar(1) DEFAULT NULL COMMENT '执行策略',
  `status` varchar(1) DEFAULT NULL COMMENT '状态',
  `entry_id` int DEFAULT NULL COMMENT 'job启动时返回的id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobs`
--

LOCK TABLES `jobs` WRITE;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
INSERT INTO `jobs` VALUES ('wvz4D6CXSw','admin',2,'2023-11-11 00:00:00','2023-11-11 00:00:00','adsa','cronJobAdown','d_1928b99619910dae5a001fa7','{\"设备下发\":\"010100000\"}',' 0/10 * * * * ?','1','0',0);
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_logins`
--

DROP TABLE IF EXISTS `log_logins`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `log_logins` (
  `info_id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(128) DEFAULT NULL COMMENT '用户名',
  `status` varchar(1) DEFAULT NULL COMMENT '状态',
  `ipaddr` varchar(255) DEFAULT NULL COMMENT 'ip地址',
  `login_location` varchar(255) DEFAULT NULL COMMENT '归属地',
  `browser` varchar(255) DEFAULT NULL COMMENT '浏览器',
  `os` varchar(255) DEFAULT NULL COMMENT '系统',
  `platform` varchar(255) DEFAULT NULL COMMENT '固件',
  `login_time` timestamp NULL DEFAULT NULL COMMENT '登录时间',
  `create_by` varchar(128) DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(128) DEFAULT NULL COMMENT '更新者',
  `remark` varchar(255) DEFAULT NULL,
  `msg` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_logins`
--

LOCK TABLES `log_logins` WRITE;
/*!40000 ALTER TABLE `log_logins` DISABLE KEYS */;
INSERT INTO `log_logins` VALUES (1,'admin','0','127.0.0.1:58941','您的IP或者域名无法解析，请核实后再查询！','Chrome 94.0.4606.71','Windows 10','Windows','2023-11-19 11:30:18','admin','','Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36','登录成功','2023-11-19 19:30:18','2023-11-19 19:30:18',NULL),(2,'admin','0','127.0.0.1:52313','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 11:34:50','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 19:34:50','2023-11-19 19:34:50',NULL),(3,'admin','0','127.0.0.1:55705','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 11:44:25','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 19:44:25','2023-11-19 19:44:25',NULL),(4,'admin','0','127.0.0.1:55707','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 11:44:43','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 19:44:43','2023-11-19 19:44:43',NULL),(5,'admin','0','127.0.0.1:55706','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 11:46:57','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 19:46:57','2023-11-19 19:46:57',NULL),(6,'admin','0','127.0.0.1:55704','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 11:51:48','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 19:51:48','2023-11-19 19:51:48',NULL),(7,'admin','0','127.0.0.1:55707','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 11:52:06','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 19:52:06','2023-11-19 19:52:06',NULL),(8,'admin','0','127.0.0.1:63381','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 12:03:46','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 20:03:46','2023-11-19 20:03:46',NULL),(9,'admin','0','127.0.0.1:58762','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 12:07:25','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 20:07:25','2023-11-19 20:07:25',NULL),(10,'admin','0','127.0.0.1:60534','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 14:42:36','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 22:42:36','2023-11-19 22:42:36',NULL),(11,'admin','0','127.0.0.1:50790','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 15:00:13','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 23:00:13','2023-11-19 23:00:13',NULL),(12,'admin','0','127.0.0.1:56994','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 15:27:38','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 23:27:38','2023-11-19 23:27:38',NULL),(13,'admin','0','127.0.0.1:53098','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 15:50:21','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 23:50:21','2023-11-19 23:50:21',NULL),(14,'admin','0','127.0.0.1:54996','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 15:56:07','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 23:56:07','2023-11-19 23:56:07',NULL),(15,'admin','0','127.0.0.1:54996','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 15:56:45','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-19 23:56:45','2023-11-19 23:56:45',NULL),(16,'admin','0','127.0.0.1:64437','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 16:03:28','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-20 00:03:28','2023-11-20 00:03:28',NULL),(17,'admin','0','127.0.0.1:62932','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 16:28:58','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-20 00:28:58','2023-11-20 00:28:58',NULL),(18,'admin','0','127.0.0.1:62932','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 16:29:16','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-20 00:29:16','2023-11-20 00:29:16',NULL),(19,'admin','0','127.0.0.1:62932','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 16:29:44','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-20 00:29:44','2023-11-20 00:29:44',NULL),(20,'admin','0','127.0.0.1:62765','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 16:43:17','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-20 00:43:17','2023-11-20 00:43:17',NULL),(21,'admin','0','127.0.0.1:62765','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 16:45:26','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-20 00:45:26','2023-11-20 00:45:26',NULL),(22,'admin','0','127.0.0.1:62765','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-19 16:46:07','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-20 00:46:07','2023-11-20 00:46:07',NULL),(23,'admin','0','127.0.0.1:57360','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-20 02:26:22','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-20 10:26:22','2023-11-20 10:26:22',NULL),(24,'admin','0','127.0.0.1:65443','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-20 02:35:09','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-20 10:35:09','2023-11-20 10:35:09',NULL),(25,'admin','0','127.0.0.1:60255','您的IP或者域名无法解析，请核实后再查询！','Chrome 112.0.0.0','Windows 10','Windows','2023-11-20 03:33:00','admin','','Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36','登录成功','2023-11-20 11:33:00','2023-11-20 11:33:00',NULL);
/*!40000 ALTER TABLE `log_logins` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_opers`
--

DROP TABLE IF EXISTS `log_opers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `log_opers` (
  `oper_id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(128) DEFAULT NULL COMMENT '操作的模块',
  `business_type` varchar(1) DEFAULT NULL COMMENT '0其它 1新增 2修改 3删除',
  `method` varchar(255) DEFAULT NULL COMMENT '请求方法',
  `oper_name` varchar(255) DEFAULT NULL COMMENT '操作人员',
  `oper_url` varchar(255) DEFAULT NULL COMMENT '操作url',
  `oper_ip` varchar(255) DEFAULT NULL COMMENT '操作IP',
  `oper_location` varchar(255) DEFAULT NULL COMMENT '操作地点',
  `oper_param` varchar(255) DEFAULT NULL COMMENT '请求参数',
  `status` varchar(1) DEFAULT NULL COMMENT '0=正常,1=异常',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_opers`
--

LOCK TABLES `log_opers` WRITE;
/*!40000 ALTER TABLE `log_opers` DISABLE KEYS */;
INSERT INTO `log_opers` VALUES (1,'删除组织信息','3','DELETE','admin','/system/organization/3','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:26:33','2023-11-19 21:26:33',NULL),(2,'删除组织信息','3','DELETE','admin','/system/organization/2','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:26:37','2023-11-19 21:26:37',NULL),(3,'添加组织信息','1','POST','admin','/system/organization','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:27:25','2023-11-19 21:27:25',NULL),(4,'添加组织信息','1','POST','admin','/system/organization','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:27:56','2023-11-19 21:27:56',NULL),(5,'添加组织信息','1','POST','admin','/system/organization','127.0.0.1:53485','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:28:08','2023-11-19 21:28:08',NULL),(6,'添加组织信息','1','POST','admin','/system/organization','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:28:32','2023-11-19 21:28:32',NULL),(7,'添加组织信息','1','POST','admin','/system/organization','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:28:56','2023-11-19 21:28:56',NULL),(8,'添加组织信息','1','POST','admin','/system/organization','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:29:07','2023-11-19 21:29:07',NULL),(9,'添加组织信息','1','POST','admin','/system/organization','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:29:29','2023-11-19 21:29:29',NULL),(10,'修改岗位信息','2','PUT','admin','/system/post','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:30:29','2023-11-19 21:30:29',NULL),(11,'修改岗位信息','2','PUT','admin','/system/post','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:30:48','2023-11-19 21:30:48',NULL),(12,'添加岗位信息','1','POST','admin','/system/post','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:32:00','2023-11-19 21:32:00',NULL),(13,'添加岗位信息','1','POST','admin','/system/post','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:33:05','2023-11-19 21:33:05',NULL),(14,'修改岗位信息','2','PUT','admin','/system/post','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:33:11','2023-11-19 21:33:11',NULL),(15,'修改岗位信息','2','PUT','admin','/system/post','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:33:14','2023-11-19 21:33:14',NULL),(16,'添加岗位信息','1','POST','admin','/system/post','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:33:48','2023-11-19 21:33:48',NULL),(17,'修改组织信息','2','PUT','admin','/system/organization','127.0.0.1:65448','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:35:46','2023-11-19 21:35:46',NULL),(18,'修改组织信息','2','PUT','admin','/system/organization','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:35:53','2023-11-19 21:35:53',NULL),(19,'修改组织信息','2','PUT','admin','/system/organization','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:35:59','2023-11-19 21:35:59',NULL),(20,'修改组织信息','2','PUT','admin','/system/organization','127.0.0.1:65448','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:36:05','2023-11-19 21:36:05',NULL),(21,'修改组织信息','2','PUT','admin','/system/organization','127.0.0.1:65448','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:36:16','2023-11-19 21:36:16',NULL),(22,'修改组织信息','2','PUT','admin','/system/organization','127.0.0.1:65448','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:36:22','2023-11-19 21:36:22',NULL),(23,'修改组织信息','2','PUT','admin','/system/organization','127.0.0.1:64399','您的IP或者域名无法解析，请核实后再查询！','','0','2023-11-19 21:36:28','2023-11-19 21:36:28',NULL);
/*!40000 ALTER TABLE `log_opers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_apis`
--

DROP TABLE IF EXISTS `sys_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_apis` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  `path` varchar(191) DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) DEFAULT NULL COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=99 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_apis`
--

LOCK TABLES `sys_apis` WRITE;
/*!40000 ALTER TABLE `sys_apis` DISABLE KEYS */;
INSERT INTO `sys_apis` VALUES (1,'2023-11-11 00:00:00','2023-11-11 00:00:00',NULL,'/system/user/list','查询用户列表（分页）','user','GET'),(2,'2023-11-11 00:00:00','2023-11-11 00:00:00',NULL,'/system/user/changeStatus','修改用户状态','user','PUT'),(3,'2023-11-11 00:00:00','2023-11-11 00:00:00',NULL,'/system/user/:userId','删除用户信息','user','DELETE'),(4,'2023-11-11 00:00:00','2023-11-11 00:00:00',NULL,'/system/organization/list','获取组织列表','organization','GET'),(5,'2023-11-11 00:00:00','2023-11-11 00:00:00',NULL,'/system/organization/:organizationId','获取组织信息','organization','GET'),(6,'2023-11-11 00:00:00','2023-11-11 00:00:00',NULL,'/system/user/avatar','上传头像','user','POST'),(7,'2023-11-11 00:00:00','2023-11-11 00:00:00',NULL,'/system/user/pwd','修改密码','user','PUT'),(8,'2023-11-11 00:00:00','2023-11-11 00:00:00',NULL,'/system/user/getById/:userId','通过id获取用户信息','user','GET'),(9,'2023-11-11 00:00:00','2023-11-11 00:00:00',NULL,'/system/user/getInit','获取初始化角色岗位信息(添加用户初始化)','user','GET'),(10,'2021-12-09 18:59:43','2021-12-09 18:59:43',NULL,'/system/user/getRoPo','获取用户角色岗位信息','user','GET'),(11,'2021-12-09 19:00:24','2021-12-09 19:00:24',NULL,'/system/user','添加用户信息','user','POST'),(12,'2021-12-09 19:00:52','2021-12-09 19:00:52',NULL,'/system/user','修改用户信息','user','PUT'),(13,'2021-12-09 19:02:30','2021-12-09 19:02:30',NULL,'/system/user/export','导出用户信息','user','GET'),(14,'2021-12-09 19:04:04','2023-09-14 14:06:35',NULL,'/system/organization/roleOrganizationTreeSelect/:roleId','获取角色部门树','organization','GET'),(15,'2021-12-09 19:04:48','2023-09-14 14:07:06',NULL,'/system/organization/organizationTree','获取所有组织树','organization','GET'),(16,'2021-12-09 19:07:37','2023-09-14 14:07:18',NULL,'/system/organization','添加组织信息','organization','POST'),(17,'2021-12-09 19:08:14','2023-09-14 14:07:28',NULL,'/system/organization','修改组织信息','organization','PUT'),(18,'2021-12-09 19:08:40','2023-09-14 14:07:41',NULL,'/system/organization/:organizationId','删除组织信息','organization','DELETE'),(19,'2021-12-09 19:09:41','2021-12-09 19:09:41',NULL,'/system/config/list','获取配置分页列表','config','GET'),(20,'2021-12-09 19:10:11','2021-12-09 19:10:11',NULL,'/system/config/configKey','获取配置列表通过ConfigKey','config','GET'),(21,'2021-12-09 19:10:45','2021-12-09 19:10:45',NULL,'/system/config/:configId','获取配置信息','config','GET'),(22,'2021-12-09 19:11:22','2021-12-09 19:11:22',NULL,'/system/config','添加配置信息','config','POST'),(23,'2021-12-09 19:11:41','2021-12-09 19:11:41',NULL,'/system/config','修改配置信息','config','PUT'),(24,'2021-12-09 19:12:28','2021-12-09 19:12:28',NULL,'/system/config/:configId','删除配置信息','config','DELETE'),(25,'2021-12-09 19:13:08','2021-12-09 19:13:08',NULL,'/system/dict/type/list','获取字典类型分页列表','dict','GET'),(26,'2021-12-09 19:13:55','2021-12-09 19:13:55',NULL,'/system/dict/type/:dictId','获取字典类型信息','dict','GET'),(27,'2021-12-09 19:14:28','2021-12-09 19:14:28',NULL,'/system/dict/type','添加字典类型信息','dict','POST'),(28,'2021-12-09 19:14:55','2021-12-09 19:14:55',NULL,'/system/dict/type','修改字典类型信息','dict','PUT'),(29,'2021-12-09 19:15:17','2021-12-09 19:15:17',NULL,'/system/dict/type/:dictId','删除字典类型信息','dict','DELETE'),(30,'2021-12-09 19:15:50','2021-12-09 19:15:50',NULL,'/system/dict/type/export','导出字典类型信息','dict','GET'),(31,'2021-12-09 19:16:26','2021-12-09 19:16:26',NULL,'/system/dict/data/list','获取字典数据分页列表','dict','GET'),(32,'2021-12-09 19:17:01','2021-12-09 19:17:01',NULL,'/system/dict/data/type','获取字典数据列表通过字典类型','dict','GET'),(33,'2021-12-09 19:17:39','2021-12-09 19:17:39',NULL,'/system/dict/data/:dictCode','获取字典数据信息','dict','GET'),(34,'2021-12-09 19:18:20','2021-12-09 19:18:20',NULL,'/system/dict/data','添加字典数据信息','dict','POST'),(35,'2021-12-09 19:18:44','2021-12-09 19:18:44',NULL,'/system/dict/data','修改字典数据信息','dict','PUT'),(36,'2021-12-09 19:19:16','2021-12-09 19:19:16',NULL,'/system/dict/data/:dictCode','删除字典数据信息','dict','DELETE'),(37,'2021-12-09 19:21:18','2021-12-09 19:21:18',NULL,'/system/menu/menuTreeSelect','获取菜单树','menu','GET'),(38,'2021-12-09 19:21:47','2021-12-09 19:21:47',NULL,'/system/menu/menuRole','获取角色菜单','menu','GET'),(39,'2021-12-09 19:22:42','2021-12-09 19:22:42',NULL,'/system/menu/roleMenuTreeSelect/:roleId','获取角色菜单树','menu','GET'),(40,'2021-12-09 19:23:17','2021-12-09 19:23:17',NULL,'/system/menu/menuPaths','获取角色菜单路径列表','menu','GET'),(41,'2021-12-09 19:23:40','2021-12-09 19:23:40',NULL,'/system/menu/list','获取菜单列表','menu','GET'),(42,'2021-12-09 19:24:09','2021-12-09 19:24:09',NULL,'/system/menu/:menuId','获取菜单信息','menu','GET'),(43,'2021-12-09 19:24:29','2021-12-09 19:24:29',NULL,'/system/menu','添加菜单信息','menu','POST'),(44,'2021-12-09 19:24:48','2021-12-09 19:24:48',NULL,'/system/menu','修改菜单信息','menu','PUT'),(45,'2021-12-09 19:25:10','2021-12-09 19:25:10',NULL,'/system/menu/:menuId','删除菜单信息','menu','DELETE'),(46,'2021-12-09 19:25:44','2021-12-09 19:27:06',NULL,'/system/post/list','获取岗位分页列表','post','GET'),(47,'2021-12-09 19:26:55','2021-12-09 19:26:55',NULL,'/system/post/:postId','获取岗位信息','post','GET'),(48,'2021-12-09 19:25:44','2021-12-09 19:25:44',NULL,'/system/post','添加岗位信息','post','POST'),(49,'2021-12-09 19:25:44','2021-12-09 19:25:44',NULL,'/system/post','修改岗位信息','post','PUT'),(50,'2021-12-09 19:25:44','2021-12-09 19:25:44',NULL,'/system/post/:postId','删除岗位信息','post','DELETE'),(51,'2021-12-09 19:25:44','2021-12-09 19:25:44',NULL,'/system/role/list','获取角色分页列表','role','GET'),(52,'2021-12-09 19:25:44','2021-12-09 19:25:44',NULL,'/system/role/:roleId','获取角色信息','role','GET'),(53,'2021-12-09 19:25:44','2021-12-09 19:25:44',NULL,'/system/role','添加角色信息','role','POST'),(54,'2021-12-09 19:25:44','2021-12-09 19:25:44',NULL,'/system/role','修改角色信息','role','PUT'),(55,'2021-12-09 19:25:44','2021-12-09 19:25:44',NULL,'/system/role/:roleId','删除角色信息','role','DELETE'),(56,'2021-12-09 19:25:44','2021-12-09 19:25:44',NULL,'/system/role/changeStatus','修改角色状态','role','PUT'),(57,'2021-12-09 19:25:44','2021-12-09 19:25:44',NULL,'/system/role/dataScope','修改角色部门权限','role','PUT'),(58,'2021-12-09 19:25:44','2021-12-09 19:25:44',NULL,'/system/role/export','导出角色信息','role','GET'),(59,'2021-12-09 19:50:57','2022-01-19 08:58:20',NULL,'/system/api/list','获取api分页列表1','api','GET'),(60,'2021-12-09 19:51:26','2021-12-09 19:51:26',NULL,'/system/api/all','获取所有api','api','GET'),(61,'2021-12-09 19:51:54','2021-12-09 19:51:54',NULL,'/system/api/getPolicyPathByRoleId','获取角色拥有的api权限','api','GET'),(62,'2021-12-09 19:52:14','2021-12-09 19:52:14',NULL,'/system/api/:id','获取api信息','api','GET'),(63,'2021-12-09 19:52:35','2021-12-09 19:52:35',NULL,'/system/api','添加api信息','api','POST'),(64,'2021-12-09 19:52:50','2021-12-09 19:52:50',NULL,'/system/api','修改api信息','api','PUT'),(65,'2021-12-09 19:53:07','2021-12-09 19:53:07',NULL,'/system/api/:id','删除api信息','api','DELETE'),(66,'2021-12-17 10:51:05','2021-12-17 10:54:22',NULL,'/log/logLogin/list','获取登录日志','log','GET'),(67,'2021-12-17 10:51:43','2021-12-17 10:54:28',NULL,'/log/logLogin/:infoId','删除日志','log','DELETE'),(68,'2021-12-17 10:53:09','2021-12-17 10:54:34',NULL,'/log/logLogin/all','清空所有','log','DELETE'),(69,'2021-12-17 10:54:07','2021-12-17 10:54:07',NULL,'/log/logOper/list','操作日志列表','log','GET'),(70,'2021-12-17 10:53:09','2021-12-17 10:53:09',NULL,'/log/logOper/:operId','删除','log','DELETE'),(71,'2021-12-17 10:53:09','2021-12-17 10:53:09',NULL,'/log/logOper/all','清空','log','DELETE'),(72,'2021-12-24 15:41:23','2021-12-24 15:41:23',NULL,'/job/list','任务列表','job','GET'),(73,'2021-12-24 15:41:54','2021-12-24 15:41:54',NULL,'/job','添加','job','POST'),(74,'2021-12-24 15:42:11','2021-12-24 15:42:11',NULL,'/job','修改任务','job','PUT'),(75,'2021-12-24 15:42:37','2021-12-24 16:32:21',NULL,'/job/:jobId','获取任务','job','GET'),(76,'2021-12-24 15:43:09','2021-12-24 16:32:05',NULL,'/job/:jobId','删除job','job','DELETE'),(77,'2021-12-24 15:43:35','2021-12-24 16:31:11',NULL,'/job/stop/:jobId','停止任务','job','GET'),(78,'2021-12-24 15:44:09','2021-12-24 16:30:38',NULL,'/job/start/:jobId','开始任务','job','GET'),(79,'2021-12-24 15:45:03','2023-08-08 14:15:59',NULL,'/job/log/list','任务日志列表','job','GET'),(80,'2021-12-24 15:45:33','2023-08-08 14:16:07',NULL,'/job/log/all','清空任务日志','job','DELETE'),(81,'2021-12-24 15:46:08','2023-08-08 14:16:15',NULL,'/job/log/:logId','删除任务日志','job','DELETE'),(82,'2021-12-24 15:45:33','2021-12-24 15:45:33',NULL,'/system/notice/list','获取通知分页列表','notice','GET'),(83,'2021-12-24 15:45:33','2021-12-24 15:45:33',NULL,'/system/notice','添加通知信息','notice','POST'),(84,'2021-12-24 15:45:33','2021-12-24 15:45:33',NULL,'/system/notice','修改通知信息','notice','PUT'),(85,'2021-12-24 15:45:33','2021-12-24 16:33:48',NULL,'/system/notice/:noticeId','删除通知信息','notice','DELETE'),(86,'2021-12-24 22:40:19','2021-12-24 22:40:19',NULL,'/job/changeStatus','修改状态','job','PUT'),(88,'2022-01-02 13:53:06','2022-07-18 10:57:58',NULL,'/develop/code/table/db/list','数据库表列表','gen','GET'),(89,'2022-01-02 13:53:44','2022-01-02 13:53:44',NULL,'/develop/code/table/list','表列表','gen','GET'),(90,'2022-01-02 13:54:10','2022-01-02 13:54:10',NULL,'/develop/code/table/info/:tableId','表信息','gen','GET'),(91,'2022-01-02 13:54:42','2022-07-18 10:58:35',NULL,'/develop/code/table/info/tableName','表名获取表信息','gen','GET'),(92,'2022-01-02 13:55:13','2022-01-02 13:55:13',NULL,'/develop/code/table/tableTree','表树','gen','GET'),(93,'2022-01-02 13:56:37','2022-01-02 13:56:37',NULL,'/develop/code/table','导入表','gen','POST'),(94,'2022-01-02 13:57:36','2022-01-02 13:57:36',NULL,'/develop/code/table','修改代码生成信息','gen','PUT'),(95,'2022-01-02 13:58:25','2022-01-02 13:58:25',NULL,'/develop/code/table/:tableId','删除表数据','gen','DELETE'),(96,'2022-01-02 13:59:07','2022-01-02 13:59:07',NULL,'/develop/code/gen/preview/:tableId','预览代码','gen','GET'),(97,'2022-01-02 13:59:43','2022-01-02 13:59:43',NULL,'/develop/code/gen/code/:tableId','生成代码','gen','GET'),(98,'2022-01-02 14:00:10','2022-07-17 01:19:42',NULL,'/develop/code/gen/configure/:tableId','生成api菜单','gen','GET');
/*!40000 ALTER TABLE `sys_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_configs`
--

DROP TABLE IF EXISTS `sys_configs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_configs` (
  `config_id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `config_name` varchar(128) DEFAULT NULL COMMENT 'ConfigName',
  `config_key` varchar(128) DEFAULT NULL COMMENT 'ConfigKey',
  `config_value` varchar(255) DEFAULT NULL COMMENT 'ConfigValue',
  `config_type` varchar(64) DEFAULT NULL COMMENT '是否系统内置0，1',
  `is_frontend` varchar(1) DEFAULT NULL COMMENT '是否前台',
  `remark` varchar(128) DEFAULT NULL COMMENT 'Remark',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`config_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_configs`
--

LOCK TABLES `sys_configs` WRITE;
/*!40000 ALTER TABLE `sys_configs` DISABLE KEYS */;
INSERT INTO `sys_configs` VALUES (1,'账号初始密码','sys.user.initPassword','123456','0','0','初始密码','2021-12-04 13:50:13','2021-12-04 13:54:52',NULL);
/*!40000 ALTER TABLE `sys_configs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict_data`
--

DROP TABLE IF EXISTS `sys_dict_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dict_data` (
  `dict_code` bigint NOT NULL AUTO_INCREMENT,
  `dict_sort` int DEFAULT NULL COMMENT '排序',
  `dict_label` varchar(64) DEFAULT NULL COMMENT '标签',
  `dict_value` varchar(64) DEFAULT NULL COMMENT '值',
  `dict_type` varchar(64) DEFAULT NULL COMMENT '字典类型',
  `status` varchar(1) DEFAULT NULL COMMENT '状态（0正常 1停用）',
  `css_class` varchar(128) DEFAULT NULL COMMENT 'CssClass',
  `list_class` varchar(128) DEFAULT NULL COMMENT 'ListClass',
  `is_default` varchar(8) DEFAULT NULL COMMENT 'IsDefault',
  `create_by` varchar(191) DEFAULT NULL,
  `update_by` varchar(191) DEFAULT NULL,
  `remark` varchar(256) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_data`
--

LOCK TABLES `sys_dict_data` WRITE;
/*!40000 ALTER TABLE `sys_dict_data` DISABLE KEYS */;
INSERT INTO `sys_dict_data` VALUES (1,0,'男','0','sys_user_sex','0','','','','admin','','男','2021-11-30 14:58:18','2021-11-30 14:58:18',NULL),(2,1,'女','1','sys_user_sex','0','','','','admin','','女生','2021-11-30 15:09:11','2021-11-30 15:10:28',NULL),(3,2,'未知','2','sys_user_sex','0','','','','admin','','未知','2021-11-30 15:09:11','2021-11-30 15:10:28',NULL),(4,0,'正常','0','sys_normal_disable','0','','','','admin','','','2021-12-01 15:58:50','2021-12-01 15:58:50',NULL),(5,1,'停用','1','sys_normal_disable','0','','','','admin','','','2021-12-01 15:59:08','2021-12-01 15:59:08',NULL),(6,0,'目录','M','sys_menu_type','0','','','','admin','','','2021-12-02 09:49:12','2021-12-02 09:49:12',NULL),(7,1,'菜单','C','sys_menu_type','0','','','','admin','','','2021-12-02 09:49:35','2021-12-02 09:49:52',NULL),(8,2,'按钮','F','sys_menu_type','0','','','','admin','','','2021-12-02 09:49:35','2021-12-02 09:49:35',NULL),(9,0,'显示','0','sys_show_hide','0','','','','admin','','','2021-12-02 09:56:40','2021-12-02 09:56:40',NULL),(10,0,'隐藏','1','sys_show_hide','0','','','','admin','','','2021-12-02 09:56:52','2021-12-02 09:56:52',NULL),(11,0,'是','0','sys_num_yes_no','0','','','','admin','','','2021-12-02 10:16:16','2021-12-02 10:16:16',NULL),(12,1,'否','1','sys_num_yes_no','0','','','','admin','','','2021-12-02 10:16:26','2021-12-02 10:16:26',NULL),(13,0,'是','0','sys_yes_no','0','','','','admin','','','2021-12-04 13:48:15','2021-12-04 13:48:15',NULL),(14,0,'否','1','sys_yes_no','0','','','','admin','','','2021-12-04 13:48:21','2021-12-04 13:48:21',NULL),(15,0,'创建(POST)','POST','sys_method_api','0','','','','admin','','','2021-12-08 17:22:05','2021-12-09 09:29:52',NULL),(16,1,'查询(GET)','GET','sys_method_api','0','','','','admin','','','2021-12-08 17:22:24','2021-12-09 09:29:59',NULL),(17,2,'修改(PUT)','PUT','sys_method_api','0','','','','admin','','','2021-12-08 17:22:40','2021-12-09 09:30:06',NULL),(18,3,'删除(DELETE)','DELETE','sys_method_api','0','','','','admin','','','2021-12-08 17:22:54','2021-12-09 09:30:13',NULL),(19,0,'成功','0','sys_common_status','0','','','','admin','','','2021-12-17 11:01:52','2021-12-17 11:01:52',NULL),(20,0,'失败','1','sys_common_status','0','','','','admin','','','2021-12-17 11:02:08','2021-12-17 11:02:08',NULL),(21,0,'其他','0','sys_oper_type','0','','','','admin','','','2021-12-17 11:30:07','2021-12-17 11:30:07',NULL),(22,0,'新增','1','sys_oper_type','0','','','','admin','','','2021-12-17 11:30:21','2021-12-17 11:30:21',NULL),(23,0,'修改','2','sys_oper_type','0','','','','admin','','','2021-12-17 11:30:32','2021-12-17 11:30:32',NULL),(24,0,'删除','3','sys_oper_type','0','','','','admin','','','2021-12-17 11:30:40','2021-12-17 11:30:40',NULL),(25,0,'默认','DEFAULT','sys_job_group','0','','','','admin','','','2021-12-24 15:15:31','2021-12-24 15:15:31',NULL),(26,1,'系统','SYSTEM','sys_job_group','0','','','','admin','','','2021-12-24 15:15:50','2021-12-24 15:15:50',NULL),(27,0,'发布通知','1','sys_notice_type','0','','','','admin','','','2021-12-26 15:24:07','2021-12-26 15:24:07',NULL),(28,0,'任免通知','2','sys_notice_type','0','','','','admin','','','2021-12-26 15:24:18','2021-12-26 15:24:18',NULL),(29,0,'事物通知','3','sys_notice_type','0','','','','admin','','','2021-12-26 15:24:46','2021-12-26 15:24:46',NULL),(30,0,'审批通知','4','sys_notice_type','0','','','','admin','','','2021-12-26 15:25:08','2021-12-26 15:25:08',NULL),(31,0,'阿里云','0','res_oss_category','0','','','','admin','','','2022-01-13 15:44:01','2022-01-13 15:44:01',NULL),(32,1,'七牛云','1','res_oss_category','0','','','','admin','','','2022-01-13 15:44:18','2022-01-13 15:44:18',NULL),(33,2,'腾讯云','2','res_oss_category','0','','','','admin','','','2022-01-13 15:44:39','2022-01-13 15:44:39',NULL),(34,0,'阿里云','0','res_sms_category','0','','','','admin','','','2022-01-13 15:47:30','2022-01-13 15:47:30',NULL),(35,1,'腾讯云','1','res_sms_category','0','','','','admin','','','2022-01-13 15:47:39','2022-01-13 15:47:39',NULL),(36,0,'163邮箱','0','res_mail_category','0','','','','admin','','','2022-01-14 15:43:42','2022-01-14 15:43:42',NULL),(37,0,'qq邮箱','1','res_mail_category','0','','','','admin','','','2022-01-14 15:44:08','2022-01-14 15:44:08',NULL),(38,0,'企业邮箱','2','res_mail_category','0','','','','admin','','','2022-01-14 15:44:20','2022-01-14 15:44:20',NULL),(39,1,'未发布','0','sys_release_type','0','','','','admin','','','2023-07-21 16:11:27','2023-07-21 16:11:27',NULL),(40,2,'已发布','1','sys_release_type','0','','','','admin','','','2023-07-21 16:11:44','2023-07-21 16:11:44',NULL);
/*!40000 ALTER TABLE `sys_dict_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict_types`
--

DROP TABLE IF EXISTS `sys_dict_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_dict_types` (
  `dict_id` bigint NOT NULL AUTO_INCREMENT,
  `dict_name` varchar(64) DEFAULT NULL COMMENT '名称',
  `dict_type` varchar(64) DEFAULT NULL COMMENT '类型',
  `status` varchar(1) DEFAULT NULL COMMENT '状态',
  `create_by` varchar(191) DEFAULT NULL,
  `update_by` varchar(191) DEFAULT NULL,
  `remark` varchar(256) DEFAULT NULL COMMENT '备注',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`dict_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_types`
--

LOCK TABLES `sys_dict_types` WRITE;
/*!40000 ALTER TABLE `sys_dict_types` DISABLE KEYS */;
INSERT INTO `sys_dict_types` VALUES (1,'用户性别','sys_user_sex','0','admin','','性别列表','2021-11-30 14:02:52','2021-11-30 14:07:55','2021-11-30 14:11:54'),(2,'用户性别','sys_user_sex','0','admin','','用户性别列表','2021-11-30 14:12:33','2021-11-30 14:12:33','2021-11-30 14:14:19'),(3,'的心','sfd','0','admin','','fs','2021-11-30 14:13:22','2021-11-30 14:13:22','2021-11-30 14:14:19'),(4,'用户性别','sys_user_sex','0','admin','','性别字典','2021-11-30 14:15:25','2021-11-30 14:15:25',NULL),(5,'df','da','0','admin','','sd','2021-11-30 15:54:33','2021-11-30 15:54:33','2021-11-30 15:54:40'),(6,'系统开关','sys_normal_disable','0','admin','','开关列表','2021-12-01 15:57:58','2021-12-01 15:57:58',NULL),(7,'菜单类型','sys_menu_type','0','admin','','菜单类型列表','2021-12-02 09:48:48','2021-12-02 09:56:12',NULL),(8,'菜单状态','sys_show_hide','0','admin','','菜单状态列表','2021-12-02 09:55:59','2021-12-02 09:55:59',NULL),(9,'数字是否','sys_num_yes_no','0','admin','','数字是否列表','2021-12-02 10:13:29','2021-12-02 10:13:40','2021-12-02 10:15:07'),(10,'数字是否','sys_num_yes_no','0','admin','','数字是否','2021-12-02 10:13:29','2021-12-02 10:13:29',NULL),(11,'状态是否','sys_yes_no','0','admin','','状态是否','2021-12-04 13:47:57','2021-12-04 13:47:57',NULL),(12,'网络请求方法','sys_method_api','0','admin','','网络请求方法列表','2021-12-08 17:21:27','2021-12-08 17:21:27',NULL),(13,'成功失败','sys_common_status','0','admin','','是否成功失败','2021-12-17 10:10:03','2021-12-17 10:10:03',NULL),(27,'操作分类','sys_oper_type','0','admin','','操作分类列表','2021-12-17 11:29:50','2021-12-17 11:29:50',NULL),(28,'任务组','sys_job_group','0','admin','','系统任务，开机自启','2021-12-24 15:14:56','2021-12-24 15:14:56',NULL),(29,'通知类型','sys_notice_type','0','admin','','通知类型列表','2021-12-26 15:23:26','2021-12-26 15:23:26',NULL),(30,'oss分类','res_oss_category','0','admin','','oss分类列表','2022-01-13 15:43:29','2022-01-13 15:43:29',NULL),(31,'sms分类','res_sms_category','0','admin','','sms分类列表','2021-12-26 15:23:26','2022-01-13 15:47:13',NULL),(32,'mail分类','res_mail_category','0','admin','','mail分类列表','2022-01-14 15:43:17','2022-01-14 15:43:17',NULL),(33,'发布状态','sys_release_type','0','admin','','发布状态','2023-07-21 16:10:38','2023-07-21 16:10:38',NULL);
/*!40000 ALTER TABLE `sys_dict_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_jobs`
--

DROP TABLE IF EXISTS `sys_jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_jobs` (
  `id` varchar(191) NOT NULL,
  `owner` varchar(64) DEFAULT NULL COMMENT '创建者,所有者',
  `org_id` bigint DEFAULT NULL COMMENT '机构ID',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `job_name` varchar(255) DEFAULT NULL,
  `target_invoke` varchar(64) DEFAULT NULL COMMENT '目标类型',
  `target_args` varchar(64) DEFAULT NULL COMMENT '目标Id',
  `job_content` json DEFAULT NULL COMMENT '任务内容',
  `cron_expression` varchar(255) DEFAULT NULL,
  `misfire_policy` varchar(1) DEFAULT NULL,
  `status` varchar(1) DEFAULT NULL,
  `entry_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_jobs`
--

LOCK TABLES `sys_jobs` WRITE;
/*!40000 ALTER TABLE `sys_jobs` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_jobs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_menus`
--

DROP TABLE IF EXISTS `sys_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_menus` (
  `menu_id` bigint NOT NULL AUTO_INCREMENT,
  `menu_name` varchar(128) DEFAULT NULL,
  `title` varchar(64) DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  `sort` int DEFAULT NULL,
  `icon` varchar(128) DEFAULT NULL,
  `path` varchar(128) DEFAULT NULL,
  `component` varchar(255) DEFAULT NULL,
  `is_iframe` varchar(1) DEFAULT NULL,
  `is_link` varchar(255) DEFAULT NULL,
  `menu_type` varchar(1) DEFAULT NULL,
  `is_hide` varchar(1) DEFAULT NULL,
  `is_keep_alive` varchar(1) DEFAULT NULL,
  `is_affix` varchar(1) DEFAULT NULL,
  `permission` varchar(32) DEFAULT NULL,
  `status` varchar(191) DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(191) DEFAULT NULL,
  `menu_group` char(1) DEFAULT 0 NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=161 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_menus`
--

LOCK TABLES `sys_menus` WRITE;
/*!40000 ALTER TABLE `sys_menus` DISABLE KEYS */;
INSERT INTO `sys_menus` VALUES (1,'系统设置','',0,0,'elementSetting','/system','Layout','1','','M','0','0','1','','0','admin','','',1,'2021-12-02 11:04:08','2021-12-28 13:32:21',NULL),((3,'用户管理','',1,1,'elementUser','/system/user','/system/user/index','1','','C','0','1','1','system:user:list','0','admin','','',1,'2021-12-02 14:07:56','2021-12-28 13:32:44',NULL),((4,'添加用户','',3,1,'','','','','','F','0','','','system:user:add','0','admin','','',1,'2021-12-03 13:36:33','2021-12-03 13:36:33',NULL),((5,'编辑用户','',3,1,'','','','','','F','0','','','system:user:edit','0','admin','','',1,'2021-12-03 13:48:13','2021-12-03 13:48:13',NULL),((6,'角色管理','',1,2,'elementUserFilled','/system/role','/system/role/index','1','','C','0','1','1','system:role:list','0','admin','','',1,'2021-12-03 13:51:55','2022-07-16 10:23:21',NULL),((7,'菜单管理','',1,2,'iconfont icon-juxingkaobei','/system/menu','/system/menu/index','1','','C','0','1','1','system:menu:list','0','admin','','',1,'2021-12-03 13:54:44','2021-12-28 13:33:19',NULL),((8,'组织管理','',1,3,'iconfont icon-jiliandongxuanzeqi','/system/organization','/system/organization/index','1','','C','0','1','1','system:organization:list','0','admin','','',1,'2021-12-03 13:58:36','2023-09-14 14:05:07',NULL),((9,'岗位管理','',1,4,'iconfont icon-neiqianshujuchucun','/system/post','/system/post/index','1','','C','0','1','1','system:post:list','0','admin','','',1,'2021-12-03 13:54:44','2021-12-28 13:40:31',NULL),((10,'字典管理','',1,5,'elementCellphone','/system/dict','/system/dict/index','1','','C','0','1','1','system:dict:list','0','admin','','',1,'2021-12-03 13:54:44','2021-12-28 13:40:50',NULL),((11,'参数管理','',1,6,'elementDocumentCopy','/system/config','/system/config/index','1','','C','0','1','1','system:config:list','0','admin','','',1,'2021-12-03 13:54:44','2021-12-28 13:41:05',NULL),((12,'个人中心','',0,10,'elementAvatar','/personal','/personal/index','1','','M','1','0','0','','0','admin','','',1,'2021-12-03 14:12:43','2023-06-27 10:09:26',NULL),((13,'添加配置','',11,1,'','','','','','F','','','','system:config:add','0','admin','','',1,'2021-12-06 17:19:19','2021-12-06 17:19:19',NULL),((14,'修改配置','',11,1,'','','','','','F','','','','system:config:edit','0','admin','','',1,'2021-12-06 17:20:30','2021-12-06 17:20:30',NULL),((15,'删除配置','',11,1,'','','','','','F','','','','system:config:delete','0','admin','','',1,'2021-12-06 17:23:52','2021-12-06 17:23:52',NULL),((16,'导出配置','',11,1,'','','','','','F','','','','system:config:export','0','admin','','',1,'2021-12-06 17:24:41','2021-12-06 17:24:41',NULL),((17,'新增角色','',6,1,'','','','','','F','','','','system:role:add','0','admin','','',1,'2021-12-06 17:43:35','2021-12-06 17:43:35',NULL),((18,'删除角色','',6,1,'','','','','','F','','','','system:role:delete','0','admin','','',1,'2021-12-06 17:44:10','2021-12-06 17:44:10',NULL),((19,'修改角色','',6,1,'','','','','','F','','','','system:role:edit','0','admin','','',1,'2021-12-06 17:44:48','2021-12-06 17:44:48',NULL),((20,'导出角色','',6,1,'','','','','','F','','','','system:role:export','0','admin','','',1,'2021-12-06 17:45:25','2021-12-06 17:45:25',NULL),((21,'添加菜单','',7,1,'','','','','','F','','','','system:menu:add','0','admin','','',1,'2021-12-06 17:46:01','2021-12-06 17:46:01',NULL),((22,'修改菜单','',7,1,'','','','','','F','','','','system:menu:edit','0','admin','','',1,'2021-12-06 17:46:24','2021-12-06 17:46:24',NULL),((23,'删除菜单','',7,1,'','','','','','F','','','','system:menu:delete','0','admin','','',1,'2021-12-06 17:46:47','2021-12-06 17:46:47',NULL),((24,'添加部门','',8,1,'','','','','','F','','','','system:organization:add','0','admin','','',1,'2021-12-07 09:33:58','2023-09-14 14:05:20',NULL),((25,'编辑部门','',8,1,'','','','','','F','','','','system:organization:edit','0','admin','','',1,'2021-12-07 09:34:39','2023-09-14 14:05:26',NULL),((26,'删除部门','',8,1,'','','','','','F','','','','system:organization:delete','0','admin','','',1,'2021-12-07 09:35:09','2023-09-14 14:05:32',NULL),((28,'添加岗位','',9,1,'','','','','','F','','','','system:post:add','0','admin','','',1,'2021-12-07 09:35:09','2021-12-07 09:35:09',NULL),((29,'编辑岗位','',9,1,'','','','','','F','','','','system:post:edit','0','admin','','',1,'2021-12-07 09:35:09','2021-12-07 09:35:09',NULL),((30,'删除岗位','',9,1,'','','','','','F','','','','system:post:delete','0','admin','','',1,'2021-12-07 09:35:09','2021-12-07 09:35:09',NULL),((31,'导出岗位','',9,1,'','','','','','F','','','','system:post:export','0','admin','','',1,'2021-12-07 09:35:09','2021-12-07 09:35:09',NULL),((32,'添加字典类型','',10,1,'','','','','','F','','','','system:dictT:add','0','admin','','',1,'2021-12-07 09:35:09','2021-12-07 09:35:09',NULL),((33,'编辑字典类型','',10,1,'','','','','','F','','','','system:dictT:edit','0','admin','','',1,'2021-12-07 09:35:09','2021-12-07 09:35:09',NULL),((34,'删除字典类型','',10,1,'','','','','','F','','','','system:dictT:delete','0','admin','','',1,'2021-12-07 09:35:09','2021-12-07 09:35:09',NULL),((35,'导出字典类型','',10,1,'','','','','','F','','','','system:dictT:export','0','admin','','',1,'2021-12-07 09:35:09','2021-12-07 09:35:09',NULL),((36,'新增字典数据','',10,1,'','','','','','F','','','','system:dictD:add','0','admin','','',1,'2021-12-07 09:35:09','2021-12-07 09:35:09',NULL),((37,'修改字典数据','',10,1,'','','','','','F','','','','system:dictD:edit','0','admin','','',1,'2021-12-07 09:48:04','2021-12-07 09:48:04',NULL),((38,'删除字典数据','',10,1,'','','','','','F','','','','system:dictD:delete','0','admin','','',1,'2021-12-07 09:48:42','2021-12-07 09:48:42',NULL),((39,'API管理','',1,2,'iconfont icon-siweidaotu','/system/api','/system/api/index','1','','C','0','1','1','system:api:list','0','admin','','',1,'2021-12-09 09:09:13','2022-07-16 10:23:42',NULL),((40,'添加api','',39,1,'','/system/api','','','','F','','','','system:api:add','0','admin','','',1,'2021-12-09 09:09:54','2021-12-09 09:09:54',NULL),((41,'编辑api','',39,1,'','/system/api','','','','F','','','','system:api:edit','0','admin','','',1,'2021-12-09 09:10:38','2021-12-09 09:10:38',NULL),((42,'删除api','',39,1,'','/system/api','','','','F','','','','system:api:delete','0','admin','','',1,'2021-12-09 09:11:11','2021-12-09 09:11:11',NULL),((43,'日志系统','',0,11,'iconfont icon-biaodan','/log','Layout','1','','M','0','1','1','','0','admin','','',1,'2021-12-02 11:04:08','2023-06-30 08:57:08',NULL),((44,'告警监控','',0,9,'iconfont icon-gongju','/tool','Layout','1','','M','0','1','1','','0','admin','','',1,'2021-12-16 16:35:15','2023-10-18 10:17:52',NULL),((45,'操作日志','',43,1,'iconfont icon-bolangnengshiyanchang','/log/operation','/log/operation/index','1','','C','0','1','1','log:operation:list','0','admin','','',1,'2021-12-16 16:42:03','2021-12-28 13:39:44',NULL),((46,'登录日志','',43,2,'iconfont icon--chaifenlie','/log/login','/log/login/index','1','','C','0','1','1','log:login:list','0','admin','','',1,'2021-12-16 16:43:28','2021-12-28 13:39:58',NULL),((47,'服务监控','',44,1,'elementCpu','/tool/monitor/','/tool/monitor/index','1','','C','0','1','1','tool:monitor:list','0','admin','','',1,'2021-12-03 14:12:43','2021-12-28 13:41:25',NULL),((49,'开发工具','',0,10,'iconfont icon-zhongduancanshu','/develop','Layout','1','','M','0','1','1','','0','admin','','',1,'2021-12-16 16:53:11','2023-06-29 16:29:23',NULL),((50,'表单构建','',49,1,'iconfont icon-zidingyibuju','/develop/form','/develop/form/index','1','','C','0','1','1','develop:form:list','0','admin','','',1,'2021-12-16 16:55:01','2022-07-12 18:56:18',NULL),((51,'代码生成','',49,2,'iconfont icon-zhongduancanshu','/develop/code','/develop/code/index','1','','C','0','1','1','develop:code:list','0','admin','','',1,'2021-12-16 16:56:48','2021-12-16 16:56:48',NULL),((52,'系统接口','',49,3,'iconfont icon-wenducanshu-05','/develop/apis','/layout/routerView/iframes','0','http://127.0.0.1:39999/swagger/doc.html','C','0','1','1','develop:apis:list','0','admin','','',1,'2021-12-16 16:58:07','2023-09-04 11:02:29',NULL),((54,'对象存储','',53,1,'iconfont icon-chazhaobiaodanliebiao','/resource/file','/resource/file/index','1','','C','0','1','1','resource:file:list','0','admin','','',1,'2021-12-16 17:06:04','2022-01-13 17:30:09',NULL),((55,'公告通知','',44,3,'elementTicket','/tool/notice','/tool/notice/index','1','','C','0','1','1','tool:notice:list','0','admin','','',1,'2021-12-16 22:09:11','2021-12-28 13:42:39',NULL),((59,'删除','',45,1,'','','','','','F','','','','log:operation:delete','0','admin','','',1,'2022-01-14 13:28:25','2022-01-14 13:28:25',NULL),((60,'清空','',45,1,'','','','','','F','','','','log:operation:clean','0','admin','','',1,'2022-01-14 13:29:24','2022-01-14 13:29:24',NULL),((63,'删除','',46,1,'','','','','','F','','','','log:login:delete','0','admin','','',1,'2022-01-14 13:30:46','2022-01-14 13:30:46',NULL),((64,'清空','',46,1,'','','','','','F','','','','log:login:clean','0','admin','','',1,'2022-01-14 13:31:06','2022-01-14 13:31:06',NULL),((69,'添加','',55,1,'','','','','','F','','','','tool:notice:add','0','admin','','',1,'2022-01-14 13:35:23','2022-01-14 13:35:23',NULL),((70,'编辑','',55,1,'','','','','','F','','','','tool:notice:edit','0','admin','','',1,'2022-01-14 13:36:04','2022-01-14 13:36:04',NULL),((71,'删除','',55,1,'','','','','','F','','','','tool:notice:delete','0','admin','','',1,'2022-01-14 13:36:26','2022-01-14 13:36:26',NULL),((72,'查看','',55,1,'','','','','','F','','','','tool:notice:view','0','admin','','',1,'2022-01-14 13:36:51','2022-01-14 13:36:51',NULL),((73,'导入','',51,1,'','','','','','F','','','','develop:code:add','0','admin','','',1,'2022-01-14 13:38:35','2022-01-14 13:38:35',NULL),((74,'编辑','',51,1,'','','','','','F','','','','develop:code:edit','0','admin','','',1,'2022-01-14 13:41:25','2022-01-14 13:41:25',NULL),((75,'删除','',51,1,'','','','','','F','','','','develop:code:delete','0','admin','','',1,'2022-01-14 13:41:42','2022-01-14 13:41:42',NULL),((76,'预览','',51,1,'','','','','','F','','','','develop:code:view','0','admin','','',1,'2022-01-14 13:42:01','2022-01-14 13:42:01',NULL),((77,'生成代码','',51,1,'','','','','','F','','','','develop:code:code','0','admin','','',1,'2022-01-14 13:42:48','2022-01-14 13:42:48',NULL),((152,'任务中心','',0,5,'iconfont icon-dayin','/job','Layout','1','','M','0','1','1','','0','admin','','',1,'2023-08-08 14:08:11','2023-10-05 14:03:38',NULL),((153,'任务中心','',152,1,'elementAlarmClock','/job/job','/job/job/index','1','','C','0','1','1','job:list','0','admin','','',1,'2023-08-08 14:10:37','2023-08-08 14:12:49',NULL),((154,'任务日志','',152,2,'elementDocument','/job/log','/job/log/index','1','','C','0','1','1','job:log:list','0','admin','','',1,'2023-08-08 14:12:37','2023-08-08 14:12:37',NULL),((155,'新增','',153,1,'','','','','','F','','','','job:add','0','admin','','',1,'2023-08-08 14:20:17','2023-08-08 14:20:17',NULL),((156,'编辑','',153,2,'','','','','','F','','','','job:edit','0','admin','','',1,'2023-08-08 14:20:44','2023-08-08 14:20:44',NULL),((157,'删除','',153,3,'','','','','','F','','','','job:delete','0','admin','','',1,'2023-08-08 14:21:03','2023-08-08 14:21:03',NULL),((158,'运行启动','',153,4,'','','','','','F','','','','job:run','0','admin','','',1,'2023-08-08 14:21:25','2023-08-08 14:21:25',NULL),((159,'删除','',154,1,'','','','','','F','','','','job:log:delete','0','admin','','',1,'2023-08-08 14:22:05','2023-08-08 14:22:05',NULL),((160,'清空','',154,2,'','','','','','F','','','','job:log:clean','0','admin','','',1,'2023-08-08 14:22:33','2023-08-08 14:22:33',NULL);
/*!40000 ALTER TABLE `sys_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_notices`
--

DROP TABLE IF EXISTS `sys_notices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_notices` (
  `notice_id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(128) DEFAULT NULL COMMENT '标题',
  `content` text COMMENT '标题',
  `notice_type` varchar(1) DEFAULT NULL COMMENT '通知类型',
  `organization_id` int DEFAULT NULL COMMENT '部门Id,部门及子部门',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  `user_name` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`notice_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_notices`
--

LOCK TABLES `sys_notices` WRITE;
/*!40000 ALTER TABLE `sys_notices` DISABLE KEYS */;
INSERT INTO `sys_notices` VALUES (1,'关于学习交流的通知','<p>发布<b>入群</b>通知&nbsp;<span style=\"color: var(--el-text-color-regular);\">467890197, 交流学习</span><span style=\"font-size: 14px; color: var(--el-text-color-regular);\"></span></p>','1',0,'2021-12-26 15:29:25','2021-12-26 16:19:48',NULL,'admin'),(2,'test','<p>sdsad</p>','1',2,'2021-12-26 16:23:13','2021-12-26 16:23:13','2021-12-26 16:31:31','admin'),(3,'版本更新通知：任务功能，通知功能完成','<p><span style=\"font-size: 14px; color: var(--el-text-color-regular);\">','1',0,'2021-12-26 17:33:47','2021-12-26 17:33:47',NULL,'admin');
/*!40000 ALTER TABLE `sys_notices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_organizations`
--

DROP TABLE IF EXISTS `sys_organizations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_organizations` (
  `organization_id` bigint NOT NULL AUTO_INCREMENT,
  `parent_id` int DEFAULT NULL COMMENT '上级组织',
  `organization_path` varchar(255) DEFAULT NULL COMMENT '组织路径',
  `organization_name` varchar(128) DEFAULT NULL COMMENT '组织名称',
  `sort` int DEFAULT NULL COMMENT '排序',
  `leader` varchar(64) DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) DEFAULT NULL COMMENT '手机',
  `email` varchar(64) DEFAULT NULL COMMENT '邮箱',
  `status` char(1) DEFAULT NULL COMMENT '状态',
  `create_by` varchar(64) DEFAULT NULL COMMENT '创建人',
  `update_by` varchar(64) DEFAULT NULL COMMENT '修改人',
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`organization_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_organizations`
--

LOCK TABLES `sys_organizations` WRITE;
/*!40000 ALTER TABLE `sys_organizations` DISABLE KEYS */;
INSERT INTO `sys_organizations` VALUES (1,0,'/0/1','我的社区',0,'admin','13966666666','13966666666@qq.com','0','admin','admin','2021-12-01 17:31:53','2021-12-02 08:56:19',NULL),(2,1,'/0/1/2','研发部',1,'admin','13900000000','13900000000@qq.com','0','admin','admin','2021-12-01 17:37:43','2021-12-02 08:55:56','2023-11-19 21:26:36'),(3,1,'/0/1/3','营销部',2,'admin','13911111111','13911111111@qq.com','0','admin','admin','2021-12-24 10:46:24','2021-12-24 10:47:15','2023-11-19 21:26:30'),(8,1,'/0/1/8','保安部',2,'','','','0','admin','admin','2023-11-19 21:27:25','2023-11-19 21:35:52',NULL),(9,1,'/0/1/9','维修部',3,'','','','0','admin','admin','2023-11-19 21:27:55','2023-11-19 21:35:58',NULL),(10,1,'/0/1/10','账务部',6,'','','','0','admin','admin','2023-11-19 21:28:07','2023-11-19 21:36:28',NULL),(11,1,'/0/1/11','办公室',5,'','','','0','admin','admin','2023-11-19 21:28:32','2023-11-19 21:36:21',NULL),(12,1,'/0/1/12','环境管理部',4,'','','','0','admin','admin','2023-11-19 21:28:56','2023-11-19 21:36:16',NULL),(13,1,'/0/1/13','综合管理部',1,'','','','0','admin','admin','2023-11-19 21:29:06','2023-11-19 21:35:45',NULL),(14,1,'/0/1/14','其他部门',7,'','','','0','admin','','2023-11-19 21:29:28','2023-11-19 21:29:28',NULL);
/*!40000 ALTER TABLE `sys_organizations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_posts`
--

DROP TABLE IF EXISTS `sys_posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_posts` (
  `post_id` bigint NOT NULL AUTO_INCREMENT,
  `post_name` varchar(128) DEFAULT NULL COMMENT '岗位名称',
  `post_code` varchar(128) DEFAULT NULL COMMENT '岗位代码',
  `sort` int DEFAULT NULL COMMENT '岗位排序',
  `status` varchar(1) DEFAULT NULL COMMENT '状态',
  `remark` varchar(255) DEFAULT NULL COMMENT '描述',
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_posts`
--

LOCK TABLES `sys_posts` WRITE;
/*!40000 ALTER TABLE `sys_posts` DISABLE KEYS */;
INSERT INTO `sys_posts` VALUES (1,'区域负责人','AREA',1,'0','区域负责人','admin','','2021-12-02 09:21:44','2023-11-19 21:30:28',NULL),(2,'项目负责人','PROJECT',2,'0','项目负责人','admin','','2021-12-02 09:21:44','2023-11-19 21:30:47',NULL),(5,'部门负责人','DEPARTMENT',3,'0','部门负责人','admin','','2023-11-19 21:31:59','2023-11-19 21:33:10',NULL),(6,'管理员','ADMIN',4,'0','管理员','admin','','2023-11-19 21:33:04','2023-11-19 21:33:14',NULL),(7,'普通员工','STAFF',5,'0','普通员工','admin','','2023-11-19 21:33:47','2023-11-19 21:33:47',NULL);
/*!40000 ALTER TABLE `sys_posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_menus`
--

DROP TABLE IF EXISTS `sys_role_menus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_role_menus` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `role_id` int DEFAULT NULL,
  `menu_id` int DEFAULT NULL,
  `role_name` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=128 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_menus`
--

LOCK TABLES `sys_role_menus` WRITE;
/*!40000 ALTER TABLE `sys_role_menus` DISABLE KEYS */;
INSERT INTO `sys_role_menus` VALUES (1,1,1,'admin'),(2,1,3,'admin'),(3,1,4,'admin'),(4,1,5,'admin'),(5,1,6,'admin'),(6,1,7,'admin'),(7,1,8,'admin'),(8,1,9,'admin'),(9,1,10,'admin'),(10,1,11,'admin'),(11,1,12,'admin'),(12,1,13,'admin'),(13,1,14,'admin'),(14,1,15,'admin'),(15,1,16,'admin'),(16,1,17,'admin'),(17,1,18,'admin'),(18,1,19,'admin'),(19,1,20,'admin'),(20,1,21,'admin'),(21,1,22,'admin'),(22,1,23,'admin'),(23,1,24,'admin'),(24,1,25,'admin'),(25,1,26,'admin'),(26,1,28,'admin'),(27,1,29,'admin'),(28,1,30,'admin'),(29,1,31,'admin'),(30,1,32,'admin'),(31,1,33,'admin'),(32,1,34,'admin'),(33,1,35,'admin'),(34,1,36,'admin'),(35,1,37,'admin'),(36,1,38,'admin'),(37,1,39,'admin'),(38,1,40,'admin'),(39,1,41,'admin'),(40,1,42,'admin'),(41,1,43,'admin'),(42,1,44,'admin'),(43,1,45,'admin'),(44,1,46,'admin'),(45,1,47,'admin'),(46,1,49,'admin'),(47,1,50,'admin'),(48,1,51,'admin'),(49,1,52,'admin'),(50,1,54,'admin'),(51,1,55,'admin'),(52,1,59,'admin'),(53,1,60,'admin'),(54,1,63,'admin'),(55,1,64,'admin'),(56,1,69,'admin'),(57,1,70,'admin'),(58,1,71,'admin'),(59,1,72,'admin'),(60,1,73,'admin'),(61,1,74,'admin'),(62,1,75,'admin'),(63,1,76,'admin'),(64,1,77,'admin'),(65,1,152,'admin'),(66,1,153,'admin'),(67,1,154,'admin'),(68,1,155,'admin'),(69,1,156,'admin'),(70,1,157,'admin'),(71,1,158,'admin'),(72,1,159,'admin'),(73,1,160,'admin');
/*!40000 ALTER TABLE `sys_role_menus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_organizations`
--

DROP TABLE IF EXISTS `sys_role_organizations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_role_organizations` (
  `role_id` int DEFAULT NULL,
  `organization_id` int DEFAULT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_organizations`
--

LOCK TABLES `sys_role_organizations` WRITE;
/*!40000 ALTER TABLE `sys_role_organizations` DISABLE KEYS */;
INSERT INTO `sys_role_organizations` VALUES (1,2,9),(1,3,10),(1,7,11),(2,2,12),(2,3,13),(2,7,14);
/*!40000 ALTER TABLE `sys_role_organizations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_roles`
--

DROP TABLE IF EXISTS `sys_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_roles` (
  `role_id` bigint NOT NULL AUTO_INCREMENT,
  `role_name` varchar(128) DEFAULT NULL COMMENT '角色名称',
  `status` varchar(1) DEFAULT NULL COMMENT '状态',
  `role_key` varchar(128) DEFAULT NULL COMMENT '角色代码',
  `data_scope` varchar(1) DEFAULT NULL COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `role_sort` int DEFAULT NULL COMMENT '角色排序',
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  `org` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_roles`
--

LOCK TABLES `sys_roles` WRITE;
/*!40000 ALTER TABLE `sys_roles` DISABLE KEYS */;
INSERT INTO `sys_roles` VALUES (1,'超管理员','0','admin','4',1,'admin','admin','超级管理','2021-12-02 16:03:26','2023-10-18 10:23:08',NULL,NULL);
/*!40000 ALTER TABLE `sys_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_tenants`
--

DROP TABLE IF EXISTS `sys_tenants`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_tenants` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  `tenant_name` varchar(255) DEFAULT NULL COMMENT '租户名',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `expire_time` datetime DEFAULT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_tenants`
--

LOCK TABLES `sys_tenants` WRITE;
/*!40000 ALTER TABLE `sys_tenants` DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_tenants` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_users`
--

DROP TABLE IF EXISTS `sys_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sys_users` (
  `user_id` bigint NOT NULL AUTO_INCREMENT,
  `nick_name` varchar(128) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `role_id` int DEFAULT NULL,
  `salt` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `sex` varchar(255) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `organization_id` int DEFAULT NULL,
  `post_id` int DEFAULT NULL,
  `create_by` varchar(128) DEFAULT NULL,
  `update_by` varchar(128) DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  `status` varchar(1) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `delete_time` datetime DEFAULT NULL,
  `username` varchar(64) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  `role_ids` varchar(255) DEFAULT NULL COMMENT '多角色',
  `post_ids` varchar(255) DEFAULT NULL COMMENT '多岗位',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_users`
--

LOCK TABLES `sys_users` WRITE;
/*!40000 ALTER TABLE `sys_users` DISABLE KEYS */;
INSERT INTO `sys_users` VALUES (1,'admin','13888888888',1,NULL,'','0','aurinerd@aurine.cn',1,1,'admin','admin',NULL,'0','2021-12-03 09:46:55','2022-02-09 13:28:49',NULL,'admin','$2a$10$cKFFTCzGOvaIHHJY2K45Zuwt8TD6oPzYi4s5MzYIBAWCLL6ZhouP2','1','1');
/*!40000 ALTER TABLE `sys_users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-11-20 12:53:53
