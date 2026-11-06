//字母异位词分组
//基本思路：首先构造一个map 这个map的key一定是这个词每个字符按升序排序的那个特定的词，然后value就是升序排序后是这个特定词的一个silce。表明这个slice中的词包含这几个特定词。也就是字母异位词。


func groupAnagrams(strs []string) [][]string {
	
    var map1 =  make(map[string][]string) //创建一个map ，key就是我说的那个升序特定词，value就是这个词中有这个字母的词就可以放在这个对应的slice里面
    for _ , value := range strs {	//遍历这个字符串数组
        s := []byte(value)		//新东西！！！这个是一个类型转换 将字符串类型按字节转换成一个slice。因为这里全是英文字母所以可以用byte 如果有中文字符的话 默认是utf-8编码，也就是说是四个字节 得用[]rune
		   sort.Slice(s , func(i , j int) bool {	 //这个也是新东西，就是一个排序函数，用这个函数 有两个参数，第一个参数表明被排序的变量 一般是一个slice 然后第二个变量就是一个函数 这里就是用来定义排序的规则的。	     
            return s[i] < s[j]   			//升序
        })
        key := string(s)					//还是类型转换 把字符slice转回string
        map1[key] = append(map1[key] , value)//把这个 append到这个map	
    }
    res := make([][]string , 0 , len(map1))	//因为函数的return值是一个[][]string所以需要构造一个 然后存。
    for _ , v := range map1 {
        res = append(res , v)
    }
    return res
}