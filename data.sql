-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 07, 2024 at 03:36 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test-enterkomputer`
--

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` int(10) UNSIGNED NOT NULL,
  `orderId` varchar(255) NOT NULL,
  `productId` int(10) UNSIGNED NOT NULL,
  `quantity` int(255) NOT NULL,
  `totalPrice` int(255) NOT NULL,
  `meja` enum('1','2','3') NOT NULL,
  `createdAt` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`id`, `orderId`, `productId`, `quantity`, `totalPrice`, `meja`, `createdAt`) VALUES
(1, '070824OZPyy', 1, 1, 12000, '1', '2024-08-07 13:35:45'),
(2, '070824OZPyy', 6, 1, 6000, '1', '2024-08-07 13:35:45'),
(3, '070824OZPyy', 10, 2, 46000, '1', '2024-08-07 13:35:45'),
(4, '070824OZPyy', 3, 1, 8000, '1', '2024-08-07 13:35:45'),
(5, '070824OZPyy', 7, 1, 15000, '1', '2024-08-07 13:35:45');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `category` enum('Makanan','Minuman','Promo') NOT NULL,
  `variant` varchar(255) DEFAULT NULL,
  `price` int(255) NOT NULL,
  `createdAt` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `category`, `variant`, `price`, `createdAt`) VALUES
(1, 'Jeruk', 'Minuman', 'Dingin', 12000, '2024-08-05 14:21:02'),
(2, 'Jeruk', 'Minuman', 'Panas', 10000, '2024-08-05 14:21:02'),
(3, 'Teh', 'Minuman', 'Manis', 8000, '2024-08-05 14:21:02'),
(4, 'Teh', 'Minuman', 'Tawar', 5000, '2024-08-05 14:21:02'),
(5, 'Kopi', 'Minuman', 'Dingin', 8000, '2024-08-05 14:21:02'),
(6, 'Kopi', 'Minuman', 'Panas', 6000, '2024-08-05 14:21:02'),
(7, 'Mie', 'Makanan', 'Goreng', 15000, '2024-08-05 14:21:02'),
(8, 'Mie', 'Minuman', 'Kuah', 15000, '2024-08-05 14:21:02'),
(9, 'Nasi', 'Makanan', 'Goreng', 15000, '2024-08-05 14:21:02'),
(10, 'Nasi Goreng + Jeruk Dingin', 'Promo', NULL, 23000, '2024-08-05 14:21:02'),
(11, 'Extra Es Batu', 'Minuman', NULL, 2000, '2024-08-05 14:21:02');

-- --------------------------------------------------------

--
-- Table structure for table `schema_migrations`
--

CREATE TABLE `schema_migrations` (
  `version` bigint(20) NOT NULL,
  `dirty` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `schema_migrations`
--

INSERT INTO `schema_migrations` (`version`, `dirty`) VALUES
(20240806042533, 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `productId` (`productId`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `schema_migrations`
--
ALTER TABLE `schema_migrations`
  ADD PRIMARY KEY (`version`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`productId`) REFERENCES `products` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
