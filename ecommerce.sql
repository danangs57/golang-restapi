-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 08 Jan 2023 pada 03.46
-- Versi server: 10.4.24-MariaDB
-- Versi PHP: 7.4.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ecommerce`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `merchs`
--

CREATE TABLE `merchs` (
  `id` bigint(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `quantity` int(255) NOT NULL,
  `seller` bigint(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `merchs`
--

INSERT INTO `merchs` (`id`, `name`, `quantity`, `seller`) VALUES
(7, 'Goku', 87, 1),
(8, 'Kulkas', 25, 1),
(9, 'Genteng', 192, 1),
(10, 'Burung', 9, 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `purchases`
--

CREATE TABLE `purchases` (
  `id` bigint(255) NOT NULL,
  `buyer_id` int(255) NOT NULL,
  `seller_id` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `name` varchar(255) NOT NULL,
  `type` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `type`) VALUES
(1, 'danang', 'seller'),
(2, 'satriani', 'buyer');

-- --------------------------------------------------------

--
-- Struktur dari tabel `user_purchase`
--

CREATE TABLE `user_purchase` (
  `id` bigint(255) NOT NULL,
  `merch_id` int(255) NOT NULL,
  `quantity` int(255) NOT NULL,
  `status` int(11) NOT NULL DEFAULT 0,
  `purchase_id` bigint(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `user_purchase`
--

INSERT INTO `user_purchase` (`id`, `merch_id`, `quantity`, `status`, `purchase_id`) VALUES
(5, 7, 3, 0, 8),
(6, 8, 2, 0, 8),
(7, 9, 2, 0, 9),
(8, 9, 2, 0, 9),
(9, 9, 2, 0, 10),
(10, 9, 2, 0, 10),
(11, 9, 2, 0, 11),
(12, 9, 2, 0, 11),
(13, 8, 3, 0, 12),
(14, 9, 2, 0, 12),
(15, 8, 3, 0, 13),
(16, 9, 2, 0, 13),
(17, 8, 3, 0, 14),
(18, 9, 2, 0, 14),
(19, 8, 3, 0, 15);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `merchs`
--
ALTER TABLE `merchs`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `purchases`
--
ALTER TABLE `purchases`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `user_purchase`
--
ALTER TABLE `user_purchase`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `merchs`
--
ALTER TABLE `merchs`
  MODIFY `id` bigint(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT untuk tabel `purchases`
--
ALTER TABLE `purchases`
  MODIFY `id` bigint(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `user_purchase`
--
ALTER TABLE `user_purchase`
  MODIFY `id` bigint(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
