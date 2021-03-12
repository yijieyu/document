## B+tree

InnoDB 存储引擎就是使用B+tree实现其索引结构
在B+tree中，素有数据记录节点都是按照键值大小顺序存放在同一层的叶子结点上，
而非叶子结点上只存储key值信息，这样可以大大加大每个节点存储的key值数量，降低B+Tree的高度


B+tree的特点
    - B+节点关键字搜索采用闭合区间
    - B+非叶节点不保存数据相关信息，只保存关键字和子节点的引用
    - B+关键字对应的数据保存在叶子节点中
    - B+ 叶子节点是顺序排列的，并且相邻节点觉有顺序引用的关系
    

B+tree 索引类型
聚集索引和辅助索引

聚集索引的B+tree 中的叶子节点存放的是整张表的行记录数据
辅助索引与聚集索引的区别在于辅助索引的叶子节点并不包含行记录的全部数据，而是存储相应行数据的聚集索引值，即主键
当通过辅助索引来查询数据时，innodb存储引擎会遍历辅助索引找到主键，然后再通过主键在聚集索引中找到完整的行记录数据


创建表时 默认使用Innodb引擎

B+Tree 在 Myisam 中的体现：
在创建好表结构并且指定搜索引擎为Myisam之后， 会在数据目录生成3个文件，分别是table_name.frm(表结构文件)
table_name.MYD(数据保存文件)，table_name.MYI(索引保存文件)

B+tree在InnoDB中的体现：
在创建好表结构并且指定搜素引擎为InnoDB之后，会在数据目录生成3个文件，分别是table_name.frm(表结构)
table_name.idb(数据与索引保存文件)

在InnoDB中，因为设计指出就是认为主键是非常重要的。是以主键为索引来组织数据的存储的，
当我们没有显示的建立主键索引的时候，搜索引擎会隐式的为我们创建一个主键索引来组织数据存储。
数据库表行中数据的物理顺序与键值的逻辑(索引)顺序相同，InnoDB就是以聚集索引来组织数据的存储的，
在叶子结点上，保存了数据的所有信息。如果这个时候建立了Name


B+Tree 在 InnoDB 中的体现：

在创建好表结构并且指定搜索引擎为 InnoDB之后，会在数据目录生成3个文件，分别是table_name.frm(表结构文件)，
table_name.idb（数据与索引保存文件）。

在 InnoDB中，因为设计之初就是认为主键是非常重要的。是以主键为索引来组织数据的存储，
当我们没有显示的建立主键索引的时候，搜索引擎会隐式的为我们建立一个主键索引以组织数据存储。
数据库表行中数据的物理顺序与键值的逻辑（索引）顺序相同，InnoDB就是以聚集索引来组织数据的存储的，
在叶子节点上，保存了数据的所有信息。如果这个时候建立了name字段的索引：






有一道MySQL的面试题，为什么MySQL的索引要使用B+树而不是其它树形结构?比如B树

因为B树不管叶子节点还是非叶子节点，都会保存数据，这样导致在非叶子节点中能保存的指针数量变少
（有些资料也称为扇出），指针少的情况下要保存大量数据，只能增加树的高度，导致IO操作变多，查询性能变低；











