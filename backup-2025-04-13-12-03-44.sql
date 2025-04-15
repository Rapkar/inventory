-- MySQL dump 10.13  Distrib 8.0.41, for Linux (x86_64)
--
-- Host: localhost    Database: Inventory
-- ------------------------------------------------------
-- Server version	8.0.41-0ubuntu0.24.04.1

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
-- Table structure for table `export_products`
--

DROP TABLE IF EXISTS `export_products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `export_products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `export_id` bigint unsigned DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `number` varchar(255) DEFAULT NULL,
  `role_price` bigint DEFAULT NULL,
  `meter_price` bigint DEFAULT NULL,
  `count` bigint DEFAULT NULL,
  `meter` bigint DEFAULT NULL,
  `total_price` bigint DEFAULT NULL,
  `inventory_number` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_export_products_export_id` (`export_id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `export_products`
--

LOCK TABLES `export_products` WRITE;
/*!40000 ALTER TABLE `export_products` DISABLE KEYS */;
INSERT INTO `export_products` VALUES (1,1,'ایزوگام شرق','10',99250,102500,100,10,2000000,1),(2,2,'ایزوگام شرق','10',99250,102500,100,10,2000000,1),(3,5,'ایزوگام شرق','',99250,102500,10,0,2017500,1),(4,6,'ایزوگام شرق','',99250,102500,10,0,1095000,1),(5,7,'ایزوگام شرق','',99250,102500,10,0,1095000,1),(6,8,'ایزوگام شرق','',99250,102500,2,0,301000,1),(7,9,'ایزوگام شرق','',99250,102500,2,0,301000,1),(8,10,'ایزوگام شرق','',99250,102500,1,0,201750,1),(9,11,'ایزوگام شرق','',99250,102500,1,0,201750,1),(10,12,'ایزوگام شرق','',99250,102500,1,0,201750,1),(11,13,'ایزوگام شرق','',99250,102500,1,0,201750,1),(12,14,'ایزوگام شرق','',99250,102500,1,0,201750,1),(13,15,'ایزوگام شرق','',99250,102500,1,0,201750,1),(14,16,'ایزوگام شرق','',99250,102500,1,0,201750,1),(15,17,'ایزوگام شرق','',99250,102500,2,0,301000,1),(16,18,'ایزوگام شرق','',99250,102500,1,0,201750,1),(17,19,'ایزوگام شرق','',99250,102500,2,0,301000,1),(18,21,'ایزوگام شرق','',99250,102500,59,0,5855750,1),(19,22,'ایزوگام شرق','',99250,102500,1,0,201750,1),(20,23,'ایزوگام شرق','',99250,102500,99,0,9928250,1),(21,24,'ایزوگام شرق','',99250,102500,1,0,201750,1),(22,25,'ایزوگام شرق','',99250,102500,1,0,99250,1),(23,26,'ایزوگام شرق','',99250,102500,1,0,99250,1),(24,27,'ایزوگام شرق','',99250,102500,1,0,99250,1),(25,28,'ایزوگام شرق','',99250,102500,1,0,99250,1),(26,29,'ایزوگام شرق','',99250,102500,1,0,99250,1),(27,30,'ایزوگام شرق','',99255,102705,5,0,1009800,1),(28,31,'ایزوگام شرق','',99255,102705,90,0,9446475,1),(29,32,'ایزوگام شرق','',99255,102705,10,0,1197960,1),(30,33,'ایزوگام شرق','',99255,102705,10,0,1300665,1),(31,33,'یک ','',1000,2000,20,0,20000,1),(32,34,'یک ','',1000,2000,2,0,2000,1),(33,34,'کالای تستی','',100,1000,4,0,400,1),(34,35,'ایزوگام شرق','',99255,102705,10,0,1197960,1),(35,36,'ایزوگام شرق','',99255,102705,1,0,99255,1),(36,37,'ایزوگام شرق','',99255,102705,10,0,992550,1),(37,37,'یک ','',1000,2000,1,0,1000,1);
/*!40000 ALTER TABLE `export_products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `exports`
--

DROP TABLE IF EXISTS `exports`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `exports` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `number` longtext,
  `phonenumber` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `total_price` bigint DEFAULT NULL,
  `tax` bigint DEFAULT NULL,
  `describe` varchar(255) DEFAULT NULL,
  `created_at` longtext,
  `inventory_number` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exports`
--

LOCK TABLES `exports` WRITE;
/*!40000 ALTER TABLE `exports` DISABLE KEYS */;
INSERT INTO `exports` VALUES (2,'رضا توانگر','9283422','09199656725','کرج -کرج=-ایران -سیسی',10000000,10,'','1404/01/21 چهارشنبه 05:33:38 ب.ظ',1),(6,'ali','I02583','02134249382','sdfsdf',1095000,200,'','1404/01/21 چهارشنبه 05:45:59 ب.ظ',1),(7,'ali','M00258','02134249383','sdfsd',1095000,2000,'gjgh','1404/01/21 چهارشنبه 05:58:09 ب.ظ',1),(8,'we','B00330','02134249381','sdfsdsd',301000,1,'','1404/01/21 چهارشنبه 06:17:58 ب.ظ',1),(9,'sef','X06967','02134249383','sefse',301000,1,'','1404/01/21 چهارشنبه 06:18:57 ب.ظ',1),(10,'ali','K01889','02134249383','qwdqwd',201750,1,'dfgfg','1404/01/21 چهارشنبه 06:28:47 ب.ظ',1),(11,'we','Z04378','02134249381','wewe',201750,1,'','1404/01/21 چهارشنبه 06:29:59 ب.ظ',1),(12,'ali','Y00074','242342','sfs',201750,1,'','1404/01/21 چهارشنبه 06:31:29 ب.ظ',1),(13,'sdv','I00850','2323','wefw',201750,10,'','1404/01/21 چهارشنبه 07:41:37 ب.ظ',1),(14,'wefwef','K02214','234','dfg',201750,0,'','1404/01/21 چهارشنبه 07:43:55 ب.ظ',1),(15,'wefwe','Y00635','2343','ergerr',201750,23,'','1404/01/21 چهارشنبه 07:50:33 ب.ظ',1),(16,'سیز','S05373','345345','erter',201750,10,'','1404/01/21 چهارشنبه 11:35:02 ب.ظ',1),(17,'فاکتور تست','Z01527','23423423','sdfas',301000,10,'wefwe','1404/01/22 پنج‌شنبه 10:25:36 ق.ظ',1),(18,'سیز','A09522','345345','sds',201750,11,'','1404/01/22 پنج‌شنبه 10:46:42 ق.ظ',1),(19,'حسین','K01082','021521','sdffw',301000,1234,'','1404/01/22 پنج‌شنبه 11:07:27 ق.ظ',1),(20,'تست','O07293','۰۲۱۳۴۲۴۹۳۸۱','سیبسی',0,1100,'','1404/01/24 شنبه 07:04:55 ب.ظ',0),(21,'تست','Z04805','۰۲۱۳۴۲۴۹۳۸۱','سیبسی',5856850,1100,'','1404/01/24 شنبه 07:06:09 ب.ظ',1),(22,'dfdft','U03134','۴۵۲۵۸۴۵۲۴۵','شسیب',201750,1200,'jkl','1404/01/24 شنبه 08:32:36 ب.ظ',1),(23,'حسین سلطانیان','P07473','۰۲۶۳۴۲۴۹۳۸۱','کرج تهران اینجا',9928250,250000,'','1404/01/25 یک‌شنبه 12:42:10 ب.ظ',1),(24,'خسته تر از همه','I04838','۰۲۱۳۴۲۴۹۳۸۵','سیبیسب',201750,1200,'س','1404/01/25 یک‌شنبه 12:55:06 ب.ظ',1),(25,'پرداخت تستی ۱','K09513','02134249381','سبی',99250,1000,'سییل','1404/01/25 یک‌شنبه 01:58:33 ب.ظ',1),(26,'sdvsv','W06680','15135','dfbdfb',99250,250,'dbdf','1404/01/25 یک‌شنبه 02:14:27 ب.ظ',1),(27,'یسبسی','M05749','۳۲۴۲۳','یسیبسل',99250,1200,'','1404/01/25 یک‌شنبه 02:34:49 ب.ظ',1),(28,'dfgfd','W09636','25000','sdcdsc',99250,1000,'ghfhjg','1404/01/25 یک‌شنبه 02:41:56 ب.ظ',1),(29,'حسین ','N04569','۰۲۱۳۴۲۴۹۳۸۱','سیزسی',99250,1200,'','1404/01/26 دوشنبه 12:34:54 ق.ظ',1),(30,'حسین امیری','Z08481','09125174845','sdf',1021800,12000,'','1404/01/26 دوشنبه 05:11:34 ب.ظ',1),(31,'امیر','L05242','09125174854','سیس',9446475,1200,'','1404/01/26 دوشنبه 05:15:32 ب.ظ',1),(32,'فاکتور مهم','V07091','09125174854','کرج حصارک شهرک یاس',1197960,100000,'','1404/01/26 دوشنبه 05:32:44 ب.ظ',1),(33,'ممد نوبری','C06957','02134249381','sdfdsf',1320665,1000,'','1404/01/26 دوشنبه 05:35:20 ب.ظ',1),(34,'فا','G02359','۲۳۴۲۳۴','۲۳ق۲۳',204360,100,'','1404/01/26 دوشنبه 06:53:28 ب.ظ',1),(35,'ممد مهرانی','C08011','۰۲۱۳۴۲۴۹۳۲۰','شسیشسی',1197960,2000,'','1404/01/26 دوشنبه 07:04:54 ب.ظ',1),(36,'ascas','X00217','025151515','نتدسیبسی',100255,2000,'','1404/01/26 دوشنبه 07:19:13 ب.ظ',1),(37,'پرداخت تستی ۱','W09623','۰۲۱۳۴۲۴۹۳۸۱','oommom',994000,225,'','1404/01/26 دوشنبه 07:33:58 ب.ظ',1);
/*!40000 ALTER TABLE `exports` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inventories`
--

DROP TABLE IF EXISTS `inventories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `inventories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `number` varchar(255) DEFAULT NULL,
  `role_price` double DEFAULT NULL,
  `meter_price` double DEFAULT NULL,
  `count` bigint DEFAULT NULL,
  `meter` bigint DEFAULT NULL,
  `inventory_number` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventories`
--

LOCK TABLES `inventories` WRITE;
/*!40000 ALTER TABLE `inventories` DISABLE KEYS */;
INSERT INTO `inventories` VALUES (1,'ایزوگام شرق','10',99255,102705,1659,7205,1),(2,'کالای تستی','',2500,1200,100,1000,1),(3,'کالای تستی','',100,1000,400,4001,1),(4,'یک ','',1000,2000,97,320,1);
/*!40000 ALTER TABLE `inventories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payments`
--

DROP TABLE IF EXISTS `payments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `method` varchar(100) DEFAULT NULL,
  `number` longtext,
  `name` varchar(100) DEFAULT NULL,
  `total_price` bigint DEFAULT NULL,
  `describe` varchar(255) DEFAULT NULL,
  `created_at` longtext,
  `export_id` bigint unsigned DEFAULT NULL,
  `status` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_exports_payments` (`export_id`),
  CONSTRAINT `fk_exports_payments` FOREIGN KEY (`export_id`) REFERENCES `exports` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payments`
--

LOCK TABLES `payments` WRITE;
/*!40000 ALTER TABLE `payments` DISABLE KEYS */;
INSERT INTO `payments` VALUES (4,'چک','232','صادرات',23,'','2025-04-01',2,'pending'),(5,'چک','4334','صادرات',333,'','2025-04-14',2,'pending'),(6,'چک','4334','صادرات',333,'','2025-04-14',2,'pending'),(7,'چک','4334','ملت',333,'','2025-02-19',2,'pending'),(8,'چک','4334','صادرات',333,'','2025-04-14',2,'rejected'),(9,'چک','4334','ملت',333,'','2025-02-19',2,'collected'),(10,'چک','4334','تجارت',333,'','2024-12-25',2,'pending'),(16,'چک','1232','توسعه',11,'','2025-04-24',18,'pending'),(17,'چک','1232','توسعه',11,'','2025-04-24',18,'pending'),(18,'چک','1232','ملت',11,'','2025-04-04',18,'collected'),(19,'چک','124000','سپhi',1200000,'','2025-04-24',19,'pending'),(20,'چک','12502551','توسعه تعاون',125000,'','2025-04-23',21,'pending'),(21,'مستقیم','','مستقیم',2,'','',22,'collected'),(22,'مستقیم','','مستقیم',25,'','',22,'collected'),(23,'مستقیم','','مستقیم',250,'','',22,'collected'),(24,'مستقیم','','مستقیم',2500,'','',22,'collected'),(26,'نقدی','###','نقدی',3,'','',23,'collected'),(27,'نقدی','###','نقدی',35,'','',23,'collected'),(28,'نقدی','###','نقدی',350,'','',23,'collected'),(29,'نقدی','###','نقدی',3500,'','',23,'collected'),(30,'نقدی','###','نقدی',35000,'','',23,'collected'),(31,'نقدی','###','نقدی',350000,'','',23,'collected'),(32,'نقدی','###','نقدی',3500000,'','',23,'collected'),(33,'چک','۳۲۳۲۲۵۱۲۱','ملی',1250000,'','2025-04-01',23,'pending'),(34,'نقدی','###','نقدی',2,'','1404/01/25 یک‌شنبه 12:55:06 ب.ظ',24,'collected'),(35,'نقدی','###','نقدی',20,'','1404/01/25 یک‌شنبه 12:55:06 ب.ظ',24,'collected'),(36,'نقدی','###','نقدی',200,'','1404/01/25 یک‌شنبه 12:55:06 ب.ظ',24,'collected'),(37,'نقدی','###','نقدی',2000,'','1404/01/25 یک‌شنبه 12:55:06 ب.ظ',24,'collected'),(38,'چک','۲۳۴۲۳','توسعه صادرات',3444,'','2025-03-30',24,'pending'),(39,'چک','۲۳۵۲۳۴','صادرات',35234,'','۱۴۰۳/۱۲/۲۷',25,'pending'),(40,'چک','24234','توسعه',1000,'','۱۴۰۴/۰۱/۴۴ ۰۲:۱۳:۵۴ ب ظ',26,'pending'),(41,'چک','34','توسعه صادرات',345,'','۱۴۰۴/۰۱/۵۵',27,'pending'),(42,'چک','1250','پارسیان',1200,'','۱۴۰۴/۰۱/۲۲',28,'pending'),(43,'نقدی','PMT-474053','نقدی',1200,'','۱۴۰۴/۰۱/۲۲',28,'collected'),(44,'چک','24234235','توسعه',12000,'','۱۴۰۴/۰۱/۳۳',29,'pending'),(45,'نقدی','PMT-157578','نقدی',12000,'','۱۴۰۴/۰۱/۳۳',29,'collected'),(46,'نقدی','PMT-451017','نقدی',2500,'','۱۴۰۴/۰۱/۳۳',30,'collected'),(47,'چک','123151','توسعه صادرات',2500,'','۱۴۰۴/۰۱/۳۳',31,'pending'),(48,'چک','123151','توسعه صادرات',2500,'','۱۴۰۴/۰۱/۳۳',31,'pending'),(49,'چک','123151','پاسارگاد',2500,'','۱۴۰۳/۱۱/۳۳',31,'pending'),(50,'نقدی','PMT-149724','نقدی',1350,'','۱۴۰۳/۱۱/۳۳',31,'collected'),(51,'چک','23232323','ملی',2000000,'','۱۴۰۴/۰۱/۵۵',32,'pending'),(52,'نقدی','PMT-452392','نقدی',100000,'','۱۴۰۴/۰۱/۵۵',32,'collected'),(53,'نقدی','PMT-50144','نقدی',1000,'','۱۴۰۴/۰۱/۶۶',33,'collected'),(54,'نقدی','PMT-22777','نقدی',1500,'','۱۴۰۴/۰۱/۳۳',34,'collected'),(55,'نقدی','PMT-643629','نقدی',1500,'','۱۴۰۴/۰۱/۳۳',34,'collected'),(56,'نقدی','PMT-359078','نقدی',1500,'','۱۴۰۴/۰۱/۳۳',34,'collected'),(57,'چک','','ملی',250000,'','۱۴۰۴/۰۱/۳۳',35,'pending'),(58,'نقدی','PMT-562471','نقدی',1250000,'','۱۴۰۴/۰۱/۳۳',35,'collected'),(59,'چک','۱۵۲۱۵۲۵','توسعه',10000,'','۱۴۰۳/۱۲/۴۴',36,'pending'),(60,'چک','1111','سرمایه',1000,'','۱۴۰۴/۰۱/۳۳',37,'pending');
/*!40000 ALTER TABLE `payments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `phonenumber` longtext,
  `address` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'hossein Soltanian','hosseinbidar7@gmail.com','$2a$14$OkZu91dlhRamYu1lQ9AtkuBN7cJLDg6AbAZQmM7VBNqA/Y535Mzu2','09125174854','','Admin'),(7,'wewe we','','','02134249381','sdfsd sdfsdsdsdfsd sdfsdsd  sdfsd sdfsdsdsdfsd sdfsdsd  sdfsd sdfsdsdsdfsd sdfsdsd','guest'),(14,'wefwe','','','2343','ergerr','guest'),(16,'فاکتور تست','','','23423423','sdfas','guest'),(18,'حسین','','','021521','sdffw','guest'),(21,'dfdft','','','۴۵۲۵۸۴۵۲۴۵','شسیب','guest'),(23,'خسته تر از همه','','','۰۲۱۳۴۲۴۹۳۸۵','سیبیسب','guest'),(24,'پرداخت تستی ۱','','','02134249381','سبی','guest'),(25,'sdvsv','','','15135','dfbdfb','guest'),(26,'یسبسی','','','۳۲۴۲۳','یسیبسل','guest'),(30,'حسین امیری','','','09125174845','sdf','guest'),(31,'امیر','','','09125174854','سیس','guest'),(32,'فاکتور مهم','','','09125174854','کرج حصارک شهرک یاس','guest'),(33,'ممد نوبری','','','02134249381','sdfdsf','guest'),(34,'فا','','','۲۳۴۲۳۴','۲۳ق۲۳','guest'),(35,'ممد مهرانی','','','۰۲۱۳۴۲۴۹۳۲۰','شسیشسی','guest'),(36,'ascas','','','025151515','نتدسیبسی','guest');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-04-15 12:03:44
