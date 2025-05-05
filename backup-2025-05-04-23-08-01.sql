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
  `rolle_price` double DEFAULT NULL,
  `meter_price` double DEFAULT NULL,
  `weight_price` double DEFAULT NULL,
  `count_price` double DEFAULT NULL,
  `barrel_price` double DEFAULT NULL,
  `roll` bigint DEFAULT NULL,
  `meter` double DEFAULT NULL,
  `weight` double DEFAULT NULL,
  `count` bigint DEFAULT NULL,
  `barrel` bigint DEFAULT NULL,
  `total_price` double DEFAULT NULL,
  `inventory_id` bigint unsigned DEFAULT NULL,
  `product_id` bigint unsigned DEFAULT NULL,
  `measurement_system` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_export_products_export_id` (`export_id`),
  KEY `idx_export_products_inventory_id` (`inventory_id`),
  KEY `idx_export_products_product_id` (`product_id`),
  CONSTRAINT `fk_export_products_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`),
  CONSTRAINT `fk_export_products_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  CONSTRAINT `fk_exports_export_products` FOREIGN KEY (`export_id`) REFERENCES `exports` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `export_products`
--

LOCK TABLES `export_products` WRITE;
/*!40000 ALTER TABLE `export_products` DISABLE KEYS */;
INSERT INTO `export_products` VALUES (1,1,' ایزوگام شرق صادراتی',99250,0,0,0,0,5,0,0,0,0,496250,2,1,''),(2,1,'ایزوگام غرب شرق مخصوص',87500,0,0,0,0,5,0,0,0,0,437500,2,2,''),(3,1,'ایزوگام شمال شرق بدون فویل',110000,0,0,0,0,10,0,0,0,0,1100000,2,3,''),(4,1,'ایزوگام سپید گام صادراتی',95000,0,0,0,0,100,0,0,0,0,9500000,2,4,''),(5,1,'ایزوگام سپیدگام صادراتی بدون فویل',105000,0,0,0,0,10,0,0,0,0,1050000,2,5,''),(6,1,'ایزوگام اصلاحی درجه ۲',0,108000,0,0,0,0,10.5,0,0,0,1134000,2,6,''),(7,1,'ایزوگام شرق طرح دار',95000,0,0,0,0,75,0,0,0,0,7125000,2,7,''),(8,1,'بشکه قیر',0,0,0,0,108000,0,0,0,0,75,8100000,2,8,''),(9,2,'ایزوگام سپیدگام صادراتی بدون فویل',0,0,0,0,0,0,0,0,1,0,0,1,5,''),(10,3,' ایزوگام شرق صادراتی',15,0,0,0,0,10,0,0,0,0,99251,1,1,''),(11,3,'بشکه قیر',0,0,0,0,250000,0,0,0,0,2,432000,1,8,''),(12,3,'ایزوگام اصلاحی درجه ۲',0,108100,0,0,0,0,45,0,0,0,4864500,1,6,''),(13,4,' ایزوگام شرق صادراتی',0,0,0,0,0,0,0,0,0,0,1985020,1,1,''),(14,4,'بشکه قیر',0,0,0,0,0,0,0,0,0,0,1080000,1,8,''),(15,4,'ایزوگام اصلاحی درجه ۲',0,108100,0,0,0,0,5,0,0,0,540500,1,6,'');
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
  `user_id` bigint unsigned DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `total_price` double DEFAULT NULL,
  `tax` bigint DEFAULT NULL,
  `describe` varchar(255) DEFAULT NULL,
  `created_at` longtext,
  `draft` tinyint(1) DEFAULT NULL,
  `inventory_id` bigint unsigned DEFAULT NULL,
  `creator_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exports_user_id` (`user_id`),
  KEY `idx_exports_inventory_id` (`inventory_id`),
  CONSTRAINT `fk_exports_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`),
  CONSTRAINT `fk_exports_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exports`
--

