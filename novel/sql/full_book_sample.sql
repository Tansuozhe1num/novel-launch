-- 示例：插入一本书、其章节与正文内容
INSERT INTO `books` (`book_name`,`author`,`desc`,`tag`,`cover`,`create_time`,`update_time`,`status`,`is_deleted`,`heat`)
VALUES ('银河回响','星航者','在星海中寻找失落文明的回声。','科幻·连载','https://placehold.co/140x180/2e86de/ffffff?text=%E9%93%B6%E6%B2%B3',UNIX_TIMESTAMP(),UNIX_TIMESTAMP(),1,0,120300);

SET @book_id := LAST_INSERT_ID();

-- 生成20章示例章节
INSERT INTO `book_chapter` (`book_id`,`title`,`is_vip`,`create_time`,`update_time`) VALUES
(@book_id,'序章·星之序曲',0,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第一章·观测站的异响',0,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第二章·折叠讯号',0,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第三章·跃迁纪录',0,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第四章·黑域边缘',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第五章·回声谱系',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第六章·深空访客',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第七章·遗迹门扉',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第八章·群星之歌',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第九章·恒星密码',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第十章·回声解构',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第十一章·叛离航路',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第十二章·记忆折射',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第十三章·边界之外',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第十四章·文明残痕',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第十五章·星门回廊',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第十六章·共振终端',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第十七章·回声之源',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第十八章·失落之城',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP()),
(@book_id,'第十九章·星海回响',1,UNIX_TIMESTAMP(),UNIX_TIMESTAMP());

-- 将章节ID收集到变量（MySQL不易一次性收集，建议用应用脚本；此处示例手动插入内容时请先查询章节ID）
-- 以下示例假设章节ID从 @chap1 到 @chap20（可通过 SELECT id FROM book_chapter WHERE book_id=@book_id ORDER BY id 获取）

-- 示例正文内容（请按查询到的章节ID替换）
-- 替换下面的 CHAP_ID_X 为实际章节ID
INSERT INTO `book_content` (`book_id`,`chapter_id`,`content`,`create_time`,`update_time`) VALUES
(@book_id, CHAP_ID_1,  '观测站的夜晚总是很长，宇宙微波在墙壁间回荡。我们第一次听见“回声”的那一刻，仿佛群星在耳畔低语。', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(@book_id, CHAP_ID_2,  '折叠讯号来自黑域边缘。它像一条被拉伸的河流，在时空的弯曲里留下涟漪。', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(@book_id, CHAP_ID_3,  '跃迁纪录记录了飞船每一次呼吸。日志里写着：“不要停下，回声在前方。”', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(@book_id, CHAP_ID_4,  '黑域不是黑暗，而是光的另一种排列。我们在那儿看见了文明的影子。', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(@book_id, CHAP_ID_5,  '回声谱系像星图，又像一首歌。每个频点对应一段历史。', UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

-- 其余章节可复制该格式补全正文内容
