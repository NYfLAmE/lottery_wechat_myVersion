DROP TABLE IF EXISTS `prize`;
CREATE TABLE `prize`
(
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品名称',
    `pic` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品图片',
    `link` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品链接',
    `type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '奖品类型，1-虚拟币，2-虚拟券，3-实物小奖，4-实物大奖',
    `data` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品数据',
    `total` int(11) NOT NULL DEFAULT '-1' COMMENT '奖品数量，0 无限量，>0限量，<0无奖品',
    `left` int(11) NOT NULL DEFAULT '0' COMMENT '剩余数量',
    `is_use` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '是否使用中，1-使用中，2-未使用',
    `probability` int(11) NOT NULL DEFAULT '0' COMMENT '中奖概率，万分之n',
    `probability_max` int(11) NOT NULL DEFAULT '0' COMMENT '中奖概率上限',
    `probability_min` int(11) NOT NULL DEFAULT '0' COMMENT '中奖概率下限',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='奖品表';