LOCK TABLES `exports` WRITE;
/*!40000 ALTER TABLE `exports` DISABLE KEYS */;
INSERT INTO `exports` VALUES (1,'حسین سلطانیان','9283422','09125174854',1,'کرج -کرج=-ایران -سیسی',10000000,10,'','1404/02/10 سه‌شنبه 04:59:09 ب.ظ',1,2,NULL),(2,'ali','A04242','02534161616',2,'سیبی سیبسی',1200,1200,'لاتلا','1404/02/11 چهارشنبه 12:06:48 ب.ظ',1,1,NULL),(3,'حسینم','Z07435','02134249381',3,'sdsd',5395751,0,'','1404/02/12 پنج‌شنبه 02:22:41 ق.ظ',0,1,NULL),(4,'سیز','S07195','345345',4,'ضصضص',6603605520,2200000000,'wewe','1404/02/12 پنج‌شنبه 10:31:09 ق.ظ',1,1,'hossein Soltanian');
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
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventories`
--

LOCK TABLES `inventories` WRITE;
/*!40000 ALTER TABLE `inventories` DISABLE KEYS */;
INSERT INTO `inventories` VALUES (1,'انبار اشتهارد'),(2,'انبار زنجان');
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
  `total_price` double DEFAULT NULL,
  `describe` varchar(255) DEFAULT NULL,
  `created_at` longtext,
  `export_id` bigint unsigned DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `inventory_id` bigint unsigned DEFAULT NULL,
  `status` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_payments_export_id` (`export_id`),
  KEY `idx_payments_user_id` (`user_id`),
  KEY `idx_payments_inventory_id` (`inventory_id`),
  CONSTRAINT `fk_exports_payments` FOREIGN KEY (`export_id`) REFERENCES `exports` (`id`),
  CONSTRAINT `fk_payments_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`),
  CONSTRAINT `fk_payments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payments`
--

LOCK TABLES `payments` WRITE;
/*!40000 ALTER TABLE `payments` DISABLE KEYS */;
INSERT INTO `payments` VALUES (1,'نقدی','PMT-367380','نقدی',1200,'','۱۴۰۴/۰۲/۱۱',2,2,1,'collected'),(2,'نقدی','PMT-304684','نقدی',251000,'','۱۴۰۴/۰۲/۱۲',3,3,1,'collected'),(3,'نقدی','PMT-935776','نقدی',250000000,'','۱۴۰۴/۰۲/۲۹',4,4,1,'collected'),(4,'چک','43523435','ملی',1000000000,'','۱۴۰۴/۰۲/۱۲',4,4,1,'pending'),(5,'چک','43523436','ملی',1000000000,'','۱۴۰۴/۰۲/۱۲',4,4,1,'pending'),(6,'چک','43623434','ملی',1000000000,'','۱۴۰۴/۰۲/۲۹',4,4,1,'pending');
/*!40000 ALTER TABLE `payments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `rolle_price` double DEFAULT NULL,
  `meter_price` double DEFAULT NULL,
  `weight_price` double DEFAULT NULL,
  `count_price` double DEFAULT NULL,
  `barrel_price` double DEFAULT NULL,
  `roll` bigint DEFAULT NULL,
  `meter` double DEFAULT NULL,
  `weight` double DEFAULT NULL,
  `count` bigint DEFAULT NULL,
  `barrel` bigint DEFAULT NULL,
  `inventory_id` bigint unsigned DEFAULT NULL,
  `measurement_system` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_inventory_id` (`inventory_id`),
  CONSTRAINT `fk_products_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,' ایزوگام شرق صادراتی',99251,0,0,0,0,140,3,0,111,3,1,'roll'),(2,'ایزوگام غرب شرق مخصوص',92510,0,0,0,0,80,0,0,2,0,1,'roll'),(3,'ایزوگام شمال شرق بدون فویل',110000,0,0,0,0,115000,0,0,0,0,2,'roll'),(4,'ایزوگام سپید گام صادراتی',95000,0,0,0,1000,120,0,0,10,3,1,'roll'),(5,'ایزوگام سپیدگام صادراتی بدون فویل',87500,0,0,0,0,80,0,0,19,0,1,'roll'),(6,'ایزوگام اصلاحی درجه ۲',0,108100,0,0,0,0,10,0,0,0,1,'meter'),(7,'ایزوگام شرق طرح دار',95000,0,0,0,0,120,0,0,0,0,2,'roll'),(8,'بشکه قیر',0,0,0,0,108000,0,0,0,0,77,1,'barrel'),(11,'یه بشکه',0,0,0,0,25100,0,0,0,0,160,1,'barrel');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
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
  `phonenumber` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_phonenumber` (`phonenumber`),
  KEY `unique` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'hossein Soltanian','hosseinbidar7@gmail.com','$2a$14$r0evCl.KCZUnXh3cw2x/be0D0E1/AoclpUVPUSSFB0Re2BAH/6C7.','09125174854','','Admin'),(2,'ali','','','02534161616','سیبی سیبسی','guest'),(3,'حسینم','','','02134249381','sdsd','guest'),(4,'سیز','','','345345','ضصضص','guest');
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

-- Dump completed on 2025-05-05 23:08:01
