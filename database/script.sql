use `midnight`;

SET @MSG = 1, @CHOOSE = 2, @NEWS = 3, @MOMENT = 4, @ONLINE = 5;
SET @OFFLINE = 6, @GAMEOVER_SUCC = 7, @GAMEOVER_FAILED = 8;

/* 故事进度 */
REPLACE INTO `story`
(`sid`, `conjunction_msg_type`, `conjunction_mid`, `time_delay`, `next_sid`)
VALUES 
(1, @ONLINE, -1, 3, 2),
(2, @MSG, 1, 5, 3),
(3, @MSG, 2, 1, 4),
(4, @CHOOSE, 1, 10, -1),
(5, @MSG, 3, 10, 6),
(6, @MSG, 4, 1, 7),
(7, @CHOOSE, 2, 10, -1),
(8, @MSG, 5, 1, 9),
(9, @CHOOSE, 3, 10, -1),
(10, @MSG, 6, 5, 11),
(11, @MSG, 7, 7, 12),
(12, @MSG, 8, 2, 13),
(13, @OFFLINE, -1, 70, 14),
(14, @ONLINE, -1, 5, 15),
(15, @MSG, 9, 5, 16),
(16, @CHOOSE, 4, 8, -1),
(17, @MSG, 10, 1, 22),
(18, @MSG, 11, 1, 19),
(19, @CHOOSE, 5, 3, -1),
(20, @OFFLINE, -1, 1, 21),
(21, @GAMEOVER_FAILED, -1, 0, 0),
(22, @NEWS, 1, 3, 23),
(23, @MSG, 12, 1, 24),
(24, @CHOOSE, 6, 4, -1),
(25, @MSG, 13, 120, 26),
(26, @MSG, 14, 1, 27),
(27, @CHOOSE, 7, 5, -1),
(28, @CHOOSE, 8, 5, -1),
(29, @OFFLINE, -1, 1, 30),
(30, @GAMEOVER_FAILED, -1, 0, 0),
(31, @MSG, 15, 120, 32),
(32, @MSG, 16, 10, 33),
(33, @MSG, 17, 1, 34),
(34, @CHOOSE, 9, 5, -1),
(35, @MSG, 18, 7, 37),
(36, @MSG, 19, 7, 37),
(37, @MSG, 20, 1, 38),
(38, @NEWS, 2, 5, 39),
(39, @MSG, 21, 6, 40),
(40, @MSG, 22, 1, 41),
(41, @CHOOSE, 10, 5, -1),
(42, @OFFLINE, -1, 3, 43),
(43, @CHOOSE, 11, 1200, -1),
(44, @ONLINE, -1, 3, 45),
(45, @MSG, 23, 5, 46),
(46, @CHOOSE, 12, 10, -1),
(47, @MSG, 24, 15, 48),
(48, @MSG, 25, 5, 49),
(49, @CHOOSE, 13, 10, -1),
(50, @MSG, 26, 7, 51),
(51, @MSG, 27, 1, 52),
(52, @CHOOSE, 14, 5, -1),
(53, @MSG, 28, 2, 54),
(54, @CHOOSE, 15, 8, -1),
(55, @MSG, 29, 5, 56),
(56, @MSG, 30, 3, 57),
(57, @CHOOSE, 16, 5, -1),
(58, @MSG, 31, 300, 59),
(59, @MSG, 32, 300, 60),
(60, @CHOOSE, 17, 10, -1),
(61, @MSG, 33, 5, 62),
(62, @CHOOSE, 18, 15, -1),
(63, @MSG, 34, 5, 64),
(64, @MSG, 35, 2, 65),
(65, @CHOOSE, 19, 15, -1),
(66, @MSG, 36, 5, 67),
(67, @MSG, 37, 3, 68),
(68, @CHOOSE, 20, 5, -1),
(69, @MSG, 38, 360, 70),
(70, @MSG, 39, 5, 71),
(71, @MSG, 40, 3, 72),
(72, @CHOOSE, 21, 300, -1),
(73, @OFFLINE, -1, 28800, 74),
(74, @ONLINE, -1, 3, 75),
(75, @MSG, 41, 3, 76),
(76, @GAMEOVER_SUCC, -1, 0, 0);
/* (Sid, MsgType, ConID, DelayTime, NextSid), */

