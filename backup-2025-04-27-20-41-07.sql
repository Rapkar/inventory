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
  `role_price` double DEFAULT NULL,
  `meter_price` double DEFAULT NULL,
  `count` bigint DEFAULT NULL,
  `meter` double DEFAULT NULL,
  `total_price` double DEFAULT NULL,
  `weight` double DEFAULT NULL,
  `weight_price` double DEFAULT NULL,
  `inventory_id` bigint unsigned DEFAULT NULL,
  `product_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_export_products_product_id` (`product_id`),
  KEY `idx_export_products_export_id` (`export_id`),
  KEY `idx_export_products_inventory_id` (`inventory_id`),
  CONSTRAINT `fk_export_products_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`),
  CONSTRAINT `fk_export_products_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  CONSTRAINT `fk_exports_export_products` FOREIGN KEY (`export_id`) REFERENCES `exports` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `export_products`
--

LOCK TABLES `export_products` WRITE;
/*!40000 ALTER TABLE `export_products` DISABLE KEYS */;
INSERT INTO `export_products` VALUES (1,1,'ایزوگام شرق',99250,102500,100,10,2000000,0,0,1,1),(6,6,'ایزوگام شرق',99250,102500,1,0,99250,0,0,1,1),(7,7,'ایزوگام شرق',99250,102500,10,0,992500,0,0,1,1),(8,8,'ایزوگام شرق',99250,102500,10,0,2017500,0,0,1,1),(9,9,'ایزوگام شرق',99250,102500,5,0,1008750,0,0,1,1),(10,10,'ایزوگام شرق',99250,102500,5,30,3571250,0,0,1,1),(11,11,'ایزوگام شرق',99250,102500,0,0,0,0,0,1,1),(12,12,'ایزوگام شرق',99250,102500,1,1,201750,0,0,1,1),(13,13,'ایزوگام شرق',99250,102500,1,1,201750,1,0,1,1),(14,14,'ایزوگام شرق',99250,102500,1,0,99250,1,0,1,1),(15,15,'ایزوگام شرق',99250,102500,0,1,102500,3,0,1,1),(16,16,'ایزوگام شرق',99250,102500,2,2,403500,5,0,1,1),(17,17,'ایزوگام شرق',99250,102500,5,5,1008750,0,0,1,1),(18,18,'ایزوگام شرق',99250,102500,1,1,201750,10,0,1,1),(19,19,'ایزوگام شرق',99250,102500,10,1,1095000,10,0,1,1),(20,20,'ایزوگام شرق',99250,102500,10,1,1095000,1,0,1,1),(21,21,'ایزوگام شرق',99250,102500,1,1,201750,1,0,1,1);
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
  `inventory_id` bigint unsigned DEFAULT NULL,
  `draft` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exports_user_id` (`user_id`),
  KEY `idx_exports_inventory_id` (`inventory_id`),
  CONSTRAINT `fk_exports_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`),
  CONSTRAINT `fk_exports_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exports`
--

