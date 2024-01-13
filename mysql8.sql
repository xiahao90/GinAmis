-- Adminer 4.8.1 MySQL 8.0.34 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(60) NOT NULL COMMENT '昵称',
  `admin` varchar(60) NOT NULL COMMENT '账号',
  `password` char(64) NOT NULL COMMENT '密码',
  `salt` char(64) NOT NULL COMMENT '加密盐',
  `role` json NOT NULL COMMENT '角色',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态',
  `superadmin` tinyint NOT NULL DEFAULT '0' COMMENT '超级管理员',
  `addtime` int NOT NULL DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `admin` (`id`, `name`, `admin`, `password`, `salt`, `role`, `status`, `superadmin`, `addtime`) VALUES
(1, '大哥', 'admin',  'e5859211254b81c4c760c30dc26f12fc14cda1a451458e2021ef1afa2d480db6', '61f8eb5de38129e4a2c97eafa1573eb7e77d513e311a3660afad11bdad62160f', '[1, 2]', 1,  1,  0),
(2, '张三', 'admin1', '06b259fd8bd29fb7592c4b4f38d3e4837a7069388c5b8a92776786c7d9f05eb1', 'd616fa68eeae3ce6af2f6a5fc180abd35935d745b6b98e52e6296ac2c08ee67a', '[1, 2]', 1,  0,  0),
(3, '72', '06f0885d', 'f42ab28bfd442f19b5e3d1be2fc94e9589f434301c32a3dff3062d0b7bf03367', '82773040dcc14bdf774a203c2d046fe80ad1d4fda746c128d9ead8c1aa1c27ff', '[]', 1,  1,  1704868466),
(4, '5e', 'ba4fcb18', '53f103027320043412ceeca1d11796b3dd642017d2ed53e15d287c2f37516dc0', 'a56bde8eae7e4d98615ea07e14d7cf763e2bd14bba6a681a42ebbba809390f6c', '[1, 2]', 0,  0,  1704868492),
(5, 'e0', '424b560f', '906421b5bb8f4a2f14bd58eedc0c7ab031b08b4a363e7e2c04dfacdfc3d49a72', '3692dae97655d74b658e2d5d435524e1620a4eb1e2d9b60e56857d8c767c29e5', '[1]',  1,  0,  1704868534),
(6, '王二', '95fed884', '1835f653f0a3a6e055f9c19109bb11a0aeaf409e775ecce71f7c36e47de8d649', 'ad7be207166285b219b5d32da505d352090259bda09b2eb17ddf29a4fce7397c', '[]', 1,  1,  1704868581),
(7, '28', 'bdc83a2d', 'd0fa57ec2617be4bbcef84eac368315240aad6fb11de739dd8c5114e6bec68a6', '936564fcb1279fa375e8774277a122d66389de5d58052f6179c917c50099f0fb', '[]', 1,  1,  1704870560),
(8, 'cf', '92a5a775', '7727d95cbf1363baf1f40f9871b58fb4335e687d47f354c0b849b608029ff78f', '34aadee0748f0199faa8ff08e7e0654a18b8835b4bedf1494fb827203f015a0f', '[]', 1,  1,  1704870563),
(9, '1c', 'f53f36ee', '5c6da091933bb368fd9addcf220ea2b035077d5c6255c7f6814f6b81432fc959', '9df213b134394359354d608855dcb0260d88f1d0d2cd194bef7b0ec555b710d1', '[]', 1,  1,  1704870566),
(10,  '6c', 'fc0a106c', '3014f00e531894a5d1e484c6c8de92e1ebaa07be61460c8011ed1b0695a877cd', '0d72ed1519e54b74350ce1c36320ce508eff35955924cd24208621b4f5c55126', '[]', 1,  1,  1704870568),
(11,  'a2', '3d7215a0', 'c04cdde2f5cc4f5ed7766f6d6391828fe87debb671bf5bae8917fdc19d8cf23c', '35daf62c2b932d7cceac6da3ea8b1a29b602bd2ddb419b2c77d7de0deeca8cb0', '[]', 1,  1,  1704870570),
(12,  'b3', 'fdb1fc03', '452bb9564f64b9da98fb40d0fa0b8bfdff472b0f9b9b00e89a46023f7b1b7e12', 'ec9125065220f972936cee967c9dfa24a54d48a449c556a6e3b7f232e6e297e7', '[1]',  1,  0,  1704870879),
(13,  '26', 'e493f8f2', '8653b3f23613379bbe92c9c1004d356ff01a4d1cedf0c206f6864c64dc82efbd', '72f662a9440cd6338b7bb42d5befc03d8bfa1dd409c9e1f789f869eb1b8cee98', '[1, 2]', 1,  0,  1704870900),
(14,  '9a', '6457f08d', 'af6283954d2f502c4f79071eaeaa4b12d591de9240870c848c731ffa54af9cc7', '16b5ddfc88fd3fc625b38580569955aac65cede95ce075e990f612e673321868', '[]', 1,  1,  1704870904),
(15,  '8d', 'fcec8c59', 'bd4f7e45bbf6aed5937cef03676f839bb334e10a30966f6da37d3e7d33238319', '77c2fd5463ac82a965e2ba26e8d6be509c9d17a121a3a36752c4e12f3cf97d4b', '[6, 8]', 1,  0,  1704870909),
(16,  'd9', 'af5d097b', '6e88d92f09ee9abd5350c7c451c49108b30a1f2593869a840ac04bea8a81153d', 'cb6d36c3ea6f071db22b7aff7cd5a7840e95f16ad6c37dbf2dee44401f0d09fb', '[]', 1,  1,  1704870928),
(17,  'de', 'd21afd3a', 'e79320483e11321d6136bf5fc0cdf0966047bee4b584d10d28dd9f79d67f3725', '5bc24fa6ca5c50c261754e0574046aa90f3b78d2eec10a485f01988ee1ff1708', '[]', 1,  1,  1704870930),
(18,  '0d', '5b077f03', 'a4de7823a2ec67ec66eea2813466f8b515550da2647b3402ff1e54f08f756d9a', '43beefa1eeaa70020494375e45ccdf2340fe4e57e7ea3554db65da52b3da41b4', '[2, 1]', 1,  0,  1704870934),
(19,  'af', 'e8857f21', 'aafa68ed297439319abcf5346339c54f3d31dfcf6a8523e81f603922e1530ec6', '235fdc67fdb26007f26faa7604b8135e08e2316bedd25856099aa05cd9d2a3b1', '[6, 9]', 1,  0,  1704870938),
(22,  '8c1ssss',  'e8720232', '93e989fcece24b29fb0cae8cb1148990d50ce1f3cf7fbc0fe35c0f0a712937b2', 'b8ce2fccc93b92dbc00da1cf59958bb0fc16e93d4bb13bd4b537e0e2ad268247', '[2]',  1,  0,  1704870953);

DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(60) NOT NULL COMMENT '角色名',
  `info` varchar(300) NOT NULL COMMENT '描述',
  `data_role` json NOT NULL COMMENT '权限',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `role` (`id`, `name`, `info`, `data_role`) VALUES
(1, '角色管理员',  '1',  '[\"HomePage\", \"RoleData\", \"RoleSchema\", \"RoleMin\", \"RoleCopy\", \"RoleEdit\"]'),
(2, '测试数据管理', '1',  '[\"HomePage\", \"HomePwdSchema\", \"HomePwd\", \"HomeRepwd\", \"TestData\", \"TestAdd\", \"TestSchema\", \"TestEdit\"]'),
(3, '111222', '1111222',  '[\"AdminData\", \"AdminSchema\", \"AdminMin\"]'),
(4, 'test1',  '1231231',  '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"AdminData\", \"AdminSchema\", \"AdminMin\"]'),
(6, '1231', '123',  '[\"AdminData\"]'),
(7, '1231', '123',  '[\"AdminData\", \"AdminAdd\"]'),
(9, '999',  '1',  '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"RoleSchema\", \"RoleMin\", \"RoleData\"]'),
(10,  '12', '2',  '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"RoleSchema\", \"RoleMin\", \"RoleData\"]'),
(11,  '10', '', '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"AdminData\", \"AdminAdd\", \"AdminEdit\", \"AdminDelete\", \"RoleDelete\", \"AdminSchema\", \"AdminMin\"]'),
(12,  '1',  '2',  '[\"HomePage\", \"AdminData\", \"AdminAdd\", \"AdminEdit\", \"AdminDelete\", \"AdminRepwd\", \"RoleData\", \"RoleAdd\", \"RoleEdit\", \"RoleDelete\", \"AdminSchema\", \"AdminMin\", \"RoleSchema\", \"RoleMin\"]'),
(13,  '1',  '2',  '[\"HomePage\", \"AdminData\", \"AdminAdd\", \"AdminEdit\", \"AdminDelete\", \"AdminRepwd\", \"RoleData\", \"RoleAdd\", \"RoleEdit\", \"RoleDelete\", \"AdminSchema\", \"AdminMin\", \"RoleSchema\", \"RoleMin\"]'),
(14,  '10', '', '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"AdminData\", \"AdminAdd\", \"AdminEdit\", \"AdminDelete\", \"RoleDelete\", \"AdminSchema\", \"AdminMin\"]'),
(15,  '10', '', '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"AdminData\", \"AdminAdd\", \"AdminEdit\", \"AdminDelete\", \"RoleDelete\", \"AdminSchema\", \"AdminMin\"]'),
(16,  '10', '', '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"AdminData\", \"AdminAdd\", \"AdminEdit\", \"AdminDelete\", \"RoleDelete\", \"AdminSchema\", \"AdminMin\"]'),
(17,  '10', '', '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"AdminData\", \"AdminAdd\", \"AdminEdit\", \"AdminDelete\", \"RoleDelete\", \"AdminSchema\", \"AdminMin\"]'),
(18,  '10', '', '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"AdminData\", \"AdminAdd\", \"AdminEdit\", \"AdminDelete\", \"RoleDelete\", \"AdminSchema\", \"AdminMin\"]'),
(20,  '10', '', '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"AdminData\", \"AdminAdd\", \"AdminEdit\", \"AdminDelete\", \"RoleDelete\", \"AdminSchema\", \"AdminMin\"]'),
(21,  '10', '', '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"AdminData\", \"AdminAdd\", \"AdminEdit\", \"AdminDelete\", \"RoleDelete\", \"AdminSchema\", \"AdminMin\"]'),
(22,  '10', '', '[\"RoleAdd\", \"RoleSchema\", \"RoleMin\", \"RoleData\", \"AdminData\", \"AdminAdd\", \"AdminEdit\", \"AdminDelete\", \"RoleDelete\", \"AdminSchema\", \"AdminMin\"]'),
(23,  '102',  'fsdfsdfsdf33', '[\"AdminData\", \"AdminAdd\"]');

