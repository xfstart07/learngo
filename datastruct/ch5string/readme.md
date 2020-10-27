# 字符串

## 抽象数据类型

ADT串( string) 

Data 
    串中元素仅由一个字符组成,相邻元素具有前驱和后继关系。
    
Operation 

* Strassign(T,* chars):生成一个其值等于字符串常量 chars的串T。
* strcopy(T,S):串S存在,由串S复制得串T。
* Clearstring(s):串S存在,将串清空。
* Stringempty(s):若串S为空,返回true,否则返回 false。
* Strength(s):返回串S的元素个数,即串的长度。
* strcompare(S,T):若S>T,返回值>0,若S=T,返回0,若S<T,返回值<0。
* Concat(T,sS1,S2):用T返回由S1和S2联接而成的新串。
* Substring(Sub,S,pos,len):
    串S存在,1≤pos≤ Strength(s), 且0≤1en≤ strength(s)-pos+1,用Sub返回串S的第pos个字符起长度为1en的子串。
* Index(S,T,pos):串S和T存在,T是非空串,1≤pos≤ Strength(s)。
若主串S中存在和串T值相同的子串,则返回它在主串S中第ps个宇特之后第一次出现的位置,否则返回0。
* Replace(S,T,V):串S、T和V存在,T是非空串。用V替換主串S中出现的所有与T相等的不重叠的子串。
* Strinsert(S,pos,T):串S和存在,1≤pos≤ strength(S)+1。
在串S的第pos个字符之前插入串T。
* Strdelete(S,pos,1en):串S存在,1≤pos≤ strength(s)-1en+1a 从串S中别除第pos个字符起长度为1en的子串。

## 串的顺序存储结构

串的顺序存储结构是用一组地址连续的存储单元来存储串中的字符序列。

## KMP 模式匹配算法

next 数组，最长公共前后缀。

# 资源

* KMP 模式匹配算法讲解: https://www.youtube.com/watch?v=dgPabAsTFa8
* KMP 模式匹配算法讲解2: https://www.youtube.com/watch?v=3IFxpozBs2I
* KMP 模式匹配算法：http://wiki.jikexueyuan.com/project/kmp-algorithm/define.html
