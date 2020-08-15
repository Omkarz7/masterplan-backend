-- MySQL dump 10.13  Distrib 5.6.46, for Linux (x86_64)
--
-- Host: localhost    Database: Construction_Industry
-- ------------------------------------------------------
-- Server version	5.6.46

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `Masterplan`
--

DROP TABLE IF EXISTS `Masterplan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Masterplan` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_id` int(11) NOT NULL,
  `sr_no` varchar(100) NOT NULL,
  `activity` varchar(500) NOT NULL,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  PRIMARY KEY (`id`),
  KEY `Masterplan_FK` (`project_id`),
  CONSTRAINT `Masterplan_FK` FOREIGN KEY (`project_id`) REFERENCES `ProjectDetails` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Masterplan`
--

LOCK TABLES `Masterplan` WRITE;
/*!40000 ALTER TABLE `Masterplan` DISABLE KEYS */;
INSERT INTO `Masterplan` VALUES (1,1,'1','House','2018-08-01','2019-02-28'),(2,1,'1.1','Foundation','2018-08-01','2019-09-04'),(3,1,'1.1.1','Digging soil','2018-08-01','2019-08-10'),(4,1,'1.1.2','Piling','2018-08-11','2019-09-04'),(5,1,'1.2','Floor','2018-09-06','2019-11-04'),(6,1,'1.2.1','Tiling','2018-09-06','2019-11-04'),(7,1,'1.3','Walls','2018-11-06','2019-11-04'),(8,1,'1.4','Roof','2019-01-15','2019-02-28'),(9,1,'1.5','Boundary wall','2018-08-01','2018-09-02');
/*!40000 ALTER TABLE `Masterplan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ProjectDetails`
--

DROP TABLE IF EXISTS `ProjectDetails`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ProjectDetails` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_manager_id` int(11) NOT NULL,
  `other_project_related_fields` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `ProjectDetails_FK` (`project_manager_id`),
  CONSTRAINT `ProjectDetails_FK` FOREIGN KEY (`project_manager_id`) REFERENCES `UserCredentials` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ProjectDetails`
--

LOCK TABLES `ProjectDetails` WRITE;
/*!40000 ALTER TABLE `ProjectDetails` DISABLE KEYS */;
INSERT INTO `ProjectDetails` VALUES (1,1,'Some Vague Details about the project');
/*!40000 ALTER TABLE `ProjectDetails` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `UserCredentials`
--

DROP TABLE IF EXISTS `UserCredentials`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `UserCredentials` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `password` varchar(100) DEFAULT NULL,
  `other_irrelevant _fields` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `UserCredentials`
--

LOCK TABLES `UserCredentials` WRITE;
/*!40000 ALTER TABLE `UserCredentials` DISABLE KEYS */;
INSERT INTO `UserCredentials` VALUES (1,'manager1','KwBcyLYQ+liZqfnlkmcbupd2oOd4x/iNubVO70hJDpQ=','irrelevant data');
/*!40000 ALTER TABLE `UserCredentials` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'Construction_Industry'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-16  0:37:24