DROP TABLE IF EXISTS `test`;
CREATE TABLE `test` (
  `id` int NOT NULL AUTO_INCREMENT,
  `addtime` int NOT NULL DEFAULT '0',
  `updatetime` int NOT NULL DEFAULT '0',
  `checkbox` tinyint NOT NULL,
  `switch` tinyint NOT NULL,
  `number` tinyint NOT NULL,
  `input` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `tag` json NOT NULL,
  `checkboxes` json NOT NULL,
  `select_1` json NOT NULL,
  `radios` tinyint NOT NULL,
  `select` tinyint NOT NULL,
  `datetime` int NOT NULL,
  `textarea` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='测试数据';

INSERT INTO `test` (`id`, `addtime`, `updatetime`, `checkbox`, `switch`, `number`, `input`, `password`, `tag`, `checkboxes`, `select_1`, `radios`, `select`, `datetime`, `textarea`) VALUES
(2, 1705162578, 1705170758, 0,  0,  10, 'fghfghdsf',  'dfgdfgdfgdfgdg', '[\"zhugeliang\", \"zhongwuyan\"]', '[1, 2, 6]',  '[1, 2, 3, 4, 4, 6]', 1,  1,  1704903619, 'gdfdgdfgdfgdfgdfg'),
(3, 1705163392, 1705167774, 1,  1,  5,  '1111111',  '11111111', '[\"zhongwuyan\", \"caocao\"]', '[1, 2]', '[1, 3, 4]',  1,  2,  1705163384, '12312312312'),
(4, 1705164820, 1705167767, 0,  0,  5,  '22222222', '2222222222', '[\"zhugeliang\", \"caocao\"]', '[1, 2]', '[4, 3]', 2,  1,  1706028817, '1231231231'),
(5, 1705166926, 1705169489, 0,  0,  8,  '会出现这是因为在', '4444444444', '[\"zhugeliang\", \"zhongwuyan\"]', '[1, 5, 3]',  '[3, 4]', 3,  4,  1706290124, '在Go语言中，当你试图向一个未初始化的map中赋值时，会出现\"assignment to entry in nil map\"的错误。这是因为在Go语言中，map是一个引用类型，它需要在使用之前进行初始化。'),
(6, 1705170732, 1705170763, 0,  0,  10, 'dfsadasdasdasdsad',  'dasdasdsadasd',  '[\"zhugeliang\", \"caocao\"]', '[1, 3]', '[5, 6]', 3,  2,  1705170724, '2222'),
(7, 1705171216, 0,  0,  0,  10, 'dfsadasdasdasdsad',  'dasdasdsadasd',  '[\"zhugeliang\", \"caocao\"]', '[1, 3]', '[5, 6]', 3,  2,  1705170724, '2222'),
(8, 1705171229, 0,  0,  0,  10, 'dfsadasdasdasdsad',  'dasdasdsadasd',  '[\"zhugeliang\", \"caocao\"]', '[1, 3]', '[5, 6]', 3,  2,  1705170724, '2222');

-- 2024-01-13 18:53:49