DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `user_name` varchar(50) NOT NULL,
  `password` varchar(100) DEFAULT NULL,
  `thumb` varchar(500) DEFAULT NULL,
  `about_me` varchar(2000) DEFAULT NULL,
  `site_title` varchar(100) DEFAULT NULL,
  `head_bg_pic` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL,
  `publish_at` datetime NOT NULL,
  `content` text,
  `thumb` varchar(500) DEFAULT NULL,
  `tag` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `guest_name` varchar(50) DEFAULT NULL,
  `content` varchar(2000) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `post_id` int(11) NOT NULL,
  `post_title` varchar(200) DEFAULT NULL,
  `reply` varchar(2000) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8

INSERT INTO `user` (`id`, `user_name`, `password`, `thumb`, `about_me`, `site_title`, `head_bg_pic`) VALUES ('1', 'Gordon', 'c3284d0f94606de1fd2af172aba15bf3', '/static/uploads/D3784F4E-4879-4ED0-BBE5-7D9DD3D70B6F.png', 'iNote是一款开源，免费，简洁的单页博客', 'iNote', null);