LOCK TABLES `exports` WRITE;
/*!40000 ALTER TABLE `exports` DISABLE KEYS */;
INSERT INTO `exports` VALUES (1,'حسین سلطانیان','9283422','09125174854',1,'کرج -کرج=-ایران -سیسی',10000000,10,'','1404/02/05 پنج‌شنبه 01:45:45 ق.ظ',1,0),(6,'hossein soltanian','O04740','09125174854',1,'اتنات',1299250,1200000,'','1404/02/05 پنج‌شنبه 09:41:42 ق.ظ',1,0),(7,'hossein soltanian','Y04835','09125174854',1,'5g',1117500,125000,'','1404/02/05 پنج‌شنبه 11:39:38 ق.ظ',1,0),(8,'dfgdf','E07273','09125174854',1,'',2137500,120000,'','1404/02/05 پنج‌شنبه 11:40:25 ق.ظ',1,0),(9,'hossein soltanian','A01977','09125174854',1,'صثثصب',13508750,12500000,'','1404/02/05 پنج‌شنبه 12:27:44 ب.ظ',1,0),(10,'hossein soltanian','A09704','09125174854',1,'',3583250,12000,'','1404/02/05 پنج‌شنبه 12:50:27 ب.ظ',1,0),(11,'hossein soltanian','U08785','09125174854',1,'gjgh',10000,10000,'','1404/02/05 پنج‌شنبه 12:51:25 ب.ظ',1,0),(12,'hossein soltanian','C00798','02134249381',2,'wef',213750,12000,'','1404/02/07 شنبه 10:00:53 ق.ظ',1,0),(13,'hossein soltanian','G06977','09125174855',3,'sdsd',213750,12000,'','1404/02/07 شنبه 10:06:14 ق.ظ',1,0),(14,'hossein soltanian','V02174','09125174854',1,'asdas',319250,220000,'sdsd','1404/02/08 یک‌شنبه 12:31:24 ب.ظ',1,1),(15,'hossein soltanian','E00910','09125174854',1,'sdsd',322500,220000,'sdsd','1404/02/08 یک‌شنبه 12:43:45 ب.ظ',1,0),(16,'hossein soltanian','V07201','09125174854',1,'dsdgdfg',403500,0,'rhstser','1404/02/08 یک‌شنبه 12:45:14 ب.ظ',1,0),(17,'hossein soltanian','M07689','09125174854',1,'fghfgh',1134250,125500,'fghfg','1404/02/08 یک‌شنبه 12:46:24 ب.ظ',1,0),(18,'hossein soltanian','K01258','09125174854',1,'یبلیبل',401750,200000,'یلیب','1404/02/08 یک‌شنبه 12:53:53 ب.ظ',1,1),(19,'hossein soltanian','Z00334','09125174854',1,'sdfdsf',9095000,8000000,'dsfsdf','1404/02/08 یک‌شنبه 12:55:39 ب.ظ',1,0),(20,'hossein soltanian','S08481','09125174854',1,'لاتلات',5095000,4000000,'dfgdf','1404/02/08 یک‌شنبه 12:59:08 ب.ظ',1,0),(21,'hossein soltanian','F03122','09125174854',1,'WQWS',25201750,25000000,'UNIUNI','1404/02/08 یک‌شنبه 01:00:59 ب.ظ',1,1);
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inventories`
--

LOCK TABLES `inventories` WRITE;
/*!40000 ALTER TABLE `inventories` DISABLE KEYS */;
INSERT INTO `inventories` VALUES (1,'انبار اشتهارد1'),(4,'انبار زنجان');
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
  KEY `idx_payments_inventory_id` (`inventory_id`),
  KEY `idx_payments_export_id` (`export_id`),
  KEY `idx_payments_user_id` (`user_id`),
  CONSTRAINT `fk_exports_payments` FOREIGN KEY (`export_id`) REFERENCES `exports` (`id`),
  CONSTRAINT `fk_payments_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`),
  CONSTRAINT `fk_payments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payments`
--

LOCK TABLES `payments` WRITE;
/*!40000 ALTER TABLE `payments` DISABLE KEYS */;
INSERT INTO `payments` VALUES (1,'نقدی','PMT-353608','نقدی',1200000,'','۱۴۰۴/۰۲/۰۵',6,1,1,'collected'),(2,'نقدی','PMT-381208','نقدی',129000,'','۱۴۰۴/۰۲/۰۵',7,1,1,'collected'),(3,'چک','1124','ملی',118000,'','۱۴۰۴/۰۲/۰۵',7,1,1,'pending'),(4,'نقدی','PMT-703820','نقدی',120000,'','۱۴۰۴/۰۲/۰۵',8,1,1,'collected'),(5,'نقدی','PMT-193509','نقدی',12500000,'','۱۴۰۴/۰۲/۰۵',9,1,1,'collected'),(6,'نقدی','PMT-570352','نقدی',49625,'','۱۴۰۴/۰۲/۰۵',10,1,1,'collected'),(7,'نقدی','PMT-171620','نقدی',10000,'','۱۴۰۴/۰۲/۰۵',11,1,1,'collected'),(8,'نقدی','PMT-967392','نقدی',13000,'','۱۴۰۴/۰۲/۰۷',12,2,1,'collected'),(9,'نقدی','PMT-729746','نقدی',25000,'','۱۴۰۴/۰۲/۰۷',13,3,1,'collected'),(10,'نقدی','PMT-990147','نقدی',1578000,'','۱۴۰۴/۰۲/۰۸',14,1,1,'collected'),(11,'نقدی','PMT-101450','نقدی',496250,'','۱۴۰۴/۰۲/۰۸',15,1,1,'collected'),(12,'نقدی','PMT-762070','نقدی',2045444,'','۱۴۰۴/۰۲/۰۸',16,1,1,'collected'),(13,'نقدی','PMT-258159','نقدی',1200,'','۱۴۰۴/۰۲/۰۸',17,1,1,'collected'),(14,'نقدی','PMT-865473','نقدی',12000000,'','۱۴۰۴/۰۲/۰۸',18,1,1,'collected'),(15,'نقدی','PMT-388880','نقدی',780000,'','۱۴۰۴/۰۲/۰۸',19,1,1,'collected'),(16,'نقدی','PMT-20843','نقدی',420000,'','۱۴۰۴/۰۲/۰۸',20,1,1,'collected'),(17,'نقدی','PMT-73660','نقدی',12300000,'','۱۴۰۴/۰۲/۰۸',21,1,1,'collected');
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
  `role_price` double DEFAULT NULL,
  `meter_price` double DEFAULT NULL,
  `weight_price` double DEFAULT NULL,
  `count` bigint DEFAULT NULL,
  `meter` double DEFAULT NULL,
  `weight` double DEFAULT NULL,
  `inventory_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_inventory_id` (`inventory_id`),
  CONSTRAINT `fk_products_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'ایزوگام شرق',99250,102500,1100,38,56,68,1),(2,'محصول خاص',1200,1300,1400,12,13,14,4);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'hossein Soltanian','hosseinbidar7@gmail.com','$2a$14$r7OEixxuTzWVz61WlGptz.af8rQ2xFUhKXTuDA5Xb0CaV6vwpOq.S','09125174854','','Admin'),(2,'hossein soltanian','','','02134249381','wef','guest'),(3,'hossein soltanian','','','09125174855','sdsd','guest');
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

-- Dump completed on 2025-04-28 20:41:07
