-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 12, 2024 at 08:27 AM
-- Server version: 10.4.17-MariaDB
-- PHP Version: 7.2.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `chat`
--

-- --------------------------------------------------------

--
-- Table structure for table `messages`
--

CREATE TABLE `messages` (
  `id` int(11) NOT NULL,
  `japanese_text` text DEFAULT NULL,
  `english_text` text DEFAULT NULL,
  `user_id` varchar(255) DEFAULT NULL,
  `company_id` varchar(255) DEFAULT NULL,
  `chat_room_id` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `messages`
--

INSERT INTO `messages` (`id`, `japanese_text`, `english_text`, `user_id`, `company_id`, `chat_room_id`, `created_at`) VALUES
(1, 'おはよう。', 'Hello', 'sano@107.jp', '107', 'Ato_107', '2024-01-11 09:13:02'),
(2, 'こんにちは、元気ですか？', 'Hello, how are you?', 'tanaka@example.com', 'XYZCorp', 'Room1', '2024-01-11 09:13:02'),
(3, '今日はいい天気ですね。', 'The weather is nice today.', 'suzuki@example.com', 'ABCInc', 'Room1', '2024-01-11 09:13:02'),
(4, '元気です、ありがとう！', 'I\'m doing well, thanks!', 'sano@107.jp', '107', 'Ato_107', '2024-01-11 09:13:02'),
(5, '週末の予定はありますか？', 'Do you have any plans for the weekend?', 'tanaka@example.com', 'XYZCorp', 'Room1', '2024-01-11 09:13:02'),
(6, 'まだです、まだ決めていません。', 'Not yet, still deciding.', 'suzuki@example.com', 'ABCInc', 'Room1', '2024-01-11 09:13:02'),
(7, '決めたら教えてくださいね。', 'Let me know when you decide.', 'tanaka@example.com', 'XYZCorp', 'Room1', '2024-01-11 09:13:02'),
(8, 'もちろん、お知らせします！', 'Sure, will do!', 'sano@107.jp', '107', 'Ato_107', '2024-01-11 09:13:02'),
(9, 'ようこそ', 'Selamat datang', 'sano@107.jp', '107', 'Ato_107', '2024-01-11 05:08:23'),
(14, '今日は雨が降りますか？(Kyō wa ame ga furimasu ka?)', 'Hari ini apakah hujan?', 'ato@107.jp', '107', 'Ato_107', '2024-01-12 00:26:40');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `messages`
--
ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id_index` (`user_id`),
  ADD KEY `company_id_index` (`company_id`),
  ADD KEY `chat_room_id_index` (`chat_room_id`),
  ADD KEY `created_at_index` (`created_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `messages`
--
ALTER TABLE `messages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
