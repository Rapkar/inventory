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
  PRIMARY KEY (`id`),
  KEY `idx_export_products_export_id` (`export_id`),
  KEY `idx_export_products_inventory_id` (`inventory_id`),
  KEY `idx_export_products_product_id` (`product_id`),
  CONSTRAINT `fk_export_products_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`),
  CONSTRAINT `fk_export_products_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  CONSTRAINT `fk_exports_export_products` FOREIGN KEY (`export_id`) REFERENCES `exports` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `export_products`
--

LOCK TABLES `export_products` WRITE;
/*!40000 ALTER TABLE `export_products` DISABLE KEYS */;
INSERT INTO `export_products` VALUES (2,1,' ایزوگام شرق صادراتی',99250,102500,0,0,0,0,10,0,100,0,2000000,1,1),(3,1,'ایزوگام غرب شرق مخصوص',87500,92500,0,0,0,0,15,0,80,0,1500000,1,1),(4,1,'ایزوگام  شرق بدون فویل',110000,115000,0,0,0,0,12,0,50,0,1800000,1,1),(5,1,'ایزوگام سپید گام صادراتی',95000,98000,0,0,0,0,18,0,120,0,2200000,1,1),(6,1,'ایزوگام سپیدگام صادراتی بدون فویل',105000,108000,0,0,0,0,20,0,75,0,2500000,1,1),(7,1,'ایزوگام اصلاحی درجه ۲',0,108000,0,0,0,0,750,0,0,0,2500000,1,1),(8,1,'ایزوگام شرق طرح دار',105000,0,0,0,0,75,20,0,0,0,2500000,1,1),(9,1,'بشکه قیر',0,0,0,0,108000,0,20,0,0,75,2500000,1,1);
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
  PRIMARY KEY (`id`),
  KEY `idx_exports_user_id` (`user_id`),
  KEY `idx_exports_inventory_id` (`inventory_id`),
  CONSTRAINT `fk_exports_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`),
  CONSTRAINT `fk_exports_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `exports`
--

LOCK TABLES `exports` WRITE;
/*!40000 ALTER TABLE `exports` DISABLE KEYS */;
INSERT INTO `exports` VALUES (1,'حسین سلطانیان','9283422','09125174854',1,'کرج -کرج=-ایران -سیسی',10000000,10,'','1404/02/10 سه‌شنبه 02:44:29 ب.ظ',0,2);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payments`
--

LOCK TABLES `payments` WRITE;
/*!40000 ALTER TABLE `payments` DISABLE KEYS */;
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
  PRIMARY KEY (`id`),
  KEY `idx_products_inventory_id` (`inventory_id`),
  CONSTRAINT `fk_products_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,' ایزوگام شرق صادراتی',99250,102500,0,0,0,0,0,0,100,0,2),(2,'ایزوگام غرب شرق مخصوص',87500,92500,0,0,0,0,0,0,80,0,2),(3,'ایزوگام  شرق بدون فویل',110000,115000,0,0,0,0,0,0,50,0,2),(4,'ایزوگام سپید گام صادراتی',95000,98000,0,0,0,0,0,0,120,0,2),(5,'ایزوگام سپیدگام صادراتی بدون فویل',105000,108000,0,0,0,0,0,0,75,0,2),(6,'ایزوگام اصلاحی درجه ۲',0,108000,0,0,0,0,750,0,0,0,2),(7,'ایزوگام شرق طرح دار',105000,0,0,0,0,75,0,0,0,0,2),(8,'بشکه قیر',0,0,0,0,108000,0,0,0,0,75,2);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'hossein Soltanian','hosseinbidar7@gmail.com','$2a$14$6KrIwjCDB76t027/L5x6t.BpOIZq5C87lLakw3QTn.CWwCbJ.FGxW','09125174854','','Admin');
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

-- Dump completed on 2025-04-29 14:46:42