/* 对方的 msg */
REPLACE INTO `msg_text` (`mid`, `content`)
VALUES
(1, '请问，你看得到吗'),
(2, '拜托了，有没有人能帮帮我'),
(3, '我也不知道，我...我什么都不记得了'),
(4, '我不知道我在什么地方，这里好可怕，刚刚…刚刚街上有人被杀了'),
(5, '是真的，就在街上，我…我都听见枪声了，还有很多人看见了，但没有一个人管'),
(6, '你是不是不相信我，我真的没有骗你，，，'),
(7, '真的，街上的人明明都看见了，但他们只是在围观，还有人在笑'),
(8, '还有枪声，枪声越来越近了'),
(9, '我真的好害怕，我现在就在一片小树林里躲着'),
(10, '周围好像有个酒馆，看上去有点年份，感觉……还是有点熟悉的'),
(11, '你可不可以正经一点，好吧抱歉，我不该打扰你的'),
(12, '不好 远处有人走过来了，我该怎么办？'),
(13, '我相信你，我尽量小心'),
(14, '就wwwww'),
(15, '我相信你，我找个地方藏起来'),
(16, '我躲进了灌木丛里，希望他们看不到我'),
(17, '天，我看到一队穿着防化服衣服的人，他们，他们正拖着一具尸体'),
(18, '我也不知道 我弄不懂这个 为什么出这种事情 我真的要崩溃了'),
(19, '今天到底怎么了，为什么出这种事情，我真的要崩溃了'),
(20, '不，不是尸体，那个人还在动'),
(21, '死了，他死了 被那些人一枪打死的'),
(22, '感觉有人跟我说话，这个声音，很熟悉，很熟悉……'),
(23, '我是不是晕过去了……'),
(24, '谢谢你，我没事'),
(25, '那些穿着防化服的人不见了，尸体也不见了，就像什么都没发生过一样，没有痕迹'),
(26, '总感觉这个地方有点熟悉，感觉之前来过，可是为什么我什么都记不起来'),
(27, '芯片 迷宫 Adam 巢穴 教堂'),
(28, '脑子里突然出现的声音，隐约好像听到了这些词语，就记录下来了'),
(29, '完全看不出有什么意义，不过那个声音倒是很温暖的感觉'),
(30, '出了这个树林，过了马路有个房子，感觉很熟悉，我想去看一下'),
(31, '嗯嗯，我一定小心'),
(32, '敲了门，没有人，不过门没有上锁，我进去看看'),
(33, '等等，我看到了我自己的照片'),
(34, '照片上另一个人好像是……我父亲，这会不会是我家'),
(35, '我的照片，还有，我的日记，这真的是我家！'),
(36, '说得对，我还是得谨慎一点，不过即便不是我自己的家，也应该是某个和我和爸爸关系不错的熟人的住处吧'),
(37, '唉，经历了这么可怕的事情，而且现在还是什么都记不清，真的不知道该怎么办'),
(38, '嗯嗯，我懂，谢谢你'),
(39, '救命！！！！一个穿黑色衣服的男人踢开门闯了进来'),
(40, '我现在躲在床底下，我好害怕'),
(41, '那个，，，请问你是？');


/* 我的 choose */
REPLACE INTO `msg_choose` (`mid`, `l_content`, `l_next_sid`, `r_content`, `r_next_sid`)
VALUES
(1, '你是谁啊', 5, '什么鬼软件，不经过我同意就加好友？', 5),
(2, '怎么会发生这种事，就在街上? ', 8, '那我觉得你可得报警了啊', 8),
(3, '那小妹妹要小心哦', 10, '怎么可能有这种事情，光天化日下？', 10),
(4, '别慌，告诉我你周围是什么样的', 17, '小树林？！有意思', 18),
(5, '好了我相信你了，告诉我你那边的一些情况', 17, '再见了您嘞', 20),
(6, '赶快找个地方先躲着', 31, '赶快逃跑，千万小心别被抓到', 25),
(7, '发生了什么？', 28, '发生了什么？', 28),
(8, '喂，你别吓我！', 29, '出什么事了，快回答我！', 29),
(9, '生化人？这是什么地方', 35, '什么？拖着尸体？', 36),
(10, '什么声音 你还好吗？', 42, '先别管声音了，那些穿防化服的人走了吗', 42),
(11, '你怎么了', 44, '喂，怎么下线了，别吓我啊', 44),
(12, '你终于上线了，还以为你发生了什么危险', 47, '谢天谢地你没出事', 47),
(13, '没有痕迹？这么快就被清理了？不会是在什么杀人组织的地盘吧', 50, '真奇怪，这到底是什么地方？而且既然无所约束的杀人，为什么要清理痕迹呢？', 50),
(14, '这些词语什么意思？', 53, '你想表达什么？', 53),
(15, '哦哦这样，看上去没有什么联系的词语', 55, '感觉像是暗号', 55),
(16, '小心点，可能那些人还在这里', 58, '这太诡异了，真的安全吗，如果去的话一定小心些', 58),
(17, '别冒险，快出来', 61, '随便进别人的地方不太好吧', 61),
(18, '照片？有可能你曾经来过这里，试试看能不能想起来', 63, '还有其它线索吗？', 63),
(19, '真的是你家吗，终于有个好消息了', 67, '别太早下结论，你的记忆并不靠谱，可能是圈套', 66),
(20, '没关系，有我在', 69, '也不知道该怎么才能让你好受，但是……', 69),
(21, '什么情况，你千万别慌，好好躲着', 73, '别害怕，看看能不能找机会跑出去', 73);

REPLACE INTO `msg_news` (`mid`, `title`, `content`)
VALUES 
(1, '国际油价持续暴跌 俄罗斯经济危机还有多久', 
'说到俄罗斯经济，就不得不提原油。原油价格与俄罗斯经济紧密相连。进入21世纪，俄罗斯人民的生活好了起来，这就是所谓的“石油福利”。领导人将俄罗斯的石油收入变为国民福利发放给老百姓，但是这种福利有一个致命的缺陷：国际油价高的时候，老百姓过的很滋润;可是油价大跌的时候，老百姓的生活甚至可以说是拮据。加之美国和欧盟对俄的经济制裁，俄罗斯的经济不容乐观。因卢布下跌，人们开始兑换外汇，加速了货币贬值速度。'),
(2, '图灵测试通过之后，下一步是？', 
'从第一次“深蓝”战胜人类，到“Alpha Go”大胜人类，再到Stephen教授研发的智能计算机首次通过图灵测试，人工智能的发展前所未有的迅速，对于人工智能是否属于生命，是否应该享有人权，社会也已经争论很久。  反对者认为，目前的人工智能并不是生命，因为它只是一个很好的模仿者，模仿了一个能够思考的人，却和人有本质的区别：它没有意识。  计算机科学家近期的研究表明，现在已经有很多相关算法，如已故人工智能专家Stephen提出的“二分心智”理论，就有可能诱导计算机产生意识。但此理论目前争议较大。');

